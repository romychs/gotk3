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

// IMenuModel is an interface type implemented by all structs
// embedding a MenuModel.
type IMenuModel interface {
	toMenuModel() *C.GMenuModel
	// Use this method to expose access to GLIB underlying object
	// in external packages.
	Native() uintptr
}

// MenuModel is a representation of GMenuModel.
type MenuModel struct {
	*Object
}

// Static cast to verify at compile time that type on the right side
// implement corresponding interface on the left.
var _ IMenuModel = &MenuModel{}

// native() returns a pointer to the underlying GMenuModel.
func (v *MenuModel) native() *C.GMenuModel {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toGMenuModel(ptr)
}

func (v *MenuModel) toMenuModel() *C.GMenuModel {
	return v.native()
}

func (v *MenuModel) Native() uintptr {
	return uintptr(unsafe.Pointer(v.native()))
}

func marshalMenuModel(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	return wrapMenuModel(wrapObject(unsafe.Pointer(c))), nil
}

func wrapMenuModel(obj *Object) *MenuModel {
	return &MenuModel{obj}
}

// IsMutable is a wrapper around g_menu_model_is_mutable().
func (v *MenuModel) IsMutable() bool {
	return gobool(C.g_menu_model_is_mutable(v.native()))
}

// GetNItems is a wrapper around g_menu_model_get_n_items().
func (v *MenuModel) GetNItems() int {
	return int(C.g_menu_model_get_n_items(v.native()))
}

// GetItemLink is a wrapper around g_menu_model_get_item_link().
func (v *MenuModel) GetItemLink(index int, link string) *MenuModel {
	cstr := C.CString(link)
	defer C.free(unsafe.Pointer(cstr))
	c := C.g_menu_model_get_item_link(v.native(), C.gint(index), (*C.gchar)(cstr))
	if c == nil {
		return nil
	}
	return wrapMenuModel(wrapObject(unsafe.Pointer(c)))
}

// ItemsChanged is a wrapper around g_menu_model_items_changed().
func (v *MenuModel) ItemsChanged(position, removed, added int) {
	C.g_menu_model_items_changed(v.native(), C.gint(position), C.gint(removed), C.gint(added))
}

// GVariant * 	g_menu_model_get_item_attribute_value ()
// gboolean 	g_menu_model_get_item_attribute ()
// GMenuAttributeIter * 	g_menu_model_iterate_item_attributes ()
// GMenuLinkIter * 	g_menu_model_iterate_item_links ()

// Menu is a representation of GMenu.
type Menu struct {
	MenuModel
}

// Static cast to verify at compile time that type on the right side
// implement corresponding interface on the left.
var _ IMenuModel = &Menu{}

// native() returns a pointer to the underlying GMenu.
func (v *Menu) native() *C.GMenu {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toGMenu(ptr)
}

func marshalMenu(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	return wrapMenu(wrapObject(unsafe.Pointer(c))), nil
}

func wrapMenu(obj *Object) *Menu {
	return &Menu{MenuModel{obj}}
}

// MenuNew is a wrapper around g_menu_new().
func MenuNew() (*Menu, error) {
	c := C.g_menu_new()
	if c == nil {
		return nil, errNilPtr
	}
	return wrapMenu(wrapObject(unsafe.Pointer(c))), nil
}

// Freeze is a wrapper around g_menu_freeze().
func (v *Menu) Freeze() {
	C.g_menu_freeze(v.native())
}

// Insert is a wrapper around g_menu_insert().
func (v *Menu) Insert(position int, label, detailedAction string) {
	var cstr1, cstr2 *C.char
	if label != "" {
		cstr1 = C.CString(label)
		defer C.free(unsafe.Pointer(cstr1))
	}

	if detailedAction != "" {
		cstr2 = C.CString(detailedAction)
		defer C.free(unsafe.Pointer(cstr2))
	}

	C.g_menu_insert(v.native(), C.gint(position), (*C.gchar)(cstr1), (*C.gchar)(cstr2))
}

// Prepend is a wrapper around g_menu_prepend().
func (v *Menu) Prepend(label, detailedAction string) {
	var cstr1, cstr2 *C.char
	if label != "" {
		cstr1 = C.CString(label)
		defer C.free(unsafe.Pointer(cstr1))
	}

	if detailedAction != "" {
		cstr2 = C.CString(detailedAction)
		defer C.free(unsafe.Pointer(cstr2))
	}

	C.g_menu_prepend(v.native(), (*C.gchar)(cstr1), (*C.gchar)(cstr2))
}

// Append is a wrapper around g_menu_append().
func (v *Menu) Append(label, detailedAction string) {
	var cstr1, cstr2 *C.char
	if label != "" {
		cstr1 = C.CString(label)
		defer C.free(unsafe.Pointer(cstr1))
	}

	if detailedAction != "" {
		cstr2 = C.CString(detailedAction)
		defer C.free(unsafe.Pointer(cstr2))
	}

	C.g_menu_append(v.native(), (*C.gchar)(cstr1), (*C.gchar)(cstr2))
}

// InsertItem is a wrapper around g_menu_insert_item().
func (v *Menu) InsertItem(position int, item *MenuItem) {
	C.g_menu_insert_item(v.native(), C.gint(position), item.native())
}

// AppendItem is a wrapper around g_menu_append_item().
func (v *Menu) AppendItem(item *MenuItem) {
	C.g_menu_append_item(v.native(), item.native())
}

// PrependItem is a wrapper around g_menu_prepend_item().
func (v *Menu) PrependItem(item *MenuItem) {
	C.g_menu_prepend_item(v.native(), item.native())
}

// InsertSection is a wrapper around g_menu_insert_section().
func (v *Menu) InsertSection(position int, label string, section IMenuModel) {
	var cstr *C.char
	if label != "" {
		cstr = C.CString(label)
		defer C.free(unsafe.Pointer(cstr))
	}

	C.g_menu_insert_section(v.native(), C.gint(position), (*C.gchar)(cstr),
		section.toMenuModel())
}

// PrependSection is a wrapper around g_menu_prepend_section().
func (v *Menu) PrependSection(label string, section IMenuModel) {
	var cstr *C.char
	if label != "" {
		cstr = C.CString(label)
		defer C.free(unsafe.Pointer(cstr))
	}

	C.g_menu_prepend_section(v.native(), (*C.gchar)(cstr),
		section.toMenuModel())
}

// AppendSection is a wrapper around g_menu_append_section().
func (v *Menu) AppendSection(label string, section IMenuModel) {
	var cstr *C.char
	if label != "" {
		cstr = C.CString(label)
		defer C.free(unsafe.Pointer(cstr))
	}

	C.g_menu_append_section(v.native(), (*C.gchar)(cstr),
		section.toMenuModel())
}

// InsertSubmenu is a wrapper around g_menu_insert_submenu().
func (v *Menu) InsertSubmenu(position int, label string, submenu IMenuModel) {
	var cstr *C.char
	if label != "" {
		cstr = C.CString(label)
		defer C.free(unsafe.Pointer(cstr))
	}

	C.g_menu_insert_submenu(v.native(), C.gint(position), (*C.gchar)(cstr),
		submenu.toMenuModel())
}

// PrependSubmenu is a wrapper around g_menu_prepend_submenu().
func (v *Menu) PrependSubmenu(label string, submenu IMenuModel) {
	var cstr *C.char
	if label != "" {
		cstr = C.CString(label)
		defer C.free(unsafe.Pointer(cstr))
	}

	C.g_menu_prepend_submenu(v.native(), (*C.gchar)(cstr),
		submenu.toMenuModel())
}

// AppendSubmenu is a wrapper around g_menu_append_submenu().
func (v *Menu) AppendSubmenu(label string, submenu IMenuModel) {
	var cstr *C.char
	if label != "" {
		cstr = C.CString(label)
		defer C.free(unsafe.Pointer((*C.gchar)(cstr)))
	}

	C.g_menu_append_submenu(v.native(), (*C.gchar)(cstr),
		submenu.toMenuModel())
}

// Remove is a wrapper around g_menu_remove().
func (v *Menu) Remove(position int) {
	C.g_menu_remove(v.native(), C.gint(position))
}

// RemoveAll is a wrapper around g_menu_remove_all().
func (v *Menu) RemoveAll() {
	C.g_menu_remove_all(v.native())
}

// MenuItem is a representation of GMenuItem.
type MenuItem struct {
	*Object
}

// native() returns a pointer to the underlying GMenuItem.
func (v *MenuItem) native() *C.GMenuItem {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toGMenuItem(ptr)
}

func marshalMenuItem(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	return wrapMenuItem(wrapObject(unsafe.Pointer(c))), nil
}

func wrapMenuItem(obj *Object) *MenuItem {
	return &MenuItem{obj}
}

// MenuItemNew is a wrapper around g_menu_item_new().
func MenuItemNew(label, detailedAction string) (*MenuItem, error) {
	var cstr1, cstr2 *C.char
	if label != "" {
		cstr1 = C.CString(label)
		defer C.free(unsafe.Pointer(cstr1))
	}

	if detailedAction != "" {
		cstr2 = C.CString(detailedAction)
		defer C.free(unsafe.Pointer(cstr2))
	}

	c := C.g_menu_item_new((*C.gchar)(cstr1), (*C.gchar)(cstr2))
	if c == nil {
		return nil, errNilPtr
	}
	return wrapMenuItem(wrapObject(unsafe.Pointer(c))), nil
}

/*
// MenuItemSectionNew is a wrapper around g_menu_item_new_section().
func MenuItemSectionNew(label string, section IMenuModel) (*MenuItem, error) {
	var cstr1 *C.gchar
	if label != "" {
		cstr1 := C.CString(label)
		defer C.free(unsafe.Pointer(cstr1))
	}

	c := C.g_menu_item_new_section((*C.gchar)(cstr1), section.toMenuModel().native())
	if c == nil {
		return nil, errNilPtr
	}
	return wrapMenuItem(wrapObject(unsafe.Pointer(c))), nil
}

// MenuItemSubmenuNew is a wrapper around g_menu_item_new_submenu().
func MenuItemSubmenuNew(label string, submenu IMenuModel) (*MenuItem, error) {
	cstr1 := C.CString(label)
	defer C.free(unsafe.Pointer(cstr1))

	c := C.g_menu_item_new_submenu((*C.gchar)(cstr1), submenu.toMenuModel().native())
	if c == nil {
		return nil, errNilPtr
	}
	return wrapMenuItem(wrapObject(unsafe.Pointer(c))), nil
}

// MenuItemFromModelNew is a wrapper around g_menu_item_new_from_model().
func MenuItemFromModelNew(model IMenuModel, index int) (*MenuItem, error) {
	c := C.g_menu_item_new_from_model(model.native(), C.gint(index))
	if c == nil {
		return nil, errNilPtr
	}
	return wrapMenuItem(wrapObject(unsafe.Pointer(c))), nil
}
*/
//SetLabel is a wrapper around g_menu_item_set_label().
func (v *MenuItem) SetLabel(label string) {
	cstr := C.CString(label)
	defer C.free(unsafe.Pointer(cstr))

	C.g_menu_item_set_label(v.native(), (*C.gchar)(cstr))
}

//SetDetailedAction is a wrapper around g_menu_item_set_detailed_action().
func (v *MenuItem) SetDetailedAction(act string) {
	cstr := C.CString(act)
	defer C.free(unsafe.Pointer(cstr))

	C.g_menu_item_set_detailed_action(v.native(), (*C.gchar)(cstr))
}

//SetSection is a wrapper around g_menu_item_set_section().
func (v *MenuItem) SetSection(section IMenuModel) {
	C.g_menu_item_set_section(v.native(),
		section.toMenuModel())
}

//SetSubmenu is a wrapper around g_menu_item_set_submenu().
func (v *MenuItem) SetSubmenu(submenu IMenuModel) {
	C.g_menu_item_set_submenu(v.native(),
		submenu.toMenuModel())
}

//GetLink is a wrapper around g_menu_item_get_link().
func (v *MenuItem) GetLink(link string) (*MenuModel, error) {
	cstr := C.CString(link)
	defer C.free(unsafe.Pointer(cstr))

	c := C.g_menu_item_get_link(v.native(), (*C.gchar)(cstr))
	if c == nil {
		return nil, errNilPtr
	}
	return wrapMenuModel(wrapObject(unsafe.Pointer(c))), nil
}

//SetLink is a wrapper around g_menu_item_Set_link().
func (v *MenuItem) SetLink(link string, model IMenuModel) {
	cstr := C.CString(link)
	defer C.free(unsafe.Pointer(cstr))

	C.g_menu_item_set_link(v.native(), (*C.gchar)(cstr),
		model.toMenuModel())
}

// void
// g_menu_item_set_attribute_value (GMenuItev *menu_item,
//                                  const gchar *attribute,
//                                  GVariant *value);
func (v *MenuItem) SetAttributeValue(attribute string, value *Variant) {
	cstr := C.CString(attribute)
	defer C.free(unsafe.Pointer(cstr))

	C.g_menu_item_set_attribute_value(v.native(), (*C.gchar)(cstr), value.native())
}

// void
// g_menu_item_set_icon (GMenuItev *menu_item,
//                       GIcon *icon);
func (v *MenuItem) SetIcon(icon *Icon) {
	C.g_menu_item_set_icon(v.native(), icon.native())
}

// void 	g_menu_item_set_action_and_target_value ()
// void 	g_menu_item_set_action_and_target ()
// GVariant * 	g_menu_item_get_attribute_value ()
// gboolean 	g_menu_item_get_attribute ()
// void 	g_menu_item_set_attribute_value ()
// void 	g_menu_item_set_attribute ()

func init() {
	tm := []TypeMarshaler{
		// Enums
		{Type(C.g_menu_model_get_type()), marshalMenuModel},
		{Type(C.g_menu_get_type()), marshalMenu},
		{Type(C.g_menu_item_get_type()), marshalMenuItem},
	}
	RegisterGValueMarshalers(tm)
}
