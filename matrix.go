package kinadv360pro

import "github.com/bgould/keyboard-firmware/keyboard"

const (
	NumRows = 5
	NumCols = 10
)

func (dev *Device) NewMatrix() *keyboard.Matrix {
	return keyboard.NewMatrix(NumRows, NumCols, dev)
}

type Device struct {
}

func NewDevice() *Device {
	return &Device{}
}

func (dev *Device) ReadRow(i uint8) keyboard.Row {
	return 0
}
