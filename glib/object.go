package glib

// #cgo pkg-config: glib-2.0 gobject-2.0
// #include <gio/gio.h>
// #include <glib.h>
// #include <glib-object.h>
// #include "glib.go.h"
import "C"
import (
	"errors"
	"fmt"
	"runtime"
	"unsafe"
)

/*
 * GObject
 */

// IObject is an interface type implemented by Object and all types which embed
// an Object.  It is meant to be used as a type for function arguments which
// require GObjects or any subclasses thereof.
type IObject interface {
	toObject() *Object
}

// Object is a representation of GLib's GObject.
type Object struct {
	gobject *C.GObject
}

// Cast underlying GObject to uintptr with protection from nil.
// No need to check for nil in successor code.
// Export this method for all other modules and successors
func (v *Object) Native() uintptr {
	return uintptr(unsafe.Pointer(v.native()))
}

// native returns a pointer to the underlying GObject.
func (v *Object) native() *C.GObject {
	if v == nil {
		return nil
	}
	return v.gobject
}

func (v *Object) toObject() *Object {
	return v
}

// newObject creates a new Object from a GObject pointer.
func newObject(p *C.GObject) *Object {
	return &Object{gobject: p}
}

// Take wraps a unsafe.Pointer as a glib.Object, taking ownership of it.
// This function is exported for visibility in other gotk3 packages and
// is not meant to be used by applications.
func Take(ptr unsafe.Pointer) *Object {
	obj := ToObject(ptr)

	if obj.IsFloating() {
		obj.RefSink()
	} else {
		obj.Ref()
	}

	runtime.SetFinalizer(obj, (*Object).Unref)
	return obj
}

// IsA is a wrapper around g_type_is_a().
func (v *Object) IsA(typ Type) bool {
	return gobool(C.g_type_is_a(C.GType(v.TypeFromInstance()), C.GType(typ)))
}

// TypeFromInstance is a wrapper around g_type_from_instance().
func (v *Object) TypeFromInstance() Type {
	c := C._g_type_from_instance(C.gpointer(v.native()))
	return Type(c)
}

// ToGObject type converts an unsafe.Pointer as a native C GObject.
// This function is exported for visibility in other gotk3 packages and
// is not meant to be used by applications.
func toGObject(ptr unsafe.Pointer) *C.GObject {
	return C.toGObject(ptr)
}

func ToObject(ptr unsafe.Pointer) *Object {
	return newObject(C.toGObject(ptr))
}

// Ref is a wrapper around g_object_ref().
func (v *Object) Ref() {
	C.g_object_ref(C.gpointer(v.gobject))
}

// Unref is a wrapper around g_object_unref().
func (v *Object) Unref() {
	C.g_object_unref(C.gpointer(v.gobject))
}

// RefSink is a wrapper around g_object_ref_sink().
func (v *Object) RefSink() {
	C.g_object_ref_sink(C.gpointer(v.gobject))
}

// IsFloating is a wrapper around g_object_is_floating().
func (v *Object) IsFloating() bool {
	c := C.g_object_is_floating(C.gpointer(v.gobject))
	return gobool(c)
}

// ForceFloating is a wrapper around g_object_force_floating().
func (v *Object) ForceFloating() {
	C.g_object_force_floating(v.gobject)
}

// StopEmission is a wrapper around g_signal_stop_emission_by_name().
func (v *Object) StopEmission(s string) {
	cstr := C.CString(s)
	defer C.free(unsafe.Pointer(cstr))
	C.g_signal_stop_emission_by_name(C.gpointer(v.gobject),
		(*C.gchar)(cstr))
}

// Set is a wrapper around g_object_set().  However, unlike
// g_object_set(), this function only sets one name value pair.  Make
// multiple calls to this function to set multiple properties.
func (v *Object) Set(name string, value interface{}) error {
	return v.SetProperty(name, value)
	/*
		cstr := C.CString(name)
		defer C.free(unsafe.Pointer(cstr))

		if _, ok := value.(Object); ok {
			value = value.(Object).GObject
		}

		// Can't call g_object_set() as it uses a variable arg list, use a
		// wrapper instead
		var p unsafe.Pointer
		switch v := value.(type) {
		case bool:
			c := gbool(v)
			p = unsafe.Pointer(&c)

		case int8:
			c := C.gint8(v)
			p = unsafe.Pointer(&c)

		case int16:
			c := C.gint16(v)
			p = unsafe.Pointer(&c)

		case int32:
			c := C.gint32(v)
			p = unsafe.Pointer(&c)

		case int64:
			c := C.gint64(v)
			p = unsafe.Pointer(&c)

		case int:
			c := C.gint(v)
			p = unsafe.Pointer(&c)

		case uint8:
			c := C.guchar(v)
			p = unsafe.Pointer(&c)

		case uint16:
			c := C.guint16(v)
			p = unsafe.Pointer(&c)

		case uint32:
			c := C.guint32(v)
			p = unsafe.Pointer(&c)

		case uint64:
			c := C.guint64(v)
			p = unsafe.Pointer(&c)

		case uint:
			c := C.guint(v)
			p = unsafe.Pointer(&c)

		case uintptr:
			p = unsafe.Pointer(C.gpointer(v))

		case float32:
			c := C.gfloat(v)
			p = unsafe.Pointer(&c)

		case float64:
			c := C.gdouble(v)
			p = unsafe.Pointer(&c)

		case string:
			cstr := C.CString(v)
			defer C.g_free(C.gpointer(unsafe.Pointer(cstr)))
			p = unsafe.Pointer(&cstr)

		default:
			if pv, ok := value.(unsafe.Pointer); ok {
				p = pv
			} else {
				val := reflect.ValueOf(value)
				switch val.Kind() {
				case reflect.Int, reflect.Int8, reflect.Int16,
					reflect.Int32, reflect.Int64:
					c := C.int(val.Int())
					p = unsafe.Pointer(&c)

				case reflect.Uintptr, reflect.Ptr, reflect.UnsafePointer:
					p = unsafe.Pointer(C.gpointer(val.Pointer()))
				}
			}
		}
		if p == nil {
			return errors.New("Unable to perform type conversion")
		}
		C._g_object_set_one(C.gpointer(v.GObject), (*C.gchar)(cstr), p)
		return nil*/
}

// GetPropertyType returns the Type of a property of the underlying GObject.
// If the property is missing it will return TYPE_INVALID and an error.
func (v *Object) GetPropertyType(name string) (Type, error) {
	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))

	paramSpec := C.g_object_class_find_property(C._g_object_get_class(v.native()),
		(*C.gchar)(cstr))
	if paramSpec == nil {
		return TYPE_INVALID, errors.New("couldn't find Property")
	}
	return Type(paramSpec.value_type), nil
}

// GetProperty is a wrapper around g_object_get_property().
func (v *Object) GetProperty(name string) (interface{}, error) {
	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))

	t, err := v.GetPropertyType(name)
	if err != nil {
		return nil, err
	}

	p, err := ValueInit(t)
	if err != nil {
		return nil, errors.New("unable to allocate value")
	}
	C.g_object_get_property(v.gobject, (*C.gchar)(cstr), p.native())
	return p.GoValue()
}

// SetProperty is a wrapper around g_object_set_property().
func (v *Object) SetProperty(name string, value interface{}) error {
	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))

	if obj, ok := value.(Object); ok {
		value = obj.gobject
	}

	p, err := GValue(value)
	if err != nil {
		return errors.New("Unable to perform type conversion")
	}
	C.g_object_set_property(C.toGObject(unsafe.Pointer(v.Native())), (*C.gchar)(cstr), p.native())
	return nil
}

// pointerVal attempts to return an unsafe.Pointer for value.
// Not all types are understood, in which case a nil Pointer
// is returned.
/*func pointerVal(value interface{}) unsafe.Pointer {
	var p unsafe.Pointer
	switch v := value.(type) {
	case bool:
		c := gbool(v)
		p = unsafe.Pointer(&c)

	case int8:
		c := C.gint8(v)
		p = unsafe.Pointer(&c)

	case int16:
		c := C.gint16(v)
		p = unsafe.Pointer(&c)

	case int32:
		c := C.gint32(v)
		p = unsafe.Pointer(&c)

	case int64:
		c := C.gint64(v)
		p = unsafe.Pointer(&c)

	case int:
		c := C.gint(v)
		p = unsafe.Pointer(&c)

	case uint8:
		c := C.guchar(v)
		p = unsafe.Pointer(&c)

	case uint16:
		c := C.guint16(v)
		p = unsafe.Pointer(&c)

	case uint32:
		c := C.guint32(v)
		p = unsafe.Pointer(&c)

	case uint64:
		c := C.guint64(v)
		p = unsafe.Pointer(&c)

	case uint:
		c := C.guint(v)
		p = unsafe.Pointer(&c)

	case uintptr:
		p = unsafe.Pointer(C.gpointer(v))

	case float32:
		c := C.gfloat(v)
		p = unsafe.Pointer(&c)

	case float64:
		c := C.gdouble(v)
		p = unsafe.Pointer(&c)

	case string:
		cstr := C.CString(v)
		defer C.free(unsafe.Pointer(cstr))
		p = unsafe.Pointer(cstr)

	default:
		if pv, ok := value.(unsafe.Pointer); ok {
			p = pv
		} else {
			val := reflect.ValueOf(value)
			switch val.Kind() {
			case reflect.Int, reflect.Int8, reflect.Int16,
				reflect.Int32, reflect.Int64:
				c := C.int(val.Int())
				p = unsafe.Pointer(&c)

			case reflect.Uintptr, reflect.Ptr, reflect.UnsafePointer:
				p = unsafe.Pointer(C.gpointer(val.Pointer()))
			}
		}
	}

	return p
}*/

/*
 * GObject Signals
 */

// Emit is a wrapper around g_signal_emitv() and emits the signal
// specified by the string s to an Object.  Arguments to callback
// functions connected to this signal must be specified in args.  Emit()
// returns an interface{} which must be type asserted as the Go
// equivalent type to the return value for native C callback.
//
// Note that this code is unsafe in that the types of values in args are
// not checked against whether they are suitable for the callback.
func (v *Object) Emit(s string, args ...interface{}) (interface{}, error) {
	cstr := C.CString(s)
	defer C.free(unsafe.Pointer(cstr))

	// Create array of this instance and arguments
	valv := C.alloc_gvalue_list(C.int(len(args)) + 1)
	defer C.free(unsafe.Pointer(valv))

	// Add args and valv
	val, err := GValue(v)
	if err != nil {
		return nil, errors.New("Error converting Object to GValue: " + err.Error())
	}
	C.val_list_insert(valv, C.int(0), val.native())
	for i := range args {
		val, err := GValue(args[i])
		if err != nil {
			return nil, fmt.Errorf("Error converting arg %d to GValue: %s", i, err.Error())
		}
		C.val_list_insert(valv, C.int(i+1), val.native())
	}

	t := v.TypeFromInstance()
	// TODO: use just the signal name
	id := C.g_signal_lookup((*C.gchar)(cstr), C.GType(t))

	ret, err := ValueAlloc()
	if err != nil {
		return nil, errors.New("Error creating Value for return value")
	}
	C.g_signal_emitv(valv, id, C.GQuark(0), ret.native())

	return ret.GoValue()
}

// HandlerBlock is a wrapper around g_signal_handler_block().
func (v *Object) HandlerBlock(handle SignalHandle) {
	C.g_signal_handler_block(C.gpointer(v.gobject), C.gulong(handle))
}

// HandlerUnblock is a wrapper around g_signal_handler_unblock().
func (v *Object) HandlerUnblock(handle SignalHandle) {
	C.g_signal_handler_unblock(C.gpointer(v.gobject), C.gulong(handle))
}

// HandlerDisconnect is a wrapper around g_signal_handler_disconnect().
func (v *Object) HandlerDisconnect(handle SignalHandle) {
	C.g_signal_handler_disconnect(C.gpointer(v.gobject), C.gulong(handle))
	C.g_closure_invalidate(signals[handle])
	delete(closures.m, signals[handle])
	delete(signals, handle)
}

// Wrapper function for new objects with reference management.
func wrapObject(ptr unsafe.Pointer) *Object {
	obj := &Object{C.toGObject(ptr)}

	if obj.IsFloating() {
		obj.RefSink()
	} else {
		obj.Ref()
	}

	runtime.SetFinalizer(obj, (*Object).Unref)
	return obj
}
