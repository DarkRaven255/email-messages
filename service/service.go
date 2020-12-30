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

	results, err := ms.messagesRepo.GetByEmail(email)
	if err != nil {
		return nil, err
	}

	return results, nil
}

func NewMessagesService(mr domain.MessagesRepository) domain.MessagesService {
	ms := &messagesService{
		messagesRepo: mr,
	}

	return ms
}
