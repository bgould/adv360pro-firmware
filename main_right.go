//go:build adv360pro_right || ledglasses_nrf52840

package main

import (
	"github.com/bgould/adv360pro-firmware/adv360pro"
)

var (
	device = adv360pro.NewDeviceRight()

	bin [20]byte
)

const (
	ProductString = "Adv360 Pro (Right)"
)
