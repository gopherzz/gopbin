package app

import (
	"github.com/gopherzz/gopbin/internal/handler"
	"github.com/gopherzz/gopbin/internal/services"
	"github.com/gopherzz/gopbin/repository"
)

type App struct {
	config *AppConfig
}

func (a *App) Start() error {
	repo := repository.NewRepository(a.config.DocumentsFolderPath)
	services := services.NewServices(repo)
	hand := handler.NewHandler(services)

	return hand.Listen(a.config.Port)
}

func NewApp(config *AppConfig) *App {
	return &App{
		config: config,
	}
}
