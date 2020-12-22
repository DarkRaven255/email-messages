package service

import (
	"email-messages/delivery/commands"
	"email-messages/domain"
	"email-messages/domain/domainmodel"
)

type messagesService struct {
	messagesRepo domain.MessagesRepository
}

func (ms *messagesService) AddMessage(cmd *commands.AddMessageCmd) error {
	return nil
}

func (ms *messagesService) SendMessages(cmd *commands.SendMessagesCmd) error {
	return nil
}

func (ms *messagesService) GetMessages(email *string) ([]*domainmodel.Model, error) {
	return nil, nil
}

func NewMessagesService(mr domain.MessagesRepository) domain.MessagesRepository {
	ms := &messagesService{
		messagesRepo: mr,
	}

	return ms
}
