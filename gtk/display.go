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

	"github.com/d2r2/gotk3/gdk"
	"github.com/d2r2/gotk3/glib"
	"github.com/d2r2/gotk3/pango"
)

/*
 * GtkLabel
 */

// Label is a representation of GTK's GtkLabel.
type Label struct {
	Widget
}

// native returns a pointer to the underlying GtkLabel.
func (v *Label) native() *C.GtkLabel {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toGtkLabel(ptr)
}

func marshalLabel(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapLabel(obj), nil
}

func wrapLabel(obj *glib.Object) *Label {
	return &Label{Widget{glib.InitiallyUnowned{obj}}}
}

// LabelNew is a wrapper around gtk_label_new().
func LabelNew(str string) (*Label, error) {
	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_label_new((*C.gchar)(cstr))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapLabel(obj), nil
}

// SetText is a wrapper around gtk_label_set_text().
func (v *Label) SetText(str string) {
	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_label_set_text(v.native(), (*C.gchar)(cstr))
}

// SetMarkup is a wrapper around gtk_label_set_markup().
func (v *Label) SetMarkup(str string) {
	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_label_set_markup(v.native(), (*C.gchar)(cstr))
}

// SetMarkupWithMnemonic is a wrapper around
// gtk_label_set_markup_with_mnemonic().
func (v *Label) SetMarkupWithMnemonic(str string) {
	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_label_set_markup_with_mnemonic(v.native(), (*C.gchar)(cstr))
}

// SetPattern is a wrapper around gtk_label_set_pattern().
func (v *Label) SetPattern(patern string) {
	cstr := C.CString(patern)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_label_set_pattern(v.native(), (*C.gchar)(cstr))
}

// SetJustify is a wrapper around gtk_label_set_justify().
func (v *Label) SetJustify(jtype Justification) {
	C.gtk_label_set_justify(v.native(), C.GtkJustification(jtype))
}

// SetEllipsize is a wrapper around gtk_label_set_ellipsize().
func (v *Label) SetEllipsize(mode pango.EllipsizeMode) {
	C.gtk_label_set_ellipsize(v.native(), C.PangoEllipsizeMode(mode))
}

// GetWidthChars is a wrapper around gtk_label_get_width_chars().
func (v *Label) GetWidthChars() int {
	c := C.gtk_label_get_width_chars(v.native())
	return int(c)
}

// SetWidthChars is a wrapper around gtk_label_set_width_chars().
func (v *Label) SetWidthChars(nChars int) {
	C.gtk_label_set_width_chars(v.native(), C.gint(nChars))
}

// GetMaxWidthChars is a wrapper around gtk_label_get_max_width_chars().
func (v *Label) GetMaxWidthChars() int {
	c := C.gtk_label_get_max_width_chars(v.native())
	return int(c)
}

// SetMaxWidthChars is a wrapper around gtk_label_set_max_width_chars().
func (v *Label) SetMaxWidthChars(nChars int) {
	C.gtk_label_set_max_width_chars(v.native(), C.gint(nChars))
}

// GetLineWrap is a wrapper around gtk_label_get_line_wrap().
func (v *Label) GetLineWrap() bool {
	c := C.gtk_label_get_line_wrap(v.native())
	return gobool(c)
}

// SetLineWrap is a wrapper around gtk_label_set_line_wrap().
func (v *Label) SetLineWrap(wrap bool) {
	C.gtk_label_set_line_wrap(v.native(), gbool(wrap))
}

// SetLineWrapMode is a wrapper around gtk_label_set_line_wrap_mode().
func (v *Label) SetLineWrapMode(wrapMode pango.WrapMode) {
	C.gtk_label_set_line_wrap_mode(v.native(), C.PangoWrapMode(wrapMode))
}

// GetSelectable is a wrapper around gtk_label_get_selectable().
func (v *Label) GetSelectable() bool {
	c := C.gtk_label_get_selectable(v.native())
	return gobool(c)
}

// GetText is a wrapper around gtk_label_get_text().
func (v *Label) GetText() (string, error) {
	c := C.gtk_label_get_text(v.native())
	if c == nil {
		return "", nilPtrErr
	}
	return goString(c), nil
}

// GetJustify is a wrapper around gtk_label_get_justify().
func (v *Label) GetJustify() Justification {
	c := C.gtk_label_get_justify(v.native())
	return Justification(c)
}

// GetEllipsize is a wrapper around gtk_label_get_ellipsize().
func (v *Label) GetEllipsize() pango.EllipsizeMode {
	c := C.gtk_label_get_ellipsize(v.native())
	return pango.EllipsizeMode(c)
}

// GetCurrentUri is a wrapper around gtk_label_get_current_uri().
func (v *Label) GetCurrentUri() string {
	c := C.gtk_label_get_current_uri(v.native())
	return goString(c)
}

// GetTrackVisitedLinks is a wrapper around gtk_label_get_track_visited_links().
func (v *Label) GetTrackVisitedLinks() bool {
	c := C.gtk_label_get_track_visited_links(v.native())
	return gobool(c)
}

// SetTrackVisitedLinks is a wrapper around gtk_label_set_track_visited_links().
func (v *Label) SetTrackVisitedLinks(trackLinks bool) {
	C.gtk_label_set_track_visited_links(v.native(), gbool(trackLinks))
}

// GetAngle is a wrapper around gtk_label_get_angle().
func (v *Label) GetAngle() float64 {
	c := C.gtk_label_get_angle(v.native())
	return float64(c)
}

// SetAngle is a wrapper around gtk_label_set_angle().
func (v *Label) SetAngle(angle float64) {
	C.gtk_label_set_angle(v.native(), C.gdouble(angle))
}

// GetSelectionBounds is a wrapper around gtk_label_get_selection_bounds().
func (v *Label) GetSelectionBounds() (start, end int, nonEmpty bool) {
	var cstart, cend C.gint
	c := C.gtk_label_get_selection_bounds(v.native(), &cstart, &cend)
	return int(cstart), int(cend), gobool(c)
}

// GetSingleLineMode is a wrapper around gtk_label_get_single_line_mode().
func (v *Label) GetSingleLineMode() bool {
	c := C.gtk_label_get_single_line_mode(v.native())
	return gobool(c)
}

// SetSingleLineMode is a wrapper around gtk_label_set_single_line_mode().
func (v *Label) SetSingleLineMode(mode bool) {
	C.gtk_label_set_single_line_mode(v.native(), gbool(mode))
}

// GetUseMarkup is a wrapper around gtk_label_get_use_markup().
func (v *Label) GetUseMarkup() bool {
	c := C.gtk_label_get_use_markup(v.native())
	return gobool(c)
}

// SetUseMarkup is a wrapper around gtk_label_set_use_markup().
func (v *Label) SetUseMarkup(use bool) {
	C.gtk_label_set_use_markup(v.native(), gbool(use))
}

// GetUseUnderline is a wrapper around gtk_label_get_use_underline().
func (v *Label) GetUseUnderline() bool {
	c := C.gtk_label_get_use_underline(v.native())
	return gobool(c)
}

// SetUseUnderline is a wrapper around gtk_label_set_use_underline().
func (v *Label) SetUseUnderline(use bool) {
	C.gtk_label_set_use_underline(v.native(), gbool(use))
}

// LabelNewWithMnemonic is a wrapper around gtk_label_new_with_mnemonic().
func LabelNewWithMnemonic(str string) (*Label, error) {
	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_label_new_with_mnemonic((*C.gchar)(cstr))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapLabel(obj), nil
}

// SelectRegion is a wrapper around gtk_label_select_region().
func (v *Label) SelectRegion(startOffset, endOffset int) {
	C.gtk_label_select_region(v.native(), C.gint(startOffset),
		C.gint(endOffset))
}

// SetSelectable is a wrapper around gtk_label_set_selectable().
func (v *Label) SetSelectable(setting bool) {
	C.gtk_label_set_selectable(v.native(), gbool(setting))
}

// SetLabel is a wrapper around gtk_label_set_label().
func (v *Label) SetLabel(str string) {
	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_label_set_label(v.native(), (*C.gchar)(cstr))
}

// GetLabel is a wrapper around gtk_label_get_label().
func (v *Label) GetLabel() string {
	c := C.gtk_label_get_label(v.native())
	if c == nil {
		return ""
	}
	return goString(c)
}

// GetMnemonicKeyval is a wrapper around gtk_label_get_mnemonic_keyval().
func (v *Label) GetMnemonicKeyval() uint {
	return uint(C.gtk_label_get_mnemonic_keyval(v.native()))
}

/*
 * GtkImage
 */

// Image is a representation of GTK's GtkImage.
type Image struct {
	Widget
}

// native returns a pointer to the underlying GtkImage.
func (v *Image) native() *C.GtkImage {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toGtkImage(ptr)
}

func marshalImage(p uintptr) (interface{}, error) {
	c := C.g_value_get_object(C.toGValue(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapImage(obj), nil
}

func wrapImage(obj *glib.Object) *Image {
	widget := wrapWidget(obj)
	return &Image{*widget}
}

// ImageNew() is a wrapper around gtk_image_new().
func ImageNew() (*Image, error) {
	c := C.gtk_image_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapImage(obj), nil
}

// ImageNewFromFile() is a wrapper around gtk_image_new_from_file().
func ImageNewFromFile(filename string) (*Image, error) {
	cstr := C.CString(filename)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_image_new_from_file((*C.gchar)(cstr))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapImage(obj), nil
}

// ImageNewFromResource() is a wrapper around gtk_image_new_from_resource().
func ImageNewFromResource(resourcePath string) (*Image, error) {
	cstr := C.CString(resourcePath)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_image_new_from_resource((*C.gchar)(cstr))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapImage(obj), nil
}

// ImageNewFromPixbuf is a wrapper around gtk_image_new_from_pixbuf().
func ImageNewFromPixbuf(pixbuf *gdk.Pixbuf) (*Image, error) {
	ptr := (*C.GdkPixbuf)(unsafe.Pointer(pixbuf.Native()))
	c := C.gtk_image_new_from_pixbuf(ptr)
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapImage(obj), nil
}

// ImageNewFromPixbuf is a wrapper around gtk_image_new_from_pixbuf().
func ImageNewFromAnimation(animation *gdk.PixbufAnimation) (*Image, error) {
	ptr := (*C.GdkPixbufAnimation)(unsafe.Pointer(animation.Native()))
	c := C.gtk_image_new_from_animation(ptr)
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapImage(obj), nil
}

// TODO(jrick) GtkIconSet
/*
func ImageNewFromIconSet() {
}
*/

// TODO(jrick) GdkPixbufAnimation
/*
func ImageNewFromAnimation() {
}
*/

// ImageNewFromIconName() is a wrapper around gtk_image_new_from_icon_name().
func ImageNewFromIconName(iconName string, size IconSize) (*Image, error) {
	cstr := C.CString(iconName)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_image_new_from_icon_name((*C.gchar)(cstr),
		C.GtkIconSize(size))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapImage(obj), nil
}

// TODO(jrick) GIcon
/*
func ImageNewFromGIcon() {
}
*/

// Clear() is a wrapper around gtk_image_clear().
func (v *Image) Clear() {
	C.gtk_image_clear(v.native())
}

// SetFromFile() is a wrapper around gtk_image_set_from_file().
func (v *Image) SetFromFile(filename string) {
	cstr := C.CString(filename)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_image_set_from_file(v.native(), (*C.gchar)(cstr))
}

// SetFromResource() is a wrapper around gtk_image_set_from_resource().
func (v *Image) SetFromResource(resourcePath string) {
	cstr := C.CString(resourcePath)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_image_set_from_resource(v.native(), (*C.gchar)(cstr))
}

// SetFromFixbuf is a wrapper around gtk_image_set_from_pixbuf().
func (v *Image) SetFromPixbuf(pixbuf *gdk.Pixbuf) {
	pbptr := (*C.GdkPixbuf)(unsafe.Pointer(pixbuf.Native()))
	C.gtk_image_set_from_pixbuf(v.native(), pbptr)
}

// TODO(jrick) GtkIconSet
/*
func (v *Image) SetFromIconSet() {
}
*/

// TODO(jrick) GdkPixbufAnimation
/*
func (v *Image) SetFromAnimation() {
}
*/

// SetFromIconName() is a wrapper around gtk_image_set_from_icon_name().
func (v *Image) SetFromIconName(iconName string, size IconSize) {
	cstr := C.CString(iconName)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_image_set_from_icon_name(v.native(), (*C.gchar)(cstr),
		C.GtkIconSize(size))
}

// TODO(jrick) GIcon
/*
func (v *Image) SetFromGIcon() {
}
*/

// SetPixelSize() is a wrapper around gtk_image_set_pixel_size().
func (v *Image) SetPixelSize(pixelSize int) {
	C.gtk_image_set_pixel_size(v.native(), C.gint(pixelSize))
}

// GetStorageType() is a wrapper around gtk_image_get_storage_type().
func (v *Image) GetStorageType() ImageType {
	c := C.gtk_image_get_storage_type(v.native())
	return ImageType(c)
}

// GetPixbuf() is a wrapper around gtk_image_get_pixbuf().
func (v *Image) GetPixbuf() *gdk.Pixbuf {
	c := C.gtk_image_get_pixbuf(v.native())
	if c == nil {
		return nil
	}

	pb := &gdk.Pixbuf{glib.Take(unsafe.Pointer(c))}
	return pb
}

// TODO(jrick) GtkIconSet
/*
func (v *Image) GetIconSet() {
}
*/

// TODO(jrick) GdkPixbufAnimation
/*
func (v *Image) GetAnimation() {
}
*/

// GetIconName() is a wrapper around gtk_image_get_icon_name().
func (v *Image) GetIconName() (string, IconSize) {
	var iconName *C.gchar
	var size C.GtkIconSize
	C.gtk_image_get_icon_name(v.native(), &iconName, &size)
	return goString(iconName), IconSize(size)
}

// TODO(jrick) GIcon
/*
func (v *Image) GetGIcon() {
}
*/

// GetPixelSize() is a wrapper around gtk_image_get_pixel_size().
func (v *Image) GetPixelSize() int {
	c := C.gtk_image_get_pixel_size(v.native())
	return int(c)
}

/*
 * GtkSpinner
 */

// Spinner is a representation of GTK's GtkSpinner.
type Spinner struct {
	Widget
}

// native returns a pointer to the underlying GtkSpinner.
func (v *Spinner) native() *C.GtkSpinner {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toGtkSpinner(ptr)
}

func marshalSpinner(p uintptr) (interface{}, error) {
	c := C.g_value_get_object(C.toGValue(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapSpinner(obj), nil
}

func wrapSpinner(obj *glib.Object) *Spinner {
	widget := wrapWidget(obj)
	return &Spinner{*widget}
}

// SpinnerNew is a wrapper around gtk_spinner_new().
func SpinnerNew() (*Spinner, error) {
	c := C.gtk_spinner_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapSpinner(obj), nil
}

// Start is a wrapper around gtk_spinner_start().
func (v *Spinner) Start() {
	C.gtk_spinner_start(v.native())
}

// Stop is a wrapper around gtk_spinner_stop().
func (v *Spinner) Stop() {
	C.gtk_spinner_stop(v.native())
}

/*
 * GtkInfoBar
 */

// InfoBar is a representation of GTK's GtkInfoBar.
type InfoBar struct {
	Box
}

func init() {
	tm := []glib.TypeMarshaler{
		{glib.Type(C.gtk_info_bar_get_type()), marshalInfoBar},
	}

	glib.RegisterGValueMarshalers(tm)

	WrapMap["GtkInfoBar"] = wrapInfoBar
}

func (v *InfoBar) native() *C.GtkInfoBar {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toGtkInfoBar(ptr)
}

func marshalInfoBar(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapInfoBar(obj), nil
}

func wrapInfoBar(obj *glib.Object) *InfoBar {
	return &InfoBar{Box{Container{Widget{glib.InitiallyUnowned{obj}}}}}
}

func InfoBarNew() (*InfoBar, error) {
	c := C.gtk_info_bar_new()
	if c == nil {
		return nil, nilPtrErr
	}

	obj := glib.Take(unsafe.Pointer(c))
	return wrapInfoBar(obj), nil
}

func (v *InfoBar) AddActionWidget(w IWidget, responseId ResponseType) {
	C.gtk_info_bar_add_action_widget(v.native(), w.toWidget(), C.gint(responseId))
}

func (v *InfoBar) AddButton(buttonText string, responseId ResponseType) {
	cstr := C.CString(buttonText)
	defer C.free(unsafe.Pointer(cstr))

	C.gtk_info_bar_add_button(v.native(), (*C.gchar)(cstr), C.gint(responseId))
}

func (v *InfoBar) SetResponseSensitive(responseId ResponseType, setting bool) {
	C.gtk_info_bar_set_response_sensitive(v.native(), C.gint(responseId), gbool(setting))
}

func (v *InfoBar) SetDefaultResponse(responseId ResponseType) {
	C.gtk_info_bar_set_default_response(v.native(), C.gint(responseId))
}

func (v *InfoBar) SetMessageType(messageType MessageType) {
	C.gtk_info_bar_set_message_type(v.native(), C.GtkMessageType(messageType))
}

func (v *InfoBar) GetMessageType() MessageType {
	messageType := C.gtk_info_bar_get_message_type(v.native())
	return MessageType(messageType)
}

func (v *InfoBar) GetActionArea() (*Widget, error) {
	c := C.gtk_info_bar_get_action_area(v.native())
	if c == nil {
		return nil, nilPtrErr
	}

	obj := glib.Take(unsafe.Pointer(c))
	return wrapWidget(obj), nil
}

func (v *InfoBar) GetContentArea() (*Box, error) {
	c := C.gtk_info_bar_get_content_area(v.native())
	if c == nil {
		return nil, nilPtrErr
	}

	obj := glib.Take(unsafe.Pointer(c))
	return wrapBox(obj), nil
}

func (v *InfoBar) GetShowCloseButton() bool {
	b := C.gtk_info_bar_get_show_close_button(v.native())
	return gobool(b)
}

func (v *InfoBar) SetShowCloseButton(setting bool) {
	C.gtk_info_bar_set_show_close_button(v.native(), gbool(setting))
}

/*
 * GtkProgressBar
 */

// ProgressBar is a representation of GTK's GtkProgressBar.
type ProgressBar struct {
	Widget
	// Interfaces
	Orientable
}

// native returns a pointer to the underlying GtkProgressBar.
func (v *ProgressBar) native() *C.GtkProgressBar {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toGtkProgressBar(ptr)
}

func marshalProgressBar(p uintptr) (interface{}, error) {
	c := C.g_value_get_object(C.toGValue(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapProgressBar(obj), nil
}

func wrapProgressBar(obj *glib.Object) *ProgressBar {
	widget := wrapWidget(obj)
	o := wrapOrientable(*glib.InterfaceFromObjectNew(obj))
	return &ProgressBar{*widget, *o}
}

// ProgressBarNew() is a wrapper around gtk_progress_bar_new().
func ProgressBarNew() (*ProgressBar, error) {
	c := C.gtk_progress_bar_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapProgressBar(obj), nil
}

// SetFraction() is a wrapper around gtk_progress_bar_set_fraction().
func (v *ProgressBar) SetFraction(fraction float64) {
	C.gtk_progress_bar_set_fraction(v.native(), C.gdouble(fraction))
}

// GetFraction() is a wrapper around gtk_progress_bar_get_fraction().
func (v *ProgressBar) GetFraction() float64 {
	c := C.gtk_progress_bar_get_fraction(v.native())
	return float64(c)
}

// SetShowText is a wrapper around gtk_progress_bar_set_show_text().
func (v *ProgressBar) SetShowText(showText bool) {
	C.gtk_progress_bar_set_show_text(v.native(), gbool(showText))
}

// GetShowText is a wrapper around gtk_progress_bar_get_show_text().
func (v *ProgressBar) GetShowText() bool {
	c := C.gtk_progress_bar_get_show_text(v.native())
	return gobool(c)
}

// SetText() is a wrapper around gtk_progress_bar_set_text().
func (v *ProgressBar) SetText(text string) {
	var cstr *C.char
	if text != "" {
		cstr := C.CString(text)
		defer C.free(unsafe.Pointer(cstr))
	}
	C.gtk_progress_bar_set_text(v.native(), (*C.gchar)(cstr))
}

// SetPulseStep is a wrapper around gtk_progress_bar_set_pulse_step().
func (v *ProgressBar) SetPulseStep(fraction float64) {
	C.gtk_progress_bar_set_pulse_step(v.native(), C.gdouble(fraction))
}

// GetPulseStep is a wrapper around gtk_progress_bar_get_pulse_step().
func (v *ProgressBar) GetPulseStep() float64 {
	c := C.gtk_progress_bar_get_pulse_step(v.native())
	return float64(c)
}

// Pulse is a wrapper arountd gtk_progress_bar_pulse().
func (v *ProgressBar) Pulse() {
	C.gtk_progress_bar_pulse(v.native())
}

// SetInverted is a wrapper around gtk_progress_bar_set_inverted().
func (v *ProgressBar) SetInverted(inverted bool) {
	C.gtk_progress_bar_set_inverted(v.native(), gbool(inverted))
}

// GetInverted is a wrapper around gtk_progress_bar_get_inverted().
func (v *ProgressBar) GetInverted() bool {
	c := C.gtk_progress_bar_get_inverted(v.native())
	return gobool(c)
}

/*
 * GtkLevelBar
 */

type LevelBar struct {
	Widget
}

func init() {
	tm := []glib.TypeMarshaler{
		{glib.Type(C.gtk_level_bar_mode_get_type()), marshalLevelBarMode},

		{glib.Type(C.gtk_level_bar_get_type()), marshalLevelBar},
	}

	glib.RegisterGValueMarshalers(tm)

	WrapMap["GtkLevelBar"] = wrapLevelBar
}

// LevelBarMode is a representation of GTK's GtkLevelBarMode.
type LevelBarMode int

const (
	LEVEL_BAR_MODE_CONTINUOUS LevelBarMode = C.GTK_LEVEL_BAR_MODE_CONTINUOUS
	LEVEL_BAR_MODE_DISCRETE   LevelBarMode = C.GTK_LEVEL_BAR_MODE_DISCRETE
)

func marshalLevelBarMode(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return LevelBarMode(c), nil
}

// native returns a pointer to the underlying GtkLevelBar.
func (v *LevelBar) native() *C.GtkLevelBar {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toGtkLevelBar(ptr)
}

func marshalLevelBar(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapLevelBar(obj), nil
}

func wrapLevelBar(obj *glib.Object) *LevelBar {
	return &LevelBar{Widget{glib.InitiallyUnowned{obj}}}
}

// LevelBarNew() is a wrapper around gtk_level_bar_new().
func LevelBarNew() (*LevelBar, error) {
	c := C.gtk_level_bar_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapLevelBar(obj), nil
}

// LevelBarNewForInterval() is a wrapper around gtk_level_bar_new_for_interval().
func LevelBarNewForInterval(min_value, max_value float64) (*LevelBar, error) {
	c := C.gtk_level_bar_new_for_interval(C.gdouble(min_value), C.gdouble(max_value))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapLevelBar(obj), nil
}

// SetMode() is a wrapper around gtk_level_bar_set_mode().
func (v *LevelBar) SetMode(m LevelBarMode) {
	C.gtk_level_bar_set_mode(v.native(), C.GtkLevelBarMode(m))
}

// GetMode() is a wrapper around gtk_level_bar_get_mode().
func (v *LevelBar) GetMode() LevelBarMode {
	return LevelBarMode(C.gtk_level_bar_get_mode(v.native()))
}

// SetValue() is a wrapper around gtk_level_bar_set_value().
func (v *LevelBar) SetValue(value float64) {
	C.gtk_level_bar_set_value(v.native(), C.gdouble(value))
}

// GetValue() is a wrapper around gtk_level_bar_get_value().
func (v *LevelBar) GetValue() float64 {
	c := C.gtk_level_bar_get_value(v.native())
	return float64(c)
}

// SetMinValue() is a wrapper around gtk_level_bar_set_min_value().
func (v *LevelBar) SetMinValue(value float64) {
	C.gtk_level_bar_set_min_value(v.native(), C.gdouble(value))
}

// GetMinValue() is a wrapper around gtk_level_bar_get_min_value().
func (v *LevelBar) GetMinValue() float64 {
	c := C.gtk_level_bar_get_min_value(v.native())
	return float64(c)
}

// SetMaxValue() is a wrapper around gtk_level_bar_set_max_value().
func (v *LevelBar) SetMaxValue(value float64) {
	C.gtk_level_bar_set_max_value(v.native(), C.gdouble(value))
}

// GetMaxValue() is a wrapper around gtk_level_bar_get_max_value().
func (v *LevelBar) GetMaxValue() float64 {
	c := C.gtk_level_bar_get_max_value(v.native())
	return float64(c)
}

const (
	LEVEL_BAR_OFFSET_LOW  string = C.GTK_LEVEL_BAR_OFFSET_LOW
	LEVEL_BAR_OFFSET_HIGH string = C.GTK_LEVEL_BAR_OFFSET_HIGH
)

// AddOffsetValue() is a wrapper around gtk_level_bar_add_offset_value().
func (v *LevelBar) AddOffsetValue(name string, value float64) {
	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_level_bar_add_offset_value(v.native(), (*C.gchar)(cstr), C.gdouble(value))
}

// RemoveOffsetValue() is a wrapper around gtk_level_bar_remove_offset_value().
func (v *LevelBar) RemoveOffsetValue(name string) {
	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_level_bar_remove_offset_value(v.native(), (*C.gchar)(cstr))
}

// GetOffsetValue() is a wrapper around gtk_level_bar_get_offset_value().
func (v *LevelBar) GetOffsetValue(name string) (float64, bool) {
	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))
	var value C.gdouble
	c := C.gtk_level_bar_get_offset_value(v.native(), (*C.gchar)(cstr), &value)
	return float64(value), gobool(c)
}

/*
 * GtkStatusbar
 */

// Statusbar is a representation of GTK's GtkStatusbar
type Statusbar struct {
	Box
}

// native returns a pointer to the underlying GtkStatusbar
func (v *Statusbar) native() *C.GtkStatusbar {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toGtkStatusbar(ptr)
}

func marshalStatusbar(p uintptr) (interface{}, error) {
	c := C.g_value_get_object(C.toGValue(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapStatusbar(obj), nil
}

func wrapStatusbar(obj *glib.Object) *Statusbar {
	box := wrapBox(obj)
	return &Statusbar{*box}
}

// StatusbarNew() is a wrapper around gtk_statusbar_new().
func StatusbarNew() (*Statusbar, error) {
	c := C.gtk_statusbar_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapStatusbar(obj), nil
}

// GetContextId() is a wrapper around gtk_statusbar_get_context_id().
func (v *Statusbar) GetContextId(contextDescription string) uint {
	cstr := C.CString(contextDescription)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_statusbar_get_context_id(v.native(), (*C.gchar)(cstr))
	return uint(c)
}

// Push() is a wrapper around gtk_statusbar_push().
func (v *Statusbar) Push(contextID uint, text string) uint {
	cstr := C.CString(text)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_statusbar_push(v.native(), C.guint(contextID),
		(*C.gchar)(cstr))
	return uint(c)
}

// Pop() is a wrapper around gtk_statusbar_pop().
func (v *Statusbar) Pop(contextID uint) {
	C.gtk_statusbar_pop(v.native(), C.guint(contextID))
}

// GetMessageArea() is a wrapper around gtk_statusbar_get_message_area().
func (v *Statusbar) GetMessageArea() (*Box, error) {
	c := C.gtk_statusbar_get_message_area(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	box := wrapBox(obj)
	return box, nil
}

// TODO: implement GtkAccelLabel
