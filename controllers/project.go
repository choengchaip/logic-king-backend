package controllers

import (
	"github.com/labstack/echo/v4"
	"logic_king_backend/cores"
	"logic_king_backend/services/projects"
	"net/http"
)

type ProjectController struct {
}

func (p *ProjectController) Pagination(ctx *cores.AppContext) echo.HandlerFunc {
	return func(c echo.Context) error {
		s := projects.NewProjectService(ctx)
		items, err := s.Pagination()
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"message": err.Error(),
			})
		}

		return c.JSON(http.StatusOK, items)
	}
}
