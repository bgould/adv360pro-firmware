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
)

//go:generate go run github.com/bgould/keyboard-firmware/hosts/usbvial/gen-def vial.json

const _debug = true

func init() {
	machine.LED = machine.P1_14
	machine.UART_TX_PIN = machine.P0_14 // PORTB
	machine.UART_RX_PIN = machine.P0_30 // PORTB
}

var (
	keymap = initKeymap()
	matrix = keyboard.NewMatrix(1, 2, keyboard.RowReaderFunc(ReadRow))
)

func main() {

	time.Sleep(1 * time.Second)

	// use the onboard LED as a status indicator
	machine.LED.Configure(machine.PinConfig{Mode: machine.PinOutput})
	machine.LED.Low()

	// create the keyboard console
	console := serial.DefaultConsole()

	configurePins()

	usb.Manufacturer = "TinyGo"
	usb.Serial = vial.MagicSerialNumber("")
	host := usbvial.NewKeyboard(VialDeviceDefinition, keymap, matrix)

	board := keyboard.New(console, host, matrix, keymap)
	board.SetDebug(_debug)

	machine.LED.High()

	var last time.Time
	for {
		if time.Since(last) > time.Second {
			println(last.String())
			last = time.Now()
		}
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
	// delay.Sleep(50 * time.Microsecond)
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
