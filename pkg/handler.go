package osscontribute

import (
	"net/http"
)

func GetRepos(c *Cache) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "405 Method Not Allowed\n", http.StatusMethodNotAllowed)
			return
		}

		data, timestamp := c.Get()
		if data == nil {
			http.Error(w, "No data found in cache\n", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Last-Modified", timestamp.Format(http.TimeFormat))
		w.WriteHeader(http.StatusOK)
		w.Write(data)
		w.Write([]byte("\n"))
	}
}
