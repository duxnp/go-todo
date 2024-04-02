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
	// e := c.Get("echo").(*echo.Echo)
	// foo := e.Reverse("get-todos")
	// _ = foo
	todos := h.repo.allTodos()
	return util.Render(c, http.StatusOK, Index(todos))
}

// Get the Echo instance that was stored in the context with the SetEchoInstance middleware
// func GetEcho(ctx context.Context) *echo.Echo {
// 	return ctx.Value("echo").(*echo.Echo)
// }

