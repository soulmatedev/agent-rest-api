package main

import (
	"agent-rest-api/configs"
	"agent-rest-api/models"

	"github.com/gin-gonic/gin"
)

func main() {
	env := models.LoadEnv()
	c := configs.LoadConfig()

	router := gin.Default()
    api := router.Group("/api")
    {
        api.POST("/agents", agentController.CreateAgent)
        api.PUT("/agents/:id", agentController.UpdateAgent)
        api.DELETE("/agents/:id", agentController.DeleteAgent)
        api.GET("/agents", agentController.GetAgents)
    }

    router.Run()
}