package http

import (
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"

	"github.com/lucasrod16/oss-contribute/internal/cache"
	"github.com/stretchr/testify/require"
)

func TestRateLimiter(t *testing.T) {
	rl := NewRateLimiter()
	c := cache.New()
	c.Set([]byte(`{"data": "some data"}`))

	req := httptest.NewRequest(http.MethodGet, "/repos", nil)
	req.Header.Set("X-Forwarded-For", "192.168.1.1")

	var wg sync.WaitGroup
	var mu sync.Mutex
	successCount := 0
	failCount := 0

	// send 20 requests concurrently to trigger the rate limit
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			rr := httptest.NewRecorder()
			rl.Limit(GetRepos(c)).ServeHTTP(rr, req)

			if rr.Code == http.StatusOK {
				mu.Lock()
				successCount++
				mu.Unlock()
			} else if rr.Code == http.StatusTooManyRequests {
				mu.Lock()
				failCount++
				mu.Unlock()
			}
		}()
	}
	wg.Wait()

	require.Equal(t, 10, successCount, "Expected 10 successful requests")
	require.Equal(t, 10, failCount, "Expected 10 failed requests")
}
