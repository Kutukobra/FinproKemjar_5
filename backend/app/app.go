package app

import (
	"github.com/Kutukobra/FinproKemjar_5/backend/app/config"
	"github.com/Kutukobra/FinproKemjar_5/backend/app/database"
	"github.com/Kutukobra/FinproKemjar_5/backend/app/handler"
	"github.com/Kutukobra/FinproKemjar_5/backend/app/repository"
	"github.com/Kutukobra/FinproKemjar_5/backend/app/service"
	"github.com/gin-gonic/gin"
)

type App struct {
	userHandler *handler.UserHandler
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
	}, nil
}

func (a *App) Routes(router *gin.Engine) {
	user := router.Group("/user")
	{
		user.GET("/")
		user.POST("/register")
		user.POST("/login")
		user.PUT("/change-password")
	}
}
