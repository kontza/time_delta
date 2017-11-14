package cmd

import (
	"github.com/spf13/cobra"
	"time"
	"github.com/spf13/viper"
	"fmt"
)

func rootRunner(cmd *cobra.Command, args []string) {
	if viper.GetBool("showVersion") {
		println("time_delta v1.0.0")
		return
	}
	if len(args) < 2 {
		logger.Fatal("Incorrect amount of arguments! There must be at least two.")
	}
	const timeForm = "15:04:05"
	var times [2]time.Time
	var err error
	for i, arg := range args {
		if i > 1 {
			break
		}
		if times[i], err = time.Parse(timeForm, arg); err != nil {
			logger.Fatalf("Failed to parse '%s': %v", arg, err)
		}
	}
	var delta time.Duration
	var alpha int
	var omega int
	midnight, _ := time.Parse(timeForm, "0:00:00")
	if times[0].After(times[1]) {
		delta = times[0].Sub(times[1])
		alpha = int(times[1].Sub(midnight).Seconds())
		omega = int(times[0].Sub(midnight).Seconds())
	} else {
		delta = times[1].Sub(times[0])
		alpha = int(times[0].Sub(midnight).Seconds())
		omega = int(times[1].Sub(midnight).Seconds())
	}
	fmt.Printf("%d - %d = %d\n", omega, alpha, int(delta.Seconds()))
	if viper.GetBool("showHours") {
		fmt.Printf("%.2f h\n", delta.Hours())
	}
	if viper.GetBool("showMinutes") {
		fmt.Printf("%.2f min\n", delta.Minutes())
	}
}
