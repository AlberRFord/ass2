package routes

import (
	"myproject/controllers"

	"github.com/gin-gonic/gin"
)

// AuthRoutes defines authentication routes
func AuthRoutes(router *gin.Engine) {
	auth := router.Group("/auth")
	{
		auth.POST("/register", controllers.Register)
	}
}
