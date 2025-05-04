package monitoring

import (
	"log"
	"net/http"
	"time"
)

// Monitor represents a monitoring instance
type Monitor struct {
	startTime time.Time
}

// NewMonitor creates a new monitoring instance
func NewMonitor() *Monitor {
	return &Monitor{
		startTime: time.Now(),
	}
}

// LogRequest logs the incoming request
func (m *Monitor) LogRequest(r *http.Request) {
	log.Printf("Request received: %s %s", r.Method, r.URL.Path)
}

// LogResponse logs the outgoing response
func (m *Monitor) LogResponse(response interface{}) {
	log.Printf("Response sent: %+v", response)
}

// LogError logs an error
func (m *Monitor) LogError(err error) {
	log.Printf("Error occurred: %v", err)
}

// Close closes the monitor and logs the total request time
func (m *Monitor) Close() {
	duration := time.Since(m.startTime)
	log.Printf("Request completed in %v", duration)
}
