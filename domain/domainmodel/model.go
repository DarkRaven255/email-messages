package domainmodel

import (
	"crypto/tls"
	"email-messages/config"
	"email-messages/delivery/commands"
	"errors"
	"time"

	"github.com/gocql/gocql"
	"gopkg.in/gomail.v2"
)

type Model struct {
	Id          gocql.UUID `cql:"id"`
	Timestamp   time.Time  `cql:"timestamp"`
	Email       string     `cql:"email"`
	Title       string     `cql:"title"`
	Content     string     `cql:"content"`
	MagicNumber int        `cql:"magic_number"`
}

func NewDomainModel(cmd *commands.AddMessageCmd) (*Model, error) {

	if &cmd.Email == nil {
		return nil, errors.New("Email can not be empty!")
	}
	if &cmd.Title == nil {
		return nil, errors.New("Title can not be empty!")
	}
	if &cmd.Content == nil {
		return nil, errors.New("Content can not be empty!")
	}
	if &cmd.MagicNumber == nil {
		return nil, errors.New("MagicNumber can not be empty!")
	}

	return &Model{
		Email:       cmd.Email,
		Title:       cmd.Title,
		Content:     cmd.Content,
		MagicNumber: cmd.MagicNumber,
	}, nil
}

func (model *Model) SendEmail() error {
	m := gomail.NewMessage()
	m.SetHeader("From", config.Cfg.EmailUsername)
	m.SetHeader("To", model.Email)
	m.SetHeader("Subject", model.Title)
	m.SetBody("text/plain", model.Content)

	d := gomail.NewDialer(config.Cfg.EmailHost, config.Cfg.EmailPort, config.Cfg.EmailUsername, config.Cfg.EmailPassword)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}
