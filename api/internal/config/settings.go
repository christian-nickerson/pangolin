package config

import (
	"strings"

	"github.com/spf13/viper"
)

type EmbeddingServer struct {
	Name           string `mapstructure:"name"`
	Port           string `mapstructure:"port"`
	ShutdownPeriod int    `mapstructure:"shutdown_period"`
	WorkerThreads  int    `mapstructure:"worker_threads"`
}

type Transformer struct {
	ModelList []string `mapstructure:"model_list"`
}

type Spacy struct {
	ModelList []string `mapstructure:"model_list"`
}

type Config struct {
	EmbeddingServer EmbeddingServer `mapstructure:"embedding-server"`
	Transformer     Transformer     `mapstructure:"transformer"`
	Spacy           Spacy           `mapstructure:"spacy"`
}

// LoadConfig reads configuration variables from toml or environment variables
func LoadConfig(name string) (config Config, err error) {
	// find settings in config or root
	viper.AddConfigPath("./internal/config")
	viper.AddConfigPath(".")

	// copy expected behaviour with Dynaconf
	replacer := strings.NewReplacer(".", "__")
	viper.SetEnvKeyReplacer(replacer)
	viper.SetEnvPrefix("PANGOLIN")

	// set up config
	viper.SetConfigName(name)
	viper.SetConfigType("toml")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
