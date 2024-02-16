//go:build adv360pro_right

package main

import "github.com/bgould/kinadv360pro-firmware/adv360pro"

var (
	device = adv360pro.NewDeviceRight()
)

const (
	ProductString = "Adv360 Pro (Right)"
)
