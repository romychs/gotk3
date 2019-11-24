// Same copyright and license as the rest of the files in this project
// This file contains accelerator related functions and structures

package gtk

// #include <gtk/gtk.h>
// #include "gtk.go.h"
import "C"
import (
	"errors"
	"runtime"
	"unsafe"

	"github.com/d2r2/gotk3/gdk"
	"github.com/d2r2/gotk3/glib"
)

/*
 * GtkTreeIter
 */

// TreeIter is a representation of GTK's GtkTreeIter.
type TreeIter struct {
	gtkTreeIter *C.GtkTreeIter
}

// native returns a pointer to the underlying GtkTreeIter.
func (v *TreeIter) native() *C.GtkTreeIter {
	if v == nil {
		return nil
	}
	return v.gtkTreeIter
}

func marshalTreeIter(p uintptr) (interface{}, error) {
	c := C.g_value_get_boxed(C.toGValue(unsafe.Pointer(p)))
	c2 := (*C.GtkTreeIter)(unsafe.Pointer(c))
	return wrapTreeIter(c2), nil
}

func wrapTreeIter(obj *C.GtkTreeIter) *TreeIter {
	return &TreeIter{obj}
}

func (v *TreeIter) free() {
	C.gtk_tree_iter_free(v.native())
}

// Copy is a wrapper around gtk_tree_iter_copy().
func (v *TreeIter) Copy() (*TreeIter, error) {
	c := C.gtk_tree_iter_copy(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	ti := wrapTreeIter(c)
	runtime.SetFinalizer(ti, (*TreeIter).free)
	return ti, nil
}

/*
 * GtkTreeModel
 */

// TreeModel is a representation of GTK's GtkTreeModel GInterface.
type TreeModel struct {
	glib.Interface
}

// ITreeModel is an interface type implemented by all structs
// embedding a TreeModel.  It is meant to be used as an argument type
// for wrapper functions that wrap around a C GTK function taking a
// GtkTreeModel.
type ITreeModel interface {
	toTreeModel() *C.GtkTreeModel
}

// native returns a pointer to the underlying GObject as a GtkTreeModel.
func (v *TreeModel) native() *C.GtkTreeModel {
	return C.toGtkTreeModel(unsafe.Pointer(v.Native()))
}

func (v *TreeModel) toTreeModel() *C.GtkTreeModel {
	if v == nil {
		return nil
	}
	return v.native()
}

func marshalTreeModel(p uintptr) (interface{}, error) {
	c := C.g_value_get_object(C.toGValue(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapTreeModel(*glib.InterfaceFromObjectNew(obj)), nil
}

func wrapTreeModel(intf glib.Interface) *TreeModel {
	return &TreeModel{intf}
}

// GetFlags is a wrapper around gtk_tree_model_get_flags().
func (v *TreeModel) GetFlags() TreeModelFlags {
	c := C.gtk_tree_model_get_flags(v.native())
	return TreeModelFlags(c)
}

// GetNColumns is a wrapper around gtk_tree_model_get_n_columns().
func (v *TreeModel) GetNColumns() int {
	c := C.gtk_tree_model_get_n_columns(v.native())
	return int(c)
}

// GetColumnType is a wrapper around gtk_tree_model_get_column_type().
func (v *TreeModel) GetColumnType(index int) glib.Type {
	c := C.gtk_tree_model_get_column_type(v.native(), C.gint(index))
	return glib.Type(c)
}

// GetIter is a wrapper around gtk_tree_model_get_iter().
func (v *TreeModel) GetIter(path *TreePath) (*TreeIter, error) {
	var iter C.GtkTreeIter
	c := C.gtk_tree_model_get_iter(v.native(), &iter, path.native())
	if !gobool(c) {
		return nil, errors.New("Unable to set iterator")
	}
	t := wrapTreeIter(&iter)
	return t, nil
}

// GetIterFromString is a wrapper around
// gtk_tree_model_get_iter_from_string().
func (v *TreeModel) GetIterFromString(path string) (*TreeIter, error) {
	var iter C.GtkTreeIter
	cstr := C.CString(path)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_tree_model_get_iter_from_string(v.native(), &iter,
		(*C.gchar)(cstr))
	if !gobool(c) {
		return nil, errors.New("Unable to set iterator")
	}
	t := wrapTreeIter(&iter)
	return t, nil
}

// GetIterFirst is a wrapper around gtk_tree_model_get_iter_first().
func (v *TreeModel) GetIterFirst() (*TreeIter, bool) {
	var iter C.GtkTreeIter
	c := C.gtk_tree_model_get_iter_first(v.native(), &iter)
	if !gobool(c) {
		return nil, false
	}
	t := wrapTreeIter(&iter)
	return t, true
}

// GetPath is a wrapper around gtk_tree_model_get_path().
func (v *TreeModel) GetPath(iter *TreeIter) (*TreePath, error) {
	c := C.gtk_tree_model_get_path(v.native(), iter.native())
	if c == nil {
		return nil, nilPtrErr
	}
	p := wrapTreePath(c)
	runtime.SetFinalizer(p, (*TreePath).free)
	return p, nil
}

// GetValue is a wrapper around gtk_tree_model_get_value().
func (v *TreeModel) GetValue(iter *TreeIter, column int) (*glib.Value, error) {
	val, err := glib.ValueAlloc()
	if err != nil {
		return nil, err
	}
	C.gtk_tree_model_get_value(
		v.native(),
		iter.native(),
		C.gint(column),
		C.toGValue(unsafe.Pointer(val.Native())))
	return val, nil
}

// IterNext is a wrapper around gtk_tree_model_iter_next().
func (v *TreeModel) IterNext(iter *TreeIter) bool {
	c := C.gtk_tree_model_iter_next(v.native(), iter.native())
	return gobool(c)
}

// IterPrevious is a wrapper around gtk_tree_model_iter_previous().
func (v *TreeModel) IterPrevious(iter *TreeIter) bool {
	c := C.gtk_tree_model_iter_previous(v.native(), iter.native())
	return gobool(c)
}

// IterNthChild is a wrapper around gtk_tree_model_iter_nth_child().
func (v *TreeModel) IterNthChild(iter *TreeIter, parent *TreeIter, n int) bool {
	c := C.gtk_tree_model_iter_nth_child(v.native(), iter.native(), parent.native(), C.gint(n))
	return gobool(c)
}

// IterChildren is a wrapper around gtk_tree_model_iter_children().
func (v *TreeModel) IterChildren(iter, child *TreeIter) bool {
	var cIter, cChild *C.GtkTreeIter
	if iter != nil {
		cIter = iter.native()
	}
	cChild = child.native()
	c := C.gtk_tree_model_iter_children(v.native(), cChild, cIter)
	return gobool(c)
}

// IterNChildren is a wrapper around gtk_tree_model_iter_n_children().
func (v *TreeModel) IterNChildren(iter *TreeIter) int {
	var cIter *C.GtkTreeIter
	if iter != nil {
		cIter = iter.native()
	}
	c := C.gtk_tree_model_iter_n_children(v.native(), cIter)
	return int(c)
}

// FilterNew is a wrapper around gtk_tree_model_filter_new().
func (v *TreeModel) FilterNew(root *TreePath) (*TreeModelFilter, error) {
	c := C.gtk_tree_model_filter_new(v.native(), root.native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapTreeModelFilter(obj), nil
}

/*
 * GtkTreeModelFilter
 */

// TreeModelFilter is a representation of GTK's GtkTreeModelFilter.
type TreeModelFilter struct {
	*glib.Object
	// Interfaces
	TreeModel
}

func (v *TreeModelFilter) native() *C.GtkTreeModelFilter {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toGtkTreeModelFilter(ptr)
}

func (v *TreeModelFilter) toTreeModelFilter() *C.GtkTreeModelFilter {
	if v == nil {
		return nil
	}
	return v.native()
}

func marshalTreeModelFilter(p uintptr) (interface{}, error) {
	c := C.g_value_get_object(C.toGValue(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapTreeModelFilter(obj), nil
}

func wrapTreeModelFilter(obj *glib.Object) *TreeModelFilter {
	tm := wrapTreeModel(*glib.InterfaceFromObjectNew(obj))
	return &TreeModelFilter{obj, *tm}
}

// SetVisibleColumn is a wrapper around gtk_tree_model_filter_set_visible_column().
func (v *TreeModelFilter) SetVisibleColumn(column int) {
	C.gtk_tree_model_filter_set_visible_column(v.native(), C.gint(column))
}

/*
 * GtkTreePath
 */

// TreePath is a representation of GTK's GtkTreePath.
type TreePath struct {
	gtkTreePath *C.GtkTreePath
}

// TreePathFromList return a TreePath from the GList
func TreePathFromList(list *glib.List) *TreePath {
	if list == nil {
		return nil
	}
	ptr := (list.Data()).(unsafe.Pointer)
	p := wrapTreePath((*C.GtkTreePath)(ptr))

	return p
}

// native returns a pointer to the underlying GtkTreePath.
func (v *TreePath) native() *C.GtkTreePath {
	if v == nil {
		return nil
	}
	return v.gtkTreePath
}

func marshalTreePath(p uintptr) (interface{}, error) {
	c := C.g_value_get_boxed(C.toGValue(unsafe.Pointer(p)))
	c2 := (*C.GtkTreePath)(unsafe.Pointer(c))
	return wrapTreePath(c2), nil
}

func wrapTreePath(obj *C.GtkTreePath) *TreePath {
	return &TreePath{obj}
}

func (v *TreePath) free() {
	C.gtk_tree_path_free(v.native())
}

// GetIndices is a wrapper around gtk_tree_path_get_indices_with_depth
func (v *TreePath) GetIndices() []int {
	var depth C.gint
	var goindices []int
	var ginthelp C.gint
	indices := uintptr(unsafe.Pointer(C.gtk_tree_path_get_indices_with_depth(v.native(), &depth)))
	size := unsafe.Sizeof(ginthelp)
	for i := 0; i < int(depth); i++ {
		goind := int(*((*C.gint)(unsafe.Pointer(indices))))
		goindices = append(goindices, goind)
		indices += size
	}
	return goindices
}

// String is a wrapper around gtk_tree_path_to_string().
func (v *TreePath) String() string {
	c := C.gtk_tree_path_to_string(v.native())
	return goString(c)
}

// TreePathNewFromString is a wrapper around gtk_tree_path_new_from_string().
func TreePathNewFromString(path string) (*TreePath, error) {
	cstr := C.CString(path)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_tree_path_new_from_string((*C.gchar)(cstr))
	if c == nil {
		return nil, nilPtrErr
	}
	t := wrapTreePath(c)
	runtime.SetFinalizer(t, func(t *TreePath) {
		t.free()
	})
	return t, nil
}

/*
 * GtkTreeSelection
 */

// TreeSelection is a representation of GTK's GtkTreeSelection.
type TreeSelection struct {
	*glib.Object
}

// native returns a pointer to the underlying GtkTreeSelection.
func (v *TreeSelection) native() *C.GtkTreeSelection {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toGtkTreeSelection(ptr)
}

func marshalTreeSelection(p uintptr) (interface{}, error) {
	c := C.g_value_get_object(C.toGValue(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapTreeSelection(obj), nil
}

func wrapTreeSelection(obj *glib.Object) *TreeSelection {
	return &TreeSelection{obj}
}

// GetSelected is a wrapper around gtk_tree_selection_get_selected().
func (v *TreeSelection) GetSelected() (model ITreeModel, iter *TreeIter, ok bool) {
	var cmodel *C.GtkTreeModel
	var citer C.GtkTreeIter
	c := C.gtk_tree_selection_get_selected(v.native(),
		&cmodel, &citer)
	obj := glib.Take(unsafe.Pointer(cmodel))
	model = wrapTreeModel(*glib.InterfaceFromObjectNew(obj))
	iter = wrapTreeIter(&citer)
	ok = gobool(c)
	return
}

// SelectPath is a wrapper around gtk_tree_selection_select_path().
func (v *TreeSelection) SelectPath(path *TreePath) {
	C.gtk_tree_selection_select_path(v.native(), path.native())
}

// UnselectPath is a wrapper around gtk_tree_selection_unselect_path().
func (v *TreeSelection) UnselectPath(path *TreePath) {
	C.gtk_tree_selection_unselect_path(v.native(), path.native())
}

// GetSelectedRows is a wrapper around gtk_tree_selection_get_selected_rows().
// All the elements of returned list are wrapped into (*gtk.TreePath) values.
//
// Please note that a runtime finalizer is only set on the head of the linked
// list, and must be kept live while accessing any item in the list, or the
// Go garbage collector will free the whole list.
func (v *TreeSelection) GetSelectedRows(model ITreeModel) *glib.List {
	var pcmodel **C.GtkTreeModel
	if model != nil {
		cmodel := model.toTreeModel()
		pcmodel = &cmodel
	}

	clist := C.gtk_tree_selection_get_selected_rows(v.native(), pcmodel)
	if clist == nil {
		return nil
	}

	glist := glib.WrapList(uintptr(unsafe.Pointer(clist)))
	glist.DataWrapper(func(ptr unsafe.Pointer) interface{} {
		p := wrapTreePath((*C.GtkTreePath)(ptr))
		return p
	})

	if glist != nil {
		runtime.SetFinalizer(glist, func(glist *glib.List) {
			glist.FreeFull(func(item interface{}) {
				path := item.(*TreePath)
				path.free()
			})
		})
	}

	return glist
}

// CountSelectedRows is a wrapper around gtk_tree_selection_count_selected_rows().
func (v *TreeSelection) CountSelectedRows() int {
	return int(C.gtk_tree_selection_count_selected_rows(v.native()))
}

// SelectIter is a wrapper around gtk_tree_selection_select_iter().
func (v *TreeSelection) SelectIter(iter *TreeIter) {
	C.gtk_tree_selection_select_iter(v.native(), iter.native())
}

// SetMode is a wrapper around gtk_tree_selection_set_mode().
func (v *TreeSelection) SetMode(m SelectionMode) {
	C.gtk_tree_selection_set_mode(v.native(), C.GtkSelectionMode(m))
}

// GetMode is a wrapper around gtk_tree_selection_get_mode().
func (v *TreeSelection) GetMode() SelectionMode {
	return SelectionMode(C.gtk_tree_selection_get_mode(v.native()))
}

/*
 * GtkTreeStore
 */

// TreeStore is a representation of GTK's GtkTreeStore.
type TreeStore struct {
	*glib.Object
	// Interfaces
	TreeModel
}

// native returns a pointer to the underlying GtkTreeStore.
func (v *TreeStore) native() *C.GtkTreeStore {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toGtkTreeStore(ptr)
}

func marshalTreeStore(p uintptr) (interface{}, error) {
	c := C.g_value_get_object(C.toGValue(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapTreeStore(obj), nil
}

func wrapTreeStore(obj *glib.Object) *TreeStore {
	tm := wrapTreeModel(*glib.InterfaceFromObjectNew(obj))
	return &TreeStore{obj, *tm}
}

func (v *TreeStore) toTreeModel() *C.GtkTreeModel {
	if v == nil {
		return nil
	}
	return C.toGtkTreeModel(unsafe.Pointer(v.Native()))
}

// TreeStoreNew is a wrapper around gtk_tree_store_newv().
func TreeStoreNew(types ...glib.Type) (*TreeStore, error) {
	gtypes := C.alloc_types(C.int(len(types)))
	for n, val := range types {
		C.set_type(gtypes, C.int(n), C.GType(val))
	}
	defer C.g_free(C.gpointer(gtypes))
	c := C.gtk_tree_store_newv(C.gint(len(types)), gtypes)
	if c == nil {
		return nil, nilPtrErr
	}

	obj := glib.Take(unsafe.Pointer(c))
	return wrapTreeStore(obj), nil
}

// Append is a wrapper around gtk_tree_store_append().
func (v *TreeStore) Append(parent *TreeIter) *TreeIter {
	var ti C.GtkTreeIter
	var cParent *C.GtkTreeIter
	if parent != nil {
		cParent = parent.native()
	}
	C.gtk_tree_store_append(v.native(), &ti, cParent)
	iter := wrapTreeIter(&ti)
	return iter
}

// Insert is a wrapper around gtk_tree_store_insert
func (v *TreeStore) Insert(parent *TreeIter, position int) *TreeIter {
	var ti C.GtkTreeIter
	var cParent *C.GtkTreeIter
	if parent != nil {
		cParent = parent.native()
	}
	C.gtk_tree_store_insert(v.native(), &ti, cParent, C.gint(position))
	iter := wrapTreeIter(&ti)
	return iter
}

// SetValue is a wrapper around gtk_tree_store_set_value()
func (v *TreeStore) SetValue(iter *TreeIter, column int, value interface{}) error {
	switch value.(type) {
	case *gdk.Pixbuf:
		pix := value.(*gdk.Pixbuf)
		C._gtk_tree_store_set(v.native(), iter.native(), C.gint(column), unsafe.Pointer(pix.Native()))

	default:
		gv, err := glib.GValue(value)
		if err != nil {
			return err
		}
		C.gtk_tree_store_set_value(v.native(), iter.native(),
			C.gint(column),
			C.toGValue(gv.Native()))
	}
	return nil
}

// Remove is a wrapper around gtk_tree_store_remove().
func (v *TreeStore) Remove(iter *TreeIter) bool {
	var ti *C.GtkTreeIter
	if iter != nil {
		ti = iter.native()
	}
	return 0 != C.gtk_tree_store_remove(v.native(), ti)
}

// Clear is a wrapper around gtk_tree_store_clear().
func (v *TreeStore) Clear() {
	C.gtk_tree_store_clear(v.native())
}

/*
 * GtkCellLayout
 */

// CellLayout is a representation of GTK's GtkCellLayout GInterface.
type CellLayout struct {
	glib.Interface
}

// ICellLayout is an interface type implemented by all structs
// embedding a CellLayout.  It is meant to be used as an argument type
// for wrapper functions that wrap around a C GTK function taking a
// GtkCellLayout.
type ICellLayout interface {
	toCellLayout() *C.GtkCellLayout
}

// native() returns a pointer to the underlying GObject as a GtkCellLayout.
func (v *CellLayout) native() *C.GtkCellLayout {
	return C.toGtkCellLayout(unsafe.Pointer(v.Native()))
}

func marshalCellLayout(p uintptr) (interface{}, error) {
	c := C.g_value_get_object(C.toGValue(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapCellLayout(*glib.InterfaceFromObjectNew(obj)), nil
}

func wrapCellLayout(intf glib.Interface) *CellLayout {
	return &CellLayout{intf}
}

func (v *CellLayout) toCellLayout() *C.GtkCellLayout {
	return v.native()
}

// PackStart is a wrapper around gtk_cell_layout_pack_start().
func (v *CellLayout) PackStart(cell ICellRenderer, expand bool) {
	C.gtk_cell_layout_pack_start(v.native(), cell.toCellRenderer(),
		gbool(expand))
}

// AddAttribute is a wrapper around gtk_cell_layout_add_attribute().
func (v *CellLayout) AddAttribute(cell ICellRenderer, attribute string, column int) {
	cstr := C.CString(attribute)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_cell_layout_add_attribute(v.native(), cell.toCellRenderer(),
		(*C.gchar)(cstr), C.gint(column))
}

/*
 * GtkCellRenderer
 */

// CellRenderer is a representation of GTK's GtkCellRenderer.
type CellRenderer struct {
	glib.InitiallyUnowned
}

// ICellRenderer is an interface type implemented by all structs
// embedding a CellRenderer.  It is meant to be used as an argument type
// for wrapper functions that wrap around a C GTK function taking a
// GtkCellRenderer.
type ICellRenderer interface {
	toCellRenderer() *C.GtkCellRenderer
}

// native returns a pointer to the underlying GtkCellRenderer.
func (v *CellRenderer) native() *C.GtkCellRenderer {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toGtkCellRenderer(ptr)
}

func (v *CellRenderer) toCellRenderer() *C.GtkCellRenderer {
	if v == nil {
		return nil
	}
	return v.native()
}

func marshalCellRenderer(p uintptr) (interface{}, error) {
	c := C.g_value_get_object(C.toGValue(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapCellRenderer(obj), nil
}

func wrapCellRenderer(obj *glib.Object) *CellRenderer {
	return &CellRenderer{glib.InitiallyUnowned{obj}}
}

// SetAlignment is a wrapper around gtk_cell_renderer_set_alignment().
func (v *CellRenderer) SetAlignment(xalign, yalign float32) {
	C.gtk_cell_renderer_set_alignment(v.native(), C.gfloat(xalign), C.gfloat(yalign))
}

// GetAlignment is a wrapper around gtk_cell_renderer_get_alignment().
func (v *CellRenderer) GetAlignment() (xalign, yalign float32) {
	var xal, yal C.gfloat
	C.gtk_cell_renderer_get_alignment(v.native(), &xal, &yal)
	return float32(xal), float32(yal)
}

/*
 * GtkCellRendererSpinner
 */

// CellRendererSpinner is a representation of GTK's GtkCellRendererSpinner.
type CellRendererSpinner struct {
	CellRenderer
}

// native returns a pointer to the underlying GtkCellRendererSpinner.
func (v *CellRendererSpinner) native() *C.GtkCellRendererSpinner {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toGtkCellRendererSpinner(ptr)
}

func marshalCellRendererSpinner(p uintptr) (interface{}, error) {
	c := C.g_value_get_object(C.toGValue(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapCellRendererSpinner(obj), nil
}

func wrapCellRendererSpinner(obj *glib.Object) *CellRendererSpinner {
	cellRenderer := wrapCellRenderer(obj)
	return &CellRendererSpinner{*cellRenderer}
}

// CellRendererSpinnerNew is a wrapper around gtk_cell_renderer_text_new().
func CellRendererSpinnerNew() (*CellRendererSpinner, error) {
	c := C.gtk_cell_renderer_spinner_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapCellRendererSpinner(obj), nil
}

/*
 * GtkCellRendererPixbuf
 */

// CellRendererPixbuf is a representation of GTK's GtkCellRendererPixbuf.
type CellRendererPixbuf struct {
	CellRenderer
}

// native returns a pointer to the underlying GtkCellRendererPixbuf.
func (v *CellRendererPixbuf) native() *C.GtkCellRendererPixbuf {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toGtkCellRendererPixbuf(ptr)
}

func marshalCellRendererPixbuf(p uintptr) (interface{}, error) {
	c := C.g_value_get_object(C.toGValue(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapCellRendererPixbuf(obj), nil
}

func wrapCellRendererPixbuf(obj *glib.Object) *CellRendererPixbuf {
	cellRenderer := wrapCellRenderer(obj)
	return &CellRendererPixbuf{*cellRenderer}
}

// CellRendererPixbufNew is a wrapper around gtk_cell_renderer_pixbuf_new().
func CellRendererPixbufNew() (*CellRendererPixbuf, error) {
	c := C.gtk_cell_renderer_pixbuf_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapCellRendererPixbuf(obj), nil
}

/*
 * GtkCellRendererText
 */

// CellRendererText is a representation of GTK's GtkCellRendererText.
type CellRendererText struct {
	CellRenderer
}

// native returns a pointer to the underlying GtkCellRendererText.
func (v *CellRendererText) native() *C.GtkCellRendererText {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toGtkCellRendererText(ptr)
}

func marshalCellRendererText(p uintptr) (interface{}, error) {
	c := C.g_value_get_object(C.toGValue(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapCellRendererText(obj), nil
}

func wrapCellRendererText(obj *glib.Object) *CellRendererText {
	cellRenderer := wrapCellRenderer(obj)
	return &CellRendererText{*cellRenderer}
}

// CellRendererTextNew is a wrapper around gtk_cell_renderer_text_new().
func CellRendererTextNew() (*CellRendererText, error) {
	c := C.gtk_cell_renderer_text_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapCellRendererText(obj), nil
}

/*
 * GtkCellRendererToggle
 */

// CellRendererToggle is a representation of GTK's GtkCellRendererToggle.
type CellRendererToggle struct {
	CellRenderer
}

// native returns a pointer to the underlying GtkCellRendererToggle.
func (v *CellRendererToggle) native() *C.GtkCellRendererToggle {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toGtkCellRendererToggle(ptr)
}

func (v *CellRendererToggle) toCellRenderer() *C.GtkCellRenderer {
	if v == nil {
		return nil
	}
	return v.CellRenderer.native()
}

func marshalCellRendererToggle(p uintptr) (interface{}, error) {
	c := C.g_value_get_object(C.toGValue(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapCellRendererToggle(obj), nil
}

func wrapCellRendererToggle(obj *glib.Object) *CellRendererToggle {
	cellRenderer := wrapCellRenderer(obj)
	return &CellRendererToggle{*cellRenderer}
}

// CellRendererToggleNew is a wrapper around gtk_cell_renderer_toggle_new().
func CellRendererToggleNew() (*CellRendererToggle, error) {
	c := C.gtk_cell_renderer_toggle_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapCellRendererToggle(obj), nil
}

// SetRadio is a wrapper around gtk_cell_renderer_toggle_set_radio().
func (v *CellRendererToggle) SetRadio(set bool) {
	C.gtk_cell_renderer_toggle_set_radio(v.native(), gbool(set))
}

// GetRadio is a wrapper around gtk_cell_renderer_toggle_get_radio().
func (v *CellRendererToggle) GetRadio() bool {
	c := C.gtk_cell_renderer_toggle_get_radio(v.native())
	return gobool(c)
}

// SetActive is a wrapper around gtk_cell_renderer_toggle_set_active().
func (v *CellRendererToggle) SetActive(active bool) {
	C.gtk_cell_renderer_toggle_set_active(v.native(), gbool(active))
}

// GetActive is a wrapper around gtk_cell_renderer_toggle_get_active().
func (v *CellRendererToggle) GetActive() bool {
	c := C.gtk_cell_renderer_toggle_get_active(v.native())
	return gobool(c)
}

// SetActivatable is a wrapper around gtk_cell_renderer_toggle_set_activatable().
func (v *CellRendererToggle) SetActivatable(activatable bool) {
	C.gtk_cell_renderer_toggle_set_activatable(v.native(),
		gbool(activatable))
}

// GetActivatable is a wrapper around gtk_cell_renderer_toggle_get_activatable().
func (v *CellRendererToggle) GetActivatable() bool {
	c := C.gtk_cell_renderer_toggle_get_activatable(v.native())
	return gobool(c)
}

/*
 * GtkIconView
 */

// IconView is a representation of GTK's GtkIconView.
type IconView struct {
	Container
}

// native returns a pointer to the underlying GtkIconView.
func (v *IconView) native() *C.GtkIconView {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toGtkIconView(ptr)
}

func marshalIconView(p uintptr) (interface{}, error) {
	c := C.g_value_get_object(C.toGValue(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapIconView(obj), nil
}

func wrapIconView(obj *glib.Object) *IconView {
	container := wrapContainer(obj)
	return &IconView{*container}
}

// IconViewNew is a wrapper around gtk_icon_view_new().
func IconViewNew() (*IconView, error) {
	c := C.gtk_icon_view_new()
	if c == nil {
		return nil, nilPtrErr
	}

	obj := glib.Take(unsafe.Pointer(c))
	return wrapIconView(obj), nil
}

// IconViewNewWithModel is a wrapper around gtk_icon_view_new_with_model().
func IconViewNewWithModel(model ITreeModel) (*IconView, error) {
	c := C.gtk_icon_view_new_with_model(model.toTreeModel())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapIconView(obj), nil
}

// GetModel is a wrapper around gtk_icon_view_get_model().
func (v *IconView) GetModel() (*TreeModel, error) {
	c := C.gtk_icon_view_get_model(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapTreeModel(*glib.InterfaceFromObjectNew(obj)), nil
}

// SetModel is a wrapper around gtk_icon_view_set_model().
func (v *IconView) SetModel(model ITreeModel) {
	C.gtk_icon_view_set_model(v.native(), model.toTreeModel())
}

// SelectPath is a wrapper around gtk_icon_view_select_path().
func (v *IconView) SelectPath(path *TreePath) {
	C.gtk_icon_view_select_path(v.native(), path.native())
}

// ScrollToPath is a wrapper around gtk_icon_view_scroll_to_path().
func (v *IconView) ScrollToPath(path *TreePath, useAlign bool, rowAlign, colAlign float32) {
	C.gtk_icon_view_scroll_to_path(v.native(), path.native(), gbool(useAlign),
		C.gfloat(rowAlign), C.gfloat(colAlign))
}

/*
 * GtkListStore
 */

// ListStore is a representation of GTK's GtkListStore.
type ListStore struct {
	*glib.Object
	// Interfaces
	TreeModel
}

// native returns a pointer to the underlying GtkListStore.
func (v *ListStore) native() *C.GtkListStore {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toGtkListStore(ptr)
}

func marshalListStore(p uintptr) (interface{}, error) {
	c := C.g_value_get_object(C.toGValue(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapListStore(obj), nil
}

func wrapListStore(obj *glib.Object) *ListStore {
	tm := wrapTreeModel(*glib.InterfaceFromObjectNew(obj))
	return &ListStore{obj, *tm}
}

func (v *ListStore) toTreeModel() *C.GtkTreeModel {
	if v == nil {
		return nil
	}
	return C.toGtkTreeModel(unsafe.Pointer(v.Native()))
}

// ListStoreNew is a wrapper around gtk_list_store_newv().
func ListStoreNew(types ...glib.Type) (*ListStore, error) {
	gtypes := C.alloc_types(C.int(len(types)))
	for n, val := range types {
		C.set_type(gtypes, C.int(n), C.GType(val))
	}
	defer C.g_free(C.gpointer(gtypes))
	c := C.gtk_list_store_newv(C.gint(len(types)), gtypes)
	if c == nil {
		return nil, nilPtrErr
	}

	obj := glib.Take(unsafe.Pointer(c))
	return wrapListStore(obj), nil
}

// Remove is a wrapper around gtk_list_store_remove().
func (v *ListStore) Remove(iter *TreeIter) bool {
	c := C.gtk_list_store_remove(v.native(), iter.native())
	return gobool(c)
}

// TODO(jrick)
/*
func (v *ListStore) SetColumnTypes(types ...glib.Type) {
}
*/

// Set() is a wrapper around gtk_list_store_set_value() but provides
// a function similar to gtk_list_store_set() in that multiple columns
// may be set by one call.  The length of columns and values slices must
// match, or Set() will return a non-nil error.
//
// As an example, a call to:
//  store.Set(iter, []int{0, 1}, []interface{}{"Foo", "Bar"})
// is functionally equivalent to calling the native C GTK function:
//  gtk_list_store_set(store, iter, 0, "Foo", 1, "Bar", -1);
func (v *ListStore) Set(iter *TreeIter, columns []int, values []interface{}) error {
	if len(columns) != len(values) {
		return errors.New("columns and values lengths do not match")
	}
	for i, val := range values {
		v.SetValue(iter, columns[i], val)
	}
	return nil
}

// SetValue is a wrapper around gtk_list_store_set_value().
func (v *ListStore) SetValue(iter *TreeIter, column int, value interface{}) error {
	switch value.(type) {
	case *gdk.Pixbuf:
		pix := value.(*gdk.Pixbuf)
		C._gtk_list_store_set(v.native(), iter.native(), C.gint(column), unsafe.Pointer(pix.Native()))

	default:
		gv, err := glib.GValue(value)
		if err != nil {
			return err
		}

		C.gtk_list_store_set_value(v.native(), iter.native(),
			C.gint(column),
			C.toGValue(unsafe.Pointer(gv.Native())))
	}

	return nil
}

// func (v *ListStore) Model(model ITreeModel) {
// 	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(model.toTreeModel()))}
//	v.TreeModel = *wrapTreeModel(obj)
//}

// SetSortColumnId is a wrapper around gtk_tree_sortable_set_sort_column_id().
func (v *ListStore) SetSortColumnId(column int, order SortType) {
	sort := C.toGtkTreeSortable(unsafe.Pointer(v.Native()))
	C.gtk_tree_sortable_set_sort_column_id(sort, C.gint(column), C.GtkSortType(order))
}

func (v *ListStore) SetCols(iter *TreeIter, cols Cols) error {
	for key, value := range cols {
		err := v.SetValue(iter, key, value)
		if err != nil {
			return err
		}
	}
	return nil
}

// Convenient map for Columns and values (See ListStore, TreeStore)
type Cols map[int]interface{}

// TODO(jrick)
/*
func (v *ListStore) InsertWithValues(iter *TreeIter, position int, columns []int, values []glib.Value) {
		var ccolumns *C.gint
		var cvalues *C.GValue

		C.gtk_list_store_insert_with_values(v.native(), iter.native(),
			C.gint(position), columns, values, C.gint(len(values)))
}
*/

// InsertBefore is a wrapper around gtk_list_store_insert_before().
func (v *ListStore) InsertBefore(sibling *TreeIter) *TreeIter {
	var ti C.GtkTreeIter
	C.gtk_list_store_insert_before(v.native(), &ti, sibling.native())
	iter := wrapTreeIter(&ti)
	return iter
}

// InsertAfter is a wrapper around gtk_list_store_insert_after().
func (v *ListStore) InsertAfter(sibling *TreeIter) *TreeIter {
	var ti C.GtkTreeIter
	C.gtk_list_store_insert_after(v.native(), &ti, sibling.native())
	iter := wrapTreeIter(&ti)
	return iter
}

// Prepend is a wrapper around gtk_list_store_prepend().
func (v *ListStore) Prepend() *TreeIter {
	var ti C.GtkTreeIter
	C.gtk_list_store_prepend(v.native(), &ti)
	iter := wrapTreeIter(&ti)
	return iter
}

// Append is a wrapper around gtk_list_store_append().
func (v *ListStore) Append() *TreeIter {
	var ti C.GtkTreeIter
	C.gtk_list_store_append(v.native(), &ti)
	iter := wrapTreeIter(&ti)
	return iter
}

// Clear is a wrapper around gtk_list_store_clear().
func (v *ListStore) Clear() {
	C.gtk_list_store_clear(v.native())
}

// IterIsValid is a wrapper around gtk_list_store_iter_is_valid().
func (v *ListStore) IterIsValid(iter *TreeIter) bool {
	c := C.gtk_list_store_iter_is_valid(v.native(), iter.native())
	return gobool(c)
}

// TODO(jrick)
/*
func (v *ListStore) Reorder(newOrder []int) {
}
*/

// Swap is a wrapper around gtk_list_store_swap().
func (v *ListStore) Swap(a, b *TreeIter) {
	C.gtk_list_store_swap(v.native(), a.native(), b.native())
}

// MoveBefore is a wrapper around gtk_list_store_move_before().
func (v *ListStore) MoveBefore(iter, position *TreeIter) {
	C.gtk_list_store_move_before(v.native(), iter.native(),
		position.native())
}

// MoveAfter is a wrapper around gtk_list_store_move_after().
func (v *ListStore) MoveAfter(iter, position *TreeIter) {
	C.gtk_list_store_move_after(v.native(), iter.native(),
		position.native())
}

/*
 * GtkTreeViewColumn
 */

// TreeViewColumn is a representation of GTK's GtkTreeViewColumn.
type TreeViewColumn struct {
	glib.InitiallyUnowned
}

// native returns a pointer to the underlying GtkTreeViewColumn.
func (v *TreeViewColumn) native() *C.GtkTreeViewColumn {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toGtkTreeViewColumn(ptr)
}

func marshalTreeViewColumn(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapTreeViewColumn(obj), nil
}

func wrapTreeViewColumn(obj *glib.Object) *TreeViewColumn {
	return &TreeViewColumn{glib.InitiallyUnowned{obj}}
}

// TreeViewColumnNew is a wrapper around gtk_tree_view_column_new().
func TreeViewColumnNew() (*TreeViewColumn, error) {
	c := C.gtk_tree_view_column_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapTreeViewColumn(obj), nil
}

// TreeViewColumnNewWithAttribute is a wrapper around
// gtk_tree_view_column_new_with_attributes() that only sets one
// attribute for one column.
func TreeViewColumnNewWithAttribute(title string, renderer ICellRenderer, attribute string, column int) (*TreeViewColumn, error) {
	t_cstr := C.CString(title)
	defer C.free(unsafe.Pointer(t_cstr))
	a_cstr := C.CString(attribute)
	defer C.free(unsafe.Pointer(a_cstr))
	c := C._gtk_tree_view_column_new_with_attributes_one((*C.gchar)(t_cstr),
		renderer.toCellRenderer(), (*C.gchar)(a_cstr), C.gint(column))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapTreeViewColumn(obj), nil
}

// AddAttribute is a wrapper around gtk_tree_view_column_add_attribute().
func (v *TreeViewColumn) AddAttribute(renderer ICellRenderer, attribute string, column int) {
	cstr := C.CString(attribute)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_tree_view_column_add_attribute(v.native(),
		renderer.toCellRenderer(), (*C.gchar)(cstr), C.gint(column))
}

// SetExpand is a wrapper around gtk_tree_view_column_set_expand().
func (v *TreeViewColumn) SetExpand(expand bool) {
	C.gtk_tree_view_column_set_expand(v.native(), gbool(expand))
}

// GetExpand is a wrapper around gtk_tree_view_column_get_expand().
func (v *TreeViewColumn) GetExpand() bool {
	c := C.gtk_tree_view_column_get_expand(v.native())
	return gobool(c)
}

// SetMinWidth is a wrapper around gtk_tree_view_column_set_min_width().
func (v *TreeViewColumn) SetMinWidth(minWidth int) {
	C.gtk_tree_view_column_set_min_width(v.native(), C.gint(minWidth))
}

// GetMinWidth is a wrapper around gtk_tree_view_column_get_min_width().
func (v *TreeViewColumn) GetMinWidth() int {
	c := C.gtk_tree_view_column_get_min_width(v.native())
	return int(c)
}

// PackStart is a wrapper around gtk_tree_view_column_pack_start().
func (v *TreeViewColumn) PackStart(cell *CellRenderer, expand bool) {
	C.gtk_tree_view_column_pack_start(v.native(), cell.native(), gbool(expand))
}

// PackEnd is a wrapper around gtk_tree_view_column_pack_end().
func (v *TreeViewColumn) PackEnd(cell *CellRenderer, expand bool) {
	C.gtk_tree_view_column_pack_end(v.native(), cell.native(), gbool(expand))
}

// Clear is a wrapper around gtk_tree_view_column_clear().
func (v *TreeViewColumn) Clear() {
	C.gtk_tree_view_column_clear(v.native())
}

// ClearAttributes is a wrapper around gtk_tree_view_column_clear_attributes().
func (v *TreeViewColumn) ClearAttributes(cell *CellRenderer) {
	C.gtk_tree_view_column_clear_attributes(v.native(), cell.native())
}

// SetSpacing is a wrapper around gtk_tree_view_column_set_spacing().
func (v *TreeViewColumn) SetSpacing(spacing int) {
	C.gtk_tree_view_column_set_spacing(v.native(), C.gint(spacing))
}

// GetSpacing is a wrapper around gtk_tree_view_column_get_spacing().
func (v *TreeViewColumn) GetSpacing() int {
	return int(C.gtk_tree_view_column_get_spacing(v.native()))
}

// SetVisible is a wrapper around gtk_tree_view_column_set_visible().
func (v *TreeViewColumn) SetVisible(visible bool) {
	C.gtk_tree_view_column_set_visible(v.native(), gbool(visible))
}

// GetVisible is a wrapper around gtk_tree_view_column_get_visible().
func (v *TreeViewColumn) GetVisible() bool {
	return gobool(C.gtk_tree_view_column_get_visible(v.native()))
}

// SetResizable is a wrapper around gtk_tree_view_column_set_resizable().
func (v *TreeViewColumn) SetResizable(resizable bool) {
	C.gtk_tree_view_column_set_resizable(v.native(), gbool(resizable))
}

// GetResizable is a wrapper around gtk_tree_view_column_get_resizable().
func (v *TreeViewColumn) GetResizable() bool {
	return gobool(C.gtk_tree_view_column_get_resizable(v.native()))
}

// GetWidth is a wrapper around gtk_tree_view_column_get_width().
func (v *TreeViewColumn) GetWidth() int {
	return int(C.gtk_tree_view_column_get_width(v.native()))
}

// SetFixedWidth is a wrapper around gtk_tree_view_column_set_fixed_width().
func (v *TreeViewColumn) SetFixedWidth(w int) {
	C.gtk_tree_view_column_set_fixed_width(v.native(), C.gint(w))
}

// GetFixedWidth is a wrapper around gtk_tree_view_column_get_fixed_width().
func (v *TreeViewColumn) GetFixedWidth() int {
	return int(C.gtk_tree_view_column_get_fixed_width(v.native()))
}

// SetMaxWidth is a wrapper around gtk_tree_view_column_set_max_width().
func (v *TreeViewColumn) SetMaxWidth(w int) {
	C.gtk_tree_view_column_set_max_width(v.native(), C.gint(w))
}

// GetMaxWidth is a wrapper around gtk_tree_view_column_get_max_width().
func (v *TreeViewColumn) GetMaxWidth() int {
	return int(C.gtk_tree_view_column_get_max_width(v.native()))
}

// Clicked is a wrapper around gtk_tree_view_column_clicked().
func (v *TreeViewColumn) Clicked() {
	C.gtk_tree_view_column_clicked(v.native())
}

// SetTitle is a wrapper around gtk_tree_view_column_set_title().
func (v *TreeViewColumn) SetTitle(t string) {
	cstr := C.CString(t)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_tree_view_column_set_title(v.native(), (*C.gchar)(cstr))
}

// GetTitle is a wrapper around gtk_tree_view_column_get_title().
func (v *TreeViewColumn) GetTitle() string {
	return goString(C.gtk_tree_view_column_get_title(v.native()))
}

// SetClickable is a wrapper around gtk_tree_view_column_set_clickable().
func (v *TreeViewColumn) SetClickable(clickable bool) {
	C.gtk_tree_view_column_set_clickable(v.native(), gbool(clickable))
}

// GetClickable is a wrapper around gtk_tree_view_column_get_clickable().
func (v *TreeViewColumn) GetClickable() bool {
	return gobool(C.gtk_tree_view_column_get_clickable(v.native()))
}

// SetReorderable is a wrapper around gtk_tree_view_column_set_reorderable().
func (v *TreeViewColumn) SetReorderable(reorderable bool) {
	C.gtk_tree_view_column_set_reorderable(v.native(), gbool(reorderable))
}

// GetReorderable is a wrapper around gtk_tree_view_column_get_reorderable().
func (v *TreeViewColumn) GetReorderable() bool {
	return gobool(C.gtk_tree_view_column_get_reorderable(v.native()))
}

// SetSortIndicator is a wrapper around gtk_tree_view_column_set_sort_indicator().
func (v *TreeViewColumn) SetSortIndicator(reorderable bool) {
	C.gtk_tree_view_column_set_sort_indicator(v.native(), gbool(reorderable))
}

// GetSortIndicator is a wrapper around gtk_tree_view_column_get_sort_indicator().
func (v *TreeViewColumn) GetSortIndicator() bool {
	return gobool(C.gtk_tree_view_column_get_sort_indicator(v.native()))
}

// SetSortColumnID is a wrapper around gtk_tree_view_column_set_sort_column_id().
func (v *TreeViewColumn) SetSortColumnID(w int) {
	C.gtk_tree_view_column_set_sort_column_id(v.native(), C.gint(w))
}

// GetSortColumnID is a wrapper around gtk_tree_view_column_get_sort_column_id().
func (v *TreeViewColumn) GetSortColumnID() int {
	return int(C.gtk_tree_view_column_get_sort_column_id(v.native()))
}

// CellIsVisible is a wrapper around gtk_tree_view_column_cell_is_visible().
func (v *TreeViewColumn) CellIsVisible() bool {
	return gobool(C.gtk_tree_view_column_cell_is_visible(v.native()))
}

// FocusCell is a wrapper around gtk_tree_view_column_focus_cell().
func (v *TreeViewColumn) FocusCell(cell *CellRenderer) {
	C.gtk_tree_view_column_focus_cell(v.native(), cell.native())
}

// QueueResize is a wrapper around gtk_tree_view_column_queue_resize().
func (v *TreeViewColumn) QueueResize() {
	C.gtk_tree_view_column_queue_resize(v.native())
}

// GetXOffset is a wrapper around gtk_tree_view_column_get_x_offset().
func (v *TreeViewColumn) GetXOffset() int {
	return int(C.gtk_tree_view_column_get_x_offset(v.native()))
}

// GtkTreeViewColumn * 	gtk_tree_view_column_new_with_area ()
// void 	gtk_tree_view_column_set_attributes ()
// void 	gtk_tree_view_column_set_cell_data_func ()
// void 	gtk_tree_view_column_set_sizing ()
// GtkTreeViewColumnSizing 	gtk_tree_view_column_get_sizing ()
// void 	gtk_tree_view_column_set_widget ()
// GtkWidget * 	gtk_tree_view_column_get_widget ()
// GtkWidget * 	gtk_tree_view_column_get_button ()
// void 	gtk_tree_view_column_set_alignment ()
// gfloat 	gtk_tree_view_column_get_alignment ()
// void 	gtk_tree_view_column_set_sort_order ()
// GtkSortType 	gtk_tree_view_column_get_sort_order ()
// void 	gtk_tree_view_column_cell_set_cell_data ()
// void 	gtk_tree_view_column_cell_get_size ()
// gboolean 	gtk_tree_view_column_cell_get_position ()
// GtkWidget * 	gtk_tree_view_column_get_tree_view ()

/*
 * GtkTreeView
 */

// TreeView is a representation of GTK's GtkTreeView.
type TreeView struct {
	Container
}

// native returns a pointer to the underlying GtkTreeView.
func (v *TreeView) native() *C.GtkTreeView {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toGtkTreeView(ptr)
}

func marshalTreeView(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapTreeView(obj), nil
}

func wrapTreeView(obj *glib.Object) *TreeView {
	return &TreeView{Container{Widget{glib.InitiallyUnowned{obj}}}}
}

func setupTreeView(c unsafe.Pointer) (*TreeView, error) {
	if c == nil {
		return nil, nilPtrErr
	}

	return wrapTreeView(glib.Take(c)), nil
}

// TreeViewNew is a wrapper around gtk_tree_view_new().
func TreeViewNew() (*TreeView, error) {
	return setupTreeView(unsafe.Pointer(C.gtk_tree_view_new()))
}

// TreeViewNewWithModel is a wrapper around gtk_tree_view_new_with_model().
func TreeViewNewWithModel(model ITreeModel) (*TreeView, error) {
	return setupTreeView(unsafe.Pointer(C.gtk_tree_view_new_with_model(model.toTreeModel())))
}

// GetModel is a wrapper around gtk_tree_view_get_model().
func (v *TreeView) GetModel() (*TreeModel, error) {
	c := C.gtk_tree_view_get_model(v.native())
	if c == nil {
		return nil, nil
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapTreeModel(*glib.InterfaceFromObjectNew(obj)), nil
}

// SetModel is a wrapper around gtk_tree_view_set_model().
func (v *TreeView) SetModel(model ITreeModel) {
	if model == nil {
		C.gtk_tree_view_set_model(v.native(), nil)
	} else {
		C.gtk_tree_view_set_model(v.native(), model.toTreeModel())
	}
}

// GetSelection is a wrapper around gtk_tree_view_get_selection().
func (v *TreeView) GetSelection() (*TreeSelection, error) {
	c := C.gtk_tree_view_get_selection(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapTreeSelection(obj), nil
}

// AppendColumn is a wrapper around gtk_tree_view_append_column().
func (v *TreeView) AppendColumn(column *TreeViewColumn) int {
	c := C.gtk_tree_view_append_column(v.native(), column.native())
	return int(c)
}

// GetPathAtPos is a wrapper around gtk_tree_view_get_path_at_pos().
func (v *TreeView) GetPathAtPos(x, y int, path *TreePath, column *TreeViewColumn, cellX, cellY *int) error {
	var ctp **C.GtkTreePath
	if path != nil {
		tp := path.native()
		ctp = &tp
	} else {
		ctp = nil
	}

	var pctvcol **C.GtkTreeViewColumn
	if column != nil {
		ctvcol := column.native()
		pctvcol = &ctvcol
	} else {
		pctvcol = nil
	}

	c := C.gtk_tree_view_get_path_at_pos(
		v.native(),
		(C.gint)(x),
		(C.gint)(y),
		ctp,
		pctvcol,
		(*C.gint)(unsafe.Pointer(cellX)),
		(*C.gint)(unsafe.Pointer(cellY)))
	if !gobool(c) {
		return errors.New("Unable to set path at position")
	}
	return nil
}

// GetLevelIndentation is a wrapper around gtk_tree_view_get_level_indentation().
func (v *TreeView) GetLevelIndentation() int {
	return int(C.gtk_tree_view_get_level_indentation(v.native()))
}

// GetShowExpanders is a wrapper around gtk_tree_view_get_show_expanders().
func (v *TreeView) GetShowExpanders() bool {
	return gobool(C.gtk_tree_view_get_show_expanders(v.native()))
}

// SetLevelIndentation is a wrapper around gtk_tree_view_set_level_indentation().
func (v *TreeView) SetLevelIndentation(indent int) {
	C.gtk_tree_view_set_level_indentation(v.native(), C.gint(indent))
}

// SetShowExpanders is a wrapper around gtk_tree_view_set_show_expanders().
func (v *TreeView) SetShowExpanders(show bool) {
	C.gtk_tree_view_set_show_expanders(v.native(), gbool(show))
}

// GetHeadersVisible is a wrapper around gtk_tree_view_get_headers_visible().
func (v *TreeView) GetHeadersVisible() bool {
	return gobool(C.gtk_tree_view_get_headers_visible(v.native()))
}

// SetHeadersVisible is a wrapper around gtk_tree_view_set_headers_visible().
func (v *TreeView) SetHeadersVisible(show bool) {
	C.gtk_tree_view_set_headers_visible(v.native(), gbool(show))
}

// ColumnsAutosize is a wrapper around gtk_tree_view_columns_autosize().
func (v *TreeView) ColumnsAutosize() {
	C.gtk_tree_view_columns_autosize(v.native())
}

// GetHeadersClickable is a wrapper around gtk_tree_view_get_headers_clickable().
func (v *TreeView) GetHeadersClickable() bool {
	return gobool(C.gtk_tree_view_get_headers_clickable(v.native()))
}

// SetHeadersClickable is a wrapper around gtk_tree_view_set_headers_clickable().
func (v *TreeView) SetHeadersClickable(show bool) {
	C.gtk_tree_view_set_headers_clickable(v.native(), gbool(show))
}

// GetActivateOnSingleClick is a wrapper around gtk_tree_view_get_activate_on_single_click().
func (v *TreeView) GetActivateOnSingleClick() bool {
	return gobool(C.gtk_tree_view_get_activate_on_single_click(v.native()))
}

// SetActivateOnSingleClick is a wrapper around gtk_tree_view_set_activate_on_single_click().
func (v *TreeView) SetActivateOnSingleClick(show bool) {
	C.gtk_tree_view_set_activate_on_single_click(v.native(), gbool(show))
}

// RemoveColumn is a wrapper around gtk_tree_view_remove_column().
func (v *TreeView) RemoveColumn(column *TreeViewColumn) int {
	return int(C.gtk_tree_view_remove_column(v.native(), column.native()))
}

// InsertColumn is a wrapper around gtk_tree_view_insert_column().
func (v *TreeView) InsertColumn(column *TreeViewColumn, pos int) int {
	return int(C.gtk_tree_view_insert_column(v.native(), column.native(), C.gint(pos)))
}

// GetNColumns is a wrapper around gtk_tree_view_get_n_columns().
func (v *TreeView) GetNColumns() uint {
	return uint(C.gtk_tree_view_get_n_columns(v.native()))
}

// GetColumn is a wrapper around gtk_tree_view_get_column().
func (v *TreeView) GetColumn(n int) *TreeViewColumn {
	c := C.gtk_tree_view_get_column(v.native(), C.gint(n))
	if c == nil {
		return nil
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapTreeViewColumn(obj)
}

// MoveColumnAfter is a wrapper around gtk_tree_view_move_column_after().
func (v *TreeView) MoveColumnAfter(column *TreeViewColumn, baseColumn *TreeViewColumn) {
	C.gtk_tree_view_move_column_after(v.native(), column.native(), baseColumn.native())
}

// SetExpanderColumn is a wrapper around gtk_tree_view_set_expander_column().
func (v *TreeView) SetExpanderColumn(column *TreeViewColumn) {
	C.gtk_tree_view_set_expander_column(v.native(), column.native())
}

// GetExpanderColumn is a wrapper around gtk_tree_view_get_expander_column().
func (v *TreeView) GetExpanderColumn() *TreeViewColumn {
	c := C.gtk_tree_view_get_expander_column(v.native())
	if c == nil {
		return nil
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapTreeViewColumn(obj)
}

// ScrollToPoint is a wrapper around gtk_tree_view_scroll_to_point().
func (v *TreeView) ScrollToPoint(treeX, treeY int) {
	C.gtk_tree_view_scroll_to_point(v.native(), C.gint(treeX), C.gint(treeY))
}

// SetCursor is a wrapper around gtk_tree_view_set_cursor().
func (v *TreeView) SetCursor(path *TreePath, focusColumn *TreeViewColumn, startEditing bool) {
	C.gtk_tree_view_set_cursor(v.native(), path.native(), focusColumn.native(), gbool(startEditing))
}

// SetCursorOnCell is a wrapper around gtk_tree_view_set_cursor_on_cell().
func (v *TreeView) SetCursorOnCell(path *TreePath, focusColumn *TreeViewColumn, focusCell *CellRenderer, startEditing bool) {
	C.gtk_tree_view_set_cursor_on_cell(v.native(), path.native(), focusColumn.native(), focusCell.native(), gbool(startEditing))
}

// GetCursor is a wrapper around gtk_tree_view_get_cursor().
func (v *TreeView) GetCursor() (p *TreePath, c *TreeViewColumn) {
	var path *C.GtkTreePath
	var col *C.GtkTreeViewColumn

	C.gtk_tree_view_get_cursor(v.native(), &path, &col)

	if path != nil {
		p = wrapTreePath(path)
		runtime.SetFinalizer(p, (*TreePath).free)
	}

	if col != nil {
		c = wrapTreeViewColumn(glib.Take(unsafe.Pointer(col)))
	}

	return
}

// RowActivated is a wrapper around gtk_tree_view_row_activated().
func (v *TreeView) RowActivated(path *TreePath, column *TreeViewColumn) {
	C.gtk_tree_view_row_activated(v.native(), path.native(), column.native())
}

// ExpandAll is a wrapper around gtk_tree_view_expand_all().
func (v *TreeView) ExpandAll() {
	C.gtk_tree_view_expand_all(v.native())
}

// CollapseAll is a wrapper around gtk_tree_view_collapse_all().
func (v *TreeView) CollapseAll() {
	C.gtk_tree_view_collapse_all(v.native())
}

// ExpandToPath is a wrapper around gtk_tree_view_expand_to_path().
func (v *TreeView) ExpandToPath(path *TreePath) {
	C.gtk_tree_view_expand_to_path(v.native(), path.native())
}

// ExpandRow is a wrapper around gtk_tree_view_expand_row().
func (v *TreeView) ExpandRow(path *TreePath, openAll bool) bool {
	return gobool(C.gtk_tree_view_expand_row(v.native(), path.native(), gbool(openAll)))
}

// CollapseRow is a wrapper around gtk_tree_view_collapse_row().
func (v *TreeView) CollapseRow(path *TreePath) bool {
	return gobool(C.gtk_tree_view_collapse_row(v.native(), path.native()))
}

// RowExpanded is a wrapper around gtk_tree_view_row_expanded().
func (v *TreeView) RowExpanded(path *TreePath) bool {
	return gobool(C.gtk_tree_view_row_expanded(v.native(), path.native()))
}

// SetReorderable is a wrapper around gtk_tree_view_set_reorderable().
func (v *TreeView) SetReorderable(b bool) {
	C.gtk_tree_view_set_reorderable(v.native(), gbool(b))
}

// GetReorderable is a wrapper around gtk_tree_view_get_reorderable().
func (v *TreeView) GetReorderable() bool {
	return gobool(C.gtk_tree_view_get_reorderable(v.native()))
}

// GetBinWindow is a wrapper around gtk_tree_view_get_bin_window().
func (v *TreeView) GetBinWindow() *gdk.Window {
	c := C.gtk_tree_view_get_bin_window(v.native())
	if c == nil {
		return nil
	}

	w := &gdk.Window{glib.Take(unsafe.Pointer(c))}
	return w
}

// SetEnableSearch is a wrapper around gtk_tree_view_set_enable_search().
func (v *TreeView) SetEnableSearch(b bool) {
	C.gtk_tree_view_set_enable_search(v.native(), gbool(b))
}

// GetEnableSearch is a wrapper around gtk_tree_view_get_enable_search().
func (v *TreeView) GetEnableSearch() bool {
	return gobool(C.gtk_tree_view_get_enable_search(v.native()))
}

// SetSearchColumn is a wrapper around gtk_tree_view_set_search_column().
func (v *TreeView) SetSearchColumn(c int) {
	C.gtk_tree_view_set_search_column(v.native(), C.gint(c))
}

// GetSearchColumn is a wrapper around gtk_tree_view_get_search_column().
func (v *TreeView) GetSearchColumn() int {
	return int(C.gtk_tree_view_get_search_column(v.native()))
}

// GetSearchEntry is a wrapper around gtk_tree_view_get_search_entry().
func (v *TreeView) GetSearchEntry() *Entry {
	c := C.gtk_tree_view_get_search_entry(v.native())
	if c == nil {
		return nil
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapEntry(obj)
}

// SetSearchEntry is a wrapper around gtk_tree_view_set_search_entry().
func (v *TreeView) SetSearchEntry(e *Entry) {
	C.gtk_tree_view_set_search_entry(v.native(), e.native())
}

// SetSearchEqualSubstringMatch sets TreeView to search by substring match.
func (v *TreeView) SetSearchEqualSubstringMatch() {
	C.gtk_tree_view_set_search_equal_func(
		v.native(),
		(C.GtkTreeViewSearchEqualFunc)(unsafe.Pointer(C.substring_match_equal_func)),
		nil,
		nil)
}

// SetFixedHeightMode is a wrapper around gtk_tree_view_set_fixed_height_mode().
func (v *TreeView) SetFixedHeightMode(b bool) {
	C.gtk_tree_view_set_fixed_height_mode(v.native(), gbool(b))
}

// GetFixedHeightMode is a wrapper around gtk_tree_view_get_fixed_height_mode().
func (v *TreeView) GetFixedHeightMode() bool {
	return gobool(C.gtk_tree_view_get_fixed_height_mode(v.native()))
}

// SetHoverSelection is a wrapper around gtk_tree_view_set_hover_selection().
func (v *TreeView) SetHoverSelection(b bool) {
	C.gtk_tree_view_set_hover_selection(v.native(), gbool(b))
}

// GetHoverSelection is a wrapper around gtk_tree_view_get_hover_selection().
func (v *TreeView) GetHoverSelection() bool {
	return gobool(C.gtk_tree_view_get_hover_selection(v.native()))
}

// SetHoverExpand is a wrapper around gtk_tree_view_set_hover_expand().
func (v *TreeView) SetHoverExpand(b bool) {
	C.gtk_tree_view_set_hover_expand(v.native(), gbool(b))
}

// GetHoverExpand is a wrapper around gtk_tree_view_get_hover_expand().
func (v *TreeView) GetHoverExpand() bool {
	return gobool(C.gtk_tree_view_get_hover_expand(v.native()))
}

// SetRubberBanding is a wrapper around gtk_tree_view_set_rubber_banding().
func (v *TreeView) SetRubberBanding(b bool) {
	C.gtk_tree_view_set_rubber_banding(v.native(), gbool(b))
}

// GetRubberBanding is a wrapper around gtk_tree_view_get_rubber_banding().
func (v *TreeView) GetRubberBanding() bool {
	return gobool(C.gtk_tree_view_get_rubber_banding(v.native()))
}

// IsRubberBandingActive is a wrapper around gtk_tree_view_is_rubber_banding_active().
func (v *TreeView) IsRubberBandingActive() bool {
	return gobool(C.gtk_tree_view_is_rubber_banding_active(v.native()))
}

// SetEnableTreeLines is a wrapper around gtk_tree_view_set_enable_tree_lines().
func (v *TreeView) SetEnableTreeLines(b bool) {
	C.gtk_tree_view_set_enable_tree_lines(v.native(), gbool(b))
}

// GetEnableTreeLines is a wrapper around gtk_tree_view_get_enable_tree_lines().
func (v *TreeView) GetEnableTreeLines() bool {
	return gobool(C.gtk_tree_view_get_enable_tree_lines(v.native()))
}

// GetTooltipColumn is a wrapper around gtk_tree_view_get_tooltip_column().
func (v *TreeView) GetTooltipColumn() int {
	return int(C.gtk_tree_view_get_tooltip_column(v.native()))
}

// SetTooltipColumn is a wrapper around gtk_tree_view_set_tooltip_column().
func (v *TreeView) SetTooltipColumn(c int) {
	C.gtk_tree_view_set_tooltip_column(v.native(), C.gint(c))
}

// void 	gtk_tree_view_set_tooltip_row ()
// void 	gtk_tree_view_set_tooltip_cell ()
// gboolean 	gtk_tree_view_get_tooltip_context ()
// void 	gtk_tree_view_set_grid_lines ()
// GtkTreeViewGridLines 	gtk_tree_view_get_grid_lines ()
// void 	(*GtkTreeDestroyCountFunc) ()
// void 	gtk_tree_view_set_destroy_count_func ()
// gboolean 	(*GtkTreeViewRowSeparatorFunc) ()
// GtkTreeViewRowSeparatorFunc 	gtk_tree_view_get_row_separator_func ()
// void 	gtk_tree_view_set_row_separator_func ()
// void 	(*GtkTreeViewSearchPositionFunc) ()
// GtkTreeViewSearchPositionFunc 	gtk_tree_view_get_search_position_func ()
// void 	gtk_tree_view_set_search_position_func ()
// void 	gtk_tree_view_set_search_equal_func ()
// GtkTreeViewSearchEqualFunc 	gtk_tree_view_get_search_equal_func ()
// void 	gtk_tree_view_map_expanded_rows ()
// GList * 	gtk_tree_view_get_columns ()
// gint 	gtk_tree_view_insert_column_with_attributes ()
// gint 	gtk_tree_view_insert_column_with_data_func ()
// void 	gtk_tree_view_set_column_drag_function ()
// void 	gtk_tree_view_scroll_to_cell ()
// gboolean 	gtk_tree_view_is_blank_at_pos ()
// void 	gtk_tree_view_get_cell_area ()
// void 	gtk_tree_view_get_background_area ()
// void 	gtk_tree_view_get_visible_rect ()
// gboolean 	gtk_tree_view_get_visible_range ()
// void 	gtk_tree_view_convert_bin_window_to_tree_coords ()
// void 	gtk_tree_view_convert_bin_window_to_widget_coords ()
// void 	gtk_tree_view_convert_tree_to_bin_window_coords ()
// void 	gtk_tree_view_convert_tree_to_widget_coords ()
// void 	gtk_tree_view_convert_widget_to_bin_window_coords ()
// void 	gtk_tree_view_convert_widget_to_tree_coords ()
// void 	gtk_tree_view_enable_model_drag_dest ()
// void 	gtk_tree_view_enable_model_drag_source ()
// void 	gtk_tree_view_unset_rows_drag_source ()
// void 	gtk_tree_view_unset_rows_drag_dest ()
// void 	gtk_tree_view_set_drag_dest_row ()
// void 	gtk_tree_view_get_drag_dest_row ()
// gboolean 	gtk_tree_view_get_dest_row_at_pos ()
// cairo_surface_t * 	gtk_tree_view_create_row_drag_icon ()
