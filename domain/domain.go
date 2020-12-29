package domain

import (
	"email-messages/delivery/commands"
	"email-messages/domain/domainmodel"

	"github.com/gocql/gocql"
)

type MessagesRepository interface {
	Create(entry *domainmodel.Model) error
	GetByMagicNumber(magicNumber *int) (*[]domainmodel.Model, error)
	GetByEmail(email *string) (*[]domainmodel.Model, error)
	Delete(id *gocql.UUID) error
}

type MessagesService interface {
	AddMessage(addMessageCmd *commands.AddMessageCmd) error
	SendMessages(sendMessagesCmd *commands.SendMessagesCmd) error
	GetMessages(email *string) (*[]domainmodel.Model, error)
}
