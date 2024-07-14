package service

import "github.com/AyBalatsan/TimeTracker/pkg/repository"

type TasksService struct {
	repo repository.Tasks
}

func NewTasksService(repo repository.Tasks) *TasksService {
	return &TasksService{repo: repo}
}
