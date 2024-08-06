package repositories

import (
	"database/sql"
	"fmt"

	"agent-rest-api/configs"
	"agent-rest-api/models"

	_ "github.com/lib/pq"
)

func NewDatabase(env *models.EnvironmentVariables, config *configs.Config) (*sql.DB, error) {
	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Postgres.DBHost, config.Postgres.DBPort, env.PostgresUser, env.PostgresPassword, config.Postgres.DBName)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return db, nil
}