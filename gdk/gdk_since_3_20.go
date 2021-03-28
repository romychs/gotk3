// +build !gtk_3_6,!gtk_3_8,!gtk_3_10,!gtk_3_12,!gtk_3_14,!gtk_3_16,!gtk_3_18
// not use this: go build -tags gtk_3_8'. Otherwise, if no build tags are used, GDK 3.20

// Go bindings for GDK 3.  Supports version 3.6 and later.
package gdk

// #cgo pkg-config: gdk-3.0
// #include <gdk/gdk.h>
// #include "gdk.go.h"
// #include "gdk_since_3_20.go.h"
import "C"
import (
	"runtime"
	"unsafe"

	"github.com/romychs/gotk3/glib"
)

func init() {
	tm := []glib.TypeMarshaler{
		// Enums
		{glib.Type(C.gdk_seat_capabilities_get_type()), marshalSeatCapabilities},
		// {glib.Type(C.gdk_event_type_get_type()), marshalEventType},
		// {glib.Type(C.gdk_interp_type_get_type()), marshalInterpType},
		// {glib.Type(C.gdk_modifier_type_get_type()), marshalModifierType},
		// {glib.Type(C.gdk_pixbuf_alpha_mode_get_type()), marshalPixbufAlphaMode},
		// {glib.Type(C.gdk_event_mask_get_type()), marshalEventMask},
		// {glib.Type(C.gdk_rectangle_get_type()), marshalRectangle},

		// Objects/Interfaces
		{glib.Type(C.gdk_seat_get_type()), marshalSeat},
	}
	glib.RegisterGValueMarshalers(tm)
}

/*
 * Constants
 */

// SeatCapabilities is a representation of GDK's GdkSeatCapabilities.
type SeatCapabilities int

const (
	SEAT_CAPABILITY_NONE          SeatCapabilities = C.GDK_SEAT_CAPABILITY_NONE
	SEAT_CAPABILITY_POINTER       SeatCapabilities = C.GDK_SEAT_CAPABILITY_POINTER
	SEAT_CAPABILITY_TOUCH         SeatCapabilities = C.GDK_SEAT_CAPABILITY_TOUCH
	SEAT_CAPABILITY_TABLET_STYLUS SeatCapabilities = C.GDK_SEAT_CAPABILITY_TABLET_STYLUS
	SEAT_CAPABILITY_KEYBOARD      SeatCapabilities = C.GDK_SEAT_CAPABILITY_KEYBOARD
	SEAT_CAPABILITY_ALL_POINTING  SeatCapabilities = C.GDK_SEAT_CAPABILITY_ALL_POINTING
	SEAT_CAPABILITY_ALL           SeatCapabilities = C.GDK_SEAT_CAPABILITY_ALL
)

func marshalSeatCapabilities(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return SeatCapabilities(c), nil
}

// native returns the underlying GdkAtom.
func (v SeatCapabilities) native() C.GdkSeatCapabilities {
	return C.GdkSeatCapabilities(v)
}

/*
 * GdkSeat
 */

// Seat is a representation of GDK's GdkSeat.
type Seat struct {
	*glib.Object
}

// native returns a pointer to the underlying GdkSeat.
func (v *Seat) native() *C.GdkSeat {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toGdkSeat(ptr)
}

func marshalSeat(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.ToObject(unsafe.Pointer(c))
	return &Seat{obj}, nil
}

func wrapSeat(obj *glib.Object) *Seat {
	return &Seat{obj}
}

// GetDisplay is a wrapper around gdk_seat_get_display().
func (v *Seat) GetDisplay() (*Display, error) {
	c := C.gdk_seat_get_display(v.native())
	if c == nil {
		return nil, nilPtrErr
	}

	return &Display{glib.Take(unsafe.Pointer(c))}, nil
}

// GetCapabilities is a wrapper around gdk_seat_get_capabilities().
func (v *Seat) GetCapabilities() SeatCapabilities {
	c := C.gdk_seat_get_capabilities(v.native())
	return SeatCapabilities(c)
}

// GetPointer is a wrapper around gdk_seat_get_pointer().
func (v *Seat) GetPointer() (*Device, error) {
	c := C.gdk_seat_get_pointer(v.native())
	if c == nil {
		return nil, nilPtrErr
	}

	return &Device{glib.Take(unsafe.Pointer(c))}, nil
}

// GetKeyboard is a wrapper around gdk_seat_get_keyboard().
func (v *Seat) GetKeyboard() (*Device, error) {
	c := C.gdk_seat_get_keyboard(v.native())
	if c == nil {
		return nil, nilPtrErr
	}

	return &Device{glib.Take(unsafe.Pointer(c))}, nil
}

// GetSlaves is a wrapper around gdk_seat_get_slaves().
// Returned list is wrapped to return *gdk.Device elements.
func (v *Seat) GetSlaves(capabilities SeatCapabilities) *glib.List {
	clist := C.gdk_seat_get_slaves(v.native(), capabilities.native())

	glist := glib.WrapList(uintptr(unsafe.Pointer(clist)))
	glist.DataWrapper(func(ptr unsafe.Pointer) interface{} {
		d := wrapDevice(glib.Take(ptr))
		return d
	})

	if glist != nil {
		runtime.SetFinalizer(glist, func(glist *glib.List) {
			glist.Free()
		})
	}

	return glist
}

/*
 * GdkDisplay
 */

// GetDefaultSeat is a wrapper around gdk_display_get_default_seat().
func (v *Display) GetDefaultSeat() (*Seat, error) {
	c := C.gdk_display_get_default_seat(v.native())
	if c == nil {
		return nil, nilPtrErr
	}

	return &Seat{glib.Take(unsafe.Pointer(c))}, nil
}

// ListSeats is a wrapper around gdk_display_list_seats().
// Returned list is wrapped to return *gdk.Seat elements.
func (v *Display) ListSeats() *glib.List {
	clist := C.gdk_display_list_seats(v.native())

	glist := glib.WrapList(uintptr(unsafe.Pointer(clist)))
	glist.DataWrapper(func(ptr unsafe.Pointer) interface{} {
		d := wrapSeat(glib.Take(ptr))
		return d
	})

	if glist != nil {
		runtime.SetFinalizer(glist, func(glist *glib.List) {
			glist.Free()
		})
	}

	return glist
}
