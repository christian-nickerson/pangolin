package configs

import (
	"testing"

	_ "github.com/christian-nickerson/pangolin/api/testing"
	"github.com/stretchr/testify/assert"
)

// test load config can read the default toml file
// and has required default values
func TestLoadConfig(t *testing.T) {
	assert := assert.New(t)
	settings, err := LoadSettings("settings")

	assert.Nil(err, "failed to load config: %v", err)

	// server.embeddings
	assert.NotEqual(settings.Server.Embeddings.Name, "", "server.embeddings.name should have a default value")
	assert.Equal(settings.Server.Embeddings.Port, 50051, "server.embedding.port should have a default of 50051")

	// server.api
	assert.NotEqual(settings.Server.Embeddings.Name, "", "server.embeddings.name should have a default value")
	assert.Equal(settings.Server.Embeddings.Port, 50051, "server.embedding.port should have a default of 50051")
}
