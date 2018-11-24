package controllers

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/olegsobchuk/go-health/configs"
	"github.com/olegsobchuk/go-health/models"
	"github.com/olegsobchuk/go-health/services/secret"
)

// Root root page
func Root(c *gin.Context) {
	c.HTML(200, "base", gin.H{
		"title":       "Root page",
		"currentUser": currentUser(c),
	})
}

// LogIn login user page
func LogIn(c *gin.Context) {
	session := sessions.Default(c)
	c.HTML(200, "login", gin.H{
		"title":  "LogIn page",
		"login":  "user123",
		"cookie": session.Get("user"),
	})
}

// CreateSession login user page
func CreateSession(c *gin.Context) {
	// bind user from form
	var user models.User
	c.ShouldBind(&user)
	// get user from DB by email
	var userDB models.User
	configs.DB.Find(&userDB, map[string]interface{}{"email": user.Email})
	// match encrypted password
	if !configs.DB.NewRecord(userDB) && secret.Check(userDB.EncPassword, user.Password) {
		session := sessions.Default(c)
		session.Set("userID", userDB.ID)
		session.Save()
		c.Redirect(http.StatusMovedPermanently, "/")
		return
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "User not found or password incorrect"})
}
