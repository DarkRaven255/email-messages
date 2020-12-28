package app

import "email-messages/domain"

type App struct {
	MessagesService domain.MessagesService
}

func NewApp(ms domain.MessagesService) *App {
	return &App{
		MessagesService: ms,
	}
}
