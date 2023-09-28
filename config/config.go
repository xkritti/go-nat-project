package config

import (
	"fmt"
	// "os"

	"github.com/joho/godotenv"
	"github.com/mitchellh/mapstructure"
)

type ENV struct {
	Port   string `mapstructure:"PORT"`
	DbHost string `mapstructure:"DB_HOST"`
	DbPort string `mapstructure:"DB_PORT"`
	DbUser string `mapstructure:"DB_USERNAME"`
	DbPass string `mapstructure:"DB_PASSWORD"`
	DbName string `mapstructure:"DB_DATABASE"`
}

var env = ENV{}

func Config() ENV {

	envMap, err := godotenv.Read(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	err = mapstructure.Decode(envMap, &env)
	if err != nil {
		fmt.Printf("Error decoding map: %v\n", err)
	}

	return *&env

}
