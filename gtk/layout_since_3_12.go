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

// +build !gtk_3_6,!gtk_3_8,!gtk_3_10
// not use this: go build -tags gtk_3_8'. Otherwise, if no build tags are used, GTK 3.10

package gtk

// #include <gtk/gtk.h>
// #include "gtk.go.h"
// #include "gtk_since_3_12.go.h"
import "C"
import (
	"runtime"
	"unsafe"

	"github.com/romychs/gotk3/glib"
)

const (
	BUTTONBOX_EXPAND ButtonBoxStyle = C.GTK_BUTTONBOX_EXPAND
)

/*
 * FlowBox
 */
type FlowBox struct {
	Container
}

func (v *FlowBox) native() *C.GtkFlowBox {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toGtkFlowBox(ptr)
}

func marshalFlowBox(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapFlowBox(obj), nil
}

func wrapFlowBox(obj *glib.Object) *FlowBox {
	return &FlowBox{Container{Widget{glib.InitiallyUnowned{obj}}}}
}

// FlowBoxNew is a wrapper around gtk_flow_box_new()
func FlowBoxNew() (*FlowBox, error) {
	c := C.gtk_flow_box_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapFlowBox(obj), nil
}

// Insert is a wrapper around gtk_flow_box_insert()
func (v *FlowBox) Insert(widget IWidget, position int) {
	C.gtk_flow_box_insert(v.native(), widget.toWidget(), C.gint(position))
}

// GetChildAtIndex is a wrapper around gtk_flow_box_get_child_at_index()
func (v *FlowBox) GetChildAtIndex(idx int) *FlowBoxChild {
	c := C.gtk_flow_box_get_child_at_index(v.native(), C.gint(idx))
	if c == nil {
		return nil
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapFlowBoxChild(obj)
}

// TODO 3.22.6 gtk_flow_box_get_child_at_pos()

// SetHAdjustment is a wrapper around gtk_flow_box_set_hadjustment()
func (v *FlowBox) SetHAdjustment(adjustment *Adjustment) {
	C.gtk_flow_box_set_hadjustment(v.native(), adjustment.native())
}

// SetVAdjustment is a wrapper around gtk_flow_box_set_vadjustment()
func (v *FlowBox) SetVAdjustment(adjustment *Adjustment) {
	C.gtk_flow_box_set_vadjustment(v.native(), adjustment.native())
}

// SetHomogeneous is a wrapper around gtk_flow_box_set_homogeneous()
func (v *FlowBox) SetHomogeneous(homogeneous bool) {
	C.gtk_flow_box_set_homogeneous(v.native(), gbool(homogeneous))
}

// GetHomogeneous is a wrapper around gtk_flow_box_get_homogeneous()
func (v *FlowBox) GetHomogeneous() bool {
	c := C.gtk_flow_box_get_homogeneous(v.native())
	return gobool(c)
}

// SetRowSpacing is a wrapper around gtk_flow_box_set_row_spacing()
func (v *FlowBox) SetRowSpacing(spacing uint) {
	C.gtk_flow_box_set_row_spacing(v.native(), C.guint(spacing))
}

// GetRowSpacing is a wrapper around gtk_flow_box_get_row_spacing()
func (v *FlowBox) GetRowSpacing() uint {
	c := C.gtk_flow_box_get_row_spacing(v.native())
	return uint(c)
}

// SetColumnSpacing is a wrapper around gtk_flow_box_set_column_spacing()
func (v *FlowBox) SetColumnSpacing(spacing uint) {
	C.gtk_flow_box_set_column_spacing(v.native(), C.guint(spacing))
}

// GetColumnSpacing is a wrapper around gtk_flow_box_get_column_spacing()
func (v *FlowBox) GetColumnSpacing() uint {
	c := C.gtk_flow_box_get_column_spacing(v.native())
	return uint(c)
}

// SetMinChildrenPerLine is a wrapper around gtk_flow_box_set_min_children_per_line()
func (v *FlowBox) SetMinChildrenPerLine(nChildren uint) {
	C.gtk_flow_box_set_min_children_per_line(v.native(), C.guint(nChildren))
}

// GetMinChildrenPerLine is a wrapper around gtk_flow_box_get_min_children_per_line()
func (v *FlowBox) GetMinChildrenPerLine() uint {
	c := C.gtk_flow_box_get_min_children_per_line(v.native())
	return uint(c)
}

// SetMaxChildrenPerLine is a wrapper around gtk_flow_box_set_max_children_per_line()
func (v *FlowBox) SetMaxChildrenPerLine(nChildren uint) {
	C.gtk_flow_box_set_max_children_per_line(v.native(), C.guint(nChildren))
}

// GetMaxChildrenPerLine is a wrapper around gtk_flow_box_get_max_children_per_line()
func (v *FlowBox) GetMaxChildrenPerLine() uint {
	c := C.gtk_flow_box_get_max_children_per_line(v.native())
	return uint(c)
}

// SetActivateOnSingleClick is a wrapper around gtk_flow_box_set_activate_on_single_click()
func (v *FlowBox) SetActivateOnSingleClick(single bool) {
	C.gtk_flow_box_set_activate_on_single_click(v.native(), gbool(single))
}

// GetActivateOnSingleClick gtk_flow_box_get_activate_on_single_click()
func (v *FlowBox) GetActivateOnSingleClick() bool {
	c := C.gtk_flow_box_get_activate_on_single_click(v.native())
	return gobool(c)
}

// TODO: gtk_flow_box_selected_foreach()

// GetSelectedChildren is a wrapper around gtk_flow_box_get_selected_children().
// Returned list is wrapped to return *gtk.FlowBoxChild elements.
func (v *FlowBox) GetSelectedChildren() *glib.List {
	clist := C.gtk_flow_box_get_selected_children(v.native())

	glist := glib.WrapList(uintptr(unsafe.Pointer(clist)))
	glist.DataWrapper(func(ptr unsafe.Pointer) interface{} {
		child := wrapFlowBoxChild(glib.Take(ptr))
		return child
	})

	if glist != nil {
		runtime.SetFinalizer(glist, func(glist *glib.List) {
			glist.Free()
		})
	}

	return glist
}

// SelectChild is a wrapper around gtk_flow_box_select_child()
func (v *FlowBox) SelectChild(child *FlowBoxChild) {
	C.gtk_flow_box_select_child(v.native(), child.native())
}

// UnselectChild is a wrapper around gtk_flow_box_unselect_child()
func (v *FlowBox) UnselectChild(child *FlowBoxChild) {
	C.gtk_flow_box_unselect_child(v.native(), child.native())
}

// SelectAll is a wrapper around gtk_flow_box_select_all()
func (v *FlowBox) SelectAll() {
	C.gtk_flow_box_select_all(v.native())
}

// UnselectAll is a wrapper around gtk_flow_box_unselect_all()
func (v *FlowBox) UnselectAll() {
	C.gtk_flow_box_unselect_all(v.native())
}

// SetSelectionMode is a wrapper around gtk_flow_box_set_selection_mode()
func (v *FlowBox) SetSelectionMode(mode SelectionMode) {
	C.gtk_flow_box_set_selection_mode(v.native(), C.GtkSelectionMode(mode))
}

// GetSelectionMode is a wrapper around gtk_flow_box_get_selection_mode()
func (v *FlowBox) GetSelectionMode() SelectionMode {
	c := C.gtk_flow_box_get_selection_mode(v.native())
	return SelectionMode(c)
}

// TODO gtk_flow_box_set_filter_func()
// TODO gtk_flow_box_invalidate_filter()
// TODO gtk_flow_box_set_sort_func()
// TODO gtk_flow_box_invalidate_sort()
// TODO 3.18 gtk_flow_box_bind_model()

/*
 * FlowBoxChild
 */
type FlowBoxChild struct {
	Bin
}

func (v *FlowBoxChild) native() *C.GtkFlowBoxChild {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toGtkFlowBoxChild(ptr)
}

func marshalFlowBoxChild(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapFlowBoxChild(obj), nil
}

func wrapFlowBoxChild(obj *glib.Object) *FlowBoxChild {
	return &FlowBoxChild{Bin{Container{Widget{glib.InitiallyUnowned{obj}}}}}
}

// FlowBoxChildNew is a wrapper around gtk_flow_box_child_new()
func FlowBoxChildNew() (*FlowBoxChild, error) {
	c := C.gtk_flow_box_child_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapFlowBoxChild(obj), nil
}

// GetIndex is a wrapper around gtk_flow_box_child_get_index()
func (v *FlowBoxChild) GetIndex() int {
	c := C.gtk_flow_box_child_get_index(v.native())
	return int(c)
}

// IsSelected is a wrapper around gtk_flow_box_child_is_selected()
func (v *FlowBoxChild) IsSelected() bool {
	c := C.gtk_flow_box_child_is_selected(v.native())
	return gobool(c)
}

// Changed is a wrapper around gtk_flow_box_child_changed()
func (v *FlowBoxChild) Changed() {
	C.gtk_flow_box_child_changed(v.native())
}

// GtkActionBar
type ActionBar struct {
	Bin
}

func init() {
	tm := []glib.TypeMarshaler{
		{glib.Type(C.gtk_action_bar_get_type()), marshalActionBar},
	}

	glib.RegisterGValueMarshalers(tm)

	//Contribute to casting
	WrapMap["GtkActionBar"] = wrapActionBar
}

func (v *ActionBar) native() *C.GtkActionBar {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toGtkActionBar(ptr)
}

func marshalActionBar(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapActionBar(obj), nil
}

func wrapActionBar(obj *glib.Object) *ActionBar {
	return &ActionBar{Bin{Container{Widget{glib.InitiallyUnowned{obj}}}}}
}

//gtk_action_bar_new()
func ActionBarNew() (*ActionBar, error) {
	c := C.gtk_action_bar_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapActionBar(obj), nil
}

//gtk_action_bar_pack_start(GtkActionBar *action_bar,GtkWidget *child)
func (v *ActionBar) PackStart(child IWidget) {
	C.gtk_action_bar_pack_start(v.native(), child.toWidget())
}

//gtk_action_bar_pack_end(GtkActionBar *action_bar,GtkWidget *child)
func (v *ActionBar) PackEnd(child IWidget) {
	C.gtk_action_bar_pack_end(v.native(), child.toWidget())
}

//gtk_action_bar_set_center_widget(GtkActionBar *action_bar,GtkWidget *center_widget)
func (v *ActionBar) SetCenterWidget(child IWidget) {
	C.gtk_action_bar_set_center_widget(v.native(), child.toWidget())
}

//gtk_action_bar_get_center_widget(GtkActionBar *action_bar)
func (v *ActionBar) GetCenterWidget() *Widget {
	c := C.gtk_action_bar_get_center_widget(v.native())
	if c == nil {
		return nil
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapWidget(obj)
}

/*
 * GtkStack
 */

// GetChildByName is a wrapper around gtk_stack_get_child_by_name().
func (v *Stack) GetChildByName(name string) *Widget {
	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_stack_get_child_by_name(v.native(), (*C.gchar)(cstr))
	if c == nil {
		return nil
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapWidget(obj)
}
