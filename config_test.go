package main

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestReadConfig_EnvVars(t *testing.T) {
	// Create a temporary directory for config files
	tmpDir, err := os.MkdirTemp("", "cronlog-test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	configPath := filepath.Join(tmpDir, "cronlog.toml")
	configContent := `
[[commands]]
path = "/usr/bin/command"
success_code = 0

[slack]
url = "http://example.com/original"
channel = "#original"
username = "original_user"
`
	if err := os.WriteFile(configPath, []byte(configContent), 0644); err != nil {
		t.Fatal(err)
	}

	t.Run("Reads from TOML", func(t *testing.T) {
		// Ensure environment variables are unset for this test
		// os.Unsetenv is not reversible by t.Setenv easily if we want to restore previous state,
		// but t.Setenv handles restoration automatically.
		// However, we can't "Unset" with t.Setenv, only set to empty string.
		// Assuming the environment is clean or empty strings are treated as unset by our logic.
		// For now, let's just set them to empty strings to be sure they don't interfere if set externally.
		t.Setenv("CRONLOG_SLACK_URL", "")
		t.Setenv("CRONLOG_SLACK_CHANNEL", "")
		t.Setenv("CRONLOG_SLACK_USERNAME", "")

		config := ReadConfig(configPath)
		expected := SlackConfig{
			Url:      "http://example.com/original",
			Channel:  "#original",
			Username: "original_user",
		}
		if !reflect.DeepEqual(config.Slack, expected) {
			t.Errorf("Expected %+v, got %+v", expected, config.Slack)
		}
	})

	t.Run("Env vars override TOML", func(t *testing.T) {
		t.Setenv("CRONLOG_SLACK_URL", "http://example.com/override")
		t.Setenv("CRONLOG_SLACK_CHANNEL", "#override")
		t.Setenv("CRONLOG_SLACK_USERNAME", "override_user")

		config := ReadConfig(configPath)
		expected := SlackConfig{
			Url:      "http://example.com/override",
			Channel:  "#override",
			Username: "override_user",
		}
		if !reflect.DeepEqual(config.Slack, expected) {
			t.Errorf("Expected %+v, got %+v", expected, config.Slack)
		}
	})

	t.Run("Env vars without TOML", func(t *testing.T) {
		t.Setenv("CRONLOG_SLACK_URL", "http://example.com/envonly")
		t.Setenv("CRONLOG_SLACK_CHANNEL", "#envonly")
		t.Setenv("CRONLOG_SLACK_USERNAME", "envonly_user")

		// Non-existent file
		config := ReadConfig(filepath.Join(tmpDir, "nonexistent.toml"))
		expected := SlackConfig{
			Url:      "http://example.com/envonly",
			Channel:  "#envonly",
			Username: "envonly_user",
		}
		if !reflect.DeepEqual(config.Slack, expected) {
			t.Errorf("Expected %+v, got %+v", expected, config.Slack)
		}
	})
}
