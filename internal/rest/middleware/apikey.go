package middleware

import (
	"net/http"

	"rest-api-template/internal/rest/response"
)

const apiKeyHeader = "api-key"

// APIKey returns a chi-compatible middleware that enforces the api-key header.
//
// If the header is absent or its value does not match expectedKey the middleware
// short-circuits with HTTP 401 and a JSON error body; no downstream handler is
// called.
//
// Usage:
//
//	router.Use(middleware.APIKey(cfg.APIKey))
func APIKey(expectedKey string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			key := r.Header.Get(apiKeyHeader)
			if key == "" || key != expectedKey {
				response.WriteErrorWithEnvelope(w, http.StatusUnauthorized,
					"UNAUTHORIZED", "unauthorized: invalid or missing api-key")
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
