// +build linux freebsd
// +build !no_wayland

package gdk

// #cgo pkg-config: gdk-wayland-3.0
// #include <gdk/gdk.h>
// #include <gdk/gdkwayland.h>
// #include "gdk_wayland.go.h"
import "C"
import (
	"unsafe"
)

// IsWaylandDisplay is a wrapper around _gdk_is_wayland_display().
func IsWaylandDisplay(display *Display) bool {
	c := C._gdk_is_wayland_display(unsafe.Pointer(display.native()))
	return gobool(c)
}
