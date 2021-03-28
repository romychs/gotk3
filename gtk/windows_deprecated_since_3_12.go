// Same copyright and license as the rest of the files in this project
// This file contains accelerator related functions and structures

// +build gtk_3_6 gtk_3_8 gtk_3_10

package gtk

// #cgo pkg-config: gtk+-3.0
// #include <stdlib.h>
// #include <gtk/gtk.h>
// #include "gtk_since_3_10.go.h"
import "C"
import (
	"unsafe"

	"github.com/romychs/gotk3/glib"
)

/*
 * GtkDialog
 */

// GetActionArea is a wrapper around gtk_dialog_get_action_area().
func (v *Dialog) GetActionArea() (*Widget, error) {
	c := C.gtk_dialog_get_action_area(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapWidget(obj), nil
}

/*
 * GtkMessageDialog
 */

// GetImage is a wrapper around gtk_message_dialog_get_image().
func (v *MessageDialog) GetImage() (*Widget, error) {
	c := C.gtk_message_dialog_get_image(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapWidget(obj), nil
}

// SetImage is a wrapper around gtk_message_dialog_set_image().
func (v *MessageDialog) SetImage(image IWidget) {
	C.gtk_message_dialog_set_image(v.native(), image.toWidget())
}
