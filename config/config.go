package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

var Cfg Config

func LoadConfig() {

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}

}
