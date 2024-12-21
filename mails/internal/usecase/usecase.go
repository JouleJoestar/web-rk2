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

func (u *Usecase) SendMail(mail entities.Mail) error {
	if !u.UserExists(mail.SenderID) {
		return entities.ErrUserNotFound
	}

	for _, receiverID := range mail.Receivers {
		if !u.UserExists(receiverID) {
			return entities.ErrUserNotFound
		}
	}

	return u.p.SendMail(mail)
}

func (u *Usecase) GetMailsByUserID(userID int) ([]entities.Mail, error) {
	if !u.UserExists(userID) {
		return nil, entities.ErrUserNotFound
	}
	return u.p.GetMailsByUserID(userID)
}

func (u *Usecase) DeleteMail(mailID int) error {
	if !u.MailExists(mailID) {
		return entities.ErrMailNotFound
	}
	return u.p.DeleteMail(mailID)
}

func (u *Usecase) UserExists(userID int) bool {
	return u.p.UserExist(userID)
}

func (u *Usecase) MailExists(mailID int) bool {
	return u.p.MailExist(mailID)
}
