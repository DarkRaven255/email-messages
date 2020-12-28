package domainmodel

import (
	"email-messages/delivery/commands"
	"errors"
	"time"

	"github.com/gocql/gocql"
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
