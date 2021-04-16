package config

import (
	"github.com/spf13/viper"
)

var Cfg Config

type Config struct {
	Server struct {
		Address string `yaml:"address" env:"SERVER_ADDRESS"`
		Port    string `yaml:"port" env:"SERVER_PORT"`
	} `yaml:"server"`
}

func GetConfig(configDir string) error {
	viper.SetConfigName("config")
	viper.AddConfigPath(configDir)

	var config Config

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	err := viper.Unmarshal(&config)
	if err != nil {
		return err
	}

	Cfg = config

	return nil
}
