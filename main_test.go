package main

import (
	"os"
	"testing"
)

// Helper function to reset environment variables after each test
func resetEnv() {
	os.Unsetenv("APP_PORT")
	os.Unsetenv("LOG_LEVEL")
}

// TestConfigResolution tests the configuration priority logic (CLI > ENV > Default).
func TestConfigResolution(t *testing.T) {
	// Test cases for different priority combinations
	tests := []struct {
		name     string
		cliPort  string
		cliLog   string
		envPort  string
		envLog   string
		expected Config
	}{
		{
			name:    "P3_DefaultValues",
			cliPort: "", cliLog: "",
			envPort: "", envLog: "",
			expected: Config{Port: "8080", LogLevel: "info"},
		},
		{
			name:    "P2_EnvironmentVariables",
			cliPort: "", cliLog: "",
			envPort: "9000", envLog: "debug",
			expected: Config{Port: "9000", LogLevel: "debug"},
		},
		{
			name:    "P1_CommandLineOverridesAll",
			cliPort: "8000", cliLog: "warn",
			envPort: "9000", envLog: "debug", // Should be ignored
			expected: Config{Port: "8000", LogLevel: "warn"},
		},
		{
			name:    "MixedPriority_CLI_Port_ENV_Log",
			cliPort: "7000", cliLog: "", // CLI Port overrides, Log is empty
			envPort: "9000", envLog: "trace",
			expected: Config{Port: "7000", LogLevel: "trace"},
		},
		{
			name:    "MixedPriority_ENV_Port_CLI_Log",
			cliPort: "", cliLog: "error", // CLI Log overrides, Port is empty
			envPort: "9000", envLog: "debug",
			expected: Config{Port: "9000", LogLevel: "error"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resetEnv()

			// Set environment variables for the current test case
			if tt.envPort != "" {
				os.Setenv("APP_PORT", tt.envPort)
			}
			if tt.envLog != "" {
				os.Setenv("LOG_LEVEL", tt.envLog)
			}

			// Resolve the configuration
			actual := resolveConfig(tt.cliPort, tt.cliLog)

			// Check Port
			if actual.Port != tt.expected.Port {
				t.Errorf("Port mismatch. Expected: %s, Got: %s", tt.expected.Port, actual.Port)
			}
			// Check LogLevel
			if actual.LogLevel != tt.expected.LogLevel {
				t.Errorf("LogLevel mismatch. Expected: %s, Got: %s", tt.expected.LogLevel, actual.LogLevel)
			}
		})
	}
}
