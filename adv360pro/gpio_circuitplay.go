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
		machine.SLIDER,
		machine.NoPin,
		machine.NoPin,
		machine.NoPin,
		machine.NoPin,
		// right
		machine.SLIDER,
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
	rowMode = machine.PinInputPullup
	colMode = machine.PinInputPulldown
)

func (m *Device) readRow(rowIndex uint8) (v keyboard.Row) {
	switch rowIndex {
	case 0:
		if machine.BUTTONA.Get() {
			i := 0x7 // left
			if m.offset == RightOffset {
				i = 0xB // right
			}
			v |= (1 << i)
		}
		if machine.BUTTONB.Get() {
			i := 0x8 // left
			if m.offset == RightOffset {
				i = 0xC // right
			}
			v |= (1 << i)
		}
		return
	case 4:
		if !machine.SLIDER.Get() {
			i := 0x6 // left
			if m.offset == RightOffset {
				i = 0xD // right
			}
			v |= (1 << i)
		}
		return
	default:
		return
	}
}
