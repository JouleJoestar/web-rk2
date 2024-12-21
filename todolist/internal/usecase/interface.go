package usecase

import "todolist/internal/entities"

type Provider interface {
	InsertTask(entities.Task) (*entities.Task, error)
	SelectAllTasks() ([]*entities.Task, error)
	UpdateTaskStatus(id int, status string) (*entities.Task, error)
	SelectTaskByID(id int) (*entities.Task, error)
}
