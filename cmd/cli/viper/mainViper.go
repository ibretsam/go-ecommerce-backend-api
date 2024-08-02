package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"server"`
	Databases []struct {
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Host     string `mapstructure:"host"`
	} `mapstructure:"databases"`
}

func main() {
	viper := viper.New()
	viper.AddConfigPath("./config/")
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")

	// Read configuration
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	// Config struct
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		fmt.Println("Unable to decode the configuration", err)
	}

	for _, db := range config.Databases {
		fmt.Println("Database User:: ", db.User)
		fmt.Println("Database Password:: ", db.Password)
		fmt.Println("Database Host:: ", db.Host)
	}
}
