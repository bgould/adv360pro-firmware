//go:build tinygo

package adv360pro

import (
	"image/color"
	"machine"

	"tinygo.org/x/drivers/ws2812"
)

type LedPos uint8

const (
	LedLeft LedPos = iota
	LedMiddle
	LedRight
)

type Indicators struct {
	ws2812 ws2812.Device
	leds   [PIXELS]color.RGBA
}

func (ind *Indicators) Configure() {
	WS2812.Configure(machine.PinConfig{Mode: machine.PinOutput})
}

func (ind *Indicators) Set(pos LedPos, c color.RGBA) {
	switch pos {
	case LedLeft:
		break
	case LedMiddle:
		break
	case LedRight:
		break
	default:
		return
	}
	ind.leds[pos] = c
}

func (ind *Indicators) Get(pos LedPos) (c color.RGBA) {
	switch pos {
	case LedLeft:
		break
	case LedMiddle:
		break
	case LedRight:
		break
	default:
		return
	}
	return ind.leds[pos]
}

func (ind *Indicators) Sync() {
	ws2812_dev.WriteColors(ind.leds[:])
}
