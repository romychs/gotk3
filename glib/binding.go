package glib

// #cgo pkg-config: glib-2.0 gobject-2.0
// #include <gio/gio.h>
// #include <glib.h>
// #include <glib-object.h>
// #include "glib.go.h"
import "C"
import "unsafe"

type BindingFlags int

const (
	BINDING_DEFAULT        BindingFlags = C.G_BINDING_DEFAULT
	BINDING_BIDIRECTIONAL               = C.G_BINDING_BIDIRECTIONAL
	BINDING_SYNC_CREATE                 = C.G_BINDING_SYNC_CREATE
	BINDING_INVERT_BOOLEAN              = C.G_BINDING_INVERT_BOOLEAN
)

// Binding is a representation of Glib's GBinding.
type Binding struct {
	*Object
}

func (v *Binding) native() *C.GBinding {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toGBinding(ptr)
}

func marshalBinding(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	return &Binding{wrapObject(unsafe.Pointer(c))}, nil
}

// BindProperty creates a binding between source property on source and target property on
// target . Whenever the source property is changed the target_property is
// updated using the same value.
func BindProperty(source *Object, sourceProperty string,
	target *Object, targetProperty string,
	flags BindingFlags) *Binding {
	srcStr := C.CString(sourceProperty)
	defer C.free(unsafe.Pointer(srcStr))
	tgtStr := C.CString(targetProperty)
	defer C.free(unsafe.Pointer(tgtStr))
	obj := C.g_object_bind_property(
		C.gpointer(source.Native()), (*C.gchar)(srcStr),
		C.gpointer(target.Native()), (*C.gchar)(tgtStr),
		C.GBindingFlags(flags),
	)
	if obj == nil {
		return nil
	}
	return &Binding{wrapObject(unsafe.Pointer(obj))}
}

// Unbind explicitly releases the binding between the source and the target property
// expressed by Binding
func (v *Binding) Unbind() {
	C.g_binding_unbind(v.native())
}

// GetSource retrieves the GObject instance used as the source of the binding
func (v *Binding) GetSource() *Object {
	obj := C.g_binding_get_source(v.native())
	if obj == nil {
		return nil
	}
	return wrapObject(unsafe.Pointer(obj))
}

// GetSourceProperty retrieves the name of the property of “source” used as the source of
// the binding.
func (v *Binding) GetSourceProperty() string {
	c := C.g_binding_get_source_property(v.native())
	return goString(c)
}

// GetTarget retrieves the GObject instance used as the target of the binding.
func (v *Binding) GetTarget() *Object {
	obj := C.g_binding_get_target(v.native())
	if obj == nil {
		return nil
	}
	return wrapObject(unsafe.Pointer(obj))
}

// GetTargetProperty retrieves the name of the property of “target” used as the target of
// the binding.
func (v *Binding) GetTargetProperty() string {
	c := C.g_binding_get_target_property(v.native())
	return goString(c)
}

// GetFlags retrieves the flags passed when constructing the GBinding.
func (v *Binding) GetFlags() BindingFlags {
	flags := C.g_binding_get_flags(v.native())
	return BindingFlags(flags)
}
