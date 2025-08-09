package api

import (
	"net/http"

	"golang-eshop-backend/configs"
	"golang-eshop-backend/internal/helpers"
	"golang-eshop-backend/internal/api/rest"
	"golang-eshop-backend/internal/api/rest/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/rs/zerolog"
)

func StartServer(cfg configs.AppConfig, myLogger *zerolog.Logger) {
	app := fiber.New()

	// make fiber use our logger for requests logging
	app.Use(logger.New(logger.Config{
		Format: "ip=${ip} method=${method} path=${path} status=${status}",
		Output: helpers.LoggerWriter{Log: myLogger},
	}))

	// register routes
	rh := &rest.RestHandler{App: app}
	setupRoutes(rh)

	// health and readiness route
	app.Get("/health", healthCheck)

	// start server
	app.Listen(cfg.ServerPort)
}

func setupRoutes(rh *rest.RestHandler) {
	handlers.SetupUserRoutes(rh)
}

func healthCheck(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "healthy",
	})
}