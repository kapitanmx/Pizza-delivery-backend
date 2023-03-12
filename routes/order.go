package routes

import (
	controller "pizza-site-backend/controllers/order"

	"github.com/gin-gonic/gin"
)

func OrderRoutes(router *gin.Engine) {
	router.POST("/orders/", controller.MakeOrder())
	router.GET("/orders/:order_id", controller.GetOrder())
	router.GET("/orders/", controller.GetAllOrders())
	router.POST("/orders/:order_id", controller.UpdateOrder())
	router.DELETE("/orders/:order_id", controller.DeleteOrder())
}
