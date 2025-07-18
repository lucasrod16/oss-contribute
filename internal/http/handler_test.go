package http

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/lucasrod16/oss-projects/internal/cache"
	"github.com/stretchr/testify/require"
)

func TestGetRepos(t *testing.T) {
	tests := []struct {
		name           string
		method         string
		expectedStatus int
		expectedBody   string
		cacheData      []byte
	}{
		{
			name:           "valid request",
			method:         http.MethodGet,
			expectedStatus: http.StatusOK,
			expectedBody:   `{"data": "some data"}`,
			cacheData:      []byte(`{"data": "some data"}`),
		},
		{
			name:           "method not allowed",
			method:         http.MethodPost,
			expectedStatus: http.StatusMethodNotAllowed,
			expectedBody:   http.StatusText(http.StatusMethodNotAllowed) + "\n",
		},
		{
			name:           "no data found in cache",
			method:         http.MethodGet,
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   "No data found in cache\n",
			cacheData:      nil,
		},
		{
			name:           "HEAD request should not return body",
			method:         http.MethodHead,
			expectedStatus: http.StatusOK,
			expectedBody:   "",
			cacheData:      []byte(`{"data": "some data"}`),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := cache.New()
			c.Set(tt.cacheData)

			req := httptest.NewRequest(tt.method, "/repos", nil)
			rr := httptest.NewRecorder()

			GetRepos(c).ServeHTTP(rr, req)

			require.Equal(t, tt.expectedStatus, rr.Code, "GetRepos handler returned wrong status code")
			require.Equal(t, tt.expectedBody, rr.Body.String(), "GetRepos handler returned unexpected body")
		})
	}
}
