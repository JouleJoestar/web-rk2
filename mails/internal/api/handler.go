package api

import (
	"mail/internal/config"
	"mail/internal/entities"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

var cfg *config.Config

// Обработчик для отправки письма
func (s *Server) SendMail(c echo.Context) error {
	var mail entities.Mail
	if err := c.Bind(&mail); err != nil {
		return c.String(http.StatusBadRequest, "Invalid input")
	}

	// Валидация структуры mail
	// В обработчике SendMail
	if err := mail.Validate(s.cfg); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	// Проверка существования отправителя
	if !s.uc.UserExists(mail.SenderID) {
		return c.String(http.StatusNotFound, entities.ErrUserNotFound.Error())
	}

	// Проверка существования получателей
	for _, receiverID := range mail.Receivers {
		if !s.uc.UserExists(receiverID) {
			return c.String(http.StatusNotFound, entities.ErrUserNotFound.Error())
		}
	}

	// Вызов метода отправки письма
	return c.String(http.StatusCreated, "Mail sent successfully")
}

// Обработчик для получения списка писем
func (s *Server) GetMails(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid user ID")
	}

	// Проверка существования пользователя
	if !s.uc.UserExists(userID) {
		return c.String(http.StatusNotFound, entities.ErrUserNotFound.Error())
	}

	mails, err := s.uc.GetMailsByUserID(userID)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, mails)
}

// Обработчик для удаления письма
func (s *Server) DeleteMail(c echo.Context) error {
	mailID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid mail ID")
	}

	// Проверка существования письма
	if !s.uc.MailExists(mailID) {
		return c.String(http.StatusNotFound, entities.ErrMailNotFound.Error())
	}

	if err := s.uc.DeleteMail(mailID); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.String(http.StatusOK, "Mail deleted successfully")
}
