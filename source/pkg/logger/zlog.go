package logger

import (
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/i-sub135/go-rest-blueprint/source/config"
	"github.com/rs/zerolog"
)

var Log zerolog.Logger

// Init initializes global logger.
// levelStr example: "debug", "info"
// prettyConsole: when true, use human-friendly console writer
// callerSkip: frames to skip so caller points to original caller (use 2 if wrapping)
func Init(prettyConsole bool) {
	// parse level
	lvl, err := zerolog.ParseLevel(config.GetConfig().Log.Level)
	if err != nil {
		lvl = zerolog.InfoLevel
	}

	zerolog.CallerMarshalFunc = func(pc uintptr, file string, line int) string {
		return fmt.Sprintf("%s:%d", file, line) // Full path instead of filepath.Base(file)
	}
	zerolog.CallerSkipFrameCount = 3
	zerolog.TimeFieldFormat = time.RFC3339

	if prettyConsole {
		out := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
		Log = zerolog.New(out).Level(lvl).With().Timestamp().Str("app", config.GetConfig().App.Name).Logger()
	} else {
		Log = zerolog.New(os.Stdout).Level(lvl).With().Timestamp().Str("app", config.GetConfig().App.Name).Logger()
	}
}

// convenience chainable functions
func Debug() *zerolog.Event { return Log.Debug() }
func Info() *zerolog.Event  { return Log.Info() }
func Warn() *zerolog.Event  { return Log.Warn() }
func Error() *zerolog.Event { return Log.Error() }

// GinZLogger returns middleware that logs request after handler runs.
func GinZLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		dur := time.Since(start)
		status := c.Writer.Status()
		method := c.Request.Method
		path := c.FullPath()
		if path == "" {
			path = c.Request.URL.Path
		}
		ip := c.ClientIP()

		var ev *zerolog.Event
		switch {
		case status >= 500:
			ev = Log.Error()
		case status >= 400:
			ev = Log.Warn()
		case status >= 300:
			ev = Log.Debug()
		default:
			ev = Log.Info()
		}

		ev.Str("method", method).
			Str("path", path).
			Str("client_ip", ip).
			Int("status", status).
			Dur("latency", dur).
			Msg("http request")
	}
}
