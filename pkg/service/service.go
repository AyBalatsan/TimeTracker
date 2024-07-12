package service

import "github.com/AyBalatsan/TimeTracker/pkg/repository"

type Service struct {
}

func NewService(repository *repository.Repository) *Service {
	return &Service{}
}
