// Package log global logging
package log

import (
	"github.com/attapon-th/go-pkgs/zlog"
	"github.com/rs/zerolog"
)

var (
	l zlog.Logger
)

func init() {
	l = zlog.NewConsole(zlog.WithColor(true), zlog.WithCaller(true))
}

// GetLogger - get global logging
//
//	@return zlog.Logger
func GetLogger() zlog.Logger {
	return l
}

// SetLogger - set default global logging
func SetLogger(logger zlog.Logger) {
	l = logger
}

// Debug -  Print sends a log event using
func Debug() *zerolog.Event {
	return l.Debug()
}

// Info -  Print sends a log event using
func Info() *zerolog.Event {
	return l.Info()
}

// Warn -  Print sends a log event using
func Warn() *zerolog.Event {
	return l.Warn()
}

// Error -  Print sends a log event using
func Error() *zerolog.Event {
	return l.Error()
}

// Fatal -  Print sends a log event using
func Fatal() *zerolog.Event {
	return l.Fatal()
}

// Panic -  Print sends a log event using
func Panic() *zerolog.Event {
	return l.Panic()
}

// Printf -  Print sends a log event using
func Printf(format string, v ...interface{}) {
	l.Printf(format, v...)
}

// Print -  Print sends a log event using
func Print(v ...interface{}) {
	l.Print(v...)
}
