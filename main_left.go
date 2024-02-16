//go:build adv360pro_left

package main

import "github.com/bgould/kinadv360pro-firmware/adv360pro"

var (
	device = adv360pro.NewDeviceLeft()
)

const (
	ProductString = "Adv360 Pro (Left)"
)
