package todo

import (
	"todo/internal/model"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

// Interfaces in Go are imlemented implicitly. Not like in Java where you have to explicitly say "implements Foo".
// The best approach will probably be to use an interface somehow like shown here  https://github.com/labstack/echo/issues/2075#issuecomment-1016819041
// Except I will want it to be more like a vertical slice pattern than a repository pattern

type todoRepo interface { // if you use interface in this package you can easily test handler with mocked db
	allTodos() []model.Todo
	getTodo(id int) model.Todo
}

type database struct {
	db *sqlx.DB
}

type handlers struct {
	repo todoRepo
}

func RegisterRoutes(parent *echo.Group, db *sqlx.DB) {
	h := handlers{
		repo: todoRepo(&database{db: db}),
	}

	g := parent.Group("/todo")

	g.GET("/", h.indexHandler).Name = "get-todos"
	g.GET("/:id", h.detailsHandler).Name = "get-todo"
}
