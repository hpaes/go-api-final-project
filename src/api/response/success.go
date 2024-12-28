package response

import (
	"encoding/json"
	"net/http"
)

type SuccessResponse struct {
	StatusCode int
	Result     any `json:"result"`
}

func NewSuccessResponse(statusCode int, result any) *SuccessResponse {
	return &SuccessResponse{
		StatusCode: statusCode,
		Result:     result,
	}
}

func (sr *SuccessResponse) Send(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(sr.StatusCode)
	if err := json.NewEncoder(w).Encode(sr.Result); err != nil {
		return
	}
}
