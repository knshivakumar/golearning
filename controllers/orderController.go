package controllers

import (
	"fmt"
	"learning-golang/models"
	"learning-golang/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OrderController struct {
	oservice services.OrderService
}

func NewOrderController(s services.OrderService) *OrderController {
	return &OrderController{oservice: s}
}

func (c *OrderController) CreateOrder(ctx *gin.Context) {
	var order models.Order

	if err := ctx.ShouldBindJSON(&order); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(order)
	orderId, error := c.oservice.CreateOrder(order)

	if error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": error})
		return
	}
	ctx.JSON(http.StatusOK, orderId)
}
