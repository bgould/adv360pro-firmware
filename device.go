package kinadv360pro

import (
	"machine"

	"github.com/bgould/keyboard-firmware/keyboard"
)

//go:generate go run github.com/bgould/keyboard-firmware/hosts/usbvial/gen-def vial.json

type Device struct {
}

func NewDevice(rows []machine.Pin, cols []machine.Pin) *Device {
	return &Device{}
}

// Initialize matrix and peripherals, returning an error if any is unavailable.
func (m *Device) Initialize() (err error) {

	// if err = m.configurePins(); err != nil {
	// 	return err
	// }

	return nil
}

// ReadRow
func (m *Device) ReadRow(rowIndex uint8) (row keyboard.Row) {

	// // set all row outputs to high except for rowIndex
	// rows := ^(uint16(1) << rowIndex)
	// m.port1.SetPins(mcp23017.Pins(rows), port1_rowMask)

	// // read input pins to determine which keys are pressed;
	// // any inputs with logic low indicate a key press at the
	// // given row,column in the matrix
	// pins, err := m.port0.GetPins()
	// if err != nil {
	// 	pins = 0
	// }
	// row = keyboard.Row((^pins) & port0_colMask)

	// // set all row outputs to high
	// m.port1.SetPins(mcp23017.Pins(^uint16(0)), port1_rowMask)

	return row
}
