package todo

import (
	"net/http"
	"todo/internal/model"
	"todo/internal/util"

	"github.com/labstack/echo/v4"
)

func (d database) getTodo(id int) model.Todo {
	var todo model.Todo
	d.db.Get(&todo, "SELECT * FROM todo t WHERE t.id = ?", id)
	return todo
}

type todoDetailsVM struct {
	ID   int `param:"id"`
	todo model.Todo
}

func (h handlers) detailsHandler(c echo.Context) error {
	var vm todoDetailsVM

	err1 := c.Bind(&vm)
	if err1 != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	// Fluent binding methods
	// err := echo.PathParamsBinder(c).
	// 	Int("id", &vm.id).
	// 	BindError()
	// if err != nil {
	// 	return c.String(http.StatusBadRequest, "bad request")
	// }

	vm.todo = h.repo.getTodo(vm.ID)
	return util.Render(c, http.StatusOK, Details(vm))
}
