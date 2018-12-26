package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/olegsobchuk/go-health/configs"
	"github.com/olegsobchuk/go-health/controllers/api/v1controllers"
	"github.com/olegsobchuk/go-health/models"
	"github.com/olegsobchuk/go-health/services/tokenizer"
)

// AttachV1 includes list of routes for API
func AttachV1(router *gin.Engine) {
	v1 := router.Group("/v1")
	{
		v1.POST("/auth", v1controllers.Auth)
		v1.Use(checkToken())
		source := v1.Group("/sources")
		{
			source.GET("/", v1controllers.IndexSources)
			source.POST("/", v1controllers.CreateSource)
			source.DELETE("/:id", v1controllers.DeleteSource)
		}
	}
}

func checkToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		authToken := c.Request.Header.Get("Authorization")
		payload, err := tokenizer.Parse(authToken)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "can't parse token"})
			c.Abort()
			return
		}
		userID := uint(payload["userID"].(float64))
		user := new(models.User)
		configs.DB.First(&user, userID)
		if configs.DB.NewRecord(user) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		} else {
			c.Set("currentUser", user)
			c.Next()
		}
	}
}
