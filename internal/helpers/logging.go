package helpers

import (
	"strings"

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