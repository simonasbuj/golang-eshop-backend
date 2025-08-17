package main

import (
	"os"

	"golang-eshop-backend/config"
	"golang-eshop-backend/internal/api"

	"github.com/rs/zerolog"
)

func main() {

	// init logger
	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()

	// load settings
	cfg, err := config.SetupEnv()
	if err != nil {
		logger.Fatal().Err(err).Msg("Couldn't initialize app config")
	}

	// start server
	logger.Info().Msg("Starting server...")
	api.StartServer(cfg)
}
