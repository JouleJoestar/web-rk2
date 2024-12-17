package usecase

import "mail/internal/entities"

// Provider интерфейс для работы с данными
type Provider interface {
	SendMail(mail entities.Mail) error
	GetMailsByUserID(userID int) ([]entities.Mail, error)
	DeleteMail(mailID int) error
	UserExist(userID int) bool
	MailExist(mailID int) bool
}
