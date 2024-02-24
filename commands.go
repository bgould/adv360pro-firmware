package main

import (
	"machine"
	"machine/usb"
	"strconv"

	"github.com/bgould/keyboard-firmware/keyboard/console"
)

func initConsole() *console.Console {
	return console.New(machine.Serial, console.Commands{
		"status": console.CommandHandlerFunc(status),
	})
}

func status(cmd console.CommandInfo) int {
	cmd.Stdout.Write([]byte("\n[Device]\n-------\n"))
	cmd.Stdout.Write([]byte("serial number: "))
	cmd.Stdout.Write([]byte(serialNumber[:]))
	cmd.Stdout.Write([]byte("\n"))
	cmd.Stdout.Write([]byte("scan rate:     "))
	cmd.Stdout.Write([]byte(strconv.Itoa(scanRate)))
	cmd.Stdout.Write([]byte("\n"))
	cmd.Stdout.Write([]byte("\n[USB]\n----\n"))
	cmd.Stdout.Write([]byte("Manufacturer: "))
	cmd.Stdout.Write([]byte(usb.Manufacturer))
	cmd.Stdout.Write([]byte("\n"))
	cmd.Stdout.Write([]byte("Product:      "))
	cmd.Stdout.Write([]byte(usb.Product))
	cmd.Stdout.Write([]byte("\n"))
	cmd.Stdout.Write([]byte("Serial:       "))
	cmd.Stdout.Write([]byte(usb.Serial))
	cmd.Stdout.Write([]byte("\n"))
	cmd.Stdout.Write([]byte("\n"))
	return 0
}
