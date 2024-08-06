package handlers

import (
	"agent-rest-api/controllers"

	"github.com/gin-gonic/gin"
)

func InitRoutes(controller *controllers.AgentController) *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.POST("/agents", controller.CreateAgent)
		api.GET("/agents", controller.GetAgents)
		api.PUT("/agents/:id", controller.UpdateAgent)
		api.DELETE("/agents/:id", controller.DeleteAgent)
	}
	return router
}