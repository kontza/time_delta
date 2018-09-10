package cmd

import (
	"fmt"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

func printParseResult(source string, result time.Time, parser string) {
	logger.Debug().Str("source", source).Str("result", fmt.Sprintf("%v", result)).Str("parser", parser).Msg("Parsed time value")
}

func rootRunner(cmd *cobra.Command, args []string) {
	const timeFormShort = "15:04:05"
	const timeFormLong = "15:04:05.000"
	var times [2]time.Time
	var err error
	usedLongForm := false
	for i, arg := range args {
		arg = strings.Replace(arg, ",", ".", 1)
		if times[i], err = time.Parse(timeFormShort, arg); err != nil {
			if times[i], err = time.Parse(timeFormLong, arg); err != nil {
				logger.Fatal().Str("reason", fmt.Sprintf("%+v", err)).Msg("Parsing failed")
			} else {
				printParseResult(arg, times[i], "timeFormLong")
			}
		} else {
			printParseResult(arg, times[i], "timeFormShort")
		}
	}
	var delta time.Duration
	var alpha float64
	var omega float64
	midnight, _ := time.Parse(timeFormShort, "0:00:00")
	if times[0].After(times[1]) {
		delta = times[0].Sub(times[1])
		alpha = times[1].Sub(midnight).Seconds()
		omega = times[0].Sub(midnight).Seconds()
	} else {
		delta = times[1].Sub(times[0])
		alpha = times[0].Sub(midnight).Seconds()
		omega = times[1].Sub(midnight).Seconds()
	}
	if usedLongForm {
		fmt.Printf("%.3f - %.3f = %.3f\n", omega, alpha, delta.Seconds())
	} else {
		fmt.Printf("%d - %d = %d\n", int(omega), int(alpha), int(delta.Seconds()))
	}
	if showHours, _ := cmd.PersistentFlags().GetBool("show-hours"); showHours {
		fmt.Printf("%.2f h\n", delta.Hours())
	}
	if showMinutes, _ := cmd.PersistentFlags().GetBool("show-hours"); showMinutes {
		fmt.Printf("%.2f min\n", delta.Minutes())
	}
}
