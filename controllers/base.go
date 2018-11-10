package controllers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/olegsobchuk/go-health/configs"
	"github.com/olegsobchuk/go-health/models"
)

func currentUser(c *gin.Context) *models.User {
	user := models.User{}
	configs.DB.First(&user, currentUserID(c))
	if configs.DB.NewRecord(user) {
		return nil
	}
	return &user
}

func currentUserID(c *gin.Context) uint {
	session := sessions.Default(c)
	ID := session.Get("userID")
	userID, ok := ID.(uint)
	if !ok {
		panic("can't convert currentUserID to uint type")
	}
	return userID
}
