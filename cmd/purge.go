/*
Copyright © 2022 Mohammed Al-Mahdawi <mohammed@al-mahdawi.is>

*/
package cmd

import (
	"log"
	"os"

	"github.com/MohammedAl-Mahdawi/glass/helpers"
	"github.com/spf13/cobra"
)

// purgeCmd represents the purge command
var purgeCmd = &cobra.Command{
	Use:   "purge",
	Short: "Delete all logs or single log",
	Long: `If -u flag not provided it will clear
All logs.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		if Uuid != "" {
			dir := helpers.GetCmdPath(Uuid)
			err := os.RemoveAll(dir)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			err := os.RemoveAll(helpers.GetGlassLogPath())
			if err != nil {
				log.Fatal(err)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(purgeCmd)

	purgeCmd.PersistentFlags().StringVarP(&Uuid, "uuid", "u", "", "UUID of the command")
}
