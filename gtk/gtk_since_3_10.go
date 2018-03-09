// Same copyright and license as the rest of the files in this project
// This file contains accelerator related functions and structures

// +build !gtk_3_6,!gtk_3_8
// not use this: go build -tags gtk_3_8'. Otherwise, if no build tags are used, GTK 3.10

package gtk

// #cgo pkg-config: gtk+-3.0
// #include <stdlib.h>
// #include <gtk/gtk.h>
// #include "gtk_since_3_10.go.h"
import "C"
import (
	"unsafe"

	"github.com/d2r2/gotk3/gdk"
	"github.com/d2r2/gotk3/glib"
)

func init() {
	tm := []glib.TypeMarshaler{
		// Enums
		{glib.Type(C.gtk_revealer_transition_type_get_type()), marshalRevealerTransitionType},
		{glib.Type(C.gtk_stack_transition_type_get_type()), marshalStackTransitionType},

		// Objects/Interfaces
		{glib.Type(C.gtk_header_bar_get_type()), marshalHeaderBar},
		{glib.Type(C.gtk_list_box_get_type()), marshalListBox},
		{glib.Type(C.gtk_list_box_row_get_type()), marshalListBoxRow},
		{glib.Type(C.gtk_revealer_get_type()), marshalRevealer},
		{glib.Type(C.gtk_search_bar_get_type()), marshalSearchBar},
		{glib.Type(C.gtk_stack_get_type()), marshalStack},
		{glib.Type(C.gtk_stack_switcher_get_type()), marshalStackSwitcher},
	}
	glib.RegisterGValueMarshalers(tm)

	//Contribute to casting
	for k, v := range map[string]WrapFn{
		"GtkHeaderBar":  wrapHeaderBar,
		"GtkListBox":    wrapListBox,
		"GtkListBoxRow": wrapListBoxRow,
		"GtkRevealer":   wrapRevealer,
		"GtkSearchBar":  wrapSearchBar,
		"GtkStack":      wrapStack,
	} {
		WrapMap[k] = v
	}
}

/*
 * Constants
 */

const (
	ALIGN_BASELINE Align = C.GTK_ALIGN_BASELINE
)

// StackTransitionType is a representation of GTK's GtkStackTransitionType.
type StackTransitionType int

const (
	STACK_TRANSITION_TYPE_NONE             StackTransitionType = C.GTK_STACK_TRANSITION_TYPE_NONE
	STACK_TRANSITION_TYPE_CROSSFADE        StackTransitionType = C.GTK_STACK_TRANSITION_TYPE_CROSSFADE
	STACK_TRANSITION_TYPE_SLIDE_RIGHT      StackTransitionType = C.GTK_STACK_TRANSITION_TYPE_SLIDE_RIGHT
	STACK_TRANSITION_TYPE_SLIDE_LEFT       StackTransitionType = C.GTK_STACK_TRANSITION_TYPE_SLIDE_LEFT
	STACK_TRANSITION_TYPE_SLIDE_UP         StackTransitionType = C.GTK_STACK_TRANSITION_TYPE_SLIDE_UP
	STACK_TRANSITION_TYPE_SLIDE_DOWN       StackTransitionType = C.GTK_STACK_TRANSITION_TYPE_SLIDE_DOWN
	STACK_TRANSITION_TYPE_SLIDE_LEFT_RIGHT StackTransitionType = C.GTK_STACK_TRANSITION_TYPE_SLIDE_LEFT_RIGHT
	STACK_TRANSITION_TYPE_SLIDE_UP_DOWN    StackTransitionType = C.GTK_STACK_TRANSITION_TYPE_SLIDE_UP_DOWN
)

func marshalStackTransitionType(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return StackTransitionType(c), nil
}

/*
 * GtkButton
 */

// ButtonNewFromIconName is a wrapper around gtk_button_new_from_icon_name().
func ButtonNewFromIconName(iconName string, size IconSize) (*Button, error) {
	cstr := C.CString(iconName)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_button_new_from_icon_name((*C.gchar)(cstr),
		C.GtkIconSize(size))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapButton(obj), nil
}

/*
 * GtkLabel
 */

// GetLines() is a wrapper around gtk_label_get_lines().
func (v *Label) GetLines() int {
	c := C.gtk_label_get_lines(v.native())
	return int(c)
}

// SetLines() is a wrapper around gtk_label_set_lines().
func (v *Label) SetLines(lines int) {
	C.gtk_label_set_lines(v.native(), C.gint(lines))
}

/*
 * GtkSearchBar
 */

// SearchBar is a representation of GTK's GtkSearchBar.
type SearchBar struct {
	Bin
}

// native returns a pointer to the underlying GtkSearchBar.
func (v *SearchBar) native() *C.GtkSearchBar {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toGtkSearchBar(ptr)
}

func marshalSearchBar(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapSearchBar(obj), nil
}

func wrapSearchBar(obj *glib.Object) *SearchBar {
	return &SearchBar{Bin{Container{Widget{glib.InitiallyUnowned{obj}}}}}
}

// SearchBarNew is a wrapper around gtk_search_bar_new()
func SearchBarNew() (*SearchBar, error) {
	c := C.gtk_search_bar_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapSearchBar(obj), nil
}

// ConnectEntry is a wrapper around gtk_search_bar_connect_entry().
func (v *SearchBar) ConnectEntry(entry IEntry) {
	C.gtk_search_bar_connect_entry(v.native(), entry.toEntry())
}

// GetSearchMode is a wrapper around gtk_search_bar_get_search_mode().
func (v *SearchBar) GetSearchMode() bool {
	c := C.gtk_search_bar_get_search_mode(v.native())
	return gobool(c)
}

// SetSearchMode is a wrapper around gtk_search_bar_set_search_mode().
func (v *SearchBar) SetSearchMode(searchMode bool) {
	C.gtk_search_bar_set_search_mode(v.native(), gbool(searchMode))
}

// GetShowCloseButton is a wrapper arounb gtk_search_bar_get_show_close_button().
func (v *SearchBar) GetShowCloseButton() bool {
	c := C.gtk_search_bar_get_show_close_button(v.native())
	return gobool(c)
}

// SetShowCloseButton is a wrapper around gtk_search_bar_set_show_close_button()
func (v *SearchBar) SetShowCloseButton(visible bool) {
	C.gtk_search_bar_set_show_close_button(v.native(), gbool(visible))
}

// HandleEvent is a wrapper around gtk_search_bar_handle_event()
func (v *SearchBar) HandleEvent(event *gdk.Event) {
	e := (*C.GdkEvent)(unsafe.Pointer(event.Native()))
	C.gtk_search_bar_handle_event(v.native(), e)
}
