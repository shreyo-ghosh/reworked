package errors

import (
	"encoding/json"
	"net/http"
)

// ErrorResponse represents an error response
type ErrorResponse struct {
	Error   string `json:"error"`
	Details string `json:"details,omitempty"`
}

// WriteError writes an error response to the HTTP response writer
func WriteError(w http.ResponseWriter, statusCode int, message string, details string) {
	response := ErrorResponse{
		Error:   message,
		Details: details,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}
