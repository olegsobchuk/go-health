package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/olegsobchuk/go-health/configs"
	"github.com/olegsobchuk/go-health/models"
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
