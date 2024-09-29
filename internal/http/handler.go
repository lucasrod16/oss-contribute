package http

import (
	"net/http"

	"github.com/lucasrod16/oss-contribute/internal/cache"
)

func GetRepos(c *cache.Cache) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
			return
		}

		data, timestamp := c.Get()
		if data == nil {
			http.Error(w, "No data found in cache", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Last-Modified", timestamp.Format(http.TimeFormat))
		w.WriteHeader(http.StatusOK)
		w.Write(data)
	}
}
