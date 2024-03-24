package repository

import (
	"database/sql"
	"todo/internal/repository/todorepo"
)

type Repositories struct {
	TodoRepo todorepo.TodoRepo
}

func InitRepositories(db *sql.DB) *Repositories {
	return &Repositories{
		TodoRepo: todorepo.NewTodoRepo(db),
	}
}
