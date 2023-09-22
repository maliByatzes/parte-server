package util

import (
	"github.com/spf13/viper"
)

type Config struct {
	DBSource    string `mapstructure:"DB_SOURCE"`
	DBUrl       string `mapstructure:"DB_URL"`
	HttpAddress string `mapstructure:"HTTP_ADDRESS"`
	TokenKey    string `mapstructure:"TOKEN_KEY"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
