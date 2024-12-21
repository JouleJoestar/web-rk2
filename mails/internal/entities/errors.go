package entities

import (
	"errors"
	"mail/internal/config"
	"strconv"
)

var (
	ErrMessageLengthInvalid = errors.New("invalid lenght")
	ErrInvalidEmailFormat   = errors.New("invalid email format")
	ErrMailNotFound         = errors.New("mail not found")
	ErrUserNotFound         = errors.New("user not found")
)

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
