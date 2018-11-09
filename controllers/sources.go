package controllers

import (
	"github.com/gin-gonic/gin"
)

// IndexSources shows list of existing sources
func IndexSources(c *gin.Context) {
	user := currentUser(c)
	c.HTML(200, "sources", gin.H{
		"currentUser": user,
	})
}
