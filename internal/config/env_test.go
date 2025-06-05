package config

import (
	"os"
	"testing"
)

// dummy .env content
const dummyEnv = `
FOO=bar
BAZ=qux
`

func writeEnvFile(name string) {
	os.WriteFile(name, []byte(dummyEnv), 0644)
}

func removeEnvFile(name string) {
	os.Remove(name)
}

func TestLoadEnv(t *testing.T) {
	tests := []struct {
		name          string
		envValue      string
		expectFoo     string
		loadDotEnv    bool
		expectToLoad  bool
	}{
		{
			name:         "development loads .env",
			envValue:     "development",
			expectFoo:    "bar",
			loadDotEnv:   true,
			expectToLoad: true,
		},
		{
			name:         "empty ENV loads .env",
			envValue:     "",
			expectFoo:    "bar",
			loadDotEnv:   true,
			expectToLoad: true,
		},
		{
			name:         "production skips .env",
			envValue:     "production",
			expectFoo:    "",
			loadDotEnv:   true,
			expectToLoad: false,
		},
		{
			name:         "no .env file in development is fine",
			envValue:     "development",
			expectFoo:    "",
			loadDotEnv:   false,
			expectToLoad: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// cleanup before
			os.Unsetenv("FOO")
			os.Unsetenv("BAZ")

			os.Setenv("ENV", tt.envValue)

			if tt.loadDotEnv {
				writeEnvFile(".env")
				defer removeEnvFile(".env")
			} else {
				os.Remove(".env") // ensure it's not there
			}

			LoadEnv()

			got := os.Getenv("FOO")
			if got != tt.expectFoo {
				t.Errorf("expected FOO=%s, got %s", tt.expectFoo, got)
			}
		})
	}
}

func TestLoadEnvWithoutSpecifyingEnv(t *testing.T) {
	// Ensure ENV, FOO, and BAZ are not set
	os.Unsetenv("ENV")
	os.Unsetenv("FOO")
	os.Unsetenv("BAZ")

	// Write a dummy .env file
	writeEnvFile(".env")
	defer removeEnvFile(".env")

	// Load the environment variables
	LoadEnv()

	// Check if the variables are loaded correctly
	if os.Getenv("FOO") != "bar" {
		t.Errorf("expected FOO=bar, got %s", os.Getenv("FOO"))
	}
	if os.Getenv("BAZ") != "qux" {
		t.Errorf("expected BAZ=qux, got %s", os.Getenv("BAZ"))
	}
}
