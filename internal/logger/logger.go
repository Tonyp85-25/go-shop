package logger

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	zlog "github.com/rs/zerolog/log"
)

func New() zerolog.Logger {
	zerolog.TimeFieldFormat = time.RFC3339

	if os.Getenv("GIN_MODE") != "release" {
		zlog.Logger = zlog.Output(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339})
	}
	return zlog.Logger
}
