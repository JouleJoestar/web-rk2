package api

import (
	"net/http"
	"strconv"
	"todolist/internal/entities"

	"github.com/labstack/echo/v4"
)

func (s *Server) CreateTask(c echo.Context) error {
	var task entities.Task
	if err := c.Bind(&task); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	createdTask, err := s.uc.CreateTask(task)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, createdTask)
}

func (s *Server) ListTasks(c echo.Context) error {
	tasks, err := s.uc.ListTasks()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, tasks)
}

func (s *Server) UpdateTaskStatus(c echo.Context) error {
	idStr := c.QueryParam("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.String(http.StatusBadRequest, "invalid task ID")
	}
	var status struct {
		Status string `json:"status"`
	}
	if err := c.Bind(&status); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	updatedTask, err := s.uc.UpdateTaskStatus(id, status.Status)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, updatedTask)
}
