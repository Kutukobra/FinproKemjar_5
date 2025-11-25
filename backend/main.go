package main

import (
	"log"

	"github.com/Kutukobra/FinproKemjar_5/backend/app"
	"github.com/Kutukobra/FinproKemjar_5/backend/app/config"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Panic(err)
	}

	app, err := app.New(cfg)
	if err != nil {
		log.Panic(err)
	}

	router := gin.Default()
	app.Routes(router)

	router.Run(":" + cfg.AppPort)

	log.Println("App Running on Port :" + cfg.AppPort)
}
