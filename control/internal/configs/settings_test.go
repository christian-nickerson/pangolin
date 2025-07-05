package configs

import (
	"os"
	"testing"

	_ "github.com/christian-nickerson/pangolin/control/testing"
	"github.com/stretchr/testify/assert"
)

// test load config can read the default toml file
// and has required default values
func TestLoadSettings(t *testing.T) {
	settings, err := Load("settings.toml")
	assert.NoError(t, err, nil, "failed to load setings file")

	// server.embeddings
	assert.NotEqual(
		t,
		settings.Server.Embeddings.Name,
		"",
		"server.embeddings.name should have a default value",
	)
	assert.Equal(
		t,
		settings.Server.Embeddings.Port,
		50051,
		"server.embedding.port should have a default of 50051",
	)

	// server.api
	assert.NotEqual(t, settings.Server.API.Name, "", "server.api.name should have a default value")
	assert.Equal(
		t,
		settings.Server.API.Port,
		3000,
		"server.api.port should have a default of 50051",
	)
}

// assert settings can be overridden by env variables
func TestSettingsOverride(t *testing.T) {
	os.Setenv("PANGOLIN_SERVER__EMBEDDINGS__NAME", "test")
	os.Setenv("PANGOLIN_SERVER__EMBEDDINGS__PORT", "455")
	settings, err := Load("settings.toml")
	assert.NoError(t, err, nil, "failed to load setings file")

	// assert overrides
	assert.Equal(t, settings.Server.Embeddings.Name, "test")
	assert.Equal(t, settings.Server.Embeddings.Port, 455)
}

// assert fileNameSplit breaks file names correctly
func TestFileNameSplitTOML(t *testing.T) {
	fileBase, fileType := fileNameSplit("settings.toml")
	assert.Equal(t, fileBase, "settings")
	assert.Equal(t, fileType, "toml")
}
