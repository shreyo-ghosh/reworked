package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// MockCommandExecutor is a mock implementation of CommandExecutor
type MockCommandExecutor struct {
	shouldError bool
}

func (m *MockCommandExecutor) Execute(name string, args ...string) ([]byte, error) {
	if m.shouldError {
		return []byte("mock error"), fmt.Errorf("mock error")
	}
	return []byte("mock success"), nil
}

func TestIsValidEnvironment(t *testing.T) {
	tests := []struct {
		name        string
		environment string
		want        bool
	}{
		{"Valid sandbox", "sandbox", true},
		{"Valid dev", "dev", true},
		{"Valid pro", "pro", true},
		{"Invalid empty", "", false},
		{"Invalid random", "random", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isValidEnvironment(tt.environment)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestDeployFunction(t *testing.T) {
	tests := []struct {
		name        string
		function    string
		env         string
		version     string
		shouldError bool
		executor    CommandExecutor
	}{
		{
			name:        "Valid deployment to sandbox",
			function:    "test-function",
			env:         "sandbox",
			version:     "1.0.0",
			shouldError: false,
			executor:    &MockCommandExecutor{shouldError: false},
		},
		{
			name:        "Invalid environment",
			function:    "test-function",
			env:         "invalid",
			version:     "1.0.0",
			shouldError: true,
			executor:    &MockCommandExecutor{shouldError: false},
		},
		{
			name:        "Command execution error",
			function:    "test-function",
			env:         "sandbox",
			version:     "1.0.0",
			shouldError: true,
			executor:    &MockCommandExecutor{shouldError: true},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := deployFunctionWithExecutor(tt.function, tt.env, tt.version, tt.executor)
			if tt.shouldError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
