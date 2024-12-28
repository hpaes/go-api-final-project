package response

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	statusCode int
	Error      string `json:"error"`
}

func NewErrorResponse(statusCode int, err error) *ErrorResponse {
	return &ErrorResponse{
		statusCode: statusCode,
		Error:      err.Error(),
	}
}

func (er *ErrorResponse) Send(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(er.statusCode)
	if err := json.NewEncoder(w).Encode(er); err != nil {
		return
	}
}
