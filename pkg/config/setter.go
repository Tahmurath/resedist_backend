package config

import (
	"github.com/spf13/viper"
	"log"
)

func Set() {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("config") // path to look for the config file in

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading config file")
	}

	err := viper.Unmarshal(&configurations)
	if err != nil {
		log.Fatal("Error parsing config file")
	}

}
