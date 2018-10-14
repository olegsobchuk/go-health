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
			"title":  "New User Page",
			"errors": err.(validator.ValidationErrors),
			"user":   user,
		})

		// Use for validations error handling
		// for _, v := range err.(validator.ValidationErrors) {

	} else {
		res, err := user.Create()
		if err != nil {
		}
		fmt.Printf("VALUE: %#v\n", res.Value)
		fmt.Printf("ERROR: %#v\n", res.Error)
		fmt.Printf("USER: %#v\n", *user)
		c.Redirect(301, "/")
	}
}
