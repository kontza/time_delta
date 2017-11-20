// Copyright © 2017 Juha Ruotsalainen <kontza@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// RootCmd represents the base command when called without any subcommands
var (
	RootCmd = &cobra.Command{
		Use:   "time_delta",
		Short: "Calculates the time difference between two given hour-minute-second times.",
		Long: `
time_delta expects the time values to be as in the following format:
	hh:mm:ss[.ms]
	|  |  |  |
	|  |  |  + milliseconds (optional)
	|  |  +--- seconds
	|  +------ minutes
	+--------- hours
Output is by default in seconds, but command line options exist to get the output
in minutes, or in hours.`,
		// Uncomment the following line if your bare application
		// has an action associated with it:
		PreRun: loggerInit,
		Run:    rootRunner,
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.time_delta.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	RootCmd.PersistentFlags().BoolP("verbose", "v", false, "show verbose logging")
	RootCmd.PersistentFlags().BoolP("version", "V", false, "show version information")
	RootCmd.PersistentFlags().BoolP("hours", "t", false, "include fractional hours in the result output")
	RootCmd.PersistentFlags().BoolP("minutes", "m", false, "include fractional minutes in the result output")
	viper.BindPFlag("beVerbose", RootCmd.PersistentFlags().Lookup("verbose"))
	viper.BindPFlag("showVersion", RootCmd.PersistentFlags().Lookup("version"))
	viper.BindPFlag("showMinutes", RootCmd.PersistentFlags().Lookup("minutes"))
	viper.BindPFlag("showHours", RootCmd.PersistentFlags().Lookup("hours"))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".time_delta" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".time_delta")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
