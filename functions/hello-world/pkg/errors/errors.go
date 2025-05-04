package errors

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// ErrorResponse represents an error response
type ErrorResponse struct {
	Error   string `json:"error"`
	Details string `json:"details,omitempty"`
}

// Error represents a wrapped error
type Error struct {
	message string
	cause   error
}

// Wrap wraps an error with a message
func Wrap(err error, message string) error {
	if err == nil {
		return nil
	}
	return &Error{
		message: message,
		cause:   err,
	}
}

// Error returns the error message
func (e *Error) Error() string {
	if e.cause == nil {
		return e.message
	}
	return fmt.Sprintf("%s: %v", e.message, e.cause)
}

// Unwrap returns the underlying error
func (e *Error) Unwrap() error {
	return e.cause
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
