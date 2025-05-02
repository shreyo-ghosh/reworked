// Package main contains a Google Cloud Function.
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"carbonquest/pkg/errors"
	"carbonquest/pkg/monitoring"
)

// Response is the structure for our HTTP response
type Response struct {
	Message string `json:"message"`
	Version string `json:"version"`
	Project string `json:"project"`
	Status  string `json:"status"`
	Time    string `json:"time"`
	Region  string `json:"region"`
}

// HelloWorld is an HTTP Cloud Function
func HelloWorld(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	functionName := "HelloWorld"

	// Initialize monitoring
	ctx := r.Context()
	projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")
	if projectID == "" {
		projectID = "calm-cab-458210-t2" // fallback to default
	}

	metricsClient, err := monitoring.NewMetricsClient(ctx, projectID)
	if err != nil {
		log.Printf("Failed to create metrics client: %v", err)
		errors.WriteError(w, http.StatusInternalServerError, "Internal Server Error", "Failed to initialize monitoring")
		return
	}

	// Log the request
	monitoring.LogInfo(fmt.Sprintf("Function triggered by request: %v", r.URL.Path))

	// Create response
	response := Response{
		Message: "Hello from CarbonQuest!",
		Version: "1.0.4",
		Project: projectID,
		Status:  "success",
		Time:    time.Now().Format(time.RFC3339),
		Region:  "us-central1",
	}

	// Set headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Write response
	if err := json.NewEncoder(w).Encode(response); err != nil {
		monitoring.LogError(err)
		errors.WriteError(w, http.StatusInternalServerError, "Internal Server Error", "Failed to encode response")
		return
	}

	// Record metrics
	duration := time.Since(startTime)
	if err := metricsClient.RecordLatency(ctx, functionName, duration); err != nil {
		monitoring.LogError(err)
	}

	monitoring.LogFunctionExecution(functionName, startTime, nil)
}
