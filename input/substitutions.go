package input

import "C"

var keyPress = map[string]string{
	"Alt_L":            "<alt_left>",
	"apostrophe":       "'",
	"backslash":        "\\",
	"BackSpace":        "<backspace>",
	"bracketleft":      "[",
	"bracketright":     "]",
	"Caps_Lock":        "<caps_lock>",
	"comma":            ",",
	"Control_L":        "<ctrl_left>",
	"Control_R":        "<ctrl_right>",
	"Delete":           "<delete>",
	"Down":             "<down>",
	"End":              "<end>",
	"equal":            "=",
	"Escape":           "<escape>",
	"F1":               "<f1>",
	"F10":              "<f10>",
	"F11":              "<f11>",
	"F12":              "<f12>",
	"F2":               "<f2>",
	"F3":               "<f3>",
	"F4":               "<f4>",
	"F5":               "<f5>",
	"F6":               "<f6>",
	"F7":               "<f7>",
	"F8":               "<f8>",
	"F9":               "<f9>",
	"grave":            "`",
	"Home":             "<home>",
	"Insert":           "<insert>",
	"ISO_Level3_Shift": "<alt_gr>",
	"KP_Add":           "+",
	"KP_Begin":         "<kp_5>",
	"KP_Delete":        "<kp_del>",
	"KP_Divide":        "/",
	"KP_Down":          "<kp_2>",
	"KP_End":           "<kp_1>",
	"KP_Enter":         "\n",
	"KP_Home":          "<kp_7>",
	"KP_Insert":        "<kp_0>",
	"KP_Left":          "<kp_4>",
	"KP_Multiply":      "*",
	"KP_Next":          "<kp_3>",
	"KP_Prior":         "<kp_9>",
	"KP_Right":         "<kp_6>",
	"KP_Subtract":      "-",
	"KP_Up":            "<kp_8>",
	"Left":             "<left>",
	"Menu":             "<menu>",
	"minus":            "-",
	"Next":             "<page_down>",
	"Num_Lock":         "<num_lock>",
	"numbersign":       "#",
	"Pause":            "<pause>",
	"period":           ".",
	"Print":            "<print_screen>",
	"Prior":            "<page_up>",
	"Return":           "\n",
	"Right":            "<right>",
	"Scroll_Lock":      "<scroll_lock>",
	"semicolon":        ";",
	"Shift_L":          "<shift_left>",
	"Shift_R":          "<shift_right>",
	"slash":            "/",
	"space":            " ",
	"Super_L":          "<meta_left>",
	"Super_R":          "<meta_right>",
	"Tab":              "\t",
	"Up":               "<up>",
}

var keyRelease = map[string]string{
	"Alt_L":     "</alt_left>",
	"Control_L": "</ctrl_left>",
	"Control_R": "</ctrl_right>",
	"Shift_L":   "</shift_left>",
	"Shift_R":   "</shift_right>",
	"Super_L":   "</meta_left>",
	"Super_R":   "</meta_right>",
}

// SubstituteKeyPress substitutes a key purely for readability in the log.
func substituteKeyPress(key *C.char) string {
	k := C.GoString(key)
	s, ok := keyPress[k]
	if !ok {
		return k
	}
	return s
}

// SubstituteKeyRelease substitutes a key purely for readability in the log.
func substituteKeyRelease(key *C.char) string {
	k := C.GoString(key)
	s, ok := keyRelease[k]
	if !ok {
		return ""
	}
	return s
}
