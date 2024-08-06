package configs

import (
	"encoding/json"
	"os"
)

const configFilePath = "configs/config.json"

type Config struct {
	Postgres *PostgresConfig `json:"postgres,omitempty"`
	Server   *ServerConfig   `json:"server,omitempty"`
}

type PostgresConfig struct {
	DBHost     string `json:"db_host"`
    DBPort     string `json:"db_port"`
    DBName     string `json:"db_name"`
    SSLMode string `json:"SSLMode,omitempty"`
}

type ServerConfig struct {
	Port string `json:"port,omitempty"`
}

func LoadConfig() *PostgresConfig {
	file, err := os.Open(configFilePath);

	if err != nil {
        panic(err)
    }
    defer file.Close()

	decoder := json.NewDecoder(file)
    config := PostgresConfig{}
    err = decoder.Decode(&config)
    if err != nil {
        panic(err)
    }

    return &config
}
