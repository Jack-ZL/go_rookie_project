package router

import (
	"github.com/Jack-ZL/go_rookie"
	"goRookie_project/api/handler"
)

func InitRoutes(gr *go_rookie.Engine) {
	gGroup := gr.Group("/admin")
	gGroup.Post("/login", handler.AdminLogin)
}
