//go:build tinygo && nrf52840

package main

import (
	"machine"
	"machine/usb"
	"time"

	"github.com/bgould/adv360pro-firmware/adv360pro/ble"
	"github.com/bgould/keyboard-firmware/hosts/usbvial"
	"github.com/bgould/keyboard-firmware/hosts/usbvial/vial"
	"github.com/bgould/keyboard-firmware/keyboard"
	"github.com/bgould/keyboard-firmware/keyboard/keycodes"
)

const (
	debug = false
)

var (
	cli    = initConsole()
	keymap = initKeymap()
	matrix = device.NewMatrix()
	host   = usbvial.NewKeyboard(VialDeviceDefinition, keymap, matrix)
	board  = keyboard.New(machine.Serial, host, matrix, keymap)
)

func init() {
	usb.Manufacturer = "Kinesis Corporation"
	usb.Product = ProductString
	usb.Serial = vial.MagicSerialNumber("")

	ble.Default.Enable()

	board.SetKeyAction(keyboard.KeyActionFunc(keyAction))
	board.SetEnterBootloaderFunc(keyboard.DefaultEnterBootloader)
	board.SetCPUResetFunc(keyboard.DefaultCPUReset)

	device.Initialize()
}

func main() {

	startBLE()

	for last, count := time.Now(), 0; true; count++ {
		now := time.Now()
		if d := now.Sub(last); d > time.Second {
			scanRate = (count * 1000) / int(d/time.Millisecond)
			count = 0
			last = now
		}
		board.Task()
		cli.Task()
		bleTask()
		time.Sleep(330 * time.Microsecond)
	}

}

var (
	scanRate int

	fn0made time.Time
	fn1prev uint8
	fn2prev uint8
)

// func configureKeyAction() keyboard.KeyActionFunc {
// return func(key keycodes.Keycode, made bool) {
func keyAction(key keycodes.Keycode, made bool) {
	switch key {

	// Handle "reset" press
	case keycodes.QK_BOOT:
		fallthrough
	case keycodes.KC_FN0:
		if made {
			println("QK_BOOT down")
			fn0made = time.Now()
		} else {
			if time.Since(fn0made) > 2*time.Second {
				println("entering bootloader")
				time.Sleep(100 * time.Millisecond)
				board.EnterBootloader()
			} else {
				println("resetting CPU")
				time.Sleep(100 * time.Millisecond)
				board.CPUReset()
			}
		}

	// Toggle function layer on key down/up
	case keycodes.KC_FN1:
		if made {
			fn1prev = board.ActiveLayer()
			board.SetActiveLayer(1)
		} else {
			board.SetActiveLayer(fn1prev)
			fn1prev = 0
		}
		if fn1prev == 1 {
			fn1prev = 0
		}

	// Toggle programming layer on key down/up
	case keycodes.KC_FN2:
		if made {
			fn2prev = board.ActiveLayer()
			board.SetActiveLayer(2)
			println("layer 2 on")
		} else {
			board.SetActiveLayer(fn2prev)
			fn2prev = 0
			println("layer 2 off")
		}
		if fn2prev == 2 {
			fn2prev = 0
		}

		// Toggle keypad layer on keypress
		// case keycodes.KC_FN0:
		// 	if fn0made && !made {
		// 		if board.ActiveLayer() == 1 {
		// 			board.SetActiveLayer(0)
		// 		} else {
		// 			board.SetActiveLayer(1)
		// 		}
		// 	}
		// 	fn0made = made}

	}
}
