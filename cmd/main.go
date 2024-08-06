package main

import (
	"agent-rest-api/configs"
	"agent-rest-api/controllers"
	"agent-rest-api/handlers"
	"agent-rest-api/models"
	"agent-rest-api/repositories"
	"agent-rest-api/services"
	"log"
)

func main() {
	env := models.LoadEnv()
	c := configs.LoadConfig()

	db, err := repositories.NewDatabase(env, c)
	if err != nil {
		log.Fatalf("Failed to initialize database: %s", err.Error())
	}

	repo := repositories.NewAgentRepository(db)
	service := services.NewAgentService(repo)
	controller := controllers.NewAgentController(service)

	router := handlers.InitRoutes(controller)
	router.Run(c.Postgres.DBHost + ":" + c.Server.Port)
}