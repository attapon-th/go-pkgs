// Package zlog - helping create zerolog and logrotation
package zlog

import (
	"io"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/goccy/go-json"
	"github.com/rs/zerolog"
)

const (

	// TimeFormatLayout Defind default time formet logger
	TimeFormatLayout = time.RFC3339
)

// Logger - A Logger represents an active logging object
type Logger = zerolog.Logger

func init() {
	zerolog.TimeFieldFormat = TimeFormatLayout
	zerolog.InterfaceMarshalFunc = json.Marshal
	zerolog.CallerMarshalFunc = callerHandler
}

// NewConsoleJSON - Set logging with console log with json format
//
//	@return Logger
func NewConsoleJSON(opts ...Option) Logger {
	cfg := Config{}
	cfg.loadOptions(opts...)
	logger := zerolog.New(os.Stdout).Level(zerolog.DebugLevel).With().Timestamp().Logger()
	if !cfg.NoCaller {
		logger = logger.With().Caller().Logger()
	}
	return logger
}

// NewConsole - Set logging with console log
//
//	@return *Logger
func NewConsole(opts ...Option) Logger {
	cfg := Config{}
	cfg.loadOptions(opts...)
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: TimeFormatLayout, NoColor: cfg.NoColor}
	logger := zerolog.New(output).Level(zerolog.DebugLevel).With().Timestamp().Logger()
	if !cfg.NoCaller {
		logger = logger.With().Caller().Logger()
	}
	return logger
}

// NewLogRollingFile - Set logging with filename. Auto rotation file every day backup 90 day
//
//	@param filePath - full path logfile
//	@return *Logger
func NewLogRollingFile(filePath string, opts ...Option) Logger {
	cfg := Config{
		ConsoleLoggingEnabled: true,
		FileLoggingEnabled:    true,
		EncodeLogsAsJSON:      true,
		Directory:             path.Dir(filePath),
		Filename:              path.Base(filePath),
		MaxSize:               1024, // MB
		MaxBackups:            90,   // Day
		MaxAge:                1,    // MaxAge the max age in days to keep a logfile
	}
	cfg.loadOptions(opts...)
	return NewCustomLogRollingFile(cfg)
}

// NewCustomLogRollingFile sets up the logging framework
func NewCustomLogRollingFile(config Config) Logger {
	var writers []io.Writer

	if config.ConsoleLoggingEnabled {
		ch := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: TimeFormatLayout}
		writers = append(writers, ch)
	}
	if config.FileLoggingEnabled {
		writers = append(writers, newRollingFile(config))
	}
	mw := io.MultiWriter(writers...)

	// zerolog.SetGlobalLevel(zerolog.DebugLevel)
	logger := zerolog.New(mw).Level(config.Level).With().Timestamp().Logger()

	return logger
}

func callerHandler(pc uintptr, file string, line int) string {
	short := file
	for i := len(file) - 1; i > 0; i-- {
		if file[i] == '/' {
			short = file[i+1:]
			break
		}
	}
	file = short
	return file + ":" + strconv.Itoa(line)
}
