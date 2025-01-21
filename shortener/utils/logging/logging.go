package logging

import (
	"os"
	"time"

	"github.com/rs/zerolog"
)

func init() {
	// Initialize the logger
	logger := zerolog.New(zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.RFC3339Nano, // Use a custom time format for console output
	}).With().Timestamp().Logger()

	zerolog.DefaultContextLogger = &logger

	logger.Debug().Msg("Global Logger initialized")
}
