//go:build tinygo

package adv360pro

import (
	"machine"
	"slices"
)

type serialNumber [16]byte

func SerialNumber() (sn serialNumber) {
	deviceID := machine.DeviceID()
	slices.Reverse(deviceID)
	bin2hex(deviceID, sn[0:16])
	return
}

func (sn *serialNumber) String() string {
	return string(sn[:])
}

func bin2hex(in []byte, out []byte) {
	const (
		chars = "0123456789ABCDEF"
	)
	for i, b := range in {
		var n = i * 2
		out[n+0] = chars[b>>4]
		out[n+1] = chars[b&15]
	}
}
