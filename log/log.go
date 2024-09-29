package log

import (
	"fmt"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	log.Logger = log.Output(zerolog.ConsoleWriter{
		Out:        os.Stderr,
		TimeFormat: time.RFC3339,
	})
}

func SetLevel(level string) {
	logLevel, err := zerolog.ParseLevel(level)
	if err != nil {
		fmt.Printf("Invalid log level '%s', fallback to 'debug'\n", level)
		logLevel = zerolog.DebugLevel
	}

	zerolog.SetGlobalLevel(logLevel)
}

func Panic(msg string) {
	log.Panic().Msg(msg)
}

func Panicf(format string, v ...any) {
	log.Panic().Msgf(format, v...)
}

func Fatal(msg string) {
	log.Fatal().Msg(msg)
}

func Fatalf(format string, v ...any) {
	log.Fatal().Msgf(format, v...)
}

func Error(msg string) {
	log.Error().Msg(msg)
}

func Errorf(format string, v ...any) {
	log.Error().Msgf(format, v...)
}

func Warn(msg string) {
	log.Warn().Msg(msg)
}

func Warnf(format string, v ...any) {
	log.Warn().Msgf(format, v...)
}

func Info(msg string) {
	log.Info().Msg(msg)
}

func Infof(format string, v ...any) {
	log.Info().Msgf(format, v...)
}

func Debug(msg string) {
	log.Debug().Msg(msg)
}

func Debugf(format string, v ...any) {
	log.Debug().Msgf(format, v...)
}

func Trace(msg string) {
	log.Trace().Msg(msg)
}

func Tracef(format string, v ...any) {
	log.Trace().Msgf(format, v...)
}
