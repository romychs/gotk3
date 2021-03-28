// +build !gtk_3_6,!gtk_3_8,!gtk_3_10,!gtk_3_12

// See: https://developer.gnome.org/gtk3/3.14/api-index-3-14.html

package gtk

// #include <gtk/gtk.h>
// #include "gtk.go.h"
import "C"
import (
	"unsafe"

	"github.com/romychs/gotk3/glib"
)

// PrefersAppMenu is a wrapper around gtk_application_prefers_app_menu().
func (v *Application) PrefersAppMenu() bool {
	return gobool(C.gtk_application_prefers_app_menu(v.native()))
}

// GetActionsForAccel is a wrapper around gtk_application_get_actions_for_accel().
func (v *Application) GetActionsForAccel(acc string) []string {
	cstr := C.CString(acc)
	defer C.free(unsafe.Pointer(cstr))

	c := C.gtk_application_get_actions_for_accel(v.native(), (*C.gchar)(cstr))
	// both pointer array and strings also should be freed.
	defer C.g_strfreev(c)

	acts := goStringArray(c)
	return acts
}

// GetMenuByID is a wrapper around gtk_application_get_menu_by_id().
func (v *Application) GetMenuByID(id string) *glib.Menu {
	cstr := C.CString(id)
	defer C.free(unsafe.Pointer(cstr))

	c := C.gtk_application_get_menu_by_id(v.native(), (*C.gchar)(cstr))
	if c == nil {
		return nil
	}
	return &glib.Menu{glib.MenuModel{glib.Take(unsafe.Pointer(c))}}
}
