package api

import (
	"log"
	"net/http"

	"golang-eshop-backend/config"
	"golang-eshop-backend/internal/api/rest"
	"golang-eshop-backend/internal/api/rest/handlers"
	"golang-eshop-backend/internal/api/rest/middleware"
	"golang-eshop-backend/internal/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/driver/postgres"
	
)

func StartServer(cfg config.AppConfig) {
	app := fiber.New()

	// register middleware
	app.Use(middleware.LoggerWithCommonValuesMiddleware())
	app.Use(middleware.RequestLoggerMiddleware())

	// init db connection and run migrations
	db, err := gorm.Open(postgres.Open(cfg.Dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent), // completely silent
	})
	if err != nil {
		log.Fatalf("database connection error")
	}
	db.AutoMigrate(&models.User{})


	// register routes
	rh := &rest.RestHandler{
		App: app,
		DB: db,
	}
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