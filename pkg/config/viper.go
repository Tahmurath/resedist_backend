package config

import (
	"github.com/spf13/viper"
	"log"
	"resedist/config"
)

var configurations config.Config

func Get() config.Config {
	return configurations
}

func Set(path string, file string) {
	viper.SetConfigName(file)   // name of config file (without extension)
	viper.SetConfigType("yaml") // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(path)   // path to look for the config file in

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading config file", err)
	}

	err := viper.Unmarshal(&configurations)
	if err != nil {
		log.Fatal("Error parsing config file")
	}

}
