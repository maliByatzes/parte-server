package config

import (
	"fmt"
	"strings"
	"time"

	"github.com/spf13/viper"
)

type (
	Config struct {
		Database DatabaseConfig
		HTTP     HTTPConfig
		Cache    CacheConfig
		Token    TokenConfig
	}

	DatabaseConfig struct {
		HostName string
		Port     uint16
		Source   string
		User     string
		Password string
		Database string
	}

	HTTPConfig struct {
		HostName     string
		Port         uint16
		ReadTimeout  time.Duration
		WriteTimeout time.Duration
		IdleTimeout  time.Duration
		TLS          struct {
			Enabled     bool
			Certificate string
			Key         string
		}
	}

	CacheConfig struct {
		HostName string
		Port     uint16
		Password string
	}

	TokenConfig struct {
		TokenKey string
	}
)

func LoadConfig() (Config, error) {
	var config Config

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("config")
	viper.AddConfigPath("../config")
	viper.AddConfigPath("../../config")

	viper.SetEnvPrefix("parte")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		return config, fmt.Errorf("could not read in the config: %v", err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		return config, err
	}

	return config, nil
}
