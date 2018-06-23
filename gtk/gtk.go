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

// Go bindings for GTK+ 3.  Supports version 3.6 and later.
//
// Functions use the same names as the native C function calls, but use
// CamelCase.  In cases where native GTK uses pointers to values to
// simulate multiple return values, Go's native multiple return values
// are used instead.  Whenever a native GTK call could return an
// unexpected NULL pointer, an additonal error is returned in the Go
// binding.
//
// GTK's C API documentation can be very useful for understanding how the
// functions in this package work and what each type is for.  This
// documentation can be found at https://developer.gnome.org/gtk3/.
//
// In addition to Go versions of the C GTK functions, every struct type
// includes a method named Native (either by direct implementation, or
// by means of struct embedding).  These methods return a uintptr of the
// native C object the binding type represents.  These pointers may be
// type switched to a native C pointer using unsafe and used with cgo
// function calls outside this package.
//
// Memory management is handled in proper Go fashion, using runtime
// finalizers to properly free memory when it is no longer needed.  Each
// time a Go type is created with a pointer to a GObject, a reference is
// added for Go, sinking the floating reference when necessary.  After
// going out of scope and the next time Go's garbage collector is run, a
// finalizer is run to remove Go's reference to the GObject.  When this
// reference count hits zero (when neither Go nor GTK holds ownership)
// the object will be freed internally by GTK.
package gtk

// #cgo pkg-config: gtk+-3.0
// #include <gtk/gtk.h>
// #include "gtk.go.h"
import "C"
import (
	"errors"
	"fmt"
	"reflect"
	"runtime"
	"sync"
	"unsafe"

	"github.com/d2r2/gotk3/gdk"
	"github.com/d2r2/gotk3/glib"
)

func init() {
	tm := []glib.TypeMarshaler{
		// Enums
		{glib.Type(C.gtk_align_get_type()), marshalAlign},
		{glib.Type(C.gtk_accel_flags_get_type()), marshalAccelFlags},
		{glib.Type(C.gtk_accel_group_get_type()), marshalAccelGroup},
		{glib.Type(C.gtk_accel_map_get_type()), marshalAccelMap},
		{glib.Type(C.gtk_arrow_placement_get_type()), marshalArrowPlacement},
		{glib.Type(C.gtk_arrow_type_get_type()), marshalArrowType},
		{glib.Type(C.gtk_assistant_page_type_get_type()), marshalAssistantPageType},
		{glib.Type(C.gtk_button_box_style_get_type()), marshalButtonBoxStyle},
		{glib.Type(C.gtk_buttons_type_get_type()), marshalButtonsType},
		{glib.Type(C.gtk_calendar_display_options_get_type()), marshalCalendarDisplayOptions},
		{glib.Type(C.gtk_dest_defaults_get_type()), marshalDestDefaults},
		{glib.Type(C.gtk_dialog_flags_get_type()), marshalDialogFlags},
		{glib.Type(C.gtk_entry_icon_position_get_type()), marshalEntryIconPosition},
		{glib.Type(C.gtk_file_chooser_action_get_type()), marshalFileChooserAction},
		{glib.Type(C.gtk_icon_lookup_flags_get_type()), marshalSortType},
		{glib.Type(C.gtk_icon_size_get_type()), marshalIconSize},
		{glib.Type(C.gtk_image_type_get_type()), marshalImageType},
		{glib.Type(C.gtk_input_hints_get_type()), marshalInputHints},
		{glib.Type(C.gtk_input_purpose_get_type()), marshalInputPurpose},
		{glib.Type(C.gtk_justification_get_type()), marshalJustification},
		{glib.Type(C.gtk_license_get_type()), marshalLicense},
		{glib.Type(C.gtk_message_type_get_type()), marshalMessageType},
		{glib.Type(C.gtk_orientation_get_type()), marshalOrientation},
		{glib.Type(C.gtk_pack_type_get_type()), marshalPackType},
		{glib.Type(C.gtk_path_type_get_type()), marshalPathType},
		{glib.Type(C.gtk_policy_type_get_type()), marshalPolicyType},
		{glib.Type(C.gtk_position_type_get_type()), marshalPositionType},
		{glib.Type(C.gtk_relief_style_get_type()), marshalReliefStyle},
		{glib.Type(C.gtk_response_type_get_type()), marshalResponseType},
		{glib.Type(C.gtk_selection_mode_get_type()), marshalSelectionMode},
		{glib.Type(C.gtk_shadow_type_get_type()), marshalShadowType},
		{glib.Type(C.gtk_sort_type_get_type()), marshalSortType},
		{glib.Type(C.gtk_state_flags_get_type()), marshalStateFlags},
		{glib.Type(C.gtk_target_flags_get_type()), marshalTargetFlags},
		{glib.Type(C.gtk_toolbar_style_get_type()), marshalToolbarStyle},
		{glib.Type(C.gtk_tree_model_flags_get_type()), marshalTreeModelFlags},
		{glib.Type(C.gtk_window_position_get_type()), marshalWindowPosition},
		{glib.Type(C.gtk_window_type_get_type()), marshalWindowType},
		{glib.Type(C.gtk_wrap_mode_get_type()), marshalWrapMode},

		// Objects/Interfaces
		{glib.Type(C.gtk_accel_group_get_type()), marshalAccelGroup},
		{glib.Type(C.gtk_accel_map_get_type()), marshalAccelMap},
		{glib.Type(C.gtk_adjustment_get_type()), marshalAdjustment},
		{glib.Type(C.gtk_application_get_type()), marshalApplication},
		{glib.Type(C.gtk_application_window_get_type()), marshalApplicationWindow},
		{glib.Type(C.gtk_assistant_get_type()), marshalAssistant},
		{glib.Type(C.gtk_bin_get_type()), marshalBin},
		{glib.Type(C.gtk_builder_get_type()), marshalBuilder},
		{glib.Type(C.gtk_button_get_type()), marshalButton},
		{glib.Type(C.gtk_box_get_type()), marshalBox},
		{glib.Type(C.gtk_calendar_get_type()), marshalCalendar},
		{glib.Type(C.gtk_cell_layout_get_type()), marshalCellLayout},
		{glib.Type(C.gtk_cell_renderer_get_type()), marshalCellRenderer},
		{glib.Type(C.gtk_cell_renderer_spinner_get_type()), marshalCellRendererSpinner},
		{glib.Type(C.gtk_cell_renderer_pixbuf_get_type()), marshalCellRendererPixbuf},
		{glib.Type(C.gtk_cell_renderer_text_get_type()), marshalCellRendererText},
		{glib.Type(C.gtk_cell_renderer_toggle_get_type()), marshalCellRendererToggle},
		{glib.Type(C.gtk_check_button_get_type()), marshalCheckButton},
		{glib.Type(C.gtk_check_menu_item_get_type()), marshalCheckMenuItem},
		{glib.Type(C.gtk_clipboard_get_type()), marshalClipboard},
		{glib.Type(C.gtk_container_get_type()), marshalContainer},
		{glib.Type(C.gtk_dialog_get_type()), marshalDialog},
		{glib.Type(C.gtk_drawing_area_get_type()), marshalDrawingArea},
		{glib.Type(C.gtk_editable_get_type()), marshalEditable},
		{glib.Type(C.gtk_entry_get_type()), marshalEntry},
		{glib.Type(C.gtk_entry_buffer_get_type()), marshalEntryBuffer},
		{glib.Type(C.gtk_entry_completion_get_type()), marshalEntryCompletion},
		{glib.Type(C.gtk_event_box_get_type()), marshalEventBox},
		{glib.Type(C.gtk_expander_get_type()), marshalExpander},
		{glib.Type(C.gtk_file_chooser_get_type()), marshalFileChooser},
		{glib.Type(C.gtk_file_chooser_button_get_type()), marshalFileChooserButton},
		{glib.Type(C.gtk_file_chooser_dialog_get_type()), marshalFileChooserDialog},
		{glib.Type(C.gtk_file_chooser_widget_get_type()), marshalFileChooserWidget},
		{glib.Type(C.gtk_font_button_get_type()), marshalFontButton},
		{glib.Type(C.gtk_frame_get_type()), marshalFrame},
		{glib.Type(C.gtk_grid_get_type()), marshalGrid},
		{glib.Type(C.gtk_icon_view_get_type()), marshalIconView},
		{glib.Type(C.gtk_image_get_type()), marshalImage},
		{glib.Type(C.gtk_label_get_type()), marshalLabel},
		{glib.Type(C.gtk_link_button_get_type()), marshalLinkButton},
		{glib.Type(C.gtk_layout_get_type()), marshalLayout},
		{glib.Type(C.gtk_list_store_get_type()), marshalListStore},
		{glib.Type(C.gtk_menu_get_type()), marshalMenu},
		{glib.Type(C.gtk_menu_bar_get_type()), marshalMenuBar},
		{glib.Type(C.gtk_menu_button_get_type()), marshalMenuButton},
		{glib.Type(C.gtk_menu_item_get_type()), marshalMenuItem},
		{glib.Type(C.gtk_menu_shell_get_type()), marshalMenuShell},
		{glib.Type(C.gtk_message_dialog_get_type()), marshalMessageDialog},
		{glib.Type(C.gtk_notebook_get_type()), marshalNotebook},
		{glib.Type(C.gtk_offscreen_window_get_type()), marshalOffscreenWindow},
		{glib.Type(C.gtk_orientable_get_type()), marshalOrientable},
		{glib.Type(C.gtk_overlay_get_type()), marshalOverlay},
		{glib.Type(C.gtk_paned_get_type()), marshalPaned},
		{glib.Type(C.gtk_progress_bar_get_type()), marshalProgressBar},
		{glib.Type(C.gtk_radio_button_get_type()), marshalRadioButton},
		{glib.Type(C.gtk_radio_menu_item_get_type()), marshalRadioMenuItem},
		{glib.Type(C.gtk_range_get_type()), marshalRange},
		{glib.Type(C.gtk_recent_chooser_get_type()), marshalRecentChooser},
		{glib.Type(C.gtk_scale_button_get_type()), marshalScaleButton},
		{glib.Type(C.gtk_scale_get_type()), marshalScale},
		{glib.Type(C.gtk_scrollbar_get_type()), marshalScrollbar},
		{glib.Type(C.gtk_scrolled_window_get_type()), marshalScrolledWindow},
		{glib.Type(C.gtk_search_entry_get_type()), marshalSearchEntry},
		{glib.Type(C.gtk_selection_data_get_type()), marshalSelectionData},
		{glib.Type(C.gtk_separator_get_type()), marshalSeparator},
		{glib.Type(C.gtk_separator_menu_item_get_type()), marshalSeparatorMenuItem},
		{glib.Type(C.gtk_separator_tool_item_get_type()), marshalSeparatorToolItem},
		{glib.Type(C.gtk_spin_button_get_type()), marshalSpinButton},
		{glib.Type(C.gtk_spinner_get_type()), marshalSpinner},
		{glib.Type(C.gtk_statusbar_get_type()), marshalStatusbar},
		{glib.Type(C.gtk_switch_get_type()), marshalSwitch},
		{glib.Type(C.gtk_text_view_get_type()), marshalTextView},
		{glib.Type(C.gtk_text_tag_get_type()), marshalTextTag},
		{glib.Type(C.gtk_text_tag_table_get_type()), marshalTextTagTable},
		{glib.Type(C.gtk_text_buffer_get_type()), marshalTextBuffer},
		{glib.Type(C.gtk_toggle_button_get_type()), marshalToggleButton},
		{glib.Type(C.gtk_toolbar_get_type()), marshalToolbar},
		{glib.Type(C.gtk_tool_button_get_type()), marshalToolButton},
		{glib.Type(C.gtk_tool_item_get_type()), marshalToolItem},
		{glib.Type(C.gtk_tree_model_get_type()), marshalTreeModel},
		{glib.Type(C.gtk_tree_selection_get_type()), marshalTreeSelection},
		{glib.Type(C.gtk_tree_store_get_type()), marshalTreeStore},
		{glib.Type(C.gtk_tree_view_get_type()), marshalTreeView},
		{glib.Type(C.gtk_tree_view_column_get_type()), marshalTreeViewColumn},
		{glib.Type(C.gtk_volume_button_get_type()), marshalVolumeButton},
		{glib.Type(C.gtk_widget_get_type()), marshalWidget},
		{glib.Type(C.gtk_window_get_type()), marshalWindow},

		// Boxed
		{glib.Type(C.gtk_target_entry_get_type()), marshalTargetEntry},
		{glib.Type(C.gtk_text_iter_get_type()), marshalTextIter},
		{glib.Type(C.gtk_text_mark_get_type()), marshalTextMark},
		{glib.Type(C.gtk_tree_iter_get_type()), marshalTreeIter},
		{glib.Type(C.gtk_tree_path_get_type()), marshalTreePath},
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
	return b != C.FALSE
}

func cGSList(clist *glib.SList) *C.GSList {
	if clist == nil {
		return nil
	}
	return (*C.GSList)(unsafe.Pointer(clist.Native()))
}

func free(str ...interface{}) {
	for _, s := range str {
		switch x := s.(type) {
		case *C.char:
			C.free(unsafe.Pointer(x))
		case []*C.char:
			for _, cp := range x {
				C.free(unsafe.Pointer(cp))
			}
			/*
				case C.gpointer:
					C.g_free(C.gpointer(c))
			*/
		default:
			fmt.Printf("utils.go free(): Unknown type: %T\n", x)
		}

	}
}

func goString(cstr *C.gchar) string {
	return C.GoString((*C.char)(cstr))
}

func goStringArray(c **C.gchar) []string {
	var strs []string

	for *c != nil {
		strs = append(strs, goString(*c))
		c = C.next_gcharptr(c)
	}

	return strs
}

// Wrapper function for TestBoolConvs since cgo can't be used with
// testing package
func testBoolConvs() error {
	b := gobool(gbool(true))
	if b != true {
		return errors.New("Unexpected bool conversion result")
	}

	cb := gbool(gobool(C.gboolean(0)))
	if cb != C.gboolean(0) {
		return errors.New("Unexpected bool conversion result")
	}

	return nil
}

/*
 * Unexported vars
 */

var nilPtrErr = errors.New("cgo returned unexpected nil pointer")

/*
 * Constants
 */

// Align is a representation of GTK's GtkAlign.
type Align int

const (
	ALIGN_FILL   Align = C.GTK_ALIGN_FILL
	ALIGN_START  Align = C.GTK_ALIGN_START
	ALIGN_END    Align = C.GTK_ALIGN_END
	ALIGN_CENTER Align = C.GTK_ALIGN_CENTER
)

func marshalAlign(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum(C.toGValue(unsafe.Pointer(p)))
	return Align(c), nil
}

// ArrowPlacement is a representation of GTK's GtkArrowPlacement.
type ArrowPlacement int

const (
	ARROWS_BOTH  ArrowPlacement = C.GTK_ARROWS_BOTH
	ARROWS_START ArrowPlacement = C.GTK_ARROWS_START
	ARROWS_END   ArrowPlacement = C.GTK_ARROWS_END
)

func marshalArrowPlacement(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum(C.toGValue(unsafe.Pointer(p)))
	return ArrowPlacement(c), nil
}

// ArrowType is a representation of GTK's GtkArrowType.
type ArrowType int

const (
	ARROW_UP    ArrowType = C.GTK_ARROW_UP
	ARROW_DOWN  ArrowType = C.GTK_ARROW_DOWN
	ARROW_LEFT  ArrowType = C.GTK_ARROW_LEFT
	ARROW_RIGHT ArrowType = C.GTK_ARROW_RIGHT
	ARROW_NONE  ArrowType = C.GTK_ARROW_NONE
)

func marshalArrowType(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum(C.toGValue(unsafe.Pointer(p)))
	return ArrowType(c), nil
}

// AssistantPageType is a representation of GTK's GtkAssistantPageType.
type AssistantPageType int

const (
	ASSISTANT_PAGE_CONTENT  AssistantPageType = C.GTK_ASSISTANT_PAGE_CONTENT
	ASSISTANT_PAGE_INTRO    AssistantPageType = C.GTK_ASSISTANT_PAGE_INTRO
	ASSISTANT_PAGE_CONFIRM  AssistantPageType = C.GTK_ASSISTANT_PAGE_CONFIRM
	ASSISTANT_PAGE_SUMMARY  AssistantPageType = C.GTK_ASSISTANT_PAGE_SUMMARY
	ASSISTANT_PAGE_PROGRESS AssistantPageType = C.GTK_ASSISTANT_PAGE_PROGRESS
	ASSISTANT_PAGE_CUSTOM   AssistantPageType = C.GTK_ASSISTANT_PAGE_CUSTOM
)

func marshalAssistantPageType(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum(C.toGValue(unsafe.Pointer(p)))
	return AssistantPageType(c), nil
}

// ButtonsType is a representation of GTK's GtkButtonsType.
type ButtonsType int

const (
	BUTTONS_NONE      ButtonsType = C.GTK_BUTTONS_NONE
	BUTTONS_OK        ButtonsType = C.GTK_BUTTONS_OK
	BUTTONS_CLOSE     ButtonsType = C.GTK_BUTTONS_CLOSE
	BUTTONS_CANCEL    ButtonsType = C.GTK_BUTTONS_CANCEL
	BUTTONS_YES_NO    ButtonsType = C.GTK_BUTTONS_YES_NO
	BUTTONS_OK_CANCEL ButtonsType = C.GTK_BUTTONS_OK_CANCEL
)

func marshalButtonsType(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum(C.toGValue(unsafe.Pointer(p)))
	return ButtonsType(c), nil
}

// CalendarDisplayOptions is a representation of GTK's GtkCalendarDisplayOptions
type CalendarDisplayOptions int

const (
	CALENDAR_SHOW_HEADING      CalendarDisplayOptions = C.GTK_CALENDAR_SHOW_HEADING
	CALENDAR_SHOW_DAY_NAMES    CalendarDisplayOptions = C.GTK_CALENDAR_SHOW_DAY_NAMES
	CALENDAR_NO_MONTH_CHANGE   CalendarDisplayOptions = C.GTK_CALENDAR_NO_MONTH_CHANGE
	CALENDAR_SHOW_WEEK_NUMBERS CalendarDisplayOptions = C.GTK_CALENDAR_SHOW_WEEK_NUMBERS
	CALENDAR_SHOW_DETAILS      CalendarDisplayOptions = C.GTK_CALENDAR_SHOW_DETAILS
)

func marshalCalendarDisplayOptions(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum(C.toGValue(unsafe.Pointer(p)))
	return CalendarDisplayOptions(c), nil
}

// DestDefaults is a representation of GTK's GtkDestDefaults.
type DestDefaults int

const (
	DEST_DEFAULT_MOTION    DestDefaults = C.GTK_DEST_DEFAULT_MOTION
	DEST_DEFAULT_HIGHLIGHT DestDefaults = C.GTK_DEST_DEFAULT_HIGHLIGHT
	DEST_DEFAULT_DROP      DestDefaults = C.GTK_DEST_DEFAULT_DROP
	DEST_DEFAULT_ALL       DestDefaults = C.GTK_DEST_DEFAULT_ALL
)

func marshalDestDefaults(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum(C.toGValue(unsafe.Pointer(p)))
	return DestDefaults(c), nil
}

// DialogFlags is a representation of GTK's GtkDialogFlags.
type DialogFlags int

const (
	DIALOG_MODAL               DialogFlags = C.GTK_DIALOG_MODAL
	DIALOG_DESTROY_WITH_PARENT DialogFlags = C.GTK_DIALOG_DESTROY_WITH_PARENT
)

func marshalDialogFlags(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum(C.toGValue(unsafe.Pointer(p)))
	return DialogFlags(c), nil
}

// EntryIconPosition is a representation of GTK's GtkEntryIconPosition.
type EntryIconPosition int

const (
	ENTRY_ICON_PRIMARY   EntryIconPosition = C.GTK_ENTRY_ICON_PRIMARY
	ENTRY_ICON_SECONDARY EntryIconPosition = C.GTK_ENTRY_ICON_SECONDARY
)

func marshalEntryIconPosition(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum(C.toGValue(unsafe.Pointer(p)))
	return EntryIconPosition(c), nil
}

// FileChooserAction is a representation of GTK's GtkFileChooserAction.
type FileChooserAction int

const (
	FILE_CHOOSER_ACTION_OPEN          FileChooserAction = C.GTK_FILE_CHOOSER_ACTION_OPEN
	FILE_CHOOSER_ACTION_SAVE          FileChooserAction = C.GTK_FILE_CHOOSER_ACTION_SAVE
	FILE_CHOOSER_ACTION_SELECT_FOLDER FileChooserAction = C.GTK_FILE_CHOOSER_ACTION_SELECT_FOLDER
	FILE_CHOOSER_ACTION_CREATE_FOLDER FileChooserAction = C.GTK_FILE_CHOOSER_ACTION_CREATE_FOLDER
)

func marshalFileChooserAction(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum(C.toGValue(unsafe.Pointer(p)))
	return FileChooserAction(c), nil
}

// IconLookupFlags is a representation of GTK's GtkIconLookupFlags.
type IconLookupFlags int

const (
	ICON_LOOKUP_NO_SVG           IconLookupFlags = C.GTK_ICON_LOOKUP_NO_SVG
	ICON_LOOKUP_FORCE_SVG                        = C.GTK_ICON_LOOKUP_FORCE_SVG
	ICON_LOOKUP_USE_BUILTIN                      = C.GTK_ICON_LOOKUP_USE_BUILTIN
	ICON_LOOKUP_GENERIC_FALLBACK                 = C.GTK_ICON_LOOKUP_GENERIC_FALLBACK
	ICON_LOOKUP_FORCE_SIZE                       = C.GTK_ICON_LOOKUP_FORCE_SIZE
)

func marshalIconLookupFlags(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum(C.toGValue(unsafe.Pointer(p)))
	return IconLookupFlags(c), nil
}

// IconSize is a representation of GTK's GtkIconSize.
type IconSize int

const (
	ICON_SIZE_INVALID       IconSize = C.GTK_ICON_SIZE_INVALID
	ICON_SIZE_MENU          IconSize = C.GTK_ICON_SIZE_MENU
	ICON_SIZE_SMALL_TOOLBAR IconSize = C.GTK_ICON_SIZE_SMALL_TOOLBAR
	ICON_SIZE_LARGE_TOOLBAR IconSize = C.GTK_ICON_SIZE_LARGE_TOOLBAR
	ICON_SIZE_BUTTON        IconSize = C.GTK_ICON_SIZE_BUTTON
	ICON_SIZE_DND           IconSize = C.GTK_ICON_SIZE_DND
	ICON_SIZE_DIALOG        IconSize = C.GTK_ICON_SIZE_DIALOG
)

func marshalIconSize(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum(C.toGValue(unsafe.Pointer(p)))
	return IconSize(c), nil
}

// ImageType is a representation of GTK's GtkImageType.
type ImageType int

const (
	IMAGE_EMPTY     ImageType = C.GTK_IMAGE_EMPTY
	IMAGE_PIXBUF    ImageType = C.GTK_IMAGE_PIXBUF
	IMAGE_STOCK     ImageType = C.GTK_IMAGE_STOCK
	IMAGE_ICON_SET  ImageType = C.GTK_IMAGE_ICON_SET
	IMAGE_ANIMATION ImageType = C.GTK_IMAGE_ANIMATION
	IMAGE_ICON_NAME ImageType = C.GTK_IMAGE_ICON_NAME
	IMAGE_GICON     ImageType = C.GTK_IMAGE_GICON
)

func marshalImageType(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum(C.toGValue(unsafe.Pointer(p)))
	return ImageType(c), nil
}

// InputHints is a representation of GTK's GtkInputHints.
type InputHints int

const (
	INPUT_HINT_NONE                InputHints = C.GTK_INPUT_HINT_NONE
	INPUT_HINT_SPELLCHECK          InputHints = C.GTK_INPUT_HINT_SPELLCHECK
	INPUT_HINT_NO_SPELLCHECK       InputHints = C.GTK_INPUT_HINT_NO_SPELLCHECK
	INPUT_HINT_WORD_COMPLETION     InputHints = C.GTK_INPUT_HINT_WORD_COMPLETION
	INPUT_HINT_LOWERCASE           InputHints = C.GTK_INPUT_HINT_LOWERCASE
	INPUT_HINT_UPPERCASE_CHARS     InputHints = C.GTK_INPUT_HINT_UPPERCASE_CHARS
	INPUT_HINT_UPPERCASE_WORDS     InputHints = C.GTK_INPUT_HINT_UPPERCASE_WORDS
	INPUT_HINT_UPPERCASE_SENTENCES InputHints = C.GTK_INPUT_HINT_UPPERCASE_SENTENCES
	INPUT_HINT_INHIBIT_OSK         InputHints = C.GTK_INPUT_HINT_INHIBIT_OSK
)

func marshalInputHints(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum(C.toGValue(unsafe.Pointer(p)))
	return InputHints(c), nil
}

// InputPurpose is a representation of GTK's GtkInputPurpose.
type InputPurpose int

const (
	INPUT_PURPOSE_FREE_FORM InputPurpose = C.GTK_INPUT_PURPOSE_FREE_FORM
	INPUT_PURPOSE_ALPHA     InputPurpose = C.GTK_INPUT_PURPOSE_ALPHA
	INPUT_PURPOSE_DIGITS    InputPurpose = C.GTK_INPUT_PURPOSE_DIGITS
	INPUT_PURPOSE_NUMBER    InputPurpose = C.GTK_INPUT_PURPOSE_NUMBER
	INPUT_PURPOSE_PHONE     InputPurpose = C.GTK_INPUT_PURPOSE_PHONE
	INPUT_PURPOSE_URL       InputPurpose = C.GTK_INPUT_PURPOSE_URL
	INPUT_PURPOSE_EMAIL     InputPurpose = C.GTK_INPUT_PURPOSE_EMAIL
	INPUT_PURPOSE_NAME      InputPurpose = C.GTK_INPUT_PURPOSE_NAME
	INPUT_PURPOSE_PASSWORD  InputPurpose = C.GTK_INPUT_PURPOSE_PASSWORD
	INPUT_PURPOSE_PIN       InputPurpose = C.GTK_INPUT_PURPOSE_PIN
)

func marshalInputPurpose(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum(C.toGValue(unsafe.Pointer(p)))
	return InputPurpose(c), nil
}

// Justify is a representation of GTK's GtkJustification.
type Justification int

const (
	JUSTIFY_LEFT   Justification = C.GTK_JUSTIFY_LEFT
	JUSTIFY_RIGHT  Justification = C.GTK_JUSTIFY_RIGHT
	JUSTIFY_CENTER Justification = C.GTK_JUSTIFY_CENTER
	JUSTIFY_FILL   Justification = C.GTK_JUSTIFY_FILL
)

func marshalJustification(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum(C.toGValue(unsafe.Pointer(p)))
	return Justification(c), nil
}

// License is a representation of GTK's GtkLicense.
type License int

const (
	LICENSE_UNKNOWN      License = C.GTK_LICENSE_UNKNOWN
	LICENSE_CUSTOM       License = C.GTK_LICENSE_CUSTOM
	LICENSE_GPL_2_0      License = C.GTK_LICENSE_GPL_2_0
	LICENSE_GPL_3_0      License = C.GTK_LICENSE_GPL_3_0
	LICENSE_LGPL_2_1     License = C.GTK_LICENSE_LGPL_2_1
	LICENSE_LGPL_3_0     License = C.GTK_LICENSE_LGPL_3_0
	LICENSE_BSD          License = C.GTK_LICENSE_BSD
	LICENSE_MIT_X11      License = C.GTK_LICENSE_MIT_X11
	LICENSE_GTK_ARTISTIC License = C.GTK_LICENSE_ARTISTIC
)

func marshalLicense(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum(C.toGValue(unsafe.Pointer(p)))
	return License(c), nil
}

// MessageType is a representation of GTK's GtkMessageType.
type MessageType int

const (
	MESSAGE_INFO     MessageType = C.GTK_MESSAGE_INFO
	MESSAGE_WARNING  MessageType = C.GTK_MESSAGE_WARNING
	MESSAGE_QUESTION MessageType = C.GTK_MESSAGE_QUESTION
	MESSAGE_ERROR    MessageType = C.GTK_MESSAGE_ERROR
	MESSAGE_OTHER    MessageType = C.GTK_MESSAGE_OTHER
)

func marshalMessageType(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum(C.toGValue(unsafe.Pointer(p)))
	return MessageType(c), nil
}

// Orientation is a representation of GTK's GtkOrientation.
type Orientation int

const (
	ORIENTATION_HORIZONTAL Orientation = C.GTK_ORIENTATION_HORIZONTAL
	ORIENTATION_VERTICAL   Orientation = C.GTK_ORIENTATION_VERTICAL
)

func marshalOrientation(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum(C.toGValue(unsafe.Pointer(p)))
	return Orientation(c), nil
}

// PackType is a representation of GTK's GtkPackType.
type PackType int

const (
	PACK_START PackType = C.GTK_PACK_START
	PACK_END   PackType = C.GTK_PACK_END
)

func marshalPackType(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum(C.toGValue(unsafe.Pointer(p)))
	return PackType(c), nil
}

// PathType is a representation of GTK's GtkPathType.
type PathType int

const (
	PATH_WIDGET       PathType = C.GTK_PATH_WIDGET
	PATH_WIDGET_CLASS PathType = C.GTK_PATH_WIDGET_CLASS
	PATH_CLASS        PathType = C.GTK_PATH_CLASS
)

func marshalPathType(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum(C.toGValue(unsafe.Pointer(p)))
	return PathType(c), nil
}

// PolicyType is a representation of GTK's GtkPolicyType.
type PolicyType int

const (
	POLICY_ALWAYS    PolicyType = C.GTK_POLICY_ALWAYS
	POLICY_AUTOMATIC PolicyType = C.GTK_POLICY_AUTOMATIC
	POLICY_NEVER     PolicyType = C.GTK_POLICY_NEVER
)

func marshalPolicyType(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum(C.toGValue(unsafe.Pointer(p)))
	return PolicyType(c), nil
}

// PositionType is a representation of GTK's GtkPositionType.
type PositionType int

const (
	POS_LEFT   PositionType = C.GTK_POS_LEFT
	POS_RIGHT  PositionType = C.GTK_POS_RIGHT
	POS_TOP    PositionType = C.GTK_POS_TOP
	POS_BOTTOM PositionType = C.GTK_POS_BOTTOM
)

func marshalPositionType(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum(C.toGValue(unsafe.Pointer(p)))
	return PositionType(c), nil
}

// ReliefStyle is a representation of GTK's GtkReliefStyle.
type ReliefStyle int

const (
	RELIEF_NORMAL ReliefStyle = C.GTK_RELIEF_NORMAL
	RELIEF_HALF   ReliefStyle = C.GTK_RELIEF_HALF
	RELIEF_NONE   ReliefStyle = C.GTK_RELIEF_NONE
)

func marshalReliefStyle(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum(C.toGValue(unsafe.Pointer(p)))
	return ReliefStyle(c), nil
}

// ResponseType is a representation of GTK's GtkResponseType.
type ResponseType int

const (
	RESPONSE_NONE         ResponseType = C.GTK_RESPONSE_NONE
	RESPONSE_REJECT       ResponseType = C.GTK_RESPONSE_REJECT
	RESPONSE_ACCEPT       ResponseType = C.GTK_RESPONSE_ACCEPT
	RESPONSE_DELETE_EVENT ResponseType = C.GTK_RESPONSE_DELETE_EVENT
	RESPONSE_OK           ResponseType = C.GTK_RESPONSE_OK
	RESPONSE_CANCEL       ResponseType = C.GTK_RESPONSE_CANCEL
	RESPONSE_CLOSE        ResponseType = C.GTK_RESPONSE_CLOSE
	RESPONSE_YES          ResponseType = C.GTK_RESPONSE_YES
	RESPONSE_NO           ResponseType = C.GTK_RESPONSE_NO
	RESPONSE_APPLY        ResponseType = C.GTK_RESPONSE_APPLY
	RESPONSE_HELP         ResponseType = C.GTK_RESPONSE_HELP
)

func marshalResponseType(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum(C.toGValue(unsafe.Pointer(p)))
	return ResponseType(c), nil
}

// SelectionMode is a representation of GTK's GtkSelectionMode.
type SelectionMode int

const (
	SELECTION_NONE     SelectionMode = C.GTK_SELECTION_NONE
	SELECTION_SINGLE   SelectionMode = C.GTK_SELECTION_SINGLE
	SELECTION_BROWSE   SelectionMode = C.GTK_SELECTION_BROWSE
	SELECTION_MULTIPLE SelectionMode = C.GTK_SELECTION_MULTIPLE
)

func marshalSelectionMode(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum(C.toGValue(unsafe.Pointer(p)))
	return SelectionMode(c), nil
}

// ShadowType is a representation of GTK's GtkShadowType.
type ShadowType int

const (
	SHADOW_NONE       ShadowType = C.GTK_SHADOW_NONE
	SHADOW_IN         ShadowType = C.GTK_SHADOW_IN
	SHADOW_OUT        ShadowType = C.GTK_SHADOW_OUT
	SHADOW_ETCHED_IN  ShadowType = C.GTK_SHADOW_ETCHED_IN
	SHADOW_ETCHED_OUT ShadowType = C.GTK_SHADOW_ETCHED_OUT
)

func marshalShadowType(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum(C.toGValue(unsafe.Pointer(p)))
	return ShadowType(c), nil
}

// SizeGroupMode is a representation of GTK's GtkSizeGroupMode
type SizeGroupMode int

const (
	SIZE_GROUP_NONE       SizeGroupMode = C.GTK_SIZE_GROUP_NONE
	SIZE_GROUP_HORIZONTAL SizeGroupMode = C.GTK_SIZE_GROUP_HORIZONTAL
	SIZE_GROUP_VERTICAL   SizeGroupMode = C.GTK_SIZE_GROUP_VERTICAL
	SIZE_GROUP_BOTH       SizeGroupMode = C.GTK_SIZE_GROUP_BOTH
)

func marshalSizeGroupMode(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum(C.toGValue(unsafe.Pointer(p)))
	return SizeGroupMode(c), nil
}

// SortType is a representation of GTK's GtkSortType.
type SortType int

const (
	SORT_ASCENDING  SortType = C.GTK_SORT_ASCENDING
	SORT_DESCENDING          = C.GTK_SORT_DESCENDING
)

func marshalSortType(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum(C.toGValue(unsafe.Pointer(p)))
	return SortType(c), nil
}

// StateFlags is a representation of GTK's GtkStateFlags.
type StateFlags int

const (
	STATE_FLAG_NORMAL       StateFlags = C.GTK_STATE_FLAG_NORMAL
	STATE_FLAG_ACTIVE       StateFlags = C.GTK_STATE_FLAG_ACTIVE
	STATE_FLAG_PRELIGHT     StateFlags = C.GTK_STATE_FLAG_PRELIGHT
	STATE_FLAG_SELECTED     StateFlags = C.GTK_STATE_FLAG_SELECTED
	STATE_FLAG_INSENSITIVE  StateFlags = C.GTK_STATE_FLAG_INSENSITIVE
	STATE_FLAG_INCONSISTENT StateFlags = C.GTK_STATE_FLAG_INCONSISTENT
	STATE_FLAG_FOCUSED      StateFlags = C.GTK_STATE_FLAG_FOCUSED
	STATE_FLAG_BACKDROP     StateFlags = C.GTK_STATE_FLAG_BACKDROP
)

func marshalStateFlags(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum(C.toGValue(unsafe.Pointer(p)))
	return StateFlags(c), nil
}

// TargetFlags is a representation of GTK's GtkTargetFlags.
type TargetFlags int

const (
	TARGET_SAME_APP     TargetFlags = C.GTK_TARGET_SAME_APP
	TARGET_SAME_WIDGET  TargetFlags = C.GTK_TARGET_SAME_WIDGET
	TARGET_OTHER_APP    TargetFlags = C.GTK_TARGET_OTHER_APP
	TARGET_OTHER_WIDGET TargetFlags = C.GTK_TARGET_OTHER_WIDGET
)

func marshalTargetFlags(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum(C.toGValue(unsafe.Pointer(p)))
	return TargetFlags(c), nil
}

// ToolbarStyle is a representation of GTK's GtkToolbarStyle.
type ToolbarStyle int

const (
	TOOLBAR_ICONS      ToolbarStyle = C.GTK_TOOLBAR_ICONS
	TOOLBAR_TEXT       ToolbarStyle = C.GTK_TOOLBAR_TEXT
	TOOLBAR_BOTH       ToolbarStyle = C.GTK_TOOLBAR_BOTH
	TOOLBAR_BOTH_HORIZ ToolbarStyle = C.GTK_TOOLBAR_BOTH_HORIZ
)

func marshalToolbarStyle(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum(C.toGValue(unsafe.Pointer(p)))
	return ToolbarStyle(c), nil
}

// TreeModelFlags is a representation of GTK's GtkTreeModelFlags.
type TreeModelFlags int

const (
	TREE_MODEL_ITERS_PERSIST TreeModelFlags = C.GTK_TREE_MODEL_ITERS_PERSIST
	TREE_MODEL_LIST_ONLY     TreeModelFlags = C.GTK_TREE_MODEL_LIST_ONLY
)

func marshalTreeModelFlags(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum(C.toGValue(unsafe.Pointer(p)))
	return TreeModelFlags(c), nil
}

// WindowPosition is a representation of GTK's GtkWindowPosition.
type WindowPosition int

const (
	WIN_POS_NONE             WindowPosition = C.GTK_WIN_POS_NONE
	WIN_POS_CENTER           WindowPosition = C.GTK_WIN_POS_CENTER
	WIN_POS_MOUSE            WindowPosition = C.GTK_WIN_POS_MOUSE
	WIN_POS_CENTER_ALWAYS    WindowPosition = C.GTK_WIN_POS_CENTER_ALWAYS
	WIN_POS_CENTER_ON_PARENT WindowPosition = C.GTK_WIN_POS_CENTER_ON_PARENT
)

func marshalWindowPosition(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum(C.toGValue(unsafe.Pointer(p)))
	return WindowPosition(c), nil
}

// WindowType is a representation of GTK's GtkWindowType.
type WindowType int

const (
	WINDOW_TOPLEVEL WindowType = C.GTK_WINDOW_TOPLEVEL
	WINDOW_POPUP    WindowType = C.GTK_WINDOW_POPUP
)

func marshalWindowType(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum(C.toGValue(unsafe.Pointer(p)))
	return WindowType(c), nil
}

// WrapMode is a representation of GTK's GtkWrapMode.
type WrapMode int

const (
	WRAP_NONE      WrapMode = C.GTK_WRAP_NONE
	WRAP_CHAR      WrapMode = C.GTK_WRAP_CHAR
	WRAP_WORD      WrapMode = C.GTK_WRAP_WORD
	WRAP_WORD_CHAR WrapMode = C.GTK_WRAP_WORD_CHAR
)

func marshalWrapMode(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum(C.toGValue(unsafe.Pointer(p)))
	return WrapMode(c), nil
}

/*
 * Init and main event loop
 */

/*
Init() is a wrapper around gtk_init() and must be called before any
other GTK calls and is used to initialize everything necessary.

In addition to setting up GTK for usage, a pointer to a slice of
strings may be passed in to parse standard GTK command line arguments.
args will be modified to remove any flags that were handled.
Alternatively, nil may be passed in to not perform any command line
parsing.
*/
func Init(args *[]string) {
	if args != nil {
		argc := C.int(len(*args))
		argv := C.make_strings(argc)
		defer C.destroy_strings(argv)

		for i, arg := range *args {
			cstr := C.CString(arg)
			C.set_string(argv, C.int(i), (*C.gchar)(cstr))
		}

		C.gtk_init((*C.int)(unsafe.Pointer(&argc)),
			(***C.char)(unsafe.Pointer(&argv)))

		unhandled := make([]string, argc)
		for i := 0; i < int(argc); i++ {
			cstr := C.get_string(argv, C.int(i))
			unhandled[i] = goString(cstr)
			C.free(unsafe.Pointer(cstr))
		}
		*args = unhandled
	} else {
		C.gtk_init(nil, nil)
	}
}

// Main() is a wrapper around gtk_main() and runs the GTK main loop,
// blocking until MainQuit() is called.
func Main() {
	C.gtk_main()
}

// MainIteration is a wrapper around gtk_main_iteration.
func MainIteration() bool {
	return gobool(C.gtk_main_iteration())
}

// MainIterationDo is a wrapper around gtk_main_iteration_do.
func MainIterationDo(blocking bool) bool {
	return gobool(C.gtk_main_iteration_do(gbool(blocking)))
}

// EventsPending is a wrapper around gtk_events_pending.
func EventsPending() bool {
	return gobool(C.gtk_events_pending())
}

// MainQuit() is a wrapper around gtk_main_quit() is used to terminate
// the GTK main loop (started by Main()).
func MainQuit() {
	C.gtk_main_quit()
}

/*
 * GtkAdjustment
 */

// Adjustment is a representation of GTK's GtkAdjustment.
type Adjustment struct {
	glib.InitiallyUnowned
}

// native returns a pointer to the underlying GtkAdjustment.
func (v *Adjustment) native() *C.GtkAdjustment {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toGtkAdjustment(ptr)
}

func marshalAdjustment(p uintptr) (interface{}, error) {
	c := C.g_value_get_object(C.toGValue(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapAdjustment(obj), nil
}

func wrapAdjustment(obj *glib.Object) *Adjustment {
	return &Adjustment{glib.InitiallyUnowned{obj}}
}

// AdjustmentNew is a wrapper around gtk_adjustment_new().
func AdjustmentNew(value, lower, upper, stepIncrement, pageIncrement, pageSize float64) (*Adjustment, error) {
	c := C.gtk_adjustment_new(C.gdouble(value),
		C.gdouble(lower),
		C.gdouble(upper),
		C.gdouble(stepIncrement),
		C.gdouble(pageIncrement),
		C.gdouble(pageSize))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapAdjustment(obj), nil
}

// GetValue is a wrapper around gtk_adjustment_get_value().
func (v *Adjustment) GetValue() float64 {
	c := C.gtk_adjustment_get_value(v.native())
	return float64(c)
}

// SetValue is a wrapper around gtk_adjustment_set_value().
func (v *Adjustment) SetValue(value float64) {
	C.gtk_adjustment_set_value(v.native(), C.gdouble(value))
}

// GetLower is a wrapper around gtk_adjustment_get_lower().
func (v *Adjustment) GetLower() float64 {
	c := C.gtk_adjustment_get_lower(v.native())
	return float64(c)
}

// GetPageSize is a wrapper around gtk_adjustment_get_page_size().
func (v *Adjustment) GetPageSize() float64 {
	return float64(C.gtk_adjustment_get_page_size(v.native()))
}

// SetPageSize is a wrapper around gtk_adjustment_set_page_size().
func (v *Adjustment) SetPageSize(value float64) {
	C.gtk_adjustment_set_page_size(v.native(), C.gdouble(value))
}

// Configure is a wrapper around gtk_adjustment_configure().
func (v *Adjustment) Configure(value, lower, upper, stepIncrement, pageIncrement, pageSize float64) {
	C.gtk_adjustment_configure(v.native(), C.gdouble(value),
		C.gdouble(lower), C.gdouble(upper), C.gdouble(stepIncrement),
		C.gdouble(pageIncrement), C.gdouble(pageSize))
}

// SetLower is a wrapper around gtk_adjustment_set_lower().
func (v *Adjustment) SetLower(value float64) {
	C.gtk_adjustment_set_lower(v.native(), C.gdouble(value))
}

// GetUpper is a wrapper around gtk_adjustment_get_upper().
func (v *Adjustment) GetUpper() float64 {
	c := C.gtk_adjustment_get_upper(v.native())
	return float64(c)
}

// SetUpper is a wrapper around gtk_adjustment_set_upper().
func (v *Adjustment) SetUpper(value float64) {
	C.gtk_adjustment_set_upper(v.native(), C.gdouble(value))
}

// GetPageIncrement is a wrapper around gtk_adjustment_get_page_increment().
func (v *Adjustment) GetPageIncrement() float64 {
	c := C.gtk_adjustment_get_page_increment(v.native())
	return float64(c)
}

// SetPageIncrement is a wrapper around gtk_adjustment_set_page_increment().
func (v *Adjustment) SetPageIncrement(value float64) {
	C.gtk_adjustment_set_page_increment(v.native(), C.gdouble(value))
}

// GetStepIncrement is a wrapper around gtk_adjustment_get_step_increment().
func (v *Adjustment) GetStepIncrement() float64 {
	c := C.gtk_adjustment_get_step_increment(v.native())
	return float64(c)
}

// SetStepIncrement is a wrapper around gtk_adjustment_set_step_increment().
func (v *Adjustment) SetStepIncrement(value float64) {
	C.gtk_adjustment_set_step_increment(v.native(), C.gdouble(value))
}

// GetMinimumIncrement is a wrapper around gtk_adjustment_get_minimum_increment().
func (v *Adjustment) GetMinimumIncrement() float64 {
	c := C.gtk_adjustment_get_minimum_increment(v.native())
	return float64(c)
}

/*
void	gtk_adjustment_clamp_page ()
void	gtk_adjustment_changed ()
void	gtk_adjustment_value_changed ()
void	gtk_adjustment_configure ()
*/

/*
 * GtkBuilder
 */

// Builder is a representation of GTK's GtkBuilder.
type Builder struct {
	*glib.Object
}

// native() returns a pointer to the underlying GtkBuilder.
func (v *Builder) native() *C.GtkBuilder {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toGtkBuilder(ptr)
}

func marshalBuilder(p uintptr) (interface{}, error) {
	c := C.g_value_get_object(C.toGValue(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return &Builder{obj}, nil
}

func wrapBuilder(obj *glib.Object) *Builder {
	return &Builder{obj}
}

// BuilderNew is a wrapper around gtk_builder_new().
func BuilderNew() (*Builder, error) {
	c := C.gtk_builder_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return &Builder{obj}, nil
}

// BuilderNewFromResource is a wrapper around gtk_builder_new_from_resource().
func BuilderNewFromResource(resourcePath string) (*Builder, error) {
	cstr := C.CString(resourcePath)
	defer C.free(unsafe.Pointer(cstr))

	c := C.gtk_builder_new_from_resource((*C.gchar)(cstr))
	if c == nil {
		return nil, nilPtrErr
	}

	obj := glib.Take(unsafe.Pointer(c))
	return wrapBuilder(obj), nil
}

// AddFromFile is a wrapper around gtk_builder_add_from_file().
func (v *Builder) AddFromFile(filename string) error {
	cstr := C.CString(filename)
	defer C.free(unsafe.Pointer(cstr))
	var err *C.GError = nil
	res := C.gtk_builder_add_from_file(v.native(), (*C.gchar)(cstr), &err)
	if res == 0 {
		defer C.g_error_free(err)
		return errors.New(goString(err.message))
	}
	return nil
}

// AddFromResource is a wrapper around gtk_builder_add_from_resource().
func (v *Builder) AddFromResource(path string) error {
	cstr := C.CString(path)
	defer C.free(unsafe.Pointer(cstr))
	var err *C.GError = nil
	res := C.gtk_builder_add_from_resource(v.native(), (*C.gchar)(cstr), &err)
	if res == 0 {
		defer C.g_error_free(err)
		return errors.New(goString(err.message))
	}
	return nil
}

// AddFromString is a wrapper around gtk_builder_add_from_string().
func (v *Builder) AddFromString(str string) error {
	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))
	length := (C.gsize)(len(str))
	var err *C.GError = nil
	res := C.gtk_builder_add_from_string(v.native(), (*C.gchar)(cstr), length, &err)
	if res == 0 {
		defer C.g_error_free(err)
		return errors.New(goString(err.message))
	}
	return nil
}

// GetObject is a wrapper around gtk_builder_get_object(). The returned result
// is an IObject, so it will need to be type-asserted to the appropriate type before
// being used. For example, to get an object and type assert it as a window:
//
//   obj, err := builder.GetObject("window")
//   if err != nil {
//       // object not found
//       return
//   }
//   if w, ok := obj.(*gtk.Window); ok {
//       // do stuff with w here
//   } else {
//       // not a *gtk.Window
//   }
//
func (v *Builder) GetObject(name string) (glib.IObject, error) {
	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_builder_get_object(v.native(), (*C.gchar)(cstr))
	if c == nil {
		return nil, nil
	}
	return cast(c)
}

var (
	builderSignals = struct {
		sync.RWMutex
		m map[*C.GtkBuilder]map[string]interface{}
	}{
		m: make(map[*C.GtkBuilder]map[string]interface{}),
	}
)

// ConnectSignals is a wrapper around gtk_builder_connect_signals_full().
func (v *Builder) ConnectSignals(signals map[string]interface{}) {
	builderSignals.Lock()
	builderSignals.m[v.native()] = signals
	builderSignals.Unlock()

	C._gtk_builder_connect_signals_full(v.native())
}

/*
 * GtkCalendar
 */

// Calendar is a representation of GTK's GtkCalendar.
type Calendar struct {
	Widget
}

// native() returns a pointer to the underlying GtkCalendar.
func (v *Calendar) native() *C.GtkCalendar {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toGtkCalendar(ptr)
}

func marshalCalendar(p uintptr) (interface{}, error) {
	c := C.g_value_get_object(C.toGValue(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapCalendar(obj), nil
}

func wrapCalendar(obj *glib.Object) *Calendar {
	widget := wrapWidget(obj)
	return &Calendar{*widget}
}

// CalendarNew is a wrapper around gtk_calendar_new().
func CalendarNew() (*Calendar, error) {
	c := C.gtk_calendar_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapCalendar(obj), nil
}

// SelectMonth is a wrapper around gtk_calendar_select_month().
func (v *Calendar) SelectMonth(month, year uint) {
	C.gtk_calendar_select_month(v.native(), C.guint(month), C.guint(year))
}

// SelectDay is a wrapper around gtk_calendar_select_day().
func (v *Calendar) SelectDay(day uint) {
	C.gtk_calendar_select_day(v.native(), C.guint(day))
}

// MarkDay is a wrapper around gtk_calendar_mark_day().
func (v *Calendar) MarkDay(day uint) {
	C.gtk_calendar_mark_day(v.native(), C.guint(day))
}

// UnmarkDay is a wrapper around gtk_calendar_unmark_day().
func (v *Calendar) UnmarkDay(day uint) {
	C.gtk_calendar_unmark_day(v.native(), C.guint(day))
}

// GetDayIsMarked is a wrapper around gtk_calendar_get_day_is_marked().
func (v *Calendar) GetDayIsMarked(day uint) bool {
	c := C.gtk_calendar_get_day_is_marked(v.native(), C.guint(day))
	return gobool(c)
}

// ClearMarks is a wrapper around gtk_calendar_clear_marks().
func (v *Calendar) ClearMarks() {
	C.gtk_calendar_clear_marks(v.native())
}

// GetDisplayOptions is a wrapper around gtk_calendar_get_display_options().
func (v *Calendar) GetDisplayOptions() CalendarDisplayOptions {
	c := C.gtk_calendar_get_display_options(v.native())
	return CalendarDisplayOptions(c)
}

// SetDisplayOptions is a wrapper around gtk_calendar_set_display_options().
func (v *Calendar) SetDisplayOptions(flags CalendarDisplayOptions) {
	C.gtk_calendar_set_display_options(v.native(),
		C.GtkCalendarDisplayOptions(flags))
}

// GetDate is a wrapper around gtk_calendar_get_date().
func (v *Calendar) GetDate() (year, month, day uint) {
	var cyear, cmonth, cday C.guint
	C.gtk_calendar_get_date(v.native(), &cyear, &cmonth, &cday)
	return uint(cyear), uint(cmonth), uint(cday)
}

// TODO gtk_calendar_set_detail_func

// GetDetailWidthChars is a wrapper around gtk_calendar_get_detail_width_chars().
func (v *Calendar) GetDetailWidthChars() int {
	c := C.gtk_calendar_get_detail_width_chars(v.native())
	return int(c)
}

// SetDetailWidthChars is a wrapper around gtk_calendar_set_detail_width_chars().
func (v *Calendar) SetDetailWidthChars(chars int) {
	C.gtk_calendar_set_detail_width_chars(v.native(), C.gint(chars))
}

// GetDetailHeightRows is a wrapper around gtk_calendar_get_detail_height_rows().
func (v *Calendar) GetDetailHeightRows() int {
	c := C.gtk_calendar_get_detail_height_rows(v.native())
	return int(c)
}

// SetDetailHeightRows is a wrapper around gtk_calendar_set_detail_height_rows().
func (v *Calendar) SetDetailHeightRows(rows int) {
	C.gtk_calendar_set_detail_height_rows(v.native(), C.gint(rows))
}

/*
 * GtkClipboard
 */

// Clipboard is a wrapper around GTK's GtkClipboard.
type Clipboard struct {
	*glib.Object
}

// native returns a pointer to the underlying GtkClipboard.
func (v *Clipboard) native() *C.GtkClipboard {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toGtkClipboard(ptr)
}

func marshalClipboard(p uintptr) (interface{}, error) {
	c := C.g_value_get_object(C.toGValue(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapClipboard(obj), nil
}

func wrapClipboard(obj *glib.Object) *Clipboard {
	return &Clipboard{obj}
}

// Store is a wrapper around gtk_clipboard_store
func (v *Clipboard) Store() {
	C.gtk_clipboard_store(v.native())
}

// ClipboardGet() is a wrapper around gtk_clipboard_get().
func ClipboardGet(atom gdk.Atom) (*Clipboard, error) {
	c := C.gtk_clipboard_get(C.GdkAtom(unsafe.Pointer(atom.Native())))
	if c == nil {
		return nil, nilPtrErr
	}

	cb := &Clipboard{glib.Take(unsafe.Pointer(c))}
	return cb, nil
}

// ClipboardGetForDisplay() is a wrapper around gtk_clipboard_get_for_display().
func ClipboardGetForDisplay(display *gdk.Display, atom gdk.Atom) (*Clipboard, error) {
	displayPtr := (*C.GdkDisplay)(unsafe.Pointer(display.Native()))
	c := C.gtk_clipboard_get_for_display(displayPtr,
		C.GdkAtom(unsafe.Pointer(atom.Native())))
	if c == nil {
		return nil, nilPtrErr
	}

	cb := &Clipboard{glib.Take(unsafe.Pointer(c))}
	return cb, nil
}

// WaitIsTextAvailable is a wrapper around gtk_clipboard_wait_is_text_available
func (v *Clipboard) WaitIsTextAvailable() bool {
	c := C.gtk_clipboard_wait_is_text_available(v.native())
	return gobool(c)
}

// WaitForText is a wrapper around gtk_clipboard_wait_for_text
func (v *Clipboard) WaitForText() string {
	c := C.gtk_clipboard_wait_for_text(v.native())
	defer C.g_free(C.gpointer(c))
	return goString(c)
}

// SetText() is a wrapper around gtk_clipboard_set_text().
func (v *Clipboard) SetText(text string) {
	cstr := C.CString(text)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_clipboard_set_text(v.native(), (*C.gchar)(cstr),
		C.gint(len(text)))
}

// WaitIsRichTextAvailable is a wrapper around gtk_clipboard_wait_is_rich_text_available
func (v *Clipboard) WaitIsRichTextAvailable(buf *TextBuffer) bool {
	c := C.gtk_clipboard_wait_is_rich_text_available(v.native(), buf.native())
	return gobool(c)
}

// WaitIsUrisAvailable is a wrapper around gtk_clipboard_wait_is_uris_available
func (v *Clipboard) WaitIsUrisAvailable() bool {
	c := C.gtk_clipboard_wait_is_uris_available(v.native())
	return gobool(c)
}

// WaitIsImageAvailable is a wrapper around gtk_clipboard_wait_is_image_available
func (v *Clipboard) WaitIsImageAvailable() bool {
	c := C.gtk_clipboard_wait_is_image_available(v.native())
	return gobool(c)
}

// SetImage is a wrapper around gtk_clipboard_set_image
func (v *Clipboard) SetImage(pixbuf *gdk.Pixbuf) {
	C.gtk_clipboard_set_image(v.native(), (*C.GdkPixbuf)(unsafe.Pointer(pixbuf.Native())))
}

// WaitForImage is a wrapper around gtk_clipboard_wait_for_image
func (v *Clipboard) WaitForImage() (*gdk.Pixbuf, error) {
	c := C.gtk_clipboard_wait_for_image(v.native())
	if c == nil {
		return nil, nilPtrErr
	}

	p := &gdk.Pixbuf{glib.Take(unsafe.Pointer(c))}
	return p, nil
}

// WaitIsTargetAvailable is a wrapper around gtk_clipboard_wait_is_target_available
func (v *Clipboard) WaitIsTargetAvailable(target gdk.Atom) bool {
	c := C.gtk_clipboard_wait_is_target_available(v.native(), C.GdkAtom(unsafe.Pointer(target.Native())))
	return gobool(c)
}

// WaitForContents is a wrapper around gtk_clipboard_wait_for_contents
func (v *Clipboard) WaitForContents(target gdk.Atom) (*SelectionData, error) {
	c := C.gtk_clipboard_wait_for_contents(v.native(), C.GdkAtom(unsafe.Pointer(target.Native())))
	if c == nil {
		return nil, nilPtrErr
	}
	sd := wrapSelectionData(c)
	runtime.SetFinalizer(sd, (*SelectionData).Free)
	return sd, nil
}

/*
 * GtkCssProvider
 */

// CssProvider is a representation of GTK's GtkCssProvider.
type CssProvider struct {
	*glib.Object
}

func marshalCssProvider(p uintptr) (interface{}, error) {
	c := C.g_value_get_object(C.toGValue(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapCssProvider(obj), nil
}

func wrapCssProvider(obj *glib.Object) *CssProvider {
	return &CssProvider{obj}
}

// native returns a pointer to the underlying GtkCssProvider.
func (v *CssProvider) native() *C.GtkCssProvider {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toGtkCssProvider(ptr)
}

func (v *CssProvider) toStyleProvider() *C.GtkStyleProvider {
	if v == nil {
		return nil
	}
	return C.toGtkStyleProvider(unsafe.Pointer(v.native()))
}

// CssProviderNew is a wrapper around gtk_css_provider_new().
func CssProviderNew() (*CssProvider, error) {
	c := C.gtk_css_provider_new()
	if c == nil {
		return nil, nilPtrErr
	}

	obj := glib.Take(unsafe.Pointer(c))
	return wrapCssProvider(obj), nil
}

// LoadFromPath is a wrapper around gtk_css_provider_load_from_path().
func (v *CssProvider) LoadFromPath(path string) error {
	cpath := C.CString(path)
	defer C.free(unsafe.Pointer(cpath))
	var gerr *C.GError
	if C.gtk_css_provider_load_from_path(v.native(), (*C.gchar)(cpath), &gerr) == 0 {
		defer C.g_error_free(gerr)
		return errors.New(goString(gerr.message))
	}
	return nil
}

// LoadFromData is a wrapper around gtk_css_provider_load_from_data().
func (v *CssProvider) LoadFromData(data string) error {
	cdata := C.CString(data)
	defer C.free(unsafe.Pointer(cdata))
	var gerr *C.GError
	if C.gtk_css_provider_load_from_data(v.native(), (*C.gchar)(unsafe.Pointer(cdata)), C.gssize(len(data)), &gerr) == 0 {
		defer C.g_error_free(gerr)
		return errors.New(goString(gerr.message))
	}
	return nil
}

// ToString is a wrapper around gtk_css_provider_to_string().
func (v *CssProvider) ToString() string {
	c := C.gtk_css_provider_to_string(v.native())
	defer C.g_free(C.gpointer(c))
	return goString((*C.gchar)(c))
}

// CssProviderGetDefault is a wrapper around gtk_css_provider_get_default().
func CssProviderGetDefault() (*CssProvider, error) {
	c := C.gtk_css_provider_get_default()
	if c == nil {
		return nil, nilPtrErr
	}

	obj := glib.Take(unsafe.Pointer(c))
	return wrapCssProvider(obj), nil
}

// GetNamed is a wrapper around gtk_css_provider_get_named().
func CssProviderGetNamed(name string, variant string) (*CssProvider, error) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	cvariant := C.CString(variant)
	defer C.free(unsafe.Pointer(cvariant))

	c := C.gtk_css_provider_get_named((*C.gchar)(cname), (*C.gchar)(cvariant))
	if c == nil {
		return nil, nilPtrErr
	}

	obj := glib.Take(unsafe.Pointer(c))
	return wrapCssProvider(obj), nil
}

/*
 * GtkDrawingArea
 */

// DrawingArea is a representation of GTK's GtkDrawingArea.
type DrawingArea struct {
	Widget
}

// native returns a pointer to the underlying GtkDrawingArea.
func (v *DrawingArea) native() *C.GtkDrawingArea {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toGtkDrawingArea(ptr)
}

func marshalDrawingArea(p uintptr) (interface{}, error) {
	c := C.g_value_get_object(C.toGValue(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapDrawingArea(obj), nil
}

func wrapDrawingArea(obj *glib.Object) *DrawingArea {
	widget := wrapWidget(obj)
	return &DrawingArea{*widget}
}

// DrawingAreaNew is a wrapper around gtk_drawing_area_new().
func DrawingAreaNew() (*DrawingArea, error) {
	c := C.gtk_drawing_area_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapDrawingArea(obj), nil
}

/*
 * GtkEventBox
 */

// EventBox is a representation of GTK's GtkEventBox.
type EventBox struct {
	Bin
}

// native returns a pointer to the underlying GtkEventBox.
func (v *EventBox) native() *C.GtkEventBox {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toGtkEventBox(ptr)
}

func marshalEventBox(p uintptr) (interface{}, error) {
	c := C.g_value_get_object(C.toGValue(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapEventBox(obj), nil
}

func wrapEventBox(obj *glib.Object) *EventBox {
	bin := wrapBin(obj)
	return &EventBox{*bin}
}

// EventBoxNew is a wrapper around gtk_event_box_new().
func EventBoxNew() (*EventBox, error) {
	c := C.gtk_event_box_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapEventBox(obj), nil
}

// SetAboveChild is a wrapper around gtk_event_box_set_above_child().
func (v *EventBox) SetAboveChild(aboveChild bool) {
	C.gtk_event_box_set_above_child(v.native(), gbool(aboveChild))
}

// GetAboveChild is a wrapper around gtk_event_box_get_above_child().
func (v *EventBox) GetAboveChild() bool {
	c := C.gtk_event_box_get_above_child(v.native())
	return gobool(c)
}

// SetVisibleWindow is a wrapper around gtk_event_box_set_visible_window().
func (v *EventBox) SetVisibleWindow(visibleWindow bool) {
	C.gtk_event_box_set_visible_window(v.native(), gbool(visibleWindow))
}

// GetVisibleWindow is a wrapper around gtk_event_box_get_visible_window().
func (v *EventBox) GetVisibleWindow() bool {
	c := C.gtk_event_box_get_visible_window(v.native())
	return gobool(c)
}

/*
 * GtkFrame
 */

// Frame is a representation of GTK's GtkFrame.
type Frame struct {
	Bin
}

// native returns a pointer to the underlying GtkFrame.
func (v *Frame) native() *C.GtkFrame {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toGtkFrame(ptr)
}

func marshalFrame(p uintptr) (interface{}, error) {
	c := C.g_value_get_object(C.toGValue(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapFrame(obj), nil
}

func wrapFrame(obj *glib.Object) *Frame {
	bin := wrapBin(obj)
	return &Frame{*bin}
}

// FrameNew is a wrapper around gtk_frame_new().
func FrameNew(label string) (*Frame, error) {
	var cstr *C.char
	if label != "" {
		cstr = C.CString(label)
		defer C.free(unsafe.Pointer(cstr))
	}
	c := C.gtk_frame_new((*C.gchar)(cstr))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapFrame(obj), nil
}

// SetLabel is a wrapper around gtk_frame_set_label().
func (v *Frame) SetLabel(label string) {
	cstr := C.CString(label)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_frame_set_label(v.native(), (*C.gchar)(cstr))
}

// SetLabelWidget is a wrapper around gtk_frame_set_label_widget().
func (v *Frame) SetLabelWidget(labelWidget IWidget) {
	C.gtk_frame_set_label_widget(v.native(), labelWidget.toWidget())
}

// SetLabelAlign is a wrapper around gtk_frame_set_label_align().
func (v *Frame) SetLabelAlign(xAlign, yAlign float32) {
	C.gtk_frame_set_label_align(v.native(), C.gfloat(xAlign),
		C.gfloat(yAlign))
}

// SetShadowType is a wrapper around gtk_frame_set_shadow_type().
func (v *Frame) SetShadowType(t ShadowType) {
	C.gtk_frame_set_shadow_type(v.native(), C.GtkShadowType(t))
}

// GetLabel is a wrapper around gtk_frame_get_label().
func (v *Frame) GetLabel() string {
	c := C.gtk_frame_get_label(v.native())
	return goString(c)
}

// GetLabelAlign is a wrapper around gtk_frame_get_label_align().
func (v *Frame) GetLabelAlign() (xAlign, yAlign float32) {
	var x, y C.gfloat
	C.gtk_frame_get_label_align(v.native(), &x, &y)
	return float32(x), float32(y)
}

// GetLabelWidget is a wrapper around gtk_frame_get_label_widget().
func (v *Frame) GetLabelWidget() (*Widget, error) {
	c := C.gtk_frame_get_label_widget(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapWidget(obj), nil
}

// GetShadowType is a wrapper around gtk_frame_get_shadow_type().
func (v *Frame) GetShadowType() ShadowType {
	c := C.gtk_frame_get_shadow_type(v.native())
	return ShadowType(c)
}

/*
 * GtkIconTheme
 */

// IconTheme is a representation of GTK's GtkIconTheme
type IconTheme struct {
	*glib.Object
}

// native returns a pointer to the underlying GtkIconTheme.
func (v *IconTheme) native() *C.GtkIconTheme {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toGtkIconTheme(ptr)
}

func marshalIconTheme(p uintptr) (interface{}, error) {
	c := C.g_value_get_object(C.toGValue(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapGrid(obj), nil
}

func wrapIconTheme(obj *glib.Object) *IconTheme {
	return &IconTheme{obj}
}

// IconThemeGetDefault is a wrapper around gtk_icon_theme_get_default().
func IconThemeGetDefault() (*IconTheme, error) {
	c := C.gtk_icon_theme_get_default()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.ToObject(unsafe.Pointer(c))
	return wrapIconTheme(obj), nil
}

// IconThemeGetForScreen is a wrapper around gtk_icon_theme_get_for_screen().
func IconThemeGetForScreen(screen gdk.Screen) (*IconTheme, error) {
	cScreen := (*C.GdkScreen)(unsafe.Pointer(screen.Native()))
	c := C.gtk_icon_theme_get_for_screen(cScreen)
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.ToObject(unsafe.Pointer(c))
	return wrapIconTheme(obj), nil
}

// LoadIcon is a wrapper around gtk_icon_theme_load_icon().
func (v *IconTheme) LoadIcon(iconName string, size int,
	flags IconLookupFlags) (*gdk.Pixbuf, error) {

	cstr := C.CString(iconName)
	defer C.free(unsafe.Pointer(cstr))
	var err *C.GError = nil
	c := C.gtk_icon_theme_load_icon(v.native(), (*C.gchar)(cstr),
		C.gint(size), C.GtkIconLookupFlags(flags), &err)
	if c == nil {
		defer C.g_error_free(err)
		return nil, errors.New(goString(err.message))
	}
	return &gdk.Pixbuf{glib.Take(unsafe.Pointer(c))}, nil
}

/*
 * GtkRange
 */

// Range is a representation of GTK's GtkRange.
type Range struct {
	Widget
}

// native returns a pointer to the underlying GtkRange.
func (v *Range) native() *C.GtkRange {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toGtkRange(ptr)
}

func marshalRange(p uintptr) (interface{}, error) {
	c := C.g_value_get_object(C.toGValue(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapRange(obj), nil
}

func wrapRange(obj *glib.Object) *Range {
	widget := wrapWidget(obj)
	return &Range{*widget}
}

// GetValue is a wrapper around gtk_range_get_value().
func (v *Range) GetValue() float64 {
	c := C.gtk_range_get_value(v.native())
	return float64(c)
}

// SetValue is a wrapper around gtk_range_set_value().
func (v *Range) SetValue(value float64) {
	C.gtk_range_set_value(v.native(), C.gdouble(value))
}

// SetIncrements() is a wrapper around gtk_range_set_increments().
func (v *Range) SetIncrements(step, page float64) {
	C.gtk_range_set_increments(v.native(), C.gdouble(step), C.gdouble(page))
}

// SetRange() is a wrapper around gtk_range_set_range().
func (v *Range) SetRange(min, max float64) {
	C.gtk_range_set_range(v.native(), C.gdouble(min), C.gdouble(max))
}

// IRecentChooser is an interface type implemented by all structs
// embedding a RecentChooser.  It is meant to be used as an argument type
// for wrapper functions that wrap around a C GTK function taking a
// GtkWidget.
type IRecentChooser interface {
	toRecentChooser() *C.GtkRecentChooser
}

/*
 * GtkRecentChooser
 */

// RecentChooser is a representation of GTK's GtkRecentChooser GInterface.
type RecentChooser struct {
	glib.Interface
}

// native returns a pointer to the underlying GtkRecentChooser.
func (v *RecentChooser) native() *C.GtkRecentChooser {
	return C.toGtkRecentChooser(unsafe.Pointer(v.Native()))
}

func marshalRecentChooser(p uintptr) (interface{}, error) {
	c := C.g_value_get_object(C.toGValue(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapRecentChooser(*glib.InterfaceFromObjectNew(obj)), nil
}

func wrapRecentChooser(intf glib.Interface) *RecentChooser {
	return &RecentChooser{intf}
}

func (v *RecentChooser) toRecentChooser() *C.GtkRecentChooser {
	return v.native()
}

func (v *RecentChooser) GetCurrentUri() string {
	c := C.gtk_recent_chooser_get_current_uri(v.native())
	defer C.g_free(C.gpointer(c))
	return goString(c)
}

func (v *RecentChooser) AddFilter(filter *RecentFilter) {
	C.gtk_recent_chooser_add_filter(v.native(), filter.native())
}

func (v *RecentChooser) RemoveFilter(filter *RecentFilter) {
	C.gtk_recent_chooser_remove_filter(v.native(), filter.native())
}

/*
 * GtkRecentChooserMenu
 */

// RecentChooserMenu is a representation of GTK's GtkRecentChooserMenu.
type RecentChooserMenu struct {
	Menu
	// Interfaces
	RecentChooser
}

// native returns a pointer to the underlying GtkRecentManager.
func (v *RecentChooserMenu) native() *C.GtkRecentChooserMenu {
	if v == nil || v.Object == nil {
		return nil
	}
	return C.toGtkRecentChooserMenu(unsafe.Pointer(v.Menu.Native()))
}

func wrapRecentChooserMenu(obj *glib.Object) *RecentChooserMenu {
	menu := wrapMenu(obj)
	rc := wrapRecentChooser(*glib.InterfaceFromObjectNew(obj))
	return &RecentChooserMenu{*menu, *rc}
}

/*
 * GtkRecentFilter
 */

// RecentFilter is a representation of GTK's GtkRecentFilter.
type RecentFilter struct {
	glib.InitiallyUnowned
}

// native returns a pointer to the underlying GtkRecentFilter.
func (v *RecentFilter) native() *C.GtkRecentFilter {
	if v == nil || v.Object == nil {
		return nil
	}
	return C.toGtkRecentFilter(unsafe.Pointer(v.Native()))
}

func wrapRecentFilter(obj *glib.Object) *RecentFilter {
	return &RecentFilter{glib.InitiallyUnowned{obj}}
}

// RecentFilterNew is a wrapper around gtk_recent_filter_new().
func RecentFilterNew() (*RecentFilter, error) {
	c := C.gtk_recent_filter_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapRecentFilter(obj), nil
}

/*
 * GtkRecentManager
 */

// RecentManager is a representation of GTK's GtkRecentManager.
type RecentManager struct {
	*glib.Object
}

// native returns a pointer to the underlying GtkRecentManager.
func (v *RecentManager) native() *C.GtkRecentManager {
	if v == nil || v.Object == nil {
		return nil
	}
	return C.toGtkRecentManager(unsafe.Pointer(v.Native()))
}

func marshalRecentManager(p uintptr) (interface{}, error) {
	c := C.g_value_get_object(C.toGValue(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapRecentManager(obj), nil
}

func wrapRecentManager(obj *glib.Object) *RecentManager {
	return &RecentManager{obj}
}

// RecentManagerGetDefault is a wrapper around gtk_recent_manager_get_default().
func RecentManagerGetDefault() (*RecentManager, error) {
	c := C.gtk_recent_manager_get_default()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	v := wrapRecentManager(obj)
	return v, nil
}

// AddItem is a wrapper around gtk_recent_manager_add_item().
func (v *RecentManager) AddItem(fileURI string) bool {
	cstr := C.CString(fileURI)
	defer C.free(unsafe.Pointer(cstr))
	cok := C.gtk_recent_manager_add_item(v.native(), (*C.gchar)(cstr))
	return gobool(cok)
}

/*
 * GtkScrollable
 */

// IScrollable is an interface type implemented by all structs
// embedding a Scrollable.  It is meant to be used as an argument type
// for wrapper functions that wrap around a C GTK function taking a
// GtkScrollable.
type IScrollable interface {
	toScrollable() *C.GtkScrollable
}

// Scrollable is a representation of GTK's GtkScrollable GInterface.
type Scrollable struct {
	glib.Interface
}

// native() returns a pointer to the underlying GObject as a GtkScrollable.
func (v *Scrollable) native() *C.GtkScrollable {
	return C.toGtkScrollable(unsafe.Pointer(v.Native()))
}

func marshalScrollable(p uintptr) (interface{}, error) {
	c := C.g_value_get_object(C.toGValue(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapScrollable(*glib.InterfaceFromObjectNew(obj)), nil
}

func wrapScrollable(intf glib.Interface) *Scrollable {
	return &Scrollable{intf}
}

func (v *Scrollable) toScrollable() *C.GtkScrollable {
	if v == nil {
		return nil
	}
	return v.native()
}

// SetHAdjustment is a wrapper around gtk_scrollable_set_hadjustment().
func (v *Scrollable) SetHAdjustment(adjustment *Adjustment) {
	C.gtk_scrollable_set_hadjustment(v.native(), adjustment.native())
}

// GetHAdjustment is a wrapper around gtk_scrollable_get_hadjustment().
func (v *Scrollable) GetHAdjustment() (*Adjustment, error) {
	c := C.gtk_scrollable_get_hadjustment(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapAdjustment(obj), nil
}

// SetVAdjustment is a wrapper around gtk_scrollable_set_vadjustment().
func (v *Scrollable) SetVAdjustment(adjustment *Adjustment) {
	C.gtk_scrollable_set_vadjustment(v.native(), adjustment.native())
}

// GetVAdjustment is a wrapper around gtk_scrollable_get_vadjustment().
func (v *Scrollable) GetVAdjustment() (*Adjustment, error) {
	c := C.gtk_scrollable_get_vadjustment(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapAdjustment(obj), nil
}

/*
 * GtkScrollbar
 */

// Scrollbar is a representation of GTK's GtkScrollbar.
type Scrollbar struct {
	Range
}

// native returns a pointer to the underlying GtkScrollbar.
func (v *Scrollbar) native() *C.GtkScrollbar {
	if v == nil || v.Object == nil {
		return nil
	}
	return C.toGtkScrollbar(unsafe.Pointer(v.Native()))
}

func marshalScrollbar(p uintptr) (interface{}, error) {
	c := C.g_value_get_object(C.toGValue(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapScrollbar(obj), nil
}

func wrapScrollbar(obj *glib.Object) *Scrollbar {
	rng := wrapRange(obj)
	return &Scrollbar{*rng}
}

// ScrollbarNew is a wrapper around gtk_scrollbar_new().
func ScrollbarNew(orientation Orientation, adjustment *Adjustment) (*Scrollbar, error) {
	c := C.gtk_scrollbar_new(C.GtkOrientation(orientation), adjustment.native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapScrollbar(obj), nil
}

/*
 * GtkScrolledWindow
 */

// ScrolledWindow is a representation of GTK's GtkScrolledWindow.
type ScrolledWindow struct {
	Bin
}

// native returns a pointer to the underlying GtkScrolledWindow.
func (v *ScrolledWindow) native() *C.GtkScrolledWindow {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toGtkScrolledWindow(ptr)
}

func marshalScrolledWindow(p uintptr) (interface{}, error) {
	c := C.g_value_get_object(C.toGValue(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapScrolledWindow(obj), nil
}

func wrapScrolledWindow(obj *glib.Object) *ScrolledWindow {
	bin := wrapBin(obj)
	return &ScrolledWindow{*bin}
}

// ScrolledWindowNew is a wrapper around gtk_scrolled_window_new().
func ScrolledWindowNew(hadjustment, vadjustment *Adjustment) (*ScrolledWindow, error) {
	var hadj, vadj *C.GtkAdjustment
	if hadjustment != nil {
		hadj = hadjustment.native()
	}
	if vadjustment != nil {
		vadj = vadjustment.native()
	}
	c := C.gtk_scrolled_window_new(hadj, vadj)
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapScrolledWindow(obj), nil
}

// SetPolicy is a wrapper around gtk_scrolled_window_set_policy().
func (v *ScrolledWindow) SetPolicy(hScrollbarPolicy, vScrollbarPolicy PolicyType) {
	C.gtk_scrolled_window_set_policy(v.native(),
		C.GtkPolicyType(hScrollbarPolicy),
		C.GtkPolicyType(vScrollbarPolicy))
}

// GetHAdjustment is a wrapper around gtk_scrolled_window_get_hadjustment().
func (v *ScrolledWindow) GetHAdjustment() *Adjustment {
	c := C.gtk_scrolled_window_get_hadjustment(v.native())
	if c == nil {
		return nil
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapAdjustment(obj)
}

// SetHAdjustment is a wrapper around gtk_scrolled_window_set_hadjustment().
func (v *ScrolledWindow) SetHAdjustment(adjustment *Adjustment) {
	C.gtk_scrolled_window_set_hadjustment(v.native(), adjustment.native())
}

// GetVAdjustment is a wrapper around gtk_scrolled_window_get_vadjustment().
func (v *ScrolledWindow) GetVAdjustment() *Adjustment {
	c := C.gtk_scrolled_window_get_vadjustment(v.native())
	if c == nil {
		return nil
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapAdjustment(obj)
}

// SetVAdjustment is a wrapper around gtk_scrolled_window_set_vadjustment().
func (v *ScrolledWindow) SetVAdjustment(adjustment *Adjustment) {
	C.gtk_scrolled_window_set_vadjustment(v.native(), adjustment.native())
}

// GetShadowType is a wrapper around gtk_scrolled_window_get_shadow_type().
func (v *ScrolledWindow) GetShadowType() ShadowType {
	c := C.gtk_scrolled_window_get_shadow_type(v.native())
	return ShadowType(c)
}

// SetShadowType is a wrapper around gtk_scrolled_window_set_vadjustment().
func (v *ScrolledWindow) SetShadowType(shadowType ShadowType) {
	C.gtk_scrolled_window_set_shadow_type(v.native(), C.GtkShadowType(shadowType))
}

// GetMinContentWidth is a wrapper around gtk_scrolled_window_get_min_content_width().
func (v *ScrolledWindow) GetMinContentWidth() int {
	c := C.gtk_scrolled_window_get_min_content_width(v.native())
	return int(c)
}

// SetMinContentWidth is a wrapper around gtk_scrolled_window_set_min_content_width().
func (v *ScrolledWindow) SetMinContentWidth(width int) {
	C.gtk_scrolled_window_set_min_content_width(v.native(), C.gint(width))
}

// GetMinContentHeight is a wrapper around gtk_scrolled_window_get_min_content_height().
func (v *ScrolledWindow) GetMinContentHeight() int {
	c := C.gtk_scrolled_window_get_min_content_height(v.native())
	return int(c)
}

// SetMinContentHeight is a wrapper around gtk_scrolled_window_set_min_content_height().
func (v *ScrolledWindow) SetMinContentHeight(height int) {
	C.gtk_scrolled_window_set_min_content_height(v.native(), C.gint(height))
}

// GetMaxContentWidth is a wrapper around gtk_scrolled_window_get_max_content_width().
func (v *ScrolledWindow) GetMaxContentWidth() int {
	c := C.gtk_scrolled_window_get_max_content_width(v.native())
	return int(c)
}

// SetMaxContentWidth is a wrapper around gtk_scrolled_window_set_max_content_width().
func (v *ScrolledWindow) SetMaxContentWidth(width int) {
	C.gtk_scrolled_window_set_max_content_width(v.native(), C.gint(width))
}

// GetMaxContentHeight is a wrapper around gtk_scrolled_window_get_max_content_height().
func (v *ScrolledWindow) GetMaxContentHeight() int {
	c := C.gtk_scrolled_window_get_max_content_height(v.native())
	return int(c)
}

// SetMaxContentHeight is a wrapper around gtk_scrolled_window_set_max_content_height().
func (v *ScrolledWindow) SetMaxContentHeight(height int) {
	C.gtk_scrolled_window_set_max_content_height(v.native(), C.gint(height))
}

// GetPropagateNaturalWidth is a wrapper around gtk_scrolled_window_get_propagate_natural_width().
func (v *ScrolledWindow) GetPropagateNaturalWidth() bool {
	c := C.gtk_scrolled_window_get_propagate_natural_width(v.native())
	return gobool(c)
}

// SetPropagateNaturalWidth is a wrapper around gtk_scrolled_window_set_propagate_natural_width().
func (v *ScrolledWindow) SetPropagateNaturalWidth(propagate bool) {
	C.gtk_scrolled_window_set_propagate_natural_width(v.native(), gbool(propagate))
}

// GetPropagateNaturalHeight is a wrapper around gtk_scrolled_window_get_propagate_natural_height().
func (v *ScrolledWindow) GetPropagateNaturalHeight() bool {
	c := C.gtk_scrolled_window_get_propagate_natural_height(v.native())
	return gobool(c)
}

// SetPropagateNaturalHeight is a wrapper around gtk_scrolled_window_set_propagate_natural_height().
func (v *ScrolledWindow) SetPropagateNaturalHeight(propagate bool) {
	C.gtk_scrolled_window_set_propagate_natural_height(v.native(), gbool(propagate))
}

/*
* GtkSelectionData
 */
type SelectionData struct {
	gtkSelectionData *C.GtkSelectionData
}

// native returns a pointer to the underlying GtkSelectionData.
func (v *SelectionData) native() *C.GtkSelectionData {
	if v == nil {
		return nil
	}
	return v.gtkSelectionData
}

func marshalSelectionData(p uintptr) (interface{}, error) {
	c := C.g_value_get_boxed(C.toGValue(unsafe.Pointer(p)))
	c2 := (*C.GtkSelectionData)(unsafe.Pointer(c))
	return wrapSelectionData(c2), nil
}

func wrapSelectionData(obj *C.GtkSelectionData) *SelectionData {
	return &SelectionData{obj}
}

// GetLength is a wrapper around gtk_selection_data_get_length
func (v *SelectionData) GetLength() int {
	return int(C.gtk_selection_data_get_length(v.native()))
}

// GetData is a wrapper around gtk_selection_data_get_data_with_length.
// It returns a slice of the correct size with the selection's data.
func (v *SelectionData) GetData() (data []byte) {
	var length C.gint
	c := C.gtk_selection_data_get_data_with_length(v.native(), &length)
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&data))
	sliceHeader.Data = uintptr(unsafe.Pointer(c))
	sliceHeader.Len = int(length)
	sliceHeader.Cap = int(length)
	return
}

func (v *SelectionData) Free() {
	C.gtk_selection_data_free(v.native())
}

/*
 * GtkSeparator
 */

// Separator is a representation of GTK's GtkSeparator.
type Separator struct {
	Widget
}

// native returns a pointer to the underlying GtkSeperator.
func (v *Separator) native() *C.GtkSeparator {
	if v == nil || v.Object == nil {
		return nil
	}
	return C.toGtkSeparator(unsafe.Pointer(v.Native()))
}

func marshalSeparator(p uintptr) (interface{}, error) {
	c := C.g_value_get_object(C.toGValue(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapSeparator(obj), nil
}

func wrapSeparator(obj *glib.Object) *Separator {
	widget := wrapWidget(obj)
	return &Separator{*widget}
}

// SeparatorNew is a wrapper around gtk_separator_new().
func SeparatorNew(orientation Orientation) (*Separator, error) {
	c := C.gtk_separator_new(C.GtkOrientation(orientation))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapSeparator(obj), nil
}

/*
 * GtkSizeGroup
 */

// SizeGroup is a representation of GTK's GtkSizeGroup
type SizeGroup struct {
	*glib.Object
}

// native() returns a pointer to the underlying GtkSizeGroup
func (v *SizeGroup) native() *C.GtkSizeGroup {
	if v == nil || v.Object == nil {
		return nil
	}
	return C.toGtkSizeGroup(unsafe.Pointer(v.Native()))
}

func marshalSizeGroup(p uintptr) (interface{}, error) {
	c := C.g_value_get_object(C.toGValue(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return &SizeGroup{obj}, nil
}

func wrapSizeGroup(obj *glib.Object) *SizeGroup {
	return &SizeGroup{obj}
}

// SizeGroupNew is a wrapper around gtk_size_group_new().
func SizeGroupNew(mode SizeGroupMode) (*SizeGroup, error) {
	c := C.gtk_size_group_new(C.GtkSizeGroupMode(mode))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapSizeGroup(obj), nil
}

func (v *SizeGroup) SetMode(mode SizeGroupMode) {
	C.gtk_size_group_set_mode(v.native(), C.GtkSizeGroupMode(mode))
}

func (v *SizeGroup) GetMode() SizeGroupMode {
	return SizeGroupMode(C.gtk_size_group_get_mode(v.native()))
}

func (v *SizeGroup) AddWidget(widget IWidget) {
	C.gtk_size_group_add_widget(v.native(), widget.toWidget())
}

func (v *SizeGroup) RemoveWidget(widget IWidget) {
	C.gtk_size_group_remove_widget(v.native(), widget.toWidget())
}

func (v *SizeGroup) GetWidgets() *glib.SList {
	c := C.gtk_size_group_get_widgets(v.native())
	if c == nil {
		return nil
	}
	return glib.WrapSList(uintptr(unsafe.Pointer(c)))
}

/*
 * GtkTargetEntry
 */

// TargetEntry is a representation of GTK's GtkTargetEntry
type TargetEntry struct {
	gtkTargetEntry *C.GtkTargetEntry
}

func (v *TargetEntry) native() *C.GtkTargetEntry {
	if v == nil {
		return nil
	}
	return v.gtkTargetEntry
}

func marshalTargetEntry(p uintptr) (interface{}, error) {
	c := C.g_value_get_boxed(C.toGValue(unsafe.Pointer(p)))
	c2 := (*C.GtkTargetEntry)(unsafe.Pointer(c))
	return wrapTargetEntry(c2), nil
}

func wrapTargetEntry(obj *C.GtkTargetEntry) *TargetEntry {
	return &TargetEntry{obj}
}

// TargetEntryNew is a wrapper aroud gtk_target_entry_new().
func TargetEntryNew(target string, flags TargetFlags, info uint) (*TargetEntry, error) {
	cstr := C.CString(target)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_target_entry_new((*C.gchar)(cstr), C.guint(flags), C.guint(info))
	if c == nil {
		return nil, nilPtrErr
	}
	t := wrapTargetEntry(c)
	runtime.SetFinalizer(t, (*TargetEntry).free)
	return t, nil
}

func (v *TargetEntry) free() {
	C.gtk_target_entry_free(v.native())
}

/*
 * GtkViewport
 */

// Viewport is a representation of GTK's GtkViewport GInterface.
type Viewport struct {
	Bin
	// Interfaces
	Scrollable
}

// IViewport is an interface type implemented by all structs
// embedding a Viewport.  It is meant to be used as an argument type
// for wrapper functions that wrap around a C GTK function taking a
// GtkViewport.
type IViewport interface {
	toViewport() *C.GtkViewport
}

// native() returns a pointer to the underlying GObject as a GtkViewport.
func (v *Viewport) native() *C.GtkViewport {
	if v == nil || v.Object == nil {
		return nil
	}
	return C.toGtkViewport(unsafe.Pointer(v.Bin.Native()))
}

func wrapViewport(obj *glib.Object) *Viewport {
	bin := wrapBin(obj)
	s := wrapScrollable(*glib.InterfaceFromObjectNew(obj))
	return &Viewport{*bin, *s}
}

func (v *Viewport) toViewport() *C.GtkViewport {
	if v == nil {
		return nil
	}
	return v.native()
}

// ViewportNew() is a wrapper around gtk_viewport_new().
func ViewportNew(hadjustment, vadjustment *Adjustment) (*Viewport, error) {
	c := C.gtk_viewport_new(hadjustment.native(), vadjustment.native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapViewport(obj), nil
}

/*
func (v *Viewport) SetHAdjustment(adjustment *Adjustment) {
	wrapScrollable(v.Object).SetHAdjustment(adjustment)
}

func (v *Viewport) GetHAdjustment() (*Adjustment, error) {
	return wrapScrollable(v.Object).GetHAdjustment()
}

func (v *Viewport) SetVAdjustment(adjustment *Adjustment) {
	wrapScrollable(v.Object).SetVAdjustment(adjustment)
}

func (v *Viewport) GetVAdjustment() (*Adjustment, error) {
	return wrapScrollable(v.Object).GetVAdjustment()
}
*/

type WrapFn interface{}

var WrapMap = map[string]WrapFn{
	"GtkAccelGroup":          wrapAccelGroup,
	"GtkAccelMao":            wrapAccelMap,
	"GtkAdjustment":          wrapAdjustment,
	"GtkApplicationWindow":   wrapApplicationWindow,
	"GtkAssistant":           wrapAssistant,
	"GtkBin":                 wrapBin,
	"GtkBox":                 wrapBox,
	"GtkButton":              wrapButton,
	"GtkCalendar":            wrapCalendar,
	"GtkCellLayout":          wrapCellLayout,
	"GtkCellRenderer":        wrapCellRenderer,
	"GtkCellRendererSpinner": wrapCellRendererSpinner,
	"GtkCellRendererPixbuf":  wrapCellRendererPixbuf,
	"GtkCellRendererText":    wrapCellRendererText,
	"GtkCellRendererToggle":  wrapCellRendererToggle,
	"GtkCheckButton":         wrapCheckButton,
	"GtkCheckMenuItem":       wrapCheckMenuItem,
	"GtkClipboard":           wrapClipboard,
	"GtkColorButton":         wrapColorButton,
	"GtkContainer":           wrapContainer,
	"GtkDialog":              wrapDialog,
	"GtkDrawingArea":         wrapDrawingArea,
	"GtkEditable":            wrapEditable,
	"GtkEntry":               wrapEntry,
	"GtkEntryBuffer":         wrapEntryBuffer,
	"GtkEntryCompletion":     wrapEntryCompletion,
	"GtkEventBox":            wrapEventBox,
	"GtkExpander":            wrapExpander,
	"GtkFrame":               wrapFrame,
	"GtkFileChooser":         wrapFileChooser,
	"GtkFileChooserButton":   wrapFileChooserButton,
	"GtkFileChooserDialog":   wrapFileChooserDialog,
	"GtkFileChooserWidget":   wrapFileChooserWidget,
	"GtkFontButton":          wrapFontButton,
	"GtkGrid":                wrapGrid,
	"GtkIconView":            wrapIconView,
	"GtkImage":               wrapImage,
	"GtkLabel":               wrapLabel,
	"GtkLayout":              wrapLayout,
	"GtkLinkButton":          wrapLinkButton,
	"GtkListStore":           wrapListStore,
	"GtkMenu":                wrapMenu,
	"GtkMenuBar":             wrapMenuBar,
	"GtkMenuButton":          wrapMenuButton,
	"GtkMenuItem":            wrapMenuItem,
	"GtkMenuShell":           wrapMenuShell,
	"GtkMessageDialog":       wrapMessageDialog,
	"GtkNotebook":            wrapNotebook,
	"GtkOffscreenWindow":     wrapOffscreenWindow,
	"GtkOrientable":          wrapOrientable,
	"GtkOverlay":             wrapOverlay,
	"GtkPaned":               wrapPaned,
	"GtkProgressBar":         wrapProgressBar,
	"GtkRadioButton":         wrapRadioButton,
	"GtkRadioMenuItem":       wrapRadioMenuItem,
	"GtkRange":               wrapRange,
	"GtkRecentChooser":       wrapRecentChooser,
	"GtkRecentChooserMenu":   wrapRecentChooserMenu,
	"GtkRecentFilter":        wrapRecentFilter,
	"GtkRecentManager":       wrapRecentManager,
	"GtkScaleButton":         wrapScaleButton,
	"GtkScale":               wrapScale,
	"GtkScrollable":          wrapScrollable,
	"GtkScrollbar":           wrapScrollbar,
	"GtkScrolledWindow":      wrapScrolledWindow,
	"GtkSearchEntry":         wrapSearchEntry,
	"GtkSeparator":           wrapSeparator,
	"GtkSeparatorMenuItem":   wrapSeparatorMenuItem,
	"GtkSeparatorToolItem":   wrapSeparatorToolItem,
	"GtkSpinButton":          wrapSpinButton,
	"GtkSpinner":             wrapSpinner,
	"GtkStatusbar":           wrapStatusbar,
	"GtkSwitch":              wrapSwitch,
	"GtkTextView":            wrapTextView,
	"GtkTextBuffer":          wrapTextBuffer,
	"GtkTextTag":             wrapTextTag,
	"GtkTextTagTable":        wrapTextTagTable,
	"GtkToggleButton":        wrapToggleButton,
	"GtkToolbar":             wrapToolbar,
	"GtkToolButton":          wrapToolButton,
	"GtkToolItem":            wrapToolItem,
	"GtkTreeModel":           wrapTreeModel,
	"GtkTreeModelFilter":     wrapTreeModelFilter,
	"GtkTreeSelection":       wrapTreeSelection,
	"GtkTreeStore":           wrapTreeStore,
	"GtkTreeView":            wrapTreeView,
	"GtkTreeViewColumn":      wrapTreeViewColumn,
	"GtkViewport":            wrapViewport,
	"GtkVolumeButton":        wrapVolumeButton,
	"GtkWidget":              wrapWidget,
	"GtkWindow":              wrapWindow,
}

// cast takes a native GObject and casts it to the appropriate Go struct.
//TODO change all wrapFns to return an IObject
func cast(c *C.GObject) (glib.IObject, error) {
	var (
		className = goString(C.object_get_class_name(c))
		obj       = glib.Take(unsafe.Pointer(c))
	)

	fn, ok := WrapMap[className]
	if !ok {
		return nil, errors.New("unrecognized class name '" + className + "'")
	}

	rf := reflect.ValueOf(fn)
	if rf.Type().Kind() != reflect.Func {
		return nil, errors.New("wraper is not a function")
	}

	v := reflect.ValueOf(obj)
	rv := rf.Call([]reflect.Value{v})

	if len(rv) != 1 {
		return nil, errors.New("wrapper did not return")
	}

	if k := rv[0].Kind(); k != reflect.Ptr {
		return nil, fmt.Errorf("wrong return type %s", k)
	}

	ret, ok := rv[0].Interface().(glib.IObject)
	if !ok {
		return nil, errors.New("did not return an IObject")
	}

	return ret, nil
}
