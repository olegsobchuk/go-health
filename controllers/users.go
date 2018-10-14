package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/olegsobchuk/go-health/models"
	validator "gopkg.in/go-playground/validator.v8"
	pg "gopkg.in/pg.v4"
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
		err := user.Create()
		if err != nil {
			dbErr, ok := err.(pg.Error)
			fmt.Printf("Column name: %#v\n", dbErr.Field('c'))
			fmt.Printf("Constraint name: %#v\n", dbErr.Field('n'))
			fmt.Printf("Where: %#v\n", dbErr.Field('W'))
			fmt.Printf("Detail: %#v\n", dbErr.Field('D'))
			fmt.Printf("Message: %#v\n", dbErr.Field('M'))
			fmt.Printf("Code: %#v\n", dbErr.Field('C'))
			fmt.Printf("Position: %#v\n", dbErr.Field('P'))
			fmt.Printf("Table name: %#v\n", dbErr.Field('t'))
			if ok && dbErr.IntegrityViolation() {
				fmt.Printf("OK: %#v\n", ok)
				fmt.Printf("Err: %#v\n", dbErr.IntegrityViolation())
			}
		}
		fmt.Printf("%#v\n", err)
		c.Redirect(301, "/")
	}
}
