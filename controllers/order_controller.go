package controllers

import (
	"net/http"
	"shipping-api/models"
	"shipping-api/services"

	"github.com/gin-gonic/gin"
)

func CreateOrder(c *gin.Context) {
	var request models.CreateOrderRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, models.APIResponse{
			Success: false,
			Message: "Validation failed",
			Error:   err.Error(),
		})
		return
	}

	orderResponse, err := services.CreateOrderService(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.APIResponse{
			Success: false,
			Message: "Failed to create order",
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.APIResponse{
		Success: true,
		Message: "Order created successfully",
		Data:    orderResponse,
	})
}

func CalculateRates(c *gin.Context) {
	var request models.RateRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, models.APIResponse{
			Success: false,
			Message: "Validation failed",
			Error:   err.Error(),
		})
		return
	}

	ratesResponse := services.CalculateRatesService(request)

	c.JSON(http.StatusOK, models.APIResponse{
		Success: true,
		Message: "Rates calculated successfully",
		Data:    ratesResponse,
	})
}
