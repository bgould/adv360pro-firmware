package kinadv360pro

import "machine"

var LeftRows = [5]machine.Pin{
	machine.P1_11,
	machine.P1_15,
	machine.P0_03,
	machine.P1_14,
	machine.P1_12,
}

/*
   diode-direction = "col2row";
   row-gpios
       = <&gpio1 11 (GPIO_ACTIVE_HIGH | GPIO_PULL_DOWN)>
       , <&gpio1 15 (GPIO_ACTIVE_HIGH | GPIO_PULL_DOWN)>
       , <&gpio0 3 (GPIO_ACTIVE_HIGH | GPIO_PULL_DOWN)>
       , <&gpio1 14 (GPIO_ACTIVE_HIGH | GPIO_PULL_DOWN)>
       , <&gpio1 12 (GPIO_ACTIVE_HIGH | GPIO_PULL_DOWN)>
       ;
*/

var LeftCols = [10]machine.Pin{
	machine.P0_25,
	machine.P0_11,
	machine.P0_02,
	machine.P0_28,
	machine.P0_29,
	machine.P0_30,
	machine.P0_31,
	machine.P1_09,
	machine.P0_12,
	machine.P0_07,
}

/*
   col-gpios
       = <&gpio0 25 GPIO_ACTIVE_HIGH>
       , <&gpio0 11 GPIO_ACTIVE_HIGH>
       , <&gpio0 2 GPIO_ACTIVE_HIGH>
       , <&gpio0 28 GPIO_ACTIVE_HIGH>
       , <&gpio0 29 GPIO_ACTIVE_HIGH>
       , <&gpio0 30 GPIO_ACTIVE_HIGH>
       , <&gpio0 31 GPIO_ACTIVE_HIGH>
       , <&gpio1 9 GPIO_ACTIVE_HIGH>
       , <&gpio0 12 GPIO_ACTIVE_HIGH>
       , <&gpio0 7 GPIO_ACTIVE_HIGH>
       ;
*/
