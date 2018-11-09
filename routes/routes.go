package routes

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/olegsobchuk/go-health/configs"
	"github.com/olegsobchuk/go-health/controllers"
	"github.com/olegsobchuk/go-health/models"
)

// Attach runs all existing routes
func Attach(router *gin.Engine) {
	router.GET("/", controllers.Root)
	router.GET("/login", controllers.LogIn)
	router.POST("/session", controllers.CreateSession)

	// users grouping
	router.GET("/users/new", controllers.NewUser)
	router.POST("/users/create", controllers.CreateUser)
	users := router.Group("/users")
	{
		users.Use(checkCurrentUser())
		users.GET("/show", controllers.ShowUser)
	}
	source := router.Group("/sources")
	{
		source.Use(checkCurrentUser())
		source.GET("/index", controllers.IndexSources)
		// source.GET("/new")
		// source.POST("/create")
		// source.GET("/edit")
		// source.PATCH("/update")
	}
}

func checkCurrentUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		userID := session.Get("userID")
		user := new(models.User)
		configs.DB.First(&user, userID)
		if configs.DB.NewRecord(user) {
			c.Redirect(http.StatusPermanentRedirect, "/login")
		} else {
			c.Next()
		}
	}
}
