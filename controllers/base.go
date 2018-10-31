package controllers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/olegsobchuk/go-health/configs"
	"github.com/olegsobchuk/go-health/models"
)

func currentUser(c *gin.Context) *models.User {
	session := sessions.Default(c)
	userID := session.Get("userID")
	user := models.User{}
	configs.DB.First(&user, userID)
	if configs.DB.NewRecord(user) {
		return nil
	}
	return &user
}
