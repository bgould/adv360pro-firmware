//go:build adv360pro_right || ledglasses_nrf52840

package main

import (
	"github.com/bgould/adv360pro-firmware/adv360pro"
	"github.com/bgould/keyboard-firmware/hosts/usbvial/vial"
)

var (
	device = adv360pro.NewDeviceRight()

	unlockKeys = []vial.Pos{{Row: 0, Col: 19}}
)

const (
	ProductString = "Adv360 Pro (Right)"
)
