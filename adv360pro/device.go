//go:build tinygo

package adv360pro

import (
	"machine"
	"time"

	"github.com/bgould/keyboard-firmware/keyboard"
)

type Device struct {
	rows   []machine.Pin
	cols   []machine.Pin
	offset int
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
		pin.Configure(machine.PinConfig{Mode: machine.PinOutput})
	}
	for _, pin := range m.cols {
		pin.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
		// pin.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})
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

// ReadRow
func (m *Device) ReadRow(rowIndex uint8) (row keyboard.Row) {
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

func delayMicros(usecs int) {
	duration := time.Duration(usecs) * time.Microsecond
	for start := time.Now(); time.Since(start) < duration; {
	}
}
