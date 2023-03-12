package main

import (
	"os"

	middleware "pizza-site-backend/middleware"
	routes "pizza-site-backend/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())
	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	routes.OrderRoutes(router)
	routes.ProductRoutes(router)
	routes.TransactionRoutes(router)

	router.Run(":" + port)
}
