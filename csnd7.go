package csnd7

/*
#cgo CFLAGS: -DUSE_DOUBLE=1
#cgo CFLAGS: -I /usr/local/include
#cgo linux CFLAGS: -DLINUX=1
#cgo LDFLAGS: -lcsound64

#include <csound/csound.h>

*/
import "C"
import "unsafe"

// Encapsulates an opaque pointer to a Csound instance
type CSOUND struct {
	Cs (*C.CSOUND)
}

type MYFLT float64

/*
 * Instantiation
 */

// Initialize Csound library with specific flags. This function is called
// internally by csoundCreate(), so there is// generally no need to use it
// explicitly unless you need to avoid default initialization that sets
// signal handlers and atexit() callbacks.
// Return value is zero on success, positive if initialisation was
// done already, and negative on error.
func Initialize(flags int) int {
	return int(C.csoundInitialize(C.int32(flags)))
}

// Creates an instance of Csound.  Returns an opaque pointer that
// must be passed to most Csound API functions.  The hostData
// parameter can be nil, or it can be a pointer to any sort of
// data; this pointer can be accessed from the Csound instance
// that is passed to callback routines.
// If not an empty string the opcodedir parameter sets an override
// for the plugin module/opcode directory search.
func Create(hostData unsafe.Pointer, opcodeDir string) CSOUND {
	var cs (*C.CSOUND)
	if opcodeDir == "" {
		cs = C.csoundCreate(hostData, nil)
	} else {
		var cstr *C.char = C.CString(opcodeDir)
		cs = C.csoundCreate(nil, cstr)
		C.free(unsafe.Pointer(cstr))
	}
	return CSOUND{cs}
}

// Destroy an instance of Csound.
func (csound *CSOUND) Destroy() {
	C.csoundDestroy(csound.Cs)
	csound.Cs = nil
}
