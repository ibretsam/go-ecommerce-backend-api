package initialize

import (
	"fmt"
	"go-ecommerce-backend-api/global"

	"github.com/spf13/viper"
)

func LoadConfig() {
	// Load configuration from file
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
	if err := viper.Unmarshal(&global.Config); err != nil {
		fmt.Println("Unable to decode the configuration", err)
	}
}
