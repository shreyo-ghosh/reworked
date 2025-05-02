package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
	// This is a basic test that verifies the function doesn't panic
	// In a real implementation, we would mock the GCP SDK calls
	err := deployFunction("test-function", "dev", "1.0.0")
	assert.NoError(t, err)
}
