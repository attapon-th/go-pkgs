package main

import (
	"github.com/attapon-th/go-pkgs/zlog"
	"github.com/attapon-th/go-pkgs/zlog/log"
)

func main() {
	logApiAcccess := zlog.NewLogRollingFile("./logs/access.log")
	logApiAcccess.Write([]byte("logApiAcccess"))
	// log.GetLogger().Write([]byte("logApiAcccess"))
	// ExampleNewConsoleColor()
	// ExampleNewConsoleJSON()
	// ExampleGlobalLogging()
	// ExampleNewLogRollingFile()
}

func ExampleGlobalLogging() {
	log.Print("Print Global logging")
	log.Debug().Msg("Debug")
	log.Info().Msg("Info")
	log.Warn().Msg("Warn")
	log.Error().Msg("Error")

	defer func() {
		x := recover()
		if x != nil {
			log.Fatal().Msgf("Fatal --> Recover: %v", x)
		}
	}()
	log.Panic().Msg("Panic")
}

func ExampleNewConsoleColor() {
	l := zlog.NewConsole(zlog.WithColor(true), zlog.WithCaller(true))
	l.Print("Print logging with json format")
	l.Debug().Msg("Debug")
	l.Info().Msg("Info")
	l.Warn().Msg("Warn")
	l.Error().Msg("Error")

	defer func() {
		x := recover()
		if x != nil {
			l.Fatal().Msgf("Fatal --> Recover: %v", x)
		}
	}()
	l.Panic().Msg("Panic")
}

func ExampleNewConsoleJSON() {
	l := zlog.NewConsoleJSON(zlog.WithCaller(true))
	l = l.With().Str("logging", "json").Logger()
	l.Print("Print logging with json format")
	l.Debug().Msg("Debug")
	l.Info().Msg("Info")
	l.Warn().Msg("Warn")
	l.Error().Msg("Error")

	defer func() {
		x := recover()
		if x != nil {
			l.Fatal().Msgf("Fatal --> Recover: %v", x)
		}
	}()
	l.Panic().Msg("Panic")
}

func ExampleNewLogRollingFile() {
	l := zlog.NewLogRollingFile("./log-rolling.log")
	l.Print("Print logging with json format")
	l.Debug().Msg("Debug")
	l.Info().Msg("Info")
	l.Warn().Msg("Warn")
	l.Error().Msg("Error")

	defer func() {
		x := recover()
		if x != nil {
			l.Fatal().Msgf("Fatal --> Recover: %v", x)
		}
	}()
	l.Panic().Msg("Panic")
}
