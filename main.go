package main

import (
	"io"
	"os"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/olegsobchuk/go-health/configs"
	"github.com/olegsobchuk/go-health/models"
	"github.com/olegsobchuk/go-health/routes"
	"github.com/olegsobchuk/go-health/routes/api"
)

const secretKey = "KJHSAD&*&ASDJSDH87asd!@01"

func main() {
	router := gin.Default()
	store := cookie.NewStore([]byte(secretKey))
	logToFile()
	router.Use(gin.Logger())
	router.Use(sessions.Sessions("_health_session", store))
	router.Static("/assets", "./assets")
	router.LoadHTMLGlob("templates/*/**")
	router.HTMLRender = buildTemplate()
	configs.InitKV()
	configs.InitDB()
	routes.Attach(router)
	api.AttachV1(router)
	var source models.Source

	source.AddToKVStorage()

	router.Run(":9000")
}

func buildTemplate() multitemplate.Renderer {
	r := multitemplate.NewRenderer()
	r.AddFromFiles("base", "templates/layouts/main.tmpl", "templates/root/root.tmpl")
	r.AddFromFiles("login", "templates/layouts/main.tmpl", "templates/session/login.tmpl")
	r.AddFromFiles("newUser", "templates/layouts/main.tmpl", "templates/users/new.tmpl")
	r.AddFromFiles("showUser", "templates/layouts/main.tmpl", "templates/users/show.tmpl")
	r.AddFromFiles("sources", "templates/layouts/main.tmpl", "templates/sources/index.tmpl")
	r.AddFromFiles("newSource", "templates/layouts/main.tmpl", "templates/sources/new.tmpl")
	return r
}

func logToFile() {
	// Logging to a file.
	f, _ := os.Create("logs/base.log")
	gin.DefaultWriter = io.MultiWriter(f)
}
