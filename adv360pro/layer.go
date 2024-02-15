package adv360pro

import (
	"github.com/bgould/keyboard-firmware/keyboard"
	"github.com/bgould/keyboard-firmware/keyboard/keycodes"
)

func Layer(
	/*-left--*/ k40, k41, k42, k43, k44, k45, k46, /* ----- */
	/*       */ k30, k31, k32, k33, k34, k35, k36, /*       */
	/*       */ k20, k21, k22, k23, k24, k25, k26, /*       */
	/*       */ k10, k11, k12, k13, k14, k15, /*   /*       */
	/*       */ k00, k01, k02, k03, k04, /*        /*       */
	/*       *\                       */ k28, k29, /*       */
	/*       *\                            */ k19, /*       */
	/*-------*\------------------*/ k07, k08, k09, /*-------*/
	/*-right-*/ k4D, k4E, k4F, k4G, k4H, k4I, k4J, /*-------*/
	/*       */ k3D, k3E, k3F, k3G, k3H, k3I, k3J, /*       */
	/*       */ k2D, k2E, k2F, k2G, k2H, k2I, k2J, /*       */
	/*       *\   */ k1E, k1F, k1G, k1H, k1I, k1J, /*       */
	/*       *\        */ k0F, k0G, k0H, k0I, k0J, /*       */
	/*       */ k2A, k2B, /*                       /*       */
	/*       */ k1A, /*                            /*       */
	/*-------*/ k0A, k0B, k0C /*---*/ keycodes.Keycode, /*--*/
) keyboard.Layer {
	return keyboard.Layer([][]keycodes.Keycode{
		/*        0x0  0x1  0x2  0x3  0x4  0x5  0x6  0x7  0x8  0x9  0xA  0xB  0xC  0xD  0xE  0xF  0xG  0xH  0xI  0xJ */
		/*************************************************************************************************************/
		/*  0 */ {k00, k01, k02, k03, k04, 0x0, 0x0, k07, k08, k09, k0A, k0B, k0C, 0x0, 0x0, k0F, k0G, k0H, k0I, k0J},
		/*  1 */ {k10, k11, k12, k13, k14, k15, 0x0, 0x0, 0x0, k19, k1A, 0x0, 0x0, 0x0, k1E, k1F, k1G, k1H, k1I, k1J},
		/*  2 */ {k20, k21, k22, k23, k24, k25, k26, 0x0, k28, k29, k2A, k2B, 0x0, k2D, k2E, k2F, k2G, k2H, k2I, k2J},
		/*  3 */ {k30, k31, k32, k33, k34, k35, k36, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, k3D, k3E, k3F, k3G, k3H, k3I, k3J},
		/*  4 */ {k40, k41, k42, k43, k44, k45, k46, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, k4D, k4E, k4F, k4G, k4H, k4I, k4J},
		/*************************************************************************************************************/
	})
}
