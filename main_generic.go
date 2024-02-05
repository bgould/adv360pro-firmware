//go:build nrf52840_generic

package main

import (
	"machine"
)

func init() {
	machine.LED = machine.P1_14
	machine.UART_TX_PIN = machine.P0_14 // PORTB
	machine.UART_RX_PIN = machine.P0_30 // PORTB
}
