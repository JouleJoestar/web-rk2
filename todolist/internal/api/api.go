package api

import (
	"strconv"

	"github.com/labstack/echo/v4"
)

type Server struct {
	server  *echo.Echo
	address string
	uc      Usecase
}

func NewServer(ip string, port int, uc Usecase) *Server {
	srv := &Server{
		uc:      uc,
		server:  echo.New(),
		address: ip + ":" + strconv.Itoa(port),
	}

	srv.server.POST("/tasks", srv.CreateTask)
	srv.server.GET("/tasks", srv.ListTasks)
	srv.server.PUT("/tasks/status", srv.UpdateTaskStatus)

	return srv
}

func (s *Server) Run() {
	s.server.Logger.Fatal(s.server.Start(s.address))
}
