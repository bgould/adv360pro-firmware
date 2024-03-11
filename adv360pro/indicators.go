//go:build tinygo

package adv360pro

import (
	"image/color"
	"machine"
)

type LedPos uint8

const (
	IndicatorLeft LedPos = iota
	IndicatorMiddle
	IndicatorRight

	IndicatorOffsetLeft  LedPos = 2
	IndicatorOffsetRight LedPos = 4

	IndicatorLayer = IndicatorRight

	IndicatorCapsLock = IndicatorLeft << IndicatorOffsetLeft
	IndicatorProfile  = IndicatorMiddle << IndicatorOffsetLeft

	IndicatorNumLock    = IndicatorLeft << IndicatorOffsetRight
	IndicatorScrollLock = IndicatorMiddle << IndicatorOffsetRight
)

var LayerColors = []color.RGBA{
	{R: 0x00, G: 0x00, B: 0x00, A: 0x00}, // default layer - off
	{R: 0x00, G: 0x00, B: 0xFF, A: 0x00}, // function layer - blue
	{R: 0x00, G: 0xFF, B: 0x00, A: 0x00}, // mod layer - green
	{R: 0x00, G: 0x00, B: 0x00, A: 0x00}, // empty layer - off
	{R: 0x00, G: 0x00, B: 0x00, A: 0x00}, // empty layer - off
}

var ProfileColors = []color.RGBA{
	{R: 0xFF, G: 0xFF, B: 0xFF, A: 0x00}, // white
	{R: 0x00, G: 0x00, B: 0xFF, A: 0x00}, // blue
	{R: 0xFF, G: 0x00, B: 0x00, A: 0x00}, // red
	{R: 0x00, G: 0xFF, B: 0x00, A: 0x00}, // green
	{R: 0x00, G: 0x00, B: 0x00, A: 0x00}, // off
}

type Indicators struct {
	Offset LedPos
	leds   [PIXELS]color.RGBA
}

func (ind *Indicators) Configure() {
	WS2812.Configure(machine.PinConfig{Mode: machine.PinOutput})
}

func (ind *Indicators) Set(pos LedPos, c color.RGBA) {
	pos, ok := ind.offsetPos(pos)
	if !ok {
		return
	}
	ind.leds[pos] = c
}

func (ind *Indicators) Get(pos LedPos) (c color.RGBA) {
	pos, ok := ind.offsetPos(pos)
	if !ok {
		return
	}
	return ind.leds[pos]
}

func (ind *Indicators) Sync() {
	ws2812_dev.WriteColors(ind.leds[:])
}

func (ind *Indicators) offsetPos(pos LedPos) (newpos LedPos, ok bool) {
	if pos > 0b11 {
		pos = pos >> ind.Offset
	}
	switch pos {
	case IndicatorLeft:
		return pos, true
	case IndicatorRight:
		return pos, true
	case IndicatorMiddle:
		return pos, true
	default:
		return pos, false
	}
}
