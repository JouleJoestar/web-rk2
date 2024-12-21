package usecase

import "mail/internal/entities"

type Provider interface {
	SendMail(mail entities.Mail) error
	GetMailsByUserID(userID int) ([]entities.Mail, error)
	DeleteMail(mailID int) error
	UserExist(userID int) bool
	MailExist(mailID int) bool
}
