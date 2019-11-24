// Same copyright and license as the rest of the files in this project
package gdk

// #cgo pkg-config: gdk-3.0
// #include <gdk/gdk.h>
// #include "gdk.go.h"
import "C"

// GetMajorVersion returns major version number.
// In general GDK repeats exactly GTK version numbering.
func GetMajorVersion() uint {
	v := C._gdk_major_version()
	return uint(v)
}

// GetMinorVersion returns minor version number.
// In general GDK repeats exactly GTK version numbering.
func GetMinorVersion() uint {
	v := C._gdk_minor_version()
	return uint(v)
}

// GetMicroVersion returns micro version number.
// In general GDK repeats exactly GTK version numbering.
func GetMicroVersion() uint {
	v := C._gdk_micro_version()
	return uint(v)
}
