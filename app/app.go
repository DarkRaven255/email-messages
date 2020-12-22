package app

import "email-messages/domain"

type App struct {
	MessagesService domain.MessagesService
}

func NewApp(es domain.MessagesService) *App {
	return &App{
		MessagesService: es,
	}
}
