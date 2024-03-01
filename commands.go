//go:build tinygo

package main

import (
	"machine"
	"machine/usb"
	"strconv"

	"github.com/bgould/keyboard-firmware/keyboard/console"
)

var (
	commands = console.Commands{}
)

func init() {
	commands["status"] = console.CommandHandlerFunc(status)
}

func initConsole() *console.Console {
	return console.New(machine.Serial, commands)
}

func status(cmd console.CommandInfo) int {

	cmd.Stdout.Write([]byte("\n[USB]\n-----\n"))
	cmd.Stdout.Write([]byte("Manufacturer: "))
	cmd.Stdout.Write([]byte(usb.Manufacturer))
	cmd.Stdout.Write([]byte("\n"))
	cmd.Stdout.Write([]byte("Product:      "))
	cmd.Stdout.Write([]byte(usb.Product))
	cmd.Stdout.Write([]byte("\n"))
	cmd.Stdout.Write([]byte("Serial:       "))
	cmd.Stdout.Write([]byte(usb.Serial))
	cmd.Stdout.Write([]byte("\n"))

	var dataStart, dataEnd [8]byte
	st, en := machine.FlashDataStart(), machine.FlashDataEnd()
	bin2hex([]byte{byte(st >> 24), byte(st >> 16), byte(st >> 8), byte(st)}, dataStart[:])
	bin2hex([]byte{byte(en >> 24), byte(en >> 16), byte(en >> 8), byte(en)}, dataEnd[:])

	cmd.Stdout.Write([]byte("\n[Device]\n--------\n"))
	cmd.Stdout.Write([]byte("Serial Number: "))
	cmd.Stdout.Write([]byte(serialNumber[:]))
	cmd.Stdout.Write([]byte("\n"))
	cmd.Stdout.Write([]byte("Scan Rate:     "))
	cmd.Stdout.Write([]byte(strconv.Itoa(scanRate)))
	cmd.Stdout.Write([]byte("\n"))
	cmd.Stdout.Write([]byte("Flash Start:   "))
	cmd.Stdout.Write(dataStart[:])
	cmd.Stdout.Write([]byte("\n"))
	cmd.Stdout.Write([]byte("Flash End:     "))
	cmd.Stdout.Write(dataEnd[:])
	cmd.Stdout.Write([]byte("\n"))

	cmd.Stdout.Write([]byte("\n"))

	return 0
}

func save(cmd console.CommandInfo) int {
	cmd.Stdout.Write([]byte("Saving keymap ...\n"))
	name := "saved.keymap"
	if n, err := saveKeymap(name); err != nil {
		cmd.Stdout.Write([]byte("Error saving keymap: "))
		cmd.Stdout.Write([]byte(err.Error()))
		cmd.Stdout.Write([]byte("\n"))
		return 1
	} else {
		cmd.Stdout.Write([]byte("Wrote "))
		cmd.Stdout.Write([]byte(strconv.Itoa(int(n))))
		cmd.Stdout.Write([]byte(" bytes. Keymap saved successfully.\n"))
		return 0
	}
}

func load(cmd console.CommandInfo) int {
	cmd.Stdout.Write([]byte("Loading keymap ...\n"))
	name := "saved.keymap"
	if n, err := loadKeymap(name); err != nil {
		cmd.Stdout.Write([]byte("Error loading keymap: "))
		cmd.Stdout.Write([]byte(err.Error()))
		cmd.Stdout.Write([]byte("\n"))
		return 1
	} else {
		cmd.Stdout.Write([]byte("Read "))
		cmd.Stdout.Write([]byte(strconv.Itoa(int(n))))
		cmd.Stdout.Write([]byte(" bytes. Keymap loaded successfully.\n"))
		return 0
	}
}

func bin2hex(in []byte, out []byte) {
	const (
		chars = "0123456789ABCDEF"
	)
	for i, b := range in {
		var n = i * 2
		out[n+0] = chars[b>>4]
		out[n+1] = chars[b&15]
	}
}
