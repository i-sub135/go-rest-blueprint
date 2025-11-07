package config_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/i-sub135/go-rest-blueprint/source/config"
)

func TestLoadConfig_BasicFunctionality(t *testing.T) {
	config.ResetConfig()
	clearEnvVars()

	err := config.LoadConfig("")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	cfg := config.GetConfig()
	if cfg.App.Mode != "local" {
		t.Errorf("Expected default App.Mode=local, got %s", cfg.App.Mode)
	}
	if cfg.App.Port != 8080 {
		t.Errorf("Expected default app.port=8080, got %d", cfg.App.Port)
	}
}

func TestLoadConfig_WithFile(t *testing.T) {
	config.ResetConfig()
	clearEnvVars()

	content := "app:\n  env: test\n  port: 9000\ndb:\n  dsn: test.db\nlog:\n  level: info"
	tmpFile := createTempYAML(t, content)
	defer os.Remove(tmpFile)

	err := config.LoadConfig(tmpFile)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	cfg := config.GetConfig()
	if cfg.App.Mode != "test" {
		t.Errorf("Expected App.Mode=test, got %s", cfg.App.Mode)
	}
	if cfg.App.Port != 9000 {
		t.Errorf("Expected app.port=9000, got %d", cfg.App.Port)
	}
}

func createTempYAML(t *testing.T, content string) string {
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "test_config.yaml")
	err := os.WriteFile(tmpFile, []byte(content), 0644)
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	return tmpFile
}

func clearEnvVars() {
	envVars := []string{"APP_ENV", "APP_PORT", "DB_DSN", "LOG_LEVEL"}
	for _, envVar := range envVars {
		os.Unsetenv(envVar)
	}
}
