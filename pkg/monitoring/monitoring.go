package monitoring

import (
	"context"
	"log"
	"time"
)

// MetricsClient represents a client for recording metrics
type MetricsClient struct {
	projectID string
}

// NewMetricsClient creates a new metrics client
func NewMetricsClient(ctx context.Context, projectID string) (*MetricsClient, error) {
	return &MetricsClient{
		projectID: projectID,
	}, nil
}

// LogFunctionExecution logs the execution of a function
func LogFunctionExecution(functionName string, startTime time.Time, err error) {
	if err != nil {
		log.Printf("Function %s failed after %v: %v", functionName, time.Since(startTime), err)
	} else {
		log.Printf("Function %s completed in %v", functionName, time.Since(startTime))
	}
}

// RecordLatency records the latency of a function call
func (c *MetricsClient) RecordLatency(ctx context.Context, functionName string, duration time.Duration) error {
	log.Printf("Latency for %s: %v", functionName, duration)
	return nil
}

// LogError logs an error with monitoring
func LogError(err error) {
	log.Printf("Error: %v", err)
}

// LogInfo logs an info message with monitoring
func LogInfo(message string) {
	log.Printf("Info: %s", message)
}
