package router

import (
	"github.com/Jack-ZL/go_rookie"
	"goRookie_project/api/handler"
)

func InitRoutes() *go_rookie.Engine {
	engine := go_rookie.Default()
	gGroup := engine.Group("/goods")
	gGroup.Get("/listGoods", handler.ListGoods)

	return engine
}
