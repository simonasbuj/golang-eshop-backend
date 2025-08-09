package middleware

import (
    "context"
    "github.com/gofiber/fiber/v2"
    "github.com/google/uuid"
)

const correlationIDKey = "correlation_id"

func CorrelationIDMiddleware() fiber.Handler {
    return func(c *fiber.Ctx) error {
		correlationID := uuid.New().String()

        ctx := context.WithValue(c.UserContext(), correlationIDKey, correlationID)
        c.SetUserContext(ctx)

        return c.Next()
    }
}

func CorrelationIDFromCtx(ctx context.Context) string {
    if v := ctx.Value(correlationIDKey); v != nil {
        if id, ok := v.(string); ok {
            return id
        }
    }
    return ""
}