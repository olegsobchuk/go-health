package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/olegsobchuk/go-health/controllers"
)

// Attach runs all existing routes
func Attach(router *gin.Engine) {
	router.GET("/", controllers.Root)
	router.GET("/login", controllers.LogIn)
	router.POST("/session", controllers.CreateSession)

	// users grouping
	users := router.Group("/users")
	{
		users.GET("/new", controllers.NewUser)
		users.POST("/create", controllers.CreateUser)
	}
}
