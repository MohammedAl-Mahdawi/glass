package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var Uuid string
var Timeout string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "glass",
	Short: "Run, report back, & keep a full history of command",
	Long: `Easily run and manage commands with Glass, 
	Glass can run any command, and it will write the 
	stdout & the stderr to log file along with everything, like 
	you can give the command an ID and follow its progress later 
	and check to see whether it's completed or not and its exit status 
	and so on.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
