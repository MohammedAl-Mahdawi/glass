package main

import (
	"os/signal"
	"syscall"

	"github.com/MohammedAl-Mahdawi/glass/cmd"
)

func main() {
	signal.Ignore(syscall.SIGHUP)
	cmd.Execute()
}
