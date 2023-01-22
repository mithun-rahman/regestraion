package config

import (
	"github.com/spf13/viper"
	"log"
)

func readConfig(fileName string) {
	viper.SetConfigFile(fileName)
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("unable to read config file: %v", err)
	}
}
