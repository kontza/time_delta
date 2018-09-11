package cmd

import (
	"fmt"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

func rootRunner(cmd *cobra.Command, args []string) {
	const timeFormShort = "15:04:05"
	var times [2]time.Time
	var err error
	for i, arg := range args {
		arg = strings.Replace(arg, ",", ".", 1)
		if times[i], err = time.Parse(timeFormShort, arg); err != nil {
			logger.Fatal().Str("reason", fmt.Sprintf("%+v", err)).Msg("Parsing failed")
		} else {
			logger.Debug().Str("source", arg).Str("result", fmt.Sprintf("%v", times[i])).Msg("Parsed time value")
		}
	}
	var delta time.Duration
	var alpha float64
	var omega float64
	midnight := time.Date(0, time.January, 1, 0, 0, 0, 0, time.UTC)
	if times[0].After(times[1]) {
		delta = times[0].Sub(times[1])
		alpha = times[1].Sub(midnight).Seconds()
		omega = times[0].Sub(midnight).Seconds()
	} else {
		delta = times[1].Sub(times[0])
		alpha = times[0].Sub(midnight).Seconds()
		omega = times[1].Sub(midnight).Seconds()
	}
	truncated := delta.Truncate(time.Second)
	useFloat := (delta != delta.Truncate(time.Second))
	logger.Debug().Dur("delta", delta).Dur("truncated", truncated).Bool("use float result", useFloat).Msg("Calculated difference")
	if useFloat {
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
