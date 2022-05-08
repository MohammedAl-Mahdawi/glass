/*
Copyright © 2022 Mohammed Al-Mahdawi <mohammed@al-mahdawi.is>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Return Glass's version",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Glass version 1.0.0")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
