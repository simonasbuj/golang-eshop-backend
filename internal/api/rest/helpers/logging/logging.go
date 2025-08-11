package logging

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)


func GetLoggerFromCtx(c *fiber.Ctx) *zerolog.Logger {
	logger, ok := c.Locals("logger").(*zerolog.Logger)
	if !ok {
		noOpLogger := zerolog.Nop()
		logger = &noOpLogger
	}
    return logger
}
