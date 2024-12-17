package entities

import (
	"errors"
	"mail/internal/config"
	"strconv"
)

var (
	ErrMessageLengthInvalid = errors.New("message length must be between 5 and 1000 characters")
	ErrInvalidEmailFormat   = errors.New("invalid email format")
	ErrMailNotFound         = errors.New("mail not found")
	ErrUserNotFound         = errors.New("user not found")
)

// Validate проверяет корректность полей структуры Mail
func (m *Mail) Validate(cfg *config.Config) error {
	if len(m.Theme) < cfg.MailThemeMinLen || len(m.Theme) > cfg.MailThemeMaxLen {
		return errors.New("theme length must be between " + strconv.Itoa(cfg.MailThemeMinLen) + " and " + strconv.Itoa(cfg.MailThemeMaxLen) + " characters")
	}
	if len(m.Text) < cfg.MailTextMinLen || len(m.Text) > cfg.MailTextMaxLen {
		return errors.New("text length must be between " + strconv.Itoa(cfg.MailTextMinLen) + " and " + strconv.Itoa(cfg.MailTextMaxLen) + " characters")
	}
	if m.SenderID <= 0 {
		return ErrUserNotFound
	}
	if len(m.Receivers) == 0 {
		return ErrUserNotFound
	}
	return nil
}
