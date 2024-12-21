package api

import (
	"mail/internal/config"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type Server struct {
	server   *echo.Echo
	address  string
	uc       Usecase
	validate *validator.Validate
	cfg      *config.Config
}

func NewServer(ip string, port int, userUc Usecase, cfg *config.Config) *Server {
	api := Server{
		uc:       userUc,
		validate: validator.New(),
		cfg:      cfg,
	}

	api.server = echo.New()
	api.server.POST("/mails", api.SendMail)
	api.server.GET("/mails/:user_id", api.GetMails)
	api.server.DELETE("/mails/:id", api.DeleteMail)

	api.address = ip + ":" + strconv.Itoa(port)

	return &api
}

func (s *Server) Run() {
	s.server.Logger.Fatal(s.server.Start(s.address))
}
