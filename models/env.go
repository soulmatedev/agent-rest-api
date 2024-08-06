package models

import (
	"fmt"

	"github.com/caarlos0/env/v9"
	"github.com/joho/godotenv"
)

type EnvironmentVariables struct {
	PostgresUser       string `env:"POSTGRES_USER,required,notEmpty"`
	PostgresPassword   string `env:"POSTGRES_PASSWORD,required,notEmpty"`
	Domain 	   		   string `env:"DOMAIN,required,notEmpty"`
}

const envFilePath = ".env.local"

func LoadEnv() *EnvironmentVariables {
	environmentVariables := EnvironmentVariables{}
	if err := godotenv.Load(envFilePath); err != nil {
		panic(fmt.Sprintf("Failed to load .env.local file: %s", err.Error()))
	}
	if err := env.Parse(&environmentVariables); err != nil {
		panic(fmt.Sprintf("Failed to parse environment variables: %s", err.Error()))
	}

	return &environmentVariables
}