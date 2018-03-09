package glib

// #cgo pkg-config: glib-2.0 gobject-2.0 gio-2.0
// #include <gio/gio.h>
// #include <glib.h>
// #include <glib-object.h>
// #include "glib.go.h"
import "C"

type Source struct {
	gsource *C.GSource
}

// native returns a pointer to the underlying GSource.
func (v *Source) native() *C.GSource {
	if v == nil {
		return nil
	}
	return v.gsource
}

// MainCurrentSource is a wrapper around g_main_current_source().
func MainCurrentSource() *Source {
	c := C.g_main_current_source()
	if c == nil {
		return nil
	}
	return &Source{c}
}
