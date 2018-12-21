// Same copyright and license as the rest of the files in this project

package gtk

// #include <gtk/gtk.h>
// #include "gtk.go.h"
import "C"

import (
	"unsafe"

	"github.com/d2r2/gotk3/gdk"
	"github.com/d2r2/gotk3/glib"
)

/*
 * GtkTextIter
 */

// TextIter is a representation of GTK's GtkTextIter
type TextIter struct {
	gtkTextIter *C.GtkTextIter
}

// native returns a pointer to the underlying GtkTextIter.
func (v *TextIter) native() *C.GtkTextIter {
	if v == nil {
		return nil
	}
	return v.gtkTextIter
}

func marshalTextIter(p uintptr) (interface{}, error) {
	c := C.g_value_get_boxed(C.toGValue(unsafe.Pointer(p)))
	c2 := (*C.GtkTreePath)(unsafe.Pointer(c))
	return wrapTreePath(c2), nil
}

func wrapTextIter(obj *C.GtkTextIter) *TextIter {
	return &TextIter{obj}
}

// GetBuffer is a wrapper around gtk_text_iter_get_buffer().
func (v *TextIter) GetBuffer() *TextBuffer {
	c := C.gtk_text_iter_get_buffer(v.native())
	if c == nil {
		return nil
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapTextBuffer(obj)
}

// GetOffset is a wrapper around gtk_text_iter_get_offset().
func (v *TextIter) GetOffset() int {
	return int(C.gtk_text_iter_get_offset(v.native()))
}

// GetLine is a wrapper around gtk_text_iter_get_line().
func (v *TextIter) GetLine() int {
	return int(C.gtk_text_iter_get_line(v.native()))
}

// GetLineOffset is a wrapper around gtk_text_iter_get_line_offset().
func (v *TextIter) GetLineOffset() int {
	return int(C.gtk_text_iter_get_line_offset(v.native()))
}

// GetLineIndex is a wrapper around gtk_text_iter_get_line_index().
func (v *TextIter) GetLineIndex() int {
	return int(C.gtk_text_iter_get_line_index(v.native()))
}

// GetVisibleLineOffset is a wrapper around gtk_text_iter_get_visible_line_offset().
func (v *TextIter) GetVisibleLineOffset() int {
	return int(C.gtk_text_iter_get_visible_line_offset(v.native()))
}

// GetVisibleLineIndex is a wrapper around gtk_text_iter_get_visible_line_index().
func (v *TextIter) GetVisibleLineIndex() int {
	return int(C.gtk_text_iter_get_visible_line_index(v.native()))
}

// GetChar is a wrapper around gtk_text_iter_get_char().
func (v *TextIter) GetChar() rune {
	return rune(C.gtk_text_iter_get_char(v.native()))
}

// GetSlice is a wrapper around gtk_text_iter_get_slice().
func (v *TextIter) GetSlice(end *TextIter) string {
	c := C.gtk_text_iter_get_slice(v.native(), end.native())
	return goString(c)
}

// GetText is a wrapper around gtk_text_iter_get_text().
func (v *TextIter) GetText(end *TextIter) string {
	c := C.gtk_text_iter_get_text(v.native(), end.native())
	return goString(c)
}

// GetVisibleSlice is a wrapper around gtk_text_iter_get_visible_slice().
func (v *TextIter) GetVisibleSlice(end *TextIter) string {
	c := C.gtk_text_iter_get_visible_slice(v.native(), end.native())
	return goString(c)
}

// GetVisibleText is a wrapper around gtk_text_iter_get_visible_text().
func (v *TextIter) GetVisibleText(end *TextIter) string {
	c := C.gtk_text_iter_get_visible_text(v.native(), end.native())
	return goString(c)
}

// EndsTag is a wrapper around gtk_text_iter_ends_tag().
func (v *TextIter) EndsTag(v1 *TextTag) bool {
	return gobool(C.gtk_text_iter_ends_tag(v.native(), v1.native()))
}

// TogglesTag is a wrapper around gtk_text_iter_toggles_tag().
func (v *TextIter) TogglesTag(v1 *TextTag) bool {
	return gobool(C.gtk_text_iter_toggles_tag(v.native(), v1.native()))
}

// HasTag is a wrapper around gtk_text_iter_has_tag().
func (v *TextIter) HasTag(v1 *TextTag) bool {
	return gobool(C.gtk_text_iter_has_tag(v.native(), v1.native()))
}

// Editable is a wrapper around gtk_text_iter_editable().
func (v *TextIter) Editable(v1 bool) bool {
	return gobool(C.gtk_text_iter_editable(v.native(), gbool(v1)))
}

// CanInsert is a wrapper around gtk_text_iter_can_insert().
func (v *TextIter) CanInsert(v1 bool) bool {
	return gobool(C.gtk_text_iter_can_insert(v.native(), gbool(v1)))
}

// StartsWord is a wrapper around gtk_text_iter_starts_word().
func (v *TextIter) StartsWord() bool {
	return gobool(C.gtk_text_iter_starts_word(v.native()))
}

// EndsWord is a wrapper around gtk_text_iter_ends_word().
func (v *TextIter) EndsWord() bool {
	return gobool(C.gtk_text_iter_ends_word(v.native()))
}

// InsideWord is a wrapper around gtk_text_iter_inside_word().
func (v *TextIter) InsideWord() bool {
	return gobool(C.gtk_text_iter_inside_word(v.native()))
}

// StartsLine is a wrapper around gtk_text_iter_starts_line().
func (v *TextIter) StartsLine() bool {
	return gobool(C.gtk_text_iter_starts_line(v.native()))
}

// EndsLine is a wrapper around gtk_text_iter_ends_line().
func (v *TextIter) EndsLine() bool {
	return gobool(C.gtk_text_iter_ends_line(v.native()))
}

// StartsSentence is a wrapper around gtk_text_iter_starts_sentence().
func (v *TextIter) StartsSentence() bool {
	return gobool(C.gtk_text_iter_starts_sentence(v.native()))
}

// EndsSentence is a wrapper around gtk_text_iter_ends_sentence().
func (v *TextIter) EndsSentence() bool {
	return gobool(C.gtk_text_iter_ends_sentence(v.native()))
}

// InsideSentence is a wrapper around gtk_text_iter_inside_sentence().
func (v *TextIter) InsideSentence() bool {
	return gobool(C.gtk_text_iter_inside_sentence(v.native()))
}

// IsCursorPosition is a wrapper around gtk_text_iter_is_cursor_position().
func (v *TextIter) IsCursorPosition() bool {
	return gobool(C.gtk_text_iter_is_cursor_position(v.native()))
}

// GetCharsInLine is a wrapper around gtk_text_iter_get_chars_in_line().
func (v *TextIter) GetCharsInLine() int {
	return int(C.gtk_text_iter_get_chars_in_line(v.native()))
}

// GetBytesInLine is a wrapper around gtk_text_iter_get_bytes_in_line().
func (v *TextIter) GetBytesInLine() int {
	return int(C.gtk_text_iter_get_bytes_in_line(v.native()))
}

// IsEnd is a wrapper around gtk_text_iter_is_end().
func (v *TextIter) IsEnd() bool {
	return gobool(C.gtk_text_iter_is_end(v.native()))
}

// IsStart is a wrapper around gtk_text_iter_is_start().
func (v *TextIter) IsStart() bool {
	return gobool(C.gtk_text_iter_is_start(v.native()))
}

// ForwardChar is a wrapper around gtk_text_iter_forward_char().
func (v *TextIter) ForwardChar() bool {
	return gobool(C.gtk_text_iter_forward_char(v.native()))
}

// BackwardChar is a wrapper around gtk_text_iter_backward_char().
func (v *TextIter) BackwardChar() bool {
	return gobool(C.gtk_text_iter_backward_char(v.native()))
}

// ForwardChars is a wrapper around gtk_text_iter_forward_chars().
func (v *TextIter) ForwardChars(v1 int) bool {
	return gobool(C.gtk_text_iter_forward_chars(v.native(), C.gint(v1)))
}

// BackwardChars is a wrapper around gtk_text_iter_backward_chars().
func (v *TextIter) BackwardChars(v1 int) bool {
	return gobool(C.gtk_text_iter_backward_chars(v.native(), C.gint(v1)))
}

// ForwardLine is a wrapper around gtk_text_iter_forward_line().
func (v *TextIter) ForwardLine() bool {
	return gobool(C.gtk_text_iter_forward_line(v.native()))
}

// BackwardLine is a wrapper around gtk_text_iter_backward_line().
func (v *TextIter) BackwardLine() bool {
	return gobool(C.gtk_text_iter_backward_line(v.native()))
}

// ForwardLines is a wrapper around gtk_text_iter_forward_lines().
func (v *TextIter) ForwardLines(v1 int) bool {
	return gobool(C.gtk_text_iter_forward_lines(v.native(), C.gint(v1)))
}

// BackwardLines is a wrapper around gtk_text_iter_backward_lines().
func (v *TextIter) BackwardLines(v1 int) bool {
	return gobool(C.gtk_text_iter_backward_lines(v.native(), C.gint(v1)))
}

// ForwardWordEnds is a wrapper around gtk_text_iter_forward_word_ends().
func (v *TextIter) ForwardWordEnds(v1 int) bool {
	return gobool(C.gtk_text_iter_forward_word_ends(v.native(), C.gint(v1)))
}

// ForwardWordEnd is a wrapper around gtk_text_iter_forward_word_end().
func (v *TextIter) ForwardWordEnd() bool {
	return gobool(C.gtk_text_iter_forward_word_end(v.native()))
}

// ForwardCursorPosition is a wrapper around gtk_text_iter_forward_cursor_position().
func (v *TextIter) ForwardCursorPosition() bool {
	return gobool(C.gtk_text_iter_forward_cursor_position(v.native()))
}

// BackwardCursorPosition is a wrapper around gtk_text_iter_backward_cursor_position().
func (v *TextIter) BackwardCursorPosition() bool {
	return gobool(C.gtk_text_iter_backward_cursor_position(v.native()))
}

// ForwardCursorPositions is a wrapper around gtk_text_iter_forward_cursor_positions().
func (v *TextIter) ForwardCursorPositions(v1 int) bool {
	return gobool(C.gtk_text_iter_forward_cursor_positions(v.native(), C.gint(v1)))
}

// BackwardCursorPositions is a wrapper around gtk_text_iter_backward_cursor_positions().
func (v *TextIter) BackwardCursorPositions(v1 int) bool {
	return gobool(C.gtk_text_iter_backward_cursor_positions(v.native(), C.gint(v1)))
}

// ForwardSentenceEnds is a wrapper around gtk_text_iter_forward_sentence_ends().
func (v *TextIter) ForwardSentenceEnds(v1 int) bool {
	return gobool(C.gtk_text_iter_forward_sentence_ends(v.native(), C.gint(v1)))
}

// ForwardSentenceEnd is a wrapper around gtk_text_iter_forward_sentence_end().
func (v *TextIter) ForwardSentenceEnd() bool {
	return gobool(C.gtk_text_iter_forward_sentence_end(v.native()))
}

// ForwardVisibleWordEnds is a wrapper around gtk_text_iter_forward_word_ends().
func (v *TextIter) ForwardVisibleWordEnds(v1 int) bool {
	return gobool(C.gtk_text_iter_forward_word_ends(v.native(), C.gint(v1)))
}

// ForwardVisibleWordEnd is a wrapper around gtk_text_iter_forward_visible_word_end().
func (v *TextIter) ForwardVisibleWordEnd() bool {
	return gobool(C.gtk_text_iter_forward_visible_word_end(v.native()))
}

// ForwardVisibleCursorPosition is a wrapper around gtk_text_iter_forward_visible_cursor_position().
func (v *TextIter) ForwardVisibleCursorPosition() bool {
	return gobool(C.gtk_text_iter_forward_visible_cursor_position(v.native()))
}

// BackwardVisibleCursorPosition is a wrapper around gtk_text_iter_backward_visible_cursor_position().
func (v *TextIter) BackwardVisibleCursorPosition() bool {
	return gobool(C.gtk_text_iter_backward_visible_cursor_position(v.native()))
}

// ForwardVisibleCursorPositions is a wrapper around gtk_text_iter_forward_visible_cursor_positions().
func (v *TextIter) ForwardVisibleCursorPositions(v1 int) bool {
	return gobool(C.gtk_text_iter_forward_visible_cursor_positions(v.native(), C.gint(v1)))
}

// BackwardVisibleCursorPositions is a wrapper around gtk_text_iter_backward_visible_cursor_positions().
func (v *TextIter) BackwardVisibleCursorPositions(v1 int) bool {
	return gobool(C.gtk_text_iter_backward_visible_cursor_positions(v.native(), C.gint(v1)))
}

// ForwardVisibleLine is a wrapper around gtk_text_iter_forward_visible_line().
func (v *TextIter) ForwardVisibleLine() bool {
	return gobool(C.gtk_text_iter_forward_visible_line(v.native()))
}

// BackwardVisibleLine is a wrapper around gtk_text_iter_backward_visible_line().
func (v *TextIter) BackwardVisibleLine() bool {
	return gobool(C.gtk_text_iter_backward_visible_line(v.native()))
}

// ForwardVisibleLines is a wrapper around gtk_text_iter_forward_visible_lines().
func (v *TextIter) ForwardVisibleLines(v1 int) bool {
	return gobool(C.gtk_text_iter_forward_visible_lines(v.native(), C.gint(v1)))
}

// BackwardVisibleLines is a wrapper around gtk_text_iter_backward_visible_lines().
func (v *TextIter) BackwardVisibleLines(v1 int) bool {
	return gobool(C.gtk_text_iter_backward_visible_lines(v.native(), C.gint(v1)))
}

// SetOffset is a wrapper around gtk_text_iter_set_offset().
func (v *TextIter) SetOffset(v1 int) {
	C.gtk_text_iter_set_offset(v.native(), C.gint(v1))
}

// SetLine is a wrapper around gtk_text_iter_set_line().
func (v *TextIter) SetLine(v1 int) {
	C.gtk_text_iter_set_line(v.native(), C.gint(v1))
}

// SetLineOffset is a wrapper around gtk_text_iter_set_line_offset().
func (v *TextIter) SetLineOffset(v1 int) {
	C.gtk_text_iter_set_line_offset(v.native(), C.gint(v1))
}

// SetLineIndex is a wrapper around gtk_text_iter_set_line_index().
func (v *TextIter) SetLineIndex(v1 int) {
	C.gtk_text_iter_set_line_index(v.native(), C.gint(v1))
}

// SetVisibleLineOffset is a wrapper around gtk_text_iter_set_visible_line_offset().
func (v *TextIter) SetVisibleLineOffset(v1 int) {
	C.gtk_text_iter_set_visible_line_offset(v.native(), C.gint(v1))
}

// SetVisibleLineIndex is a wrapper around gtk_text_iter_set_visible_line_index().
func (v *TextIter) SetVisibleLineIndex(v1 int) {
	C.gtk_text_iter_set_visible_line_index(v.native(), C.gint(v1))
}

// ForwardToEnd is a wrapper around gtk_text_iter_forward_to_end().
func (v *TextIter) ForwardToEnd() {
	C.gtk_text_iter_forward_to_end(v.native())
}

// ForwardToLineEnd is a wrapper around gtk_text_iter_forward_to_line_end().
func (v *TextIter) ForwardToLineEnd() bool {
	return gobool(C.gtk_text_iter_forward_to_line_end(v.native()))
}

// ForwardToTagToggle is a wrapper around gtk_text_iter_forward_to_tag_toggle().
func (v *TextIter) ForwardToTagToggle(v1 *TextTag) bool {
	return gobool(C.gtk_text_iter_forward_to_tag_toggle(v.native(), v1.native()))
}

// BackwardToTagToggle is a wrapper around gtk_text_iter_backward_to_tag_toggle().
func (v *TextIter) BackwardToTagToggle(v1 *TextTag) bool {
	return gobool(C.gtk_text_iter_backward_to_tag_toggle(v.native(), v1.native()))
}

// Equal is a wrapper around gtk_text_iter_equal().
func (v *TextIter) Equal(v1 *TextIter) bool {
	return gobool(C.gtk_text_iter_equal(v.native(), v1.native()))
}

// Compare is a wrapper around gtk_text_iter_compare().
func (v *TextIter) Compare(v1 *TextIter) int {
	return int(C.gtk_text_iter_compare(v.native(), v1.native()))
}

// InRange is a wrapper around gtk_text_iter_in_range().
func (v *TextIter) InRange(v1 *TextIter, v2 *TextIter) bool {
	return gobool(C.gtk_text_iter_in_range(v.native(), v1.native(), v2.native()))
}

// void 	gtk_text_iter_order ()
// gboolean 	(*GtkTextCharPredicate) ()
// gboolean 	gtk_text_iter_forward_find_char ()
// gboolean 	gtk_text_iter_backward_find_char ()
// gboolean 	gtk_text_iter_forward_search ()
// gboolean 	gtk_text_iter_backward_search ()
// gboolean 	gtk_text_iter_get_attributes ()
// GtkTextIter * 	gtk_text_iter_copy ()
// void 	gtk_text_iter_assign ()
// void 	gtk_text_iter_free ()
// GdkPixbuf * 	gtk_text_iter_get_pixbuf ()
// GSList * 	gtk_text_iter_get_marks ()
// GSList * 	gtk_text_iter_get_toggled_tags ()
// GtkTextChildAnchor * 	gtk_text_iter_get_child_anchor ()
// GSList * 	gtk_text_iter_get_tags ()
// PangoLanguage * 	gtk_text_iter_get_language ()

/*
 * GtkTextMark
 */

// TextMark is a representation of GTK's GtkTextMark
type TextMark struct {
	gtkTextMark *C.GtkTextMark
}

// native returns a pointer to the underlying GtkTextMark.
func (v *TextMark) native() *C.GtkTextMark {
	if v == nil {
		return nil
	}
	return v.gtkTextMark
}

func marshalTextMark(p uintptr) (interface{}, error) {
	c := C.g_value_get_boxed(C.toGValue(unsafe.Pointer(p)))
	c2 := (*C.GtkTextMark)(unsafe.Pointer(c))
	return wrapTextMark(c2), nil
}

func wrapTextMark(obj *C.GtkTextMark) *TextMark {
	return &TextMark{obj}
}

// TextWindowType is a representation of GTK's GtkTextWindowType.
type TextWindowType int

const (
	TEXT_WINDOW_WIDGET TextWindowType = C.GTK_TEXT_WINDOW_WIDGET
	TEXT_WINDOW_TEXT   TextWindowType = C.GTK_TEXT_WINDOW_TEXT
	TEXT_WINDOW_LEFT   TextWindowType = C.GTK_TEXT_WINDOW_LEFT
	TEXT_WINDOW_RIGHT  TextWindowType = C.GTK_TEXT_WINDOW_RIGHT
	TEXT_WINDOW_TOP    TextWindowType = C.GTK_TEXT_WINDOW_TOP
	TEXT_WINDOW_BOTTOM TextWindowType = C.GTK_TEXT_WINDOW_BOTTOM
)

/*
 * GtkTextTag
 */

type TextTag struct {
	*glib.Object
}

// native returns a pointer to the underlying GObject as a GtkTextTag.
func (v *TextTag) native() *C.GtkTextTag {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toGtkTextTag(ptr)
}

func marshalTextTag(p uintptr) (interface{}, error) {
	c := C.g_value_get_object(C.toGValue(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapTextTag(obj), nil
}

func wrapTextTag(obj *glib.Object) *TextTag {
	return &TextTag{obj}
}

func TextTagNew(name string) (*TextTag, error) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	c := C.gtk_text_tag_new((*C.gchar)(cname))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapTextTag(obj), nil
}

// GetPriority() is a wrapper around gtk_text_tag_get_priority().
func (v *TextTag) GetPriority() int {
	return int(C.gtk_text_tag_get_priority(v.native()))
}

// SetPriority() is a wrapper around gtk_text_tag_set_priority().
func (v *TextTag) SetPriority(priority int) {
	C.gtk_text_tag_set_priority(v.native(), C.gint(priority))
}

// Event() is a wrapper around gtk_text_tag_event().
func (v *TextTag) Event(eventObject *glib.Object, event *gdk.Event, iter *TextIter) bool {
	ok := C.gtk_text_tag_event(v.native(),
		C.toGObject(unsafe.Pointer(eventObject.Native())),
		(*C.GdkEvent)(unsafe.Pointer(event.Native())),
		iter.native(),
	)
	return gobool(ok)
}

/*
 * GtkTextTagTable
 */

type TextTagTable struct {
	*glib.Object
}

// native returns a pointer to the underlying GObject as a GtkTextTagTable.
func (v *TextTagTable) native() *C.GtkTextTagTable {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toGtkTextTagTable(ptr)
}

func marshalTextTagTable(p uintptr) (interface{}, error) {
	c := C.g_value_get_object(C.toGValue(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapTextTagTable(obj), nil
}

func wrapTextTagTable(obj *glib.Object) *TextTagTable {
	return &TextTagTable{obj}
}

func TextTagTableNew() (*TextTagTable, error) {
	c := C.gtk_text_tag_table_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapTextTagTable(obj), nil
}

// Add() is a wrapper around gtk_text_tag_table_add().
func (v *TextTagTable) Add(tag *TextTag) {
	C.gtk_text_tag_table_add(v.native(), tag.native())
	//return gobool(c) // TODO version-separate
}

// Lookup() is a wrapper around gtk_text_tag_table_lookup().
func (v *TextTagTable) Lookup(name string) (*TextTag, error) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	c := C.gtk_text_tag_table_lookup(v.native(), (*C.gchar)(cname))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapTextTag(obj), nil
}

// Remove() is a wrapper around gtk_text_tag_table_remove().
func (v *TextTagTable) Remove(tag *TextTag) {
	C.gtk_text_tag_table_remove(v.native(), tag.native())
}

/*
 * GtkTextBuffer
 */

// TextBuffer is a representation of GTK's GtkTextBuffer.
type TextBuffer struct {
	*glib.Object
}

// native returns a pointer to the underlying GtkTextBuffer.
func (v *TextBuffer) native() *C.GtkTextBuffer {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toGtkTextBuffer(ptr)
}

func marshalTextBuffer(p uintptr) (interface{}, error) {
	c := C.g_value_get_object(C.toGValue(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapTextBuffer(obj), nil
}

func wrapTextBuffer(obj *glib.Object) *TextBuffer {
	return &TextBuffer{obj}
}

// TextBufferNew() is a wrapper around gtk_text_buffer_new().
func TextBufferNew(table *TextTagTable) (*TextBuffer, error) {
	c := C.gtk_text_buffer_new(table.native())
	if c == nil {
		return nil, nilPtrErr
	}

	obj := glib.Take(unsafe.Pointer(c))
	return wrapTextBuffer(obj), nil
}

// ApplyTag() is a wrapper around gtk_text_buffer_apply_tag().
func (v *TextBuffer) ApplyTag(tag *TextTag, start, end *TextIter) {
	C.gtk_text_buffer_apply_tag(v.native(), tag.native(), start.native(), end.native())
}

// ApplyTagByName() is a wrapper around gtk_text_buffer_apply_tag_by_name().
func (v *TextBuffer) ApplyTagByName(name string, start, end *TextIter) {
	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_text_buffer_apply_tag_by_name(v.native(), (*C.gchar)(cstr),
		start.native(), end.native())
}

// Delete() is a wrapper around gtk_text_buffer_delete().
func (v *TextBuffer) Delete(start, end *TextIter) {
	C.gtk_text_buffer_delete(v.native(), start.native(), end.native())
}

func (v *TextBuffer) GetBounds() (start, end *TextIter) {
	var tis C.GtkTextIter
	var tie C.GtkTextIter
	C.gtk_text_buffer_get_bounds(v.native(), &tis, &tie)
	start = wrapTextIter(&tis)
	end = wrapTextIter(&tie)
	return start, end
}

// GetCharCount() is a wrapper around gtk_text_buffer_get_char_count().
func (v *TextBuffer) GetCharCount() int {
	return int(C.gtk_text_buffer_get_char_count(v.native()))
}

// GetIterAtOffset() is a wrapper around gtk_text_buffer_get_iter_at_offset().
func (v *TextBuffer) GetIterAtOffset(charOffset int) *TextIter {
	var iter C.GtkTextIter
	C.gtk_text_buffer_get_iter_at_offset(v.native(), &iter, C.gint(charOffset))
	ti := wrapTextIter(&iter)
	return ti
}

// GetStartIter() is a wrapper around gtk_text_buffer_get_start_iter().
func (v *TextBuffer) GetStartIter() *TextIter {
	var iter C.GtkTextIter
	C.gtk_text_buffer_get_start_iter(v.native(), &iter)
	ti := wrapTextIter(&iter)
	return ti
}

// GetEndIter() is a wrapper around gtk_text_buffer_get_end_iter().
func (v *TextBuffer) GetEndIter() *TextIter {
	var iter C.GtkTextIter
	C.gtk_text_buffer_get_end_iter(v.native(), &iter)
	ti := wrapTextIter(&iter)
	return ti
}

// GetLineCount() is a wrapper around gtk_text_buffer_get_line_count().
func (v *TextBuffer) GetLineCount() int {
	return int(C.gtk_text_buffer_get_line_count(v.native()))
}

// GetModified() is a wrapper around gtk_text_buffer_get_modified().
func (v *TextBuffer) GetModified() bool {
	return gobool(C.gtk_text_buffer_get_modified(v.native()))
}

// GetTagTable() is a wrapper around gtk_text_buffer_get_tag_table().
func (v *TextBuffer) GetTagTable() (*TextTagTable, error) {
	c := C.gtk_text_buffer_get_tag_table(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapTextTagTable(obj), nil
}

func (v *TextBuffer) GetText(start, end *TextIter, includeHiddenChars bool) (string, error) {
	c := C.gtk_text_buffer_get_text(v.native(), start.native(), end.native(),
		gbool(includeHiddenChars))
	if c == nil {
		return "", nilPtrErr
	}
	gostr := goString(c)
	C.free(unsafe.Pointer(c))
	return gostr, nil
}

// Insert() is a wrapper around gtk_text_buffer_insert().
func (v *TextBuffer) Insert(iter *TextIter, text string) {
	cstr := C.CString(text)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_text_buffer_insert(v.native(), iter.native(), (*C.gchar)(cstr), C.gint(len(text)))
}

// InsertAtCursor() is a wrapper around gtk_text_buffer_insert_at_cursor().
func (v *TextBuffer) InsertAtCursor(text string) {
	cstr := C.CString(text)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_text_buffer_insert_at_cursor(v.native(), (*C.gchar)(cstr), C.gint(len(text)))
}

// RemoveTag() is a wrapper around gtk_text_buffer_remove_tag().
func (v *TextBuffer) RemoveTag(tag *TextTag, start, end *TextIter) {
	C.gtk_text_buffer_remove_tag(v.native(), tag.native(), start.native(), end.native())
}

// SetModified() is a wrapper around gtk_text_buffer_set_modified().
func (v *TextBuffer) SetModified(setting bool) {
	C.gtk_text_buffer_set_modified(v.native(), gbool(setting))
}

func (v *TextBuffer) SetText(text string) {
	cstr := C.CString(text)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_text_buffer_set_text(v.native(), (*C.gchar)(cstr),
		C.gint(len(text)))
}

// GetIterAtMark() is a wrapper around gtk_text_buffer_get_iter_at_mark().
func (v *TextBuffer) GetIterAtMark(mark *TextMark) *TextIter {
	var iter C.GtkTextIter
	C.gtk_text_buffer_get_iter_at_mark(v.native(), &iter, mark.native())
	ti := wrapTextIter(&iter)
	return ti
}

// CreateMark() is a wrapper around gtk_text_buffer_create_mark().
func (v *TextBuffer) CreateMark(mark_name string, where *TextIter, left_gravity bool) *TextMark {
	cstr := C.CString(mark_name)
	defer C.free(unsafe.Pointer(cstr))
	ret := C.gtk_text_buffer_create_mark(v.native(), (*C.gchar)(cstr), where.native(), gbool(left_gravity))
	tm := wrapTextMark(ret)
	return tm
}

// InsertChildAnchor() is a wrapper around gtk_text_buffer_insert_child_anchor().
func (v *TextBuffer) InsertChildAnchor(iter *TextIter, anchor *TextChildAnchor) {
	C.gtk_text_buffer_insert_child_anchor(v.native(), iter.native(), anchor.native())
}

// CreateChildAnchor() is a wrapper around gtk_text_buffer_insert_child_anchor().
func (v *TextBuffer) CreateChildAnchor(iter *TextIter) (*TextChildAnchor, error) {
	c := C.gtk_text_buffer_create_child_anchor(v.native(), iter.native())
	if c == nil {
		return nil, nilPtrErr
	}

	obj := glib.Take(unsafe.Pointer(c))
	return wrapTextChildAnchor(obj), nil
}

/*
 * GtkTextChildAnchor
 */

// TextChildAnchor is a representation of GTK's GtkTextChildAnchor.
type TextChildAnchor struct {
	*glib.Object
}

// native returns a pointer to the underlying GtkTextBuffer.
func (v *TextChildAnchor) native() *C.GtkTextChildAnchor {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toGtkTextChildAnchor(ptr)
}

func marshalTextChildAnchor(p uintptr) (interface{}, error) {
	c := C.g_value_get_object(C.toGValue(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapTextChildAnchor(obj), nil
}

func wrapTextChildAnchor(obj *glib.Object) *TextChildAnchor {
	return &TextChildAnchor{obj}
}

// TextChildAnchorNew() is a wrapper around gtk_text_child_anchor_new().
func TextChildAnchorNew() (*TextChildAnchor, error) {
	c := C.gtk_text_child_anchor_new()
	if c == nil {
		return nil, nilPtrErr
	}

	obj := glib.Take(unsafe.Pointer(c))
	return wrapTextChildAnchor(obj), nil
}

/*
 * GtkTextView
 */

// TextView is a representation of GTK's GtkTextView
type TextView struct {
	Container
}

// native returns a pointer to the underlying GtkTextView.
func (v *TextView) native() *C.GtkTextView {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toGtkTextView(ptr)
}

func marshalTextView(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapTextView(obj), nil
}

func wrapTextView(obj *glib.Object) *TextView {
	return &TextView{Container{Widget{glib.InitiallyUnowned{obj}}}}
}

// TextViewNew is a wrapper around gtk_text_view_new().
func TextViewNew() (*TextView, error) {
	c := C.gtk_text_view_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapTextView(obj), nil
}

// TextViewNewWithBuffer is a wrapper around gtk_text_view_new_with_buffer().
func TextViewNewWithBuffer(buf *TextBuffer) (*TextView, error) {
	cbuf := buf.native()
	c := C.gtk_text_view_new_with_buffer(cbuf)
	obj := glib.Take(unsafe.Pointer(c))
	return wrapTextView(obj), nil
}

// GetBuffer is a wrapper around gtk_text_view_get_buffer().
func (v *TextView) GetBuffer() (*TextBuffer, error) {
	c := C.gtk_text_view_get_buffer(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapTextBuffer(obj), nil
}

// SetBuffer is a wrapper around gtk_text_view_set_buffer().
func (v *TextView) SetBuffer(buffer *TextBuffer) {
	C.gtk_text_view_set_buffer(v.native(), buffer.native())
}

// SetEditable is a wrapper around gtk_text_view_set_editable().
func (v *TextView) SetEditable(editable bool) {
	C.gtk_text_view_set_editable(v.native(), gbool(editable))
}

// GetEditable is a wrapper around gtk_text_view_get_editable().
func (v *TextView) GetEditable() bool {
	c := C.gtk_text_view_get_editable(v.native())
	return gobool(c)
}

// SetWrapMode is a wrapper around gtk_text_view_set_wrap_mode().
func (v *TextView) SetWrapMode(wrapMode WrapMode) {
	C.gtk_text_view_set_wrap_mode(v.native(), C.GtkWrapMode(wrapMode))
}

// GetWrapMode is a wrapper around gtk_text_view_get_wrap_mode().
func (v *TextView) GetWrapMode() WrapMode {
	return WrapMode(C.gtk_text_view_get_wrap_mode(v.native()))
}

// SetCursorVisible is a wrapper around gtk_text_view_set_cursor_visible().
func (v *TextView) SetCursorVisible(visible bool) {
	C.gtk_text_view_set_cursor_visible(v.native(), gbool(visible))
}

// GetCursorVisible is a wrapper around gtk_text_view_get_cursor_visible().
func (v *TextView) GetCursorVisible() bool {
	c := C.gtk_text_view_get_cursor_visible(v.native())
	return gobool(c)
}

// SetOverwrite is a wrapper around gtk_text_view_set_overwrite().
func (v *TextView) SetOverwrite(overwrite bool) {
	C.gtk_text_view_set_overwrite(v.native(), gbool(overwrite))
}

// GetOverwrite is a wrapper around gtk_text_view_get_overwrite().
func (v *TextView) GetOverwrite() bool {
	c := C.gtk_text_view_get_overwrite(v.native())
	return gobool(c)
}

// SetJustification is a wrapper around gtk_text_view_set_justification().
func (v *TextView) SetJustification(justify Justification) {
	C.gtk_text_view_set_justification(v.native(), C.GtkJustification(justify))
}

// GetJustification is a wrapper around gtk_text_view_get_justification().
func (v *TextView) GetJustification() Justification {
	c := C.gtk_text_view_get_justification(v.native())
	return Justification(c)
}

// SetAcceptsTab is a wrapper around gtk_text_view_set_accepts_tab().
func (v *TextView) SetAcceptsTab(acceptsTab bool) {
	C.gtk_text_view_set_accepts_tab(v.native(), gbool(acceptsTab))
}

// GetAcceptsTab is a wrapper around gtk_text_view_get_accepts_tab().
func (v *TextView) GetAcceptsTab() bool {
	c := C.gtk_text_view_get_accepts_tab(v.native())
	return gobool(c)
}

// SetPixelsAboveLines is a wrapper around gtk_text_view_set_pixels_above_lines().
func (v *TextView) SetPixelsAboveLines(px int) {
	C.gtk_text_view_set_pixels_above_lines(v.native(), C.gint(px))
}

// GetPixelsAboveLines is a wrapper around gtk_text_view_get_pixels_above_lines().
func (v *TextView) GetPixelsAboveLines() int {
	c := C.gtk_text_view_get_pixels_above_lines(v.native())
	return int(c)
}

// SetPixelsBelowLines is a wrapper around gtk_text_view_set_pixels_below_lines().
func (v *TextView) SetPixelsBelowLines(px int) {
	C.gtk_text_view_set_pixels_below_lines(v.native(), C.gint(px))
}

// GetPixelsBelowLines is a wrapper around gtk_text_view_get_pixels_below_lines().
func (v *TextView) GetPixelsBelowLines() int {
	c := C.gtk_text_view_get_pixels_below_lines(v.native())
	return int(c)
}

// SetPixelsInsideWrap is a wrapper around gtk_text_view_set_pixels_inside_wrap().
func (v *TextView) SetPixelsInsideWrap(px int) {
	C.gtk_text_view_set_pixels_inside_wrap(v.native(), C.gint(px))
}

// GetPixelsInsideWrap is a wrapper around gtk_text_view_get_pixels_inside_wrap().
func (v *TextView) GetPixelsInsideWrap() int {
	c := C.gtk_text_view_get_pixels_inside_wrap(v.native())
	return int(c)
}

// SetLeftMargin is a wrapper around gtk_text_view_set_left_margin().
func (v *TextView) SetLeftMargin(margin int) {
	C.gtk_text_view_set_left_margin(v.native(), C.gint(margin))
}

// GetLeftMargin is a wrapper around gtk_text_view_get_left_margin().
func (v *TextView) GetLeftMargin() int {
	c := C.gtk_text_view_get_left_margin(v.native())
	return int(c)
}

// SetRightMargin is a wrapper around gtk_text_view_set_right_margin().
func (v *TextView) SetRightMargin(margin int) {
	C.gtk_text_view_set_right_margin(v.native(), C.gint(margin))
}

// GetRightMargin is a wrapper around gtk_text_view_get_right_margin().
func (v *TextView) GetRightMargin() int {
	c := C.gtk_text_view_get_right_margin(v.native())
	return int(c)
}

// SetIndent is a wrapper around gtk_text_view_set_indent().
func (v *TextView) SetIndent(indent int) {
	C.gtk_text_view_set_indent(v.native(), C.gint(indent))
}

// GetIndent is a wrapper around gtk_text_view_get_indent().
func (v *TextView) GetIndent() int {
	c := C.gtk_text_view_get_indent(v.native())
	return int(c)
}

// SetInputHints is a wrapper around gtk_text_view_set_input_hints().
func (v *TextView) SetInputHints(hints InputHints) {
	C.gtk_text_view_set_input_hints(v.native(), C.GtkInputHints(hints))
}

// GetInputHints is a wrapper around gtk_text_view_get_input_hints().
func (v *TextView) GetInputHints() InputHints {
	c := C.gtk_text_view_get_input_hints(v.native())
	return InputHints(c)
}

// SetInputPurpose is a wrapper around gtk_text_view_set_input_purpose().
func (v *TextView) SetInputPurpose(purpose InputPurpose) {
	C.gtk_text_view_set_input_purpose(v.native(),
		C.GtkInputPurpose(purpose))
}

// GetInputPurpose is a wrapper around gtk_text_view_get_input_purpose().
func (v *TextView) GetInputPurpose() InputPurpose {
	c := C.gtk_text_view_get_input_purpose(v.native())
	return InputPurpose(c)
}

// ScrollToMark is a wrapper around gtk_text_view_scroll_to_mark().
func (v *TextView) ScrollToMark(mark *TextMark, within_margin float64, use_align bool, xalign, yalign float64) {
	C.gtk_text_view_scroll_to_mark(v.native(), mark.native(), C.gdouble(within_margin), gbool(use_align), C.gdouble(xalign), C.gdouble(yalign))
}

// ScrollToIter is a wrapper around gtk_text_view_scroll_to_iter().
func (v *TextView) ScrollToIter(iter *TextIter, within_margin float64, use_align bool, xalign, yalign float64) bool {
	return gobool(C.gtk_text_view_scroll_to_iter(v.native(), iter.native(), C.gdouble(within_margin), gbool(use_align), C.gdouble(xalign), C.gdouble(yalign)))
}

// ScrollMarkOnscreen is a wrapper around gtk_text_view_scroll_mark_onscreen().
func (v *TextView) ScrollMarkOnscreen(mark *TextMark) {
	C.gtk_text_view_scroll_mark_onscreen(v.native(), mark.native())
}

// MoveMarkOnscreen is a wrapper around gtk_text_view_move_mark_onscreen().
func (v *TextView) MoveMarkOnscreen(mark *TextMark) bool {
	return gobool(C.gtk_text_view_move_mark_onscreen(v.native(), mark.native()))
}

// PlaceCursorOnscreen is a wrapper around gtk_text_view_place_cursor_onscreen().
func (v *TextView) PlaceCursorOnscreen() bool {
	return gobool(C.gtk_text_view_place_cursor_onscreen(v.native()))
}

// GetVisibleRect is a wrapper around gtk_text_view_get_visible_rect().
func (v *TextView) GetVisibleRect() *gdk.Rectangle {
	var rect C.GdkRectangle
	C.gtk_text_view_get_visible_rect(v.native(), &rect)
	return gdk.WrapRectangle(uintptr(unsafe.Pointer(&rect)))
}

// GetIterLocation is a wrapper around gtk_text_view_get_iter_location().
func (v *TextView) GetIterLocation(iter *TextIter) *gdk.Rectangle {
	var rect C.GdkRectangle
	C.gtk_text_view_get_iter_location(v.native(), iter.native(), &rect)
	return gdk.WrapRectangle(uintptr(unsafe.Pointer(&rect)))
}

// GetCursorLocations is a wrapper around gtk_text_view_get_cursor_locations().
func (v *TextView) GetCursorLocations(iter *TextIter) (strong, weak *gdk.Rectangle) {
	var strongRect, weakRect C.GdkRectangle
	C.gtk_text_view_get_cursor_locations(v.native(), iter.native(), &strongRect, &weakRect)
	return gdk.WrapRectangle(uintptr(unsafe.Pointer(&strongRect))), gdk.WrapRectangle(uintptr(unsafe.Pointer(&weakRect)))
}

// GetLineAtY is a wrapper around gtk_text_view_get_line_at_y().
func (v *TextView) GetLineAtY(y int) (*TextIter, int) {
	var iter C.GtkTextIter
	var line_top C.gint
	C.gtk_text_view_get_line_at_y(v.native(), &iter, C.gint(y), &line_top)
	ti := wrapTextIter(&iter)
	return ti, int(line_top)
}

// GetLineYrange is a wrapper around gtk_text_view_get_line_yrange().
func (v *TextView) GetLineYrange(iter *TextIter) (y, height int) {
	var yx, heightx C.gint
	C.gtk_text_view_get_line_yrange(v.native(), iter.native(), &yx, &heightx)
	return int(yx), int(heightx)
}

// GetIterAtLocation is a wrapper around gtk_text_view_get_iter_at_location().
func (v *TextView) GetIterAtLocation(x, y int) *TextIter {
	var iter C.GtkTextIter
	C.gtk_text_view_get_iter_at_location(v.native(), &iter, C.gint(x), C.gint(y))
	ti := wrapTextIter(&iter)
	return ti
}

// GetIterAtPosition is a wrapper around gtk_text_view_get_iter_at_position().
func (v *TextView) GetIterAtPosition(x, y int) (*TextIter, int) {
	var iter C.GtkTextIter
	var trailing C.gint
	C.gtk_text_view_get_iter_at_position(v.native(), &iter, &trailing, C.gint(x), C.gint(y))
	ti := wrapTextIter(&iter)
	return ti, int(trailing)
}

// BufferToWindowCoords is a wrapper around gtk_text_view_buffer_to_window_coords().
func (v *TextView) BufferToWindowCoords(win TextWindowType, buffer_x, buffer_y int) (window_x, window_y int) {
	var wx, wy C.gint
	C.gtk_text_view_buffer_to_window_coords(v.native(), C.GtkTextWindowType(win), C.gint(buffer_x), C.gint(buffer_y), &wx, &wy)
	return int(wx), int(wy)
}

// WindowToBufferCoords is a wrapper around gtk_text_view_window_to_buffer_coords().
func (v *TextView) WindowToBufferCoords(win TextWindowType, window_x, window_y int) (buffer_x, buffer_y int) {
	var bx, by C.gint
	C.gtk_text_view_window_to_buffer_coords(v.native(), C.GtkTextWindowType(win), C.gint(window_x), C.gint(window_y), &bx, &by)
	return int(bx), int(by)
}

// GetWindow is a wrapper around gtk_text_view_get_window().
func (v *TextView) GetWindow(win TextWindowType) *gdk.Window {
	c := C.gtk_text_view_get_window(v.native(), C.GtkTextWindowType(win))
	if c == nil {
		return nil
	}
	return &gdk.Window{glib.Take(unsafe.Pointer(c))}
}

// GetWindowType is a wrapper around gtk_text_view_get_window_type().
func (v *TextView) GetWindowType(w *gdk.Window) TextWindowType {
	return TextWindowType(C.gtk_text_view_get_window_type(v.native(), (*C.GdkWindow)(unsafe.Pointer(w.Native()))))
}

// SetBorderWindowSize is a wrapper around gtk_text_view_set_border_window_size().
func (v *TextView) SetBorderWindowSize(tp TextWindowType, size int) {
	C.gtk_text_view_set_border_window_size(v.native(), C.GtkTextWindowType(tp), C.gint(size))
}

// GetBorderWindowSize is a wrapper around gtk_text_view_get_border_window_size().
func (v *TextView) GetBorderWindowSize(tp TextWindowType) int {
	return int(C.gtk_text_view_get_border_window_size(v.native(), C.GtkTextWindowType(tp)))
}

// ForwardDisplayLine is a wrapper around gtk_text_view_forward_display_line().
func (v *TextView) ForwardDisplayLine(iter *TextIter) bool {
	return gobool(C.gtk_text_view_forward_display_line(v.native(), iter.native()))
}

// BackwardDisplayLine is a wrapper around gtk_text_view_backward_display_line().
func (v *TextView) BackwardDisplayLine(iter *TextIter) bool {
	return gobool(C.gtk_text_view_backward_display_line(v.native(), iter.native()))
}

// ForwardDisplayLineEnd is a wrapper around gtk_text_view_forward_display_line_end().
func (v *TextView) ForwardDisplayLineEnd(iter *TextIter) bool {
	return gobool(C.gtk_text_view_forward_display_line_end(v.native(), iter.native()))
}

// BackwardDisplayLineStart is a wrapper around gtk_text_view_backward_display_line_start().
func (v *TextView) BackwardDisplayLineStart(iter *TextIter) bool {
	return gobool(C.gtk_text_view_backward_display_line_start(v.native(), iter.native()))
}

// StartsDisplayLine is a wrapper around gtk_text_view_starts_display_line().
func (v *TextView) StartsDisplayLine(iter *TextIter) bool {
	return gobool(C.gtk_text_view_starts_display_line(v.native(), iter.native()))
}

// MoveVisually is a wrapper around gtk_text_view_move_visually().
func (v *TextView) MoveVisually(iter *TextIter, count int) bool {
	return gobool(C.gtk_text_view_move_visually(v.native(), iter.native(), C.gint(count)))
}

// AddChildAtAnchor is a wrapper around gtk_text_view_add_child_at_anchor().
func (v *TextView) AddChildAtAnchor(child IWidget, anchor *TextChildAnchor) {
	C.gtk_text_view_add_child_at_anchor(v.native(), child.toWidget(), anchor.native())
}

// AddChildInWindow is a wrapper around gtk_text_view_add_child_in_window().
func (v *TextView) AddChildInWindow(child IWidget, tp TextWindowType, xpos, ypos int) {
	C.gtk_text_view_add_child_in_window(v.native(), child.toWidget(), C.GtkTextWindowType(tp), C.gint(xpos), C.gint(ypos))
}

// MoveChild is a wrapper around gtk_text_view_move_child().
func (v *TextView) MoveChild(child IWidget, xpos, ypos int) {
	C.gtk_text_view_move_child(v.native(), child.toWidget(), C.gint(xpos), C.gint(ypos))
}

// ImContextFilterKeypress is a wrapper around gtk_text_view_im_context_filter_keypress().
func (v *TextView) ImContextFilterKeypress(event *gdk.EventKey) bool {
	return gobool(C.gtk_text_view_im_context_filter_keypress(v.native(), (*C.GdkEventKey)(unsafe.Pointer(event.Native()))))
}

// ResetImContext is a wrapper around gtk_text_view_reset_im_context().
func (v *TextView) ResetImContext() {
	C.gtk_text_view_reset_im_context(v.native())
}

// GtkAdjustment * 	gtk_text_view_get_hadjustment ()  -- DEPRECATED
// GtkAdjustment * 	gtk_text_view_get_vadjustment ()  -- DEPRECATED
// void 	gtk_text_view_add_child_at_anchor ()
// GtkTextChildAnchor * 	gtk_text_child_anchor_new ()
// GList * 	gtk_text_child_anchor_get_widgets ()
// gboolean 	gtk_text_child_anchor_get_deleted ()
// void 	gtk_text_view_set_top_margin () -- SINCE 3.18
// gint 	gtk_text_view_get_top_margin () -- SINCE 3.18
// void 	gtk_text_view_set_bottom_margin ()  -- SINCE 3.18
// gint 	gtk_text_view_get_bottom_margin ()  -- SINCE 3.18
// void 	gtk_text_view_set_tabs () -- PangoTabArray
// PangoTabArray * 	gtk_text_view_get_tabs () -- PangoTabArray
// GtkTextAttributes * 	gtk_text_view_get_default_attributes () -- GtkTextAttributes
// void 	gtk_text_view_set_monospace () -- SINCE 3.16
// gboolean 	gtk_text_view_get_monospace () -- SINCE 3.16
