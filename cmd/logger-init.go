package cmd

import (
	"os"

	"github.com/romana/rlog"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func loggerInit(cmd *cobra.Command, args []string) {
	logLevel := "INFO"
	hideTime := "yes"
	if viper.GetBool("beVerbose") {
		logLevel = "DEBUG"
		hideTime = "no"
	}
	os.Setenv("RLOG_LOG_LEVEL", logLevel)
	os.Setenv("RLOG_LOG_NOTIME", hideTime)
	os.Setenv("RLOG_TIME_FORMAT", "2006/01/06 15:04:05.000")
	rlog.UpdateEnv()
	rlog.Debug("Logging level:", logLevel)
}
