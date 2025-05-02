package main

import (
	"context"
	"encoding/json"
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
}

// HelloWorld is an HTTP Cloud Function
func HelloWorld(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	var err error
	defer func() {
		monitoring.LogFunctionExecution("HelloWorld", startTime, err)
	}()

	// Log the request
	log.Printf("Function triggered by request: %v", r.URL.Path)

	// Get project ID from environment
	projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")
	if projectID == "" {
		projectID = "calm-cab-458210-t2" // fallback to default
	}

	// Initialize monitoring client
	ctx := context.Background()
	metricsClient, err := monitoring.NewMetricsClient(ctx, projectID)
	if err != nil {
		errors.WriteError(w, http.StatusInternalServerError, "Failed to initialize monitoring", err.Error())
		return
	}

	// Create response
	response := Response{
		Message: "Hello from CarbonQuest!",
		Version: "1.0.0",
		Project: projectID,
		Status:  "success",
		Time:    time.Now().Format(time.RFC3339),
	}

	// Set headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Write response
	if err = json.NewEncoder(w).Encode(response); err != nil {
		errors.WriteError(w, http.StatusInternalServerError, "Failed to encode response", err.Error())
		return
	}

	// Record metrics
	if err = metricsClient.RecordLatency(ctx, "HelloWorld", time.Since(startTime)); err != nil {
		log.Printf("Failed to record metrics: %v", err)
	}

	log.Printf("Function executed successfully for project: %s", projectID)
}
