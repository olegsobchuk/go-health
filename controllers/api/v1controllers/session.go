package v1controllers

import (
  "github.com/gin-gonic/gin"
  "net/http"
  "fmt"

  "github.com/olegsobchuk/go-health/models"
  "github.com/olegsobchuk/go-health/services/tokenizer"
  "github.com/olegsobchuk/go-health/configs"
  "github.com/olegsobchuk/go-health/services/secret"
)

// Auth authenticates users
func Auth(c *gin.Context) {
  var user models.User
  c.ShouldBindJSON(&user)
  var userDB models.User
  configs.DB.Find(&userDB, map[string]interface{}{"email": user.Email})
  if !configs.DB.NewRecord(userDB) && secret.Check(userDB.EncPassword, user.Password) {
    token, err := tokenizer.BuildNew(user.ID)
    fmt.Printf("%+v\n", err)
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
