package configs

import (
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
func LoadSettings(fileName string) (settings Settings, err error) {
	// find settings in configs or root
	viper.AddConfigPath("./internal/configs")
	viper.AddConfigPath(".")

	// copy expected behaviour with Dynaconf
	replacer := strings.NewReplacer(".", "__")
	viper.SetEnvKeyReplacer(replacer)
	viper.SetEnvPrefix("PANGOLIN")

	// set up config
	viper.SetConfigName(fileName)
	viper.SetConfigType("toml")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&settings)
	return
}
