package ble

import (
	"fmt"
	"testing"

	"github.com/bgould/keyboard-firmware/keyboard"
)

func TestMatrixPacketCreate(t *testing.T) {

	const (
		MatrixRows = 5
		MatrixCols = 20
	)

	matrices := map[string][MatrixRows]keyboard.Row{
		"blank": {
			0b0000000000_0000000000,
			0b0000000000_0000000000,
			0b0000000000_0000000000,
			0b0000000000_0000000000,
			0b0000000000_0000000000,
		},
		"all": {
			0b1111111111_1111111111,
			0b1111111111_1111111111,
			0b1111111111_1111111111,
			0b1111111111_1111111111,
			0b1111111111_1111111111,
		},
		"test": {
			0b0000000000_0000000000,
			0b0000000000_0000000000,
			0b0000000000_0000011110,
			0b0000000000_0000011110,
			0b0000000000_0000000000,
		},
		"test2": {
			0b111111111_1111111111,
			0b111111111_1111111111,
			0b111111111_1111111111,
			0b111111111_1111111111,
			0b111111111_1111111111,
		},
	}

	var pkt MatrixPacket

	for k, v := range matrices {
		test := v
		fmt.Printf("testing (%v): %v\n", k, test)
		pkt.EncodeFrom(MatrixRows, MatrixCols, test[:])
		if pkt.Rows() != 5 {
			t.Fail()
		}
		if pkt.Cols() != 20 {
			t.Fail()
		}
		fmt.Printf("encoded (%v): %v\n", k, pkt)

		roundTrip := [MatrixRows]keyboard.Row{}
		ok, rows, cols := pkt.DecodeTo(roundTrip[:])
		if !ok {
			t.Fail()
		}
		if rows != 5 {
			t.Fail()
		}
		if cols != 20 {
			t.Fail()
		}
		fmt.Printf("finding (%v): %v\n", k, roundTrip)
		if roundTrip != test {
			t.Fail()
		}
	}

	//fmt.Printf("%v", pkt.String())
}
