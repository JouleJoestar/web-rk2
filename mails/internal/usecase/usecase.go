package usecase

import (
	"mail/internal/entities"
)

type Usecase struct {
	p Provider
}

func NewUsecase(p Provider) *Usecase {
	return &Usecase{p: p}
}

// Метод для отправки письма
func (u *Usecase) SendMail(mail entities.Mail) error {
	// Проверка существования отправителя
	if !u.UserExists(mail.SenderID) {
		return entities.ErrUserNotFound
	}

	// Проверка существования получателей
	for _, receiverID := range mail.Receivers {
		if !u.UserExists(receiverID) {
			return entities.ErrUserNotFound
		}
	}

	// Вызов метода отправки письма
	return u.p.SendMail(mail)
}

// Метод для получения списка писем для пользователя
func (u *Usecase) GetMailsByUserID(userID int) ([]entities.Mail, error) {
	// Проверка существования пользователя
	if !u.UserExists(userID) {
		return nil, entities.ErrUserNotFound
	}
	return u.p.GetMailsByUserID(userID)
}

// Метод для удаления письма
func (u *Usecase) DeleteMail(mailID int) error {
	// Проверка существования письма
	if !u.MailExists(mailID) {
		return entities.ErrMailNotFound
	}
	return u.p.DeleteMail(mailID)
}

// Метод для проверки существования пользователя
func (u *Usecase) UserExists(userID int) bool {
	// Вызов метода провайдера для проверки существования пользователя
	return u.p.UserExist(userID)
}

// Метод для проверки существования письма
func (u *Usecase) MailExists(mailID int) bool {
	// Вызов метода провайдера для проверки существования письма
	return u.p.MailExist(mailID)
}
