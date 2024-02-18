//go:build adv360pro || nrf52840_generic

package main

import (
	"machine/usb"
)

func init() {
	usb.VendorID = 0x29ea
	usb.ProductID = 0x0362
}
