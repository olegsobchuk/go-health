package v1controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/olegsobchuk/go-health/configs"
	"github.com/olegsobchuk/go-health/models"
	"github.com/olegsobchuk/go-health/services/secret"
	"github.com/olegsobchuk/go-health/services/tokenizer"
)

// Auth authenticates users
func Auth(c *gin.Context) {
	var user models.User
	c.ShouldBindJSON(&user)
	var userDB models.User
	configs.DB.Find(&userDB, map[string]interface{}{"email": user.Email})
	if !configs.DB.NewRecord(userDB) && secret.Check(userDB.EncPassword, user.Password) {
		token, err := tokenizer.BuildNew(userDB.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "issue with token"})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{"token": token})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "unauthorized"})
}
