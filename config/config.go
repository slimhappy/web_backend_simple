package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName("config.yaml") // name of config file (without extension)
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
}
