package ble

import (
	"fmt"

	"github.com/bgould/keyboard-firmware/keyboard"
)

type MatrixPacket [20]byte

func NewMatrixPacket(rows, cols int) MatrixPacket {
	return MatrixPacket{uint8(rows), uint8(cols)}
}

func (pkt *MatrixPacket) Rows() int {
	return int(pkt[0])
}

func (pkt *MatrixPacket) Cols() int {
	return int(pkt[1])
}

func (pkt *MatrixPacket) DecodeTo(matrix []keyboard.Row) (ok bool, rows, cols uint8) {
	const (
		hdrsz = 4
	)
	rows, cols = pkt[0], pkt[1]
	bpr := (cols / 8) + 1 // bytes per row
	for row, n := uint8(0), uint8(len(matrix)); row < rows && row < n; row++ {
		state := keyboard.Row(0)
		for col := uint8(0); col < bpr; col++ {
			state |= keyboard.Row(pkt[hdrsz+(row*bpr)+col]) << (col * 8)
		}
		matrix[row] = state & (0xFFFFFFFF >> (32 - cols))
	}
	return true, rows, cols
}

func (pkt *MatrixPacket) EncodeFrom(rows, cols uint8, matrix []keyboard.Row) {
	const (
		hdrsz = 4
	)
	pkt[0] = rows
	pkt[1] = cols
	bpr := (cols / 8) + 1 // bytes per row
	for row, n := uint8(0), uint8(len(matrix)); row < rows && row < n; row++ {
		state := matrix[row] & (0xFFFFFFFF >> (32 - cols))
		for col := uint8(0); col < bpr; col++ {
			b := uint8(state >> (col * 8))
			pkt[hdrsz+(row*bpr)+col] = b
		}
	}
}

func (pkt *MatrixPacket) String() string {
	return fmt.Sprintf(
		"%02x %02x %02x %02x %02x %02x %02x %02x %02x %02x %02x %02x %02x %02x %02x %02x %02x %02x %02x %02x ",
		pkt[0], pkt[1], pkt[2], pkt[3], pkt[4], pkt[5], pkt[6], pkt[7], pkt[8], pkt[9],
		pkt[10], pkt[11], pkt[12], pkt[13], pkt[14], pkt[15], pkt[16], pkt[17], pkt[18], pkt[19],
	)
}
