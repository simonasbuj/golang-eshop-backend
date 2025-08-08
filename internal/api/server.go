package api

import (
	"net/http"
	"strings"

	"golang-eshop-backend/configs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/rs/zerolog"
)

func StartServer(cfg configs.AppConfig, myLogger *zerolog.Logger) {
	app := fiber.New()

	// make fiber use our logger for requests logging
	app.Use(logger.New(logger.Config{
		Format: "ip=${ip} method=${method} path=${path} status=${status}",
		Output: loggerWriter{log: myLogger},
	}))

	// register routes
	app.Get("/health", healthCheck)

	// start server
	app.Listen(cfg.ServerPort)
}

type loggerWriter struct {
	log *zerolog.Logger
}

func (lw loggerWriter) Write(p []byte) (n int, err error) {
	line := strings.TrimSpace(string(p))

    fields := map[string]string{}
    for _, part := range strings.Split(line, " ") {
        kv := strings.SplitN(part, "=", 2)
        if len(kv) == 2 {
            fields[kv[0]] = kv[1]
        }
    }

    lw.log.Info().
        Str("ip", fields["ip"]).
        Str("method", fields["method"]).
        Str("path", fields["path"]).
        Str("response_status", fields["status"]).
        Msg("Request")
		
    return len(p), nil
}

func healthCheck(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "healthy",
	})
}