// Same copyright and license as the rest of the files in this project
// This file contains style related functions and structures

package gtk

// #include <gtk/gtk.h>
// #include "gtk.go.h"
import "C"
import (
	"runtime"
	"unsafe"

	"github.com/romychs/gotk3/glib"
)

// Actionable is a representation of GTK's GtkActionable GInterface.
type Actionable struct {
	glib.Interface
}

// native() returns a pointer to the underlying GtkActionable.
func (v *Actionable) native() *C.GtkActionable {
	return C.toGtkActionable(unsafe.Pointer(v.Native()))
}

/*
func (v* Actionable) toActionable() *Actionable {
	return v
}
*/

func marshalActionable(p uintptr) (interface{}, error) {
	c := C.g_value_get_object(C.toGValue(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapActionable(*glib.InterfaceFromObjectNew(obj)), nil
}

func wrapActionable(intf glib.Interface) *Actionable {
	return &Actionable{intf}
}

// const gchar *
// gtk_actionable_get_action_name (GtkActionable *actionable);
func (v *Actionable) GetActionName() string {
	c := C.gtk_actionable_get_action_name(v.native())
	defer C.g_free(C.gpointer(c))
	return goString(c)
}

// void
// gtk_actionable_set_action_name (GtkActionable *actionable,
//                                 const gchar *action_name);
func (v *Actionable) SetActionName(actionName string) {
	cstr := C.CString(actionName)
	defer C.free(unsafe.Pointer(cstr))

	C.gtk_actionable_set_action_name(v.native(), (*C.gchar)(cstr))
}

// GVariant *
// gtk_actionable_get_action_target_value
//                               (GtkActionable *actionable);
func (v *Actionable) GetActionTargetValue() (*glib.Variant, error) {
	c := C.gtk_actionable_get_action_target_value(v.native())
	if c == nil {
		return nil, nilPtrErr
	}

	return glib.WrapVariant(unsafe.Pointer(c)), nil
}

// void
// gtk_actionable_set_action_target_value
//                                (GtkActionable *actionable,
//                                 GVariant *target_value);
func (v *Actionable) SetActionTargetValue(targetValue *glib.Variant) {
	C.gtk_actionable_set_action_target_value(v.native(),
		C.toGVariant(targetValue.Native()))
}

// void
// gtk_actionable_set_detailed_action_name
//                                (GtkActionable *actionable,
//                                 const gchar *detailed_action_name);
func (v *Actionable) SetDetailedActionName(detailedActionName string) {
	cstr := C.CString(detailedActionName)
	defer C.free(unsafe.Pointer(cstr))

	C.gtk_actionable_set_detailed_action_name(v.native(), (*C.gchar)(cstr))
}

// ApplicationInhibitFlags is a representation of GTK's GtkApplicationInhibitFlags.
type ApplicationInhibitFlags int

const (
	APPLICATION_INHIBIT_LOGOUT  ApplicationInhibitFlags = C.GTK_APPLICATION_INHIBIT_LOGOUT
	APPLICATION_INHIBIT_SWITCH  ApplicationInhibitFlags = C.GTK_APPLICATION_INHIBIT_SWITCH
	APPLICATION_INHIBIT_SUSPEND ApplicationInhibitFlags = C.GTK_APPLICATION_INHIBIT_SUSPEND
	APPLICATION_INHIBIT_IDLE    ApplicationInhibitFlags = C.GTK_APPLICATION_INHIBIT_IDLE
)

/*
 * GtkApplication
 */

// Application is a representation of GTK's GtkApplication.
type Application struct {
	glib.Application
	// Interfaces
	glib.ActionMap
}

// native returns a pointer to the underlying GtkApplication.
func (v *Application) native() *C.GtkApplication {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toGtkApplication(ptr)
}

func marshalApplication(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapApplication(obj), nil
}

func wrapApplication(obj *glib.Object) *Application {
	actionMap := glib.ToActionMap(obj)
	return &Application{glib.Application{obj}, *actionMap}
}

// ApplicationNew is a wrapper around gtk_application_new().
func ApplicationNew(appID string, flags glib.ApplicationFlags) (*Application, error) {
	cstr := C.CString(appID)
	defer C.free(unsafe.Pointer(cstr))

	c := C.gtk_application_new((*C.gchar)(cstr), C.GApplicationFlags(flags))
	if c == nil {
		return nil, nilPtrErr
	}

	obj := glib.Take(unsafe.Pointer(c))
	return wrapApplication(obj), nil
}

// AddWindow is a wrapper around gtk_application_add_window().
func (v *Application) AddWindow(w *Window) {
	C.gtk_application_add_window(v.native(), w.native())
}

// RemoveWindow is a wrapper around gtk_application_remove_window().
func (v *Application) RemoveWindow(w *Window) {
	C.gtk_application_remove_window(v.native(), w.native())
}

// GetWindowByID is a wrapper around gtk_application_get_window_by_id().
func (v *Application) GetWindowByID(id uint) *Window {
	c := C.gtk_application_get_window_by_id(v.native(), C.guint(id))
	if c == nil {
		return nil
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapWindow(obj)
}

// GetActiveWindow is a wrapper around gtk_application_get_active_window().
func (v *Application) GetActiveWindow() *Window {
	c := C.gtk_application_get_active_window(v.native())
	if c == nil {
		return nil
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapWindow(obj)
}

// Uninhibit is a wrapper around gtk_application_uninhibit().
func (v *Application) Uninhibit(cookie uint) {
	C.gtk_application_uninhibit(v.native(), C.guint(cookie))
}

// GetAppMenu is a wrapper around gtk_application_get_app_menu().
func (v *Application) GetAppMenu() *glib.MenuModel {
	c := C.gtk_application_get_app_menu(v.native())
	if c == nil {
		return nil
	}
	return &glib.MenuModel{glib.Take(unsafe.Pointer(c))}
}

// SetAppMenu is a wrapper around gtk_application_set_app_menu().
func (v *Application) SetAppMenu(model glib.IMenuModel) {
	C.gtk_application_set_app_menu(v.native(),
		C.toGMenuModel(unsafe.Pointer(model.Native())))
}

// GetMenubar is a wrapper around gtk_application_get_menubar().
func (v *Application) GetMenubar() *glib.MenuModel {
	c := C.gtk_application_get_menubar(v.native())
	if c == nil {
		return nil
	}
	return &glib.MenuModel{glib.Take(unsafe.Pointer(c))}
}

// SetMenubar is a wrapper around gtk_application_set_menubar().
func (v *Application) SetMenubar(model glib.IMenuModel) {
	C.gtk_application_set_menubar(v.native(),
		C.toGMenuModel(unsafe.Pointer(model.Native())))
}

// IsInhibited is a wrapper around gtk_application_is_inhibited().
func (v *Application) IsInhibited(flags ApplicationInhibitFlags) bool {
	return gobool(C.gtk_application_is_inhibited(v.native(), C.GtkApplicationInhibitFlags(flags)))
}

// Inhibited is a wrapper around gtk_application_inhibit().
func (v *Application) Inhibited(w *Window, flags ApplicationInhibitFlags, reason string) uint {
	cstr := C.CString(reason)
	defer C.free(unsafe.Pointer(cstr))

	return uint(C.gtk_application_inhibit(v.native(), w.native(),
		C.GtkApplicationInhibitFlags(flags), (*C.gchar)(cstr)))
}

// void 	gtk_application_add_accelerator () // deprecated and uses a gvariant paramater
// void 	gtk_application_remove_accelerator () // deprecated and uses a gvariant paramater

// GetWindows is a wrapper around gtk_application_get_windows().
// Returned list is wrapped to return *gtk.Window elements.
func (v *Application) GetWindows() *glib.List {
	clist := C.gtk_application_get_windows(v.native())

	glist := glib.WrapList(uintptr(unsafe.Pointer(clist)))
	glist.DataWrapper(func(ptr unsafe.Pointer) interface{} {
		w := wrapWindow(glib.Take(ptr))
		return w
	})

	if glist != nil {
		runtime.SetFinalizer(glist, func(glist *glib.List) {
			glist.Free()
		})
	}

	return glist
}

/*
 * GtkApplicationWindow
 */

// ApplicationWindow is a representation of GTK's GtkApplicationWindow.
type ApplicationWindow struct {
	Window
	// Interfaces
	glib.ActionMap
}

// native returns a pointer to the underlying GtkApplicationWindow.
func (v *ApplicationWindow) native() *C.GtkApplicationWindow {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toGtkApplicationWindow(ptr)
}

func marshalApplicationWindow(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapApplicationWindow(obj), nil
}

func wrapApplicationWindow(obj *glib.Object) *ApplicationWindow {
	window := wrapWindow(obj)
	actionMap := glib.ToActionMap(obj)
	return &ApplicationWindow{*window, *actionMap}
}

// ApplicationWindowNew is a wrapper around gtk_application_window_new().
func ApplicationWindowNew(app *Application) (*ApplicationWindow, error) {
	c := C.gtk_application_window_new(app.native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapApplicationWindow(obj), nil
}

// SetShowMenubar is a wrapper around gtk_application_window_set_show_menubar().
func (v *ApplicationWindow) SetShowMenubar(b bool) {
	C.gtk_application_window_set_show_menubar(v.native(), gbool(b))
}

// GetShowMenubar is a wrapper around gtk_application_window_get_show_menubar().
func (v *ApplicationWindow) GetShowMenubar() bool {
	return gobool(C.gtk_application_window_get_show_menubar(v.native()))
}

// GetID is a wrapper around gtk_application_window_get_id().
func (v *ApplicationWindow) GetID() uint {
	return uint(C.gtk_application_window_get_id(v.native()))
}
