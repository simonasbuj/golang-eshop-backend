package middleware

import (
    "context"

    "github.com/gofiber/fiber/v2"
    "github.com/google/uuid"
)

func CorrelationIDMiddleware() fiber.Handler {
    return func(c *fiber.Ctx) error {
		correlationID := uuid.New().String()

        ctx := context.WithValue(c.UserContext(), CorrelationIDKey, correlationID)
        c.SetUserContext(ctx)

        return c.Next()
    }
}
