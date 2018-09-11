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

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var (
	cfgFile string
	rootCmd = &cobra.Command{
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
		Run:     rootRunner,
		Args:    cobra.MinimumNArgs(2),
		Version: "v1.3.1",
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(loggerInit)
	rootCmd.PersistentFlags().BoolP("structured", "s", false, "use structured logging")
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "show verbose logging")
	rootCmd.PersistentFlags().BoolP("show-hours", "t", false, "include fractional hours in the result output")
	rootCmd.PersistentFlags().BoolP("show-minutes", "m", false, "include fractional minutes in the result output")
}
