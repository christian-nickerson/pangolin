package configs

import (
	"testing"

	_ "github.com/christian-nickerson/pangolin/api/testing"
)

// test load config can read the default toml file
// and has required default values
func TestLoadConfig(t *testing.T) {
	settings, err := LoadSettings("settings")
	if err != nil {
		t.Fatalf("failed to load config: %v", err)
	}

	if settings.EmbeddingServer.Name == "" {
		t.Fatalf("embeddinger-server.name should have a default value")
	}
	if settings.EmbeddingServer.Port != 50051 {
		t.Fatalf("embeddinger-server.port should have a default of 50051")
	}
}
