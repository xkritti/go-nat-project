package config

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/mitchellh/mapstructure"
)

type Environment struct {
	Port   string `mapstructure:"PORT"`
	DbHost string `mapstructure:"DB_HOST"`
	DbPort string `mapstructure:"DB_PORT"`
	DbUser string `mapstructure:"DB_USERNAME"`
	DbPass string `mapstructure:"DB_PASSWORD"`
	DbName string `mapstructure:"DB_DATABASE"`
}

var ENV Environment = Environment{}

func Config() Environment {
	envMap, err := godotenv.Read(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	err = mapstructure.Decode(envMap, &ENV)
	if err != nil {
		fmt.Printf("Error decoding map: %v\n", err)
	}
	return *&ENV
}
