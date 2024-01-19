//go:build tinygo && nrf52840

package main

import (
	"machine"
	"machine/usb"
	"time"

	"github.com/bgould/keyboard-firmware/hosts/serial"
	"github.com/bgould/keyboard-firmware/hosts/usbvial"
	"github.com/bgould/keyboard-firmware/hosts/usbvial/vial"
	"github.com/bgould/keyboard-firmware/keyboard"
	"tinygo.org/x/drivers/delay"
)

const _debug = true

var (
	// keymap = CircuitPlaygroundDefaultKeymap()
	matrix = keyboard.NewMatrix(1, 2, keyboard.RowReaderFunc(ReadRow))
)

func main() {

	// use the onboard LED as a status indicator
	machine.LED.Configure(machine.PinConfig{Mode: machine.PinOutput})
	machine.LED.Low()

	// create the keyboard console
	console := serial.DefaultConsole()

	configurePins()

	usb.Serial = vial.MagicSerialNumber("")
	host := usbvial.NewKeyboard(VialDeviceDefinition, keymap, matrix)

	board := keyboard.New(console, host, matrix, keymap)
	board.SetDebug(_debug)

	machine.LED.High()

	for {
		board.Task()
	}

}

func configurePins() {
	// for _, pin := range pins {
	// 	pin.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})
	// 	println("configured pin", pin, pin.Get())
	// }
}

func ReadRow(rowIndex uint8) keyboard.Row {
	delay.Sleep(50 * time.Microsecond)
	// switch rowIndex {
	// case 0:
	// 	v := keyboard.Row(0)
	// 	for i := range pins {
	// 		if pins[i].Get() {
	// 			v |= (1 << i)
	// 		}
	// 	}
	// 	return v
	// default:
	// 	return 0
	// }
	return 0
}
