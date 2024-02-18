//go:build tinygo && nrf52840_generic

package adv360pro

import (
	"machine"

	"github.com/bgould/keyboard-firmware/keyboard"
)

const (
	LED     = machine.P0_17
	BATTERY = machine.P0_04
	WS2812  = machine.P0_20
	PIXELS  = 3
)

var (
	rows = [MatrixRows * 2]machine.Pin{
		// left
		machine.P1_11,
		machine.P1_15,
		machine.P0_03,
		machine.P1_14,
		machine.P1_12,
		// right
		machine.P0_19,
		machine.P0_05,
		machine.P0_31,
		machine.P0_30,
		machine.P0_29,
	}

	cols = [MatrixCols]machine.Pin{
		// left
		machine.P0_25,
		machine.P0_11,
		machine.P0_02,
		machine.P0_28,
		machine.P0_29,
		machine.P0_30,
		machine.P0_31,
		machine.P1_09,
		machine.P0_12,
		machine.P0_07,
		// right
		machine.P0_12,
		machine.P1_09,
		machine.P0_07,
		machine.P1_11,
		machine.P1_10,
		machine.P1_13,
		machine.P1_15,
		machine.P0_03,
		machine.P0_02,
		machine.P0_28,
	}
	colMode = machine.PinInputPullup
)

func (m *Device) readRow(rowIndex uint8) (row keyboard.Row) {
	for i, pin := range m.rows {
		v := i != int(rowIndex)
		// v := i == int(rowIndex)
		pin.Set(v)
	}
	delayMicros(5)
	for i, pin := range m.cols {
		v := pin.Get()
		// if v {
		if !v {
			row |= (1 << (i + m.offset))
		}
	}
	return row
}
