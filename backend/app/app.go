package app

import (
	"github.com/Kutukobra/FinproKemjar_5/backend/app/config"
	"github.com/Kutukobra/FinproKemjar_5/backend/app/database"
	"github.com/Kutukobra/FinproKemjar_5/backend/app/handler"
	"github.com/Kutukobra/FinproKemjar_5/backend/app/repository"
	"github.com/Kutukobra/FinproKemjar_5/backend/app/service"
)

type App struct {
	userHandler *handler.UserHandler
	pageHandler *handler.PageHandler
}

func New(cfg *config.Config) (*App, error) {
	pgDatabase, err := database.NewPostgresDatabase(cfg.PostgresConnectionString)
	if err != nil {
		return nil, err
	}

	userRepository := repository.NewPGUserRepository(pgDatabase)
	userService := service.NewUserService(userRepository, cfg.BcryptCost)

	return &App{
		userHandler: handler.NewUserHandler(userService),
		pageHandler: handler.NewPageHandler(userService),
	}, nil
}
