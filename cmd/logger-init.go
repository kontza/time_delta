package cmd

import (
	"os"

	"github.com/rs/zerolog"
)

var (
	logger zerolog.Logger
)

func loggerInit() {
	logLevel := zerolog.InfoLevel
	zerolog.TimeFieldFormat = " "

	verbose, _ := rootCmd.PersistentFlags().GetBool("verbose")
	structured, _ := rootCmd.PersistentFlags().GetBool("structured")

	if verbose || structured {
		logLevel = zerolog.DebugLevel
		zerolog.TimeFieldFormat = "2006/01/06 15:04:05.000"
	}
	zerolog.SetGlobalLevel(logLevel)
	if structured {
		logger = zerolog.New(os.Stderr).With().Timestamp().Logger()
	} else {
		logger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr}).With().Timestamp().Logger()
	}
	logger.Debug().Str("Logging level", zerolog.GlobalLevel().String()).Msg("Current")
}
