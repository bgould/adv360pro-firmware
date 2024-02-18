//go:build adv360pro_ble_primary

package main

import (
	"time"

	"github.com/bgould/adv360pro-firmware/adv360pro"
	"github.com/bgould/adv360pro-firmware/adv360pro/ble"
	"github.com/bgould/keyboard-firmware/keyboard"
)

func startBLE() {
	time.Sleep(time.Second)
	ble.Default.EnablePrimary(receiveNotification)
	println("connected to split", ble.Default.Mode())
}

func receiveNotification(b []byte) {
	if debug {
		print("recv --> ")
		for i, n := 0, len(b); i < n; i++ {
			print(" ", b[i])
		}
		println()
	}
	if len(b) != 20 {
		return
	}
	var overlay [adv360pro.MatrixRows]keyboard.Row
	var packet ble.MatrixPacket
	copy(packet[:], b)
	if ok, rows, cols := packet.DecodeTo(overlay[:]); !ok {
		return
	} else if rows != adv360pro.MatrixRows {
		return
	} else if cols != adv360pro.MatrixCols {
		return
	}
	device.SetOverlay(overlay)
}

func bleTask() {

}
