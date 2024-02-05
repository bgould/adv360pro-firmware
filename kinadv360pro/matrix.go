package kinadv360pro

import "github.com/bgould/keyboard-firmware/keyboard"

const (
	NumRows = 10
	NumCols = 10
)

func (dev *Device) NewMatrix() *keyboard.Matrix {
	return keyboard.NewMatrix(NumRows, NumCols, dev)
}
