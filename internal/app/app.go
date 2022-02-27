package app

import (
	"github.com/gopherzz/gopbin/internal/handler"
	"github.com/gopherzz/gopbin/internal/repository"
	"github.com/gopherzz/gopbin/internal/services"
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
