package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/olegsobchuk/go-health/models"
)

// NewUser build new user form
func NewUser(c *gin.Context) {
	var user = models.User{}
	c.HTML(200, "newUser", gin.H{
		"title": "New User Page",
		"user":  user,
	})
}

// CreateUser create new user
func CreateUser(c *gin.Context) {
	var user = new(models.User)
	if err := c.ShouldBind(user); err != nil {
		c.HTML(http.StatusUnauthorized, "newUser", gin.H{
			"title":  "New User Page",
			"errors": err,
			"user":   user,
		})

		// Use for validations error handling
		// for _, v := range err.(validator.ValidationErrors) {

	} else {
		_, err := user.Create()
		if err != nil {
			// set message *something come up* and redirect
		}
		c.Redirect(301, "/")
	}
}

// ShowUser shows information about user
func ShowUser(c *gin.Context) {
	user := c.MustGet("currentUser")
	c.HTML(200, "showUser", gin.H{
		"currentUser": user,
	})
}
