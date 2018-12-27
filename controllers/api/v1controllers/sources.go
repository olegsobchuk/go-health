package v1controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/olegsobchuk/go-health/configs"
	"github.com/olegsobchuk/go-health/models"
)

// IndexSources returns list of existing Sources, related to current user
func IndexSources(c *gin.Context) {
	sources := []models.Source{}
	currentUser, _ := c.Get("currentUser")
	configs.DB.Model(currentUser.(*models.User)).Related(&sources)
	c.JSON(http.StatusOK, gin.H{"sources": sources})
}

// CreateSource creates source
func CreateSource(c *gin.Context) {
	source := new(models.Source)
	if err := c.BindJSON(source); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"errors": err.Error(),
			"source": source,
		})
	} else {
		currentUser, exists := c.Get("currentUser")
		if exists {
			source.UserID = currentUser.(*models.User).ID
			_, err := source.Create()
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"status": "bad_request"})
			} else {
				c.JSON(http.StatusCreated, gin.H{"status": "OK"})
			}
		} else {
			c.JSON(http.StatusUnauthorized, "/")
		}
	}
}

// DeleteSource deletes existing Source
func DeleteSource(c *gin.Context) {
	currentUser, exists := c.Get("currentUser")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is invalid"})
		return
	}
	source := models.Source{ID: uint(id)}

	result := configs.DB.Model(&source).Where("user_id = ?", currentUser.(*models.User).ID).Delete(&source)
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"error":   result.Error,
		"releted": result.RowsAffected,
	})
}
