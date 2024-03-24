package todorepo

import (
	"database/sql"
	"todo/internal/model"
)

type TodoRepo interface {
	CreateTodo(t model.Todo) (int, error)
}

type todoRepo struct {
	db *sql.DB
}

func NewTodoRepo(db *sql.DB) TodoRepo {
	return &todoRepo{
		db: db,
	}
}

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
