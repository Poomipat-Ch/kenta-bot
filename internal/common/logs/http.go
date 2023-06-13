package logs

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog"
)

func NewStructuredLogger(logger *zerolog.Logger) func(next http.Handler) http.Handler {
	return middleware.RequestLogger(&StructuredLogger{logger})
}

type StructuredLogger struct {
	Logger *zerolog.Logger
}

// NewLogEntry implements middleware.LogFormatter.
func (l *StructuredLogger) NewLogEntry(r *http.Request) middleware.LogEntry {
	entry := &StructuredLoggerEntry{Logger: zerolog.New(l.Logger)}

	logContext := entry.Logger.With()

	if reqID := middleware.GetReqID(r.Context()); reqID != "" {
		logContext.Str("request_id", reqID)
	}

	logContext.Str("http_method", r.Method)

	logContext.Str("remote_addr", r.RemoteAddr)
	logContext.Str("uri", r.RequestURI)

	entry.Logger = logContext.Logger()

	entry.Logger.Info().Msg("request started")

	return entry
}

type StructuredLoggerEntry struct {
	Logger zerolog.Logger
}

func (l *StructuredLoggerEntry) Write(status, bytes int, header http.Header, elapsed time.Duration, extra interface{}) {
	l.Logger.Info().
		Int("status", status).
		Int("bytes_length", bytes).
		Float64("elapsed", float64(elapsed.Round(time.Millisecond/100))).
		Msg("request complete")
}

func (l *StructuredLoggerEntry) Panic(v interface{}, stack []byte) {
	l.Logger.Error().Msgf("panic: %v", v)
}
