package repository

import (
	"github.com/AyBalatsan/TimeTracker/models"
	"github.com/go-sqlx/sqlx"
)

type Users interface {
	GetAllUser() ([]models.People, error)
	CreateUser(user models.People) (int, error)
}
type Tasks interface{}
type Repository struct {
	Users
	Tasks
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Users: NewUserPostgres(db),
		Tasks: NewTasksPostgres(db),
	}
}
