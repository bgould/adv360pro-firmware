//go:build tinygo

package adv360pro

import (
	"machine"
	"time"

	"github.com/bgould/keyboard-firmware/keyboard"
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

type Device struct {
	rows    []machine.Pin
	cols    []machine.Pin
	offset  int
	overlay [MatrixRows]keyboard.Row
}

func NewDeviceLeft() *Device {
	return newDevice(LeftRows, LeftCols, LeftOffset)
}

func NewDeviceRight() *Device {
	return newDevice(RightRows, RightCols, RightOffset)
}

func newDevice(rows []machine.Pin, cols []machine.Pin, offset int) *Device {
	return &Device{rows: rows, cols: cols, offset: offset}
}

// Initialize matrix and peripherals, returning an error if any is unavailable.
func (m *Device) Initialize() (err error) {
	for _, pin := range m.rows {
		if pin != machine.NoPin {
			pin.Configure(machine.PinConfig{Mode: machine.PinOutput})
		}
	}
	for _, pin := range m.cols {
		if pin != machine.NoPin {
			pin.Configure(machine.PinConfig{Mode: colMode})
		}
	}
	return nil
}

const (
	MatrixRows = 5
	MatrixCols = 20
)

func (dev *Device) NewMatrix() *keyboard.Matrix {
	return keyboard.NewMatrix(MatrixRows, MatrixCols, dev)
}

func (dev *Device) SetOverlay(overlay [MatrixRows]keyboard.Row) {
	dev.overlay = overlay
}

// ReadRow
func (m *Device) ReadRow(rowIndex uint8) keyboard.Row {
	const (
		maskLeft = 0b00000000_00000000_00000011_11111111
		maskRght = 0b00000000_00001111_11111100_00000000
	)
	row := m.readRow(rowIndex)
	if m.offset == RightOffset {
		return (row & maskRght) | (m.overlay[rowIndex] & maskLeft)
	} else {
		return (row & maskLeft) | (m.overlay[rowIndex] & maskRght)
	}
}

func delayMicros(usecs int) {
	duration := time.Duration(usecs) * time.Microsecond
	for start := time.Now(); time.Since(start) < duration; {
	}
}
