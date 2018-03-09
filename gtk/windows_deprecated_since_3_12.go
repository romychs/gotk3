// Same copyright and license as the rest of the files in this project
// This file contains accelerator related functions and structures

// +build gtk_3_6 gtk_3_8 gtk_3_10

package gtk

// #cgo pkg-config: gtk+-3.0
// #include <stdlib.h>
// #include <gtk/gtk.h>
// #include "gtk_since_3_10.go.h"
import "C"

/*
 * GtkDialog
 */

// GetActionArea() is a wrapper around gtk_dialog_get_action_area().
func (v *Dialog) GetActionArea() (*Box, error) {
	c := C.gtk_dialog_get_action_area(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	box := wrapBox(obj)
	return box, nil
}

/*
 * GtkMessageDialog
 */

// Wrap around gtk_message_dialog_set_image()
func (v *MessageDialog) SetImage(image *Image) {
	C.gtk_message_dialog_set_image(v.native(), image.Widget.native())
}
