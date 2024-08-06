package configs

import (
	"encoding/json"
	"os"
)

type Config struct {
	Postgres *PostgresConfig `json:"postgres,omitempty"`
	Server   *ServerConfig   `json:"server,omitempty"`
}

type PostgresConfig struct {
	DBHost  string `json:"host"`
	DBPort  string `json:"port"`
	DBName  string `json:"DBName"`
	SSLMode string `json:"SSLMode,omitempty"`
}

type ServerConfig struct {
	Port string `json:"port,omitempty"`
}

func LoadConfig() *Config {
	file, err := os.Open("configs/config.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	config := &Config{}
	err = decoder.Decode(config)
	if err != nil {
		panic(err)
	}

	return config
}
