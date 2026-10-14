package calculator

import (
	"encoding/json"
	"net/http"

	"rest-api-template/internal/rest/response"
	pgstore "rest-api-template/internal/store/pg"

	"github.com/rs/zerolog"
)

// MultiplyRequest is the JSON request body for POST /calculator/v1/multiply.
type MultiplyRequest struct {
	Value1 float64 `json:"value1"`
	Value2 float64 `json:"value2"`
}

// MultiplyResponse is the JSON response body for a successful multiply operation.
type MultiplyResponse struct {
	TransactionUUID string  `json:"transaction-uuid"`
	Value1          float64 `json:"value1"`
	Value2          float64 `json:"value2"`
	Operation       string  `json:"operation"`
	Result          float64 `json:"result"`
}

// Multiply returns the HTTP handler for POST /calculator/v1/multiply.
//
// HTTP status codes:
//   - 200 — success
//   - 400 — value1 or value2 is outside the range [-1000, 1000]
//   - 401 — missing or wrong api-key (enforced by upstream middleware)
//   - 500 — any unhandled storage error
func Multiply(logger zerolog.Logger, store *pgstore.CalculatorTxStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req MultiplyRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			response.WriteErrorWithEnvelope(w, http.StatusBadRequest,
				"INVALID_REQUEST_BODY", "invalid request body")
			return
		}

		if req.Value1 < -1000 || req.Value1 > 1000 {
			response.WriteErrorWithEnvelope(w, http.StatusBadRequest,
				"VALUE1_OUT_OF_RANGE", "value1 must be in range [-1000, 1000]")
			return
		}
		if req.Value2 < -1000 || req.Value2 > 1000 {
			response.WriteErrorWithEnvelope(w, http.StatusBadRequest,
				"VALUE2_OUT_OF_RANGE", "value2 must be in range [-1000, 1000]")
			return
		}

		result := req.Value1 * req.Value2

		txUUID, err := store.Create(r.Context(), pgstore.CreateTxParams{
			Value1:    req.Value1,
			Value2:    req.Value2,
			Operation: "MULTIPLY",
			Result:    result,
		})
		if err != nil {
			logger.Error().Err(err).Msg("failed to persist calculator transaction")
			response.WriteErrorWithEnvelope(w, http.StatusInternalServerError,
				"INTERNAL_SERVER_ERROR", "internal server error")
			return
		}

		response.WriteJSONWithEnvelope(w, http.StatusOK, MultiplyResponse{
			TransactionUUID: txUUID.String(),
			Value1:          req.Value1,
			Value2:          req.Value2,
			Operation:       "MULTIPLY",
			Result:          result,
		})
	}
}
