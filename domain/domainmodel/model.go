package domainmodel

import (
	"email-messages/delivery/commands"
	"time"
)

type Model struct {
	ID          int       `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	Email       string    `json:"email"`
	Title       string    `json:"title"`
	Content     string    `json:"content"`
	MagicNumber int       `json:"magic_number"`
}

func NewDomainModel(cmd *commands.AddMessageCmd) *Model {
	return nil
}
