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

// reportCmd represents the report command
var reportCmd = &cobra.Command{
	Use:   "report",
	Short: "Return the status of command",
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

		fmt.Println(l.Status)
	},
}

func init() {
	rootCmd.AddCommand(reportCmd)

	reportCmd.PersistentFlags().StringVarP(&Uuid, "uuid", "u", "", "The command uuid")
}
