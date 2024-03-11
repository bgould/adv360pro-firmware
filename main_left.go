//go:build adv360pro_left || !(adv360pro_right || ledglasses_nrf52840)

package main

import (
	"github.com/bgould/adv360pro-firmware/adv360pro"
	"github.com/bgould/keyboard-firmware/hosts/usbvial/vial"
)

var (
	device     = adv360pro.NewDeviceLeft()
	unlockKeys = []vial.Pos{{Row: 0, Col: 0}}
)

const (
	ProductString = "Adv360 Pro (Left)"
)

func init() {
	device.Indicators.Offset = adv360pro.IndicatorOffsetLeft
}

// func init() {
// 	board.SetEventReceiver(keyboard.EventReceiverFunc(eventReceiver))
// }

// var (
// 	bin [20]byte
// )

// func eventReceiver(ev keyboard.Event) (bool, error) {
// 	bin = [20]byte{}
// 	for i, n := 0, int(board.MatrixRows()); i < n; i++ {
// 		if (i+1)*3 < 20 {
// 			state := board.GetMatrixRowState(i)
// 			bin[i*3+0] = byte(state >> 0xF)
// 			bin[i*3+1] = byte(state >> 0x8)
// 			bin[i*3+2] = byte(state >> 0x0)
// 		}
// 	}
// 	print("matrix --> ")
// 	for i := 0; i < len(bin); i++ {
// 		print(" ", bin[i])
// 	}
// 	println()
// 	return true, nil
// }
