package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/olegsobchuk/go-health/configs"
	"github.com/olegsobchuk/go-health/models"
	validator "gopkg.in/go-playground/validator.v8"
)

// IndexSources shows list of existing sources
func IndexSources(c *gin.Context) {
	user := models.User{ID: currentUserID(c)}
	var sources []models.Source
	configs.DB.Model(&user).Limit(models.SourcesPerPage).Related(&sources)
	c.HTML(200, "sources", gin.H{
		"title":       "sources",
		"sources":     sources,
		"currentUser": user,
	})
}

// NewSource render form for New Source
func NewSource(c *gin.Context) {
	source := new(models.Source)
	c.HTML(200, "newSource", gin.H{
		"title":       "New Source",
		"source":      source,
		"currentUser": currentUser,
	})
}

// CreateSource creates Source
func CreateSource(c *gin.Context) {
	fmt.Printf("%#v\n", c.PostForm("status"))
	source := new(models.Source)
	if err := c.ShouldBind(source); err != nil {
		c.HTML(http.StatusUnauthorized, "newSource", gin.H{
			"title":  "New Source",
			"errors": err.(validator.ValidationErrors),
			"source": source,
		})
	} else {
		source.UserID = currentUserID(c)
		_, err := source.Create()
		if err != nil {
			// set message *something come up* and redirect
		}
		c.Redirect(301, "/sources/index")
	}
}

// DestroySource deletes Source
func DestroySource(c *gin.Context) {
	// should be implemented via AJAX
	user := currentUser
	id := c.Query("id")
	configs.DB.Model(&user).Association("Sources").Delete(id)
	// configs.DB.Model(&user).Related(&sources).Delete(id) // TODO: remove if previous works
	c.Redirect(301, "/sources/index")
}
