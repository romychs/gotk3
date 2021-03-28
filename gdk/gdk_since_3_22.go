// +build !gtk_3_6,!gtk_3_8,!gtk_3_10,!gtk_3_12,!gtk_3_14,!gtk_3_16,!gtk_3_18,!gtk_3_20
// not use this: go build -tags gtk_3_8'. Otherwise, if no build tags are used, GDK 3.22

// Go bindings for GDK 3.  Supports version 3.6 and later.
package gdk

// #cgo pkg-config: gdk-3.0
// #include <gdk/gdk.h>
// #include "gdk.go.h"
// #include "gdk_since_3_22.go.h"
import "C"
import (
	"unsafe"

	"github.com/romychs/gotk3/glib"
)

func init() {
	tm := []glib.TypeMarshaler{
		// Enums
		{glib.Type(C.gdk_subpixel_layout_get_type()), marshalSubpixelLayout},
		// {glib.Type(C.gdk_colorspace_get_type()), marshalColorspace},
		// {glib.Type(C.gdk_event_type_get_type()), marshalEventType},
		// {glib.Type(C.gdk_interp_type_get_type()), marshalInterpType},
		// {glib.Type(C.gdk_modifier_type_get_type()), marshalModifierType},
		// {glib.Type(C.gdk_pixbuf_alpha_mode_get_type()), marshalPixbufAlphaMode},
		// {glib.Type(C.gdk_event_mask_get_type()), marshalEventMask},
		// {glib.Type(C.gdk_rectangle_get_type()), marshalRectangle},

		// Objects/Interfaces
		{glib.Type(C.gdk_monitor_get_type()), marshalMonitor},
	}
	glib.RegisterGValueMarshalers(tm)
}

/*
 * Constants
 */

// SubpixelLayout is a representation of GDK's GdkSubpixelLayout.
type SubpixelLayout int

const (
	SUBPIXEL_LAYOUT_UNKNOWN        SubpixelLayout = C.GDK_SUBPIXEL_LAYOUT_UNKNOWN
	SUBPIXEL_LAYOUT_NONE           SubpixelLayout = C.GDK_SUBPIXEL_LAYOUT_NONE
	SUBPIXEL_LAYOUT_HORIZONTAL_RGB SubpixelLayout = C.GDK_SUBPIXEL_LAYOUT_HORIZONTAL_RGB
	SUBPIXEL_LAYOUT_HORIZONTAL_BGR SubpixelLayout = C.GDK_SUBPIXEL_LAYOUT_HORIZONTAL_BGR
	SUBPIXEL_LAYOUT_VERTICAL_RGB   SubpixelLayout = C.GDK_SUBPIXEL_LAYOUT_VERTICAL_RGB
	SUBPIXEL_LAYOUT_VERTICAL_BGR   SubpixelLayout = C.GDK_SUBPIXEL_LAYOUT_VERTICAL_BGR
)

func marshalSubpixelLayout(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return SubpixelLayout(c), nil
}

/*
 * GdkMonitor
 */

// Monitor is a representation of GDK's GdkMonitor.
type Monitor struct {
	*glib.Object
}

// native returns a pointer to the underlying GdkMonitor.
func (v *Monitor) native() *C.GdkMonitor {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toGdkMonitor(ptr)
}

func marshalMonitor(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.ToObject(unsafe.Pointer(c))
	return &Monitor{obj}, nil
}

// GetDisplay is a wrapper around gdk_monitor_get_display().
func (v *Monitor) GetDisplay() (*Display, error) {
	c := C.gdk_monitor_get_display(v.native())
	if c == nil {
		return nil, nilPtrErr
	}

	return &Display{glib.Take(unsafe.Pointer(c))}, nil
}

// GetGeometry is a wrapper around gdk_monitor_get_geometry().
func (v *Monitor) GetGeometry() *Rectangle {
	var r C.GdkRectangle
	C.gdk_monitor_get_geometry(v.native(), &r)

	return wrapRectangle(&r)
}

// GetWorkarea is a wrapper around gdk_monitor_get_width_mm().
func (v *Monitor) GetWorkarea() *Rectangle {
	var r C.GdkRectangle
	C.gdk_monitor_get_workarea(v.native(), &r)

	return wrapRectangle(&r)
}

// GetWidthMm is a wrapper around gdk_monitor_get_width_mm().
func (v *Monitor) GetWidthMm() int {
	c := C.gdk_monitor_get_width_mm(v.native())
	return int(c)
}

// GetHeightMm is a wrapper around gdk_monitor_get_height_mm().
func (v *Monitor) GetHeightMm() int {
	c := C.gdk_monitor_get_height_mm(v.native())
	return int(c)
}

// GetManufacturer is a wrapper around gdk_monitor_get_manufacturer().
func (v *Monitor) GetManufacturer() (string, error) {
	c := C.gdk_monitor_get_manufacturer(v.native())
	if c == nil {
		return "", nilPtrErr
	}
	return goString(c), nil
}

// GetModel is a wrapper around gdk_monitor_get_model().
func (v *Monitor) GetModel() (string, error) {
	c := C.gdk_monitor_get_model(v.native())
	if c == nil {
		return "", nilPtrErr
	}
	return goString(c), nil
}

// GetScaleFactor is a wrapper around gdk_monitor_get_scale_factor().
func (v *Monitor) GetScaleFactor() int {
	c := C.gdk_monitor_get_scale_factor(v.native())
	return int(c)
}

// GetRefreshRate is a wrapper around gdk_monitor_get_refresh_rate().
func (v *Monitor) GetRefreshRate() int {
	c := C.gdk_monitor_get_refresh_rate(v.native())
	return int(c)
}

// GetSubpixelLayout is a wrapper around gdk_monitor_get_subpixel_layout().
func (v *Monitor) GetSubpixelLayout() SubpixelLayout {
	c := C.gdk_monitor_get_subpixel_layout(v.native())
	return SubpixelLayout(c)
}

// GetRefreshRate is a wrapper around gdk_monitor_get_refresh_rate().
func (v *Monitor) IsPrimary() bool {
	c := C.gdk_monitor_is_primary(v.native())
	return gobool(c)
}

/*
 * GdkDisplay
 */

// GetMonitorsNumber is a wrapper around gdk_display_get_n_monitors().
func (v *Display) GetMonitorsNumber() int {
	c := C.gdk_display_get_n_monitors(v.native())
	return int(c)
}

// GetMonitor is a wrapper around gdk_display_get_primary_monitor().
func (v *Display) GetMonitor(n int) (*Monitor, error) {
	c := C.gdk_display_get_monitor(v.native(), C.int(n))
	if c == nil {
		return nil, nilPtrErr
	}

	return &Monitor{glib.Take(unsafe.Pointer(c))}, nil
}

// GetPrimaryMonitor is a wrapper around gdk_display_get_primary_monitor().
func (v *Display) GetPrimaryMonitor() (*Monitor, error) {
	c := C.gdk_display_get_primary_monitor(v.native())
	if c == nil {
		return nil, nilPtrErr
	}

	return &Monitor{glib.Take(unsafe.Pointer(c))}, nil
}

// GetMonitorAtPoint is a wrapper around gdk_display_get_monitor_at_point().
func (v *Display) GetMonitorAtPoint(x, y int) (*Monitor, error) {
	c := C.gdk_display_get_monitor_at_point(v.native(), C.int(x), C.int(y))
	if c == nil {
		return nil, nilPtrErr
	}

	return &Monitor{glib.Take(unsafe.Pointer(c))}, nil
}

// GetMonitorAtWindow is a wrapper around gdk_display_get_monitor_at_window().
func (v *Display) GetMonitorAtWindow(window *Window) (*Monitor, error) {
	c := C.gdk_display_get_monitor_at_window(v.native(), window.native())
	if c == nil {
		return nil, nilPtrErr
	}

	return &Monitor{glib.Take(unsafe.Pointer(c))}, nil
}
