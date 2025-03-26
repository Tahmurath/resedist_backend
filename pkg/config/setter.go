package config

import (
	"log"

	"github.com/spf13/viper"
)

func Set(path string) {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(path)     // path to look for the config file in

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading config file", err)
	}

	err := viper.Unmarshal(&configurations)
	if err != nil {
		log.Fatal("Error parsing config file")
	}

}
