package api

import (
  "github.com/gin-gonic/gin"
  "github.com/olegsobchuk/go-health/controllers/api/v1controllers"
)

// AttachV1 includes list of routes for API
func AttachV1(router *gin.Engine) {
  v1 := router.Group("/v1")
  {
    v1.POST("/auth", v1controllers.Auth)
  }
}
