package glib

// #cgo pkg-config: glib-2.0 gobject-2.0
// #include <gio/gio.h>
// #include <glib.h>
// #include <glib-object.h>
import "C"
import "unsafe"

// Abstract respresentation for any GTK+ GInterface.
// Should be used instead of Object type in all GInterface
// implementations.
type Interface struct {
	ginterface C.gpointer
}

func InterfaceNew(ptr unsafe.Pointer) *Interface {
	p := unsafe.Pointer(ptr)
	v := &Interface{(C.gpointer)(p)}
	return v
}

func InterfaceFromObjectNew(obj *Object) *Interface {
	p := unsafe.Pointer(obj.Native())
	return InterfaceNew(p)
}

func (v *Interface) Native() uintptr {
	if v == nil {
		return 0
	}
	return uintptr(unsafe.Pointer(v.ginterface))
}
