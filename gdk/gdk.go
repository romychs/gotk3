// Copyright (c) 2013-2014 Conformal Systems <info@conformal.com>
//
// This file originated from: http://opensource.conformal.com/
//
// Permission to use, copy, modify, and distribute this software for any
// purpose with or without fee is hereby granted, provided that the above
// copyright notice and this permission notice appear in all copies.
//
// THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
// WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
// MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
// ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
// WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
// ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
// OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.

// Go bindings for GDK 3.  Supports version 3.6 and later.
package gdk

// #cgo pkg-config: gdk-3.0 gio-2.0
// #include <gdk/gdk.h>
// #include "gdk.go.h"
// #include "gdk_atom.go.h"
import "C"

import (
	"errors"
	"runtime"
	"unsafe"

	"github.com/d2r2/gotk3/glib"
)

func init() {
	tm := []glib.TypeMarshaler{
		// Enums
		{glib.Type(C.gdk_drag_action_get_type()), marshalDragAction},
		{glib.Type(C.gdk_colorspace_get_type()), marshalColorspace},
		{glib.Type(C.gdk_interp_type_get_type()), marshalInterpType},
		{glib.Type(C.gdk_pixbuf_rotation_get_type()), marshalPixbufRotation},
		{glib.Type(C.gdk_modifier_type_get_type()), marshalModifierType},
		{glib.Type(C.gdk_window_state_get_type()), marshalWindowState},
		{glib.Type(C.gdk_window_hints_get_type()), marshalWindowHints},
		{glib.Type(C.gdk_window_type_hint_get_type()), marshalWindowTypeHint},
		{glib.Type(C.gdk_grab_status_get_type()), marshalGrabStatus},
		{glib.Type(C.gdk_grab_ownership_get_type()), marshalGrabOwnership},
		{glib.Type(C.gdk_device_type_get_type()), marshalDeviceType},
		{glib.Type(C.gdk_gravity_get_type()), marshalGravity},
		{glib.Type(C.gdk_event_type_get_type()), marshalEventType},
		{glib.Type(C.gdk_event_mask_get_type()), marshalEventMask},
		{glib.Type(C.gdk_pixbuf_alpha_mode_get_type()), marshalPixbufAlphaMode},

		// Objects/Interfaces
		{glib.Type(C.gdk_device_get_type()), marshalDevice},
		{glib.Type(C.gdk_cursor_get_type()), marshalCursor},
		{glib.Type(C.gdk_device_manager_get_type()), marshalDeviceManager},
		{glib.Type(C.gdk_display_get_type()), marshalDisplay},
		{glib.Type(C.gdk_drag_context_get_type()), marshalDragContext},
		{glib.Type(C.gdk_pixbuf_get_type()), marshalPixbuf},
		{glib.Type(C.gdk_screen_get_type()), marshalScreen},
		{glib.Type(C.gdk_visual_get_type()), marshalVisual},
		{glib.Type(C.gdk_window_get_type()), marshalWindow},

		// Boxed
		{glib.Type(C.gdk_event_get_type()), marshalEvent},
		{glib.Type(C.gdk_rectangle_get_type()), marshalRectangle},
		{glib.Type(C.gdk_rgba_get_type()), marshalRGBA},
	}
	glib.RegisterGValueMarshalers(tm)
}

/*
 * Type conversions
 */

func gbool(b bool) C.gboolean {
	if b {
		return C.gboolean(1)
	}
	return C.gboolean(0)
}
func gobool(b C.gboolean) bool {
	if b != 0 {
		return true
	}
	return false
}

func goString(cstr *C.gchar) string {
	return C.GoString((*C.char)(cstr))
}

/*
 * Unexported vars
 */

var nilPtrErr = errors.New("cgo returned unexpected nil pointer")

/*
 * Constants
 */

// DragAction is a representation of GDK's GdkDragAction.
type DragAction int

const (
	ACTION_DEFAULT DragAction = C.GDK_ACTION_DEFAULT
	ACTION_COPY    DragAction = C.GDK_ACTION_COPY
	ACTION_MOVE    DragAction = C.GDK_ACTION_MOVE
	ACTION_LINK    DragAction = C.GDK_ACTION_LINK
	ACTION_PRIVATE DragAction = C.GDK_ACTION_PRIVATE
	ACTION_ASK     DragAction = C.GDK_ACTION_ASK
)

func marshalDragAction(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return DragAction(c), nil
}

// Colorspace is a representation of GDK's GdkColorspace.
type Colorspace int

const (
	COLORSPACE_RGB Colorspace = C.GDK_COLORSPACE_RGB
)

func marshalColorspace(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return Colorspace(c), nil
}

// InterpType is a representation of GDK's GdkInterpType.
type InterpType int

const (
	INTERP_NEAREST  InterpType = C.GDK_INTERP_NEAREST
	INTERP_TILES    InterpType = C.GDK_INTERP_TILES
	INTERP_BILINEAR InterpType = C.GDK_INTERP_BILINEAR
	INTERP_HYPER    InterpType = C.GDK_INTERP_HYPER
)

func marshalInterpType(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return InterpType(c), nil
}

// PixbufRotation is a representation of GDK's GdkPixbufRotation.
type PixbufRotation int

const (
	PIXBUF_ROTATE_NONE             PixbufRotation = C.GDK_PIXBUF_ROTATE_NONE
	PIXBUF_ROTATE_COUNTERCLOCKWISE PixbufRotation = C.GDK_PIXBUF_ROTATE_COUNTERCLOCKWISE
	PIXBUF_ROTATE_UPSIDEDOWN       PixbufRotation = C.GDK_PIXBUF_ROTATE_UPSIDEDOWN
	PIXBUF_ROTATE_CLOCKWISE        PixbufRotation = C.GDK_PIXBUF_ROTATE_CLOCKWISE
)

func marshalPixbufRotation(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return PixbufRotation(c), nil
}

// ModifierType is a representation of GDK's GdkModifierType.
type ModifierType uint

const (
	GDK_SHIFT_MASK    ModifierType = C.GDK_SHIFT_MASK
	GDK_LOCK_MASK     ModifierType = C.GDK_LOCK_MASK
	GDK_CONTROL_MASK  ModifierType = C.GDK_CONTROL_MASK
	GDK_MOD1_MASK     ModifierType = C.GDK_MOD1_MASK
	GDK_MOD2_MASK     ModifierType = C.GDK_MOD2_MASK
	GDK_MOD3_MASK     ModifierType = C.GDK_MOD3_MASK
	GDK_MOD4_MASK     ModifierType = C.GDK_MOD4_MASK
	GDK_MOD5_MASK     ModifierType = C.GDK_MOD5_MASK
	GDK_BUTTON1_MASK  ModifierType = C.GDK_BUTTON1_MASK
	GDK_BUTTON2_MASK  ModifierType = C.GDK_BUTTON2_MASK
	GDK_BUTTON3_MASK  ModifierType = C.GDK_BUTTON3_MASK
	GDK_BUTTON4_MASK  ModifierType = C.GDK_BUTTON4_MASK
	GDK_BUTTON5_MASK  ModifierType = C.GDK_BUTTON5_MASK
	GDK_SUPER_MASK    ModifierType = C.GDK_SUPER_MASK
	GDK_HYPER_MASK    ModifierType = C.GDK_HYPER_MASK
	GDK_META_MASK     ModifierType = C.GDK_META_MASK
	GDK_RELEASE_MASK  ModifierType = C.GDK_RELEASE_MASK
	GDK_MODIFIER_MASK ModifierType = C.GDK_MODIFIER_MASK
)

func marshalModifierType(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return ModifierType(c), nil
}

// Selections

var (
	SELECTION_PRIMARY       = Atom{C._GDK_SELECTION_PRIMARY}
	SELECTION_SECONDARY     = Atom{C._GDK_SELECTION_SECONDARY}
	SELECTION_CLIPBOARD     = Atom{C._GDK_SELECTION_CLIPBOARD}
	TARGET_BITMAP           = Atom{C._GDK_TARGET_BITMAP}
	TARGET_COLORMAP         = Atom{C._GDK_TARGET_COLORMAP}
	TARGET_DRAWABLE         = Atom{C._GDK_TARGET_DRAWABLE}
	TARGET_PIXMAP           = Atom{C._GDK_TARGET_PIXMAP}
	TARGET_STRING           = Atom{C._GDK_TARGET_STRING}
	SELECTION_TYPE_ATOM     = Atom{C._GDK_SELECTION_TYPE_ATOM}
	SELECTION_TYPE_BITMAP   = Atom{C._GDK_SELECTION_TYPE_BITMAP}
	SELECTION_TYPE_COLORMAP = Atom{C._GDK_SELECTION_TYPE_COLORMAP}
	SELECTION_TYPE_DRAWABLE = Atom{C._GDK_SELECTION_TYPE_DRAWABLE}
	SELECTION_TYPE_INTEGER  = Atom{C._GDK_SELECTION_TYPE_INTEGER}
	SELECTION_TYPE_PIXMAP   = Atom{C._GDK_SELECTION_TYPE_PIXMAP}
	SELECTION_TYPE_WINDOW   = Atom{C._GDK_SELECTION_TYPE_WINDOW}
	SELECTION_TYPE_STRING   = Atom{C._GDK_SELECTION_TYPE_STRING}
)

// WindowState is a representation of GDK's GdkWindowState
type WindowState int

const (
	WINDOW_STATE_WITHDRAWN  WindowState = C.GDK_WINDOW_STATE_WITHDRAWN
	WINDOW_STATE_ICONIFIED  WindowState = C.GDK_WINDOW_STATE_ICONIFIED
	WINDOW_STATE_MAXIMIZED  WindowState = C.GDK_WINDOW_STATE_MAXIMIZED
	WINDOW_STATE_STICKY     WindowState = C.GDK_WINDOW_STATE_STICKY
	WINDOW_STATE_FULLSCREEN WindowState = C.GDK_WINDOW_STATE_FULLSCREEN
	WINDOW_STATE_ABOVE      WindowState = C.GDK_WINDOW_STATE_ABOVE
	WINDOW_STATE_BELOW      WindowState = C.GDK_WINDOW_STATE_BELOW
	WINDOW_STATE_FOCUSED    WindowState = C.GDK_WINDOW_STATE_FOCUSED
	WINDOW_STATE_TILED      WindowState = C.GDK_WINDOW_STATE_TILED
)

func marshalWindowState(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return WindowState(c), nil
}

// WindowHints is a representation of GDK's GdkWindowHints
type WindowHints int

const (
	HINT_POS         WindowHints = C.GDK_HINT_POS
	HINT_MIN_SIZE    WindowHints = C.GDK_HINT_MIN_SIZE
	HINT_MAX_SIZE    WindowHints = C.GDK_HINT_MAX_SIZE
	HINT_BASE_SIZE   WindowHints = C.GDK_HINT_BASE_SIZE
	HINT_ASPECT      WindowHints = C.GDK_HINT_ASPECT
	HINT_RESIZE_INC  WindowHints = C.GDK_HINT_RESIZE_INC
	HINT_WIN_GRAVITY WindowHints = C.GDK_HINT_WIN_GRAVITY
	HINT_USER_POS    WindowHints = C.GDK_HINT_USER_POS
	HINT_USER_SIZE   WindowHints = C.GDK_HINT_USER_SIZE
)

func marshalWindowHints(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return WindowHints(c), nil
}

// WindowTypeHint is a representation of GDK's GdkWindowTypeHint
type WindowTypeHint int

const (
	WINDOW_TYPE_HINT_NORMAL        WindowTypeHint = C.GDK_WINDOW_TYPE_HINT_NORMAL
	WINDOW_TYPE_HINT_DIALOG        WindowTypeHint = C.GDK_WINDOW_TYPE_HINT_DIALOG
	WINDOW_TYPE_HINT_MENU          WindowTypeHint = C.GDK_WINDOW_TYPE_HINT_MENU
	WINDOW_TYPE_HINT_TOOLBAR       WindowTypeHint = C.GDK_WINDOW_TYPE_HINT_TOOLBAR
	WINDOW_TYPE_HINT_SPLASHSCREEN  WindowTypeHint = C.GDK_WINDOW_TYPE_HINT_SPLASHSCREEN
	WINDOW_TYPE_HINT_UTILITY       WindowTypeHint = C.GDK_WINDOW_TYPE_HINT_UTILITY
	WINDOW_TYPE_HINT_DOCK          WindowTypeHint = C.GDK_WINDOW_TYPE_HINT_DOCK
	WINDOW_TYPE_HINT_DESKTOP       WindowTypeHint = C.GDK_WINDOW_TYPE_HINT_DESKTOP
	WINDOW_TYPE_HINT_DROPDOWN_MENU WindowTypeHint = C.GDK_WINDOW_TYPE_HINT_DROPDOWN_MENU
	WINDOW_TYPE_HINT_POPUP_MENU    WindowTypeHint = C.GDK_WINDOW_TYPE_HINT_POPUP_MENU
	WINDOW_TYPE_HINT_TOOLTIP       WindowTypeHint = C.GDK_WINDOW_TYPE_HINT_TOOLTIP
	WINDOW_TYPE_HINT_NOTIFICATION  WindowTypeHint = C.GDK_WINDOW_TYPE_HINT_NOTIFICATION
	WINDOW_TYPE_HINT_COMBO         WindowTypeHint = C.GDK_WINDOW_TYPE_HINT_COMBO
	WINDOW_TYPE_HINT_DND           WindowTypeHint = C.GDK_WINDOW_TYPE_HINT_DND
)

func marshalWindowTypeHint(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return WindowTypeHint(c), nil
}

// CURRENT_TIME is a representation of GDK_CURRENT_TIME

const CURRENT_TIME = C.GDK_CURRENT_TIME

// GrabStatus is a representation of GdkGrabStatus

type GrabStatus int

const (
	GRAB_SUCCESS         GrabStatus = C.GDK_GRAB_SUCCESS
	GRAB_ALREADY_GRABBED GrabStatus = C.GDK_GRAB_ALREADY_GRABBED
	GRAB_INVALID_TIME    GrabStatus = C.GDK_GRAB_INVALID_TIME
	GRAB_FROZEN          GrabStatus = C.GDK_GRAB_FROZEN
)

func marshalGrabStatus(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return GrabStatus(c), nil
}

// GrabOwnership is a representation of GdkGrabOwnership

type GrabOwnership int

const (
	OWNERSHIP_NONE        GrabOwnership = C.GDK_OWNERSHIP_NONE
	OWNERSHIP_WINDOW      GrabOwnership = C.GDK_OWNERSHIP_WINDOW
	OWNERSHIP_APPLICATION GrabOwnership = C.GDK_OWNERSHIP_APPLICATION
)

func marshalGrabOwnership(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return GrabOwnership(c), nil
}

// DeviceType is a representation of GdkDeviceType

type DeviceType int

const (
	DEVICE_TYPE_MASTER   DeviceType = C.GDK_DEVICE_TYPE_MASTER
	DEVICE_TYPE_SLAVE    DeviceType = C.GDK_DEVICE_TYPE_SLAVE
	DEVICE_TYPE_FLOATING DeviceType = C.GDK_DEVICE_TYPE_FLOATING
)

func marshalDeviceType(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return DeviceType(c), nil
}

// Gravity is a representation of GdkGravity

type Gravity int

const (
	GRAVITY_NORTH_WEST Gravity = C.GDK_GRAVITY_NORTH_WEST
	GRAVITY_NORTH      Gravity = C.GDK_GRAVITY_NORTH
	GRAVITY_NORTH_EAST Gravity = C.GDK_GRAVITY_NORTH_EAST
	GRAVITY_WEST       Gravity = C.GDK_GRAVITY_WEST
	GRAVITY_CENTER     Gravity = C.GDK_GRAVITY_CENTER
	GRAVITY_EAST       Gravity = C.GDK_GRAVITY_EAST
	GRAVITY_SOUTH_WEST Gravity = C.GDK_GRAVITY_SOUTH_WEST
	GRAVITY_SOUTH      Gravity = C.GDK_GRAVITY_SOUTH
	GRAVITY_SOUTH_EAST Gravity = C.GDK_GRAVITY_SOUTH_EAST
	GRAVITY_STATIC     Gravity = C.GDK_GRAVITY_STATIC
)

func marshalGravity(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return Gravity(c), nil
}

// EventPropagation constants

const (
	GDK_EVENT_PROPAGATE bool = C.GDK_EVENT_PROPAGATE != 0
	GDK_EVENT_STOP      bool = C.GDK_EVENT_STOP != 0
)

/*
 * GdkAtom
 */

// Atom is a representation of GDK's GdkAtom.
type Atom struct {
	gdkAtom C.GdkAtom
}

// native returns the underlying GdkAtom.
func (v Atom) native() C.GdkAtom {
	return v.gdkAtom
}

func (v Atom) Native() uintptr {
	return uintptr(unsafe.Pointer(v.native()))
}

func (v Atom) Name() string {
	c := C.gdk_atom_name(v.native())
	defer C.g_free(C.gpointer(c))
	return goString(c)
}

// AtomIntern is a wrapper around gdk_atom_intern
func AtomIntern(atomName string, onlyIfExists bool) *Atom {
	cstr := C.CString(atomName)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gdk_atom_intern((*C.gchar)(cstr), gbool(onlyIfExists))
	return &Atom{c}
}

/*
 * GdkDevice
 */

// Device is a representation of GDK's GdkDevice.
type Device struct {
	*glib.Object
}

// native returns a pointer to the underlying GdkDevice.
func (v *Device) native() *C.GdkDevice {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toGdkDevice(ptr)
}

func marshalDevice(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.ToObject(unsafe.Pointer(c))
	return &Device{obj}, nil
}

func wrapDevice(obj *glib.Object) *Device {
	return &Device{obj}
}

/*
 * GdkCursor
 */

// Cursor is a representation of GdkCursor.
type Cursor struct {
	*glib.Object
}

// CursorNewFromName is a wrapper around gdk_cursor_new_from_name().
func CursorNewFromName(display *Display, name string) (*Cursor, error) {
	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gdk_cursor_new_from_name(display.native(), (*C.gchar)(cstr))
	if c == nil {
		return nil, nilPtrErr
	}

	return &Cursor{glib.Take(unsafe.Pointer(c))}, nil
}

// native returns a pointer to the underlying GdkCursor.
func (v *Cursor) native() *C.GdkCursor {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toGdkCursor(ptr)
}

func marshalCursor(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.ToObject(unsafe.Pointer(c))
	return &Cursor{obj}, nil
}

/*
 * GdkDeviceManager
 */

// DeviceManager is a representation of GDK's GdkDeviceManager.
type DeviceManager struct {
	*glib.Object
}

// native returns a pointer to the underlying GdkDeviceManager.
func (v *DeviceManager) native() *C.GdkDeviceManager {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toGdkDeviceManager(ptr)
}

func marshalDeviceManager(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.ToObject(unsafe.Pointer(c))
	return &DeviceManager{obj}, nil
}

// GetDisplay is a wrapper around gdk_device_manager_get_display().
func (v *DeviceManager) GetDisplay() (*Display, error) {
	c := C.gdk_device_manager_get_display(v.native())
	if c == nil {
		return nil, nilPtrErr
	}

	return &Display{glib.Take(unsafe.Pointer(c))}, nil
}

/*
 * GdkDisplay
 */

// Display is a representation of GDK's GdkDisplay.
type Display struct {
	*glib.Object
}

// native returns a pointer to the underlying GdkDisplay.
func (v *Display) native() *C.GdkDisplay {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toGdkDisplay(ptr)
}

func marshalDisplay(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.ToObject(unsafe.Pointer(c))
	return &Display{obj}, nil
}

func toDisplay(s *C.GdkDisplay) (*Display, error) {
	if s == nil {
		return nil, nilPtrErr
	}
	obj := glib.ToObject(unsafe.Pointer(s))
	return &Display{obj}, nil
}

// DisplayOpen is a wrapper around gdk_display_open().
func DisplayOpen(displayName string) (*Display, error) {
	cstr := C.CString(displayName)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gdk_display_open((*C.gchar)(cstr))
	if c == nil {
		return nil, nilPtrErr
	}

	return &Display{glib.Take(unsafe.Pointer(c))}, nil
}

// DisplayGetDefault is a wrapper around gdk_display_get_default().
func DisplayGetDefault() (*Display, error) {
	c := C.gdk_display_get_default()
	if c == nil {
		return nil, nilPtrErr
	}

	return &Display{glib.Take(unsafe.Pointer(c))}, nil
}

// GetName is a wrapper around gdk_display_get_name().
func (v *Display) GetName() (string, error) {
	c := C.gdk_display_get_name(v.native())
	if c == nil {
		return "", nilPtrErr
	}
	return goString(c), nil
}

// GetDefaultScreen is a wrapper around gdk_display_get_default_screen().
func (v *Display) GetDefaultScreen() (*Screen, error) {
	c := C.gdk_display_get_default_screen(v.native())
	if c == nil {
		return nil, nilPtrErr
	}

	return &Screen{glib.Take(unsafe.Pointer(c))}, nil
}

// DeviceIsGrabbed is a wrapper around gdk_display_device_is_grabbed().
func (v *Display) DeviceIsGrabbed(device *Device) bool {
	c := C.gdk_display_device_is_grabbed(v.native(), device.native())
	return gobool(c)
}

// Beep is a wrapper around gdk_display_beep().
func (v *Display) Beep() {
	C.gdk_display_beep(v.native())
}

// Sync is a wrapper around gdk_display_sync().
func (v *Display) Sync() {
	C.gdk_display_sync(v.native())
}

// Flush is a wrapper around gdk_display_flush().
func (v *Display) Flush() {
	C.gdk_display_flush(v.native())
}

// Close is a wrapper around gdk_display_close().
func (v *Display) Close() {
	C.gdk_display_close(v.native())
}

// IsClosed is a wrapper around gdk_display_is_closed().
func (v *Display) IsClosed() bool {
	c := C.gdk_display_is_closed(v.native())
	return gobool(c)
}

// GetEvent is a wrapper around gdk_display_get_event().
func (v *Display) GetEvent() (*Event, error) {
	c := C.gdk_display_get_event(v.native())
	if c == nil {
		return nil, nilPtrErr
	}

	//The finalizer is not on the glib.Object but on the event.
	e := &Event{c}
	runtime.SetFinalizer(e, (*Event).free)
	return e, nil
}

// PeekEvent is a wrapper around gdk_display_peek_event().
func (v *Display) PeekEvent() (*Event, error) {
	c := C.gdk_display_peek_event(v.native())
	if c == nil {
		return nil, nilPtrErr
	}

	// The finalizer is not on the glib.Object but on the event.
	e := &Event{c}
	runtime.SetFinalizer(e, (*Event).free)
	return e, nil
}

// PutEvent is a wrapper around gdk_display_put_event().
func (v *Display) PutEvent(event *Event) {
	C.gdk_display_put_event(v.native(), event.native())
}

// HasPending is a wrapper around gdk_display_has_pending().
func (v *Display) HasPending() bool {
	c := C.gdk_display_has_pending(v.native())
	return gobool(c)
}

// SetDoubleClickTime is a wrapper around gdk_display_set_double_click_time().
func (v *Display) SetDoubleClickTime(msec uint) {
	C.gdk_display_set_double_click_time(v.native(), C.guint(msec))
}

// SetDoubleClickDistance is a wrapper around gdk_display_set_double_click_distance().
func (v *Display) SetDoubleClickDistance(distance uint) {
	C.gdk_display_set_double_click_distance(v.native(), C.guint(distance))
}

// SupportsColorCursor is a wrapper around gdk_display_supports_cursor_color().
func (v *Display) SupportsColorCursor() bool {
	c := C.gdk_display_supports_cursor_color(v.native())
	return gobool(c)
}

// SupportsCursorAlpha is a wrapper around gdk_display_supports_cursor_alpha().
func (v *Display) SupportsCursorAlpha() bool {
	c := C.gdk_display_supports_cursor_alpha(v.native())
	return gobool(c)
}

// GetDefaultCursorSize is a wrapper around gdk_display_get_default_cursor_size().
func (v *Display) GetDefaultCursorSize() uint {
	c := C.gdk_display_get_default_cursor_size(v.native())
	return uint(c)
}

// GetMaximalCursorSize is a wrapper around gdk_display_get_maximal_cursor_size().
func (v *Display) GetMaximalCursorSize() (width, height uint) {
	var w, h C.guint
	C.gdk_display_get_maximal_cursor_size(v.native(), &w, &h)
	return uint(w), uint(h)
}

// GetDefaultGroup is a wrapper around gdk_display_get_default_group().
func (v *Display) GetDefaultGroup() (*Window, error) {
	c := C.gdk_display_get_default_group(v.native())
	if c == nil {
		return nil, nilPtrErr
	}

	return &Window{glib.Take(unsafe.Pointer(c))}, nil
}

// SupportsSelectionNotification is a wrapper around
// gdk_display_supports_selection_notification().
func (v *Display) SupportsSelectionNotification() bool {
	c := C.gdk_display_supports_selection_notification(v.native())
	return gobool(c)
}

// RequestSelectionNotification is a wrapper around
// gdk_display_request_selection_notification().
func (v *Display) RequestSelectionNotification(selection Atom) bool {
	c := C.gdk_display_request_selection_notification(v.native(),
		selection.native())
	return gobool(c)
}

// SupportsClipboardPersistence is a wrapper around
// gdk_display_supports_clipboard_persistence().
func (v *Display) SupportsClipboardPersistence() bool {
	c := C.gdk_display_supports_clipboard_persistence(v.native())
	return gobool(c)
}

// StoreClipboard is a wrapper around gdk_display_store_clipboard().
func (v *Display) StoreClipboard(clipboardWindow *Window, time uint32, targets ...Atom) {
	atoms := make([]C.GdkAtom, len(targets))
	for i, target := range targets {
		atoms[i] = target.gdkAtom
	}
	C.gdk_display_store_clipboard(v.native(), clipboardWindow.native(), C.guint32(time),
		(*C.GdkAtom)(&atoms[0]), C.gint(len(targets)))
}

// SupportsShapes is a wrapper around gdk_display_supports_shapes().
func (v *Display) SupportsShapes() bool {
	c := C.gdk_display_supports_shapes(v.native())
	return gobool(c)
}

// SupportsInputShapes is a wrapper around gdk_display_supports_input_shapes().
func (v *Display) SupportsInputShapes() bool {
	c := C.gdk_display_supports_input_shapes(v.native())
	return gobool(c)
}

// TODO(jrick) glib.AppLaunchContext GdkAppLaunchContext
func (v *Display) GetAppLaunchContext() {
	panic("Not implemented")
}

// NotifyStartupComplete is a wrapper around gdk_display_notify_startup_complete().
func (v *Display) NotifyStartupComplete(startupID string) {
	cstr := C.CString(startupID)
	defer C.free(unsafe.Pointer(cstr))
	C.gdk_display_notify_startup_complete(v.native(), (*C.gchar)(cstr))
}

/*
 * GDK Keyval
 */

// KeyvalFromName is a wrapper around gdk_keyval_from_name().
func KeyvalFromName(keyvalName string) uint {
	cstr := C.CString(keyvalName)
	defer C.free(unsafe.Pointer(cstr))
	return uint(C.gdk_keyval_from_name((*C.gchar)(cstr)))
}

func KeyvalConvertCase(v uint) (lower, upper uint) {
	var l, u C.guint
	l = 0
	u = 0
	C.gdk_keyval_convert_case(C.guint(v), &l, &u)
	return uint(l), uint(u)
}

func KeyvalIsLower(v uint) bool {
	return gobool(C.gdk_keyval_is_lower(C.guint(v)))
}

func KeyvalIsUpper(v uint) bool {
	return gobool(C.gdk_keyval_is_upper(C.guint(v)))
}

func KeyvalToLower(v uint) uint {
	return uint(C.gdk_keyval_to_lower(C.guint(v)))
}

func KeyvalToUpper(v uint) uint {
	return uint(C.gdk_keyval_to_upper(C.guint(v)))
}

func KeyvalToUnicode(v uint) rune {
	return rune(C.gdk_keyval_to_unicode(C.guint(v)))
}

func UnicodeToKeyval(v rune) uint {
	return uint(C.gdk_unicode_to_keyval(C.guint32(v)))
}

/*
 * GdkDragContext
 */

// DragContext is a representation of GDK's GdkDragContext.
type DragContext struct {
	*glib.Object
}

// native returns a pointer to the underlying GdkDragContext.
func (v *DragContext) native() *C.GdkDragContext {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toGdkDragContext(ptr)
}

func marshalDragContext(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.ToObject(unsafe.Pointer(c))
	return &DragContext{obj}, nil
}

// ListTargets is a representation of gdk_drag_context_list_targets().
func (v *DragContext) ListTargets() *glib.List {
	clist := C.gdk_drag_context_list_targets(v.native())

	glist := glib.WrapList(uintptr(unsafe.Pointer(clist)))
	glist.DataWrapper(func(ptr unsafe.Pointer) interface{} {
		atom := Atom{C.GdkAtom(ptr)}
		return atom
	})

	if glist != nil {
		runtime.SetFinalizer(glist, func(glist *glib.List) {
			glist.Free()
		})
	}

	return glist
}

// RGBA is a representation of GDK's GdkRGBA type.
type RGBA struct {
	rgba *C.GdkRGBA
}

func marshalRGBA(p uintptr) (interface{}, error) {
	c := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	c2 := (*C.GdkRGBA)(unsafe.Pointer(c))
	return wrapRGBA(c2), nil
}

func wrapRGBA(obj *C.GdkRGBA) *RGBA {
	return &RGBA{obj}
}

func NewRGBA(values ...float64) *RGBA {
	cval := C.GdkRGBA{}
	v := &RGBA{&cval}
	if len(values) > 0 {
		v.rgba.red = C.gdouble(values[0])
	}
	if len(values) > 1 {
		v.rgba.green = C.gdouble(values[1])
	}
	if len(values) > 2 {
		v.rgba.blue = C.gdouble(values[2])
	}
	if len(values) > 3 {
		v.rgba.alpha = C.gdouble(values[3])
	}
	return v
}

func (v *RGBA) Floats() []float64 {
	return []float64{float64(v.rgba.red), float64(v.rgba.green), float64(v.rgba.blue), float64(v.rgba.alpha)}
}

func (v *RGBA) Native() uintptr {
	return uintptr(unsafe.Pointer(v.rgba))
}

// Parse is a representation of gdk_rgba_parse().
func (v *RGBA) Parse(spec string) bool {
	cstr := C.CString(spec)
	defer C.free(unsafe.Pointer(cstr))

	return gobool(C.gdk_rgba_parse(v.rgba, (*C.gchar)(cstr)))
}

// String is a representation of gdk_rgba_to_string().
func (v *RGBA) String() string {
	c := C.gdk_rgba_to_string(v.rgba)
	return goString(c)
}

// GdkRGBA * 	gdk_rgba_copy ()
// void 	gdk_rgba_free ()
// gboolean 	gdk_rgba_equal ()
// guint 	gdk_rgba_hash ()

/*
 * GdkRectangle
 */

// Rectangle is a representation of GDK's GdkRectangle type.
type Rectangle struct {
	gdkRectangle *C.GdkRectangle
}

func marshalRectangle(p uintptr) (interface{}, error) {
	c := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	c2 := (*C.GdkRectangle)(unsafe.Pointer(c))
	return wrapRectangle(c2), nil
}

func wrapRectangle(obj *C.GdkRectangle) *Rectangle {
	return &Rectangle{obj}
}

func WrapRectangle(p uintptr) *Rectangle {
	return wrapRectangle((*C.GdkRectangle)(unsafe.Pointer(p)))
}

func (v *Rectangle) Native() uintptr {
	if v == nil {
		return 0
	}
	return uintptr(unsafe.Pointer(v.gdkRectangle))
}

// Native() returns a pointer to the underlying GdkRectangle.
func (v *Rectangle) native() *C.GdkRectangle {
	if v == nil {
		return nil
	}
	return v.gdkRectangle
}

// GetX returns x field of the underlying GdkRectangle.
func (v *Rectangle) GetX() int {
	return int(v.native().x)
}

// SetX set x field of the underlying GdkRectangle.
func (v *Rectangle) SetX(x int) {
	v.native().x = C.int(x)
}

// GetY returns y field of the underlying GdkRectangle.
func (v *Rectangle) GetY() int {
	return int(v.native().y)
}

// SetY set y field of the underlying GdkRectangle.
func (v *Rectangle) SetY(y int) {
	v.native().y = C.int(y)
}

// GetWidth returns width field of the underlying GdkRectangle.
func (v *Rectangle) GetWidth() int {
	return int(v.native().width)
}

// SetWidth set width field of the underlying GdkRectangle.
func (v *Rectangle) SetWidth(width int) {
	v.native().width = C.int(width)
}

// GetHeight returns height field of the underlying GdkRectangle.
func (v *Rectangle) GetHeight() int {
	return int(v.native().height)
}

// SetHeight set height field of the underlying GdkRectangle.
func (v *Rectangle) SetHeight(height int) {
	v.native().height = C.int(height)
}

/*
 * GdkGeometry
 */

// Geometry is a representation of GDK's GdkGeometry type.
type Geometry struct {
	gdkGeometry *C.GdkGeometry
}

func marshalGeometry(p uintptr) (interface{}, error) {
	c := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	c2 := (*C.GdkGeometry)(unsafe.Pointer(c))
	return wrapGeometry(c2), nil
}

func wrapGeometry(obj *C.GdkGeometry) *Geometry {
	return &Geometry{obj}
}

func WrapGeometry(p uintptr) *Geometry {
	return wrapGeometry((*C.GdkGeometry)(unsafe.Pointer(p)))
}

func (v *Geometry) Native() uintptr {
	if v == nil {
		return 0
	}
	return uintptr(unsafe.Pointer(v.gdkGeometry))
}

// native returns a pointer to the underlying GdkGeometry.
func (v *Geometry) native() *C.GdkGeometry {
	if v == nil {
		return nil
	}
	return v.gdkGeometry
}

// GetMinWidth returns min_width field of the underlying GdkGeometry.
func (v *Geometry) GetMinWidth() int {
	return int(v.native().min_width)
}

// SetMinWidth set min_width field of the underlying GdkGeometry.
func (v *Geometry) SetMinWidth(minWidth int) {
	v.native().min_width = C.gint(minWidth)
}

// GetMinHeight returns min_height field of the underlying GdkGeometry.
func (v *Geometry) GetMinHeight() int {
	return int(v.native().min_height)
}

// SetMinHeight set min_height field of the underlying GdkGeometry.
func (v *Geometry) SetMinHeight(minHeight int) {
	v.native().min_height = C.gint(minHeight)
}

// GetMaxWidth returns max_width field of the underlying GdkGeometry.
func (v *Geometry) GetMaxWidth() int {
	return int(v.native().max_width)
}

// SetMaxWidth set max_width field of the underlying GdkGeometry.
func (v *Geometry) SetMaxWidth(maxWidth int) {
	v.native().max_width = C.gint(maxWidth)
}

// GetMaxHeight returns max_height field of the underlying GdkGeometry.
func (v *Geometry) GetMaxHeight() int {
	return int(v.native().max_height)
}

// SetMaxHeight set max_height field of the underlying GdkGeometry.
func (v *Geometry) SetMaxHeight(maxHeight int) {
	v.native().max_height = C.gint(maxHeight)
}

// GetBaseWidth returns base_width field of the underlying GdkGeometry.
func (v *Geometry) GetBaseWidth() int {
	return int(v.native().base_width)
}

// SetBaseWidth set base_width field of the underlying GdkGeometry.
func (v *Geometry) SetBaseWidth(baseWidth int) {
	v.native().base_width = C.gint(baseWidth)
}

// GetBaseHeight returns base_height field of the underlying GdkGeometry.
func (v *Geometry) GetBaseHeight() int {
	return int(v.native().base_height)
}

// SetBaseHeight set base_height field of the underlying GdkGeometry.
func (v *Geometry) SetBaseHeight(baseHeight int) {
	v.native().base_height = C.gint(baseHeight)
}

// GetWidthInc returns width_inc field of the underlying GdkGeometry.
func (v *Geometry) GetWidthInc() int {
	return int(v.native().width_inc)
}

// SetWidthInc set width_inc field of the underlying GdkGeometry.
func (v *Geometry) SetWidthInc(widthInc int) {
	v.native().width_inc = C.gint(widthInc)
}

// GetHeightInc returns height_inc field of the underlying GdkGeometry.
func (v *Geometry) GetHeightInc() int {
	return int(v.native().height_inc)
}

// SetHeightInc set height_inc field of the underlying GdkGeometry.
func (v *Geometry) SetHeightInc(heightInc int) {
	v.native().height_inc = C.gint(heightInc)
}

// GetMinAspect returns min_aspect field of the underlying GdkGeometry.
func (v *Geometry) GetMinAspect() float64 {
	return float64(v.native().min_aspect)
}

// SetMinAspect set min_aspect field of the underlying GdkGeometry.
func (v *Geometry) SetMinAspect(minAspect float64) {
	v.native().min_aspect = C.gdouble(minAspect)
}

// GetMaxAspect returns max_aspect field of the underlying GdkGeometry.
func (v *Geometry) GetMaxAspect() float64 {
	return float64(v.native().max_aspect)
}

// SetMaxAspect set max_aspect field of the underlying GdkGeometry.
func (v *Geometry) SetMaxAspect(maxAspect float64) {
	v.native().max_aspect = C.gdouble(maxAspect)
}

// GetWinGravity returns win_gravity field of the underlying GdkGeometry.
func (v *Geometry) GetWinGravity() Gravity {
	return Gravity(v.native().win_gravity)
}

// SetWinGravity set win_gravity field of the underlying GdkGeometry.
func (v *Geometry) SetWinGravity(winGravity Gravity) {
	v.native().win_gravity = C.GdkGravity(winGravity)
}

/*
 * GdkVisual
 */

// Visual is a representation of GDK's GdkVisual.
type Visual struct {
	*glib.Object
}

func (v *Visual) native() *C.GdkVisual {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toGdkVisual(ptr)
}

func marshalVisual(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.ToObject(unsafe.Pointer(c))
	return &Visual{obj}, nil
}

/*
 * GdkWindow
 */

// Window is a representation of GDK's GdkWindow.
type Window struct {
	*glib.Object
}

// native returns a pointer to the underlying GdkWindow.
func (v *Window) native() *C.GdkWindow {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toGdkWindow(ptr)
}

func marshalWindow(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.ToObject(unsafe.Pointer(c))
	return &Window{obj}, nil
}

func toWindow(s *C.GdkWindow) (*Window, error) {
	if s == nil {
		return nil, nilPtrErr
	}
	obj := glib.ToObject(unsafe.Pointer(s))
	return &Window{obj}, nil
}

// SetCursor is a wrapper around gdk_window_set_cursor().
func (v *Window) SetCursor(cursor *Cursor) {
	C.gdk_window_set_cursor(v.native(), cursor.native())
}
