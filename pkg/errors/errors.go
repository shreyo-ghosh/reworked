package errors

import (
	"encoding/json"
	"net/http"
)

// ErrorResponse represents a standardized error response
type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Details string `json:"details,omitempty"`
}

// WriteError writes a standardized error response
func WriteError(w http.ResponseWriter, code int, message string, details string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	response := ErrorResponse{
		Code:    code,
		Message: message,
		Details: details,
	}

	json.NewEncoder(w).Encode(response)
}

// HandleError handles common error cases
func HandleError(w http.ResponseWriter, err error) {
	switch err {
	case nil:
		return
	default:
		WriteError(w, http.StatusInternalServerError, "Internal Server Error", err.Error())
	}
}
