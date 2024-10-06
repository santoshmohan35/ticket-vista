package configuration

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Port int `json:"port"`
}

func LoadConfig() (Config, error) {
	var config Config
	viper.AddConfigPath("./internal/configuration")
	viper.SetConfigName("config") // Register config file name (no extension)
	viper.SetConfigType("json")   // Look for specific type

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %s", err)
		return config, err
	}
	err := viper.Unmarshal(&config)
	if err != nil {
		return config, err
	}
	fmt.Println("configurations loaded ", config.Port)
	return config, nil
}
