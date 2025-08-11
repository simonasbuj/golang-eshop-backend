package api

import (
	"net/http"

	"golang-eshop-backend/configs"
	"golang-eshop-backend/internal/api/rest"
	"golang-eshop-backend/internal/api/rest/handlers"
	"golang-eshop-backend/internal/api/rest/middleware"

	"github.com/gofiber/fiber/v2"
)

func StartServer(cfg configs.AppConfig) {
	app := fiber.New()

	// register middleware
	app.Use(middleware.LoggerWithCommonValuesMiddleware())
	app.Use(middleware.RequestLoggerMiddleware())


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