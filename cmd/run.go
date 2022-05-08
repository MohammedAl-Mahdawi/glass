package cmd

import (
	"bytes"
	"context"
	"log"
	"os/exec"
	"strings"
	"time"

	"github.com/MohammedAl-Mahdawi/glass/helpers"
	"github.com/MohammedAl-Mahdawi/glass/types"
	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run [flags] -- COMMAND [args...]",
	Short: "Run command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		subCommand := helpers.GetCommand(cmd, args)
		if len(subCommand) == 0 {
			log.Fatal("No command to run")
		}

		if Timeout == "" {
			// Default timeout is 12 hours
			Timeout = "12h"
		}

		dr, err := time.ParseDuration(Timeout)
		if err != nil {
			log.Fatal(err)
		}

		ctx, cancel := context.WithTimeout(context.Background(), dr)
		defer cancel()

		cd := exec.Command(subCommand[0], subCommand[1:]...)

		var stdout bytes.Buffer
		var stderr bytes.Buffer

		cd.Stdout = &stdout
		cd.Stderr = &stderr

		if err := cd.Start(); err != nil {
			log.Fatal(err)
		}

		dir := helpers.GetCmdPath(Uuid)
		logPath := dir + "/log.yml"

		// Send signal on command run complete/error
		ch := make(chan int)
		var er error

		go func() {
			if err := cd.Wait(); err != nil {
				er = err
				ch <- 1
				return
			}

			select {
			case <-ch:
				return
			default:
				ch <- 0
			}
		}()

		createdAt := time.Now()

		output := types.Log{
			Command:     strings.Join(subCommand, " "),
			CreatedAt:   createdAt,
			CompletedAt: time.Time{},
			ExitStatus:  0,
			Status:      "started",
			Log:         "",
			UUID:        Uuid,
		}

		err = helpers.WriteToLog(&output, logPath)
		if err != nil {
			log.Fatal(err)
		}

		select {
		case <-ctx.Done():
			_ = cd.Process.Kill()
			out := stdout.String() + "\n" + stderr.String() + "\n" + "Command timed out!"
			output.CompletedAt = time.Now()
			output.Log = out

			output.Status = "timeout"
			if exitError, ok := er.(*exec.ExitError); ok {
				ec := exitError.ExitCode()
				output.ExitStatus = ec
			} else {
				output.ExitStatus = 888
			}

			err2 := helpers.WriteToLog(&output, logPath)
			if err2 != nil {
				log.Fatal(err2)
			}
		case v := <-ch:
			out := stdout.String() + "\n" + stderr.String()
			output.CompletedAt = time.Now()
			output.Log = out
			if v == 1 {
				output.Status = "err"
				if exitError, ok := er.(*exec.ExitError); ok {
					ec := exitError.ExitCode()
					output.ExitStatus = ec
				} else {
					output.ExitStatus = 999
				}

				err2 := helpers.WriteToLog(&output, logPath)
				if err2 != nil {
					log.Fatal(err2)
				}
			} else {
				output.Status = "completed"

				err = helpers.WriteToLog(&output, logPath)
				if err != nil {
					log.Fatal(err)
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
	runCmd.PersistentFlags().StringVarP(&Uuid, "uuid", "u", "", "UUID of the command")
	runCmd.PersistentFlags().StringVarP(&Timeout, "timeout", "t", "", "Timeout of the command, defaults to 12h")
}
