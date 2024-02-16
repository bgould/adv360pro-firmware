package main

import (
	"github.com/bgould/keyboard-firmware/keyboard"
	"github.com/bgould/kinadv360pro-firmware/adv360pro"

	. "github.com/bgould/keyboard-firmware/keyboard/keycodes"
)

const (
	_______ = KC_NO
)

// FN0 -> Tap to Reset, Hold 2+ secs for Bootloader
// FN1 -> Toggle "Function" layer on and off on key up/down
// FN2 -> Toggle "Programming" layer on and off on key up/down

func initKeymap() keyboard.Keymap {
	return keyboard.Keymap([]keyboard.Layer{
		// 0 - Default Layer
		adv360pro.Layer(
			/*******************************************************************************************************************************************/
			/*******************************************************************************************************************************************/
			/**/ KC_EQL_, KC_N1__, KC_N2__, KC_N3__, KC_N4__, KC_N5__, KC_FN2_ /* */, KC_FN2_, KC_N6__, KC_N7__, KC_N8__, KC_N9__, KC_N0__, KC_MINS, /**/
			/**/ KC_TAB_, KC_Q___, KC_W___, KC_E___, KC_R___, KC_T___, KC_F1__ /* */, KC_F5__, KC_Y___, KC_U___, KC_I___, KC_O___, KC_P___, KC_BSLS, /**/
			/**/ KC_RCTL, KC_A___, KC_S___, KC_D___, KC_F___, KC_G___, KC_F2__ /* */, KC_PSCR, KC_H___, KC_J___, KC_K___, KC_L___, KC_SCLN, KC_QUOT, /**/
			/**/ KC_LSFT, KC_Z___, KC_X___, KC_C___, KC_V___, KC_B___ /*                   */, KC_N___, KC_M___, KC_COMM, KC_DOT_, KC_SLSH, KC_RSFT, /**/
			/**/ KC_FN1_, KC_GRV_, KC_INS_, KC_LEFT, KC_RGHT /*                                     */, KC_UP__, KC_DOWN, KC_LBRC, KC_RBRC, KC_FN1_, /**/
			/**\                                           */ KC_ESC_, KC_LGUI /* */, KC_LALT, KC_LCTL, /*                                           /**/
			/**\                                                    */ KC_HOME /* */, KC_PGUP, /*                                                    /**/
			/**\                                  */ KC_BSPC, KC_DEL_, KC_END_ /* */, KC_PGDN, KC_ENT_, KC_SPC_, /*                                  /**/
			/*******************************************************************************************************************************************/
			/*******************************************************************************************************************************************/
		),
		// 1 - Function Layer
		adv360pro.Layer(
			/*******************************************************************************************************************************************/
			/*******************************************************************************************************************************************/
			/**/ KC_F1__, KC_F2__, KC_F3__, KC_F4__, KC_F5__, KC_F6__, KC_FN2_ /* */, KC_FN2_, KC_F7__, KC_F8__, KC_F9__, KC_F10_, KC_F11_, KC_F12_, /**/
			/**/ _______, _______, _______, _______, _______, _______, _______ /* */, _______, _______, _______, _______, _______, _______, _______, /**/
			/**/ _______, _______, _______, _______, _______, _______, _______ /* */, _______, _______, _______, _______, _______, _______, _______, /**/
			/**/ _______, _______, _______, _______, _______, _______ /*                   */, _______, _______, _______, _______, _______, _______, /**/
			/**/ KC_FN1_, _______, _______, _______, _______ /*                                     */, _______, _______, _______, _______, KC_FN1_, /**/
			/**\                                           */ _______, _______ /* */, _______, _______, /*                                           /**/
			/**\                                                    */ _______ /* */, _______, /*                                                    /**/
			/**\                                  */ _______, _______, _______ /* */, _______, _______, _______, /*                                  /**/
			/*******************************************************************************************************************************************/
			/*******************************************************************************************************************************************/
		),
		// 2 - Programming Layer
		adv360pro.Layer(
			/*******************************************************************************************************************************************/
			/*******************************************************************************************************************************************/
			/**/ _______, _______, _______, _______, _______, _______, KC_FN2_ /* */, KC_FN2_, _______, _______, _______, _______, _______, _______, /**/
			/**/ _______, _______, _______, _______, _______, _______, _______ /* */, _______, _______, _______, _______, _______, _______, _______, /**/
			/**/ _______, _______, _______, _______, _______, _______, KC_FN0_ /* */, KC_FN0_, _______, _______, _______, _______, _______, _______, /**/
			/**/ _______, _______, _______, _______, _______, _______ /*                   */, _______, _______, _______, _______, _______, _______, /**/
			/**/ KC_FN1_, _______, _______, _______, _______ /*                                     */, _______, _______, _______, _______, KC_FN1_, /**/
			/**\                                           */ _______, _______ /* */, _______, _______, /*                                           /**/
			/**\                                                    */ _______ /* */, _______, /*                                                    /**/
			/**\                                  */ _______, _______, _______ /* */, _______, _______, _______, /*                                  /**/
			/*******************************************************************************************************************************************/
			/*******************************************************************************************************************************************/
		),
		// 4 - Empty Layer
		adv360pro.Layer(
			/*******************************************************************************************************************************************/
			/*******************************************************************************************************************************************/
			/**/ _______, _______, _______, _______, _______, _______, _______ /* */, _______, _______, _______, _______, _______, _______, _______, /**/
			/**/ _______, _______, _______, _______, _______, _______, _______ /* */, _______, _______, _______, _______, _______, _______, _______, /**/
			/**/ _______, _______, _______, _______, _______, _______, _______ /* */, _______, _______, _______, _______, _______, _______, _______, /**/
			/**/ _______, _______, _______, _______, _______, _______ /*                   */, _______, _______, _______, _______, _______, _______, /**/
			/**/ _______, _______, _______, _______, _______ /*                                     */, _______, _______, _______, _______, _______, /**/
			/**\                                           */ _______, _______ /* */, _______, _______, /*                                           /**/
			/**\                                                    */ _______ /* */, _______, /*                                                    /**/
			/**\                                  */ _______, _______, _______ /* */, _______, _______, _______, /*                                  /**/
			/*******************************************************************************************************************************************/
			/*******************************************************************************************************************************************/
		),
		// 5 - Empty Layer
		adv360pro.Layer(
			/*******************************************************************************************************************************************/
			/**/ _______, _______, _______, _______, _______, _______, _______ /* */, _______, _______, _______, _______, _______, _______, _______, /**/
			/**/ _______, _______, _______, _______, _______, _______, _______ /* */, _______, _______, _______, _______, _______, _______, _______, /**/
			/**/ _______, _______, _______, _______, _______, _______, _______ /* */, _______, _______, _______, _______, _______, _______, _______, /**/
			/**/ _______, _______, _______, _______, _______, _______ /*                   */, _______, _______, _______, _______, _______, _______, /**/
			/**/ _______, _______, _______, _______, _______ /*                                     */, _______, _______, _______, _______, _______, /**/
			/**\                                           */ _______, _______ /* */, _______, _______, /*                                           /**/
			/**\                                                    */ _______ /* */, _______, /*                                                    /**/
			/**\                                  */ _______, _______, _______ /* */, _______, _______, _______, /*                                  /**/
			/*******************************************************************************************************************************************/
		),
		// n -Layer
		// adv360pro.Layer(
		//	/*******************************************************************************************************************************************/
		// 	/*******************************************************************************************************************************************/
		// 	/**/ _______, _______, _______, _______, _______, _______, _______, /* */ _______, _______, _______, _______, _______, _______, _______, /**/
		// 	/**/ _______, _______, _______, _______, _______, _______, _______, /* */ _______, _______, _______, _______, _______, _______, _______, /**/
		// 	/**/ _______, _______, _______, _______, _______, _______, _______, /* */ _______, _______, _______, _______, _______, _______, _______, /**/
		// 	/**/ _______, _______, _______, _______, _______, _______, /*                   */ _______, _______, _______, _______, _______, _______, /**/
		// 	/**/ _______, _______, _______, _______, _______, /*                                     */ _______, _______, _______, _______, _______, /**/
		// 	/**\                                           */ _______, _______, /* */ _______, _______, /*                                           /**/
		// 	/**\                                                    */ _______, /* */ _______, /*                                                    /**/
		// 	/**\                                  */ _______, _______, _______, /* */ _______, _______, _______, /*                                  /**/
		// 	/*******************************************************************************************************************************************/
		//	/*******************************************************************************************************************************************/
		// ),
	})
}
