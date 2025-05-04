// Package function contains a Google Cloud Function.
package function

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/shreyo-ghosh/carbonquest-assignment/functions/hello-world/pkg/errors"
	"github.com/shreyo-ghosh/carbonquest-assignment/functions/hello-world/pkg/monitoring"
)

// Response represents the structure of our HTTP response
type Response struct {
	Message   string `json:"message"`
	Version   string `json:"version"`
	Time      string `json:"time"`
	Timestamp string `json:"timestamp"`
	Status    string `json:"status"`
}

// HelloWorld is an HTTP Cloud Function that returns a greeting.
func HelloWorld(w http.ResponseWriter, r *http.Request) {
	// Set response headers
	w.Header().Set("Content-Type", "application/json")

	// Initialize monitoring
	monitor := monitoring.NewMonitor()
	defer monitor.Close()

	// Create response
	response := Response{
		Message:   "Hello from Cloud Function!",
		Version:   "1.0.2",
		Time:      time.Now().Format(time.RFC3339),
		Timestamp: fmt.Sprintf("%d", time.Now().Unix()),
		Status:    "success",
	}

	// Log the request
	monitor.LogRequest(r)

	// Encode and send response
	if err := json.NewEncoder(w).Encode(response); err != nil {
		monitor.LogError(errors.Wrap(err, "failed to encode response"))
		http.Error(w, fmt.Sprintf("Error encoding response: %v", err), http.StatusInternalServerError)
		return
	}

	// Log successful response
	monitor.LogResponse(response)
}
