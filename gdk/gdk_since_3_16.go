// +build !gtk_3_6,!gtk_3_8,!gtk_3_10,!gtk_3_12,!gtk_3_14
// not use this: go build -tags gtk_3_8'. Otherwise, if no build tags are used, GDK 3.20

// Go bindings for GDK 3.  Supports version 3.6 and later.
package gdk

// #cgo pkg-config: gdk-3.0
// #include <gdk/gdk.h>
// #include "gdk.go.h"
// #include "gdk_since_3_20.go.h"
import "C"

const (
	GRAB_FAILED GrabStatus = C.GDK_GRAB_FAILED
)
