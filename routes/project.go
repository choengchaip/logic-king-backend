package routes

import (
	"logic_king_backend/controllers"
	"logic_king_backend/cores"
)

func NewProjectRoute(ctx *cores.AppContext) {
	controller := controllers.ProjectController{}

	ctx.Echo.GET("/projects", controller.Pagination(ctx))
}
