package domain

import (
	"email-messages/delivery/commands"
	"email-messages/domain/domainmodel"
)

type MessagesRepository interface {
	Create(entry *domainmodel.Model) error
	GetByMagicNumber(magicNumber *string) ([]*domainmodel.Model, error)
	Delete() error
}

type MessagesService interface {
	AddMessage(addMessageCmd *commands.AddMessageCmd) error
	SendMessages(sendMessagesCmd *commands.SendMessagesCmd) error
	GetMessages(email *string) ([]*domainmodel.Model, error)
}
