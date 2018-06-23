// Same copyright and license as the rest of the files in this project
// This file contains accelerator related functions and structures

package gtk

// #include <gtk/gtk.h>
// #include "gtk.go.h"
import "C"
import (
	"unsafe"

	"github.com/d2r2/gotk3/glib"
)

/*
 * GtkMenuShell
 */

// MenuShell is a representation of GTK's GtkMenuShell.
type MenuShell struct {
	Container
}

// native returns a pointer to the underlying GtkMenuShell.
func (v *MenuShell) native() *C.GtkMenuShell {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toGtkMenuShell(ptr)
}

func marshalMenuShell(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapMenuShell(obj), nil
}

func wrapMenuShell(obj *glib.Object) *MenuShell {
	return &MenuShell{Container{Widget{glib.InitiallyUnowned{obj}}}}
}

// Append is a wrapper around gtk_menu_shell_append().
func (v *MenuShell) Append(child IMenuItem) {
	C.gtk_menu_shell_append(v.native(), child.toWidget())
}

// Prepend is a wrapper around gtk_menu_shell_prepend().
func (v *MenuShell) Prepend(child IMenuItem) {
	C.gtk_menu_shell_prepend(v.native(), child.toWidget())
}

// Insert is a wrapper around gtk_menu_shell_insert().
func (v *MenuShell) Insert(child IMenuItem, position int) {
	C.gtk_menu_shell_insert(v.native(), child.toWidget(), C.gint(position))
}

// Deactivate is a wrapper around gtk_menu_shell_deactivate().
func (v *MenuShell) Deactivate() {
	C.gtk_menu_shell_deactivate(v.native())
}

// SelectItem is a wrapper around gtk_menu_shell_select_item().
func (v *MenuShell) SelectItem(child IMenuItem) {
	C.gtk_menu_shell_select_item(v.native(), child.toWidget())
}

// SelectFirst is a wrapper around gtk_menu_shell_select_first().
func (v *MenuShell) SelectFirst(searchSensitive bool) {
	C.gtk_menu_shell_select_first(v.native(), gbool(searchSensitive))
}

// Deselect is a wrapper around gtk_menu_shell_deselect().
func (v *MenuShell) Deselect() {
	C.gtk_menu_shell_deselect(v.native())
}

// ActivateItem is a wrapper around gtk_menu_shell_activate_item().
func (v *MenuShell) ActivateItem(child IMenuItem, forceDeactivate bool) {
	C.gtk_menu_shell_activate_item(v.native(), child.toWidget(), gbool(forceDeactivate))
}

// Cancel is a wrapper around gtk_menu_shell_cancel().
func (v *MenuShell) Cancel() {
	C.gtk_menu_shell_cancel(v.native())
}

// SetTakeFocus is a wrapper around gtk_menu_shell_set_take_focus().
func (v *MenuShell) SetTakeFocus(takeFocus bool) {
	C.gtk_menu_shell_set_take_focus(v.native(), gbool(takeFocus))
}

// gboolean 	gtk_menu_shell_get_take_focus ()
// GtkWidget * 	gtk_menu_shell_get_selected_item ()
// GtkWidget * 	gtk_menu_shell_get_parent_shell ()
// void 	gtk_menu_shell_bind_model ()

/*
 * GtkMenu
 */

// Menu is a representation of GTK's GtkMenu.
type Menu struct {
	MenuShell
}

// IMenu is an interface type implemented by all structs embedding
// a Menu.  It is meant to be used as an argument type for wrapper
// functions that wrap around a C GTK function taking a
// GtkMenu.
type IMenu interface {
	toMenu() *C.GtkMenu
	toWidget() *C.GtkWidget
}

// native() returns a pointer to the underlying GtkMenu.
func (v *Menu) native() *C.GtkMenu {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toGtkMenu(ptr)
}

func (v *Menu) toMenu() *C.GtkMenu {
	if v == nil {
		return nil
	}
	return v.native()
}

func marshalMenu(p uintptr) (interface{}, error) {
	c := C.g_value_get_object(C.toGValue(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapMenu(obj), nil
}

func wrapMenu(obj *glib.Object) *Menu {
	container := wrapContainer(obj)
	return &Menu{MenuShell{*container}}
}

// MenuNew() is a wrapper around gtk_menu_new().
func MenuNew() (*Menu, error) {
	c := C.gtk_menu_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapMenu(obj), nil
}

// GtkWidget* gtk_menu_new_from_model        (GMenuModel *model);
func MenuFromModelNew(model glib.IMenuModel) (*Menu, error) {
	c := C.gtk_menu_new_from_model(C.toGMenuModel(unsafe.Pointer(model.Native())))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapMenu(obj), nil
}

// Popdown() is a wrapper around gtk_menu_popdown().
func (v *Menu) Popdown() {
	C.gtk_menu_popdown(v.native())
}

// ReorderChild() is a wrapper around gtk_menu_reorder_child().
func (v *Menu) ReorderChild(child IWidget, position int) {
	C.gtk_menu_reorder_child(v.native(), child.toWidget(), C.gint(position))
}

// SetAccelGroup is a wrapper around gtk_menu_set_accel_group().
func (v *Menu) SetAccelGroup(accelGroup *AccelGroup) {
	C.gtk_menu_set_accel_group(v.native(), accelGroup.native())
}

// GetAccelGroup is a wrapper around gtk_menu_get_accel_group().
func (v *Menu) GetAccelGroup() *AccelGroup {
	c := C.gtk_menu_get_accel_group(v.native())
	if c == nil {
		return nil
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapAccelGroup(obj)
}

// SetAccelPath is a wrapper around gtk_menu_set_accel_path().
func (v *Menu) SetAccelPath(path string) {
	cstr := C.CString(path)
	defer C.free(unsafe.Pointer(cstr))

	C.gtk_menu_set_accel_path(v.native(), (*C.gchar)(cstr))
}

// GetAccelPath is a wrapper around gtk_menu_get_accel_path().
func (v *Menu) GetAccelPath() string {
	c := C.gtk_menu_get_accel_path(v.native())
	return goString(c)
}

/*
 * GtkMenuBar
 */

// MenuBar is a representation of GTK's GtkMenuBar.
type MenuBar struct {
	MenuShell
}

// native() returns a pointer to the underlying GtkMenuBar.
func (v *MenuBar) native() *C.GtkMenuBar {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toGtkMenuBar(ptr)
}

func marshalMenuBar(p uintptr) (interface{}, error) {
	c := C.g_value_get_object(C.toGValue(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapMenuBar(obj), nil
}

func wrapMenuBar(obj *glib.Object) *MenuBar {
	container := wrapContainer(obj)
	return &MenuBar{MenuShell{*container}}
}

// MenuBarNew() is a wrapper around gtk_menu_bar_new().
func MenuBarNew() (*MenuBar, error) {
	c := C.gtk_menu_bar_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapMenuBar(obj), nil
}

// GtkWidget *
// gtk_menu_bar_new_from_model (GMenuModel *model);
func MenuBarFromModelNew(model glib.IMenuModel) (*MenuBar, error) {
	c := C.gtk_menu_bar_new_from_model(C.toGMenuModel(unsafe.Pointer(model.Native())))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapMenuBar(obj), nil
}

/*
 * GtkMenuItem
 */

// MenuItem is a representation of GTK's GtkMenuItem.
type MenuItem struct {
	Bin
	// Interfaces
	Actionable
}

// IMenuItem is an interface type implemented by all structs
// embedding a MenuItem.  It is meant to be used as an argument type
// for wrapper functions that wrap around a C GTK function taking a
// GtkMenuItem.
type IMenuItem interface {
	toMenuItem() *C.GtkMenuItem
	toWidget() *C.GtkWidget
}

// native returns a pointer to the underlying GtkMenuItem.
func (v *MenuItem) native() *C.GtkMenuItem {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toGtkMenuItem(ptr)
}

func (v *MenuItem) toMenuItem() *C.GtkMenuItem {
	if v == nil {
		return nil
	}
	return v.native()
}

func marshalMenuItem(p uintptr) (interface{}, error) {
	c := C.g_value_get_object(C.toGValue(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapMenuItem(obj), nil
}

func wrapMenuItem(obj *glib.Object) *MenuItem {
	bin := wrapBin(obj)
	actionable := wrapActionable(*glib.InterfaceFromObjectNew(obj))
	return &MenuItem{*bin, *actionable}
}

// MenuItemNew() is a wrapper around gtk_menu_item_new().
func MenuItemNew() (*MenuItem, error) {
	c := C.gtk_menu_item_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapMenuItem(obj), nil
}

// MenuItemNewWithLabel() is a wrapper around gtk_menu_item_new_with_label().
func MenuItemNewWithLabel(label string) (*MenuItem, error) {
	cstr := C.CString(label)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_menu_item_new_with_label((*C.gchar)(cstr))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapMenuItem(obj), nil
}

// MenuItemNewWithMnemonic() is a wrapper around
// gtk_menu_item_new_with_mnemonic().
func MenuItemNewWithMnemonic(label string) (*MenuItem, error) {
	cstr := C.CString(label)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_menu_item_new_with_mnemonic((*C.gchar)(cstr))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapMenuItem(obj), nil
}

// SetSubmenu() is a wrapper around gtk_menu_item_set_submenu().
func (v *MenuItem) SetSubmenu(submenu IWidget) {
	C.gtk_menu_item_set_submenu(v.native(), submenu.toWidget())
}

// Sets text on the menu_item label
func (v *MenuItem) SetLabel(label string) {
	cstr := C.CString(label)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_menu_item_set_label(v.native(), (*C.gchar)(cstr))
}

// Gets text on the menu_item label
func (v *MenuItem) GetLabel() string {
	c := C.gtk_menu_item_get_label(v.native())
	return goString(c)
}

// SetAccelPath is a wrapper around gtk_menu_item_set_accel_path().
func (v *MenuItem) SetAccelPath(path string) {
	cstr := C.CString(path)
	defer C.free(unsafe.Pointer(cstr))

	C.gtk_menu_item_set_accel_path(v.native(), (*C.gchar)(cstr))
}

// GetAccelPath is a wrapper around gtk_menu_item_get_accel_path().
func (v *MenuItem) GetAccelPath() string {
	c := C.gtk_menu_item_get_accel_path(v.native())
	return goString(c)
}

/*
 * GtkSeparatorMenuItem
 */

// SeparatorMenuItem is a representation of GTK's GtkSeparatorMenuItem.
type SeparatorMenuItem struct {
	MenuItem
}

// native returns a pointer to the underlying GtkSeparatorMenuItem.
func (v *SeparatorMenuItem) native() *C.GtkSeparatorMenuItem {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toGtkSeparatorMenuItem(ptr)
}

func marshalSeparatorMenuItem(p uintptr) (interface{}, error) {
	c := C.g_value_get_object(C.toGValue(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapSeparatorMenuItem(obj), nil
}

func wrapSeparatorMenuItem(obj *glib.Object) *SeparatorMenuItem {
	menuItem := wrapMenuItem(obj)
	return &SeparatorMenuItem{*menuItem}
}

// SeparatorMenuItemNew is a wrapper around gtk_separator_menu_item_new().
func SeparatorMenuItemNew() (*SeparatorMenuItem, error) {
	c := C.gtk_separator_menu_item_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapSeparatorMenuItem(obj), nil
}

/*
 * GtkCheckMenuItem
 */

type CheckMenuItem struct {
	MenuItem
}

// native returns a pointer to the underlying GtkCheckMenuItem.
func (v *CheckMenuItem) native() *C.GtkCheckMenuItem {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toGtkCheckMenuItem(ptr)
}

func marshalCheckMenuItem(p uintptr) (interface{}, error) {
	c := C.g_value_get_object(C.toGValue(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapCheckMenuItem(obj), nil
}

func wrapCheckMenuItem(obj *glib.Object) *CheckMenuItem {
	menuItem := wrapMenuItem(obj)
	return &CheckMenuItem{*menuItem}
}

// CheckMenuItemNew is a wrapper around gtk_check_menu_item_new().
func CheckMenuItemNew() (*CheckMenuItem, error) {
	c := C.gtk_check_menu_item_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapCheckMenuItem(obj), nil
}

// CheckMenuItemNewWithLabel is a wrapper around
// gtk_check_menu_item_new_with_label().
func CheckMenuItemNewWithLabel(label string) (*CheckMenuItem, error) {
	cstr := C.CString(label)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_check_menu_item_new_with_label((*C.gchar)(cstr))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapCheckMenuItem(obj), nil
}

// CheckMenuItemNewWithMnemonic is a wrapper around
// gtk_check_menu_item_new_with_mnemonic().
func CheckMenuItemNewWithMnemonic(label string) (*CheckMenuItem, error) {
	cstr := C.CString(label)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_check_menu_item_new_with_mnemonic((*C.gchar)(cstr))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapCheckMenuItem(obj), nil
}

// GetActive is a wrapper around gtk_check_menu_item_get_active().
func (v *CheckMenuItem) GetActive() bool {
	c := C.gtk_check_menu_item_get_active(v.native())
	return gobool(c)
}

// SetActive is a wrapper around gtk_check_menu_item_set_active().
func (v *CheckMenuItem) SetActive(isActive bool) {
	C.gtk_check_menu_item_set_active(v.native(), gbool(isActive))
}

// Toggled is a wrapper around gtk_check_menu_item_toggled().
func (v *CheckMenuItem) Toggled() {
	C.gtk_check_menu_item_toggled(v.native())
}

// GetInconsistent is a wrapper around gtk_check_menu_item_get_inconsistent().
func (v *CheckMenuItem) GetInconsistent() bool {
	c := C.gtk_check_menu_item_get_inconsistent(v.native())
	return gobool(c)
}

// SetInconsistent is a wrapper around gtk_check_menu_item_set_inconsistent().
func (v *CheckMenuItem) SetInconsistent(setting bool) {
	C.gtk_check_menu_item_set_inconsistent(v.native(), gbool(setting))
}

// SetDrawAsRadio is a wrapper around gtk_check_menu_item_set_draw_as_radio().
func (v *CheckMenuItem) SetDrawAsRadio(drawAsRadio bool) {
	C.gtk_check_menu_item_set_draw_as_radio(v.native(), gbool(drawAsRadio))
}

// GetDrawAsRadio is a wrapper around gtk_check_menu_item_get_draw_as_radio().
func (v *CheckMenuItem) GetDrawAsRadio() bool {
	c := C.gtk_check_menu_item_get_draw_as_radio(v.native())
	return gobool(c)
}

/*
 * GtkRadioMenuItem
 */

// RadioMenuItem is a representation of GTK's GtkRadioMenuItem.
type RadioMenuItem struct {
	CheckMenuItem
}

// native returns a pointer to the underlying GtkRadioMenuItem.
func (v *RadioMenuItem) native() *C.GtkRadioMenuItem {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toGtkRadioMenuItem(ptr)
}

func marshalRadioMenuItem(p uintptr) (interface{}, error) {
	c := C.g_value_get_object(C.toGValue(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapRadioMenuItem(obj), nil
}

func wrapRadioMenuItem(obj *glib.Object) *RadioMenuItem {
	checkMenuItem := wrapCheckMenuItem(obj)
	return &RadioMenuItem{*checkMenuItem}
}

// RadioMenuItemNew is a wrapper around gtk_radio_menu_item_new().
func RadioMenuItemNew(group *glib.SList) (*RadioMenuItem, error) {
	c := C.gtk_radio_menu_item_new(cGSList(group))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapRadioMenuItem(obj), nil
}

// RadioMenuItemNewWithLabel is a wrapper around
// gtk_radio_menu_item_new_with_label().
func RadioMenuItemNewWithLabel(group *glib.SList, label string) (*RadioMenuItem, error) {
	cstr := C.CString(label)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_radio_menu_item_new_with_label(cGSList(group), (*C.gchar)(cstr))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapRadioMenuItem(obj), nil
}

// RadioMenuItemNewWithMnemonic is a wrapper around
// gtk_radio_menu_item_new_with_mnemonic().
func RadioMenuItemNewWithMnemonic(group *glib.SList, label string) (*RadioMenuItem, error) {
	cstr := C.CString(label)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_radio_menu_item_new_with_mnemonic(cGSList(group), (*C.gchar)(cstr))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapRadioMenuItem(obj), nil
}

// RadioMenuItemNewFromWidget is a wrapper around
// gtk_radio_menu_item_new_from_widget().
func RadioMenuItemNewFromWidget(group *RadioMenuItem) (*RadioMenuItem, error) {
	c := C.gtk_radio_menu_item_new_from_widget(group.native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapRadioMenuItem(obj), nil
}

// RadioMenuItemNewWithLabelFromWidget is a wrapper around
// gtk_radio_menu_item_new_with_label_from_widget().
func RadioMenuItemNewWithLabelFromWidget(group *RadioMenuItem, label string) (*RadioMenuItem, error) {
	cstr := C.CString(label)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_radio_menu_item_new_with_label_from_widget(group.native(),
		(*C.gchar)(cstr))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapRadioMenuItem(obj), nil
}

// RadioMenuItemNewWithMnemonicFromWidget is a wrapper around
// gtk_radio_menu_item_new_with_mnemonic_from_widget().
func RadioMenuItemNewWithMnemonicFromWidget(group *RadioMenuItem, label string) (*RadioMenuItem, error) {
	cstr := C.CString(label)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_radio_menu_item_new_with_mnemonic_from_widget(group.native(),
		(*C.gchar)(cstr))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapRadioMenuItem(obj), nil
}

// SetGroup is a wrapper around gtk_radio_menu_item_set_group().
func (v *RadioMenuItem) SetGroup(group *glib.SList) {
	C.gtk_radio_menu_item_set_group(v.native(), cGSList(group))
}

// GetGroup is a wrapper around gtk_radio_menu_item_get_group().
func (v *RadioMenuItem) GetGroup() (*glib.SList, error) {
	c := C.gtk_radio_menu_item_get_group(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	return glib.WrapSList(uintptr(unsafe.Pointer(c))), nil
}

/*
 * GtkToolbar
 */

// Toolbar is a representation of GTK's GtkToolbar.
type Toolbar struct {
	Container
}

// native returns a pointer to the underlying GtkToolbar.
func (v *Toolbar) native() *C.GtkToolbar {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toGtkToolbar(ptr)
}

func marshalToolbar(p uintptr) (interface{}, error) {
	c := C.g_value_get_object(C.toGValue(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapToolbar(obj), nil
}

func wrapToolbar(obj *glib.Object) *Toolbar {
	container := wrapContainer(obj)
	return &Toolbar{*container}
}

// ToolbarNew is a wrapper around gtk_toolbar_new().
func ToolbarNew() (*Toolbar, error) {
	c := C.gtk_toolbar_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapToolbar(obj), nil
}

// Insert is a wrapper around gtk_toolbar_insert().
func (v *Toolbar) Insert(item IToolItem, pos int) {
	C.gtk_toolbar_insert(v.native(), item.toToolItem(), C.gint(pos))
}

// GetItemIndex is a wrapper around gtk_toolbar_get_item_index().
func (v *Toolbar) GetItemIndex(item IToolItem) int {
	c := C.gtk_toolbar_get_item_index(v.native(), item.toToolItem())
	return int(c)
}

// GetNItems is a wrapper around gtk_toolbar_get_n_items().
func (v *Toolbar) GetNItems() int {
	c := C.gtk_toolbar_get_n_items(v.native())
	return int(c)
}

// GetNthItem is a wrapper around gtk_toolbar_get_nth_item().
func (v *Toolbar) GetNthItem(n int) *ToolItem {
	c := C.gtk_toolbar_get_nth_item(v.native(), C.gint(n))
	if c == nil {
		return nil
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapToolItem(obj)
}

// GetDropIndex is a wrapper around gtk_toolbar_get_drop_index().
func (v *Toolbar) GetDropIndex(x, y int) int {
	c := C.gtk_toolbar_get_drop_index(v.native(), C.gint(x), C.gint(y))
	return int(c)
}

// SetDropHighlightItem is a wrapper around
// gtk_toolbar_set_drop_highlight_item().
func (v *Toolbar) SetDropHighlightItem(toolItem IToolItem, index int) {
	C.gtk_toolbar_set_drop_highlight_item(v.native(),
		toolItem.toToolItem(), C.gint(index))
}

// SetShowArrow is a wrapper around gtk_toolbar_set_show_arrow().
func (v *Toolbar) SetShowArrow(showArrow bool) {
	C.gtk_toolbar_set_show_arrow(v.native(), gbool(showArrow))
}

// UnsetIconSize is a wrapper around gtk_toolbar_unset_icon_size().
func (v *Toolbar) UnsetIconSize() {
	C.gtk_toolbar_unset_icon_size(v.native())
}

// GetShowArrow is a wrapper around gtk_toolbar_get_show_arrow().
func (v *Toolbar) GetShowArrow() bool {
	c := C.gtk_toolbar_get_show_arrow(v.native())
	return gobool(c)
}

// GetStyle is a wrapper around gtk_toolbar_get_style().
func (v *Toolbar) GetStyle() ToolbarStyle {
	c := C.gtk_toolbar_get_style(v.native())
	return ToolbarStyle(c)
}

// GetIconSize is a wrapper around gtk_toolbar_get_icon_size().
func (v *Toolbar) GetIconSize() IconSize {
	c := C.gtk_toolbar_get_icon_size(v.native())
	return IconSize(c)
}

// GetReliefStyle is a wrapper around gtk_toolbar_get_relief_style().
func (v *Toolbar) GetReliefStyle() ReliefStyle {
	c := C.gtk_toolbar_get_relief_style(v.native())
	return ReliefStyle(c)
}

// SetStyle is a wrapper around gtk_toolbar_set_style().
func (v *Toolbar) SetStyle(style ToolbarStyle) {
	C.gtk_toolbar_set_style(v.native(), C.GtkToolbarStyle(style))
}

// SetIconSize is a wrapper around gtk_toolbar_set_icon_size().
func (v *Toolbar) SetIconSize(iconSize IconSize) {
	C.gtk_toolbar_set_icon_size(v.native(), C.GtkIconSize(iconSize))
}

// UnsetStyle is a wrapper around gtk_toolbar_unset_style().
func (v *Toolbar) UnsetStyle() {
	C.gtk_toolbar_unset_style(v.native())
}

/*
 * GtkToolItem
 */

// ToolItem is a representation of GTK's GtkToolItem.
type ToolItem struct {
	Bin
}

// IToolItem is an interface type implemented by all structs embedding
// a ToolItem.  It is meant to be used as an argument type for wrapper
// functions that wrap around a C GTK function taking a GtkToolItem.
type IToolItem interface {
	toToolItem() *C.GtkToolItem
}

// native returns a pointer to the underlying GtkToolItem.
func (v *ToolItem) native() *C.GtkToolItem {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toGtkToolItem(ptr)
}

func (v *ToolItem) toToolItem() *C.GtkToolItem {
	return v.native()
}

func marshalToolItem(p uintptr) (interface{}, error) {
	c := C.g_value_get_object(C.toGValue(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapToolItem(obj), nil
}

func wrapToolItem(obj *glib.Object) *ToolItem {
	bin := wrapBin(obj)
	return &ToolItem{*bin}
}

// ToolItemNew is a wrapper around gtk_tool_item_new().
func ToolItemNew() (*ToolItem, error) {
	c := C.gtk_tool_item_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapToolItem(obj), nil
}

// SetHomogeneous is a wrapper around gtk_tool_item_set_homogeneous().
func (v *ToolItem) SetHomogeneous(homogeneous bool) {
	C.gtk_tool_item_set_homogeneous(v.native(), gbool(homogeneous))
}

// GetHomogeneous is a wrapper around gtk_tool_item_get_homogeneous().
func (v *ToolItem) GetHomogeneous() bool {
	c := C.gtk_tool_item_get_homogeneous(v.native())
	return gobool(c)
}

// SetExpand is a wrapper around gtk_tool_item_set_expand().
func (v *ToolItem) SetExpand(expand bool) {
	C.gtk_tool_item_set_expand(v.native(), gbool(expand))
}

// GetExpand is a wrapper around gtk_tool_item_get_expand().
func (v *ToolItem) GetExpand() bool {
	c := C.gtk_tool_item_get_expand(v.native())
	return gobool(c)
}

// SetTooltipText is a wrapper around gtk_tool_item_set_tooltip_text().
func (v *ToolItem) SetTooltipText(text string) {
	var cstr *C.char
	if text != "" {
		cstr = C.CString(text)
		defer C.free(unsafe.Pointer(cstr))
	}
	C.gtk_tool_item_set_tooltip_text(v.native(), (*C.gchar)(cstr))
}

// SetTooltipMarkup is a wrapper around gtk_tool_item_set_tooltip_markup().
func (v *ToolItem) SetTooltipMarkup(text string) {
	cstr := C.CString(text)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_tool_item_set_tooltip_markup(v.native(), (*C.gchar)(cstr))
}

// SetUseDragWindow is a wrapper around gtk_tool_item_set_use_drag_window().
func (v *ToolItem) SetUseDragWindow(useDragWindow bool) {
	C.gtk_tool_item_set_use_drag_window(v.native(), gbool(useDragWindow))
}

// GetUseDragWindow is a wrapper around gtk_tool_item_get_use_drag_window().
func (v *ToolItem) GetUseDragWindow() bool {
	c := C.gtk_tool_item_get_use_drag_window(v.native())
	return gobool(c)
}

// SetVisibleHorizontal is a wrapper around
// gtk_tool_item_set_visible_horizontal().
func (v *ToolItem) SetVisibleHorizontal(visibleHorizontal bool) {
	C.gtk_tool_item_set_visible_horizontal(v.native(),
		gbool(visibleHorizontal))
}

// GetVisibleHorizontal is a wrapper around
// gtk_tool_item_get_visible_horizontal().
func (v *ToolItem) GetVisibleHorizontal() bool {
	c := C.gtk_tool_item_get_visible_horizontal(v.native())
	return gobool(c)
}

// SetVisibleVertical is a wrapper around gtk_tool_item_set_visible_vertical().
func (v *ToolItem) SetVisibleVertical(visibleVertical bool) {
	C.gtk_tool_item_set_visible_vertical(v.native(), gbool(visibleVertical))
}

// GetVisibleVertical is a wrapper around gtk_tool_item_get_visible_vertical().
func (v *ToolItem) GetVisibleVertical() bool {
	c := C.gtk_tool_item_get_visible_vertical(v.native())
	return gobool(c)
}

// SetIsImportant is a wrapper around gtk_tool_item_set_is_important().
func (v *ToolItem) SetIsImportant(isImportant bool) {
	C.gtk_tool_item_set_is_important(v.native(), gbool(isImportant))
}

// GetIsImportant is a wrapper around gtk_tool_item_get_is_important().
func (v *ToolItem) GetIsImportant() bool {
	c := C.gtk_tool_item_get_is_important(v.native())
	return gobool(c)
}

// TODO: gtk_tool_item_get_ellipsize_mode

// GetIconSize is a wrapper around gtk_tool_item_get_icon_size().
func (v *ToolItem) GetIconSize() IconSize {
	c := C.gtk_tool_item_get_icon_size(v.native())
	return IconSize(c)
}

// GetOrientation is a wrapper around gtk_tool_item_get_orientation().
func (v *ToolItem) GetOrientation() Orientation {
	c := C.gtk_tool_item_get_orientation(v.native())
	return Orientation(c)
}

// GetToolbarStyle is a wrapper around gtk_tool_item_get_toolbar_style().
func (v *ToolItem) gtk_tool_item_get_toolbar_style() ToolbarStyle {
	c := C.gtk_tool_item_get_toolbar_style(v.native())
	return ToolbarStyle(c)
}

// GetReliefStyle is a wrapper around gtk_tool_item_get_relief_style().
func (v *ToolItem) GetReliefStyle() ReliefStyle {
	c := C.gtk_tool_item_get_relief_style(v.native())
	return ReliefStyle(c)
}

// GetTextAlignment is a wrapper around gtk_tool_item_get_text_alignment().
func (v *ToolItem) GetTextAlignment() float32 {
	c := C.gtk_tool_item_get_text_alignment(v.native())
	return float32(c)
}

// GetTextOrientation is a wrapper around gtk_tool_item_get_text_orientation().
func (v *ToolItem) GetTextOrientation() Orientation {
	c := C.gtk_tool_item_get_text_orientation(v.native())
	return Orientation(c)
}

// RetrieveProxyMenuItem is a wrapper around
// gtk_tool_item_retrieve_proxy_menu_item()
func (v *ToolItem) RetrieveProxyMenuItem() *MenuItem {
	c := C.gtk_tool_item_retrieve_proxy_menu_item(v.native())
	if c == nil {
		return nil
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapMenuItem(obj)
}

// SetProxyMenuItem is a wrapper around gtk_tool_item_set_proxy_menu_item().
func (v *ToolItem) SetProxyMenuItem(menuItemId string, menuItem IMenuItem) {
	cstr := C.CString(menuItemId)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_tool_item_set_proxy_menu_item(v.native(), (*C.gchar)(cstr),
		C.toGtkWidget(unsafe.Pointer(menuItem.toMenuItem())))
}

// RebuildMenu is a wrapper around gtk_tool_item_rebuild_menu().
func (v *ToolItem) RebuildMenu() {
	C.gtk_tool_item_rebuild_menu(v.native())
}

// ToolbarReconfigured is a wrapper around gtk_tool_item_toolbar_reconfigured().
func (v *ToolItem) ToolbarReconfigured() {
	C.gtk_tool_item_toolbar_reconfigured(v.native())
}

// TODO: gtk_tool_item_get_text_size_group

/*
 * GtkToolButton
 */

// ToolButton is a representation of GTK's GtkToolButton.
type ToolButton struct {
	ToolItem
	// Interfaces
	Actionable
}

// native returns a pointer to the underlying GtkToolButton.
func (v *ToolButton) native() *C.GtkToolButton {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toGtkToolButton(ptr)
}

func marshalToolButton(p uintptr) (interface{}, error) {
	c := C.g_value_get_object(C.toGValue(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapToolButton(obj), nil
}

func wrapToolButton(obj *glib.Object) *ToolButton {
	toolItem := wrapToolItem(obj)
	actionable := wrapActionable(*glib.InterfaceFromObjectNew(obj))
	return &ToolButton{*toolItem, *actionable}
}

// ToolButtonNew is a wrapper around gtk_tool_button_new().
func ToolButtonNew(iconWidget IWidget, label string) (*ToolButton, error) {
	// label could be empty
	var cstr *C.char
	if label != "" {
		cstr = C.CString(label)
		defer C.free(unsafe.Pointer(cstr))
	}
	// iconWidget could be empty
	var wdg *C.GtkWidget
	if iconWidget != nil {
		wdg = iconWidget.toWidget()
	}
	c := C.gtk_tool_button_new(wdg, (*C.gchar)(cstr))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapToolButton(obj), nil
}

// SetLabel is a wrapper around gtk_tool_button_set_label().
func (v *ToolButton) SetLabel(label string) {
	cstr := C.CString(label)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_tool_button_set_label(v.native(), (*C.gchar)(cstr))
}

// GetLabel is a wrapper aroud gtk_tool_button_get_label().
func (v *ToolButton) GetLabel() string {
	c := C.gtk_tool_button_get_label(v.native())
	return goString(c)
}

// SetUseUnderline is a wrapper around gtk_tool_button_set_use_underline().
func (v *ToolButton) SetGetUnderline(useUnderline bool) {
	C.gtk_tool_button_set_use_underline(v.native(), gbool(useUnderline))
}

// GetUseUnderline is a wrapper around gtk_tool_button_get_use_underline().
func (v *ToolButton) GetuseUnderline() bool {
	c := C.gtk_tool_button_get_use_underline(v.native())
	return gobool(c)
}

// SetIconName is a wrapper around gtk_tool_button_set_icon_name().
func (v *ToolButton) SetIconName(iconName string) {
	cstr := C.CString(iconName)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_tool_button_set_icon_name(v.native(), (*C.gchar)(cstr))
}

// GetIconName is a wrapper around gtk_tool_button_get_icon_name().
func (v *ToolButton) GetIconName() string {
	c := C.gtk_tool_button_get_icon_name(v.native())
	return goString(c)
}

// SetIconWidget is a wrapper around gtk_tool_button_set_icon_widget().
func (v *ToolButton) SetIconWidget(iconWidget IWidget) {
	C.gtk_tool_button_set_icon_widget(v.native(), iconWidget.toWidget())
}

// GetIconWidget is a wrapper around gtk_tool_button_get_icon_widget().
func (v *ToolButton) GetIconWidget() *Widget {
	c := C.gtk_tool_button_get_icon_widget(v.native())
	if c == nil {
		return nil
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapWidget(obj)
}

// SetLabelWidget is a wrapper around gtk_tool_button_set_label_widget().
func (v *ToolButton) SetLabelWidget(labelWidget IWidget) {
	C.gtk_tool_button_set_label_widget(v.native(), labelWidget.toWidget())
}

// GetLabelWidget is a wrapper around gtk_tool_button_get_label_widget().
func (v *ToolButton) GetLabelWidget() *Widget {
	c := C.gtk_tool_button_get_label_widget(v.native())
	if c == nil {
		return nil
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapWidget(obj)
}

/*
 * GtkSeparatorToolItem
 */

// SeparatorToolItem is a representation of GTK's GtkSeparatorToolItem.
type SeparatorToolItem struct {
	ToolItem
	// Interfaces
	Actionable
}

// native returns a pointer to the underlying GtkSeparatorToolItem.
func (v *SeparatorToolItem) native() *C.GtkSeparatorToolItem {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toGtkSeparatorToolItem(ptr)
}

func marshalSeparatorToolItem(p uintptr) (interface{}, error) {
	c := C.g_value_get_object(C.toGValue(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapSeparatorToolItem(obj), nil
}

func wrapSeparatorToolItem(obj *glib.Object) *SeparatorToolItem {
	toolItem := wrapToolItem(obj)
	actionable := wrapActionable(*glib.InterfaceFromObjectNew(obj))
	return &SeparatorToolItem{*toolItem, *actionable}
}

// SeparatorToolItemNew is a wrapper around gtk_separator_tool_item_new().
func SeparatorToolItemNew() (*SeparatorToolItem, error) {
	c := C.gtk_separator_tool_item_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapSeparatorToolItem(obj), nil
}

// SetDraw is a wrapper around gtk_separator_tool_item_set_draw().
func (v *SeparatorToolItem) SetDraw(draw bool) {
	C.gtk_separator_tool_item_set_draw(v.native(), gbool(draw))
}

// GetDraw is a wrapper around gtk_separator_tool_item_get_draw().
func (v *SeparatorToolItem) GetDraw() bool {
	c := C.gtk_separator_tool_item_get_draw(v.native())
	return gobool(c)
}
