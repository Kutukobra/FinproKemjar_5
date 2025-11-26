package config

import (
	"errors"
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

	appPort := os.Getenv("PORT")
	if appPort == "" {
		appPort = ":8080"
	}

	pgConnectionString := os.Getenv("PG_STRING")
	if pgConnectionString == "" {
		return nil, errors.New("config: no postgresql connection string found")
	}

	bcryptCost, err := strconv.Atoi(os.Getenv("BCRYPT_COST"))
	if err != nil {
		bcryptCost = 12 // Default
	}

	return &Config{
		PostgresConnectionString: pgConnectionString,
		BcryptCost:               bcryptCost,
		AppPort:                  appPort,
	}, nil
}
