// Same copyright and license as the rest of the files in this project
// This file contains accelerator related functions and structures

package gtk

// #include <gtk/gtk.h>
// #include "gtk.go.h"
import "C"
import (
	"errors"
	"fmt"
	"unsafe"

	"github.com/d2r2/gotk3/cairo"
	"github.com/d2r2/gotk3/gdk"
	"github.com/d2r2/gotk3/glib"
)

/*
 * GtkWindow
 */

// Window is a representation of GTK's GtkWindow.
type Window struct {
	Bin
}

// IWindow is an interface type implemented by all structs embedding a
// Window.  It is meant to be used as an argument type for wrapper
// functions that wrap around a C GTK function taking a GtkWindow.
type IWindow interface {
	toWindow() *C.GtkWindow
}

// native returns a pointer to the underlying GtkWindow.
func (v *Window) native() *C.GtkWindow {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toGtkWindow(ptr)
}

func (v *Window) toWindow() *C.GtkWindow {
	if v == nil {
		return nil
	}
	return v.native()
}

func marshalWindow(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapWindow(obj), nil
}

func wrapWindow(obj *glib.Object) *Window {
	return &Window{Bin{Container{Widget{glib.InitiallyUnowned{obj}}}}}
}

// WindowNew is a wrapper around gtk_window_new().
func WindowNew(t WindowType) (*Window, error) {
	c := C.gtk_window_new(C.GtkWindowType(t))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapWindow(obj), nil
}

// SetTitle is a wrapper around gtk_window_set_title().
func (v *Window) SetTitle(title string) {
	cstr := C.CString(title)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_window_set_title(v.native(), (*C.gchar)(cstr))
}

// SetResizable is a wrapper around gtk_window_set_resizable().
func (v *Window) SetResizable(resizable bool) {
	C.gtk_window_set_resizable(v.native(), gbool(resizable))
}

// GetResizable is a wrapper around gtk_window_get_resizable().
func (v *Window) GetResizable() bool {
	c := C.gtk_window_get_resizable(v.native())
	return gobool(c)
}

// ActivateFocus is a wrapper around gtk_window_activate_focus().
func (v *Window) ActivateFocus() bool {
	c := C.gtk_window_activate_focus(v.native())
	return gobool(c)
}

// ActivateDefault is a wrapper around gtk_window_activate_default().
func (v *Window) ActivateDefault() bool {
	c := C.gtk_window_activate_default(v.native())
	return gobool(c)
}

// SetModal is a wrapper around gtk_window_set_modal().
func (v *Window) SetModal(modal bool) {
	C.gtk_window_set_modal(v.native(), gbool(modal))
}

// SetDefaultSize is a wrapper around gtk_window_set_default_size().
func (v *Window) SetDefaultSize(width, height int) {
	C.gtk_window_set_default_size(v.native(), C.gint(width), C.gint(height))
}

// GetScreen is a wrapper around gtk_window_get_screen().
func (v *Window) GetScreen() (*gdk.Screen, error) {
	c := C.gtk_window_get_screen(v.native())
	if c == nil {
		return nil, nilPtrErr
	}

	s := &gdk.Screen{glib.Take(unsafe.Pointer(c))}
	return s, nil
}

// SetIcon is a wrapper around gtk_window_set_icon().
func (v *Window) SetIcon(icon *gdk.Pixbuf) {
	iconPtr := (*C.GdkPixbuf)(unsafe.Pointer(icon.Native()))
	C.gtk_window_set_icon(v.native(), iconPtr)
}

// WindowSetDefaultIcon is a wrapper around gtk_window_set_default_icon().
func WindowSetDefaultIcon(icon *gdk.Pixbuf) {
	iconPtr := (*C.GdkPixbuf)(unsafe.Pointer(icon.Native()))
	C.gtk_window_set_default_icon(iconPtr)
}

// TODO(jrick) GdkGeometry GdkWindowHints.
/*
func (v *Window) SetGeometryHints() {
}
*/

// SetGravity is a wrapper around gtk_window_set_gravity().
func (v *Window) SetGravity(gravity gdk.GdkGravity) {
	C.gtk_window_set_gravity(v.native(), C.GdkGravity(gravity))
}

// TODO(jrick) GdkGravity.
/*
func (v *Window) GetGravity() {
}
*/

// SetPosition is a wrapper around gtk_window_set_position().
func (v *Window) SetPosition(position WindowPosition) {
	C.gtk_window_set_position(v.native(), C.GtkWindowPosition(position))
}

// SetTransientFor is a wrapper around gtk_window_set_transient_for().
func (v *Window) SetTransientFor(parent IWindow) {
	var pw *C.GtkWindow = nil
	if parent != nil {
		pw = parent.toWindow()
	}
	C.gtk_window_set_transient_for(v.native(), pw)
}

// SetDestroyWithParent is a wrapper around
// gtk_window_set_destroy_with_parent().
func (v *Window) SetDestroyWithParent(setting bool) {
	C.gtk_window_set_destroy_with_parent(v.native(), gbool(setting))
}

// SetHideTitlebarWhenMaximized is a wrapper around
// gtk_window_set_hide_titlebar_when_maximized().
func (v *Window) SetHideTitlebarWhenMaximized(setting bool) {
	C.gtk_window_set_hide_titlebar_when_maximized(v.native(),
		gbool(setting))
}

// IsActive is a wrapper around gtk_window_is_active().
func (v *Window) IsActive() bool {
	c := C.gtk_window_is_active(v.native())
	return gobool(c)
}

// HasToplevelFocus is a wrapper around gtk_window_has_toplevel_focus().
func (v *Window) HasToplevelFocus() bool {
	c := C.gtk_window_has_toplevel_focus(v.native())
	return gobool(c)
}

// GetFocus is a wrapper around gtk_window_get_focus().
func (v *Window) GetFocus() (*Widget, error) {
	c := C.gtk_window_get_focus(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapWidget(obj), nil
}

// SetFocus is a wrapper around gtk_window_set_focus().
func (v *Window) SetFocus(widget IWidget) {
	C.gtk_window_set_focus(v.native(), widget.toWidget())
}

// GetDefaultWidget is a wrapper arround gtk_window_get_default_widget().
func (v *Window) GetDefaultWidget() *Widget {
	c := C.gtk_window_get_default_widget(v.native())
	if c == nil {
		return nil
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapWidget(obj)
}

// SetDefault is a wrapper arround gtk_window_set_default().
func (v *Window) SetDefault(widget IWidget) {
	C.gtk_window_set_default(v.native(), widget.toWidget())
}

// Present is a wrapper around gtk_window_present().
func (v *Window) Present() {
	C.gtk_window_present(v.native())
}

// PresentWithTime is a wrapper around gtk_window_present_with_time().
func (v *Window) PresentWithTime(ts uint32) {
	C.gtk_window_present_with_time(v.native(), C.guint32(ts))
}

// Iconify is a wrapper around gtk_window_iconify().
func (v *Window) Iconify() {
	C.gtk_window_iconify(v.native())
}

// Deiconify is a wrapper around gtk_window_deiconify().
func (v *Window) Deiconify() {
	C.gtk_window_deiconify(v.native())
}

// Stick is a wrapper around gtk_window_stick().
func (v *Window) Stick() {
	C.gtk_window_stick(v.native())
}

// Unstick is a wrapper around gtk_window_unstick().
func (v *Window) Unstick() {
	C.gtk_window_unstick(v.native())
}

// Maximize is a wrapper around gtk_window_maximize().
func (v *Window) Maximize() {
	C.gtk_window_maximize(v.native())
}

// Unmaximize is a wrapper around gtk_window_unmaximize().
func (v *Window) Unmaximize() {
	C.gtk_window_unmaximize(v.native())
}

// Fullscreen is a wrapper around gtk_window_fullscreen().
func (v *Window) Fullscreen() {
	C.gtk_window_fullscreen(v.native())
}

// Unfullscreen is a wrapper around gtk_window_unfullscreen().
func (v *Window) Unfullscreen() {
	C.gtk_window_unfullscreen(v.native())
}

// SetKeepAbove is a wrapper around gtk_window_set_keep_above().
func (v *Window) SetKeepAbove(setting bool) {
	C.gtk_window_set_keep_above(v.native(), gbool(setting))
}

// SetKeepBelow is a wrapper around gtk_window_set_keep_below().
func (v *Window) SetKeepBelow(setting bool) {
	C.gtk_window_set_keep_below(v.native(), gbool(setting))
}

// SetDecorated is a wrapper around gtk_window_set_decorated().
func (v *Window) SetDecorated(setting bool) {
	C.gtk_window_set_decorated(v.native(), gbool(setting))
}

// SetDeletable is a wrapper around gtk_window_set_deletable().
func (v *Window) SetDeletable(setting bool) {
	C.gtk_window_set_deletable(v.native(), gbool(setting))
}

// SetTypeHint is a wrapper around gtk_window_set_type_hint().
func (v *Window) SetTypeHint(typeHint gdk.WindowTypeHint) {
	C.gtk_window_set_type_hint(v.native(), C.GdkWindowTypeHint(typeHint))
}

// SetSkipTaskbarHint is a wrapper around gtk_window_set_skip_taskbar_hint().
func (v *Window) SetSkipTaskbarHint(setting bool) {
	C.gtk_window_set_skip_taskbar_hint(v.native(), gbool(setting))
}

// SetSkipPagerHint is a wrapper around gtk_window_set_skip_pager_hint().
func (v *Window) SetSkipPagerHint(setting bool) {
	C.gtk_window_set_skip_pager_hint(v.native(), gbool(setting))
}

// SetUrgencyHint is a wrapper around gtk_window_set_urgency_hint().
func (v *Window) SetUrgencyHint(setting bool) {
	C.gtk_window_set_urgency_hint(v.native(), gbool(setting))
}

// SetAcceptFocus is a wrapper around gtk_window_set_accept_focus().
func (v *Window) SetAcceptFocus(setting bool) {
	C.gtk_window_set_accept_focus(v.native(), gbool(setting))
}

// SetFocusOnMap is a wrapper around gtk_window_set_focus_on_map().
func (v *Window) SetFocusOnMap(setting bool) {
	C.gtk_window_set_focus_on_map(v.native(), gbool(setting))
}

// SetStartupID is a wrapper around gtk_window_set_startup_id().
func (v *Window) SetStartupID(sid string) {
	cstr := C.CString(sid)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_window_set_startup_id(v.native(), (*C.gchar)(cstr))
}

// SetRole is a wrapper around gtk_window_set_role().
func (v *Window) SetRole(s string) {
	cstr := C.CString(s)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_window_set_role(v.native(), (*C.gchar)(cstr))
}

// GetDecorated is a wrapper around gtk_window_get_decorated().
func (v *Window) GetDecorated() bool {
	c := C.gtk_window_get_decorated(v.native())
	return gobool(c)
}

// GetDeletable is a wrapper around gtk_window_get_deletable().
func (v *Window) GetDeletable() bool {
	c := C.gtk_window_get_deletable(v.native())
	return gobool(c)
}

// WindowGetDefaultIconName is a wrapper around gtk_window_get_default_icon_name().
func WindowGetDefaultIconName() (string, error) {
	return stringReturn(C.gtk_window_get_default_icon_name())
}

// GetDefaultSize is a wrapper around gtk_window_get_default_size().
func (v *Window) GetDefaultSize() (width, height int) {
	var w, h C.gint
	C.gtk_window_get_default_size(v.native(), &w, &h)
	return int(w), int(h)
}

// GetDestroyWithParent is a wrapper around
// gtk_window_get_destroy_with_parent().
func (v *Window) GetDestroyWithParent() bool {
	c := C.gtk_window_get_destroy_with_parent(v.native())
	return gobool(c)
}

// GetHideTitlebarWhenMaximized is a wrapper around
// gtk_window_get_hide_titlebar_when_maximized().
func (v *Window) GetHideTitlebarWhenMaximized() bool {
	c := C.gtk_window_get_hide_titlebar_when_maximized(v.native())
	return gobool(c)
}

// GetIcon is a wrapper around gtk_window_get_icon().
func (v *Window) GetIcon() (*gdk.Pixbuf, error) {
	c := C.gtk_window_get_icon(v.native())
	if c == nil {
		return nil, nilPtrErr
	}

	p := &gdk.Pixbuf{glib.Take(unsafe.Pointer(c))}
	return p, nil
}

// GetIconName is a wrapper around gtk_window_get_icon_name().
func (v *Window) GetIconName() (string, error) {
	return stringReturn(C.gtk_window_get_icon_name(v.native()))
}

// GetModal is a wrapper around gtk_window_get_modal().
func (v *Window) GetModal() bool {
	c := C.gtk_window_get_modal(v.native())
	return gobool(c)
}

// GetPosition is a wrapper around gtk_window_get_position().
func (v *Window) GetPosition() (root_x, root_y int) {
	var x, y C.gint
	C.gtk_window_get_position(v.native(), &x, &y)
	return int(x), int(y)
}

func stringReturn(c *C.gchar) (string, error) {
	if c == nil {
		return "", nilPtrErr
	}
	return goString(c), nil
}

// GetRole is a wrapper around gtk_window_get_role().
func (v *Window) GetRole() (string, error) {
	return stringReturn(C.gtk_window_get_role(v.native()))
}

// GetSize is a wrapper around gtk_window_get_size().
func (v *Window) GetSize() (width, height int) {
	var w, h C.gint
	C.gtk_window_get_size(v.native(), &w, &h)
	return int(w), int(h)
}

// GetTitle is a wrapper around gtk_window_get_title().
func (v *Window) GetTitle() (string, error) {
	return stringReturn(C.gtk_window_get_title(v.native()))
}

// GetTransientFor is a wrapper around gtk_window_get_transient_for().
func (v *Window) GetTransientFor() (*Window, error) {
	c := C.gtk_window_get_transient_for(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapWindow(obj), nil
}

// GetAttachedTo is a wrapper around gtk_window_get_attached_to().
func (v *Window) GetAttachedTo() (*Widget, error) {
	c := C.gtk_window_get_attached_to(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapWidget(obj), nil
}

// GetTypeHint is a wrapper around gtk_window_get_type_hint().
func (v *Window) GetTypeHint() gdk.WindowTypeHint {
	c := C.gtk_window_get_type_hint(v.native())
	return gdk.WindowTypeHint(c)
}

// GetSkipTaskbarHint is a wrapper around gtk_window_get_skip_taskbar_hint().
func (v *Window) GetSkipTaskbarHint() bool {
	c := C.gtk_window_get_skip_taskbar_hint(v.native())
	return gobool(c)
}

// GetSkipPagerHint is a wrapper around gtk_window_get_skip_pager_hint().
func (v *Window) GetSkipPagerHint() bool {
	c := C.gtk_window_get_skip_taskbar_hint(v.native())
	return gobool(c)
}

// GetUrgencyHint is a wrapper around gtk_window_get_urgency_hint().
func (v *Window) GetUrgencyHint() bool {
	c := C.gtk_window_get_urgency_hint(v.native())
	return gobool(c)
}

// GetAcceptFocus is a wrapper around gtk_window_get_accept_focus().
func (v *Window) GetAcceptFocus() bool {
	c := C.gtk_window_get_accept_focus(v.native())
	return gobool(c)
}

// GetFocusOnMap is a wrapper around gtk_window_get_focus_on_map().
func (v *Window) GetFocusOnMap() bool {
	c := C.gtk_window_get_focus_on_map(v.native())
	return gobool(c)
}

// HasGroup is a wrapper around gtk_window_has_group().
func (v *Window) HasGroup() bool {
	c := C.gtk_window_has_group(v.native())
	return gobool(c)
}

// Move is a wrapper around gtk_window_move().
func (v *Window) Move(x, y int) {
	C.gtk_window_move(v.native(), C.gint(x), C.gint(y))
}

// Resize is a wrapper around gtk_window_resize().
func (v *Window) Resize(width, height int) {
	C.gtk_window_resize(v.native(), C.gint(width), C.gint(height))
}

// WindowSetDefaultIconFromFile is a wrapper around gtk_window_set_default_icon_from_file().
func WindowSetDefaultIconFromFile(file string) error {
	cstr := C.CString(file)
	defer C.free(unsafe.Pointer(cstr))
	var err *C.GError = nil
	res := C.gtk_window_set_default_icon_from_file((*C.gchar)(cstr), &err)
	if res == 0 {
		defer C.g_error_free(err)
		return errors.New(goString(err.message))
	}
	return nil
}

// WindowSetDefaultIconName is a wrapper around gtk_window_set_default_icon_name().
func WindowSetDefaultIconName(s string) {
	cstr := C.CString(s)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_window_set_default_icon_name((*C.gchar)(cstr))
}

// SetIconFromFile is a wrapper around gtk_window_set_icon_from_file().
func (v *Window) SetIconFromFile(file string) error {
	cstr := C.CString(file)
	defer C.free(unsafe.Pointer(cstr))
	var err *C.GError = nil
	res := C.gtk_window_set_icon_from_file(v.native(), (*C.gchar)(cstr), &err)
	if res == 0 {
		defer C.g_error_free(err)
		return errors.New(goString(err.message))
	}
	return nil
}

// SetIconName is a wrapper around gtk_window_set_icon_name().
func (v *Window) SetIconName(name string) {
	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_window_set_icon_name(v.native(), (*C.gchar)(cstr))
}

// SetAutoStartupNotification is a wrapper around
// gtk_window_set_auto_startup_notification().
// This doesn't seem write.  Might need to rethink?
/*
func (v *Window) SetAutoStartupNotification(setting bool) {
	C.gtk_window_set_auto_startup_notification(gbool(setting))
}
*/

// GetMnemonicsVisible is a wrapper around
// gtk_window_get_mnemonics_visible().
func (v *Window) GetMnemonicsVisible() bool {
	c := C.gtk_window_get_mnemonics_visible(v.native())
	return gobool(c)
}

// SetMnemonicsVisible is a wrapper around
// gtk_window_get_mnemonics_visible().
func (v *Window) SetMnemonicsVisible(setting bool) {
	C.gtk_window_set_mnemonics_visible(v.native(), gbool(setting))
}

// GetFocusVisible is a wrapper around gtk_window_get_focus_visible().
func (v *Window) GetFocusVisible() bool {
	c := C.gtk_window_get_focus_visible(v.native())
	return gobool(c)
}

// SetFocusVisible is a wrapper around gtk_window_set_focus_visible().
func (v *Window) SetFocusVisible(setting bool) {
	C.gtk_window_set_focus_visible(v.native(), gbool(setting))
}

// GetApplication is a wrapper around gtk_window_get_application().
func (v *Window) GetApplication() (*Application, error) {
	c := C.gtk_window_get_application(v.native())
	if c == nil {
		return nil, nilPtrErr
	}

	obj := glib.Take(unsafe.Pointer(c))
	return wrapApplication(obj), nil
}

// SetApplication is a wrapper around gtk_window_set_application().
func (v *Window) SetApplication(a *Application) {
	C.gtk_window_set_application(v.native(), a.native())
}

// ActivateKey is a wrapper around gtk_window_activate_key().
func (v *Window) ActivateKey(event *gdk.EventKey) bool {
	c := C.gtk_window_activate_key(v.native(), (*C.GdkEventKey)(unsafe.Pointer(event.Native())))
	return gobool(c)
}

// AddMnemonic is a wrapper around gtk_window_add_mnemonic().
func (v *Window) AddMnemonic(keyval uint, target *Widget) {
	C.gtk_window_add_mnemonic(v.native(), C.guint(keyval), target.native())
}

// RemoveMnemonic is a wrapper around gtk_window_remove_mnemonic().
func (v *Window) RemoveMnemonic(keyval uint, target *Widget) {
	C.gtk_window_remove_mnemonic(v.native(), C.guint(keyval), target.native())
}

// ActivateMnemonic is a wrapper around gtk_window_mnemonic_activate().
func (v *Window) ActivateMnemonic(keyval uint, mods gdk.ModifierType) bool {
	c := C.gtk_window_mnemonic_activate(v.native(), C.guint(keyval), C.GdkModifierType(mods))
	return gobool(c)
}

// GetMnemonicModifier is a wrapper around gtk_window_get_mnemonic_modifier().
func (v *Window) GetMnemonicModifier() gdk.ModifierType {
	c := C.gtk_window_get_mnemonic_modifier(v.native())
	return gdk.ModifierType(c)
}

// SetMnemonicModifier is a wrapper around gtk_window_set_mnemonic_modifier().
func (v *Window) SetMnemonicModifier(mods gdk.ModifierType) {
	C.gtk_window_set_mnemonic_modifier(v.native(), C.GdkModifierType(mods))
}

// TODO gtk_window_begin_move_drag().
// TODO gtk_window_begin_resize_drag().
// TODO gtk_window_get_default_icon_list().
// TODO gtk_window_get_group().
// TODO gtk_window_get_icon_list().
// TODO gtk_window_get_window_type().
// TODO gtk_window_list_toplevels().
// TODO gtk_window_parse_geometry().
// TODO gtk_window_propogate_key_event().
// TODO gtk_window_set_attached_to().
// TODO gtk_window_set_default_icon_list().
// TODO gtk_window_set_icon_list().
// TODO gtk_window_set_screen().
// TODO gtk_window_get_resize_grip_area().

/*
 * GtkAssistant
 */

// Assistant is a representation of GTK's GtkAssistant.
type Assistant struct {
	Window
}

// native returns a pointer to the underlying GtkAssistant.
func (v *Assistant) native() *C.GtkAssistant {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toGtkAssistant(ptr)
}

func marshalAssistant(p uintptr) (interface{}, error) {
	c := C.g_value_get_object(C.toGValue(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapAssistant(obj), nil
}

func wrapAssistant(obj *glib.Object) *Assistant {
	window := wrapWindow(obj)
	return &Assistant{*window}
}

// AssistantNew is a wrapper around gtk_assistant_new().
func AssistantNew() (*Assistant, error) {
	c := C.gtk_assistant_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapAssistant(obj), nil
}

// GetCurrentPage is a wrapper around gtk_assistant_get_current_page().
func (v *Assistant) GetCurrentPage() int {
	c := C.gtk_assistant_get_current_page(v.native())
	return int(c)
}

// SetCurrentPage is a wrapper around gtk_assistant_set_current_page().
func (v *Assistant) SetCurrentPage(pageNum int) {
	C.gtk_assistant_set_current_page(v.native(), C.gint(pageNum))
}

// GetNPages is a wrapper around gtk_assistant_get_n_pages().
func (v *Assistant) GetNPages() int {
	c := C.gtk_assistant_get_n_pages(v.native())
	return int(c)
}

// GetNthPage is a wrapper around gtk_assistant_get_nth_page().
func (v *Assistant) GetNthPage(pageNum int) (*Widget, error) {
	c := C.gtk_assistant_get_nth_page(v.native(), C.gint(pageNum))
	if c == nil {
		return nil, fmt.Errorf("page %d is out of bounds", pageNum)
	}

	obj := glib.Take(unsafe.Pointer(c))
	return wrapWidget(obj), nil
}

// PrependPage is a wrapper around gtk_assistant_prepend_page().
func (v *Assistant) PrependPage(page IWidget) int {
	c := C.gtk_assistant_prepend_page(v.native(), page.toWidget())
	return int(c)
}

// AppendPage is a wrapper around gtk_assistant_append_page().
func (v *Assistant) AppendPage(page IWidget) int {
	c := C.gtk_assistant_append_page(v.native(), page.toWidget())
	return int(c)
}

// InsertPage is a wrapper around gtk_assistant_insert_page().
func (v *Assistant) InsertPage(page IWidget, position int) int {
	c := C.gtk_assistant_insert_page(v.native(), page.toWidget(),
		C.gint(position))
	return int(c)
}

// RemovePage is a wrapper around gtk_assistant_remove_page().
func (v *Assistant) RemovePage(pageNum int) {
	C.gtk_assistant_remove_page(v.native(), C.gint(pageNum))
}

// TODO: gtk_assistant_set_forward_page_func

// SetPageType is a wrapper around gtk_assistant_set_page_type().
func (v *Assistant) SetPageType(page IWidget, ptype AssistantPageType) {
	C.gtk_assistant_set_page_type(v.native(), page.toWidget(),
		C.GtkAssistantPageType(ptype))
}

// GetPageType is a wrapper around gtk_assistant_get_page_type().
func (v *Assistant) GetPageType(page IWidget) AssistantPageType {
	c := C.gtk_assistant_get_page_type(v.native(), page.toWidget())
	return AssistantPageType(c)
}

// SetPageTitle is a wrapper around gtk_assistant_set_page_title().
func (v *Assistant) SetPageTitle(page IWidget, title string) {
	cstr := C.CString(title)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_assistant_set_page_title(v.native(), page.toWidget(),
		(*C.gchar)(cstr))
}

// GetPageTitle is a wrapper around gtk_assistant_get_page_title().
func (v *Assistant) GetPageTitle(page IWidget) string {
	c := C.gtk_assistant_get_page_title(v.native(), page.toWidget())
	return goString(c)
}

// SetPageComplete is a wrapper around gtk_assistant_set_page_complete().
func (v *Assistant) SetPageComplete(page IWidget, complete bool) {
	C.gtk_assistant_set_page_complete(v.native(), page.toWidget(),
		gbool(complete))
}

// GetPageComplete is a wrapper around gtk_assistant_get_page_complete().
func (v *Assistant) GetPageComplete(page IWidget) bool {
	c := C.gtk_assistant_get_page_complete(v.native(), page.toWidget())
	return gobool(c)
}

// AddActionWidget is a wrapper around gtk_assistant_add_action_widget().
func (v *Assistant) AddActionWidget(child IWidget) {
	C.gtk_assistant_add_action_widget(v.native(), child.toWidget())
}

// RemoveActionWidget is a wrapper around gtk_assistant_remove_action_widget().
func (v *Assistant) RemoveActionWidget(child IWidget) {
	C.gtk_assistant_remove_action_widget(v.native(), child.toWidget())
}

// UpdateButtonsState is a wrapper around gtk_assistant_update_buttons_state().
func (v *Assistant) UpdateButtonsState() {
	C.gtk_assistant_update_buttons_state(v.native())
}

// Commit is a wrapper around gtk_assistant_commit().
func (v *Assistant) Commit() {
	C.gtk_assistant_commit(v.native())
}

// NextPage is a wrapper around gtk_assistant_next_page().
func (v *Assistant) NextPage() {
	C.gtk_assistant_next_page(v.native())
}

// PreviousPage is a wrapper around gtk_assistant_previous_page().
func (v *Assistant) PreviousPage() {
	C.gtk_assistant_previous_page(v.native())
}

/*
 * GtkOffscreenWindow
 */

// OffscreenWindow is a representation of GTK's GtkOffscreenWindow.
type OffscreenWindow struct {
	Window
}

// native returns a pointer to the underlying GtkOffscreenWindow.
func (v *OffscreenWindow) native() *C.GtkOffscreenWindow {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toGtkOffscreenWindow(ptr)
}

func marshalOffscreenWindow(p uintptr) (interface{}, error) {
	c := C.g_value_get_object(C.toGValue(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapOffscreenWindow(obj), nil
}

func wrapOffscreenWindow(obj *glib.Object) *OffscreenWindow {
	window := wrapWindow(obj)
	return &OffscreenWindow{*window}
}

// OffscreenWindowNew is a wrapper around gtk_offscreen_window_new().
func OffscreenWindowNew() (*OffscreenWindow, error) {
	c := C.gtk_offscreen_window_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapOffscreenWindow(obj), nil
}

// GetSurface is a wrapper around gtk_offscreen_window_get_surface().
// The returned surface is safe to use over window resizes.
func (v *OffscreenWindow) GetSurface() (*cairo.Surface, error) {
	c := C.gtk_offscreen_window_get_surface(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	cairoPtr := (uintptr)(unsafe.Pointer(c))
	s := cairo.NewSurface(cairoPtr, true)
	return s, nil
}

// GetPixbuf is a wrapper around gtk_offscreen_window_get_pixbuf().
func (v *OffscreenWindow) GetPixbuf() (*gdk.Pixbuf, error) {
	c := C.gtk_offscreen_window_get_pixbuf(v.native())
	if c == nil {
		return nil, nilPtrErr
	}

	// Pixbuf is returned with ref count of 1, so don't increment.
	// Is it a floating reference?
	pb := &gdk.Pixbuf{glib.Take(unsafe.Pointer(c))}
	return pb, nil
}

/*
 * GtkDialog
 */

// Dialog is a representation of GTK's GtkDialog.
type Dialog struct {
	Window
}

// native returns a pointer to the underlying GtkDialog.
func (v *Dialog) native() *C.GtkDialog {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toGtkDialog(ptr)
}

func marshalDialog(p uintptr) (interface{}, error) {
	c := C.g_value_get_object(C.toGValue(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapDialog(obj), nil
}

func wrapDialog(obj *glib.Object) *Dialog {
	window := wrapWindow(obj)
	return &Dialog{*window}
}

// DialogNew() is a wrapper around gtk_dialog_new().
func DialogNew() (*Dialog, error) {
	c := C.gtk_dialog_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapDialog(obj), nil
}

// DialogWithButtonsNew is a wrapper around gtk_dialog_new_with_buttons().
func DialogWithButtonsNew(title string, parent IWindow, flags DialogFlags,
	buttonText string, buttonResponse ResponseType) (*Dialog, error) {
	cstr1 := C.CString(title)
	defer C.free(unsafe.Pointer(cstr1))
	cstr2 := C.CString(buttonText)
	defer C.free(unsafe.Pointer(cstr2))
	c := C._gtk_dialog_new_with_buttons(cstr1, parent.toWindow(),
		(C.GtkDialogFlags)(flags), cstr2, (C.GtkResponseType)(buttonResponse))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapDialog(obj), nil
}

// DialogWithFlagsNew is an alternative wrapper around gtk_dialog_new_with_buttons().
func DialogWithFlagsNew(title string, parent IWindow, flags DialogFlags) (*Dialog, error) {
	cstr1 := C.CString(title)
	defer C.free(unsafe.Pointer(cstr1))
	c := C._gtk_dialog_new_with_buttons(cstr1, parent.toWindow(),
		(C.GtkDialogFlags)(flags), nil, 0)
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapDialog(obj), nil
}

// Run() is a wrapper around gtk_dialog_run().
func (v *Dialog) Run() ResponseType {
	c := C.gtk_dialog_run(v.native())
	return ResponseType(c)
}

// Response() is a wrapper around gtk_dialog_response().
func (v *Dialog) Response(response ResponseType) {
	C.gtk_dialog_response(v.native(), C.gint(response))
}

// AddButton() is a wrapper around gtk_dialog_add_button().  text may
// be either the literal button text, or if using GTK 3.8 or earlier, a
// Stock type converted to a string.
func (v *Dialog) AddButton(text string, id ResponseType) (*Button, error) {
	cstr := C.CString(text)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_dialog_add_button(v.native(), (*C.gchar)(cstr), C.gint(id))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapButton(obj), nil
}

// AddActionWidget() is a wrapper around gtk_dialog_add_action_widget().
func (v *Dialog) AddActionWidget(child IWidget, id ResponseType) {
	C.gtk_dialog_add_action_widget(v.native(), child.toWidget(), C.gint(id))
}

// SetDefaultResponse() is a wrapper around gtk_dialog_set_default_response().
func (v *Dialog) SetDefaultResponse(id ResponseType) {
	C.gtk_dialog_set_default_response(v.native(), C.gint(id))
}

// SetResponseSensitive() is a wrapper around
// gtk_dialog_set_response_sensitive().
func (v *Dialog) SetResponseSensitive(id ResponseType, setting bool) {
	C.gtk_dialog_set_response_sensitive(v.native(), C.gint(id),
		gbool(setting))
}

// GetResponseForWidget() is a wrapper around
// gtk_dialog_get_response_for_widget().
func (v *Dialog) GetResponseForWidget(widget IWidget) ResponseType {
	c := C.gtk_dialog_get_response_for_widget(v.native(), widget.toWidget())
	return ResponseType(c)
}

// GetWidgetForResponse() is a wrapper around
// gtk_dialog_get_widget_for_response().
func (v *Dialog) GetWidgetForResponse(id ResponseType) (*Widget, error) {
	c := C.gtk_dialog_get_widget_for_response(v.native(), C.gint(id))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapWidget(obj), nil
}

// GetContentArea() is a wrapper around gtk_dialog_get_content_area().
func (v *Dialog) GetContentArea() (*Box, error) {
	c := C.gtk_dialog_get_content_area(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	box := wrapBox(obj)
	return box, nil
}

func init() {
	tm := []glib.TypeMarshaler{
		{glib.Type(C.gtk_about_dialog_get_type()), marshalAboutDialog},
	}

	glib.RegisterGValueMarshalers(tm)

	WrapMap["GtkAboutDialog"] = wrapAboutDialog
}

/*
 * GtkAboutDialog
 */

// AboutDialog is a representation of GTK's GtkAboutDialog.
type AboutDialog struct {
	Dialog
}

// native returns a pointer to the underlying GtkAboutDialog.
func (v *AboutDialog) native() *C.GtkAboutDialog {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toGtkAboutDialog(ptr)
}

func marshalAboutDialog(p uintptr) (interface{}, error) {
	c := C.g_value_get_object(C.toGValue(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapAboutDialog(obj), nil
}

func wrapAboutDialog(obj *glib.Object) *AboutDialog {
	return &AboutDialog{Dialog{Window{Bin{Container{Widget{glib.InitiallyUnowned{obj}}}}}}}
}

// AboutDialogNew is a wrapper around gtk_about_dialog_new().
func AboutDialogNew() (*AboutDialog, error) {
	c := C.gtk_about_dialog_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapAboutDialog(obj), nil
}

// GetComments is a wrapper around gtk_about_dialog_get_comments().
func (v *AboutDialog) GetComments() string {
	c := C.gtk_about_dialog_get_comments(v.native())
	return goString(c)
}

// SetComments is a wrapper around gtk_about_dialog_set_comments().
func (v *AboutDialog) SetComments(comments string) {
	cstr := C.CString(comments)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_about_dialog_set_comments(v.native(), (*C.gchar)(cstr))
}

// GetCopyright is a wrapper around gtk_about_dialog_get_copyright().
func (v *AboutDialog) GetCopyright() string {
	c := C.gtk_about_dialog_get_copyright(v.native())
	return goString(c)
}

// SetCopyright is a wrapper around gtk_about_dialog_set_copyright().
func (v *AboutDialog) SetCopyright(copyright string) {
	cstr := C.CString(copyright)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_about_dialog_set_copyright(v.native(), (*C.gchar)(cstr))
}

// GetLicense is a wrapper around gtk_about_dialog_get_license().
func (v *AboutDialog) GetLicense() string {
	c := C.gtk_about_dialog_get_license(v.native())
	return goString(c)
}

// SetLicense is a wrapper around gtk_about_dialog_set_license().
func (v *AboutDialog) SetLicense(license string) {
	cstr := C.CString(license)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_about_dialog_set_license(v.native(), (*C.gchar)(cstr))
}

// GetLicenseType is a wrapper around gtk_about_dialog_get_license_type().
func (v *AboutDialog) GetLicenseType() License {
	c := C.gtk_about_dialog_get_license_type(v.native())
	return License(c)
}

// SetLicenseType is a wrapper around gtk_about_dialog_set_license_type().
func (v *AboutDialog) SetLicenseType(license License) {
	C.gtk_about_dialog_set_license_type(v.native(), C.GtkLicense(license))
}

// GetLogo is a wrapper around gtk_about_dialog_get_logo().
func (v *AboutDialog) GetLogo() (*gdk.Pixbuf, error) {
	c := C.gtk_about_dialog_get_logo(v.native())
	if c == nil {
		return nil, nilPtrErr
	}

	p := &gdk.Pixbuf{glib.Take(unsafe.Pointer(c))}
	return p, nil
}

// SetLogo is a wrapper around gtk_about_dialog_set_logo().
func (v *AboutDialog) SetLogo(logo *gdk.Pixbuf) {
	logoPtr := (*C.GdkPixbuf)(unsafe.Pointer(logo.Native()))
	C.gtk_about_dialog_set_logo(v.native(), logoPtr)
}

// GetLogoIconName is a wrapper around gtk_about_dialog_get_logo_icon_name().
func (v *AboutDialog) GetLogoIconName() string {
	c := C.gtk_about_dialog_get_logo_icon_name(v.native())
	return goString(c)
}

// SetLogoIconName is a wrapper around gtk_about_dialog_set_logo_icon_name().
func (v *AboutDialog) SetLogoIconName(name string) {
	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_about_dialog_set_logo_icon_name(v.native(), (*C.gchar)(cstr))
}

// GetProgramName is a wrapper around gtk_about_dialog_get_program_name().
func (v *AboutDialog) GetProgramName() string {
	c := C.gtk_about_dialog_get_program_name(v.native())
	return goString(c)
}

// SetProgramName is a wrapper around gtk_about_dialog_set_program_name().
func (v *AboutDialog) SetProgramName(name string) {
	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_about_dialog_set_program_name(v.native(), (*C.gchar)(cstr))
}

// GetAuthors is a wrapper around gtk_about_dialog_get_authors().
func (v *AboutDialog) GetAuthors() []string {
	cauthors := C.gtk_about_dialog_get_authors(v.native())
	authors := goStringArray(cauthors)
	return authors
}

// SetAuthors is a wrapper around gtk_about_dialog_set_authors().
func (v *AboutDialog) SetAuthors(authors []string) {
	cauthors := C.make_strings(C.int(len(authors) + 1))
	for i, author := range authors {
		cstr := C.CString(author)
		defer C.free(unsafe.Pointer(cstr))
		C.set_string(cauthors, C.int(i), (*C.gchar)(cstr))
	}

	C.set_string(cauthors, C.int(len(authors)), nil)
	C.gtk_about_dialog_set_authors(v.native(), cauthors)
	C.destroy_strings(cauthors)
}

// GetArtists is a wrapper around gtk_about_dialog_get_artists().
func (v *AboutDialog) GetArtists() []string {
	cartists := C.gtk_about_dialog_get_artists(v.native())
	artists := goStringArray(cartists)
	return artists
}

// SetArtists is a wrapper around gtk_about_dialog_set_artists().
func (v *AboutDialog) SetArtists(artists []string) {
	cartists := C.make_strings(C.int(len(artists) + 1))
	for i, artist := range artists {
		cstr := C.CString(artist)
		defer C.free(unsafe.Pointer(cstr))
		C.set_string(cartists, C.int(i), (*C.gchar)(cstr))
	}

	C.set_string(cartists, C.int(len(artists)), nil)
	C.gtk_about_dialog_set_artists(v.native(), cartists)
	C.destroy_strings(cartists)
}

// GetDocumenters is a wrapper around gtk_about_dialog_get_documenters().
func (v *AboutDialog) GetDocumenters() []string {
	cdocumenters := C.gtk_about_dialog_get_documenters(v.native())
	documenters := goStringArray(cdocumenters)
	return documenters
}

// SetDocumenters is a wrapper around gtk_about_dialog_set_documenters().
func (v *AboutDialog) SetDocumenters(documenters []string) {
	cdocumenters := C.make_strings(C.int(len(documenters) + 1))
	for i, doc := range documenters {
		cstr := C.CString(doc)
		defer C.free(unsafe.Pointer(cstr))
		C.set_string(cdocumenters, C.int(i), (*C.gchar)(cstr))
	}

	C.set_string(cdocumenters, C.int(len(documenters)), nil)
	C.gtk_about_dialog_set_documenters(v.native(), cdocumenters)
	C.destroy_strings(cdocumenters)
}

// GetTranslatorCredits is a wrapper around gtk_about_dialog_get_translator_credits().
func (v *AboutDialog) GetTranslatorCredits() string {
	c := C.gtk_about_dialog_get_translator_credits(v.native())
	return goString(c)
}

// SetTranslatorCredits is a wrapper around gtk_about_dialog_set_translator_credits().
func (v *AboutDialog) SetTranslatorCredits(translatorCredits string) {
	cstr := C.CString(translatorCredits)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_about_dialog_set_translator_credits(v.native(), (*C.gchar)(cstr))
}

// GetVersion is a wrapper around gtk_about_dialog_get_version().
func (v *AboutDialog) GetVersion() string {
	c := C.gtk_about_dialog_get_version(v.native())
	return goString(c)
}

// SetVersion is a wrapper around gtk_about_dialog_set_version().
func (v *AboutDialog) SetVersion(version string) {
	cstr := C.CString(version)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_about_dialog_set_version(v.native(), (*C.gchar)(cstr))
}

// GetWebsite is a wrapper around gtk_about_dialog_get_website().
func (v *AboutDialog) GetWebsite() string {
	c := C.gtk_about_dialog_get_website(v.native())
	return goString(c)
}

// SetWebsite is a wrapper around gtk_about_dialog_set_website().
func (v *AboutDialog) SetWebsite(website string) {
	cstr := C.CString(website)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_about_dialog_set_website(v.native(), (*C.gchar)(cstr))
}

// GetWebsiteLabel is a wrapper around gtk_about_dialog_get_website_label().
func (v *AboutDialog) GetWebsiteLabel() string {
	c := C.gtk_about_dialog_get_website_label(v.native())
	return goString(c)
}

// SetWebsiteLabel is a wrapper around gtk_about_dialog_set_website_label().
func (v *AboutDialog) SetWebsiteLabel(websiteLabel string) {
	cstr := C.CString(websiteLabel)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_about_dialog_set_website_label(v.native(), (*C.gchar)(cstr))
}

// GetWrapLicense is a wrapper around gtk_about_dialog_get_wrap_license().
func (v *AboutDialog) GetWrapLicense() bool {
	return gobool(C.gtk_about_dialog_get_wrap_license(v.native()))
}

// SetWrapLicense is a wrapper around gtk_about_dialog_set_wrap_license().
func (v *AboutDialog) SetWrapLicense(wrapLicense bool) {
	C.gtk_about_dialog_set_wrap_license(v.native(), gbool(wrapLicense))
}

// AddCreditSection is a wrapper around gtk_about_dialog_add_credit_section().
func (v *AboutDialog) AddCreditSection(sectionName string, people []string) {
	cname := C.CString(sectionName)
	defer C.free(unsafe.Pointer(cname))

	cpeople := C.make_strings(C.int(len(people)) + 1)
	defer C.destroy_strings(cpeople)
	for i, p := range people {
		cp := C.CString(p)
		defer C.free(unsafe.Pointer(cp))
		C.set_string(cpeople, C.int(i), (*C.gchar)(cp))
	}
	C.set_string(cpeople, C.int(len(people)), nil)

	C.gtk_about_dialog_add_credit_section(v.native(), (*C.gchar)(cname), cpeople)
}

func init() {
	tm := []glib.TypeMarshaler{
		{glib.Type(C.gtk_app_chooser_get_type()), marshalAppChooser},
		{glib.Type(C.gtk_app_chooser_button_get_type()), marshalAppChooserButton},
		{glib.Type(C.gtk_app_chooser_widget_get_type()), marshalAppChooserWidget},
		{glib.Type(C.gtk_app_chooser_dialog_get_type()), marshalAppChooserDialog},
	}

	glib.RegisterGValueMarshalers(tm)

	WrapMap["GtkAppChooser"] = wrapAppChooser
	WrapMap["GtkAppChooserButton"] = wrapAppChooserButton
	WrapMap["GtkAppChooserWidget"] = wrapAppChooserWidget
	WrapMap["GtkAppChooserDialog"] = wrapAppChooserDialog
}

/*
 * GtkAppChooser
 */

// AppChooser is a representation of GTK's GtkAppChooser GInterface.
type AppChooser struct {
	glib.Interface
}

// IAppChooser is an interface type implemented by all structs
// embedding an AppChooser. It is meant to be used as an argument type
// for wrapper functions that wrap around a C GTK function taking a
// GtkAppChooser.
type IAppChooser interface {
	toAppChooser() *C.GtkAppChooser
}

// native returns a pointer to the underlying GtkAppChooser.
func (v *AppChooser) native() *C.GtkAppChooser {
	return C.toGtkAppChooser(unsafe.Pointer(v.Native()))
}

func marshalAppChooser(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapAppChooser(*glib.InterfaceFromObjectNew(obj)), nil
}

func wrapAppChooser(intf glib.Interface) *AppChooser {
	return &AppChooser{intf}
}

func (v *AppChooser) toAppChooser() *C.GtkAppChooser {
	return v.native()
}

// TODO: Needs gio/GAppInfo implementation first
// gtk_app_chooser_get_app_info ()

// GetContentType is a wrapper around gtk_app_chooser_get_content_type().
func (v *AppChooser) GetContentType() string {
	cstr := C.gtk_app_chooser_get_content_type(v.native())
	defer C.free(unsafe.Pointer(cstr))
	return goString(cstr)
}

// Refresh is a wrapper around gtk_app_chooser_refresh().
func (v *AppChooser) Refresh() {
	C.gtk_app_chooser_refresh(v.native())
}

/*
 * GtkAppChooserButton
 */

// AppChooserButton is a representation of GTK's GtkAppChooserButton.
type AppChooserButton struct {
	ComboBox
	// Interfaces
	AppChooser
}

// native returns a pointer to the underlying GtkAppChooserButton.
func (v *AppChooserButton) native() *C.GtkAppChooserButton {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toGtkAppChooserButton(ptr)
}

func marshalAppChooserButton(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapAppChooserButton(obj), nil
}

func wrapAppChooserButton(obj *glib.Object) *AppChooserButton {
	comboBox := wrapComboBox(obj)
	ac := wrapAppChooser(*glib.InterfaceFromObjectNew(obj))
	return &AppChooserButton{*comboBox, *ac}
}

// AppChooserButtonNew() is a wrapper around gtk_app_chooser_button_new().
func AppChooserButtonNew(content_type string) (*AppChooserButton, error) {
	cstr := C.CString(content_type)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_app_chooser_button_new((*C.gchar)(cstr))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapAppChooserButton(obj), nil
}

// TODO: Needs gio/GIcon implemented first
// gtk_app_chooser_button_append_custom_item ()

// AppendSeparator() is a wrapper around gtk_app_chooser_button_append_separator().
func (v *AppChooserButton) AppendSeparator() {
	C.gtk_app_chooser_button_append_separator(v.native())
}

// SetActiveCustomItem() is a wrapper around gtk_app_chooser_button_set_active_custom_item().
func (v *AppChooserButton) SetActiveCustomItem(name string) {
	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_app_chooser_button_set_active_custom_item(v.native(), (*C.gchar)(cstr))
}

// GetShowDefaultItem() is a wrapper around gtk_app_chooser_button_get_show_default_item().
func (v *AppChooserButton) GetShowDefaultItem() bool {
	return gobool(C.gtk_app_chooser_button_get_show_default_item(v.native()))
}

// SetShowDefaultItem() is a wrapper around gtk_app_chooser_button_set_show_default_item().
func (v *AppChooserButton) SetShowDefaultItem(setting bool) {
	C.gtk_app_chooser_button_set_show_default_item(v.native(), gbool(setting))
}

// GetShowDialogItem() is a wrapper around gtk_app_chooser_button_get_show_dialog_item().
func (v *AppChooserButton) GetShowDialogItem() bool {
	return gobool(C.gtk_app_chooser_button_get_show_dialog_item(v.native()))
}

// SetShowDialogItem() is a wrapper around gtk_app_chooser_button_set_show_dialog_item().
func (v *AppChooserButton) SetShowDialogItem(setting bool) {
	C.gtk_app_chooser_button_set_show_dialog_item(v.native(), gbool(setting))
}

// GetHeading() is a wrapper around gtk_app_chooser_button_get_heading().
// In case when gtk_app_chooser_button_get_heading() returns a nil string,
// GetHeading() returns a non-nil error.
func (v *AppChooserButton) GetHeading() (string, error) {
	cstr := C.gtk_app_chooser_button_get_heading(v.native())
	if cstr == nil {
		return "", nilPtrErr
	}
	defer C.free(unsafe.Pointer(cstr))
	return goString(cstr), nil
}

// SetHeading() is a wrapper around gtk_app_chooser_button_set_heading().
func (v *AppChooserButton) SetHeading(heading string) {
	cstr := C.CString(heading)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_app_chooser_button_set_heading(v.native(), (*C.gchar)(cstr))
}

/*
 * GtkAppChooserWidget
 */

// AppChooserWidget is a representation of GTK's GtkAppChooserWidget.
type AppChooserWidget struct {
	Box
	// Interfaces
	AppChooser
}

// native returns a pointer to the underlying GtkAppChooserWidget.
func (v *AppChooserWidget) native() *C.GtkAppChooserWidget {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toGtkAppChooserWidget(ptr)
}

func marshalAppChooserWidget(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapAppChooserWidget(obj), nil
}

func wrapAppChooserWidget(obj *glib.Object) *AppChooserWidget {
	box := wrapBox(obj)
	ac := wrapAppChooser(*glib.InterfaceFromObjectNew(obj))
	return &AppChooserWidget{*box, *ac}
}

// AppChooserWidgetNew() is a wrapper around gtk_app_chooser_widget_new().
func AppChooserWidgetNew(content_type string) (*AppChooserWidget, error) {
	cstr := C.CString(content_type)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_app_chooser_widget_new((*C.gchar)(cstr))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapAppChooserWidget(obj), nil
}

// GetShowDefault() is a wrapper around gtk_app_chooser_widget_get_show_default().
func (v *AppChooserWidget) GetShowDefault() bool {
	return gobool(C.gtk_app_chooser_widget_get_show_default(v.native()))
}

// SetShowDefault() is a wrapper around gtk_app_chooser_widget_set_show_default().
func (v *AppChooserWidget) SetShowDefault(setting bool) {
	C.gtk_app_chooser_widget_set_show_default(v.native(), gbool(setting))
}

// GetShowRecommended() is a wrapper around gtk_app_chooser_widget_get_show_recommended().
func (v *AppChooserWidget) GetShowRecommended() bool {
	return gobool(C.gtk_app_chooser_widget_get_show_recommended(v.native()))
}

// SetShowRecommended() is a wrapper around gtk_app_chooser_widget_set_show_recommended().
func (v *AppChooserWidget) SetShowRecommended(setting bool) {
	C.gtk_app_chooser_widget_set_show_recommended(v.native(), gbool(setting))
}

// GetShowFallback() is a wrapper around gtk_app_chooser_widget_get_show_fallback().
func (v *AppChooserWidget) GetShowFallback() bool {
	return gobool(C.gtk_app_chooser_widget_get_show_fallback(v.native()))
}

// SetShowFallback() is a wrapper around gtk_app_chooser_widget_set_show_fallback().
func (v *AppChooserWidget) SetShowFallback(setting bool) {
	C.gtk_app_chooser_widget_set_show_fallback(v.native(), gbool(setting))
}

// GetShowOther() is a wrapper around gtk_app_chooser_widget_get_show_other().
func (v *AppChooserWidget) GetShowOther() bool {
	return gobool(C.gtk_app_chooser_widget_get_show_other(v.native()))
}

// SetShowOther() is a wrapper around gtk_app_chooser_widget_set_show_other().
func (v *AppChooserWidget) SetShowOther(setting bool) {
	C.gtk_app_chooser_widget_set_show_other(v.native(), gbool(setting))
}

// GetShowAll() is a wrapper around gtk_app_chooser_widget_get_show_all().
func (v *AppChooserWidget) GetShowAll() bool {
	return gobool(C.gtk_app_chooser_widget_get_show_all(v.native()))
}

// SetShowAll() is a wrapper around gtk_app_chooser_widget_set_show_all().
func (v *AppChooserWidget) SetShowAll(setting bool) {
	C.gtk_app_chooser_widget_set_show_all(v.native(), gbool(setting))
}

// GetDefaultText() is a wrapper around gtk_app_chooser_widget_get_default_text().
// In case when gtk_app_chooser_widget_get_default_text() returns a nil string,
// GetDefaultText() returns a non-nil error.
func (v *AppChooserWidget) GetDefaultText() (string, error) {
	cstr := C.gtk_app_chooser_widget_get_default_text(v.native())
	if cstr == nil {
		return "", nilPtrErr
	}
	defer C.free(unsafe.Pointer(cstr))
	return goString(cstr), nil
}

// SetDefaultText() is a wrapper around gtk_app_chooser_widget_set_default_text().
func (v *AppChooserWidget) SetDefaultText(text string) {
	cstr := C.CString(text)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_app_chooser_widget_set_default_text(v.native(), (*C.gchar)(cstr))
}

/*
 * GtkAppChooserDialog
 */

// AppChooserDialog is a representation of GTK's GtkAppChooserDialog.
type AppChooserDialog struct {
	Dialog
	// Interfaces
	AppChooser
}

// native returns a pointer to the underlying GtkAppChooserButton.
func (v *AppChooserDialog) native() *C.GtkAppChooserDialog {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toGtkAppChooserDialog(ptr)
}

func marshalAppChooserDialog(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapAppChooserDialog(obj), nil
}

func wrapAppChooserDialog(obj *glib.Object) *AppChooserDialog {
	dialog := wrapDialog(obj)
	ac := wrapAppChooser(*glib.InterfaceFromObjectNew(obj))
	return &AppChooserDialog{*dialog, *ac}
}

// AppChooserDialogNew() is a wrapper around gtk_app_chooser_dialog_new().
func AppChooserDialogNew(parent *Window, flags DialogFlags, file *glib.File) (*AppChooserDialog, error) {
	c := C.gtk_app_chooser_dialog_new(parent.native(), C.GtkDialogFlags(flags),
		C.toGFile(unsafe.Pointer(file.Native())))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapAppChooserDialog(obj), nil
}

// AppChooserDialogNewForContentType() is a wrapper around gtk_app_chooser_dialog_new_for_content_type().
func AppChooserDialogNewForContentType(parent *Window, flags DialogFlags, content_type string) (*AppChooserDialog, error) {
	cstr := C.CString(content_type)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_app_chooser_dialog_new_for_content_type(parent.native(), C.GtkDialogFlags(flags), (*C.gchar)(cstr))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapAppChooserDialog(obj), nil
}

// GetWidget() is a wrapper around gtk_app_chooser_dialog_get_widget().
func (v *AppChooserDialog) GetWidget() *AppChooserWidget {
	c := C.gtk_app_chooser_dialog_get_widget(v.native())
	obj := glib.Take(unsafe.Pointer(c))
	return wrapAppChooserWidget(obj)
}

// GetHeading() is a wrapper around gtk_app_chooser_dialog_get_heading().
// In case when gtk_app_chooser_dialog_get_heading() returns a nil string,
// GetHeading() returns a non-nil error.
func (v *AppChooserDialog) GetHeading() (string, error) {
	cstr := C.gtk_app_chooser_dialog_get_heading(v.native())
	if cstr == nil {
		return "", nilPtrErr
	}
	defer C.free(unsafe.Pointer(cstr))
	return goString(cstr), nil
}

// SetHeading() is a wrapper around gtk_app_chooser_dialog_set_heading().
func (v *AppChooserDialog) SetHeading(heading string) {
	cstr := C.CString(heading)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_app_chooser_dialog_set_heading(v.native(), (*C.gchar)(cstr))
}

func init() {
	tm := []glib.TypeMarshaler{
		{glib.Type(C.gtk_color_chooser_get_type()), marshalColorChooser},
		{glib.Type(C.gtk_color_chooser_dialog_get_type()), marshalColorChooserDialog},
	}

	glib.RegisterGValueMarshalers(tm)

	WrapMap["GtkColorChooser"] = wrapColorChooser
	WrapMap["GtkColorChooserDialog"] = wrapColorChooserDialog
}

/*
 * GtkColorChooser
 */

// ColorChooser is a representation of GTK's GtkColorChooser GInterface.
type ColorChooser struct {
	*glib.Object
}

// IColorChooser is an interface type implemented by all structs
// embedding an ColorChooser. It is meant to be used as an argument type
// for wrapper functions that wrap around a C GTK function taking a
// GtkColorChooser.
type IColorChooser interface {
	toColorChooser() *C.GtkColorChooser
}

// native returns a pointer to the underlying GtkAppChooser.
func (v *ColorChooser) native() *C.GtkColorChooser {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toGtkColorChooser(ptr)
}

func marshalColorChooser(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapColorChooser(obj), nil
}

func wrapColorChooser(obj *glib.Object) *ColorChooser {
	return &ColorChooser{obj}
}

func (v *ColorChooser) toColorChooser() *C.GtkColorChooser {
	if v == nil {
		return nil
	}
	return v.native()
}

// GetRGBA() is a wrapper around gtk_color_chooser_get_rgba().
func (v *ColorChooser) GetRGBA() *gdk.RGBA {
	gdkColor := gdk.NewRGBA()
	C.gtk_color_chooser_get_rgba(v.native(), (*C.GdkRGBA)(unsafe.Pointer(gdkColor.Native())))
	return gdkColor
}

// SetRGBA() is a wrapper around gtk_color_chooser_set_rgba().
func (v *ColorChooser) SetRGBA(gdkColor *gdk.RGBA) {
	C.gtk_color_chooser_set_rgba(v.native(), (*C.GdkRGBA)(unsafe.Pointer(gdkColor.Native())))
}

// GetUseAlpha() is a wrapper around gtk_color_chooser_get_use_alpha().
func (v *ColorChooser) GetUseAlpha() bool {
	return gobool(C.gtk_color_chooser_get_use_alpha(v.native()))
}

// SetUseAlpha() is a wrapper around gtk_color_chooser_set_use_alpha().
func (v *ColorChooser) SetUseAlpha(use_alpha bool) {
	C.gtk_color_chooser_set_use_alpha(v.native(), gbool(use_alpha))
}

// AddPalette() is a wrapper around gtk_color_chooser_add_palette().
func (v *ColorChooser) AddPalette(orientation Orientation, colors_per_line int, colors []*gdk.RGBA) {
	n_colors := len(colors)
	var c_colors []C.GdkRGBA
	for _, c := range colors {
		c_colors = append(c_colors, *(*C.GdkRGBA)(unsafe.Pointer(c.Native())))
	}
	C.gtk_color_chooser_add_palette(
		v.native(),
		C.GtkOrientation(orientation),
		C.gint(colors_per_line),
		C.gint(n_colors),
		&c_colors[0],
	)
}

/*
 * GtkColorChooserDialog
 */

// ColorChooserDialog is a representation of GTK's GtkColorChooserDialog.
type ColorChooserDialog struct {
	Dialog
	// Interfaces
	ColorChooser
}

// native returns a pointer to the underlying GtkColorChooserButton.
func (v *ColorChooserDialog) native() *C.GtkColorChooserDialog {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toGtkColorChooserDialog(ptr)
}

func marshalColorChooserDialog(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapColorChooserDialog(obj), nil
}

func wrapColorChooserDialog(obj *glib.Object) *ColorChooserDialog {
	dialog := wrapDialog(obj)
	cc := wrapColorChooser(obj)
	return &ColorChooserDialog{*dialog, *cc}
}

// ColorChooserDialogNew() is a wrapper around gtk_color_chooser_dialog_new().
func ColorChooserDialogNew(title string, parent *Window) (*ColorChooserDialog, error) {
	cstr := C.CString(title)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_color_chooser_dialog_new((*C.gchar)(cstr), parent.native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapColorChooserDialog(obj), nil
}

/*
 * GtkColorButton
 */

// ColorButton is a representation of GTK's GtkColorButton.
type ColorButton struct {
	Button
	// Interfaces
	ColorChooser
}

// Native returns a pointer to the underlying GtkColorButton.
func (v *ColorButton) native() *C.GtkColorButton {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toGtkColorButton(ptr)
}

func wrapColorButton(obj *glib.Object) *ColorButton {
	button := wrapButton(obj)
	cc := wrapColorChooser(obj)
	return &ColorButton{*button, *cc}
}

// ColorButtonNew is a wrapper around gtk_color_button_new().
func ColorButtonNew() (*ColorButton, error) {
	c := C.gtk_color_button_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapColorButton(obj), nil
}

// ColorButtonNewWithRGBA is a wrapper around gtk_color_button_new_with_rgba().
func ColorButtonNewWithRGBA(gdkColor *gdk.RGBA) (*ColorButton, error) {
	c := C.gtk_color_button_new_with_rgba((*C.GdkRGBA)(unsafe.Pointer(gdkColor.Native())))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapColorButton(obj), nil
}

/*
 * GtkMessageDialog
 */

// MessageDialog is a representation of GTK's GtkMessageDialog.
type MessageDialog struct {
	Dialog
}

// native returns a pointer to the underlying GtkMessageDialog.
func (v *MessageDialog) native() *C.GtkMessageDialog {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toGtkMessageDialog(ptr)
}

func marshalMessageDialog(p uintptr) (interface{}, error) {
	c := C.g_value_get_object(C.toGValue(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapMessageDialog(obj), nil
}

func wrapMessageDialog(obj *glib.Object) *MessageDialog {
	dlg := wrapDialog(obj)
	return &MessageDialog{*dlg}
}

// MessageDialogNew() is a wrapper around gtk_message_dialog_new().
// The text is created and formatted by the format specifier and any
// additional arguments.
func MessageDialogNew(parent IWindow, flags DialogFlags, mType MessageType, buttons ButtonsType,
	format *string, a ...interface{}) (*MessageDialog, error) {
	var cstr *C.char
	if format != nil {
		s := fmt.Sprintf(*format, a...)
		cstr = C.CString(s)
		defer C.free(unsafe.Pointer(cstr))
	}
	var w *C.GtkWindow = nil
	if parent != nil {
		w = parent.toWindow()
	}
	c := C._gtk_message_dialog_new(w, C.GtkDialogFlags(flags), C.GtkMessageType(mType),
		C.GtkButtonsType(buttons), cstr)
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapMessageDialog(obj), nil
}

// MessageDialogNewWithMarkup is a wrapper around
// gtk_message_dialog_new_with_markup().
func MessageDialogNewWithMarkup(parent IWindow, flags DialogFlags, mType MessageType, buttons ButtonsType,
	format string, a ...interface{}) (*MessageDialog, error) {
	s := fmt.Sprintf(format, a...)
	cstr := C.CString(s)
	defer C.free(unsafe.Pointer(cstr))
	var w *C.GtkWindow = nil
	if parent != nil {
		w = parent.toWindow()
	}
	c := C._gtk_message_dialog_new_with_markup(w,
		C.GtkDialogFlags(flags), C.GtkMessageType(mType),
		C.GtkButtonsType(buttons), cstr)
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapMessageDialog(obj), nil
}

// SetMarkup is a wrapper around gtk_message_dialog_set_markup().
func (v *MessageDialog) SetMarkup(str string) {
	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_message_dialog_set_markup(v.native(), (*C.gchar)(cstr))
}

// FormatSecondaryText is a wrapper around
// gtk_message_dialog_format_secondary_text().
func (v *MessageDialog) FormatSecondaryText(format string, a ...interface{}) {
	s := fmt.Sprintf(format, a...)
	cstr := C.CString(s)
	defer C.free(unsafe.Pointer(cstr))
	C._gtk_message_dialog_format_secondary_text(v.native(),
		(*C.gchar)(cstr))
}

// FormatSecondaryMarkup is a wrapper around
// gtk_message_dialog_format_secondary_text().
func (v *MessageDialog) FormatSecondaryMarkup(format string, a ...interface{}) {
	s := fmt.Sprintf(format, a...)
	cstr := C.CString(s)
	defer C.free(unsafe.Pointer(cstr))
	C._gtk_message_dialog_format_secondary_markup(v.native(),
		(*C.gchar)(cstr))

}

// Wrap around gtk_message_dialog_get_message_area()
func (v *MessageDialog) GetMessageArea() (*Box, error) {
	c := C.gtk_message_dialog_get_message_area(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapBox(obj), nil
}

/*
 * GtkFileChooser
 */

// FileChoser is a representation of GTK's GtkFileChooser GInterface.
type FileChooser struct {
	glib.Interface
}

// native returns a pointer to the underlying GObject as a GtkFileChooser.
func (v *FileChooser) native() *C.GtkFileChooser {
	return C.toGtkFileChooser(unsafe.Pointer(v.Native()))
}

func marshalFileChooser(p uintptr) (interface{}, error) {
	c := C.g_value_get_object(C.toGValue(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapFileChooser(*glib.InterfaceFromObjectNew(obj)), nil
}

func wrapFileChooser(intf glib.Interface) *FileChooser {
	return &FileChooser{intf}
}

// SelectAll is a wrapper around gtk_file_chooser_select_all().
func (v *FileChooser) SelectAll() {
	C.gtk_file_chooser_select_all(v.native())
}

// UnselectAll is a wrapper around gtk_file_chooser_unselect_all().
func (v *FileChooser) UnselectAll() {
	C.gtk_file_chooser_unselect_all(v.native())
}

// GetFilename is a wrapper around gtk_file_chooser_get_filename().
func (v *FileChooser) GetFilename() string {
	c := C.gtk_file_chooser_get_filename(v.native())
	defer C.g_free(C.gpointer(c))
	return goString(c)
}

// SetFilename is a wrapper around gtk_file_chooser_set_filename().
func (v *FileChooser) SetFilename(filename string) bool {
	var cstr *C.char
	if filename != "" {
		cstr = C.CString(filename)
		defer C.free(unsafe.Pointer(cstr))
	}
	c := C.gtk_file_chooser_set_filename(v.native(), cstr)
	return gobool(c)
}

// SetCurrentName is a wrapper around gtk_file_chooser_set_current_name().
func (v *FileChooser) SetCurrentName(name string) {
	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_file_chooser_set_current_name(v.native(), (*C.gchar)(cstr))
	return
}

// GetCurrentFolder is a wrapper around gtk_file_chooser_get_current_folder().
func (v *FileChooser) GetCurrentFolder() (string, error) {
	c := C.gtk_file_chooser_get_current_folder(v.native())
	if c == nil {
		return "", nilPtrErr
	}
	defer C.free(unsafe.Pointer(c))
	return goString(c), nil
}

// SetCurrentFolder is a wrapper around gtk_file_chooser_set_current_folder().
func (v *FileChooser) SetCurrentFolder(folder string) bool {
	cstr := C.CString(folder)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_file_chooser_set_current_folder(v.native(), (*C.gchar)(cstr))
	return gobool(c)
}

// SetPreviewWidget is a wrapper around gtk_file_chooser_set_preview_widget().
func (v *FileChooser) SetPreviewWidget(widget IWidget) {
	C.gtk_file_chooser_set_preview_widget(v.native(), widget.toWidget())
}

// SetPreviewWidgetActive is a wrapper around gtk_file_chooser_set_preview_widget_active().
func (v *FileChooser) SetPreviewWidgetActive(active bool) {
	C.gtk_file_chooser_set_preview_widget_active(v.native(), gbool(active))
}

// GetPreviewFilename is a wrapper around gtk_file_chooser_get_preview_filename().
func (v *FileChooser) GetPreviewFilename() string {
	c := C.gtk_file_chooser_get_preview_filename(v.native())
	defer C.free(unsafe.Pointer(c))
	return goString((*C.gchar)(c))
}

// AddFilter is a wrapper around gtk_file_chooser_add_filter().
func (v *FileChooser) AddFilter(filter *FileFilter) {
	C.gtk_file_chooser_add_filter(v.native(), filter.native())
}

// GetURI is a wrapper around gtk_file_chooser_get_uri().
func (v *FileChooser) GetURI() string {
	c := C.gtk_file_chooser_get_uri(v.native())
	defer C.g_free(C.gpointer(c))
	return goString(c)
}

// AddShortcutFolder is a wrapper around gtk_file_chooser_add_shortcut_folder().
func (v *FileChooser) AddShortcutFolder(folder string) bool {
	cstr := C.CString(folder)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_file_chooser_add_shortcut_folder(v.native(), cstr, nil)
	return gobool(c)
}

/*
 * GtkFileChooserButton
 */

// FileChooserButton is a representation of GTK's GtkFileChooserButton.
type FileChooserButton struct {
	Box
	// Interfaces
	FileChooser
}

// native returns a pointer to the underlying GtkFileChooserButton.
func (v *FileChooserButton) native() *C.GtkFileChooserButton {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toGtkFileChooserButton(ptr)
}

func marshalFileChooserButton(p uintptr) (interface{}, error) {
	c := C.g_value_get_object(C.toGValue(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapFileChooserButton(obj), nil
}

func wrapFileChooserButton(obj *glib.Object) *FileChooserButton {
	box := wrapBox(obj)
	fc := wrapFileChooser(*glib.InterfaceFromObjectNew(obj))
	return &FileChooserButton{*box, *fc}
}

// FileChooserButtonNew is a wrapper around gtk_file_chooser_button_new().
func FileChooserButtonNew(title string, action FileChooserAction) (*FileChooserButton, error) {
	cstr := C.CString(title)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_file_chooser_button_new((*C.gchar)(cstr),
		(C.GtkFileChooserAction)(action))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapFileChooserButton(obj), nil
}

/*
 * GtkFileChooserDialog
 */

// FileChooserDialog is a representation of GTK's GtkFileChooserDialog.
type FileChooserDialog struct {
	Dialog
	// Interfaces
	FileChooser
}

// native returns a pointer to the underlying GtkFileChooserDialog.
func (v *FileChooserDialog) native() *C.GtkFileChooserDialog {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toGtkFileChooserDialog(ptr)
}

func marshalFileChooserDialog(p uintptr) (interface{}, error) {
	c := C.g_value_get_object(C.toGValue(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapFileChooserDialog(obj), nil
}

func wrapFileChooserDialog(obj *glib.Object) *FileChooserDialog {
	dlg := wrapDialog(obj)
	fc := wrapFileChooser(*glib.InterfaceFromObjectNew(obj))
	return &FileChooserDialog{*dlg, *fc}
}

// FileChooserDialogNewWith1Button is a wrapper around gtk_file_chooser_dialog_new() with one button.
func FileChooserDialogNewWith1Button(
	title string,
	parent *Window,
	action FileChooserAction,
	first_button_text string,
	first_button_id ResponseType) (*FileChooserDialog, error) {
	c_title := C.CString(title)
	defer C.free(unsafe.Pointer(c_title))
	c_first_button_text := C.CString(first_button_text)
	defer C.free(unsafe.Pointer(c_first_button_text))
	c := C.gtk_file_chooser_dialog_new_1(
		(*C.gchar)(c_title), parent.native(), C.GtkFileChooserAction(action),
		(*C.gchar)(c_first_button_text), C.int(first_button_id))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapFileChooserDialog(obj), nil
}

// FileChooserDialogNewWith2Buttons is a wrapper around gtk_file_chooser_dialog_new() with two buttons.
func FileChooserDialogNewWith2Buttons(
	title string,
	parent *Window,
	action FileChooserAction,
	first_button_text string,
	first_button_id ResponseType,
	second_button_text string,
	second_button_id ResponseType) (*FileChooserDialog, error) {
	c_title := C.CString(title)
	defer C.free(unsafe.Pointer(c_title))
	c_first_button_text := C.CString(first_button_text)
	defer C.free(unsafe.Pointer(c_first_button_text))
	c_second_button_text := C.CString(second_button_text)
	defer C.free(unsafe.Pointer(c_second_button_text))
	c := C.gtk_file_chooser_dialog_new_2(
		(*C.gchar)(c_title), parent.native(), C.GtkFileChooserAction(action),
		(*C.gchar)(c_first_button_text), C.int(first_button_id),
		(*C.gchar)(c_second_button_text), C.int(second_button_id))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapFileChooserDialog(obj), nil
}

/*
 * GtkFileChooserWidget
 */

// FileChooserWidget is a representation of GTK's GtkFileChooserWidget.
type FileChooserWidget struct {
	Box
	// Interfaces
	FileChooser
}

// native returns a pointer to the underlying GtkFileChooserWidget.
func (v *FileChooserWidget) native() *C.GtkFileChooserWidget {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toGtkFileChooserWidget(ptr)
}

func marshalFileChooserWidget(p uintptr) (interface{}, error) {
	c := C.g_value_get_object(C.toGValue(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapFileChooserWidget(obj), nil
}

func wrapFileChooserWidget(obj *glib.Object) *FileChooserWidget {
	box := wrapBox(obj)
	fc := wrapFileChooser(*glib.InterfaceFromObjectNew(obj))
	return &FileChooserWidget{*box, *fc}
}

// FileChooserWidgetNew is a wrapper around gtk_file_chooser_widget_new().
func FileChooserWidgetNew(action FileChooserAction) (*FileChooserWidget, error) {
	c := C.gtk_file_chooser_widget_new((C.GtkFileChooserAction)(action))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapFileChooserWidget(obj), nil
}

/*
 * GtkFileFilter
 */

// FileChoser is a representation of GTK's GtkFileFilter GInterface.
type FileFilter struct {
	glib.Interface
}

// native returns a pointer to the underlying GObject as a GtkFileFilter.
func (v *FileFilter) native() *C.GtkFileFilter {
	return C.toGtkFileFilter(unsafe.Pointer(v.Native()))
}

func marshalFileFilter(p uintptr) (interface{}, error) {
	c := C.g_value_get_object(C.toGValue(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapFileFilter(*glib.InterfaceFromObjectNew(obj)), nil
}

func wrapFileFilter(intf glib.Interface) *FileFilter {
	return &FileFilter{intf}
}

// FileFilterNew is a wrapper around gtk_file_filter_new().
func FileFilterNew() (*FileFilter, error) {
	c := C.gtk_file_filter_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapFileFilter(*glib.InterfaceFromObjectNew(obj)), nil
}

// SetName is a wrapper around gtk_file_filter_set_name().
func (v *FileFilter) SetName(name string) {
	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_file_filter_set_name(v.native(), (*C.gchar)(cstr))
}

// AddMimeType is a wrapper around gtk_file_filter_add_mime_type().
func (v *FileFilter) AddMimeType(mimeType string) {
	cstr := C.CString(mimeType)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_file_filter_add_mime_type(v.native(), (*C.gchar)(cstr))
}

// AddPattern is a wrapper around gtk_file_filter_add_pattern().
func (v *FileFilter) AddPattern(pattern string) {
	cstr := C.CString(pattern)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_file_filter_add_pattern(v.native(), (*C.gchar)(cstr))
}

// AddPixbufFormats is a wrapper around gtk_file_filter_add_pixbuf_formats().
func (v *FileFilter) AddPixbufFormats() {
	C.gtk_file_filter_add_pixbuf_formats(v.native())
}

/*
 * GtkFontButton
 */

// FontButton is a representation of GTK's GtkFontButton.
type FontButton struct {
	Button
}

// native returns a pointer to the underlying GtkFontButton.
func (v *FontButton) native() *C.GtkFontButton {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toGtkFontButton(ptr)
}

func marshalFontButton(p uintptr) (interface{}, error) {
	c := C.g_value_get_object(C.toGValue(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapFontButton(obj), nil
}

func wrapFontButton(obj *glib.Object) *FontButton {
	button := wrapButton(obj)
	return &FontButton{*button}
}

// FontButtonNew is a wrapper around gtk_font_button_new().
func FontButtonNew() (*FontButton, error) {
	c := C.gtk_font_button_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapFontButton(obj), nil
}

// FontButtonNewWithFont is a wrapper around gtk_font_button_new_with_font().
func FontButtonNewWithFont(fontname string) (*FontButton, error) {
	cstr := C.CString(fontname)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_font_button_new_with_font((*C.gchar)(cstr))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapFontButton(obj), nil
}
