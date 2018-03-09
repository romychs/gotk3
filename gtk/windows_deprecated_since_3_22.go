// Same copyright and license as the rest of the files in this project
// This file contains accelerator related functions and structures

//+build gtk_3_6 gtk_3_8 gtk_3_10 gtk_3_12 gtk_3_14 gtk_3_16 gtk_3_18 gtk_3_20

package gtk

// #cgo pkg-config: gtk+-3.0
// #include <stdlib.h>
// #include <gtk/gtk.h>
// #include "gtk_since_3_10.go.h"
import "C"
import "unsafe"

/*
 * GtkFontButton
 */

// GetFontName is a wrapper around gtk_font_button_get_font_name().
func (v *FontButton) GetFontName() string {
	c := C.gtk_font_button_get_font_name(v.native())
	return goString(c)
}

// SetFontName is a wrapper around gtk_font_button_set_font_name().
func (v *FontButton) SetFontName(fontname string) bool {
	cstr := C.CString(fontname)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_font_button_set_font_name(v.native(), (*C.gchar)(cstr))
	return gobool(c)
}
