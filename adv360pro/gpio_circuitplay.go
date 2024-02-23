//go:build tinygo && circuitplay_bluefruit

package adv360pro

import (
	"machine"

	"github.com/bgould/keyboard-firmware/keyboard"
)

const (
	BACKLIGHT = machine.LED
	// BATTERY = machine.P0_04
	WS2812 = machine.WS2812
	PIXELS = 10
	LED    = BACKLIGHT
)

var (
	BACKLIGHT_PWM = machine.PWM0
)

var (
	rows = [MatrixRows * 2]machine.Pin{
		// left
		machine.NoPin,
		machine.NoPin,
		machine.NoPin,
		machine.NoPin,
		machine.NoPin,
		// right
		machine.NoPin,
		machine.NoPin,
		machine.NoPin,
		machine.NoPin,
		machine.NoPin,
	}

	cols = [MatrixCols]machine.Pin{
		// left
		machine.BUTTONA,
		machine.BUTTONB,
		machine.NoPin,
		machine.NoPin,
		machine.NoPin,
		machine.NoPin,
		machine.NoPin,
		machine.NoPin,
		machine.NoPin,
		machine.NoPin,
		// right
		machine.BUTTONA,
		machine.BUTTONB,
		machine.NoPin,
		machine.NoPin,
		machine.NoPin,
		machine.NoPin,
		machine.NoPin,
		machine.NoPin,
		machine.NoPin,
		machine.NoPin,
	}
	colMode = machine.PinInputPulldown
)

func (m *Device) readRow(rowIndex uint8) keyboard.Row {
	switch rowIndex {
	case 0:
		v := keyboard.Row(0)
		if rowIndex == 0 {
			if machine.BUTTONA.Get() {
				i := 0x0 // left
				if m.offset == RightOffset {
					i = 0xF // right
				}
				v |= (1 << i)
			}
			if machine.BUTTONB.Get() {
				i := 0x1 // left
				if m.offset == RightOffset {
					i = 0x10 // right
				}
				v |= (1 << i)
			}
		}
		return v
	default:
		return 0
	}
}
