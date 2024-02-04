//go:build tinygo && nrf52840

package main

import (
	"machine"
	"time"
)

func main() {
	for {
		time.Sleep(2 * time.Second)
		println(
			"flash data start:", machine.FlashDataStart(),
			"flash data end:", machine.FlashDataEnd(),
		)
	}
}
