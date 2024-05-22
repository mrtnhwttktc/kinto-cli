package log

import (
	"io"
	"log/slog"
	"os"

	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

// can be updated using the Set method, LogLevel.Set(slog.LevelDebug)
var LogLevel = &slog.LevelVar{} // INFO

// init initializes the logger with a rotating file handler.
func init() {
	rotatingLogger := getLumberjackConfig()

	opts := &slog.HandlerOptions{
		Level: LogLevel,
	}

	handler := slog.NewJSONHandler(rotatingLogger, opts)
	logger := slog.New(handler)
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

// VerboseLogger will set the logger to verbose mode, printing to stdout and a rotating file. The log level will be set to the level specified in the config file or the level set in the flags.
func VerboseLogger() {
	rotatingLogger := getLumberjackConfig()

	opts := &slog.HandlerOptions{
		Level: LogLevel,
	}

	multiWriter := io.MultiWriter(os.Stdout, rotatingLogger)
	handler := slog.NewJSONHandler(multiWriter, opts)
	slog.SetDefault(slog.New(handler))
}
