package api

import (
	"mail/internal/entities"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (s *Server) SendMail(c echo.Context) error {
	var mail entities.Mail
	if err := c.Bind(&mail); err != nil {
		return c.String(http.StatusBadRequest, "Invalid input")
	}

	if err := mail.Validate(s.cfg); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	if !s.uc.UserExists(mail.SenderID) {
		return c.String(http.StatusNotFound, entities.ErrUserNotFound.Error())
	}

	for _, receiverID := range mail.Receivers {
		if !s.uc.UserExists(receiverID) {
			return c.String(http.StatusNotFound, entities.ErrUserNotFound.Error())
		}
	}

	if err := s.uc.SendMail(mail); err != nil {
		return c.String(http.StatusInternalServerError, "Failed to send mail: "+err.Error())
	}

	return c.String(http.StatusCreated, "Mail sent successfully")
}

func (s *Server) GetMails(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid user ID")
	}

	if !s.uc.UserExists(userID) {
		return c.String(http.StatusNotFound, entities.ErrUserNotFound.Error())
	}

	mails, err := s.uc.GetMailsByUserID(userID)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	for i := range mails {
		mails[i].Receivers = nil
	}

	return c.JSON(http.StatusOK, mails)
}

func (s *Server) DeleteMail(c echo.Context) error {
	mailID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid mail ID")
	}

	if !s.uc.MailExists(mailID) {
		return c.String(http.StatusNotFound, entities.ErrMailNotFound.Error())
	}

	if err := s.uc.DeleteMail(mailID); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.String(http.StatusOK, "Mail deleted successfully")
}
