package main

import (
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/olegsobchuk/go-health/configs"
	"github.com/olegsobchuk/go-health/routes"
)

const secretKey = "KJHSAD&*&ASDJSDH87asd!@01"

func main() {
	router := gin.Default()
	store := cookie.NewStore([]byte(secretKey))
	router.Use(sessions.Sessions("_health_session", store))
	router.Static("/assets", "./assets")
	router.LoadHTMLGlob("templates/*/**")
	router.HTMLRender = buildTemplate()
	configs.Init()
	routes.Attach(router)

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
