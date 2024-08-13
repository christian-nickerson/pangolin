package configs

import (
	"log"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

type ServerConfig struct {
	Name string `mapstructure:"name"`
	Port int    `mapstructure:"port"`
}

type Server struct {
	Embeddings ServerConfig `mapstructure:"embeddings"`
	API        ServerConfig `mapstructure:"api"`
}

type Settings struct {
	Server Server `mapstructure:"server"`
}

// LoadSettings reads configurations from a toml file or environment variables
// and returns a Settings struct of all setting variables
func Load(fileName string) (settings Settings) {
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

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Loading settings failed", err)
	}

	err = viper.Unmarshal(&settings)
	if err != nil {
		log.Fatal("Failed to fit file to settings:", err)
	}
	return
}

// take a file name and return the base name and file type from extension
func fileNameSplit(fileName string) (fileBase string, fileType string) {
	fileExtension := filepath.Ext(fileName)
	fileType = strings.Trim(fileExtension, ".")
	fileBase = strings.TrimSuffix(fileName, fileExtension)
	return
}
