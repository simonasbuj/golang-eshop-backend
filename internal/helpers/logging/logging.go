package logging

import (
	"golang-eshop-backend/internal/api/rest/middleware"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

type LoggerWriter struct {
	Log *zerolog.Logger
}

func (lw LoggerWriter) Write(p []byte) (n int, err error) {
	line := strings.TrimSpace(string(p))

    fields := map[string]string{}
    for _, part := range strings.Split(line, " ") {
        kv := strings.SplitN(part, "=", 2)
        if len(kv) == 2 {
            fields[kv[0]] = kv[1]
        }
    }

    lw.Log.Info().
        Str("ip", fields["ip"]).
        Str("method", fields["method"]).
        Str("path", fields["path"]).
        Str("response_status", fields["status"]).
        Msg("Request")
		
    return len(p), nil
}


func putCtxValuesIntoLogger(logger *zerolog.Logger, ctx *fiber.Ctx) *zerolog.Logger {

    corrID := string(middleware.CorrelationIDKey)

    l := logger.With().Str(corrID, getCtxValue(ctx, corrID)).Logger()
    return &l
}

func getCtxValue(c *fiber.Ctx, valueName string) string {
    if id, ok := c.UserContext().Value(valueName).(string); ok {
        return id
    }
    return ""
}

func LogInfo(logger *zerolog.Logger, ctx *fiber.Ctx, msg string) {
    l := putCtxValuesIntoLogger(logger, ctx)
    l.Info().Msg(msg)
}

func LogError(logger *zerolog.Logger, ctx *fiber.Ctx, err error, msg string) {
    l := putCtxValuesIntoLogger(logger, ctx)
    l.Error().Err(err).Msg(msg)
}

func GetLoggerFromCtx(c *fiber.Ctx) *zerolog.Logger {
    logger := c.Locals("logger").(*zerolog.Logger)
    return logger
}