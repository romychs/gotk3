package gtk

// #include <gtk/gtk.h>
// #include "gtk.go.h"
import "C"
import (
	"errors"
	"runtime"
	"sync"
	"unsafe"

	"github.com/d2r2/gotk3/cairo"
	"github.com/d2r2/gotk3/glib"
	"github.com/d2r2/gotk3/pango"
)

func init() {
	tm := []glib.TypeMarshaler{
		// Enums
		{glib.Type(C.gtk_page_orientation_get_type()), marshalPageOrientation},
		{glib.Type(C.gtk_print_error_get_type()), marshalPrintError},
		{glib.Type(C.gtk_print_operation_action_get_type()), marshalPrintOperationAction},
		{glib.Type(C.gtk_print_operation_result_get_type()), marshalPrintOperationResult},
		{glib.Type(C.gtk_print_status_get_type()), marshalPrintStatus},
		{glib.Type(C.gtk_unit_get_type()), marshalUnit},

		// Objects/Interfaces
		{glib.Type(C.gtk_number_up_layout_get_type()), marshalNumberUpLayout},
		{glib.Type(C.gtk_page_orientation_get_type()), marshalPageOrientation},
		{glib.Type(C.gtk_page_set_get_type()), marshalPageSet},
		{glib.Type(C.gtk_page_setup_get_type()), marshalPageSetup},
		{glib.Type(C.gtk_print_context_get_type()), marshalPrintContext},
		{glib.Type(C.gtk_print_duplex_get_type()), marshalPrintDuplex},
		{glib.Type(C.gtk_print_operation_get_type()), marshalPrintOperation},
		{glib.Type(C.gtk_print_operation_preview_get_type()), marshalPrintOperationPreview},
		{glib.Type(C.gtk_print_pages_get_type()), marshalPrintPages},
		{glib.Type(C.gtk_print_quality_get_type()), marshalPrintQuality},
		{glib.Type(C.gtk_print_settings_get_type()), marshalPrintSettings},

		// Boxed
		{glib.Type(C.gtk_paper_size_get_type()), marshalPaperSize},
	}

	glib.RegisterGValueMarshalers(tm)

	WrapMap["GtkPageSetup"] = wrapPageSetup
	WrapMap["GtkPrintContext"] = wrapPrintContext
	WrapMap["GtkPrintOperation"] = wrapPrintOperation
	WrapMap["GtkPrintOperationPreview"] = wrapPrintOperationPreview
	WrapMap["GtkPrintSettings"] = wrapPrintSettings
}

/*
 * Constants
 */

// NumberUpLayout is a representation of GTK's GtkNumberUpLayout.
type NumberUpLayout int

const (
	NUMBER_UP_LAYOUT_LEFT_TO_RIGHT_TOP_TO_BOTTOM NumberUpLayout = C.GTK_NUMBER_UP_LAYOUT_LEFT_TO_RIGHT_TOP_TO_BOTTOM
	NUMBER_UP_LAYOUT_LEFT_TO_RIGHT_BOTTOM_TO_TOP NumberUpLayout = C.GTK_NUMBER_UP_LAYOUT_LEFT_TO_RIGHT_BOTTOM_TO_TOP
	NUMBER_UP_LAYOUT_RIGHT_TO_LEFT_TOP_TO_BOTTOM NumberUpLayout = C.GTK_NUMBER_UP_LAYOUT_RIGHT_TO_LEFT_TOP_TO_BOTTOM
	NUMBER_UP_LAYOUT_RIGHT_TO_LEFT_BOTTOM_TO_TOP NumberUpLayout = C.GTK_NUMBER_UP_LAYOUT_RIGHT_TO_LEFT_BOTTOM_TO_TOP
	NUMBER_UP_LAYOUT_TOP_TO_BOTTOM_LEFT_TO_RIGHT NumberUpLayout = C.GTK_NUMBER_UP_LAYOUT_TOP_TO_BOTTOM_LEFT_TO_RIGHT
	NUMBER_UP_LAYOUT_TOP_TO_BOTTOM_RIGHT_TO_LEFT NumberUpLayout = C.GTK_NUMBER_UP_LAYOUT_TOP_TO_BOTTOM_RIGHT_TO_LEFT
	NUMBER_UP_LAYOUT_BOTTOM_TO_TOP_LEFT_TO_RIGHT NumberUpLayout = C.GTK_NUMBER_UP_LAYOUT_BOTTOM_TO_TOP_LEFT_TO_RIGHT
	NUMBER_UP_LAYOUT_BOTTOM_TO_TOP_RIGHT_TO_LEFT NumberUpLayout = C.GTK_NUMBER_UP_LAYOUT_BOTTOM_TO_TOP_RIGHT_TO_LEFT
)

func marshalNumberUpLayout(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return NumberUpLayout(c), nil
}

// PageOrientation is a representation of GTK's GtkPageOrientation.
type PageOrientation int

const (
	PAGE_ORIENTATION_PORTRAIT          PageOrientation = C.GTK_PAGE_ORIENTATION_PORTRAIT
	PAGE_ORIENTATION_LANDSCAPE         PageOrientation = C.GTK_PAGE_ORIENTATION_LANDSCAPE
	PAGE_ORIENTATION_REVERSE_PORTRAIT  PageOrientation = C.GTK_PAGE_ORIENTATION_REVERSE_PORTRAIT
	PAGE_ORIENTATION_REVERSE_LANDSCAPE PageOrientation = C.GTK_PAGE_ORIENTATION_REVERSE_LANDSCAPE
)

func marshalPageOrientation(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return PageOrientation(c), nil
}

// PrintDuplex is a representation of GTK's GtkPrintDuplex.
type PrintDuplex int

const (
	PRINT_DUPLEX_SIMPLEX    PrintDuplex = C.GTK_PRINT_DUPLEX_SIMPLEX
	PRINT_DUPLEX_HORIZONTAL PrintDuplex = C.GTK_PRINT_DUPLEX_HORIZONTAL
	PRINT_DUPLEX_VERTICAL   PrintDuplex = C.GTK_PRINT_DUPLEX_VERTICAL
)

func marshalPrintDuplex(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return PrintDuplex(c), nil
}

// PrintPages is a representation of GTK's GtkPrintPages.
type PrintPages int

const (
	PRINT_PAGES_ALL       PrintPages = C.GTK_PRINT_PAGES_ALL
	PRINT_PAGES_CURRENT   PrintPages = C.GTK_PRINT_PAGES_CURRENT
	PRINT_PAGES_RANGES    PrintPages = C.GTK_PRINT_PAGES_RANGES
	PRINT_PAGES_SELECTION PrintPages = C.GTK_PRINT_PAGES_SELECTION
)

func marshalPrintPages(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return PrintPages(c), nil
}

// PageSet is a representation of GTK's GtkPageSet.
type PageSet int

const (
	PAGE_SET_ALL  PageSet = C.GTK_PAGE_SET_ALL
	PAGE_SET_EVEN PageSet = C.GTK_PAGE_SET_EVEN
	PAGE_SET_ODD  PageSet = C.GTK_PAGE_SET_ODD
)

func marshalPageSet(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return PageSet(c), nil
}

// PrintOperationAction is a representation of GTK's GtkPrintError.
type PrintError int

const (
	PRINT_ERROR_GENERAL        PrintError = C.GTK_PRINT_ERROR_GENERAL
	PRINT_ERROR_INTERNAL_ERROR PrintError = C.GTK_PRINT_ERROR_INTERNAL_ERROR
	PRINT_ERROR_NOMEM          PrintError = C.GTK_PRINT_ERROR_NOMEM
	PRINT_ERROR_INVALID_FILE   PrintError = C.GTK_PRINT_ERROR_INVALID_FILE
)

func marshalPrintError(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return PrintError(c), nil
}

// PrintOperationAction is a representation of GTK's GtkPrintOperationAction.
type PrintOperationAction int

const (
	PRINT_OPERATION_ACTION_PRINT_DIALOG PrintOperationAction = C.GTK_PRINT_OPERATION_ACTION_PRINT_DIALOG
	PRINT_OPERATION_ACTION_PRINT        PrintOperationAction = C.GTK_PRINT_OPERATION_ACTION_PRINT
	PRINT_OPERATION_ACTION_PREVIEW      PrintOperationAction = C.GTK_PRINT_OPERATION_ACTION_PREVIEW
	PRINT_OPERATION_ACTION_EXPORT       PrintOperationAction = C.GTK_PRINT_OPERATION_ACTION_EXPORT
)

func marshalPrintOperationAction(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return PrintOperationAction(c), nil
}

// PrintOperationResult is a representation of GTK's GtkPrintOperationResult.
type PrintOperationResult int

const (
	PRINT_OPERATION_RESULT_ERROR       PrintOperationResult = C.GTK_PRINT_OPERATION_RESULT_ERROR
	PRINT_OPERATION_RESULT_APPLY       PrintOperationResult = C.GTK_PRINT_OPERATION_RESULT_APPLY
	PRINT_OPERATION_RESULT_CANCEL      PrintOperationResult = C.GTK_PRINT_OPERATION_RESULT_CANCEL
	PRINT_OPERATION_RESULT_IN_PROGRESS PrintOperationResult = C.GTK_PRINT_OPERATION_RESULT_IN_PROGRESS
)

func marshalPrintOperationResult(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return PrintOperationResult(c), nil
}

// PrintStatus is a representation of GTK's GtkPrintStatus.
type PrintStatus int

const (
	PRINT_STATUS_INITIAL          PrintStatus = C.GTK_PRINT_STATUS_INITIAL
	PRINT_STATUS_PREPARING        PrintStatus = C.GTK_PRINT_STATUS_PREPARING
	PRINT_STATUS_GENERATING_DATA  PrintStatus = C.GTK_PRINT_STATUS_GENERATING_DATA
	PRINT_STATUS_SENDING_DATA     PrintStatus = C.GTK_PRINT_STATUS_SENDING_DATA
	PRINT_STATUS_PENDING          PrintStatus = C.GTK_PRINT_STATUS_PENDING
	PRINT_STATUS_PENDING_ISSUE    PrintStatus = C.GTK_PRINT_STATUS_PENDING_ISSUE
	PRINT_STATUS_PRINTING         PrintStatus = C.GTK_PRINT_STATUS_PRINTING
	PRINT_STATUS_FINISHED         PrintStatus = C.GTK_PRINT_STATUS_FINISHED
	PRINT_STATUS_FINISHED_ABORTED PrintStatus = C.GTK_PRINT_STATUS_FINISHED_ABORTED
)

func marshalPrintStatus(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return PrintStatus(c), nil
}

// PrintQuality is a representation of GTK's GtkPrintQuality.
type PrintQuality int

const (
	PRINT_QUALITY_LOW    PrintQuality = C.GTK_PRINT_QUALITY_LOW
	PRINT_QUALITY_NORMAL PrintQuality = C.GTK_PRINT_QUALITY_NORMAL
	PRINT_QUALITY_HIGH   PrintQuality = C.GTK_PRINT_QUALITY_HIGH
	PRINT_QUALITY_DRAFT  PrintQuality = C.GTK_PRINT_QUALITY_DRAFT
)

func marshalPrintQuality(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return PrintQuality(c), nil
}

// Unit is a representation of GTK's GtkUnit.
type Unit int

const (
	GTK_UNIT_NONE   Unit = C.GTK_UNIT_NONE
	GTK_UNIT_POINTS Unit = C.GTK_UNIT_POINTS
	GTK_UNIT_INCH   Unit = C.GTK_UNIT_INCH
	GTK_UNIT_MM     Unit = C.GTK_UNIT_MM
)

func marshalUnit(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return Unit(c), nil
}

/*
 * GtkPageSetup
 */
type PageSetup struct {
	*glib.Object
}

func (v *PageSetup) native() *C.GtkPageSetup {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toGtkPageSetup(ptr)
}

func marshalPageSetup(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapPageSetup(obj), nil
}

func wrapPageSetup(obj *glib.Object) *PageSetup {
	return &PageSetup{obj}
}

// PageSetupNew is a wrapper around gtk_page_setup_new().
func PageSetupNew() (*PageSetup, error) {
	c := C.gtk_page_setup_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapPageSetup(obj), nil
}

// Copy is a wrapper around gtk_page_setup_copy().
func (v *PageSetup) Copy() (*PageSetup, error) {
	c := C.gtk_page_setup_copy(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapPageSetup(obj), nil
}

// GetOrientation is a wrapper around gtk_page_setup_get_orientation().
func (v *PageSetup) GetOrientation() PageOrientation {
	c := C.gtk_page_setup_get_orientation(v.native())
	return PageOrientation(c)
}

// SetOrientation is a wrapper around gtk_page_setup_set_orientation().
func (v *PageSetup) SetOrientation(orientation PageOrientation) {
	C.gtk_page_setup_set_orientation(v.native(), C.GtkPageOrientation(orientation))
}

// GetPaperSize is a wrapper around gtk_page_setup_get_paper_size().
func (v *PageSetup) GetPaperSize() *PaperSize {
	c := C.gtk_page_setup_get_paper_size(v.native())
	p := wrapPaperSize(c)
	runtime.SetFinalizer(p, (*PaperSize).free)
	return p
}

// SetPaperSize is a wrapper around gtk_page_setup_set_paper_size().
func (v *PageSetup) SetPaperSize(size *PaperSize) {
	C.gtk_page_setup_set_paper_size(v.native(), size.native())
}

// GetTopMargin is a wrapper around gtk_page_setup_get_top_margin().
func (v *PageSetup) GetTopMargin(unit Unit) float64 {
	c := C.gtk_page_setup_get_top_margin(v.native(), C.GtkUnit(unit))
	return float64(c)
}

// SetTopMargin is a wrapper around gtk_page_setup_set_top_margin().
func (v *PageSetup) SetTopMargin(margin float64, unit Unit) {
	C.gtk_page_setup_set_top_margin(v.native(), C.gdouble(margin), C.GtkUnit(unit))
}

// GetBottomMargin is a wrapper around gtk_page_setup_get_bottom_margin().
func (v *PageSetup) GetBottomMargin(unit Unit) float64 {
	c := C.gtk_page_setup_get_bottom_margin(v.native(), C.GtkUnit(unit))
	return float64(c)
}

// SetBottomMargin is a wrapper around gtk_page_setup_set_bottom_margin().
func (v *PageSetup) SetBottomMargin(margin float64, unit Unit) {
	C.gtk_page_setup_set_bottom_margin(v.native(), C.gdouble(margin), C.GtkUnit(unit))
}

// GetLeftMargin is a wrapper around gtk_page_setup_get_left_margin().
func (v *PageSetup) GetLeftMargin(unit Unit) float64 {
	c := C.gtk_page_setup_get_left_margin(v.native(), C.GtkUnit(unit))
	return float64(c)
}

// SetLeftMargin is a wrapper around gtk_page_setup_set_left_margin().
func (v *PageSetup) SetLeftMargin(margin float64, unit Unit) {
	C.gtk_page_setup_set_left_margin(v.native(), C.gdouble(margin), C.GtkUnit(unit))
}

// GetRightMargin is a wrapper around gtk_page_setup_get_right_margin().
func (v *PageSetup) GetRightMargin(unit Unit) float64 {
	c := C.gtk_page_setup_get_right_margin(v.native(), C.GtkUnit(unit))
	return float64(c)
}

// SetRightMargin is a wrapper around gtk_page_setup_set_right_margin().
func (v *PageSetup) SetRightMargin(margin float64, unit Unit) {
	C.gtk_page_setup_set_right_margin(v.native(), C.gdouble(margin), C.GtkUnit(unit))
}

// SetPaperSizeAndDefaultMargins is a wrapper around gtk_page_setup_set_paper_size_and_default_margins().
func (v *PageSetup) SetPaperSizeAndDefaultMargins(size *PaperSize) {
	C.gtk_page_setup_set_paper_size_and_default_margins(v.native(), size.native())
}

// GetPaperWidth is a wrapper around gtk_page_setup_get_paper_width().
func (v *PageSetup) GetPaperWidth(unit Unit) float64 {
	c := C.gtk_page_setup_get_paper_width(v.native(), C.GtkUnit(unit))
	return float64(c)
}

// GetPaperHeight is a wrapper around gtk_page_setup_get_paper_height().
func (v *PageSetup) GetPaperHeight(unit Unit) float64 {
	c := C.gtk_page_setup_get_paper_height(v.native(), C.GtkUnit(unit))
	return float64(c)
}

// GetPageWidth is a wrapper around gtk_page_setup_get_page_width().
func (v *PageSetup) GetPageWidth(unit Unit) float64 {
	c := C.gtk_page_setup_get_page_width(v.native(), C.GtkUnit(unit))
	return float64(c)
}

// GetPageHeight is a wrapper around gtk_page_setup_get_page_height().
func (v *PageSetup) GetPageHeight(unit Unit) float64 {
	c := C.gtk_page_setup_get_page_height(v.native(), C.GtkUnit(unit))
	return float64(c)
}

// PageSetupNewFromFile is a wrapper around gtk_page_setup_new_from_file().
func PageSetupNewFromFile(fileName string) (*PageSetup, error) {
	cstr := C.CString(fileName)
	defer C.free(unsafe.Pointer(cstr))
	var err *C.GError = nil
	c := C.gtk_page_setup_new_from_file((*C.gchar)(cstr), &err)
	if c == nil {
		defer C.g_error_free(err)
		return nil, errors.New(goString(err.message))
	}
	obj := glib.Take(unsafe.Pointer(c))
	return &PageSetup{obj}, nil

}

// PageSetupNewFromKeyFile() is a wrapper around gtk_page_setup_new_from_key_file().

// PageSetupLoadFile is a wrapper around gtk_page_setup_load_file().
func (v *PageSetup) PageSetupLoadFile(name string) error {
	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))
	var err *C.GError = nil
	res := C.gtk_page_setup_load_file(v.native(), cstr, &err)
	if !gobool(res) {
		defer C.g_error_free(err)
		return errors.New(goString(err.message))
	}
	return nil
}

// PageSetupLoadKeyFile() is a wrapper around gtk_page_setup_load_key_file().

// PageSetupToFile is a wrapper around gtk_page_setup_to_file().
func (v *PageSetup) PageSetupToFile(name string) error {
	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))
	var err *C.GError = nil
	res := C.gtk_page_setup_to_file(v.native(), cstr, &err)
	if !gobool(res) {
		defer C.g_error_free(err)
		return errors.New(goString(err.message))
	}
	return nil
}

// PageSetupToKeyFile() is a wrapper around gtk_page_setup_to_key_file().

/*
 * GtkPaperSize
 */

// PaperSize is a representation of GTK's GtkPaperSize
type PaperSize struct {
	gtkPaperSize *C.GtkPaperSize
}

// native returns a pointer to the underlying GtkPaperSize.
func (v *PaperSize) native() *C.GtkPaperSize {
	if v == nil {
		return nil
	}
	return v.gtkPaperSize
}

func marshalPaperSize(p uintptr) (interface{}, error) {
	c := C.g_value_get_boxed(C.toGValue(unsafe.Pointer(p)))
	c2 := (*C.GtkPaperSize)(unsafe.Pointer(c))
	return wrapPaperSize(c2), nil
}

func wrapPaperSize(obj *C.GtkPaperSize) *PaperSize {
	return &PaperSize{obj}
}

const (
	UNIT_PIXEL           int    = C.GTK_UNIT_PIXEL
	PAPER_NAME_A3        string = C.GTK_PAPER_NAME_A3
	PAPER_NAME_A4        string = C.GTK_PAPER_NAME_A4
	PAPER_NAME_A5        string = C.GTK_PAPER_NAME_A5
	PAPER_NAME_B5        string = C.GTK_PAPER_NAME_B5
	PAPER_NAME_LETTER    string = C.GTK_PAPER_NAME_LETTER
	PAPER_NAME_EXECUTIVE string = C.GTK_PAPER_NAME_EXECUTIVE
	PAPER_NAME_LEGAL     string = C.GTK_PAPER_NAME_LEGAL
)

// PaperSizeNew is a wrapper around gtk_paper_size_new().
func PaperSizeNew(name string) (*PaperSize, error) {
	var cstr *C.char

	if name != "" {
		cstr := C.CString(name)
		defer C.free(unsafe.Pointer(cstr))
	}

	c := C.gtk_paper_size_new((*C.gchar)(cstr))
	if c == nil {
		return nil, nilPtrErr
	}

	ps := wrapPaperSize(c)
	runtime.SetFinalizer(ps, (*PaperSize).free)
	return ps, nil
}

// PaperSizeNewFromPPD is a wrapper around gtk_paper_size_new_from_ppd().
func PaperSizeNewFromPPD(name, displayName string, width, height float64) (*PaperSize, error) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	cDisplayName := C.CString(displayName)
	defer C.free(unsafe.Pointer(cDisplayName))
	c := C.gtk_paper_size_new_from_ppd((*C.gchar)(cName), (*C.gchar)(cDisplayName),
		C.gdouble(width), C.gdouble(height))
	if c == nil {
		return nil, nilPtrErr
	}

	ps := wrapPaperSize(c)
	runtime.SetFinalizer(ps, (*PaperSize).free)
	return ps, nil
}

// PaperSizeNewCustom is a wrapper around gtk_paper_size_new_custom().
func PaperSizeNewCustom(name, displayName string, width, height float64, unit Unit) (*PaperSize, error) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	cDisplayName := C.CString(displayName)
	defer C.free(unsafe.Pointer(cDisplayName))
	c := C.gtk_paper_size_new_custom((*C.gchar)(cName), (*C.gchar)(cDisplayName),
		C.gdouble(width), C.gdouble(height), C.GtkUnit(unit))
	if c == nil {
		return nil, nilPtrErr
	}

	ps := wrapPaperSize(c)
	runtime.SetFinalizer(ps, (*PaperSize).free)
	return ps, nil
}

// Copy is a wrapper around gtk_paper_size_copy().
func (v *PaperSize) Copy() (*PaperSize, error) {
	c := C.gtk_paper_size_copy(v.native())
	if c == nil {
		return nil, nilPtrErr
	}

	p := wrapPaperSize(c)
	runtime.SetFinalizer(p, (*PaperSize).free)
	return p, nil
}

// free is a wrapper around gtk_paper_size_free().
func (v *PaperSize) free() {
	C.gtk_paper_size_free(v.native())
}

// IsEqual is a wrapper around gtk_paper_size_is_equal().
func (v *PaperSize) IsEqual(other *PaperSize) bool {
	c := C.gtk_paper_size_is_equal(v.native(), other.native())
	return gobool(c)
}

// PaperSizeGetPaperSizes is a wrapper around gtk_paper_size_get_paper_sizes().
// Returned list is wrapped to return *gtk.PaperSize elements.
func PaperSizeGetPaperSizes(includeCustom bool) *glib.List {
	clist := C.gtk_paper_size_get_paper_sizes(gbool(includeCustom))
	if clist == nil {
		return nil
	}

	glist := glib.WrapList(uintptr(unsafe.Pointer(clist)))
	glist.DataWrapper(func(ptr unsafe.Pointer) interface{} {
		ps := wrapPaperSize((*C.GtkPaperSize)(ptr))
		return ps
	})

	if glist != nil {
		runtime.SetFinalizer(glist, func(glist *glib.List) {
			glist.FreeFull(func(item interface{}) {
				ps := item.(*PaperSize)
				C.gtk_paper_size_free(ps.native())
			})
		})
	}

	return glist
}

// GetName is a wrapper around gtk_paper_size_get_name().
func (v *PaperSize) GetName() string {
	c := C.gtk_paper_size_get_name(v.native())
	return goString(c)
}

// GetDisplayName is a wrapper around gtk_paper_size_get_display_name().
func (v *PaperSize) GetDisplayName() string {
	c := C.gtk_paper_size_get_display_name(v.native())
	return goString(c)
}

// GetPPDName is a wrapper around gtk_paper_size_get_ppd_name().
func (v *PaperSize) GetPPDName() (string, error) {
	c := C.gtk_paper_size_get_ppd_name(v.native())
	if c == nil {
		return "", nilPtrErr
	}
	return goString(c), nil
}

// GetWidth is a wrapper around gtk_paper_size_get_width().
func (v *PaperSize) GetWidth(unit Unit) float64 {
	c := C.gtk_paper_size_get_width(v.native(), C.GtkUnit(unit))
	return float64(c)
}

// GetHeight is a wrapper around gtk_paper_size_get_height().
func (v *PaperSize) GetHeight(unit Unit) float64 {
	c := C.gtk_paper_size_get_width(v.native(), C.GtkUnit(unit))
	return float64(c)
}

// IsCustom is a wrapper around gtk_paper_size_is_custom().
func (v *PaperSize) IsCustom() bool {
	c := C.gtk_paper_size_is_custom(v.native())
	return gobool(c)
}

// SetSize is a wrapper around gtk_paper_size_set_size().
func (v *PaperSize) SetSize(width, height float64, unit Unit) {
	C.gtk_paper_size_set_size(v.native(), C.gdouble(width), C.gdouble(height), C.GtkUnit(unit))
}

// GetDefaultTopMargin is a wrapper around gtk_paper_size_get_default_top_margin().
func (v *PaperSize) GetDefaultTopMargin(unit Unit) float64 {
	c := C.gtk_paper_size_get_default_top_margin(v.native(), C.GtkUnit(unit))
	return float64(c)
}

// GetDefaultBottomMargin is a wrapper around gtk_paper_size_get_default_bottom_margin().
func (v *PaperSize) GetDefaultBottomMargin(unit Unit) float64 {
	c := C.gtk_paper_size_get_default_bottom_margin(v.native(), C.GtkUnit(unit))
	return float64(c)
}

// GetDefaultLeftMargin is a wrapper around gtk_paper_size_get_default_left_margin().
func (v *PaperSize) GetDefaultLeftMargin(unit Unit) float64 {
	c := C.gtk_paper_size_get_default_left_margin(v.native(), C.GtkUnit(unit))
	return float64(c)
}

// GetDefaultRightMargin is a wrapper around gtk_paper_size_get_default_right_margin().
func (v *PaperSize) GetDefaultRightMargin(unit Unit) float64 {
	c := C.gtk_paper_size_get_default_right_margin(v.native(), C.GtkUnit(unit))
	return float64(c)
}

// PaperSizeGetDefault is a wrapper around gtk_paper_size_get_default().
func PaperSizeGetDefaultRightMargin(unit Unit) string {
	c := C.gtk_paper_size_get_default()
	return goString(c)
}

// PaperSizeNewFromKeyFile() is a wrapper around gtk_paper_size_new_from_key_file().
// PaperSizeToKeyFile() is a wrapper around gtk_paper_size_to_key_file().

/*
 * GtkPrintContext
 */

// PrintContext is a representation of GTK's GtkPrintContext.
type PrintContext struct {
	*glib.Object
}

// native() returns a pointer to the underlying GtkPrintContext.
func (v *PrintContext) native() *C.GtkPrintContext {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toGtkPrintContext(ptr)
}

func marshalPrintContext(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapPrintContext(obj), nil
}

func wrapPrintContext(obj *glib.Object) *PrintContext {
	return &PrintContext{obj}
}

// GetCairoContext is a wrapper around gtk_print_context_get_cairo_context().
func (v *PrintContext) GetCairoContext() *cairo.Context {
	c := C.gtk_print_context_get_cairo_context(v.native())
	return cairo.WrapContext(uintptr(unsafe.Pointer(c)))
}

// SetCairoContext is a wrapper around gtk_print_context_set_cairo_context().
func (v *PrintContext) SetCairoContext(cr *cairo.Context, dpiX, dpiY float64) {
	C.gtk_print_context_set_cairo_context(v.native(),
		(*C.cairo_t)(unsafe.Pointer(cr.Native())),
		C.double(dpiX), C.double(dpiY))
}

// GetPageSetup is a wrapper around gtk_print_context_get_page_setup().
func (v *PrintContext) GetPageSetup() *PageSetup {
	c := C.gtk_print_context_get_page_setup(v.native())
	obj := glib.Take(unsafe.Pointer(c))
	return wrapPageSetup(obj)
}

// GetWidth is a wrapper around gtk_print_context_get_width().
func (v *PrintContext) GetWidth() float64 {
	c := C.gtk_print_context_get_width(v.native())
	return float64(c)
}

// GetHeight is a wrapper around gtk_print_context_get_height().
func (v *PrintContext) GetHeight() float64 {
	c := C.gtk_print_context_get_height(v.native())
	return float64(c)
}

// GetDpiX is a wrapper around gtk_print_context_get_dpi_x().
func (v *PrintContext) GetDpiX() float64 {
	c := C.gtk_print_context_get_dpi_x(v.native())
	return float64(c)
}

// GetDpiY is a wrapper around gtk_print_context_get_dpi_y().
func (v *PrintContext) GetDpiY() float64 {
	c := C.gtk_print_context_get_dpi_y(v.native())
	return float64(c)
}

// GetPangoFontMap is a wrapper around gtk_print_context_get_pango_fontmap().
func (v *PrintContext) GetPangoFontMap() *pango.FontMap {
	c := C.gtk_print_context_get_pango_fontmap(v.native())
	return pango.WrapFontMap(uintptr(unsafe.Pointer(c)))
}

// CreatePangoContext is a wrapper around gtk_print_context_create_pango_context().
func (v *PrintContext) CreatePangoContext() *pango.Context {
	c := C.gtk_print_context_create_pango_context(v.native())
	return pango.WrapContext(uintptr(unsafe.Pointer(c)))
}

// CreatePangoLayout is a wrapper around gtk_print_context_create_pango_layout().
func (v *PrintContext) CreatePangoLayout() *pango.Layout {
	c := C.gtk_print_context_create_pango_layout(v.native())
	return pango.WrapLayout(uintptr(unsafe.Pointer(c)))
}

// GetHardMargins is a wrapper around gtk_print_context_get_hard_margins().
func (v *PrintContext) GetHardMargins() (float64, float64, float64, float64, error) {
	var top, bottom, left, right C.gdouble
	c := C.gtk_print_context_get_hard_margins(v.native(), &top, &bottom, &left, &right)
	if gobool(c) == false {
		return 0.0, 0.0, 0.0, 0.0, errors.New("unable to retrieve hard margins")
	}
	return float64(top), float64(bottom), float64(left), float64(right), nil
}

/*
 * GtkPrintOperation
 */
type PrintOperation struct {
	*glib.Object
	// Interfaces
	PrintOperationPreview
}

func (v *PrintOperation) native() *C.GtkPrintOperation {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toGtkPrintOperation(ptr)
}

func (v *PrintOperation) toPrintOperationPreview() *C.GtkPrintOperationPreview {
	if v == nil {
		return nil
	}
	return C.toGtkPrintOperationPreview(unsafe.Pointer(v.Native()))
}

func marshalPrintOperation(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapPrintOperation(obj), nil
}

func wrapPrintOperation(obj *glib.Object) *PrintOperation {
	pop := wrapPrintOperationPreview(*glib.InterfaceFromObjectNew(obj))
	return &PrintOperation{obj, *pop}
}

// PrintOperationNew is a wrapper around gtk_print_operation_new().
func PrintOperationNew() (*PrintOperation, error) {
	c := C.gtk_print_operation_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapPrintOperation(obj), nil
}

// SetAllowAsync is a wrapper around gtk_print_operation_set_allow_async().
func (v *PrintOperation) PrintOperationSetAllowAsync(allowSync bool) {
	C.gtk_print_operation_set_allow_async(v.native(), gbool(allowSync))
}

// GetError is a wrapper around gtk_print_operation_get_error().
func (v *PrintOperation) PrintOperationGetError() error {
	var err *C.GError
	C.gtk_print_operation_get_error(v.native(), &err)
	defer C.g_error_free(err)
	return errors.New(goString(err.message))
}

// SetDefaultPageSetup is a wrapper around gtk_print_operation_set_default_page_setup().
func (v *PrintOperation) SetDefaultPageSetup(ps *PageSetup) {
	C.gtk_print_operation_set_default_page_setup(v.native(), ps.native())
}

// GetDefaultPageSetup is a wrapper around gtk_print_operation_get_default_page_setup().
func (v *PrintOperation) GetDefaultPageSetup() (*PageSetup, error) {
	c := C.gtk_print_operation_get_default_page_setup(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapPageSetup(obj), nil
}

// SetPrintSettings is a wrapper around gtk_print_operation_set_print_settings().
func (v *PrintOperation) SetPrintSettings(ps *PrintSettings) {
	C.gtk_print_operation_set_print_settings(v.native(), ps.native())
}

// GetPrintSettings is a wrapper around gtk_print_operation_get_print_settings().
func (v *PrintOperation) GetPrintSettings() (*PrintSettings, error) {
	c := C.gtk_print_operation_get_print_settings(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapPrintSettings(obj), nil
}

// SetJobName is a wrapper around gtk_print_operation_set_job_name().
func (v *PrintOperation) SetJobName(name string) {
	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_print_operation_set_job_name(v.native(), (*C.gchar)(cstr))
}

// SetNPages is a wrapper around gtk_print_operation_set_n_pages().
func (v *PrintOperation) SetNPages(pages int) {
	C.gtk_print_operation_set_n_pages(v.native(), C.gint(pages))
}

// GetNPagesToPrint is a wrapper around gtk_print_operation_get_n_pages_to_print().
func (v *PrintOperation) GetNPagesToPrint() int {
	c := C.gtk_print_operation_get_n_pages_to_print(v.native())
	return int(c)
}

// SetCurrentPage is a wrapper around gtk_print_operation_set_current_page().
func (v *PrintOperation) SetCurrentPage(page int) {
	C.gtk_print_operation_set_current_page(v.native(), C.gint(page))
}

// SetUseFullPage is a wrapper around gtk_print_operation_set_use_full_page().
func (v *PrintOperation) SetUseFullPage(full bool) {
	C.gtk_print_operation_set_use_full_page(v.native(), gbool(full))
}

// SetUnit is a wrapper around gtk_print_operation_set_unit().
func (v *PrintOperation) SetUnit(unit Unit) {
	C.gtk_print_operation_set_unit(v.native(), C.GtkUnit(unit))
}

// SetExportFilename is a wrapper around gtk_print_operation_set_export_filename().
func (v *PrintOperation) SetExportFilename(name string) {
	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_print_operation_set_export_filename(v.native(), (*C.gchar)(cstr))
}

// SetShowProgress is a wrapper around gtk_print_operation_set_show_progress().
func (v *PrintOperation) SetShowProgress(show bool) {
	C.gtk_print_operation_set_show_progress(v.native(), gbool(show))
}

// SetTrackPrintStatus is a wrapper around gtk_print_operation_set_track_print_status().
func (v *PrintOperation) SetTrackPrintStatus(progress bool) {
	C.gtk_print_operation_set_track_print_status(v.native(), gbool(progress))
}

// SetCustomTabLabel is a wrapper around gtk_print_operation_set_custom_tab_label().
func (v *PrintOperation) SetCustomTabLabel(label string) {
	cstr := C.CString(label)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_print_operation_set_custom_tab_label(v.native(), (*C.gchar)(cstr))
}

// Run is a wrapper around gtk_print_operation_run().
func (v *PrintOperation) Run(action PrintOperationAction, parent *Window) (PrintOperationResult, error) {
	var err *C.GError = nil
	c := C.gtk_print_operation_run(v.native(), C.GtkPrintOperationAction(action), parent.native(), &err)
	res := PrintOperationResult(c)
	if res == PRINT_OPERATION_RESULT_ERROR {
		defer C.g_error_free(err)
		return res, errors.New(goString(err.message))
	}
	return res, nil
}

// Cancel is a wrapper around gtk_print_operation_cancel().
func (v *PrintOperation) Cancel() {
	C.gtk_print_operation_cancel(v.native())
}

// DrawPageFinish is a wrapper around gtk_print_operation_draw_page_finish().
func (v *PrintOperation) DrawPageFinish() {
	C.gtk_print_operation_draw_page_finish(v.native())
}

// SetDeferDrawing is a wrapper around gtk_print_operation_set_defer_drawing().
func (v *PrintOperation) SetDeferDrawing() {
	C.gtk_print_operation_set_defer_drawing(v.native())
}

// GetStatus is a wrapper around gtk_print_operation_get_status().
func (v *PrintOperation) GetStatus() PrintStatus {
	c := C.gtk_print_operation_get_status(v.native())
	return PrintStatus(c)
}

// GetStatusString is a wrapper around gtk_print_operation_get_status_string().
func (v *PrintOperation) GetStatusString() string {
	c := C.gtk_print_operation_get_status_string(v.native())
	return goString(c)
}

// IsFinished is a wrapper around gtk_print_operation_is_finished().
func (v *PrintOperation) IsFinished() bool {
	c := C.gtk_print_operation_is_finished(v.native())
	return gobool(c)
}

// SetSupportSelection is a wrapper around gtk_print_operation_set_support_selection().
func (v *PrintOperation) SetSupportSelection(selection bool) {
	C.gtk_print_operation_set_support_selection(v.native(), gbool(selection))
}

// GetSupportSelection is a wrapper around gtk_print_operation_get_support_selection().
func (v *PrintOperation) GetSupportSelection() bool {
	c := C.gtk_print_operation_get_support_selection(v.native())
	return gobool(c)
}

// SetHasSelection is a wrapper around gtk_print_operation_set_has_selection().
func (v *PrintOperation) SetHasSelection(selection bool) {
	C.gtk_print_operation_set_has_selection(v.native(), gbool(selection))
}

// GetHasSelection is a wrapper around gtk_print_operation_get_has_selection().
func (v *PrintOperation) GetHasSelection() bool {
	c := C.gtk_print_operation_get_has_selection(v.native())
	return gobool(c)
}

// SetEmbedPageSetup is a wrapper around gtk_print_operation_set_embed_page_setup().
func (v *PrintOperation) SetEmbedPageSetup(embed bool) {
	C.gtk_print_operation_set_embed_page_setup(v.native(), gbool(embed))
}

// GetEmbedPageSetup is a wrapper around gtk_print_operation_get_embed_page_setup().
func (v *PrintOperation) GetEmbedPageSetup() bool {
	c := C.gtk_print_operation_get_embed_page_setup(v.native())
	return gobool(c)
}

// PrintRunPageSetupDialog is a wrapper around gtk_print_run_page_setup_dialog().
func PrintRunPageSetupDialog(parent *Window, pageSetup *PageSetup, settings *PrintSettings) *PageSetup {
	c := C.gtk_print_run_page_setup_dialog(parent.native(), pageSetup.native(), settings.native())
	obj := glib.Take(unsafe.Pointer(c))
	return wrapPageSetup(obj)
}

type PageSetupDoneCallback func(setup *PageSetup, userData uintptr)

type pageSetupDoneCallbackData struct {
	fn   PageSetupDoneCallback
	data uintptr
}

var (
	pageSetupDoneCallbackRegistry = struct {
		sync.RWMutex
		next int
		m    map[int]pageSetupDoneCallbackData
	}{
		next: 1,
		m:    make(map[int]pageSetupDoneCallbackData),
	}
)

// PrintRunPageSetupDialogAsync is a wrapper around gtk_print_run_page_setup_dialog_async().
func PrintRunPageSetupDialogAsync(parent *Window, setup *PageSetup,
	settings *PrintSettings, cb PageSetupDoneCallback, data uintptr) {

	pageSetupDoneCallbackRegistry.Lock()
	id := pageSetupDoneCallbackRegistry.next
	pageSetupDoneCallbackRegistry.next++
	pageSetupDoneCallbackRegistry.m[id] =
		pageSetupDoneCallbackData{fn: cb, data: data}
	pageSetupDoneCallbackRegistry.Unlock()

	C._gtk_print_run_page_setup_dialog_async(parent.native(), setup.native(),
		settings.native(), C.gpointer(uintptr(id)))
}

/*
 * GtkPrintOperationPreview
 */

// PrintOperationPreview is a representation of GTK's GtkPrintOperationPreview GInterface.
type PrintOperationPreview struct {
	glib.Interface
}

// IPrintOperationPreview is an interface type implemented by all structs
// embedding a PrintOperationPreview.  It is meant to be used as an argument type
// for wrapper functions that wrap around a C GTK function taking a
// GtkPrintOperationPreview.
type IPrintOperationPreview interface {
	toPrintOperationPreview() *C.GtkPrintOperationPreview
}

// native() returns a pointer to the underlying GObject as a GtkPrintOperationPreview.
func (v *PrintOperationPreview) native() *C.GtkPrintOperationPreview {
	return C.toGtkPrintOperationPreview(unsafe.Pointer(v.Native()))
}

func marshalPrintOperationPreview(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapPrintOperationPreview(*glib.InterfaceFromObjectNew(obj)), nil
}

func wrapPrintOperationPreview(intf glib.Interface) *PrintOperationPreview {
	return &PrintOperationPreview{intf}
}

func (v *PrintOperationPreview) toPrintOperationPreview() *C.GtkPrintOperationPreview {
	if v == nil {
		return nil
	}
	return v.native()
}

// RenderPage is a wrapper around gtk_print_operation_preview_render_page().
func (v *PrintOperationPreview) RenderPage(page int) {
	C.gtk_print_operation_preview_render_page(v.native(), C.gint(page))
}

// EndPreview is a wrapper around gtk_print_operation_preview_end_preview().
func (v *PrintOperationPreview) EndPreview() {
	C.gtk_print_operation_preview_end_preview(v.native())
}

// IsSelected is a wrapper around gtk_print_operation_preview_is_selected().
func (v *PrintOperationPreview) IsSelected(page int) bool {
	c := C.gtk_print_operation_preview_is_selected(v.native(), C.gint(page))
	return gobool(c)
}

/*
 * GtkPrintSettings
 */

type PrintSettings struct {
	*glib.Object
}

func (v *PrintSettings) native() *C.GtkPrintSettings {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toGtkPrintSettings(ptr)
}

func marshalPrintSettings(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapPrintSettings(obj), nil
}

func wrapPrintSettings(obj *glib.Object) *PrintSettings {
	return &PrintSettings{obj}
}

const (
	PRINT_SETTINGS_PRINTER              string = C.GTK_PRINT_SETTINGS_PRINTER
	PRINT_SETTINGS_ORIENTATION          string = C.GTK_PRINT_SETTINGS_ORIENTATION
	PRINT_SETTINGS_PAPER_FORMAT         string = C.GTK_PRINT_SETTINGS_PAPER_FORMAT
	PRINT_SETTINGS_PAPER_WIDTH          string = C.GTK_PRINT_SETTINGS_PAPER_WIDTH
	PRINT_SETTINGS_PAPER_HEIGHT         string = C.GTK_PRINT_SETTINGS_PAPER_HEIGHT
	PRINT_SETTINGS_USE_COLOR            string = C.GTK_PRINT_SETTINGS_USE_COLOR
	PRINT_SETTINGS_COLLATE              string = C.GTK_PRINT_SETTINGS_COLLATE
	PRINT_SETTINGS_REVERSE              string = C.GTK_PRINT_SETTINGS_REVERSE
	PRINT_SETTINGS_DUPLEX               string = C.GTK_PRINT_SETTINGS_DUPLEX
	PRINT_SETTINGS_QUALITY              string = C.GTK_PRINT_SETTINGS_QUALITY
	PRINT_SETTINGS_N_COPIES             string = C.GTK_PRINT_SETTINGS_N_COPIES
	PRINT_SETTINGS_NUMBER_UP            string = C.GTK_PRINT_SETTINGS_NUMBER_UP
	PRINT_SETTINGS_NUMBER_UP_LAYOUT     string = C.GTK_PRINT_SETTINGS_NUMBER_UP_LAYOUT
	PRINT_SETTINGS_RESOLUTION           string = C.GTK_PRINT_SETTINGS_RESOLUTION
	PRINT_SETTINGS_RESOLUTION_X         string = C.GTK_PRINT_SETTINGS_RESOLUTION_X
	PRINT_SETTINGS_RESOLUTION_Y         string = C.GTK_PRINT_SETTINGS_RESOLUTION_Y
	PRINT_SETTINGS_PRINTER_LPI          string = C.GTK_PRINT_SETTINGS_PRINTER_LPI
	PRINT_SETTINGS_SCALE                string = C.GTK_PRINT_SETTINGS_SCALE
	PRINT_SETTINGS_PRINT_PAGES          string = C.GTK_PRINT_SETTINGS_PRINT_PAGES
	PRINT_SETTINGS_PAGE_RANGES          string = C.GTK_PRINT_SETTINGS_PAGE_RANGES
	PRINT_SETTINGS_PAGE_SET             string = C.GTK_PRINT_SETTINGS_PAGE_SET
	PRINT_SETTINGS_DEFAULT_SOURCE       string = C.GTK_PRINT_SETTINGS_DEFAULT_SOURCE
	PRINT_SETTINGS_MEDIA_TYPE           string = C.GTK_PRINT_SETTINGS_MEDIA_TYPE
	PRINT_SETTINGS_DITHER               string = C.GTK_PRINT_SETTINGS_DITHER
	PRINT_SETTINGS_FINISHINGS           string = C.GTK_PRINT_SETTINGS_FINISHINGS
	PRINT_SETTINGS_OUTPUT_BIN           string = C.GTK_PRINT_SETTINGS_OUTPUT_BIN
	PRINT_SETTINGS_OUTPUT_DIR           string = C.GTK_PRINT_SETTINGS_OUTPUT_DIR
	PRINT_SETTINGS_OUTPUT_BASENAME      string = C.GTK_PRINT_SETTINGS_OUTPUT_BASENAME
	PRINT_SETTINGS_OUTPUT_FILE_FORMAT   string = C.GTK_PRINT_SETTINGS_OUTPUT_FILE_FORMAT
	PRINT_SETTINGS_OUTPUT_URI           string = C.GTK_PRINT_SETTINGS_OUTPUT_URI
	PRINT_SETTINGS_WIN32_DRIVER_EXTRA   string = C.GTK_PRINT_SETTINGS_WIN32_DRIVER_EXTRA
	PRINT_SETTINGS_WIN32_DRIVER_VERSION string = C.GTK_PRINT_SETTINGS_WIN32_DRIVER_VERSION
)

// PrintSettingsNew is a wrapper around gtk_print_settings_new().
func PrintSettingsNew() (*PrintSettings, error) {
	c := C.gtk_print_settings_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapPrintSettings(obj), nil
}

// Copy is a wrapper around gtk_print_settings_copy().
func (v *PrintSettings) Copy() (*PrintSettings, error) {
	c := C.gtk_print_settings_copy(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapPrintSettings(obj), nil
}

// HasKey is a wrapper around gtk_print_settings_has_key().
func (v *PrintSettings) HasKey(key string) bool {
	cstr := C.CString(key)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_print_settings_has_key(v.native(), (*C.gchar)(cstr))
	return gobool(c)
}

// Get is a wrapper around gtk_print_settings_get().
func (v *PrintSettings) Get(key string) string {
	cstr := C.CString(key)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_print_settings_get(v.native(), (*C.gchar)(cstr))
	return goString(c)
}

// Set is a wrapper around gtk_print_settings_set().
// TODO: Since value can't be nil, we can't unset values here.
func (v *PrintSettings) Set(key, value string) {
	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))
	cValue := C.CString(value)
	defer C.free(unsafe.Pointer(cValue))
	C.gtk_print_settings_set(v.native(), (*C.gchar)(cKey), (*C.gchar)(cValue))
}

// Unset is a wrapper around gtk_print_settings_unset().
func (v *PrintSettings) Unset(key string) {
	cstr := C.CString(key)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_print_settings_unset(v.native(), (*C.gchar)(cstr))
}

type PrintSettingsCallback func(key, value string, userData uintptr)

type printSettingsCallbackData struct {
	fn       PrintSettingsCallback
	userData uintptr
}

var (
	printSettingsCallbackRegistry = struct {
		sync.RWMutex
		next int
		m    map[int]printSettingsCallbackData
	}{
		next: 1,
		m:    make(map[int]printSettingsCallbackData),
	}
)

// Foreach is a wrapper around gtk_print_settings_foreach().
func (v *PrintSettings) ForEach(cb PrintSettingsCallback, userData uintptr) {
	printSettingsCallbackRegistry.Lock()
	id := printSettingsCallbackRegistry.next
	printSettingsCallbackRegistry.next++
	printSettingsCallbackRegistry.m[id] =
		printSettingsCallbackData{fn: cb, userData: userData}
	printSettingsCallbackRegistry.Unlock()

	C._gtk_print_settings_foreach(v.native(), C.gpointer(uintptr(id)))
}

// GetBool is a wrapper around gtk_print_settings_get_bool().
func (v *PrintSettings) GetBool(key string) bool {
	cstr := C.CString(key)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_print_settings_get_bool(v.native(), (*C.gchar)(cstr))
	return gobool(c)
}

// SetBool is a wrapper around gtk_print_settings_set_bool().
func (v *PrintSettings) SetBool(key string, value bool) {
	cstr := C.CString(key)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_print_settings_set_bool(v.native(), (*C.gchar)(cstr), gbool(value))
}

// GetDouble is a wrapper around gtk_print_settings_get_double().
func (v *PrintSettings) GetDouble(key string) float64 {
	cstr := C.CString(key)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_print_settings_get_double(v.native(), (*C.gchar)(cstr))
	return float64(c)
}

// GetDoubleWithDefault is a wrapper around gtk_print_settings_get_double_with_default().
func (v *PrintSettings) GetDoubleWithDefault(key string, def float64) float64 {
	cstr := C.CString(key)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_print_settings_get_double_with_default(v.native(),
		(*C.gchar)(cstr), C.gdouble(def))
	return float64(c)
}

// SetDouble is a wrapper around gtk_print_settings_set_double().
func (v *PrintSettings) SetDouble(key string, value float64) {
	cstr := C.CString(key)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_print_settings_set_double(v.native(), (*C.gchar)(cstr), C.gdouble(value))
}

// GetLength is a wrapper around gtk_print_settings_get_length().
func (v *PrintSettings) GetLength(key string, unit Unit) float64 {
	cstr := C.CString(key)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_print_settings_get_length(v.native(), (*C.gchar)(cstr), C.GtkUnit(unit))
	return float64(c)
}

// SetLength is a wrapper around gtk_print_settings_set_length().
func (v *PrintSettings) SetLength(key string, value float64, unit Unit) {
	cstr := C.CString(key)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_print_settings_set_length(v.native(), (*C.gchar)(cstr), C.gdouble(value), C.GtkUnit(unit))
}

// GetInt is a wrapper around gtk_print_settings_get_int().
func (v *PrintSettings) GetInt(key string) int {
	cstr := C.CString(key)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_print_settings_get_int(v.native(), (*C.gchar)(cstr))
	return int(c)
}

// GetIntWithDefault is a wrapper around gtk_print_settings_get_int_with_default().
func (v *PrintSettings) GetIntWithDefault(key string, def int) int {
	cstr := C.CString(key)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_print_settings_get_int_with_default(v.native(), (*C.gchar)(cstr), C.gint(def))
	return int(c)
}

// SetInt is a wrapper around gtk_print_settings_set_int().
func (v *PrintSettings) SetInt(key string, value int) {
	cstr := C.CString(key)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_print_settings_set_int(v.native(), (*C.gchar)(cstr), C.gint(value))
}

// GetPrinter is a wrapper around gtk_print_settings_get_printer().
func (v *PrintSettings) GetPrinter() string {
	c := C.gtk_print_settings_get_printer(v.native())
	return goString(c)
}

// SetPrinter is a wrapper around gtk_print_settings_set_printer().
func (v *PrintSettings) SetPrinter(printer string) {
	cstr := C.CString(printer)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_print_settings_set_printer(v.native(), (*C.gchar)(cstr))
}

// GetOrientation is a wrapper around gtk_print_settings_get_orientation().
func (v *PrintSettings) GetOrientation() PageOrientation {
	c := C.gtk_print_settings_get_orientation(v.native())
	return PageOrientation(c)
}

// SetOrientation is a wrapper around gtk_print_settings_set_orientation().
func (v *PrintSettings) SetOrientation(orientation PageOrientation) {
	C.gtk_print_settings_set_orientation(v.native(), C.GtkPageOrientation(orientation))
}

// GetPaperSize is a wrapper around gtk_print_settings_get_paper_size().
func (v *PrintSettings) GetPaperSize() (*PaperSize, error) {
	c := C.gtk_print_settings_get_paper_size(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	p := wrapPaperSize(c)
	runtime.SetFinalizer(p, (*PaperSize).free)
	return p, nil
}

// SetPaperSize is a wrapper around gtk_print_settings_set_paper_size().
func (v *PrintSettings) SetPaperSize(size *PaperSize) {
	C.gtk_print_settings_set_paper_size(v.native(), size.native())
}

// GetPaperWidth is a wrapper around gtk_print_settings_get_paper_width().
func (v *PrintSettings) GetPaperWidth(unit Unit) float64 {
	c := C.gtk_print_settings_get_paper_width(v.native(), C.GtkUnit(unit))
	return float64(c)
}

// SetPaperWidth is a wrapper around gtk_print_settings_set_paper_width().
func (v *PrintSettings) SetPaperWidth(width float64, unit Unit) {
	C.gtk_print_settings_set_paper_width(v.native(), C.gdouble(width), C.GtkUnit(unit))
}

// GetPaperHeight is a wrapper around gtk_print_settings_get_paper_height().
func (v *PrintSettings) GetPaperHeight(unit Unit) float64 {
	c := C.gtk_print_settings_get_paper_height(v.native(), C.GtkUnit(unit))
	return float64(c)
}

// SetPaperHeight is a wrapper around gtk_print_settings_set_paper_height().
func (v *PrintSettings) SetPaperHeight(width float64, unit Unit) {
	C.gtk_print_settings_set_paper_height(v.native(), C.gdouble(width), C.GtkUnit(unit))
}

// GetUseColor is a wrapper around gtk_print_settings_get_use_color().
func (v *PrintSettings) GetUseColor() bool {
	c := C.gtk_print_settings_get_use_color(v.native())
	return gobool(c)
}

// SetUseColor is a wrapper around gtk_print_settings_set_use_color().
func (v *PrintSettings) SetUseColor(color bool) {
	C.gtk_print_settings_set_use_color(v.native(), gbool(color))
}

// GetCollate is a wrapper around gtk_print_settings_get_collate().
func (v *PrintSettings) GetCollate() bool {
	c := C.gtk_print_settings_get_collate(v.native())
	return gobool(c)
}

// SetCollate is a wrapper around gtk_print_settings_set_collate().
func (v *PrintSettings) SetCollate(collate bool) {
	C.gtk_print_settings_set_collate(v.native(), gbool(collate))
}

// GetReverse is a wrapper around gtk_print_settings_get_reverse().
func (v *PrintSettings) GetReverse() bool {
	c := C.gtk_print_settings_get_reverse(v.native())
	return gobool(c)
}

// SetReverse is a wrapper around gtk_print_settings_set_reverse().
func (v *PrintSettings) SetReverse(reverse bool) {
	C.gtk_print_settings_set_reverse(v.native(), gbool(reverse))
}

// GetDuplex is a wrapper around gtk_print_settings_get_duplex().
func (v *PrintSettings) GetDuplex() PrintDuplex {
	c := C.gtk_print_settings_get_duplex(v.native())
	return PrintDuplex(c)
}

// SetDuplex is a wrapper around gtk_print_settings_set_duplex().
func (v *PrintSettings) SetDuplex(duplex PrintDuplex) {
	C.gtk_print_settings_set_duplex(v.native(), C.GtkPrintDuplex(duplex))
}

// GetQuality is a wrapper around gtk_print_settings_get_quality().
func (v *PrintSettings) GetQuality() PrintQuality {
	c := C.gtk_print_settings_get_quality(v.native())
	return PrintQuality(c)
}

// SetQuality is a wrapper around gtk_print_settings_set_quality().
func (v *PrintSettings) SetQuality(quality PrintQuality) {
	C.gtk_print_settings_set_quality(v.native(), C.GtkPrintQuality(quality))
}

// GetNCopies is a wrapper around gtk_print_settings_get_n_copies().
func (v *PrintSettings) GetNCopies() int {
	c := C.gtk_print_settings_get_n_copies(v.native())
	return int(c)
}

// SetNCopies is a wrapper around gtk_print_settings_set_n_copies().
func (v *PrintSettings) SetNCopies(copies int) {
	C.gtk_print_settings_set_n_copies(v.native(), C.gint(copies))
}

// GetNmberUp is a wrapper around gtk_print_settings_get_number_up().
func (v *PrintSettings) GetNmberUp() int {
	c := C.gtk_print_settings_get_number_up(v.native())
	return int(c)
}

// SetNumberUp is a wrapper around gtk_print_settings_set_number_up().
func (v *PrintSettings) SetNumberUp(numberUp int) {
	C.gtk_print_settings_set_number_up(v.native(), C.gint(numberUp))
}

// GetNumberUpLayout is a wrapper around gtk_print_settings_get_number_up_layout().
func (v *PrintSettings) GetNumberUpLayout() NumberUpLayout {
	c := C.gtk_print_settings_get_number_up_layout(v.native())
	return NumberUpLayout(c)
}

// SetNumberUpLayout is a wrapper around gtk_print_settings_set_number_up_layout().
func (v *PrintSettings) SetNumberUpLayout(numberUpLayout NumberUpLayout) {
	C.gtk_print_settings_set_number_up_layout(v.native(), C.GtkNumberUpLayout(numberUpLayout))
}

// GetResolution is a wrapper around gtk_print_settings_get_resolution().
func (v *PrintSettings) GetResolution() int {
	c := C.gtk_print_settings_get_resolution(v.native())
	return int(c)
}

// SetResolution is a wrapper around gtk_print_settings_set_resolution().
func (v *PrintSettings) SetResolution(resolution int) {
	C.gtk_print_settings_set_resolution(v.native(), C.gint(resolution))
}

// SetResolutionXY is a wrapper around gtk_print_settings_set_resolution_xy().
func (v *PrintSettings) SetResolutionXY(resolutionX, resolutionY int) {
	C.gtk_print_settings_set_resolution_xy(v.native(), C.gint(resolutionX), C.gint(resolutionY))
}

// GetResolutionX is a wrapper around gtk_print_settings_get_resolution_x().
func (v *PrintSettings) GetResolutionX() int {
	c := C.gtk_print_settings_get_resolution_x(v.native())
	return int(c)
}

// GetResolutionY is a wrapper around gtk_print_settings_get_resolution_y().
func (v *PrintSettings) GetResolutionY() int {
	c := C.gtk_print_settings_get_resolution_y(v.native())
	return int(c)
}

// GetPrinterLpi is a wrapper around gtk_print_settings_get_printer_lpi().
func (v *PrintSettings) GetPrinterLpi() float64 {
	c := C.gtk_print_settings_get_printer_lpi(v.native())
	return float64(c)
}

// SetPrinterLpi is a wrapper around gtk_print_settings_set_printer_lpi().
func (v *PrintSettings) SetPrinterLpi(lpi float64) {
	C.gtk_print_settings_set_printer_lpi(v.native(), C.gdouble(lpi))
}

// GetScale is a wrapper around gtk_print_settings_get_scale().
func (v *PrintSettings) GetScale() float64 {
	c := C.gtk_print_settings_get_scale(v.native())
	return float64(c)
}

// SetScale is a wrapper around gtk_print_settings_set_scale().
func (v *PrintSettings) SetScale(scale float64) {
	C.gtk_print_settings_set_scale(v.native(), C.gdouble(scale))
}

// GetPrintPages is a wrapper around gtk_print_settings_get_print_pages().
func (v *PrintSettings) GetPrintPages() PrintPages {
	c := C.gtk_print_settings_get_print_pages(v.native())
	return PrintPages(c)
}

// SetPrintPages is a wrapper around gtk_print_settings_set_print_pages().
func (v *PrintSettings) SetPrintPages(pages PrintPages) {
	C.gtk_print_settings_set_print_pages(v.native(), C.GtkPrintPages(pages))
}

// GetPageSet is a wrapper around gtk_print_settings_get_page_set().
func (v *PrintSettings) GetPageSet(pages PrintPages) PageSet {
	c := C.gtk_print_settings_get_page_set(v.native())
	return PageSet(c)
}

// SetPageSet is a wrapper around gtk_print_settings_set_page_set().
func (v *PrintSettings) SetPageSet(pageSet PageSet) {
	C.gtk_print_settings_set_page_set(v.native(), C.GtkPageSet(pageSet))
}

// GetDefaultSource is a wrapper around gtk_print_settings_get_default_source().
func (v *PrintSettings) GetDefaultSource() string {
	c := C.gtk_print_settings_get_default_source(v.native())
	return goString(c)
}

// SetSefaultSource is a wrapper around gtk_print_settings_set_default_source().
func (v *PrintSettings) SetSefaultSource(defaultSource string) {
	cstr := C.CString(defaultSource)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_print_settings_set_default_source(v.native(), (*C.gchar)(cstr))
}

// GetMediaType is a wrapper around gtk_print_settings_get_media_type().
func (v *PrintSettings) GetMediaType() string {
	c := C.gtk_print_settings_get_media_type(v.native())
	return goString(c)
}

// SetMediaType is a wrapper around gtk_print_settings_set_media_type().
func (v *PrintSettings) SetMediaType(mediaType string) {
	cstr := C.CString(mediaType)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_print_settings_set_media_type(v.native(), (*C.gchar)(cstr))
}

// GetDither is a wrapper around gtk_print_settings_get_dither().
func (v *PrintSettings) GetDither() string {
	c := C.gtk_print_settings_get_dither(v.native())
	return goString(c)
}

// SetDither is a wrapper around gtk_print_settings_set_dither().
func (v *PrintSettings) SetDither(dither string) {
	cstr := C.CString(dither)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_print_settings_set_dither(v.native(), (*C.gchar)(cstr))
}

// GetFinishings is a wrapper around gtk_print_settings_get_finishings().
func (v *PrintSettings) GetFinishings() string {
	c := C.gtk_print_settings_get_finishings(v.native())
	return goString(c)
}

// SetFinishings is a wrapper around gtk_print_settings_set_finishings().
func (v *PrintSettings) SetFinishings(dither string) {
	cstr := C.CString(dither)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_print_settings_set_finishings(v.native(), (*C.gchar)(cstr))
}

// GetOutputBin is a wrapper around gtk_print_settings_get_output_bin().
func (v *PrintSettings) GetOutputBin() string {
	c := C.gtk_print_settings_get_output_bin(v.native())
	return goString(c)
}

// SetOutputBin is a wrapper around gtk_print_settings_set_output_bin().
func (v *PrintSettings) SetOutputBin(bin string) {
	cstr := C.CString(bin)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_print_settings_set_output_bin(v.native(), (*C.gchar)(cstr))
}

// PrintSettingsNewFromFile is a wrapper around gtk_print_settings_new_from_file().
func PrintSettingsNewFromFile(name string) (*PrintSettings, error) {
	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))
	var err *C.GError
	c := C.gtk_print_settings_new_from_file((*C.gchar)(cstr), &err)
	if c == nil {
		defer C.g_error_free(err)
		return nil, errors.New(goString(err.message))
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapPrintSettings(obj), nil
}

// PrintSettingsNewFromKeyFile() is a wrapper around gtk_print_settings_new_from_key_file().

// LoadFile is a wrapper around gtk_print_settings_load_file().
func (v *PrintSettings) LoadFile(name string) error {
	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))
	var err *C.GError
	c := C.gtk_print_settings_load_file(v.native(), (*C.gchar)(cstr), &err)
	if gobool(c) == false {
		defer C.g_error_free(err)
		return errors.New(goString(err.message))
	}
	return nil
}

// LoadKeyFile() is a wrapper around gtk_print_settings_load_key_file().

// ToFile is a wrapper around gtk_print_settings_to_file().
func (v *PrintSettings) ToFile(name string) error {
	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))
	var err *C.GError
	c := C.gtk_print_settings_to_file(v.native(), (*C.gchar)(cstr), &err)
	if gobool(c) == false {
		return errors.New(goString(err.message))
	}
	return nil
}

// ToKeyFile() is a wrapper around gtk_print_settings_to_key_file().
