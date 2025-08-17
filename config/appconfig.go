package config

import(
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	ServerPort 	string
	Dsn			string
}

func SetupEnv() (AppConfig, error) {
	if os.Getenv("APP_ENV") == "local" {
		godotenv.Load()
	}

	serverPort := os.Getenv("SERVER_PORT")
	if len(serverPort) < 1 {
		return AppConfig{}, errors.New("SERVER_PORT env variable not found")
	}

	dsn := os.Getenv("DSN")
	if len(serverPort) < 1 {
		return AppConfig{}, errors.New("DSN env variable not found")
	}

	return AppConfig{ ServerPort: serverPort, Dsn: dsn }, nil
}
