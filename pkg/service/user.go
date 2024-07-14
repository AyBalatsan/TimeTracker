package service

import (
	"github.com/AyBalatsan/TimeTracker/models"
	"github.com/AyBalatsan/TimeTracker/pkg/repository"
)

const salt = "fdklskl324kfdo" //TODO вынести в .ENV

type UsersService struct {
	repo repository.Users
}

func NewUsersService(repo repository.Users) *UsersService {
	return &UsersService{repo: repo}
}

func (s *UsersService) CreateUser(user models.People) (int, error) {
	return s.repo.CreateUser(user)
}
func (s *UsersService) GetAllUser() ([]models.People, error) {
	return s.repo.GetAllUser()
}
