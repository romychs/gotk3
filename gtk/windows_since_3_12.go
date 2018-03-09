// +build !gtk_3_6,!gtk_3_8,!gtk_3_10

package gtk

// #cgo pkg-config: gtk+-3.0
// #include <gtk/gtk.h>
import "C"

const (
	DIALOG_USE_HEADER_BAR DialogFlags = C.GTK_DIALOG_USE_HEADER_BAR
)

// IsMaximized is a wrapper around gtk_window_is_maximized().
func (v *Window) IsMaximized() bool {
	c := C.gtk_window_is_maximized(v.native())
	return gobool(c)
}
