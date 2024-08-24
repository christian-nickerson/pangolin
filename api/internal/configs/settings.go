package configs

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

// Settings strucuture
type Settings struct {
	Server   Server   `mapstructure:"server"`
	Metadata Metadata `mapstructure:"metadata"`
}

type Server struct {
	Embeddings ServerConfig `mapstructure:"embeddings"`
	API        ServerConfig `mapstructure:"api"`
}

type Metadata struct {
	Database DatabaseConfig `mapstructure:"database"`
}

// Config types
type ServerConfig struct {
	Name string `mapstructure:"name"`
	Port int    `mapstructure:"port"`
}

type DatabaseConfig struct {
	Type     string `mapstructure:"type"`
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	DbName   string `mapstructure:"dbname"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

// Load reads configurations from a toml file or environment variables
// and returns a Settings struct of all setting variables
func Load(fileName string) (Settings, error) {
	var settings Settings

	// find settings in configs or root
	viper.AddConfigPath("./internal/configs")
	viper.AddConfigPath(".")

	// copy expected behaviour with Dynaconf
	replacer := strings.NewReplacer(".", "__")
	viper.SetEnvKeyReplacer(replacer)
	viper.SetEnvPrefix("PANGOLIN")

	// set up config
	fileBase, fileType := fileNameSplit(fileName)
	viper.SetConfigName(fileBase)
	viper.SetConfigType(fileType)
	viper.AutomaticEnv()

	// read and load settings file
	if err := viper.ReadInConfig(); err != nil {
		return settings, fmt.Errorf("unable to read settings file, %v", err)
	}

	if err := viper.Unmarshal(&settings); err != nil {
		return settings, fmt.Errorf("unable to load settings file, %v", err)
	}

	return settings, nil
}

// take a file name and return the base name and file type from extension
func fileNameSplit(fileName string) (string, string) {
	fileExtension := filepath.Ext(fileName)
	fileType := strings.Trim(fileExtension, ".")
	fileBase := strings.TrimSuffix(fileName, fileExtension)
	return fileBase, fileType
}
