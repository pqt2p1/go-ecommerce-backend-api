package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"server"`
	Database []struct {  
		User string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Host string `mapstructure:"host"`
	} `mapstructure:"database"`
}

func main() {
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
	fmt.Println("Server port is: ", viper.GetInt("server.port"))
	var config Config
	if err :=  viper.Unmarshal(&config); err != nil {
		panic(fmt.Errorf("Failed to unmarshal configuration: %w \n", err))
	}

	fmt.Println("Server port is: ", config.Server.Port)

	// read database configuration
	for _, db := range config.Database {
		fmt.Printf("Database user is: %s, password is: %s, host is: %s \n", db.User, db.Password, db.Host)
	}

}
