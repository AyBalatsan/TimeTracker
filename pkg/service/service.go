package service

import (
	"github.com/AyBalatsan/TimeTracker/models"
	"github.com/AyBalatsan/TimeTracker/pkg/repository"
)

type Users interface {
	CreateUser(user models.People) (int, error)
	GetAllUser() ([]models.People, error)
}
type Tasks interface {
}

type Service struct {
	Users
	Tasks
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		Users: NewUsersService(repository.Users),
		Tasks: NewTasksService(repository.Users),
	}
}
