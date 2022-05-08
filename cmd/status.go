/*
Copyright © 2022 Mohammed Al-Mahdawi <mohammed@al-mahdawi.is>

*/
package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/MohammedAl-Mahdawi/glass/helpers"
	"github.com/MohammedAl-Mahdawi/glass/types"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Return exit code of the command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		dir := helpers.GetCmdPath(Uuid)
		logPath := dir + "/log.yml"

		file, err := os.Open(logPath)
		if err != nil {
			log.Fatal(err)
		}

		data, _ := ioutil.ReadAll(file)

		l := types.Log{}

		err = yaml.Unmarshal(data, &l)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(l.ExitStatus)
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)

	statusCmd.PersistentFlags().StringVarP(&Uuid, "uuid", "u", "", "UUID of the command")
}
