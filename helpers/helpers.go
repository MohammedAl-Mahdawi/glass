package helpers

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/MohammedAl-Mahdawi/glass/types"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

func GetCmdPath(uuid string) string {
	dir, err := filepath.Abs(GetGlassLogPath() + "/" + uuid)
	if err != nil {
		log.Fatal(err)
	}
	os.MkdirAll(dir, os.ModePerm)

	return dir
}

func GetGlassLogPath() string {
	dir := os.Getenv("GLASS_LOG_PATH")
	if len(dir) == 0 {
		d, err := filepath.Abs("./glass")
		if err != nil {
			log.Fatal(err)
		}
		dir = d
	}

	return dir
}

func WriteToLog(output *types.Log, logPath string) error {
	d, err := yaml.Marshal(&output)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(logPath, d, 0777)
	if err != nil {
		return err
	}

	return nil
}

// GetCommand get the command after --
func GetCommand(cmd *cobra.Command, args []string) []string {
	argsLenAtDash := cmd.ArgsLenAtDash()

	if argsLenAtDash > -1 {
		return args[argsLenAtDash:]
	}

	return []string{}
}
