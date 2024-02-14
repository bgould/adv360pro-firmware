package main

import (
	"machine"
	"strconv"

	"github.com/bgould/keyboard-firmware/keyboard/console"
)

func initConsole() *console.Console {
	return console.New(machine.Serial, console.Commands{
		"status": console.CommandHandlerFunc(status),
	})
}

func status(cmd console.CommandInfo) int {
	cmd.Stdout.Write([]byte("status: "))
	cmd.Stdout.Write([]byte(strconv.Itoa(scanRate)))
	cmd.Stdout.Write([]byte("\n"))
	return 0
}
