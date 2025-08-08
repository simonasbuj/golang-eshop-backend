package configs

import(
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	ServerPort string
}

func SetupEnv() (AppConfig, error) {
	if os.Getenv("APP_ENV") == "local" {
		godotenv.Load()
	}

	serverPort := os.Getenv("SERVER_PORT")
	if len(serverPort) < 1 {
		return AppConfig{}, errors.New("serverPort env variable not found")
	}

	return AppConfig{ ServerPort: serverPort }, nil
}
