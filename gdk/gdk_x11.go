// +build linux freebsd
// +build !no_x11

package gdk

// #cgo pkg-config: gdk-x11-3.0
// #include <gdk/gdk.h>
// #include <gdk/gdkx.h>
// #include "gdk_x11.go.h"
import "C"
import (
	"unsafe"
)

// IsX11Display is a wrapper around _gdk_is_x11_display().
func IsX11Display(display *Display) bool {
	c := C._gdk_is_x11_display(unsafe.Pointer(display.native()))
	return gobool(c)
}
