package routes

import (
	controller "pizza-site-backend/controllers/product"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(router *gin.Engine) {
	router.POST("/products/", controller.CreateProduct())
	router.GET("/products/:product_id", controller.GetProducts())
	router.GET("/products/", controller.GetAllProducts())
	router.POST("/products/:product_id", controller.ModifyProduct())
	router.DELETE("/products/:product_id", controller.DeleteProduct())
}
