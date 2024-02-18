//go:build adv360pro_ble_secondary

package main

import (
	"time"

	"github.com/bgould/adv360pro-firmware/adv360pro"
	"github.com/bgould/adv360pro-firmware/adv360pro/ble"
	"github.com/bgould/keyboard-firmware/keyboard"
)

func init() {
	board.SetEventReceiver(keyboard.EventReceiverFunc(eventReceiver))
}

func startBLE() {
	time.Sleep(time.Second)
	if debug {
		println("enabling secondary")
	}
	ble.Default.EnableSecondary()
	if debug {
		println("enabled secondary device", ble.Default.Mode())
	}

	// go func() {
	// 	for {
	// 		if receivedEvent {
	// 			sendMatrixPacket()
	// 			receivedEvent = false
	// 		}
	// 		time.Sleep(10 * time.Millisecond)
	// 	}
	// }()
}

var (
	lastSent   time.Time
	lastMatrix [adv360pro.MatrixRows]keyboard.Row
)

func bleTask() {
	// now := time.Now()
	// // if now.Sub(lastSent) > 5*time.Millisecond {
	// 	sendMatrixPacket()
	// 	lastSent = now
	// }
}

// var receivedEvent bool

func eventReceiver(ev keyboard.Event) (bool, error) {
	sendMatrixPacket()
	return true, nil
	// receivedEvent = true

	// var packet ble.MatrixPacket
	// var matrix [adv360pro.MatrixRows]keyboard.Row

	// for i := 0; i < adv360pro.MatrixRows; i++ {
	// 	matrix[i] = keyboard.Row(board.GetMatrixRowState(i))
	// }
	// packet.EncodeFrom(adv360pro.MatrixRows, adv360pro.MatrixCols, matrix[:])
	// // for i, n := 0, int(board.MatrixRows()); i < n; i++ {
	// // 	if (i+1)*3 < 20 {
	// // 		state := board.GetMatrixRowState(i)
	// // 		bin[i*3+0] = byte(state >> 0xF)
	// // 		bin[i*3+1] = byte(state >> 0x8)
	// // 		bin[i*3+2] = byte(state >> 0x0)
	// // 	}
	// // }
	// if debug {
	// 	print("send --> ")
	// 	for i := 0; i < len(packet); i++ {
	// 		print(" ", packet[i])
	// 	}
	// }
	// println()
	// if ok, err := ble.Default.Tx(packet[:]); !ok {
	// 	if err != nil {
	// 		println("tx error:", err.Error())
	// 	} else {
	// 		println("tx not sent")
	// 	}
	// }

	// return true, nil
}

func getMatrix() (matrix [adv360pro.MatrixRows]keyboard.Row) {
	for i := 0; i < adv360pro.MatrixRows; i++ {
		matrix[i] = keyboard.Row(board.GetMatrixRowState(i))
	}
	return
}

func sendMatrixPacket() (bool, error) {
	var packet ble.MatrixPacket
	matrix := getMatrix()
	packet.EncodeFrom(adv360pro.MatrixRows, adv360pro.MatrixCols, matrix[:])
	if debug {
		print("send --> ")
		for i := 0; i < len(packet); i++ {
			print(" ", packet[i])
		}
		println()
	}
	if ok, err := ble.Default.Tx(packet[:]); !ok {
		if err != nil {
			if debug {
				println("tx error:", err.Error())
			}
			return false, err
		} else {
			if debug {
				println("tx not sent")
			}
			return false, nil
		}
	}
	lastMatrix = matrix
	lastSent = time.Now()
	return true, nil
}
