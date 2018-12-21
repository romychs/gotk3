package glib

// #cgo pkg-config: glib-2.0 gobject-2.0
// #include <gio/gio.h>
// #include <glib.h>
// #include <glib-object.h>
// #include "glib.go.h"
import "C"
import "unsafe"

// Notification is a representation of GNotification.
type Notification struct {
	*Object
}

// native() returns a pointer to the underlying GNotification.
func (v *Notification) native() *C.GNotification {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toGNotification(ptr)
}

func marshalNotification(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	return wrapNotification(wrapObject(unsafe.Pointer(c))), nil
}

func wrapNotification(obj *Object) *Notification {
	return &Notification{obj}
}

// NotificationNew is a wrapper around g_notification_new().
func NotificationNew(title string) (*Notification, error) {
	cstr := C.CString(title)
	defer C.free(unsafe.Pointer(cstr))

	c := C.g_notification_new((*C.gchar)(cstr))
	if c == nil {
		return nil, errNilPtr
	}
	return wrapNotification(wrapObject(unsafe.Pointer(c))), nil
}

// SetTitle is a wrapper around g_notification_set_title().
func (v *Notification) SetTitle(title string) {
	cstr := C.CString(title)
	defer C.free(unsafe.Pointer(cstr))

	C.g_notification_set_title(v.native(), (*C.gchar)(cstr))
}

// SetBody is a wrapper around g_notification_set_body().
func (v *Notification) SetBody(body string) {
	cstr := C.CString(body)
	defer C.free(unsafe.Pointer(cstr))

	C.g_notification_set_body(v.native(), (*C.gchar)(cstr))
}

// SetIcon is a wrapper around g_notification_set_icon
func (v *Notification) SetIcon(icon *Icon) {
	C.g_notification_set_icon(v.native(), icon.native())
}

// SetDefaultAction is a wrapper around g_notification_set_default_action().
func (v *Notification) SetDefaultAction(detailedAction string) {
	cstr := C.CString(detailedAction)
	defer C.free(unsafe.Pointer(cstr))

	C.g_notification_set_default_action(v.native(), (*C.gchar)(cstr))
}

// AddButton is a wrapper around g_notification_add_button().
func (v *Notification) AddButton(label, detailedAction string) {
	cstr1 := C.CString(label)
	defer C.free(unsafe.Pointer(cstr1))

	cstr2 := C.CString(detailedAction)
	defer C.free(unsafe.Pointer(cstr2))

	C.g_notification_add_button(v.native(), (*C.gchar)(cstr1), (*C.gchar)(cstr2))
}

// void 	g_notification_set_default_action_and_target () // requires varargs
// void 	g_notification_set_default_action_and_target_value () // requires variant
// void 	g_notification_add_button_with_target () // requires varargs
// void 	g_notification_add_button_with_target_value () //requires variant
// void 	g_notification_set_urgent () // Deprecated, so not implemented
// void 	g_notification_set_icon () // Requires support for GIcon, which we don't have yet.
