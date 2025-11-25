package config

import (
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
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	bcryptCost, _ := strconv.Atoi(os.Getenv("BCRYPT_COST"))

	return &Config{
		PostgresConnectionString: os.Getenv("PG_CONN_STRING"),
		BcryptCost:               bcryptCost,
	}, nil
}
