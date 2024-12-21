package usecase

import (
	"todolist/internal/entities"
)

type Usecase struct {
	p Provider
}

func NewUsecase(p Provider) *Usecase {
	return &Usecase{p: p}
}

func (u *Usecase) CreateTask(task entities.Task) (*entities.Task, error) {
	createdTask, err := u.p.InsertTask(task)
	if err != nil {
		return nil, err
	}
	return createdTask, nil
}

func (u *Usecase) ListTasks() ([]*entities.Task, error) {
	tasks, err := u.p.SelectAllTasks()
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (u *Usecase) UpdateTaskStatus(id int, status string) (*entities.Task, error) {
	task, err := u.p.SelectTaskByID(id)
	if err != nil {
		return nil, err
	}
	if task == nil {
		return nil, entities.ErrTaskNotFound
	}
	if task.Status == "done" {
		return nil, entities.ErrTaskAlreadyResolved
	}
	if (task.Status == "new" && status != "in progress" && status != "done") || (task.Status == "in progress" && status != "done") {
		return nil, entities.ErrInvalidTaskStatus
	}
	updatedTask, err := u.p.UpdateTaskStatus(id, status)
	if err != nil {
		return nil, err
	}
	return updatedTask, nil
}
