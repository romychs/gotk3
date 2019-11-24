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

	"github.com/d2r2/gotk3/glib"
)

/*
 * GtkOrientable
 */

// Orientable is a representation of GTK's GtkOrientable GInterface.
type Orientable struct {
	glib.Interface
}

// IOrientable is an interface type implemented by all structs
// embedding an Orientable.  It is meant to be used as an argument type
// for wrapper functions that wrap around a C GTK function taking a
// GtkOrientable.
type IOrientable interface {
	toOrientable() *C.GtkOrientable
}

// native returns a pointer to the underlying GObject as a GtkOrientable.
func (v *Orientable) native() *C.GtkOrientable {
	return C.toGtkOrientable(unsafe.Pointer(v.Native()))
}

func marshalOrientable(p uintptr) (interface{}, error) {
	c := C.g_value_get_object(C.toGValue(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapOrientable(*glib.InterfaceFromObjectNew(obj)), nil
}

func wrapOrientable(intf glib.Interface) *Orientable {
	return &Orientable{intf}
}

// GetOrientation is a wrapper around gtk_orientable_get_orientation().
func (v *Orientable) GetOrientation() Orientation {
	c := C.gtk_orientable_get_orientation(v.native())
	return Orientation(c)
}

// SetOrientation is a wrapper around gtk_orientable_set_orientation().
func (v *Orientable) SetOrientation(orientation Orientation) {
	C.gtk_orientable_set_orientation(v.native(),
		C.GtkOrientation(orientation))
}

/*
 * GtkBox
 */

// Box is a representation of GTK's GtkBox.
type Box struct {
	Container
}

// native returns a pointer to the underlying GtkBox.
func (v *Box) native() *C.GtkBox {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toGtkBox(ptr)
}

func marshalBox(p uintptr) (interface{}, error) {
	c := C.g_value_get_object(C.toGValue(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapBox(obj), nil
}

func wrapBox(obj *glib.Object) *Box {
	container := wrapContainer(obj)
	return &Box{*container}
}

func (v *Box) toOrientable() *C.GtkOrientable {
	if v == nil {
		return nil
	}
	return C.toGtkOrientable(unsafe.Pointer(v.Native()))
}

// GetOrientation is a wrapper around C.gtk_orientable_get_orientation() for a GtkBox
func (v *Box) GetOrientation() Orientation {
	return Orientation(C.gtk_orientable_get_orientation(v.toOrientable()))
}

// SetOrientation is a wrapper around C.gtk_orientable_set_orientation() for a GtkBox
func (v *Box) SetOrientation(o Orientation) {
	C.gtk_orientable_set_orientation(v.toOrientable(), C.GtkOrientation(o))
}

// BoxNew is a wrapper around gtk_box_new().
func BoxNew(orientation Orientation, spacing int) (*Box, error) {
	c := C.gtk_box_new(C.GtkOrientation(orientation), C.gint(spacing))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapBox(obj), nil
}

// PackStart is a wrapper around gtk_box_pack_start().
func (v *Box) PackStart(child IWidget, expand, fill bool, padding uint) {
	C.gtk_box_pack_start(v.native(), child.toWidget(), gbool(expand),
		gbool(fill), C.guint(padding))
}

// PackEnd is a wrapper around gtk_box_pack_end().
func (v *Box) PackEnd(child IWidget, expand, fill bool, padding uint) {
	C.gtk_box_pack_end(v.native(), child.toWidget(), gbool(expand),
		gbool(fill), C.guint(padding))
}

// GetHomogeneous is a wrapper around gtk_box_get_homogeneous().
func (v *Box) GetHomogeneous() bool {
	c := C.gtk_box_get_homogeneous(v.native())
	return gobool(c)
}

// SetHomogeneous is a wrapper around gtk_box_set_homogeneous().
func (v *Box) SetHomogeneous(homogeneous bool) {
	C.gtk_box_set_homogeneous(v.native(), gbool(homogeneous))
}

// GetSpacing is a wrapper around gtk_box_get_spacing().
func (v *Box) GetSpacing() int {
	c := C.gtk_box_get_spacing(v.native())
	return int(c)
}

// SetSpacing is a wrapper around gtk_box_set_spacing()
func (v *Box) SetSpacing(spacing int) {
	C.gtk_box_set_spacing(v.native(), C.gint(spacing))
}

// ReorderChild is a wrapper around gtk_box_reorder_child().
func (v *Box) ReorderChild(child IWidget, position int) {
	C.gtk_box_reorder_child(v.native(), child.toWidget(), C.gint(position))
}

// QueryChildPacking is a wrapper around gtk_box_query_child_packing().
func (v *Box) QueryChildPacking(child IWidget) (expand, fill bool, padding uint, packType PackType) {
	var cexpand, cfill C.gboolean
	var cpadding C.guint
	var cpackType C.GtkPackType

	C.gtk_box_query_child_packing(v.native(), child.toWidget(), &cexpand,
		&cfill, &cpadding, &cpackType)
	return gobool(cexpand), gobool(cfill), uint(cpadding), PackType(cpackType)
}

// SetChildPacking is a wrapper around gtk_box_set_child_packing().
func (v *Box) SetChildPacking(child IWidget, expand, fill bool, padding uint, packType PackType) {
	C.gtk_box_set_child_packing(v.native(), child.toWidget(), gbool(expand),
		gbool(fill), C.guint(padding), C.GtkPackType(packType))
}

/*
 * GtkGrid
 */

// Grid is a representation of GTK's GtkGrid.
type Grid struct {
	Container
	// Interfaces
	Orientable
}

// native returns a pointer to the underlying GtkGrid.
func (v *Grid) native() *C.GtkGrid {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toGtkGrid(ptr)
}

func (v *Grid) toOrientable() *C.GtkOrientable {
	if v == nil {
		return nil
	}
	return C.toGtkOrientable(unsafe.Pointer(v.Native()))
}

func marshalGrid(p uintptr) (interface{}, error) {
	c := C.g_value_get_object(C.toGValue(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapGrid(obj), nil
}

func wrapGrid(obj *glib.Object) *Grid {
	container := wrapContainer(obj)
	o := wrapOrientable(*glib.InterfaceFromObjectNew(obj))
	return &Grid{*container, *o}
}

// GridNew is a wrapper around gtk_grid_new().
func GridNew() (*Grid, error) {
	c := C.gtk_grid_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapGrid(obj), nil
}

// Attach is a wrapper around gtk_grid_attach().
func (v *Grid) Attach(child IWidget, left, top, width, height int) {
	C.gtk_grid_attach(v.native(), child.toWidget(), C.gint(left),
		C.gint(top), C.gint(width), C.gint(height))
}

// AttachNextTo is a wrapper around gtk_grid_attach_next_to().
func (v *Grid) AttachNextTo(child, sibling IWidget, side PositionType, width, height int) {
	C.gtk_grid_attach_next_to(v.native(), child.toWidget(),
		sibling.toWidget(), C.GtkPositionType(side), C.gint(width),
		C.gint(height))
}

// GetChildAt is a wrapper around gtk_grid_get_child_at().
func (v *Grid) GetChildAt(left, top int) (*Widget, error) {
	c := C.gtk_grid_get_child_at(v.native(), C.gint(left), C.gint(top))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapWidget(obj), nil
}

// InsertRow is a wrapper around gtk_grid_insert_row().
func (v *Grid) InsertRow(position int) {
	C.gtk_grid_insert_row(v.native(), C.gint(position))
}

// InsertColumn is a wrapper around gtk_grid_insert_column().
func (v *Grid) InsertColumn(position int) {
	C.gtk_grid_insert_column(v.native(), C.gint(position))
}

// InsertNextTo is a wrapper around gtk_grid_insert_next_to()
func (v *Grid) InsertNextTo(sibling IWidget, side PositionType) {
	C.gtk_grid_insert_next_to(v.native(), sibling.toWidget(),
		C.GtkPositionType(side))
}

// SetRowHomogeneous is a wrapper around gtk_grid_set_row_homogeneous().
func (v *Grid) SetRowHomogeneous(homogeneous bool) {
	C.gtk_grid_set_row_homogeneous(v.native(), gbool(homogeneous))
}

// GetRowHomogeneous is a wrapper around gtk_grid_get_row_homogeneous().
func (v *Grid) GetRowHomogeneous() bool {
	c := C.gtk_grid_get_row_homogeneous(v.native())
	return gobool(c)
}

// SetRowSpacing is a wrapper around gtk_grid_set_row_spacing().
func (v *Grid) SetRowSpacing(spacing uint) {
	C.gtk_grid_set_row_spacing(v.native(), C.guint(spacing))
}

// GetRowSpacing is a wrapper around gtk_grid_get_row_spacing().
func (v *Grid) GetRowSpacing() uint {
	c := C.gtk_grid_get_row_spacing(v.native())
	return uint(c)
}

// SetColumnHomogeneous is a wrapper around gtk_grid_set_column_homogeneous().
func (v *Grid) SetColumnHomogeneous(homogeneous bool) {
	C.gtk_grid_set_column_homogeneous(v.native(), gbool(homogeneous))
}

// GetColumnHomogeneous is a wrapper around gtk_grid_get_column_homogeneous().
func (v *Grid) GetColumnHomogeneous() bool {
	c := C.gtk_grid_get_column_homogeneous(v.native())
	return gobool(c)
}

// SetColumnSpacing is a wrapper around gtk_grid_set_column_spacing().
func (v *Grid) SetColumnSpacing(spacing uint) {
	C.gtk_grid_set_column_spacing(v.native(), C.guint(spacing))
}

// GetColumnSpacing is a wrapper around gtk_grid_get_column_spacing().
func (v *Grid) GetColumnSpacing() uint {
	c := C.gtk_grid_get_column_spacing(v.native())
	return uint(c)
}

/*
 * GtkOverlay
 */

// Overlay is a representation of GTK's GtkOverlay.
type Overlay struct {
	Bin
}

// native returns a pointer to the underlying GtkOverlay.
func (v *Overlay) native() *C.GtkOverlay {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toGtkOverlay(ptr)
}

func marshalOverlay(p uintptr) (interface{}, error) {
	c := C.g_value_get_object(C.toGValue(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapOverlay(obj), nil
}

func wrapOverlay(obj *glib.Object) *Overlay {
	bin := wrapBin(obj)
	return &Overlay{*bin}
}

// OverlayNew is a wrapper around gtk_overlay_new().
func OverlayNew() (*Overlay, error) {
	c := C.gtk_overlay_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapOverlay(obj), nil
}

// AddOverlay is a wrapper around gtk_overlay_add_overlay().
func (v *Overlay) AddOverlay(widget IWidget) {
	C.gtk_overlay_add_overlay(v.native(), widget.toWidget())
}

// ButtonBoxStyle is a representation of GTK's GtkButtonBoxStyle.
type ButtonBoxStyle int

const (
	BUTTONBOX_SPREAD ButtonBoxStyle = C.GTK_BUTTONBOX_SPREAD
	BUTTONBOX_EDGE   ButtonBoxStyle = C.GTK_BUTTONBOX_EDGE
	BUTTONBOX_START  ButtonBoxStyle = C.GTK_BUTTONBOX_START
	BUTTONBOX_END    ButtonBoxStyle = C.GTK_BUTTONBOX_END
	BUTTONBOX_CENTER ButtonBoxStyle = C.GTK_BUTTONBOX_CENTER
)

func marshalButtonBoxStyle(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum(C.toGValue(unsafe.Pointer(p)))
	return ButtonBoxStyle(c), nil
}

/*
 * GtkButtonBox
 */

// ButtonBox is a representation of GTK's GtkButtonBox
type ButtonBox struct {
	Box
}

// native returns a pointer to underlying GtkButtonBox
func (v *ButtonBox) native() *C.GtkButtonBox {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toGtkButtonBox(ptr)
}

func marshalButtonBox(p uintptr) (interface{}, error) {
	c := C.g_value_get_object(C.toGValue(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapButtonBox(obj), nil
}

func wrapButtonBox(obj *glib.Object) *ButtonBox {
	box := wrapBox(obj)
	return &ButtonBox{*box}
}

// ButtonBoxNew is a wrapper around gtk_button_box_new().
func ButtonBoxNew(orientation Orientation) (*ButtonBox, error) {
	c := C.gtk_button_box_new(C.GtkOrientation(orientation))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapButtonBox(obj), nil
}

// SetLayout is a wrapper around gtk_button_box_set_layout().
func (v *ButtonBox) SetLayout(layoutStyle ButtonBoxStyle) {
	C.gtk_button_box_set_layout(v.native(), C.GtkButtonBoxStyle(layoutStyle))
}

// GetLayout is a wrapper around gtk_button_box_get_layout().
func (v *ButtonBox) GetLayout() ButtonBoxStyle {
	c := C.gtk_button_box_get_layout(v.native())
	return ButtonBoxStyle(c)
}

/*
 * GtkPaned
 */

// Paned is a representation of GTK's GtkPaned.
type Paned struct {
	Bin
}

// native returns a pointer to the underlying GtkPaned.
func (v *Paned) native() *C.GtkPaned {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toGtkPaned(ptr)
}

func marshalPaned(p uintptr) (interface{}, error) {
	c := C.g_value_get_object(C.toGValue(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapPaned(obj), nil
}

func wrapPaned(obj *glib.Object) *Paned {
	bin := wrapBin(obj)
	return &Paned{*bin}
}

// PanedNew is a wrapper around gtk_paned_new().
func PanedNew(orientation Orientation) (*Paned, error) {
	c := C.gtk_paned_new(C.GtkOrientation(orientation))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapPaned(obj), nil
}

// Add1 is a wrapper around gtk_paned_add1().
func (v *Paned) Add1(child IWidget) {
	C.gtk_paned_add1(v.native(), child.toWidget())
}

// Add2 is a wrapper around gtk_paned_add2().
func (v *Paned) Add2(child IWidget) {
	C.gtk_paned_add2(v.native(), child.toWidget())
}

// Pack1 is a wrapper around gtk_paned_pack1().
func (v *Paned) Pack1(child IWidget, resize, shrink bool) {
	C.gtk_paned_pack1(v.native(), child.toWidget(), gbool(resize), gbool(shrink))
}

// Pack2 is a wrapper around gtk_paned_pack2().
func (v *Paned) Pack2(child IWidget, resize, shrink bool) {
	C.gtk_paned_pack2(v.native(), child.toWidget(), gbool(resize), gbool(shrink))
}

// SetPosition is a wrapper around gtk_paned_set_position().
func (v *Paned) SetPosition(position int) {
	C.gtk_paned_set_position(v.native(), C.gint(position))
}

// GetChild1 is a wrapper around gtk_paned_get_child1().
func (v *Paned) GetChild1() (*Widget, error) {
	c := C.gtk_paned_get_child1(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapWidget(obj), nil
}

// GetChild2 is a wrapper around gtk_paned_get_child2().
func (v *Paned) GetChild2() (*Widget, error) {
	c := C.gtk_paned_get_child2(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapWidget(obj), nil
}

// GetHandleWindow is a wrapper around gtk_paned_get_handle_window().
func (v *Paned) GetHandleWindow() (*Window, error) {
	c := C.gtk_paned_get_handle_window(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapWindow(obj), nil
}

// GetPosition is a wrapper around gtk_paned_get_position().
func (v *Paned) GetPosition() int {
	return int(C.gtk_paned_get_position(v.native()))
}

// added by terrak
/*
 * GtkLayout
 */

// Layout is a representation of GTK's GtkLayout.
type Layout struct {
	Container
}

// native returns a pointer to the underlying GtkDrawingArea.
func (v *Layout) native() *C.GtkLayout {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toGtkLayout(ptr)
}

func marshalLayout(p uintptr) (interface{}, error) {
	c := C.g_value_get_object(C.toGValue(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapLayout(obj), nil
}

func wrapLayout(obj *glib.Object) *Layout {
	container := wrapContainer(obj)
	return &Layout{*container}
}

// LayoutNew is a wrapper around gtk_layout_new().
func LayoutNew(hadjustment, vadjustment *Adjustment) (*Layout, error) {
	c := C.gtk_layout_new(hadjustment.native(), vadjustment.native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapLayout(obj), nil
}

// Put is a wrapper around gtk_layout_put().
func (v *Layout) Put(w IWidget, x, y int) {
	C.gtk_layout_put(v.native(), w.toWidget(), C.gint(x), C.gint(y))
}

// Move is a wrapper around gtk_layout_move().
func (v *Layout) Move(w IWidget, x, y int) {
	C.gtk_layout_move(v.native(), w.toWidget(), C.gint(x), C.gint(y))
}

// SetSize is a wrapper around gtk_layout_set_size
func (v *Layout) SetSize(width, height uint) {
	C.gtk_layout_set_size(v.native(), C.guint(width), C.guint(height))
}

// GetSize is a wrapper around gtk_layout_get_size
func (v *Layout) GetSize() (width, height uint) {
	var w, h C.guint
	C.gtk_layout_get_size(v.native(), &w, &h)
	return uint(w), uint(h)
}

/*
 * GtkNotebook
 */

// Notebook is a representation of GTK's GtkNotebook.
type Notebook struct {
	Container
}

// native returns a pointer to the underlying GtkNotebook.
func (v *Notebook) native() *C.GtkNotebook {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toGtkNotebook(ptr)
}

func marshalNotebook(p uintptr) (interface{}, error) {
	c := C.g_value_get_object(C.toGValue(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapNotebook(obj), nil
}

func wrapNotebook(obj *glib.Object) *Notebook {
	container := wrapContainer(obj)
	return &Notebook{*container}
}

// NotebookNew is a wrapper around gtk_notebook_new().
func NotebookNew() (*Notebook, error) {
	c := C.gtk_notebook_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapNotebook(obj), nil
}

// AppendPage is a wrapper around gtk_notebook_append_page().
func (v *Notebook) AppendPage(child IWidget, tabLabel IWidget) int {
	c := C.gtk_notebook_append_page(v.native(), child.toWidget(),
		tabLabel.toWidget())
	return int(c)
}

// AppendPageMenu is a wrapper around gtk_notebook_append_page_menu().
func (v *Notebook) AppendPageMenu(child IWidget, tabLabel IWidget, menuLabel IWidget) int {
	c := C.gtk_notebook_append_page_menu(v.native(), child.toWidget(),
		tabLabel.toWidget(), menuLabel.toWidget())
	return int(c)
}

// PrependPage is a wrapper around gtk_notebook_prepend_page().
func (v *Notebook) PrependPage(child IWidget, tabLabel IWidget) int {
	c := C.gtk_notebook_prepend_page(v.native(), child.toWidget(),
		tabLabel.toWidget())
	return int(c)
}

// PrependPageMenu is a wrapper around gtk_notebook_prepend_page_menu().
func (v *Notebook) PrependPageMenu(child IWidget, tabLabel IWidget, menuLabel IWidget) int {
	c := C.gtk_notebook_prepend_page_menu(v.native(), child.toWidget(),
		tabLabel.toWidget(), menuLabel.toWidget())
	return int(c)
}

// InsertPage is a wrapper around gtk_notebook_insert_page().
func (v *Notebook) InsertPage(child IWidget, tabLabel IWidget, position int) int {
	c := C.gtk_notebook_insert_page(v.native(), child.toWidget(),
		tabLabel.toWidget(), C.gint(position))
	return int(c)
}

// InsertPageMenu is a wrapper around gtk_notebook_insert_page_menu().
func (v *Notebook) InsertPageMenu(child IWidget, tabLabel IWidget, menuLabel IWidget, position int) int {
	c := C.gtk_notebook_insert_page_menu(v.native(), child.toWidget(),
		tabLabel.toWidget(), menuLabel.toWidget(), C.gint(position))
	return int(c)
}

// RemovePage is a wrapper around gtk_notebook_remove_page().
func (v *Notebook) RemovePage(pageNum int) {
	C.gtk_notebook_remove_page(v.native(), C.gint(pageNum))
}

// PageNum is a wrapper around gtk_notebook_page_num().
func (v *Notebook) PageNum(child IWidget) int {
	c := C.gtk_notebook_page_num(v.native(), child.toWidget())
	return int(c)
}

// NextPage is a wrapper around gtk_notebook_next_page().
func (v *Notebook) NextPage() {
	C.gtk_notebook_next_page(v.native())
}

// PrevPage is a wrapper around gtk_notebook_prev_page().
func (v *Notebook) PrevPage() {
	C.gtk_notebook_prev_page(v.native())
}

// ReorderChild is a wrapper around gtk_notebook_reorder_child().
func (v *Notebook) ReorderChild(child IWidget, position int) {
	C.gtk_notebook_reorder_child(v.native(), child.toWidget(),
		C.gint(position))
}

// SetTabPos is a wrapper around gtk_notebook_set_tab_pos().
func (v *Notebook) SetTabPos(pos PositionType) {
	C.gtk_notebook_set_tab_pos(v.native(), C.GtkPositionType(pos))
}

// SetShowTabs is a wrapper around gtk_notebook_set_show_tabs().
func (v *Notebook) SetShowTabs(showTabs bool) {
	C.gtk_notebook_set_show_tabs(v.native(), gbool(showTabs))
}

// SetShowBorder is a wrapper around gtk_notebook_set_show_border().
func (v *Notebook) SetShowBorder(showBorder bool) {
	C.gtk_notebook_set_show_border(v.native(), gbool(showBorder))
}

// SetScrollable is a wrapper around gtk_notebook_set_scrollable().
func (v *Notebook) SetScrollable(scrollable bool) {
	C.gtk_notebook_set_scrollable(v.native(), gbool(scrollable))
}

// PopupEnable is a wrapper around gtk_notebook_popup_enable().
func (v *Notebook) PopupEnable() {
	C.gtk_notebook_popup_enable(v.native())
}

// PopupDisable is a wrapper around gtk_notebook_popup_disable().
func (v *Notebook) PopupDisable() {
	C.gtk_notebook_popup_disable(v.native())
}

// GetCurrentPage is a wrapper around gtk_notebook_get_current_page().
func (v *Notebook) GetCurrentPage() int {
	c := C.gtk_notebook_get_current_page(v.native())
	return int(c)
}

// GetMenuLabel is a wrapper around gtk_notebook_get_menu_label().
func (v *Notebook) GetMenuLabel(child IWidget) (*Widget, error) {
	c := C.gtk_notebook_get_menu_label(v.native(), child.toWidget())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapWidget(obj), nil
}

// GetNthPage is a wrapper around gtk_notebook_get_nth_page().
func (v *Notebook) GetNthPage(pageNum int) (*Widget, error) {
	c := C.gtk_notebook_get_nth_page(v.native(), C.gint(pageNum))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapWidget(obj), nil
}

// GetNPages is a wrapper around gtk_notebook_get_n_pages().
func (v *Notebook) GetNPages() int {
	c := C.gtk_notebook_get_n_pages(v.native())
	return int(c)
}

// GetTabLabel is a wrapper around gtk_notebook_get_tab_label().
func (v *Notebook) GetTabLabel(child IWidget) (*Widget, error) {
	c := C.gtk_notebook_get_tab_label(v.native(), child.toWidget())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapWidget(obj), nil
}

// SetMenuLabel is a wrapper around gtk_notebook_set_menu_label().
func (v *Notebook) SetMenuLabel(child, menuLabel IWidget) {
	C.gtk_notebook_set_menu_label(v.native(), child.toWidget(),
		menuLabel.toWidget())
}

// SetMenuLabelText is a wrapper around gtk_notebook_set_menu_label_text().
func (v *Notebook) SetMenuLabelText(child IWidget, menuText string) {
	cstr := C.CString(menuText)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_notebook_set_menu_label_text(v.native(), child.toWidget(),
		(*C.gchar)(cstr))
}

// SetTabLabel is a wrapper around gtk_notebook_set_tab_label().
func (v *Notebook) SetTabLabel(child, tabLabel IWidget) {
	C.gtk_notebook_set_tab_label(v.native(), child.toWidget(),
		tabLabel.toWidget())
}

// SetTabLabelText is a wrapper around gtk_notebook_set_tab_label_text().
func (v *Notebook) SetTabLabelText(child IWidget, tabText string) {
	cstr := C.CString(tabText)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_notebook_set_tab_label_text(v.native(), child.toWidget(),
		(*C.gchar)(cstr))
}

// SetTabReorderable is a wrapper around gtk_notebook_set_tab_reorderable().
func (v *Notebook) SetTabReorderable(child IWidget, reorderable bool) {
	C.gtk_notebook_set_tab_reorderable(v.native(), child.toWidget(),
		gbool(reorderable))
}

// SetTabDetachable is a wrapper around gtk_notebook_set_tab_detachable().
func (v *Notebook) SetTabDetachable(child IWidget, detachable bool) {
	C.gtk_notebook_set_tab_detachable(v.native(), child.toWidget(),
		gbool(detachable))
}

// GetMenuLabelText is a wrapper around gtk_notebook_get_menu_label_text().
func (v *Notebook) GetMenuLabelText(child IWidget) (string, error) {
	c := C.gtk_notebook_get_menu_label_text(v.native(), child.toWidget())
	if c == nil {
		return "", errors.New("No menu label for widget")
	}
	return goString(c), nil
}

// GetScrollable is a wrapper around gtk_notebook_get_scrollable().
func (v *Notebook) GetScrollable() bool {
	c := C.gtk_notebook_get_scrollable(v.native())
	return gobool(c)
}

// GetShowBorder is a wrapper around gtk_notebook_get_show_border().
func (v *Notebook) GetShowBorder() bool {
	c := C.gtk_notebook_get_show_border(v.native())
	return gobool(c)
}

// GetShowTabs is a wrapper around gtk_notebook_get_show_tabs().
func (v *Notebook) GetShowTabs() bool {
	c := C.gtk_notebook_get_show_tabs(v.native())
	return gobool(c)
}

// GetTabLabelText is a wrapper around gtk_notebook_get_tab_label_text().
func (v *Notebook) GetTabLabelText(child IWidget) (string, error) {
	c := C.gtk_notebook_get_tab_label_text(v.native(), child.toWidget())
	if c == nil {
		return "", errors.New("No tab label for widget")
	}
	return goString(c), nil
}

// GetTabPos is a wrapper around gtk_notebook_get_tab_pos().
func (v *Notebook) GetTabPos() PositionType {
	c := C.gtk_notebook_get_tab_pos(v.native())
	return PositionType(c)
}

// GetTabReorderable is a wrapper around gtk_notebook_get_tab_reorderable().
func (v *Notebook) GetTabReorderable(child IWidget) bool {
	c := C.gtk_notebook_get_tab_reorderable(v.native(), child.toWidget())
	return gobool(c)
}

// GetTabDetachable is a wrapper around gtk_notebook_get_tab_detachable().
func (v *Notebook) GetTabDetachable(child IWidget) bool {
	c := C.gtk_notebook_get_tab_detachable(v.native(), child.toWidget())
	return gobool(c)
}

// SetCurrentPage is a wrapper around gtk_notebook_set_current_page().
func (v *Notebook) SetCurrentPage(pageNum int) {
	C.gtk_notebook_set_current_page(v.native(), C.gint(pageNum))
}

// SetGroupName is a wrapper around gtk_notebook_set_group_name().
func (v *Notebook) SetGroupName(groupName string) {
	cstr := C.CString(groupName)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_notebook_set_group_name(v.native(), (*C.gchar)(cstr))
}

// GetGroupName is a wrapper around gtk_notebook_get_group_name().
func (v *Notebook) GetGroupName() (string, error) {
	c := C.gtk_notebook_get_group_name(v.native())
	if c == nil {
		return "", errors.New("No group name")
	}
	return goString(c), nil
}

// SetActionWidget is a wrapper around gtk_notebook_set_action_widget().
func (v *Notebook) SetActionWidget(widget IWidget, packType PackType) {
	C.gtk_notebook_set_action_widget(v.native(), widget.toWidget(),
		C.GtkPackType(packType))
}

// GetActionWidget is a wrapper around gtk_notebook_get_action_widget().
func (v *Notebook) GetActionWidget(packType PackType) (*Widget, error) {
	c := C.gtk_notebook_get_action_widget(v.native(),
		C.GtkPackType(packType))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapWidget(obj), nil
}

/*
 * GtkExpander
 */

// Expander is a representation of GTK's GtkExpander.
type Expander struct {
	Bin
}

// native returns a pointer to the underlying GtkExpander.
func (v *Expander) native() *C.GtkExpander {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toGtkExpander(ptr)
}

func marshalExpander(p uintptr) (interface{}, error) {
	c := C.g_value_get_object(C.toGValue(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapExpander(obj), nil
}

func wrapExpander(obj *glib.Object) *Expander {
	bin := wrapBin(obj)
	return &Expander{*bin}
}

// ExpanderNew is a wrapper around gtk_expander_new().
func ExpanderNew(label string) (*Expander, error) {
	var cstr *C.gchar
	if label != "" {
		cstr = C.CString(label)
		defer C.free(unsafe.Pointer(cstr))
	}
	c := C.gtk_expander_new((*C.gchar)(cstr))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapExpander(obj), nil
}

// SetExpanded is a wrapper around gtk_expander_set_expanded().
func (v *Expander) SetExpanded(expanded bool) {
	C.gtk_expander_set_expanded(v.native(), gbool(expanded))
}

// GetExpanded is a wrapper around gtk_expander_get_expanded().
func (v *Expander) GetExpanded() bool {
	c := C.gtk_expander_get_expanded(v.native())
	return gobool(c)
}

// SetLabel is a wrapper around gtk_expander_set_label().
func (v *Expander) SetLabel(label string) {
	var cstr *C.char
	if label != "" {
		cstr = C.CString(label)
		defer C.free(unsafe.Pointer(cstr))
	}
	C.gtk_expander_set_label(v.native(), (*C.gchar)(cstr))
}

// GetLabel is a wrapper around gtk_expander_get_label().
func (v *Expander) GetLabel() string {
	c := C.gtk_expander_get_label(v.native())
	return goString(c)
}

// SetLabelWidget is a wrapper around gtk_expander_set_label_widget().
func (v *Expander) SetLabelWidget(widget IWidget) {
	C.gtk_expander_set_label_widget(v.native(), widget.toWidget())
}

// GetLabelWidget is a wrapper around gtk_expander_get_label_widget().
func (v *Expander) GetLabelWidget() (*Widget, error) {
	c := C.gtk_expander_get_label_widget(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapWidget(obj), nil
}

// SetUseUnderline is a wrapper around gtk_expander_set_use_underline().
func (v *Expander) SetUseUnderline(useUnderline bool) {
	C.gtk_expander_set_use_underline(v.native(), gbool(useUnderline))
}

// GetUseUnderline is a wrapper around gtk_expander_get_use_underline().
func (v *Expander) GetUseUnderline() bool {
	c := C.gtk_expander_get_use_underline(v.native())
	return gobool(c)
}

// SetUseMarkup is a wrapper around gtk_expander_set_use_markup().
func (v *Expander) SetUseMarkup(useMarkup bool) {
	C.gtk_expander_set_use_markup(v.native(), gbool(useMarkup))
}

// GetUseMarkup is a wrapper around gtk_expander_get_use_markup().
func (v *Expander) GetUseMarkup() bool {
	c := C.gtk_expander_get_use_markup(v.native())
	return gobool(c)
}

// SetLabelFill is a wrapper around gtk_expander_set_label_fill().
func (v *Expander) SetLabelFill(labelFill bool) {
	C.gtk_expander_set_label_fill(v.native(), gbool(labelFill))
}

// GetLabelFill is a wrapper around gtk_expander_get_label_fill().
func (v *Expander) GetLabelFill() bool {
	c := C.gtk_expander_get_label_fill(v.native())
	return gobool(c)
}

// SetResizeToplevel is a wrapper around gtk_expander_set_resize_toplevel().
func (v *Expander) SetResizeToplevel(resizeToplevel bool) {
	C.gtk_expander_set_resize_toplevel(v.native(), gbool(resizeToplevel))
}

// GetResizeToplevel is a wrapper around gtk_expander_get_resize_toplevel().
func (v *Expander) GetResizeToplevel() bool {
	c := C.gtk_expander_get_resize_toplevel(v.native())
	return gobool(c)
}

// TODO: implement GtkAspectFrame

// TODO: implement GtkFixed
