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

package gtk

// #include <gtk/gtk.h>
// #include "gtk.go.h"
import "C"
import (
	"errors"
	"unsafe"

	"github.com/d2r2/gotk3/cairo"
	"github.com/d2r2/gotk3/gdk"
	"github.com/d2r2/gotk3/glib"
)

// IWidget is an interface type implemented by all structs
// embedding a Widget.  It is meant to be used as an argument type
// for wrapper functions that wrap around a C GTK function taking a
// GtkWidget.
type IWidget interface {
	toWidget() *C.GtkWidget
	GetWidget() *Widget
	Set(string, interface{}) error
}

/*
type IWidgetable interface {
	toWidget() *C.GtkWidget
}
*/

/*
 * GtkWidget
 */

// Widget is a representation of GTK's GtkWidget.
type Widget struct {
	glib.InitiallyUnowned
}

var _ IWidget = &Widget{}

//var _ IWidgetable = &Widget{}

// native returns a pointer to the underlying GtkWidget.
func (v *Widget) native() *C.GtkWidget {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toGtkWidget(ptr)
}

func marshalWidget(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapWidget(obj), nil
}

func wrapWidget(obj *glib.Object) *Widget {
	return &Widget{glib.InitiallyUnowned{obj}}
}

func (v *Widget) toWidget() *C.GtkWidget {
	if v == nil {
		return nil
	}
	return v.native()
}

func (v *Widget) GetWidget() *Widget {
	return v
}

// Destroy is a wrapper around gtk_widget_destroy().
func (v *Widget) Destroy() {
	C.gtk_widget_destroy(v.native())
}

func (v *Widget) HideOnDelete() {
	C._gtk_widget_hide_on_delete(v.native())
}

/* TODO
func (v *Widget) DragDestSet(flags DestDefaults, targets []TargetEntry, actions gdk.DragAction) {
	C.gtk_drag_dest_set(v.native(), C.GtkDestDefaults(flags), (*C.GtkTargetEntry)(&targets[0]),
		C.gint(len(targets)), C.GdkDragAction(actions))
}
*/

// ResetStyle is a wrapper around gtk_widget_reset_style().
func (v *Widget) ResetStyle() {
	C.gtk_widget_reset_style(v.native())
}

// InDestruction is a wrapper around gtk_widget_in_destruction().
func (v *Widget) InDestruction() bool {
	return gobool(C.gtk_widget_in_destruction(v.native()))
}

// TODO(jrick) this may require some rethinking
/*
func (v *Widget) Destroyed(widgetPointer **Widget) {
}
*/

// Unparent is a wrapper around gtk_widget_unparent().
func (v *Widget) Unparent() {
	C.gtk_widget_unparent(v.native())
}

// Show is a wrapper around gtk_widget_show().
func (v *Widget) Show() {
	C.gtk_widget_show(v.native())
}

// Hide is a wrapper around gtk_widget_hide().
func (v *Widget) Hide() {
	C.gtk_widget_hide(v.native())
}

// GetCanFocus is a wrapper around gtk_widget_get_can_focus().
func (v *Widget) GetCanFocus() bool {
	c := C.gtk_widget_get_can_focus(v.native())
	return gobool(c)
}

// SetCanFocus is a wrapper around gtk_widget_set_can_focus().
func (v *Widget) SetCanFocus(canFocus bool) {
	C.gtk_widget_set_can_focus(v.native(), gbool(canFocus))
}

// GetCanDefault is a wrapper around gtk_widget_get_can_default().
func (v *Widget) GetCanDefault() bool {
	c := C.gtk_widget_get_can_default(v.native())
	return gobool(c)
}

// SetCanDefault is a wrapper around gtk_widget_set_can_default().
func (v *Widget) SetCanDefault(canDefault bool) {
	C.gtk_widget_set_can_default(v.native(), gbool(canDefault))
}

// GetMapped is a wrapper around gtk_widget_get_mapped().
func (v *Widget) GetMapped() bool {
	c := C.gtk_widget_get_mapped(v.native())
	return gobool(c)
}

// SetMapped is a wrapper around gtk_widget_set_mapped().
func (v *Widget) SetMapped(mapped bool) {
	C.gtk_widget_set_can_focus(v.native(), gbool(mapped))
}

// GetRealized is a wrapper around gtk_widget_get_realized().
func (v *Widget) GetRealized() bool {
	c := C.gtk_widget_get_realized(v.native())
	return gobool(c)
}

// SetRealized is a wrapper around gtk_widget_set_realized().
func (v *Widget) SetRealized(realized bool) {
	C.gtk_widget_set_realized(v.native(), gbool(realized))
}

// GetHasWindow is a wrapper around gtk_widget_get_has_window().
func (v *Widget) GetHasWindow() bool {
	c := C.gtk_widget_get_has_window(v.native())
	return gobool(c)
}

// SetHasWindow is a wrapper around gtk_widget_set_has_window().
func (v *Widget) SetHasWindow(hasWindow bool) {
	C.gtk_widget_set_has_window(v.native(), gbool(hasWindow))
}

// GetStyleContext is a wrapper around gtk_widget_get_style_context().
func (v *Widget) GetStyleContext() (*StyleContext, error) {
	c := C.gtk_widget_get_style_context(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapStyleContext(obj), nil
}

// ShowNow is a wrapper around gtk_widget_show_now().
func (v *Widget) ShowNow() {
	C.gtk_widget_show_now(v.native())
}

// ShowAll is a wrapper around gtk_widget_show_all().
func (v *Widget) ShowAll() {
	C.gtk_widget_show_all(v.native())
}

// SetNoShowAll is a wrapper around gtk_widget_set_no_show_all().
func (v *Widget) SetNoShowAll(noShowAll bool) {
	C.gtk_widget_set_no_show_all(v.native(), gbool(noShowAll))
}

// GetNoShowAll is a wrapper around gtk_widget_get_no_show_all().
func (v *Widget) GetNoShowAll() bool {
	c := C.gtk_widget_get_no_show_all(v.native())
	return gobool(c)
}

// Map is a wrapper around gtk_widget_map().
func (v *Widget) Map() {
	C.gtk_widget_map(v.native())
}

// Unmap is a wrapper around gtk_widget_unmap().
func (v *Widget) Unmap() {
	C.gtk_widget_unmap(v.native())
}

// QueueDrawArea is a wrapper aroung gtk_widget_queue_draw_area().
func (v *Widget) QueueDrawArea(x, y, w, h int) {
	C.gtk_widget_queue_draw_area(v.native(), C.gint(x), C.gint(y), C.gint(w), C.gint(h))
}

// AddAccelerator is a wrapper around gtk_widget_add_accelerator().
func (v *Widget) AddAccelerator(signal string, group *AccelGroup, key uint, mods gdk.ModifierType, flags AccelFlags) {
	csignal := C.CString(signal)
	defer C.free(unsafe.Pointer(csignal))

	C.gtk_widget_add_accelerator(v.native(), (*C.gchar)(csignal), group.native(),
		C.guint(key), C.GdkModifierType(mods), C.GtkAccelFlags(flags))
}

// RemoveAccelerator is a wrapper around gtk_widget_remove_accelerator().
func (v *Widget) RemoveAccelerator(group *AccelGroup, key uint, mods gdk.ModifierType) bool {
	return gobool(C.gtk_widget_remove_accelerator(v.native(), group.native(),
		C.guint(key), C.GdkModifierType(mods)))
}

// SetAccelPath is a wrapper around gtk_widget_set_accel_path().
func (v *Widget) SetAccelPath(path string, group *AccelGroup) {
	cstr := C.CString(path)
	defer C.free(unsafe.Pointer(cstr))

	C.gtk_widget_set_accel_path(v.native(), (*C.gchar)(cstr), group.native())
}

// CanActivateAccel is a wrapper around gtk_widget_can_activate_accel().
func (v *Widget) CanActivateAccel(signalId uint) bool {
	return gobool(C.gtk_widget_can_activate_accel(v.native(), C.guint(signalId)))
}

// AddAccelGroup() is a wrapper around gtk_window_add_accel_group().
func (v *Window) AddAccelGroup(accelGroup *AccelGroup) {
	C.gtk_window_add_accel_group(v.native(), accelGroup.native())
}

// RemoveAccelGroup() is a wrapper around gtk_window_add_accel_group().
func (v *Window) RemoveAccelGroup(accelGroup *AccelGroup) {
	C.gtk_window_remove_accel_group(v.native(), accelGroup.native())
}

//void gtk_widget_realize(GtkWidget *widget);
//void gtk_widget_unrealize(GtkWidget *widget);
//void gtk_widget_draw(GtkWidget *widget, cairo_t *cr);
//void gtk_widget_queue_resize(GtkWidget *widget);
//void gtk_widget_queue_resize_no_redraw(GtkWidget *widget);
//GdkFrameClock *gtk_widget_get_frame_clock(GtkWidget *widget);
//guint gtk_widget_add_tick_callback (GtkWidget *widget,
//                                    GtkTickCallback callback,
//                                    gpointer user_data,
//                                    GDestroyNotify notify);
//void gtk_widget_remove_tick_callback(GtkWidget *widget, guint id);

// TODO(jrick) GtkAllocation
/*
func (v *Widget) SizeAllocate() {
}
*/

// Allocation is a representation of GTK's GtkAllocation type.
type Allocation struct {
	gdk.Rectangle
}

// Native returns a pointer to the underlying GtkAllocation.
func (v *Allocation) native() *C.GtkAllocation {
	return (*C.GtkAllocation)(unsafe.Pointer(v.Native()))
}

// GetAllocatedWidth() is a wrapper around gtk_widget_get_allocated_width().
func (v *Widget) GetAllocatedWidth() int {
	return int(C.gtk_widget_get_allocated_width(v.native()))
}

// GetAllocatedHeight() is a wrapper around gtk_widget_get_allocated_height().
func (v *Widget) GetAllocatedHeight() int {
	return int(C.gtk_widget_get_allocated_height(v.native()))
}

// Event() is a wrapper around gtk_widget_event().
func (v *Widget) Event(event *gdk.Event) bool {
	c := C.gtk_widget_event(v.native(),
		(*C.GdkEvent)(unsafe.Pointer(event.Native())))
	return gobool(c)
}

// Activate() is a wrapper around gtk_widget_activate().
func (v *Widget) Activate() bool {
	return gobool(C.gtk_widget_activate(v.native()))
}

// TODO(jrick) GdkRectangle
/*
func (v *Widget) Intersect() {
}
*/

// IsFocus() is a wrapper around gtk_widget_is_focus().
func (v *Widget) IsFocus() bool {
	return gobool(C.gtk_widget_is_focus(v.native()))
}

// GrabFocus() is a wrapper around gtk_widget_grab_focus().
func (v *Widget) GrabFocus() {
	C.gtk_widget_grab_focus(v.native())
}

// GrabDefault() is a wrapper around gtk_widget_grab_default().
func (v *Widget) GrabDefault() {
	C.gtk_widget_grab_default(v.native())
}

// SetName() is a wrapper around gtk_widget_set_name().
func (v *Widget) SetName(name string) {
	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_widget_set_name(v.native(), (*C.gchar)(cstr))
}

// GetName() is a wrapper around gtk_widget_get_name().  A non-nil
// error is returned in the case that gtk_widget_get_name returns NULL to
// differentiate between NULL and an empty string.
func (v *Widget) GetName() (string, error) {
	c := C.gtk_widget_get_name(v.native())
	if c == nil {
		return "", nilPtrErr
	}
	return goString(c), nil
}

// GetSensitive is a wrapper around gtk_widget_get_sensitive().
func (v *Widget) GetSensitive() bool {
	c := C.gtk_widget_get_sensitive(v.native())
	return gobool(c)
}

// IsSensitive is a wrapper around gtk_widget_is_sensitive().
func (v *Widget) IsSensitive() bool {
	c := C.gtk_widget_is_sensitive(v.native())
	return gobool(c)
}

// SetSensitive is a wrapper around gtk_widget_set_sensitive().
func (v *Widget) SetSensitive(sensitive bool) {
	C.gtk_widget_set_sensitive(v.native(), gbool(sensitive))
}

// GetVisible is a wrapper around gtk_widget_get_visible().
func (v *Widget) GetVisible() bool {
	c := C.gtk_widget_get_visible(v.native())
	return gobool(c)
}

// SetVisible is a wrapper around gtk_widget_set_visible().
func (v *Widget) SetVisible(visible bool) {
	C.gtk_widget_set_visible(v.native(), gbool(visible))
}

// SetParent is a wrapper around gtk_widget_set_parent().
func (v *Widget) SetParent(parent IWidget) {
	C.gtk_widget_set_parent(v.native(), parent.toWidget())
}

// GetParent is a wrapper around gtk_widget_get_parent().
func (v *Widget) GetParent() (*Widget, error) {
	c := C.gtk_widget_get_parent(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapWidget(obj), nil
}

// SetSizeRequest is a wrapper around gtk_widget_set_size_request().
func (v *Widget) SetSizeRequest(width, height int) {
	C.gtk_widget_set_size_request(v.native(), C.gint(width), C.gint(height))
}

// GetSizeRequest is a wrapper around gtk_widget_get_size_request().
func (v *Widget) GetSizeRequest() (width, height int) {
	var w, h C.gint
	C.gtk_widget_get_size_request(v.native(), &w, &h)
	return int(w), int(h)
}

// SetParentWindow is a wrapper around gtk_widget_set_parent_window().
func (v *Widget) SetParentWindow(parentWindow *gdk.Window) {
	C.gtk_widget_set_parent_window(v.native(),
		(*C.GdkWindow)(unsafe.Pointer(parentWindow.Native())))
}

// GetParentWindow is a wrapper around gtk_widget_get_parent_window().
func (v *Widget) GetParentWindow() (*gdk.Window, error) {
	c := C.gtk_widget_get_parent_window(v.native())
	if v == nil {
		return nil, nilPtrErr
	}

	w := &gdk.Window{glib.Take(unsafe.Pointer(c))}
	return w, nil
}

// SetEvents is a wrapper around gtk_widget_set_events().
func (v *Widget) SetEvents(events int) {
	C.gtk_widget_set_events(v.native(), C.gint(events))
}

// GetEvents is a wrapper around gtk_widget_get_events().
func (v *Widget) GetEvents() int {
	return int(C.gtk_widget_get_events(v.native()))
}

// AddEvents is a wrapper around gtk_widget_add_events().
func (v *Widget) AddEvents(events int) {
	C.gtk_widget_add_events(v.native(), C.gint(events))
}

// HasDefault is a wrapper around gtk_widget_has_default().
func (v *Widget) HasDefault() bool {
	c := C.gtk_widget_has_default(v.native())
	return gobool(c)
}

// HasFocus is a wrapper around gtk_widget_has_focus().
func (v *Widget) HasFocus() bool {
	c := C.gtk_widget_has_focus(v.native())
	return gobool(c)
}

// HasVisibleFocus is a wrapper around gtk_widget_has_visible_focus().
func (v *Widget) HasVisibleFocus() bool {
	c := C.gtk_widget_has_visible_focus(v.native())
	return gobool(c)
}

// HasGrab is a wrapper around gtk_widget_has_grab().
func (v *Widget) HasGrab() bool {
	c := C.gtk_widget_has_grab(v.native())
	return gobool(c)
}

// IsDrawable is a wrapper around gtk_widget_is_drawable().
func (v *Widget) IsDrawable() bool {
	c := C.gtk_widget_is_drawable(v.native())
	return gobool(c)
}

// IsToplevel is a wrapper around gtk_widget_is_toplevel().
func (v *Widget) IsToplevel() bool {
	c := C.gtk_widget_is_toplevel(v.native())
	return gobool(c)
}

// TODO(jrick) GdkEventMask
/*
func (v *Widget) SetDeviceEvents() {
}
*/

// TODO(jrick) GdkEventMask
/*
func (v *Widget) GetDeviceEvents() {
}
*/

// TODO(jrick) GdkEventMask
/*
func (v *Widget) AddDeviceEvents() {
}
*/

// SetDeviceEnabled is a wrapper around gtk_widget_set_device_enabled().
func (v *Widget) SetDeviceEnabled(device *gdk.Device, enabled bool) {
	C.gtk_widget_set_device_enabled(v.native(),
		(*C.GdkDevice)(unsafe.Pointer(device.Native())), gbool(enabled))
}

// GetDeviceEnabled is a wrapper around gtk_widget_get_device_enabled().
func (v *Widget) GetDeviceEnabled(device *gdk.Device) bool {
	c := C.gtk_widget_get_device_enabled(v.native(),
		(*C.GdkDevice)(unsafe.Pointer(device.Native())))
	return gobool(c)
}

// GetToplevel is a wrapper around gtk_widget_get_toplevel().
func (v *Widget) GetToplevel() (*Widget, error) {
	c := C.gtk_widget_get_toplevel(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapWidget(obj), nil
}

// GetTooltipText is a wrapper around gtk_widget_get_tooltip_text().
// A non-nil error is returned in the case that
// gtk_widget_get_tooltip_text returns NULL to differentiate between NULL
// and an empty string.
func (v *Widget) GetTooltipText() string {
	c := C.gtk_widget_get_tooltip_text(v.native())
	defer C.g_free(C.gpointer(c))
	return goString(c)
}

// SetTooltipText is a wrapper around gtk_widget_set_tooltip_text().
func (v *Widget) SetTooltipText(text string) {
	var cstr *C.char
	if text != "" {
		cstr = C.CString(text)
		defer C.free(unsafe.Pointer(cstr))
	}
	C.gtk_widget_set_tooltip_text(v.native(), (*C.gchar)(cstr))
}

// GetTooltipMarkup is a wrapper around gtk_widget_get_tooltip_markup().
func (v *Widget) GetTooltipMarkup() string {
	c := C.gtk_widget_get_tooltip_markup(v.native())
	defer C.g_free(C.gpointer(c))
	return goString(c)
}

// SetTooltipMarkup is a wrapper around gtk_widget_set_tooltip_markup().
func (v *Widget) SetTooltipMarkup(text string) {
	var cstr *C.char
	if text != "" {
		cstr = C.CString(text)
		defer C.free(unsafe.Pointer(cstr))
	}
	C.gtk_widget_set_tooltip_markup(v.native(), (*C.gchar)(cstr))
}

// GetHAlign is a wrapper around gtk_widget_get_halign().
func (v *Widget) GetHAlign() Align {
	c := C.gtk_widget_get_halign(v.native())
	return Align(c)
}

// SetHAlign is a wrapper around gtk_widget_set_halign().
func (v *Widget) SetHAlign(align Align) {
	C.gtk_widget_set_halign(v.native(), C.GtkAlign(align))
}

// GetVAlign is a wrapper around gtk_widget_get_valign().
func (v *Widget) GetVAlign() Align {
	c := C.gtk_widget_get_valign(v.native())
	return Align(c)
}

// SetVAlign is a wrapper around gtk_widget_set_valign().
func (v *Widget) SetVAlign(align Align) {
	C.gtk_widget_set_valign(v.native(), C.GtkAlign(align))
}

// GetMarginTop is a wrapper around gtk_widget_get_margin_top().
func (v *Widget) GetMarginTop() int {
	c := C.gtk_widget_get_margin_top(v.native())
	return int(c)
}

// SetMarginTop is a wrapper around gtk_widget_set_margin_top().
func (v *Widget) SetMarginTop(margin int) {
	C.gtk_widget_set_margin_top(v.native(), C.gint(margin))
}

// GetMarginBottom is a wrapper around gtk_widget_get_margin_bottom().
func (v *Widget) GetMarginBottom() int {
	c := C.gtk_widget_get_margin_bottom(v.native())
	return int(c)
}

// SetMarginBottom is a wrapper around gtk_widget_set_margin_bottom().
func (v *Widget) SetMarginBottom(margin int) {
	C.gtk_widget_set_margin_bottom(v.native(), C.gint(margin))
}

// GetHExpand is a wrapper around gtk_widget_get_hexpand().
func (v *Widget) GetHExpand() bool {
	c := C.gtk_widget_get_hexpand(v.native())
	return gobool(c)
}

// SetHExpand is a wrapper around gtk_widget_set_hexpand().
func (v *Widget) SetHExpand(expand bool) {
	C.gtk_widget_set_hexpand(v.native(), gbool(expand))
}

// GetVExpand is a wrapper around gtk_widget_get_vexpand().
func (v *Widget) GetVExpand() bool {
	c := C.gtk_widget_get_vexpand(v.native())
	return gobool(c)
}

// SetVExpand is a wrapper around gtk_widget_set_vexpand().
func (v *Widget) SetVExpand(expand bool) {
	C.gtk_widget_set_vexpand(v.native(), gbool(expand))
}

// TranslateCoordinates is a wrapper around gtk_widget_translate_coordinates().
func (v *Widget) TranslateCoordinates(dest IWidget, srcX, srcY int) (destX, destY int, e error) {
	var cdestX, cdestY C.gint
	c := C.gtk_widget_translate_coordinates(v.native(), dest.toWidget(), C.gint(srcX), C.gint(srcY), &cdestX, &cdestY)
	if !gobool(c) {
		return 0, 0, errors.New("translate coordinates failed")
	}
	return int(cdestX), int(cdestY), nil
}

// SetVisual is a wrapper around gtk_widget_set_visual().
func (v *Widget) SetVisual(visual *gdk.Visual) {
	C.gtk_widget_set_visual(v.native(),
		(*C.GdkVisual)(unsafe.Pointer(visual.Native())))
}

// SetAppPaintable is a wrapper around gtk_widget_set_app_paintable().
func (v *Widget) SetAppPaintable(paintable bool) {
	C.gtk_widget_set_app_paintable(v.native(), gbool(paintable))
}

// GetAppPaintable is a wrapper around gtk_widget_get_app_paintable().
func (v *Widget) GetAppPaintable() bool {
	c := C.gtk_widget_get_app_paintable(v.native())
	return gobool(c)
}

// QueueDraw is a wrapper around gtk_widget_queue_draw().
func (v *Widget) QueueDraw() {
	C.gtk_widget_queue_draw(v.native())
}

// GetAllocation is a wrapper around gtk_widget_get_allocation().
func (v *Widget) GetAllocation() *Allocation {
	var a Allocation
	C.gtk_widget_get_allocation(v.native(), a.native())
	return &a
}

// SetAllocation is a wrapper around gtk_widget_set_allocation().
func (v *Widget) SetAllocation(allocation *Allocation) {
	C.gtk_widget_set_allocation(v.native(), allocation.native())
}

// SizeAllocate is a wrapper around gtk_widget_size_allocate().
func (v *Widget) SizeAllocate(allocation *Allocation) {
	C.gtk_widget_size_allocate(v.native(), allocation.native())
}

// SetStateFlags is a wrapper around gtk_widget_set_state_flags().
func (v *Widget) SetStateFlags(stateFlags StateFlags, clear bool) {
	C.gtk_widget_set_state_flags(v.native(), C.GtkStateFlags(stateFlags), gbool(clear))
}

// GetWindow is a wrapper around gtk_widget_get_window().
func (v *Widget) GetWindow() (*gdk.Window, error) {
	c := C.gtk_widget_get_window(v.native())
	if c == nil {
		return nil, nilPtrErr
	}

	w := &gdk.Window{glib.Take(unsafe.Pointer(c))}
	return w, nil
}

// GetPreferredHeight is a wrapper around gtk_widget_get_preferred_height().
func (v *Widget) GetPreferredHeight() (int, int) {
	var minimum, natural C.gint
	C.gtk_widget_get_preferred_height(v.native(), &minimum, &natural)
	return int(minimum), int(natural)
}

// GetPreferredWidth is a wrapper around gtk_widget_get_preferred_width().
func (v *Widget) GetPreferredWidth() (int, int) {
	var minimum, natural C.gint
	C.gtk_widget_get_preferred_width(v.native(), &minimum, &natural)
	return int(minimum), int(natural)
}

/*
 * GtkContainer
 */

// Container is a representation of GTK's GtkContainer.
type Container struct {
	Widget
}

// native returns a pointer to the underlying GtkContainer.
func (v *Container) native() *C.GtkContainer {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toGtkContainer(ptr)
}

func marshalContainer(p uintptr) (interface{}, error) {
	c := C.g_value_get_object(C.toGValue(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapContainer(obj), nil
}

func wrapContainer(obj *glib.Object) *Container {
	widget := wrapWidget(obj)
	return &Container{*widget}
}

func (v *Container) GetContainer() *Container {
	return v
}

// Add is a wrapper around gtk_container_add().
func (v *Container) Add(w IWidget) {
	C.gtk_container_add(v.native(), w.toWidget())
}

// Remove is a wrapper around gtk_container_remove().
func (v *Container) Remove(w IWidget) {
	C.gtk_container_remove(v.native(), w.toWidget())
}

// TODO: gtk_container_add_with_properties

// CheckResize is a wrapper around gtk_container_check_resize().
func (v *Container) CheckResize() {
	C.gtk_container_check_resize(v.native())
}

// TODO: gtk_container_foreach

// GetChildren is a wrapper around gtk_container_get_children().
func (v *Container) GetChildren() *glib.List {
	clist := C.gtk_container_get_children(v.native())
	glist := glib.WrapList(uintptr(unsafe.Pointer(clist)))
	glist.DataWrapper(func(ptr unsafe.Pointer) interface{} {
		return wrapWidget(glib.Take(ptr))
	})

	return glist
}

// TODO: gtk_container_get_path_for_child

// GetFocusChild is a wrapper around gtk_container_get_focus_child().
func (v *Container) GetFocusChild() *Widget {
	c := C.gtk_container_get_focus_child(v.native())
	if c == nil {
		return nil
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapWidget(obj)
}

// SetFocusChild is a wrapper around gtk_container_set_focus_child().
func (v *Container) SetFocusChild(child IWidget) {
	C.gtk_container_set_focus_child(v.native(), child.toWidget())
}

// GetFocusVAdjustment is a wrapper around
// gtk_container_get_focus_vadjustment().
func (v *Container) GetFocusVAdjustment() *Adjustment {
	c := C.gtk_container_get_focus_vadjustment(v.native())
	if c == nil {
		return nil
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapAdjustment(obj)
}

// SetFocusVAdjustment is a wrapper around
// gtk_container_set_focus_vadjustment().
func (v *Container) SetFocusVAdjustment(adjustment *Adjustment) {
	C.gtk_container_set_focus_vadjustment(v.native(), adjustment.native())
}

// GetFocusHAdjustment is a wrapper around
// gtk_container_get_focus_hadjustment().
func (v *Container) GetFocusHAdjustment() *Adjustment {
	c := C.gtk_container_get_focus_hadjustment(v.native())
	if c == nil {
		return nil
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapAdjustment(obj)
}

// SetFocusHAdjustment is a wrapper around
// gtk_container_set_focus_hadjustment().
func (v *Container) SetFocusHAdjustment(adjustment *Adjustment) {
	C.gtk_container_set_focus_hadjustment(v.native(), adjustment.native())
}

// ChildType is a wrapper around gtk_container_child_type().
func (v *Container) ChildType() glib.Type {
	c := C.gtk_container_child_type(v.native())
	return glib.Type(c)
}

// TODO: gtk_container_child_get_valist
// TODO: gtk_container_child_set_valist

// ChildNotify is a wrapper around gtk_container_child_notify().
func (v *Container) ChildNotify(child IWidget, childProperty string) {
	cstr := C.CString(childProperty)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_container_child_notify(v.native(), child.toWidget(),
		(*C.gchar)(cstr))
}

// ChildGetProperty is a wrapper around gtk_container_child_get_property().
func (v *Container) ChildGetProperty(child IWidget, name string, valueType glib.Type) (interface{}, error) {
	gv, e := glib.ValueInit(valueType)
	if e != nil {
		return nil, e
	}
	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))

	C.gtk_container_child_get_property(v.native(), child.toWidget(), (*C.gchar)(cstr),
		C.toGValue(unsafe.Pointer(gv.Native())))
	return gv.GoValue()
}

// ChildSetProperty is a wrapper around gtk_container_child_set_property().
func (v *Container) ChildSetProperty(child IWidget, name string, value interface{}) error {
	gv, e := glib.GValue(value)
	if e != nil {
		return e
	}
	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))

	C.gtk_container_child_set_property(v.native(), child.toWidget(), (*C.gchar)(cstr),
		C.toGValue(unsafe.Pointer(gv.Native())))
	return nil
}

// TODO: gtk_container_forall

// GetBorderWidth is a wrapper around gtk_container_get_border_width().
func (v *Container) GetBorderWidth() uint {
	c := C.gtk_container_get_border_width(v.native())
	return uint(c)
}

// SetBorderWidth is a wrapper around gtk_container_set_border_width().
func (v *Container) SetBorderWidth(borderWidth uint) {
	C.gtk_container_set_border_width(v.native(), C.guint(borderWidth))
}

// PropagateDraw is a wrapper around gtk_container_propagate_draw().
func (v *Container) PropagateDraw(child IWidget, cr *cairo.Context) {
	context := (*C.cairo_t)(unsafe.Pointer(cr.Native()))
	C.gtk_container_propagate_draw(v.native(), child.toWidget(), context)
}

// GdkCairoSetSourcePixBuf() is a wrapper around gdk_cairo_set_source_pixbuf().
func GdkCairoSetSourcePixBuf(cr *cairo.Context, pixbuf *gdk.Pixbuf, pixbufX, pixbufY float64) {
	context := (*C.cairo_t)(unsafe.Pointer(cr.Native()))
	ptr := (*C.GdkPixbuf)(unsafe.Pointer(pixbuf.Native()))
	C.gdk_cairo_set_source_pixbuf(context, ptr, C.gdouble(pixbufX), C.gdouble(pixbufY))
}

// GetFocusChain is a wrapper around gtk_container_get_focus_chain().
func (v *Container) GetFocusChain() ([]*Widget, bool) {
	var cwlist *C.GList
	c := C.gtk_container_get_focus_chain(v.native(), &cwlist)

	var widgets []*Widget
	wlist := glib.WrapList(uintptr(unsafe.Pointer(cwlist)))
	for ; wlist.Data() != nil; wlist = wlist.Next() {
		widgets = append(widgets, wrapWidget(glib.Take(wlist.Data().(unsafe.Pointer))))
	}
	return widgets, gobool(c)
}

// SetFocusChain is a wrapper around gtk_container_set_focus_chain().
func (v *Container) SetFocusChain(focusableWidgets []IWidget) {
	var list *glib.List
	for _, w := range focusableWidgets {
		data := uintptr(unsafe.Pointer(w.toWidget()))
		list = list.Append(data)
	}
	glist := (*C.GList)(unsafe.Pointer(list))
	C.gtk_container_set_focus_chain(v.native(), glist)
}

/*
 * GtkBin
 */

// Bin is a representation of GTK's GtkBin.
type Bin struct {
	Container
}

// native returns a pointer to the underlying GtkBin.
func (v *Bin) native() *C.GtkBin {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toGtkBin(ptr)
}

func marshalBin(p uintptr) (interface{}, error) {
	c := C.g_value_get_object(C.toGValue(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapBin(obj), nil
}

func wrapBin(obj *glib.Object) *Bin {
	container := wrapContainer(obj)
	return &Bin{*container}
}

// GetChild is a wrapper around gtk_bin_get_child().
func (v *Bin) GetChild() (*Widget, error) {
	c := C.gtk_bin_get_child(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapWidget(obj), nil
}
