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
	uc       Usecase // Интерфейс для работы с пользователями
	validate *validator.Validate
	cfg      *config.Config // Добавлено поле для конфигурации
}

func NewServer(ip string, port int, userUc Usecase, cfg *config.Config) *Server {
	api := Server{
		uc:       userUc,
		validate: validator.New(),
		cfg:      cfg, // Инициализация валидатора
	}

	api.server = echo.New()
	api.server.POST("/mails", api.SendMail)         // Отправка письма
	api.server.GET("/mails/:user_id", api.GetMails) // Получение списка писем для пользователя
	api.server.DELETE("/mails/:id", api.DeleteMail) // Удаление письма

	api.address = ip + ":" + strconv.Itoa(port)

	return &api
}

func (s *Server) Run() {
	s.server.Logger.Fatal(s.server.Start(s.address))
}
