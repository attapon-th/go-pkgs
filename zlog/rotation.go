package zlog

import (
	"io"
	"os"
	"path"

	"github.com/rs/zerolog/log"
	"gopkg.in/natefinch/lumberjack.v2"
)

func newRollingFile(config Config) io.Writer {
	if err := os.MkdirAll(config.Directory, 0744); err != nil {
		log.Error().Err(err).Str("path", config.Directory).Msg("can't create log directory")
		return nil
	}

	return &lumberjack.Logger{
		Filename:   path.Join(config.Directory, config.Filename),
		MaxBackups: config.MaxBackups, // files
		MaxSize:    config.MaxSize,    // megabytes
		MaxAge:     config.MaxAge,     // days
		Compress:   true,
		LocalTime:  true,
	}
}
