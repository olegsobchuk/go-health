package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/olegsobchuk/go-health/models"
	validator "gopkg.in/go-playground/validator.v8"
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
			"title": "New User Page",
			"error": err,
			"user":  user,
		})

		// TODO error handling
		for _, v := range err.(validator.ValidationErrors) {
			fmt.Println(v.Field)
			fmt.Println(v.Tag)
		}

		fmt.Printf("%#v", err.(validator.ValidationErrors))
	} else {
		c.JSON(http.StatusOK, gin.H{"user": user})
	}
}
