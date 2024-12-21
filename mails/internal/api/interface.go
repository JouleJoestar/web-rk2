package api

import "mail/internal/entities"

type Usecase interface {
	SendMail(mail entities.Mail) error
	GetMailsByUserID(userID int) ([]entities.Mail, error)
	DeleteMail(mailID int) error
	UserExists(userID int) bool
	MailExists(mailID int) bool
}
