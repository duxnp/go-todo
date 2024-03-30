package repository

import (
	"todo/internal/repository/todorepo"

	"github.com/jmoiron/sqlx"
)

type Repositories struct {
	TodoRepo todorepo.TodoRepo
}

func InitRepositories(db *sqlx.DB) *Repositories {
	return &Repositories{
		TodoRepo: todorepo.NewTodoRepo(db),
	}
}
