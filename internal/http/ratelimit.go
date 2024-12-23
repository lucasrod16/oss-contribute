package http

import (
	"log"
	"net"
	"net/http"
	"strings"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

type client struct {
	limiter  *rate.Limiter
	lastSeen time.Time
	mu       sync.Mutex
}

// rateLimiter holds rate limiters per client IP address.
type rateLimiter struct {
	mu      sync.Mutex
	clients map[string]*client
}

func NewRateLimiter() *rateLimiter {
	rl := &rateLimiter{
		clients: make(map[string]*client),
	}
	go rl.cleanupStaleClients(10)
	return rl
}

// Limit applies rate limiting to the given HTTP handler.
func (rl *rateLimiter) Limit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cl := rl.getClientLimiter(getClientIP(r))

		if !cl.limiter.Allow() {
			http.Error(w, http.StatusText(http.StatusTooManyRequests), http.StatusTooManyRequests)
			return
		}

		cl.mu.Lock()
		cl.lastSeen = time.Now()
		cl.mu.Unlock()

		next.ServeHTTP(w, r)
	})
}

func (rl *rateLimiter) getClientLimiter(ip string) *client {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	cl, exists := rl.clients[ip]
	if !exists {
		cl = &client{
			limiter:  rate.NewLimiter(5, 10),
			lastSeen: time.Now(),
		}
		rl.clients[ip] = cl
	}
	return cl
}

// cleanupStaleClients removes clients that haven't requested in the last specified duration to conserve memory.
func (rl *rateLimiter) cleanupStaleClients(minutes time.Duration) {
	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()

	for range ticker.C {
		rl.mu.Lock()
		for ip, cl := range rl.clients {
			if time.Since(cl.lastSeen) > minutes*time.Minute {
				delete(rl.clients, ip)
			}
		}
		rl.mu.Unlock()
	}
}

// getClientIP extracts the client's IP address from the request.
func getClientIP(r *http.Request) string {
	// handle "X-Forwarded-For" header used by proxies and load balancers.
	xff := r.Header.Get("X-Forwarded-For")
	if xff != "" {
		ips := strings.Split(xff, ",")
		return strings.TrimSpace(ips[0])
	}

	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		log.Printf("could not determine client IP: %v\n", err)
		return ""
	}
	return ip
}
