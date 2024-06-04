package logger

import (
	"io"
	"log/slog"
	"os"

	v "github.com/mrtnhwttktc/kinto-cli/internal/version"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

// can be updated using the Set method, LogLevel.Set(slog.LevelDebug)
var LogLevel = &slog.LevelVar{} // INFO

func SetLogLevel(loglevel string) {
	switch loglevel {
	case "debug":
		LogLevel.Set(slog.LevelDebug)
	case "info":
		LogLevel.Set(slog.LevelInfo)
	case "warn":
		LogLevel.Set(slog.LevelWarn)
	case "error":
		LogLevel.Set(slog.LevelError)
	default:
		LogLevel.Set(slog.LevelInfo)
	}
}

// SetDefaultLogger initializes the logger with a JSON handler and a rotating file handler.
func SetDefaultLogger() {
	rotatingLogger := getLumberjackConfig()

	opts := &slog.HandlerOptions{
		Level: LogLevel,
	}

	handler := slog.NewJSONHandler(rotatingLogger, opts)
	logger := slog.New(handler)
	logger = setDefaultKeys(logger)
	slog.SetDefault(logger)
}

// getLumberjackConfig returns a configuration for the lumberjack logger.
func getLumberjackConfig() *lumberjack.Logger {
	return &lumberjack.Logger{
		Filename:   "/tmp/ktcli/ktcli.log",
		MaxSize:    1024, // megabytes
		MaxBackups: 3,
		MaxAge:     28, // days
	}
}

// VerboseLogger configures the logger to operate in verbose mode.
// In this mode, log messages are outputted to both stdout and a rotating file.
// The log level is determined by the configuration file or the level set in the flags. Defalut is INFO.
// Additionally, the source of each log message is included in the log entry.
func VerboseLogger() {
	rotatingLogger := getLumberjackConfig()

	opts := &slog.HandlerOptions{
		Level:     LogLevel,
		AddSource: true,
	}

	multiWriter := io.MultiWriter(os.Stderr, rotatingLogger)
	handler := slog.NewJSONHandler(multiWriter, opts)

	logger := slog.New(handler)
	logger = setDefaultKeys(logger)
	slog.SetDefault(logger)
}

func setDefaultKeys(logger *slog.Logger) *slog.Logger {
	defaultKeys := slog.Group(
		"binary", // group name
		slog.String("version", v.Version),
	)
	return logger.With(defaultKeys)
}
