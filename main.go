//go:build tinygo && nrf52840

package main

import (
	"image/color"
	"machine"
	"machine/usb"
	"time"

	"github.com/bgould/keyboard-firmware/hosts/serial"
	"github.com/bgould/keyboard-firmware/hosts/usbvial"
	"github.com/bgould/keyboard-firmware/hosts/usbvial/vial"
	"github.com/bgould/keyboard-firmware/keyboard"
	"github.com/bgould/kinadv360pro-firmware/kinadv360pro"
	"tinygo.org/x/drivers/ws2812"
)

const _debug = true

var (
	keymap = initKeymap()
	device = kinadv360pro.NewDevice(kinadv360pro.LeftRows[:], kinadv360pro.LeftCols[:])
	matrix = keyboard.NewMatrix(kinadv360pro.NumRows, kinadv360pro.NumCols, device)
)

func init() {
	usb.VendorID = 0x29ea
	usb.ProductID = 0x0362
	usb.Manufacturer = "Kinesis Corporation"
	usb.Product = "Adv360 Pro"
	usb.Serial = vial.MagicSerialNumber("")
}

func main() {

	time.Sleep(2 * time.Second)

	// use the onboard LED as a status indicator
	// machine.LED.Configure(machine.PinConfig{Mode: machine.PinOutput})
	// machine.LED.Low()

	// create the keyboard console
	console := serial.DefaultConsole()

	println("initializing")
	device.Initialize()

	configureNeo()
	// configurePins()

	host := usbvial.NewKeyboard(VialDeviceDefinition, keymap, matrix)

	board := keyboard.New(console, host, matrix, keymap)
	board.SetDebug(_debug)

	// machine.LED.High()

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

var leds [3]color.RGBA

func configureNeo() {
	// led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	kinadv360pro.WS2812.Configure(machine.PinConfig{Mode: machine.PinOutput})

	ws := ws2812.NewWS2812(kinadv360pro.WS2812)
	for i := range leds {
		leds[i] = color.RGBA{R: 0xff, G: 0xff, B: 0xff}
	}
	ws.WriteColors(leds[:])
	// rg := false

	// for {
	// 	rg = !rg
	// 	for i := range leds {
	// 		rg = !rg
	// 		if rg {
	// 			// Alpha channel is not supported by WS2812 so we leave it out
	// 			leds[i] = color.RGBA{R: 0xff, G: 0x00, B: 0x00}
	// 		} else {
	// 			leds[i] = color.RGBA{R: 0x00, G: 0xff, B: 0x00}
	// 		}
	// 	}

	// 	ws.WriteColors(leds[:])
	// 	led.Set(rg)
	// 	time.Sleep(100 * time.Millisecond)
	// }
}
