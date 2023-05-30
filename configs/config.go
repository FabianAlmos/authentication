package configs

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Port                        string
	AccessTokenSecret           string
	AccessTokenLifetimeMinutes  int
	RefreshTokenSecret          string
	RefreshTokenLifetimeMinutes int
}

func NewConfig() *Config {
	err := godotenv.Load("configs/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	ATLM, err := strconv.Atoi(os.Getenv("ACCESS_TOKEN_LIFETIME_MINUTES"))
	if err != nil {
		log.Fatal("Error converting ENVVAR to int")
	}
	RTLM, err := strconv.Atoi(os.Getenv("REFRESH_TOKEN_LIFETIME_MINUTES"))
	if err != nil {
		log.Fatal("Error converting ENVVAR to int")
	}

	return &Config{
		Port:                        os.Getenv("PORT"),
		AccessTokenSecret:           os.Getenv("ACCESS_TOKEN_SECRET"),
		AccessTokenLifetimeMinutes:  ATLM,
		RefreshTokenSecret:          os.Getenv("REFRESH_TOKEN_SECRET"),
		RefreshTokenLifetimeMinutes: RTLM,
	}
}
