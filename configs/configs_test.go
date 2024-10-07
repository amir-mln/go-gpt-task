package configs_test

import (
	"go-gpt-task/configs"
	"os"
	"path"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	dummyApiKey := "dummy_api_key"
	os.Setenv(configs.EnvVarAPIKey, dummyApiKey)

	wd, err := os.Getwd()
	if err != nil {
		t.Fatal("Failed to read working directory.", err)
	}

	cfg, err := configs.Load(path.Join(wd, "..", ".env"))
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}

	if cfg.APIKey != dummyApiKey {
		t.Errorf("Expected %q to be %q, got %q", configs.EnvVarAPIKey, dummyApiKey, cfg.APIKey)
	}
}
