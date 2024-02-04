package kinadv360pro

import (
	"github.com/bgould/keyboard-firmware/keyboard"
	"github.com/bgould/keyboard-firmware/keyboard/keycodes"
)

func Layer(
	k00, k10, k20, k30, k40, k50, kC0, /*  */
	k01, k11, k21, k31, k41, k51, kD0, /*  */
	k02, k12, k22, k32, k42, k52, kE0, /*  */
	k03, k13, k23, k33, k43, k53,
	kD2, k14, k24, k34, k54,
	/*                    */ k56, k55, /*  */
	/*                         */ k35, /*  */
	/*               */ k36, k46, k25, /*  */
	kC1, k60, k70, k80, k90, kA0, kB0,
	kD1, k61, k71, k81, k91, kA1, kB1,
	kE1, k62, k72, k82, k92, kA2, kB2,
	k63, k73, k83, k93, kA3, kB3,
	k64, k84, k94, kA4, kE2,
	k96, k85,
	k86,
	k66, k75 /**/, k65 keycodes.Keycode,
) keyboard.Layer {
	return keyboard.Layer([][]keycodes.Keycode{
		/*        0x0  0x1  0x2  0x3  0x4  0x5  0x6  */
		/*********************************************/
		/*  0 */ {k00, k01, k02, k03, 0x0, 0x0, 0x0},
		/*  1 */ {k10, k11, k12, k13, k14, 0x0, 0x0},
		/*  2 */ {k20, k21, k22, k23, k24, k25, 0x0},
		/*  3 */ {k30, k31, k32, k33, k34, k35, k36},
		/*  4 */ {k40, k41, k42, k43, 0x0, 0x0, k46},
		/*  5 */ {k50, k51, k52, k53, k54, k55, k56},
		/*  6 */ {k60, k61, k62, k63, k64, k65, k66},
		/*  7 */ {k70, k71, k72, k73, 0x0, k75, 0x0},
		/*  8 */ {k80, k81, k82, k83, k84, k85, k86},
		/*  9 */ {k90, k91, k92, k93, k94, 0x0, k96},
		/* 10 */ {kA0, kA1, kA2, kA3, kA4, 0x0, 0x0},
		/* 11 */ {kB0, kB1, kB2, kB3, 0x0, 0x0, 0x0},
		/* 12 */ {kC0, kC1, 0x0, 0x0, 0x0, 0x0, 0x0},
		/* 13 */ {kD0, kD1, kD2, 0x0, 0x0, 0x0, 0x0},
		/* 14 */ {kE0, kE1, kE2, 0x0, 0x0, 0x0, 0x0},
		/*********************************************/
	})
}
