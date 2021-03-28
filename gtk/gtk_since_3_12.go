// +build !gtk_3_6,!gtk_3_8,!gtk_3_10
// not use this: go build -tags gtk_3_8'. Otherwise, if no build tags are used, GTK 3.10

// See: https://developer.gnome.org/gtk3/3.12/api-index-3-12.html

package gtk

// #cgo pkg-config: gtk+-3.0
// #include <gtk/gtk.h>
import "C"

import (
	"unsafe"

	"github.com/romychs/gotk3/glib"
)

func init() {
	tm := []glib.TypeMarshaler{
		// Objects/Interfaces
		{glib.Type(C.gtk_flow_box_get_type()), marshalFlowBox},
		{glib.Type(C.gtk_flow_box_child_get_type()), marshalFlowBoxChild},
	}
	glib.RegisterGValueMarshalers(tm)

	WrapMap["GtkFlowBox"] = wrapFlowBox
	WrapMap["GtkFlowBoxChild"] = wrapFlowBoxChild
}

// SetPopover is a wrapper around gtk_menu_button_set_popover().
func (v *MenuButton) SetPopover(popover *Popover) {
	C.gtk_menu_button_set_popover(v.native(), popover.toWidget())
}

// GetPopover is a wrapper around gtk_menu_button_get_popover().
func (v *MenuButton) GetPopover() *Popover {
	c := C.gtk_menu_button_get_popover(v.native())
	if c == nil {
		return nil
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapPopover(obj)
}
