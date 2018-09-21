package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/olegsobchuk/go-health/models"
)

// Root root page
func Root(c *gin.Context) {
	session := sessions.Default(c)
	session.Set("user", 22)
	session.Save()
	c.HTML(200, "base", gin.H{
		"title": "Root page",
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
	var user models.User
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	fmt.Printf("%v", user)
	if user.Email != "uder123" && user.Password != "123456" {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		fmt.Println("err")
	}
	c.Redirect(http.StatusMovedPermanently, "/")
}
