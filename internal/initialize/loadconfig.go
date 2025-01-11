package initialize

import (
	"fmt"

	"github.com/pqt2p1/go-ecommerce-backend-api/global"
	"github.com/spf13/viper"
)

func LoadConfig() {
	viper := viper.New()
	viper.AddConfigPath("./config/")
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")

	// read configuration
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Failed to read configuration: %w \n", err))
	}
	// read server configuration
	if err := viper.Unmarshal(&global.Config); err != nil {
		panic(fmt.Errorf("Failed to unmarshal configuration: %w \n", err))
	}
}
