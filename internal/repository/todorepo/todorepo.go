package todorepo

import (
	"todo/internal/model"

	"github.com/jmoiron/sqlx"
)

// In Go, interfaces are implemented implicitly
type TodoRepo interface {
	CreateTodo(t model.Todo) (int, error)
	GetTodo(id uint) model.Todo
}

type todoRepo struct {
	db *sqlx.DB
}

func NewTodoRepo(db *sqlx.DB) TodoRepo {
	return &todoRepo{
		db: db,
	}
}

// func (r todoRepo) CreateTodo(t model.Todo) (int, error) {
// 	res, err := r.db.Exec("INSERT INTO todo VALUES (NULL, ?, ?)", t.Title, t.Description)

// 	if err != nil {
// 		return 0, err
// 	}

// 	var id int64
// 	if id, err = res.LastInsertId(); err != nil {
// 		return 0, err
// 	}
// 	return int(id), nil
// }

func (r todoRepo) CreateTodo(t model.Todo) (int, error) {
	res, err := r.db.Exec("INSERT INTO todo VALUES (NULL, ?, ?)", t.Title, t.Description)

	if err != nil {
		return 0, err
	}

	var id int64
	if id, err = res.LastInsertId(); err != nil {
		return 0, err
	}
	return int(id), nil
}

func (r todoRepo) GetTodo(id uint) model.Todo {
	t := model.Todo{}
	// err := r.db.Get(&t, "SELECT * FROM todo t WHERE t.id = ?", id)
	r.db.Get(&t, "SELECT * FROM todo t WHERE t.id = ?", id)
	return t
}

func (r todoRepo) AllTodos() []model.Todo {
	todos := []model.Todo{}
	r.db.Select(&todos, "SELECT * FROM todo t")
	return todos
}
