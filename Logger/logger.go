// Package logger provides functions to set up a new logger
package logger

import (
	"fmt"
	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"net/http"
	"path"
	"runtime"
	"time"
)

const (
	logFormatJSON    = "json"
	logFormatText    = "text"
	logFormatConsole = "console"
)

const (
	LogLevelError = "error"
	LogLevelWarn  = "warn"
	LogLevelFatal = "fatal"
	LogLevelPanic = "panic"
	LogLevelDebug = "debug"
	LogLevelInfo  = "info"
)

// Logger is a small wrapper around a zap.Logger.
type Logger struct {
	*zap.Logger
}

// SetLogger returns a middleware that logs the start and end of each request, along
// with some useful data about what was requested, what the response status was,
// and how long it took to return.
// Inspired by https://github.com/treastech/logger
func SetLogger(l *Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

			t1 := time.Now()
			defer func() {
				l.Info("served request",
					zap.String("proto", r.Proto),
					zap.String("method", r.Method),
					zap.String("path", r.URL.Path),
					zap.Duration("lat", time.Since(t1)),
					zap.Int("status", ww.Status()),
					zap.Int("size", ww.BytesWritten()),
					zap.String("reqId", middleware.GetReqID(r.Context())))
			}()

			next.ServeHTTP(ww, r)
		}
		return http.HandlerFunc(fn)
	}
}

// New creates a new Logger with given logLevel and logFormat as part of a permanent field of the logger.
func New(logLevel, logFormat, env string) (*Logger, error) {
	if logFormat == logFormatText {
		logFormat = logFormatConsole
	}

	zapConfig := zap.NewProductionConfig()
	zapConfig.Encoding = logFormat

	var level zapcore.Level
	err := level.UnmarshalText([]byte(logLevel))
	if err != nil {
		return nil, err
	}

	date := time.Now().Format("20060102")
	_, filename, _, _ := runtime.Caller(1)
	envPath := path.Join(path.Dir(filename), "../log/sys.log."+date)

	zapConfig.Level = zap.NewAtomicLevelAt(level)
	zapConfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	zapConfig.OutputPaths = []string{"stdout", envPath}

	logger, err := zapConfig.Build()
	if err != nil {
		return nil, fmt.Errorf("could not build logger: %w", err)
	}

	zap.ReplaceGlobals(logger)

	return &Logger{Logger: logger}, nil
}
