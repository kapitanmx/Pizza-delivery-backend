package routes

import (
	controller "pizza-site-backend/controllers/transaction"

	"github.com/gin-gonic/gin"
)

func TransactionRoutes(router *gin.Engine) {
	router.POST("/transactions/", controller.GenerateTransaction())
	router.GET("/transactions/:transaction_id", controller.GetTransaction())
	router.GET("/transactions/", controller.GetAllTransactions())
}
