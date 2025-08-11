package middleware

import (
    "github.com/gofiber/fiber/v2"
    "github.com/google/uuid"

    "github.com/rs/zerolog/log"
)


type ctxKey string

const (
    CorrelationIDKey ctxKey = "correlationID"
    LoggerKey ctxKey = "logger"        
)

func LoggerMiddleware() fiber.Handler {
    return func(c *fiber.Ctx) error {
		correlationID := uuid.New().String()

        logger := log.With().Str(string(CorrelationIDKey), correlationID).Logger()

        c.Locals("logger", &logger)
        
        return c.Next()
    }
}
