package main

import (
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"github.com/olegsobchuk/go-health/routes"
)

func main() {
	router := gin.Default()
	router.Static("/assets", "./assets")
	router.LoadHTMLGlob("templates/*/**")
	router.HTMLRender = buildTemplate()
	routes.Attach(router)

	router.Run(":9000")
}

func buildTemplate() multitemplate.Renderer {
	r := multitemplate.NewRenderer()
	r.AddFromFiles("base", "templates/layouts/main.tmpl", "templates/root/root.tmpl")
	r.AddFromFiles("login", "templates/layouts/main.tmpl", "templates/session/login.tmpl")
	return r
}
