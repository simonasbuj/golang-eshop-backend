package logging

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)


func GetLoggerFromCtx(c *fiber.Ctx) *zerolog.Logger {
    logger := c.Locals("logger").(*zerolog.Logger)
    return logger
}
