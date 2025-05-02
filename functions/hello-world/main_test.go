package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloWorld(t *testing.T) {
	// Set up test environment
	os.Setenv("GOOGLE_CLOUD_PROJECT", "test-project")
	defer os.Unsetenv("GOOGLE_CLOUD_PROJECT")

	// Create test request
	req := httptest.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()

	// Call the function
	HelloWorld(rr, req)

	// Check status code and headers
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))

	// Parse response
	var response Response
	err := json.NewDecoder(rr.Body).Decode(&response)
	assert.NoError(t, err)

	// Verify response fields
	assert.Equal(t, "Hello from CarbonQuest!", response.Message)
	assert.Equal(t, "1.0.0", response.Version)
	assert.Equal(t, "test-project", response.Project)
}
