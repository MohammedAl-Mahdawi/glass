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

// logCmd represents the log command
var logCmd = &cobra.Command{
	Use:   "log",
	Short: "Return the command log",
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

		fmt.Println(l.Log)
	},
}

func init() {
	rootCmd.AddCommand(logCmd)
	logCmd.PersistentFlags().StringVarP(&Uuid, "uuid", "u", "", "The command uuid")
}
