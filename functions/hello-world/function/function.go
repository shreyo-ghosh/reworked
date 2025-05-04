// Package function contains a Google Cloud Function.
package function

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Response represents the structure of our HTTP response
type Response struct {
	Message string `json:"message"`
	Version string `json:"version"`
	Time    string `json:"time"`
}

// HelloWorld is an HTTP Cloud Function that returns a greeting.
func HelloWorld(w http.ResponseWriter, r *http.Request) {
	// Set response headers
	w.Header().Set("Content-Type", "application/json")

	// Create response
	response := Response{
		Message: "Hello from Cloud Function!",
		Version: "1.0.0",
		Time:    time.Now().Format(time.RFC3339),
	}

	// Encode and send response
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, fmt.Sprintf("Error encoding response: %v", err), http.StatusInternalServerError)
		return
	}
}
