package adv360pro

import "machine"

func SerialNumber() string {
	return bin2hex(machine.DeviceID())
}

func bin2hex(in []byte) string {
	const (
		chars = "0123456789ABCDEF"
	)
	out := make([]byte, len(in)*2)
	for i, b := range in {
		n := (len(in) - 1 - i) * 2
		out[n+0] = chars[b>>4]
		out[n+1] = chars[b&15]
	}
	return string(out)
}
