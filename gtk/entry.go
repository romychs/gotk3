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
	"unsafe"

	"github.com/d2r2/gotk3/glib"
)

/*
 * GtkEditable
 */

// Editable is a representation of GTK's GtkEditable GInterface.
type Editable struct {
	glib.Interface
}

// IEditable is an interface type implemented by all structs
// embedding an Editable.  It is meant to be used as an argument type
// for wrapper functions that wrap around a C GTK function taking a
// GtkEditable.
type IEditable interface {
	toEditable() *C.GtkEditable
}

// native() returns a pointer to the underlying GObject as a GtkEditable.
func (v *Editable) native() *C.GtkEditable {
	return C.toGtkEditable(unsafe.Pointer(v.Native()))
}

func marshalEditable(p uintptr) (interface{}, error) {
	c := C.g_value_get_object(C.toGValue(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	intf := glib.InterfaceFromObjectNew(obj)
	return wrapEditable(*intf), nil
}

func wrapEditable(intf glib.Interface) *Editable {
	return &Editable{intf}
}

func (v *Editable) toEditable() *C.GtkEditable {
	if v == nil {
		return nil
	}
	return v.native()
}

// SelectRegion is a wrapper around gtk_editable_select_region().
func (v *Editable) SelectRegion(startPos, endPos int) {
	C.gtk_editable_select_region(v.native(), C.gint(startPos),
		C.gint(endPos))
}

// GetSelectionBounds is a wrapper around gtk_editable_get_selection_bounds().
func (v *Editable) GetSelectionBounds() (start, end int, nonEmpty bool) {
	var cstart, cend C.gint
	c := C.gtk_editable_get_selection_bounds(v.native(), &cstart, &cend)
	return int(cstart), int(cend), gobool(c)
}

// InsertText is a wrapper around gtk_editable_insert_text(). The returned
// int is the position after the inserted text.
func (v *Editable) InsertText(newText string, position int) int {
	cstr := C.CString(newText)
	defer C.free(unsafe.Pointer(cstr))
	pos := new(C.gint)
	*pos = C.gint(position)
	C.gtk_editable_insert_text(v.native(), (*C.gchar)(cstr),
		C.gint(len(newText)), pos)
	return int(*pos)
}

// DeleteText is a wrapper around gtk_editable_delete_text().
func (v *Editable) DeleteText(startPos, endPos int) {
	C.gtk_editable_delete_text(v.native(), C.gint(startPos), C.gint(endPos))
}

// GetChars is a wrapper around gtk_editable_get_chars().
func (v *Editable) GetChars(startPos, endPos int) string {
	c := C.gtk_editable_get_chars(v.native(), C.gint(startPos),
		C.gint(endPos))
	defer C.g_free(C.gpointer(c))
	return goString(c)
}

// CutClipboard is a wrapper around gtk_editable_cut_clipboard().
func (v *Editable) CutClipboard() {
	C.gtk_editable_cut_clipboard(v.native())
}

// CopyClipboard is a wrapper around gtk_editable_copy_clipboard().
func (v *Editable) CopyClipboard() {
	C.gtk_editable_copy_clipboard(v.native())
}

// PasteClipboard is a wrapper around gtk_editable_paste_clipboard().
func (v *Editable) PasteClipboard() {
	C.gtk_editable_paste_clipboard(v.native())
}

// DeleteSelection is a wrapper around gtk_editable_delete_selection().
func (v *Editable) DeleteSelection() {
	C.gtk_editable_delete_selection(v.native())
}

// SetPosition is a wrapper around gtk_editable_set_position().
func (v *Editable) SetPosition(position int) {
	C.gtk_editable_set_position(v.native(), C.gint(position))
}

// GetPosition is a wrapper around gtk_editable_get_position().
func (v *Editable) GetPosition() int {
	c := C.gtk_editable_get_position(v.native())
	return int(c)
}

// SetEditable is a wrapper around gtk_editable_set_editable().
func (v *Editable) SetEditable(isEditable bool) {
	C.gtk_editable_set_editable(v.native(), gbool(isEditable))
}

// GetEditable is a wrapper around gtk_editable_get_editable().
func (v *Editable) GetEditable() bool {
	c := C.gtk_editable_get_editable(v.native())
	return gobool(c)
}

/*
 * GtkEntry
 */

// Entry is a representation of GTK's GtkEntry.
type Entry struct {
	Widget
	// Interfaces
	Editable
}

type IEntry interface {
	toEntry() *C.GtkEntry
}

func (v *Entry) toEntry() *C.GtkEntry {
	return v.native()
}

// native returns a pointer to the underlying GtkEntry.
func (v *Entry) native() *C.GtkEntry {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toGtkEntry(ptr)
}

func marshalEntry(p uintptr) (interface{}, error) {
	c := C.g_value_get_object(C.toGValue(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapEntry(obj), nil
}

func wrapEntry(obj *glib.Object) *Entry {
	widget := wrapWidget(obj)
	e := wrapEditable(*glib.InterfaceFromObjectNew(obj))
	return &Entry{*widget, *e}
}

// EntryNew() is a wrapper around gtk_entry_new().
func EntryNew() (*Entry, error) {
	c := C.gtk_entry_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapEntry(obj), nil
}

// EntryNewWithBuffer() is a wrapper around gtk_entry_new_with_buffer().
func EntryNewWithBuffer(buffer *EntryBuffer) (*Entry, error) {
	c := C.gtk_entry_new_with_buffer(buffer.native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapEntry(obj), nil
}

// GetBuffer() is a wrapper around gtk_entry_get_buffer().
func (v *Entry) GetBuffer() (*EntryBuffer, error) {
	c := C.gtk_entry_get_buffer(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return &EntryBuffer{obj}, nil
}

// SetBuffer() is a wrapper around gtk_entry_set_buffer().
func (v *Entry) SetBuffer(buffer *EntryBuffer) {
	C.gtk_entry_set_buffer(v.native(), buffer.native())
}

// SetText() is a wrapper around gtk_entry_set_text().
func (v *Entry) SetText(text string) {
	cstr := C.CString(text)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_entry_set_text(v.native(), (*C.gchar)(cstr))
}

// GetText() is a wrapper around gtk_entry_get_text().
func (v *Entry) GetText() (string, error) {
	c := C.gtk_entry_get_text(v.native())
	if c == nil {
		return "", nilPtrErr
	}
	return goString(c), nil
}

// GetTextLength() is a wrapper around gtk_entry_get_text_length().
func (v *Entry) GetTextLength() uint16 {
	c := C.gtk_entry_get_text_length(v.native())
	return uint16(c)
}

// TODO(jrick) GdkRectangle
/*
func (v *Entry) GetTextArea() {
}
*/

// SetVisibility() is a wrapper around gtk_entry_set_visibility().
func (v *Entry) SetVisibility(visible bool) {
	C.gtk_entry_set_visibility(v.native(), gbool(visible))
}

// SetInvisibleChar() is a wrapper around gtk_entry_set_invisible_char().
func (v *Entry) SetInvisibleChar(ch rune) {
	C.gtk_entry_set_invisible_char(v.native(), C.gunichar(ch))
}

// UnsetInvisibleChar() is a wrapper around gtk_entry_unset_invisible_char().
func (v *Entry) UnsetInvisibleChar() {
	C.gtk_entry_unset_invisible_char(v.native())
}

// SetMaxLength() is a wrapper around gtk_entry_set_max_length().
func (v *Entry) SetMaxLength(len int) {
	C.gtk_entry_set_max_length(v.native(), C.gint(len))
}

// GetActivatesDefault() is a wrapper around gtk_entry_get_activates_default().
func (v *Entry) GetActivatesDefault() bool {
	c := C.gtk_entry_get_activates_default(v.native())
	return gobool(c)
}

// GetHasFrame() is a wrapper around gtk_entry_get_has_frame().
func (v *Entry) GetHasFrame() bool {
	c := C.gtk_entry_get_has_frame(v.native())
	return gobool(c)
}

// GetWidthChars() is a wrapper around gtk_entry_get_width_chars().
func (v *Entry) GetWidthChars() int {
	c := C.gtk_entry_get_width_chars(v.native())
	return int(c)
}

// SetActivatesDefault() is a wrapper around gtk_entry_set_activates_default().
func (v *Entry) SetActivatesDefault(setting bool) {
	C.gtk_entry_set_activates_default(v.native(), gbool(setting))
}

// SetHasFrame() is a wrapper around gtk_entry_set_has_frame().
func (v *Entry) SetHasFrame(setting bool) {
	C.gtk_entry_set_has_frame(v.native(), gbool(setting))
}

// SetWidthChars() is a wrapper around gtk_entry_set_width_chars().
func (v *Entry) SetWidthChars(nChars int) {
	C.gtk_entry_set_width_chars(v.native(), C.gint(nChars))
}

// GetInvisibleChar() is a wrapper around gtk_entry_get_invisible_char().
func (v *Entry) GetInvisibleChar() rune {
	c := C.gtk_entry_get_invisible_char(v.native())
	return rune(c)
}

// SetAlignment() is a wrapper around gtk_entry_set_alignment().
func (v *Entry) SetAlignment(xalign float32) {
	C.gtk_entry_set_alignment(v.native(), C.gfloat(xalign))
}

// GetAlignment() is a wrapper around gtk_entry_get_alignment().
func (v *Entry) GetAlignment() float32 {
	c := C.gtk_entry_get_alignment(v.native())
	return float32(c)
}

// SetPlaceholderText() is a wrapper around gtk_entry_set_placeholder_text().
func (v *Entry) SetPlaceholderText(text string) {
	cstr := C.CString(text)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_entry_set_placeholder_text(v.native(), (*C.gchar)(cstr))
}

// GetPlaceholderText() is a wrapper around gtk_entry_get_placeholder_text().
func (v *Entry) GetPlaceholderText() (string, error) {
	c := C.gtk_entry_get_placeholder_text(v.native())
	if c == nil {
		return "", nilPtrErr
	}
	return goString(c), nil
}

// SetOverwriteMode() is a wrapper around gtk_entry_set_overwrite_mode().
func (v *Entry) SetOverwriteMode(overwrite bool) {
	C.gtk_entry_set_overwrite_mode(v.native(), gbool(overwrite))
}

// GetOverwriteMode() is a wrapper around gtk_entry_get_overwrite_mode().
func (v *Entry) GetOverwriteMode() bool {
	c := C.gtk_entry_get_overwrite_mode(v.native())
	return gobool(c)
}

// TODO(jrick) Pangolayout
/*
func (v *Entry) GetLayout() {
}
*/

// GetLayoutOffsets() is a wrapper around gtk_entry_get_layout_offsets().
func (v *Entry) GetLayoutOffsets() (x, y int) {
	var gx, gy C.gint
	C.gtk_entry_get_layout_offsets(v.native(), &gx, &gy)
	return int(gx), int(gy)
}

// LayoutIndexToTextIndex() is a wrapper around
// gtk_entry_layout_index_to_text_index().
func (v *Entry) LayoutIndexToTextIndex(layoutIndex int) int {
	c := C.gtk_entry_layout_index_to_text_index(v.native(),
		C.gint(layoutIndex))
	return int(c)
}

// TextIndexToLayoutIndex() is a wrapper around
// gtk_entry_text_index_to_layout_index().
func (v *Entry) TextIndexToLayoutIndex(textIndex int) int {
	c := C.gtk_entry_text_index_to_layout_index(v.native(),
		C.gint(textIndex))
	return int(c)
}

// TODO(jrick) PandoAttrList
/*
func (v *Entry) SetAttributes() {
}
*/

// TODO(jrick) PandoAttrList
/*
func (v *Entry) GetAttributes() {
}
*/

// GetMaxLength() is a wrapper around gtk_entry_get_max_length().
func (v *Entry) GetMaxLength() int {
	c := C.gtk_entry_get_max_length(v.native())
	return int(c)
}

// GetVisibility() is a wrapper around gtk_entry_get_visibility().
func (v *Entry) GetVisibility() bool {
	c := C.gtk_entry_get_visibility(v.native())
	return gobool(c)
}

// SetCompletion() is a wrapper around gtk_entry_set_completion().
func (v *Entry) SetCompletion(completion *EntryCompletion) {
	C.gtk_entry_set_completion(v.native(), completion.native())
}

// GetCompletion() is a wrapper around gtk_entry_get_completion().
func (v *Entry) GetCompletion() (*EntryCompletion, error) {
	c := C.gtk_entry_get_completion(v.native())
	if c == nil {
		return nil, nilPtrErr
	}

	e := &EntryCompletion{glib.Take(unsafe.Pointer(c))}
	return e, nil
}

// SetCursorHAdjustment() is a wrapper around
// gtk_entry_set_cursor_hadjustment().
func (v *Entry) SetCursorHAdjustment(adjustment *Adjustment) {
	C.gtk_entry_set_cursor_hadjustment(v.native(), adjustment.native())
}

// GetCursorHAdjustment() is a wrapper around
// gtk_entry_get_cursor_hadjustment().
func (v *Entry) GetCursorHAdjustment() (*Adjustment, error) {
	c := C.gtk_entry_get_cursor_hadjustment(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return &Adjustment{glib.InitiallyUnowned{obj}}, nil
}

// SetProgressFraction() is a wrapper around gtk_entry_set_progress_fraction().
func (v *Entry) SetProgressFraction(fraction float64) {
	C.gtk_entry_set_progress_fraction(v.native(), C.gdouble(fraction))
}

// GetProgressFraction() is a wrapper around gtk_entry_get_progress_fraction().
func (v *Entry) GetProgressFraction() float64 {
	c := C.gtk_entry_get_progress_fraction(v.native())
	return float64(c)
}

// SetProgressPulseStep() is a wrapper around
// gtk_entry_set_progress_pulse_step().
func (v *Entry) SetProgressPulseStep(fraction float64) {
	C.gtk_entry_set_progress_pulse_step(v.native(), C.gdouble(fraction))
}

// GetProgressPulseStep() is a wrapper around
// gtk_entry_get_progress_pulse_step().
func (v *Entry) GetProgressPulseStep() float64 {
	c := C.gtk_entry_get_progress_pulse_step(v.native())
	return float64(c)
}

// ProgressPulse() is a wrapper around gtk_entry_progress_pulse().
func (v *Entry) ProgressPulse() {
	C.gtk_entry_progress_pulse(v.native())
}

// TODO(jrick) GdkEventKey
/*
func (v *Entry) IMContextFilterKeypress() {
}
*/

// ResetIMContext() is a wrapper around gtk_entry_reset_im_context().
func (v *Entry) ResetIMContext() {
	C.gtk_entry_reset_im_context(v.native())
}

// TODO(jrick) GdkPixbuf
/*
func (v *Entry) SetIconFromPixbuf() {
}
*/

// SetIconFromIconName() is a wrapper around
// gtk_entry_set_icon_from_icon_name().
func (v *Entry) SetIconFromIconName(iconPos EntryIconPosition, name string) {
	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_entry_set_icon_from_icon_name(v.native(),
		C.GtkEntryIconPosition(iconPos), (*C.gchar)(cstr))
}

// TODO(jrick) GIcon
/*
func (v *Entry) SetIconFromGIcon() {
}
*/

// GetIconStorageType() is a wrapper around gtk_entry_get_icon_storage_type().
func (v *Entry) GetIconStorageType(iconPos EntryIconPosition) ImageType {
	c := C.gtk_entry_get_icon_storage_type(v.native(),
		C.GtkEntryIconPosition(iconPos))
	return ImageType(c)
}

// TODO(jrick) GdkPixbuf
/*
func (v *Entry) GetIconPixbuf() {
}
*/

// GetIconName() is a wrapper around gtk_entry_get_icon_name().
func (v *Entry) GetIconName(iconPos EntryIconPosition) (string, error) {
	c := C.gtk_entry_get_icon_name(v.native(),
		C.GtkEntryIconPosition(iconPos))
	if c == nil {
		return "", nilPtrErr
	}
	return goString(c), nil
}

// TODO(jrick) GIcon
/*
func (v *Entry) GetIconGIcon() {
}
*/

// SetIconActivatable() is a wrapper around gtk_entry_set_icon_activatable().
func (v *Entry) SetIconActivatable(iconPos EntryIconPosition, activatable bool) {
	C.gtk_entry_set_icon_activatable(v.native(),
		C.GtkEntryIconPosition(iconPos), gbool(activatable))
}

// GetIconActivatable() is a wrapper around gtk_entry_get_icon_activatable().
func (v *Entry) GetIconActivatable(iconPos EntryIconPosition) bool {
	c := C.gtk_entry_get_icon_activatable(v.native(),
		C.GtkEntryIconPosition(iconPos))
	return gobool(c)
}

// SetIconSensitive() is a wrapper around gtk_entry_set_icon_sensitive().
func (v *Entry) SetIconSensitive(iconPos EntryIconPosition, sensitive bool) {
	C.gtk_entry_set_icon_sensitive(v.native(),
		C.GtkEntryIconPosition(iconPos), gbool(sensitive))
}

// GetIconSensitive() is a wrapper around gtk_entry_get_icon_sensitive().
func (v *Entry) GetIconSensitive(iconPos EntryIconPosition) bool {
	c := C.gtk_entry_get_icon_sensitive(v.native(),
		C.GtkEntryIconPosition(iconPos))
	return gobool(c)
}

// GetIconAtPos() is a wrapper around gtk_entry_get_icon_at_pos().
func (v *Entry) GetIconAtPos(x, y int) int {
	c := C.gtk_entry_get_icon_at_pos(v.native(), C.gint(x), C.gint(y))
	return int(c)
}

// SetIconTooltipText() is a wrapper around gtk_entry_set_icon_tooltip_text().
func (v *Entry) SetIconTooltipText(iconPos EntryIconPosition, tooltip string) {
	var cstr *C.char
	if tooltip != "" {
		cstr = C.CString(tooltip)
		defer C.free(unsafe.Pointer(cstr))
	}
	// Null cstr is allowed by API to unset tooltip, but raise log message
	// Atk-CRITICAL **: atk_object_set_description: assertion 'ATK_IS_OBJECT (accessible)' failed
	C.gtk_entry_set_icon_tooltip_text(v.native(),
		C.GtkEntryIconPosition(iconPos), (*C.gchar)(cstr))
}

// GetIconTooltipText() is a wrapper around gtk_entry_get_icon_tooltip_text().
func (v *Entry) GetIconTooltipText(iconPos EntryIconPosition) string {
	c := C.gtk_entry_get_icon_tooltip_text(v.native(),
		C.GtkEntryIconPosition(iconPos))
	defer C.g_free(C.gpointer(c))
	return goString(c)
}

// SetIconTooltipMarkup() is a wrapper around
// gtk_entry_set_icon_tooltip_markup().
func (v *Entry) SetIconTooltipMarkup(iconPos EntryIconPosition, tooltip string) {
	cstr := C.CString(tooltip)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_entry_set_icon_tooltip_markup(v.native(),
		C.GtkEntryIconPosition(iconPos), (*C.gchar)(cstr))
}

// GetIconTooltipMarkup() is a wrapper around
// gtk_entry_get_icon_tooltip_markup().
func (v *Entry) GetIconTooltipMarkup(iconPos EntryIconPosition) string {
	c := C.gtk_entry_get_icon_tooltip_markup(v.native(),
		C.GtkEntryIconPosition(iconPos))
	defer C.g_free(C.gpointer(c))
	return goString(c)
}

// TODO(jrick) GdkDragAction
/*
func (v *Entry) SetIconDragSource() {
}
*/

// GetCurrentIconDragSource() is a wrapper around
// gtk_entry_get_current_icon_drag_source().
func (v *Entry) GetCurrentIconDragSource() int {
	c := C.gtk_entry_get_current_icon_drag_source(v.native())
	return int(c)
}

// TODO(jrick) GdkRectangle
/*
func (v *Entry) GetIconArea() {
}
*/

// SetInputPurpose() is a wrapper around gtk_entry_set_input_purpose().
func (v *Entry) SetInputPurpose(purpose InputPurpose) {
	C.gtk_entry_set_input_purpose(v.native(), C.GtkInputPurpose(purpose))
}

// GetInputPurpose() is a wrapper around gtk_entry_get_input_purpose().
func (v *Entry) GetInputPurpose() InputPurpose {
	c := C.gtk_entry_get_input_purpose(v.native())
	return InputPurpose(c)
}

// SetInputHints() is a wrapper around gtk_entry_set_input_hints().
func (v *Entry) SetInputHints(hints InputHints) {
	C.gtk_entry_set_input_hints(v.native(), C.GtkInputHints(hints))
}

// GetInputHints() is a wrapper around gtk_entry_get_input_hints().
func (v *Entry) GetInputHints() InputHints {
	c := C.gtk_entry_get_input_hints(v.native())
	return InputHints(c)
}

/*
 * GtkEntryBuffer
 */

// EntryBuffer is a representation of GTK's GtkEntryBuffer.
type EntryBuffer struct {
	*glib.Object
}

// native returns a pointer to the underlying GtkEntryBuffer.
func (v *EntryBuffer) native() *C.GtkEntryBuffer {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toGtkEntryBuffer(ptr)
}

func marshalEntryBuffer(p uintptr) (interface{}, error) {
	c := C.g_value_get_object(C.toGValue(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapEntryBuffer(obj), nil
}

func wrapEntryBuffer(obj *glib.Object) *EntryBuffer {
	return &EntryBuffer{obj}
}

// EntryBufferNew() is a wrapper around gtk_entry_buffer_new().
func EntryBufferNew(initialChars string, nInitialChars int) (*EntryBuffer, error) {
	cstr := C.CString(initialChars)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_entry_buffer_new((*C.gchar)(cstr), C.gint(nInitialChars))
	if c == nil {
		return nil, nilPtrErr
	}

	obj := glib.Take(unsafe.Pointer(c))
	return wrapEntryBuffer(obj), nil
}

// GetText() is a wrapper around gtk_entry_buffer_get_text().  A
// non-nil error is returned in the case that gtk_entry_buffer_get_text
// returns NULL to differentiate between NULL and an empty string.
func (v *EntryBuffer) GetText() (string, error) {
	c := C.gtk_entry_buffer_get_text(v.native())
	if c == nil {
		return "", nilPtrErr
	}
	return goString(c), nil
}

// SetText() is a wrapper around gtk_entry_buffer_set_text().
func (v *EntryBuffer) SetText(text string) {
	cstr := C.CString(text)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_entry_buffer_set_text(v.native(), (*C.gchar)(cstr),
		C.gint(len(text)))
}

// GetBytes() is a wrapper around gtk_entry_buffer_get_bytes().
func (v *EntryBuffer) GetBytes() uint {
	c := C.gtk_entry_buffer_get_bytes(v.native())
	return uint(c)
}

// GetLength() is a wrapper around gtk_entry_buffer_get_length().
func (v *EntryBuffer) GetLength() uint {
	c := C.gtk_entry_buffer_get_length(v.native())
	return uint(c)
}

// GetMaxLength() is a wrapper around gtk_entry_buffer_get_max_length().
func (v *EntryBuffer) GetMaxLength() int {
	c := C.gtk_entry_buffer_get_max_length(v.native())
	return int(c)
}

// SetMaxLength() is a wrapper around gtk_entry_buffer_set_max_length().
func (v *EntryBuffer) SetMaxLength(maxLength int) {
	C.gtk_entry_buffer_set_max_length(v.native(), C.gint(maxLength))
}

// InsertText() is a wrapper around gtk_entry_buffer_insert_text().
func (v *EntryBuffer) InsertText(position uint, text string) uint {
	cstr := C.CString(text)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_entry_buffer_insert_text(v.native(), C.guint(position),
		(*C.gchar)(cstr), C.gint(len(text)))
	return uint(c)
}

// DeleteText() is a wrapper around gtk_entry_buffer_delete_text().
func (v *EntryBuffer) DeleteText(position uint, nChars int) uint {
	c := C.gtk_entry_buffer_delete_text(v.native(), C.guint(position),
		C.gint(nChars))
	return uint(c)
}

// EmitDeletedText() is a wrapper around gtk_entry_buffer_emit_deleted_text().
func (v *EntryBuffer) EmitDeletedText(pos, nChars uint) {
	C.gtk_entry_buffer_emit_deleted_text(v.native(), C.guint(pos),
		C.guint(nChars))
}

// EmitInsertedText() is a wrapper around gtk_entry_buffer_emit_inserted_text().
func (v *EntryBuffer) EmitInsertedText(pos uint, text string) {
	cstr := C.CString(text)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_entry_buffer_emit_inserted_text(v.native(), C.guint(pos),
		(*C.gchar)(cstr), C.guint(len(text)))
}

/*
 * GtkEntryCompletion
 */

// EntryCompletion is a representation of GTK's GtkEntryCompletion.
type EntryCompletion struct {
	*glib.Object
}

// native returns a pointer to the underlying GtkEntryCompletion.
func (v *EntryCompletion) native() *C.GtkEntryCompletion {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toGtkEntryCompletion(ptr)
}

func marshalEntryCompletion(p uintptr) (interface{}, error) {
	c := C.g_value_get_object(C.toGValue(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapEntryCompletion(obj), nil
}

func wrapEntryCompletion(obj *glib.Object) *EntryCompletion {
	return &EntryCompletion{obj}
}

/*
 * GtkScale
 */

// Scale is a representation of GTK's GtkScale.
type Scale struct {
	Range
}

// native returns a pointer to the underlying GtkScale.
func (v *Scale) native() *C.GtkScale {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toGtkScale(ptr)
}

func marshalScale(p uintptr) (interface{}, error) {
	c := C.g_value_get_object(C.toGValue(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapScale(obj), nil
}

func wrapScale(obj *glib.Object) *Scale {
	rng := wrapRange(obj)
	return &Scale{*rng}
}

// ScaleNew is a wrapper around gtk_scale_new().
func ScaleNew(orientation Orientation, adjustment *Adjustment) (*Scale, error) {
	c := C.gtk_scale_new(C.GtkOrientation(orientation), adjustment.native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapScale(obj), nil
}

// ScaleNewWithRange is a wrapper around gtk_scale_new_with_range().
func ScaleNewWithRange(orientation Orientation, min, max, step float64) (*Scale, error) {
	c := C.gtk_scale_new_with_range(C.GtkOrientation(orientation),
		C.gdouble(min), C.gdouble(max), C.gdouble(step))

	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapScale(obj), nil
}

// SetDrawValue() is a wrapper around gtk_scale_set_draw_value().
func (v *Scale) SetDrawValue(drawValue bool) {
	C.gtk_scale_set_draw_value(v.native(), gbool(drawValue))
}

/*
 * GtkSpinButton
 */

// SpinButton is a representation of GTK's GtkSpinButton.
type SpinButton struct {
	Entry
}

// native returns a pointer to the underlying GtkSpinButton.
func (v *SpinButton) native() *C.GtkSpinButton {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toGtkSpinButton(ptr)
}

func marshalSpinButton(p uintptr) (interface{}, error) {
	c := C.g_value_get_object(C.toGValue(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapSpinButton(obj), nil
}

func wrapSpinButton(obj *glib.Object) *SpinButton {
	entry := wrapEntry(obj)
	return &SpinButton{*entry}
}

// Configure() is a wrapper around gtk_spin_button_configure().
func (v *SpinButton) Configure(adjustment *Adjustment, climbRate float64, digits uint) {
	C.gtk_spin_button_configure(v.native(), adjustment.native(),
		C.gdouble(climbRate), C.guint(digits))
}

// SpinButtonNew() is a wrapper around gtk_spin_button_new().
func SpinButtonNew(adjustment *Adjustment, climbRate float64, digits uint) (*SpinButton, error) {
	c := C.gtk_spin_button_new(adjustment.native(),
		C.gdouble(climbRate), C.guint(digits))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapSpinButton(obj), nil
}

// SpinButtonNewWithRange() is a wrapper around
// gtk_spin_button_new_with_range().
func SpinButtonNewWithRange(min, max, step float64) (*SpinButton, error) {
	c := C.gtk_spin_button_new_with_range(C.gdouble(min), C.gdouble(max),
		C.gdouble(step))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapSpinButton(obj), nil
}

// GetValueAsInt() is a wrapper around gtk_spin_button_get_value_as_int().
func (v *SpinButton) GetValueAsInt() int {
	c := C.gtk_spin_button_get_value_as_int(v.native())
	return int(c)
}

// SetValue() is a wrapper around gtk_spin_button_set_value().
func (v *SpinButton) SetValue(value float64) {
	C.gtk_spin_button_set_value(v.native(), C.gdouble(value))
}

// GetValue() is a wrapper around gtk_spin_button_get_value().
func (v *SpinButton) GetValue() float64 {
	c := C.gtk_spin_button_get_value(v.native())
	return float64(c)
}

// GetAdjustment() is a wrapper around gtk_spin_button_get_adjustment
func (v *SpinButton) GetAdjustment() *Adjustment {
	c := C.gtk_spin_button_get_adjustment(v.native())
	if c == nil {
		return nil
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapAdjustment(obj)
}

// SetRange is a wrapper around gtk_spin_button_set_range().
func (v *SpinButton) SetRange(min, max float64) {
	C.gtk_spin_button_set_range(v.native(), C.gdouble(min), C.gdouble(max))
}

// SetIncrements() is a wrapper around gtk_spin_button_set_increments().
func (v *SpinButton) SetIncrements(step, page float64) {
	C.gtk_spin_button_set_increments(v.native(), C.gdouble(step), C.gdouble(page))
}

/*
 * GtkSearchEntry
 */

// SearchEntry is a reprensentation of GTK's GtkSearchEntry.
type SearchEntry struct {
	Entry
}

// native returns a pointer to the underlying GtkSearchEntry.
func (v *SearchEntry) native() *C.GtkSearchEntry {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toGtkSearchEntry(ptr)
}

func marshalSearchEntry(p uintptr) (interface{}, error) {
	c := C.g_value_get_object(C.toGValue(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapSearchEntry(obj), nil
}

func wrapSearchEntry(obj *glib.Object) *SearchEntry {
	entry := wrapEntry(obj)
	return &SearchEntry{*entry}
}

// SearchEntryNew is a wrapper around gtk_search_entry_new().
func SearchEntryNew() (*SearchEntry, error) {
	c := C.gtk_search_entry_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapSearchEntry(obj), nil
}

// TODO: implement GtkSearchBar
