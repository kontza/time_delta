package cmd

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/romana/rlog"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func rootRunner(cmd *cobra.Command, args []string) {
	const timeFormShort = "15:04:05"
	const timeFormLong = "15:04:05.000"
	var times [2]time.Time
	var err error
	pat := regexp.MustCompile(":|\\.")
	usedLongForm := false
	for i, arg := range args {
		arg = strings.Replace(arg, ",", ".", -1)
		if i > 1 {
			break
		}
		timeForm := timeFormShort
		partCount := len(pat.Split(arg, -1))
		rlog.Debugf("Component count: %d, arg: %s", partCount, arg)
		if partCount > 3 {
			timeForm = timeFormLong
			usedLongForm = true
		}
		if times[i], err = time.Parse(timeForm, arg); err != nil {
			rlog.Criticalf("Failed to parse '%s': %v", arg, err)
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
	if viper.GetBool("showHours") {
		fmt.Printf("%.2f h\n", delta.Hours())
	}
	if viper.GetBool("showMinutes") {
		fmt.Printf("%.2f min\n", delta.Minutes())
	}
}
