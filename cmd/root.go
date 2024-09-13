/*
Copyright © 2024 Enrique Marín

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

const (
	VERSION = "1.0.0"

	Reset   = "\033[0m"
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
	Cyan    = "\033[36m"
	Gray    = "\033[37m"
	White   = "\033[97m"
	Orange  = "\x1b[38:5:202m"
	Maroon  = "\x1b[38:5:1m"
	Lime    = "\x1b[38:5:46m"
	Bold    = "\033[1m"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "task",
	Version: VERSION,
	Short:   "Task is a command line application for managing to-do tasks.",
	Long: `
Task is a simple yet elegant command-line application
for managing to-do tasks,
allowing computer terminal lovers to work
without leaving their favorite environment.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(db *gorm.DB) {
	rootCmd.AddCommand(newAddCommand(db))
	rootCmd.AddCommand(newListCommand(db))
	rootCmd.AddCommand(newDetailCommand(db))
	rootCmd.AddCommand(newUpdateCommand(db))
	rootCmd.AddCommand(newToggledCommand(db))
	rootCmd.AddCommand(newDeleteCommand(db))

	err := rootCmd.Execute()
	if err != nil {
		log.Println(err)

		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.go-cli-bubbletea-todoapp.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
