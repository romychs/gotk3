// +build !glib_2_40,!glib_2_42,!glib_2_44

// See: https://developer.gnome.org/glib/2.46/api-index-2-46.html

package glib

// #cgo pkg-config: glib-2.0 gobject-2.0
// #include <gio/gio.h>
// #include <glib.h>
// #include <glib-object.h>
// #include "glib.go.h"
import "C"

func (v *SettingsSchema) ListKeys() []string {
	c := C.g_settings_schema_list_keys(v.native())
	// both pointer array and strings should be freed.
	defer C.g_strfreev(c)

	strs := goStringArray(c)
	return strs
}
