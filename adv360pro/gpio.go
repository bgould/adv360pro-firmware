//go:build tinygo && nrf52840

package adv360pro

import (
	"machine"
)

const (
	LED     = machine.P0_17
	BATTERY = machine.P0_04
	WS2812  = machine.P0_20
	PIXELS  = 3
)

var (
	// LeftCols is a slice of GPIO pins for the left hand device with the matrix columns
	LeftCols = cols[0 : MatrixCols/2]

	// RightCols is a slice of GPIO pins for the right hand device with the matrix columns
	RightCols = cols[MatrixCols/2 : MatrixCols]

	// LeftRows is a slice of GPIO pins for the left hand device with the matrix rows
	LeftRows = rows[0:MatrixRows]

	// RightRows is a slice of GPIO pins for the right hand device with the matrix rows
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
