package zlog

import (
	"io"
	"os"
	"path"
	"strconv"
	"time"

	jsonitor "github.com/json-iterator/go"
	"github.com/rs/zerolog"
)

const (

	// TimeFormatLayout Defind default time formet logger
	TimeFormatLayout = time.RFC3339
)

// Logger - A Logger represents an active logging object
type Logger struct {
	cfg Config
	zerolog.Logger
}

func init() {
	zerolog.TimeFieldFormat = TimeFormatLayout
	zerolog.InterfaceMarshalFunc = jsonitor.Marshal
	zerolog.CallerMarshalFunc = callerHandler
}

// NewConsoleJSON - Set logging with console log with json format
//
//	@return Logger
func NewConsoleJSON() Logger {
	logger := zerolog.New(os.Stdout).Level(zerolog.DebugLevel).With().Timestamp().Logger()
	return Logger{Logger: logger}
}

// NewConsole - Set logging with console log
//
//	@return *Logger
func NewConsole() Logger {
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: TimeFormatLayout}
	logger := zerolog.New(output).Level(zerolog.DebugLevel).With().Timestamp().Logger()
	return Logger{Logger: logger}
}

// NewConsoleNoColor - Set logging with console log and set no color
//
//	@return *Logger
func NewConsoleNoColor() Logger {
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: TimeFormatLayout, NoColor: true}
	logger := zerolog.New(output).Level(zerolog.DebugLevel).With().Timestamp().Logger()
	return Logger{Logger: logger}
}

// NewLogRollingFile - Set logging with filename. Auto rotation file every day backup 90 day
//
//	@param filePath - full path logfile
//	@return *Logger
func NewLogRollingFile(filePath string) Logger {
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
	return NewCustomLogRollingFile(cfg)
}

// NewCustomLogRollingFile sets up the logging framework
func NewCustomLogRollingFile(config Config) Logger {
	var writers []io.Writer

	if config.ConsoleLoggingEnabled {
		writers = append(writers, NewConsole())
	}
	if config.FileLoggingEnabled {
		writers = append(writers, newRollingFile(config))
	}
	mw := io.MultiWriter(writers...)

	// zerolog.SetGlobalLevel(zerolog.DebugLevel)
	logger := zerolog.New(mw).Level(config.Level).With().Timestamp().Logger()

	return Logger{
		Logger: logger,
		cfg:    config,
	}
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

// EnableCaller = Caller adds the file:line
//
//	@return Logger
func (l Logger) EnableCaller() Logger {
	l.Logger = l.With().Caller().Logger()
	return l
}

// GetZerolog - get zerolog logging
//
//	@return zerolog.Logger
func (l Logger) GetZerolog() zerolog.Logger {
	return l.Logger
}

// SetZerolog - Set logger by zerolog
//
//	@param z
//	@return Logger
func (l Logger) SetZerolog(z zerolog.Logger) Logger {
	l.Logger = z
	return l
}
