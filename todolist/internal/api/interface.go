package api

import "todolist/internal/entities"

type Usecase interface {
	CreateTask(task entities.Task) (*entities.Task, error)
	ListTasks() ([]*entities.Task, error)
	UpdateTaskStatus(id int, status string) (*entities.Task, error)
}
