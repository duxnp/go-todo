package todo

import (
	"net/http"
	"todo/internal/model"
	"todo/internal/util"

	"github.com/labstack/echo/v4"
)

func (d database) allTodos() []model.Todo {
	todos := []model.Todo{}
	d.db.Select(&todos, "SELECT * FROM todo t")
	return todos
}

func (h handlers) indexHandler(c echo.Context) error {
	todos := h.repo.allTodos()
	return util.Render(c, http.StatusOK, Index(todos, c.Echo()))
}
