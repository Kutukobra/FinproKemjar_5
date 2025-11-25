package config

import (
	"errors"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	PostgresConnectionString string
	BcryptCost               int
	AppPort                  string
}

func Load() (*Config, error) {
	godotenv.Load()

	appPort := os.Getenv("APP_PORT")
	if appPort == "" {
		appPort = "3000"
	}

	pgConnectionString := os.Getenv("PG_URI")
	if pgConnectionString == "" {
		return nil, errors.New("config: no postgresql connection string found")
	}

	bcryptCost, err := strconv.Atoi(os.Getenv("BCRYPT_COST"))
	if err != nil {
		bcryptCost = 12 // Default
	}

	log.Println("ENVS: " + appPort + pgConnectionString)

	return &Config{
		PostgresConnectionString: pgConnectionString,
		BcryptCost:               bcryptCost,
		AppPort:                  appPort,
	}, nil
}
