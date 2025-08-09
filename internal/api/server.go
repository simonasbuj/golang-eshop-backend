package api

import (
	"net/http"

	"golang-eshop-backend/configs"
	"golang-eshop-backend/internal/api/rest"
	"golang-eshop-backend/internal/api/rest/handlers"
	"golang-eshop-backend/internal/api/rest/middleware"
	"golang-eshop-backend/internal/helpers/logging"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/rs/zerolog"
)

func StartServer(cfg configs.AppConfig, myLogger *zerolog.Logger) {
	app := fiber.New()


	// middleware
	// make fiber use our logger for requests logging
	app.Use(logger.New(logger.Config{
		Format: "ip=${ip} method=${method} path=${path} status=${status}",
		Output: logging.LoggerWriter{Log: myLogger},
	}))

	// add correlation_id to context for each request
	app.Use(middleware.CorrelationIDMiddleware())


	// register routes
	rh := &rest.RestHandler{App: app}
	setupRoutes(rh, myLogger)

	// health and readiness route
	app.Get("/health", healthCheck)

	// start server
	app.Listen(cfg.ServerPort)
}

func setupRoutes(rh *rest.RestHandler, logger *zerolog.Logger) {
	handlers.SetupUserRoutes(rh, logger)
}

func healthCheck(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "healthy",
	})
}