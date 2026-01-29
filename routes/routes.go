package routes

import (
	"shipping-api/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	api := router.Group("/api/v1")

	orderRoutes := api.Group("/orders")
	{
		orderRoutes.POST("/create", controllers.CreateOrder)
		orderRoutes.POST("/rates", controllers.CalculateRates)
	}

	trackingRoutes := api.Group("/tracking")
	{
		trackingRoutes.POST("/poll", controllers.PollTracking)
	}

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "healthy",
			"service": "Shipping API",
		})
	})
}
