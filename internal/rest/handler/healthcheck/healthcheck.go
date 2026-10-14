package healthcheck

import (
	"encoding/json"
	"net/http"
	"rest-api-template/internal/rest/constant"
)

// HealthcheckResponse is the JSON response body for GET /healthcheck.
type HealthcheckResponse struct {
	Status string `json:"status"`
}

// Healthcheck returns the HTTP handler for GET /healthcheck.
//
// HTTP status codes:
//   - 200 — service is healthy
func Healthcheck() http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", constant.MimeApplicationJSON)
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(HealthcheckResponse{Status: "ok"})
	}
}
