// +build gtk_3_6 gtk_3_8 gtk_3_10 gtk_3_12 gtk_3_14 gtk_3_16 gtk_3_18 gtk_3_20

package gtk

// #cgo pkg-config: gtk+-3.0
// #include <stdlib.h>
// #include <gtk/gtk.h>
import "C"
import (
	"errors"
	"unsafe"

	"github.com/d2r2/gotk3/gdk"
)

// ShowUri is a wrapper for gtk_show_uri().
func ShowUri(screen *gdk.Screen, uri string) error {
	cstr := C.CString(uri)
	defer C.free(unsafe.Pointer(cstr))

	var err *C.GError

	c := C.gtk_show_uri(C.toGdkScreen(screen.Native()),
		cstr, C.gtk_get_current_event_time(), &err)
	if !gobool(c) {
		defer C.g_error_free(err)
		return errors.New(goString(err.message))
	}
	return nil
}
