//go:build tinygo && ledglasses_nrf52840

package adv360pro

import (
	"machine"

	"github.com/bgould/keyboard-firmware/keyboard"
)

const (
	BACKLIGHT = machine.LED
	BATTERY   = machine.D20
	WS2812    = machine.WS2812
	PIXELS    = 1
	LED       = BACKLIGHT
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
		machine.BUTTON,
		machine.NoPin,
		machine.NoPin,
		machine.NoPin,
		machine.NoPin,
		machine.NoPin,
		machine.NoPin,
		machine.NoPin,
		machine.NoPin,
		machine.NoPin,
		// right
		machine.BUTTON,
		machine.NoPin,
		machine.NoPin,
		machine.NoPin,
		machine.NoPin,
		machine.NoPin,
		machine.NoPin,
		machine.NoPin,
		machine.NoPin,
		machine.NoPin,
	}
	colMode = machine.PinInputPullup
)

func (m *Device) readRow(rowIndex uint8) keyboard.Row {
	switch rowIndex {
	case 4:
		v := keyboard.Row(0)
		if !machine.BUTTON.Get() {
			i := 0x6 // left
			if m.offset == RightOffset {
				i = 0xD // right
			}
			v |= (1 << i)
		}
		return v
	default:
		return 0
	}
}
