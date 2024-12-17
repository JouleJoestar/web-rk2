package provider

import (
	"database/sql"
	"fmt"
	"log"
	"mail/internal/entities"
)

type Provider struct {
	conn *sql.DB
}

func NewProvider(host string, port int, user, password, dbName string) *Provider {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbName)

	// Создание соединения с сервером postgres
	conn, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	return &Provider{conn: conn}
}

// Метод для отправки письма
func (p *Provider) SendMail(mail entities.Mail) error {
	for _, receiverID := range mail.Receivers {
		_, err := p.conn.Exec(`INSERT INTO mails (theme, text, image, id_sender, id_receiver) VALUES ($1, $2, $3, $4, $5)`,
			mail.Theme, mail.Text, mail.Image, mail.SenderID, receiverID)
		if err != nil {
			return err
		}
	}
	return nil
}

// Метод для получения списка писем для пользователя
func (p *Provider) GetMailsByUserID(userID int) ([]entities.Mail, error) {
	mails := []entities.Mail{}
	rows, err := p.conn.Query(`SELECT id, theme, text, image, id_sender FROM mails WHERE id_receiver = $1`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var mail entities.Mail
		if err := rows.Scan(&mail.ID, &mail.Theme, &mail.Text, &mail.Image, &mail.SenderID); err != nil {
			return nil, err
		}
		mails = append(mails, mail)
	}
	return mails, nil
}

// Метод для удаления письма
func (p *Provider) DeleteMail(mailID int) error {
	_, err := p.conn.Exec(`DELETE FROM mails WHERE id = $1`, mailID)
	return err
}

// Метод для проверки существования пользователя
func (p *Provider) UserExist(userID int) bool {
	var exists bool
	err := p.conn.QueryRow(`SELECT EXISTS(SELECT 1 FROM users WHERE id = $1)`, userID).Scan(&exists)
	if err != nil {
		log.Println("Error checking user existence:", err)
		return false
	}
	return exists
}

// Метод для проверки существования письма
func (p *Provider) MailExist(mailID int) bool {
	var exists bool
	err := p.conn.QueryRow(`SELECT EXISTS(SELECT 1 FROM mails WHERE id = $1)`, mailID).Scan(&exists)
	if err != nil {
		log.Println("Error checking mail existence:", err)
		return false
	}
	return exists
}
