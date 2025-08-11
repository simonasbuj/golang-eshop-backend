package middleware

import (
    "time"

    "github.com/gofiber/fiber/v2"
    "github.com/google/uuid"
    "github.com/rs/zerolog"
    "github.com/rs/zerolog/log"
)


type ctxKey string

const (
    CorrelationIDKey ctxKey = "correlationID"
    LoggerKey ctxKey = "logger"        
)

func LoggerWithCommonValuesMiddleware() fiber.Handler {
    return func(c *fiber.Ctx) error {
		correlationID := uuid.New().String()

        logger := log.With().
            Str(string(CorrelationIDKey), correlationID).
            Logger()

        c.Locals("logger", &logger)
        
        return c.Next()
    }
}

func RequestLoggerMiddleware() fiber.Handler {
    return func(c *fiber.Ctx) error {
        start := time.Now()
        err := c.Next()
        duration := time.Since(start)

        logger, ok := c.Locals("logger").(*zerolog.Logger)
        if !ok {
            noOpLogger := zerolog.Nop()
            logger = &noOpLogger
        }

        logger.Info().
            Str("ip", c.IP()).
            Str("method", c.Method()).
            Str("path", c.Path()).
            Int("status", c.Response().StatusCode()).
            Dur("duration_ms", duration).
            Msg("Request")

        return err
    }
}

