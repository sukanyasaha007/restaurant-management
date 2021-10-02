package routes

import (
	"golang-restaurant-management/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(inComingRoutes *gin.Engine) {
	inComingRoutes.GET("/users", controllers.GetUsers())
	inComingRoutes.GET("/users/:user_id", controllers.GetUser())
	inComingRoutes.POSt("/users/signup", controllers.SignUp())
	inComingRoutes.POST("/users/login", controllers.Login())
}
