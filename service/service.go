package service

import (
	"email-messages/delivery/commands"
	"email-messages/domain"
	"email-messages/domain/domainmodel"
	"errors"
)

type messagesService struct {
	messagesRepo domain.MessagesRepository
}

func (ms *messagesService) AddMessage(cmd *commands.AddMessageCmd) error {

	message, err := domainmodel.NewDomainModel(cmd)
	if err != nil {
		return err
	}

	err = ms.messagesRepo.Create(message)
	if err != nil {
		return err
	}

	return nil
}

func (ms *messagesService) SendMessages(cmd *commands.SendMessagesCmd) error {

	if &cmd.MagicNumber == nil {
		return errors.New("MagicNumber can not be empty!")
	}

	results, err := ms.messagesRepo.GetByMagicNumber(&cmd.MagicNumber)
	if err != nil {
		return err
	}

	for _, result := range *results {
		err := result.SendEmail()
		if err != nil {
			return err
		}
		err = ms.messagesRepo.Delete(&result.Id)
		if err != nil {
			return err
		}
	}

	return nil
}

func (ms *messagesService) GetMessages(email *string) (*[]domainmodel.Model, error) {
	return nil, nil
}

func NewMessagesService(mr domain.MessagesRepository) domain.MessagesService {
	ms := &messagesService{
		messagesRepo: mr,
	}

	return ms
}
