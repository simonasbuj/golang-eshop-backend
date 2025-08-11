package api

import (
	"net/http"

	"golang-eshop-backend/configs"
	"golang-eshop-backend/internal/api/rest"
	"golang-eshop-backend/internal/api/rest/handlers"
	"golang-eshop-backend/internal/api/rest/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

func StartServer(cfg configs.AppConfig, myLogger *zerolog.Logger) {
	app := fiber.New()

	// register middleware
	app.Use(middleware.LoggerWithCommonValuesMiddleware())
	app.Use(middleware.RequestLoggerMiddleware())


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