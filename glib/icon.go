package glib

// #cgo pkg-config: glib-2.0 gobject-2.0
// #include <gio/gio.h>
// #include <glib.h>
// #include <glib-object.h>
// #include "glib.go.h"
import "C"
import (
	"unsafe"
)

// Icon is a representation of GInterface.
type Icon struct {
	Interface
}

// native() returns a pointer to the underlying GThemedIcon.
func (v *Icon) native() *C.GIcon {
	return C.toGIcon(unsafe.Pointer(v.Native()))
}

func marshalIcon(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := Take(unsafe.Pointer(c))
	icon := wrapIcon(*InterfaceFromObjectNew(obj))
	return icon, nil
}

func wrapIcon(intf Interface) *Icon {
	return &Icon{intf}
}

/*
func IconForStringNew(path string) (*Icon, error) {
	cstr := C.CString(path)
	defer C.free(unsafe.Pointer(cstr))

	var err *C.GError = nil
	c := C.g_icon_new_for_string((*C.gchar)(cstr), &err)
	if err != nil {
		defer C.g_error_free(err)
		return nil, errors.New(goString(err.message))
	}

	obj := Take(unsafe.Pointer(c))
	icon := wrapIcon(InterfaceFromObjectNew(obj))
	return icon, nil
}
*/

// ThemedIcon is a representation of GThemedIcon.
type ThemedIcon struct {
	*Object
	// Interfaces
	Icon
}

// native() returns a pointer to the underlying GThemedIcon.
func (v *ThemedIcon) native() *C.GThemedIcon {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toGThemedIcon(ptr)
}

func (v *ThemedIcon) Native() uintptr {
	return uintptr(unsafe.Pointer(v.native()))
}

func marshalThemedIcon(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := Take(unsafe.Pointer(c))
	return wrapThemedIcon(obj), nil
}

func wrapThemedIcon(obj *Object) *ThemedIcon {
	icon := wrapIcon(*InterfaceFromObjectNew(obj))
	return &ThemedIcon{obj, *icon}
}

// GIcon *
// g_themed_icon_new (const char *iconname);
func ThemedIconNew(iconName string) (*ThemedIcon, error) {
	cstr := C.CString(iconName)
	defer C.free(unsafe.Pointer(cstr))

	c := C.g_themed_icon_new(cstr)
	if c == nil {
		return nil, errNilPtr
	}

	return wrapThemedIcon(wrapObject(unsafe.Pointer(c))), nil
}

// FileIcon is a representation of GFileIcon.
type FileIcon struct {
	*Object
	// Interfaces
	Icon
}

// native() returns a pointer to the underlying GFileIcon.
func (v *FileIcon) native() *C.GFileIcon {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toGFileIcon(ptr)
}

func (v *FileIcon) Native() uintptr {
	return uintptr(unsafe.Pointer(v.native()))
}

func marshalFileIcon(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := Take(unsafe.Pointer(c))
	return wrapFileIcon(obj), nil
}

func wrapFileIcon(obj *Object) *FileIcon {
	icon := wrapIcon(*InterfaceFromObjectNew(obj))
	return &FileIcon{obj, *icon}
}

func FileIconNew(file *File) (*FileIcon, error) {
	c := C.g_file_icon_new(file.native())
	if c == nil {
		return nil, errNilPtr
	}

	return wrapFileIcon(wrapObject(unsafe.Pointer(c))), nil
}
