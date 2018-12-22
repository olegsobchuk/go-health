package controllers

import (
  "net/http"
  "strconv"

  "github.com/gin-gonic/gin"
  "github.com/olegsobchuk/go-health/configs"
  "github.com/olegsobchuk/go-health/models"
  validator "gopkg.in/go-playground/validator.v9"
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
    "currentUser": currentUser(c),
  })
}

// CreateSource creates Source
func CreateSource(c *gin.Context) {
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
    c.Redirect(http.StatusMovedPermanently, "/sources/index")
  }
}

// UpdateSource updates existing Source
func UpdateSource(c *gin.Context) {

}

// DestroySource delete existing Source
func DestroySource(c *gin.Context) {
  curUser := currentUser(c)
  sourceIDString := c.Param("id")
  sourceID, err := strconv.ParseUint(sourceIDString, 10, 64)
  if err == nil {
    source := models.Source{ID: uint(sourceID)}
    configs.DB.Model(curUser).Association("Sources").Find(&source)
    configs.DB.Delete(&source)
  }
  c.Redirect(http.StatusMovedPermanently, "/")
}
