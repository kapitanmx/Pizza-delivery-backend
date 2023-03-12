package routes

import (
	controller "pizza-site-backend/controllers/user"
	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	router.POST("/users/signup", controller.SignUp())
	router.POST("/users/login", controller.Login())
	router.GET("/users/:user_id", controller.GetUser())
	router.GET("/users/", controller.GetAllUsers())
	router.POST("/users/:user_id", controller.EditUser())
	router.DELETE("/users/:user_id", controllers.DeleteUser())
}
