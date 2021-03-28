package gdk

// #cgo pkg-config: gdk-3.0
// #include <gdk/gdk.h>
// #include "gdk.go.h"
import "C"
import (
	"unsafe"

	"github.com/romychs/gotk3/glib"
)

/*
 * GdkScreen
 */

// Screen is a representation of GDK's GdkScreen.
type Screen struct {
	*glib.Object
}

// native returns a pointer to the underlying GdkScreen.
func (v *Screen) native() *C.GdkScreen {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toGdkScreen(ptr)
}

func marshalScreen(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.ToObject(unsafe.Pointer(c))
	return &Screen{obj}, nil
}

func toScreen(s *C.GdkScreen) (*Screen, error) {
	if s == nil {
		return nil, nilPtrErr
	}
	obj := glib.ToObject(unsafe.Pointer(s))
	return &Screen{obj}, nil
}

func (v *Screen) Native() uintptr {
	return uintptr(unsafe.Pointer(v.native()))
}

// GetRGBAVisual is a wrapper around gdk_screen_get_rgba_visual().
func (v *Screen) GetRGBAVisual() (*Visual, error) {
	c := C.gdk_screen_get_rgba_visual(v.native())
	if c == nil {
		return nil, nilPtrErr
	}

	return &Visual{glib.Take(unsafe.Pointer(c))}, nil
}

// GetSystemVisual is a wrapper around gdk_screen_get_system_visual().
func (v *Screen) GetSystemVisual() (*Visual, error) {
	c := C.gdk_screen_get_system_visual(v.native())
	if c == nil {
		return nil, nilPtrErr
	}

	return &Visual{glib.Take(unsafe.Pointer(c))}, nil
}

// ScreenGetDefault is a wrapper around gdk_screen_get_default().
func ScreenGetDefault() (*Screen, error) {
	return toScreen(C.gdk_screen_get_default())
}

// IsComposited is a wrapper around gdk_screen_is_composited().
func (v *Screen) IsComposited() bool {
	return gobool(C.gdk_screen_is_composited(v.native()))
}

// GetRootWindow is a wrapper around gdk_screen_get_root_window().
func (v *Screen) GetRootWindow() (*Window, error) {
	return toWindow(C.gdk_screen_get_root_window(v.native()))
}

// GetDisplay is a wrapper around gdk_screen_get_display().
func (v *Screen) GetDisplay() (*Display, error) {
	return toDisplay(C.gdk_screen_get_display(v.native()))
}

func toString(c *C.gchar) (string, error) {
	if c == nil {
		return "", nilPtrErr
	}
	return goString(c), nil
}

// GetResolution is a wrapper around gdk_screen_get_resolution().
func (v *Screen) GetResolution() float64 {
	return float64(C.gdk_screen_get_resolution(v.native()))
}

// SetResolution is a wrapper around gdk_screen_set_resolution().
func (v *Screen) SetResolution(r float64) {
	C.gdk_screen_set_resolution(v.native(), C.gdouble(r))
}

// void 	gdk_screen_set_font_options ()
// gboolean 	gdk_screen_get_setting ()
// const cairo_font_options_t * 	gdk_screen_get_font_options ()
// GList * 	gdk_screen_get_window_stack ()
// GList * 	gdk_screen_list_visuals ()
// GList * 	gdk_screen_get_toplevel_windows ()
// void 	gdk_screen_get_monitor_geometry ()
// void 	gdk_screen_get_monitor_workarea ()
