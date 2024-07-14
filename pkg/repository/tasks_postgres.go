package repository

import "github.com/go-sqlx/sqlx"

type TasksPostgres struct {
	db *sqlx.DB
}

func NewTasksPostgres(db *sqlx.DB) *TasksPostgres {
	return &TasksPostgres{db: db}
}
