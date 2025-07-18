package http

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"

	"github.com/lucasrod16/oss-projects/internal/cache"
	"github.com/stretchr/testify/require"
	"golang.org/x/sync/errgroup"
)

func TestRateLimiter(t *testing.T) {
	rl := NewRateLimiter()
	c := cache.New()
	c.Set([]byte(`{"data": "some data"}`))

	req := httptest.NewRequest(http.MethodGet, "/repos", nil)
	req.Header.Set("X-Forwarded-For", "192.168.1.1")

	var g errgroup.Group
	var mu sync.Mutex

	successCount := 0
	failCount := 0

	// send 20 requests concurrently to trigger the rate limit
	for range 20 {
		g.Go(func() error {
			rr := httptest.NewRecorder()
			rl.Limit(GetRepos(c)).ServeHTTP(rr, req)

			mu.Lock()
			defer mu.Unlock()

			switch rr.Code {
			case http.StatusOK:
				successCount++
			case http.StatusTooManyRequests:
				failCount++
			default:
				return fmt.Errorf("unexpected status code: %d", rr.Code)
			}
			return nil
		})
	}
	require.NoError(t, g.Wait())
	require.Equal(t, 10, successCount, "Expected 10 successful requests")
	require.Equal(t, 10, failCount, "Expected 10 failed requests")
}
