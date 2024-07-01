package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

var Cfg Config

func LoadConfig(directory string) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(directory)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
    if config.Username == "" || config.Password == "" {
        log.Fatalf("Username and password must be set in config file")
    }

	Cfg = config
}
