//go:build tinygo && nrf52840

package adv360pro

import (
	"machine"
)

var (
	BATTERY = machine.P0_04
	LED     = machine.P0_17
	WS2812  = machine.P0_20
	PIXELS  = 3

	LeftCols = cols[0 : MatrixCols/2]

	RightCols = cols[MatrixCols/2 : MatrixCols]

	LeftRows = rows[0:MatrixRows]

	RightRows = rows[MatrixRows : MatrixRows*2]
)

const (
	LeftOffset  = 0
	RightOffset = 10
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

	pwm = machine.PWM0
)
