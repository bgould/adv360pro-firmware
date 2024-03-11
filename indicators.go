package main

import (
	"image/color"

	"github.com/bgould/adv360pro-firmware/adv360pro"
	"github.com/bgould/keyboard-firmware/keyboard"
)

// |        |        |        | layer  | layer  | caps   | nlck   | slck   |
type IndicatorState keyboard.LEDs

func (s *IndicatorState) SetLayer(i uint8) {
	*s |= IndicatorState((i & 0b0011) << 3)
}

func (s *IndicatorState) GetLayer() uint8 {
	return (uint8(*s>>3) & 0b0011)
}

func (s *IndicatorState) Get(led keyboard.LED) bool {
	return (*keyboard.LEDs)(s).Get(led)
}

func (s *IndicatorState) Set(led keyboard.LED, on bool) {
	(*keyboard.LEDs)(s).Set(led, on)
}

func currentIndicatorState() IndicatorState {
	leds := IndicatorState(board.LEDs())
	leds.SetLayer(board.ActiveLayer())
	return leds
}

func syncIndicators(oldState IndicatorState) IndicatorState {
	leds := currentIndicatorState()
	if leds != oldState {
		layerColor := adv360pro.LayerColors[leds.GetLayer()]
		var caps, nlck, slck color.RGBA
		if leds.Get(keyboard.LEDCapsLock) {
			caps = adv360pro.ProfileColors[1]
		}
		if leds.Get(keyboard.LEDNumLock) {
			nlck = adv360pro.ProfileColors[1]

		}
		if leds.Get(keyboard.LEDScrollLock) {
			slck = adv360pro.ProfileColors[1]
		}
		device.Indicators.Set(adv360pro.IndicatorLayer, layerColor)
		device.Indicators.Set(adv360pro.IndicatorCapsLock, caps)
		device.Indicators.Set(adv360pro.IndicatorNumLock, nlck)
		device.Indicators.Set(adv360pro.IndicatorScrollLock, slck)
		device.Indicators.Sync()
	}
	return leds
}
