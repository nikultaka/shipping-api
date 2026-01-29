package controllers

import (
	"net/http"
	"shipping-api/models"
	"shipping-api/services"
	"time"

	"github.com/gin-gonic/gin"
)

func PollTracking(c *gin.Context) {

	responses, err := services.PollTracking()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.APIResponse{
			Success: false,
			Message: "Failed to poll tracking",
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.APIResponse{
		Success: true,
		Message: "Tracking poll completed",
		Data: gin.H{
			"total_orders_updated": len(responses),
			"tracking_updates":     responses,
			"timestamp":            time.Now().Format(time.RFC3339),
		},
	})
}
