// Same copyright and license as the rest of the files in this project
// This file contains style related functions and structures

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
 * GtkButton
 */

// Button is a representation of GTK's GtkButton.
type Button struct {
	Bin
	// Interfaces
	Actionable
}

// native() returns a pointer to the underlying GtkButton.
func (v *Button) native() *C.GtkButton {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toGtkButton(ptr)
}

func marshalButton(p uintptr) (interface{}, error) {
	c := C.g_value_get_object(C.toGValue(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapButton(obj), nil
}

func wrapButton(obj *glib.Object) *Button {
	bin := wrapBin(obj)
	actionable := wrapActionable(*glib.InterfaceFromObjectNew(obj))
	return &Button{*bin, *actionable}
}

// ButtonNew() is a wrapper around gtk_button_new().
func ButtonNew() (*Button, error) {
	c := C.gtk_button_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapButton(obj), nil
}

// ButtonNewWithLabel() is a wrapper around gtk_button_new_with_label().
func ButtonNewWithLabel(label string) (*Button, error) {
	cstr := C.CString(label)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_button_new_with_label((*C.gchar)(cstr))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapButton(obj), nil
}

// ButtonNewWithMnemonic() is a wrapper around gtk_button_new_with_mnemonic().
func ButtonNewWithMnemonic(label string) (*Button, error) {
	cstr := C.CString(label)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_button_new_with_mnemonic((*C.gchar)(cstr))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapButton(obj), nil
}

// Clicked() is a wrapper around gtk_button_clicked().
func (v *Button) Clicked() {
	C.gtk_button_clicked(v.native())
}

// SetRelief() is a wrapper around gtk_button_set_relief().
func (v *Button) SetRelief(newStyle ReliefStyle) {
	C.gtk_button_set_relief(v.native(), C.GtkReliefStyle(newStyle))
}

// GetRelief() is a wrapper around gtk_button_get_relief().
func (v *Button) GetRelief() ReliefStyle {
	c := C.gtk_button_get_relief(v.native())
	return ReliefStyle(c)
}

// SetLabel() is a wrapper around gtk_button_set_label().
func (v *Button) SetLabel(label string) {
	cstr := C.CString(label)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_button_set_label(v.native(), (*C.gchar)(cstr))
}

// GetLabel() is a wrapper around gtk_button_get_label().
func (v *Button) GetLabel() (string, error) {
	c := C.gtk_button_get_label(v.native())
	if c == nil {
		return "", nilPtrErr
	}
	return goString(c), nil
}

// SetUseUnderline() is a wrapper around gtk_button_set_use_underline().
func (v *Button) SetUseUnderline(useUnderline bool) {
	C.gtk_button_set_use_underline(v.native(), gbool(useUnderline))
}

// GetUseUnderline() is a wrapper around gtk_button_get_use_underline().
func (v *Button) GetUseUnderline() bool {
	c := C.gtk_button_get_use_underline(v.native())
	return gobool(c)
}

// SetImage() is a wrapper around gtk_button_set_image().
func (v *Button) SetImage(image IWidget) {
	C.gtk_button_set_image(v.native(), image.toWidget())
}

// GetImage() is a wrapper around gtk_button_get_image().
func (v *Button) GetImage() (*Widget, error) {
	c := C.gtk_button_get_image(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapWidget(obj), nil
}

// SetImagePosition() is a wrapper around gtk_button_set_image_position().
func (v *Button) SetImagePosition(position PositionType) {
	C.gtk_button_set_image_position(v.native(), C.GtkPositionType(position))
}

// GetImagePosition() is a wrapper around gtk_button_get_image_position().
func (v *Button) GetImagePosition() PositionType {
	c := C.gtk_button_get_image_position(v.native())
	return PositionType(c)
}

// SetAlwaysShowImage() is a wrapper around gtk_button_set_always_show_image().
func (v *Button) SetAlwaysShowImage(alwaysShow bool) {
	C.gtk_button_set_always_show_image(v.native(), gbool(alwaysShow))
}

// GetAlwaysShowImage() is a wrapper around gtk_button_get_always_show_image().
func (v *Button) GetAlwaysShowImage() bool {
	c := C.gtk_button_get_always_show_image(v.native())
	return gobool(c)
}

// GetEventWindow() is a wrapper around gtk_button_get_event_window().
func (v *Button) GetEventWindow() (*gdk.Window, error) {
	c := C.gtk_button_get_event_window(v.native())
	if c == nil {
		return nil, nilPtrErr
	}

	obj := glib.Take(unsafe.Pointer(c))
	return &gdk.Window{obj}, nil
}

/*
 * GtkCheckButton
 */

// CheckButton is a wrapper around GTK's GtkCheckButton.
type CheckButton struct {
	ToggleButton
}

// native returns a pointer to the underlying GtkCheckButton.
func (v *CheckButton) native() *C.GtkCheckButton {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toGtkCheckButton(ptr)
}

func marshalCheckButton(p uintptr) (interface{}, error) {
	c := C.g_value_get_object(C.toGValue(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapCheckButton(obj), nil
}

func wrapCheckButton(obj *glib.Object) *CheckButton {
	toggleButton := wrapToggleButton(obj)
	return &CheckButton{*toggleButton}
}

// CheckButtonNew is a wrapper around gtk_check_button_new().
func CheckButtonNew() (*CheckButton, error) {
	c := C.gtk_check_button_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapCheckButton(obj), nil
}

// CheckButtonNewWithLabel is a wrapper around
// gtk_check_button_new_with_label().
func CheckButtonNewWithLabel(label string) (*CheckButton, error) {
	cstr := C.CString(label)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_check_button_new_with_label((*C.gchar)(cstr))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapCheckButton(obj), nil
}

// CheckButtonNewWithMnemonic is a wrapper around
// gtk_check_button_new_with_mnemonic().
func CheckButtonNewWithMnemonic(label string) (*CheckButton, error) {
	cstr := C.CString(label)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_check_button_new_with_mnemonic((*C.gchar)(cstr))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapCheckButton(obj), nil
}

/*
 * GtkRadioButton
 */

// RadioButton is a representation of GTK's GtkRadioButton.
type RadioButton struct {
	CheckButton
}

// native returns a pointer to the underlying GtkRadioButton.
func (v *RadioButton) native() *C.GtkRadioButton {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toGtkRadioButton(ptr)
}

func marshalRadioButton(p uintptr) (interface{}, error) {
	c := C.g_value_get_object(C.toGValue(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapRadioButton(obj), nil
}

func wrapRadioButton(obj *glib.Object) *RadioButton {
	checkButton := wrapCheckButton(obj)
	return &RadioButton{*checkButton}
}

// RadioButtonNew is a wrapper around gtk_radio_button_new().
func RadioButtonNew(group *glib.SList) (*RadioButton, error) {
	c := C.gtk_radio_button_new(cGSList(group))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapRadioButton(obj), nil
}

// RadioButtonNewFromWidget is a wrapper around
// gtk_radio_button_new_from_widget().
func RadioButtonNewFromWidget(radioGroupMember *RadioButton) (*RadioButton, error) {
	c := C.gtk_radio_button_new_from_widget(radioGroupMember.native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapRadioButton(obj), nil
}

// RadioButtonNewWithLabel is a wrapper around
// gtk_radio_button_new_with_label().
func RadioButtonNewWithLabel(group *glib.SList, label string) (*RadioButton, error) {
	cstr := C.CString(label)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_radio_button_new_with_label(cGSList(group), (*C.gchar)(cstr))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapRadioButton(obj), nil
}

// RadioButtonNewWithLabelFromWidget is a wrapper around
// gtk_radio_button_new_with_label_from_widget().
func RadioButtonNewWithLabelFromWidget(radioGroupMember *RadioButton, label string) (*RadioButton, error) {
	cstr := C.CString(label)
	defer C.free(unsafe.Pointer(cstr))
	var cradio *C.GtkRadioButton
	if radioGroupMember != nil {
		cradio = radioGroupMember.native()
	}
	c := C.gtk_radio_button_new_with_label_from_widget(cradio, (*C.gchar)(cstr))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapRadioButton(obj), nil
}

// RadioButtonNewWithMnemonic is a wrapper around
// gtk_radio_button_new_with_mnemonic()
func RadioButtonNewWithMnemonic(group *glib.SList, label string) (*RadioButton, error) {
	cstr := C.CString(label)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_radio_button_new_with_mnemonic(cGSList(group), (*C.gchar)(cstr))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapRadioButton(obj), nil
}

// RadioButtonNewWithMnemonicFromWidget is a wrapper around
// gtk_radio_button_new_with_mnemonic_from_widget().
func RadioButtonNewWithMnemonicFromWidget(radioGroupMember *RadioButton, label string) (*RadioButton, error) {
	cstr := C.CString(label)
	defer C.free(unsafe.Pointer(cstr))
	var cradio *C.GtkRadioButton
	if radioGroupMember != nil {
		cradio = radioGroupMember.native()
	}
	c := C.gtk_radio_button_new_with_mnemonic_from_widget(cradio,
		(*C.gchar)(cstr))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapRadioButton(obj), nil
}

// SetGroup is a wrapper around gtk_radio_button_set_group().
func (v *RadioButton) SetGroup(group *glib.SList) {
	C.gtk_radio_button_set_group(v.native(), cGSList(group))
}

// GetGroup is a wrapper around gtk_radio_button_get_group().
func (v *RadioButton) GetGroup() (*glib.SList, error) {
	c := C.gtk_radio_button_get_group(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	return glib.WrapSList(uintptr(unsafe.Pointer(c))), nil
}

// JoinGroup is a wrapper around gtk_radio_button_join_group().
func (v *RadioButton) JoinGroup(groupSource *RadioButton) {
	var cgroup *C.GtkRadioButton
	if groupSource != nil {
		cgroup = groupSource.native()
	}
	C.gtk_radio_button_join_group(v.native(), cgroup)
}

/*
 * GtkToggleButton
 */

// ToggleButton is a representation of GTK's GtkToggleButton.
type ToggleButton struct {
	Button
}

// native returns a pointer to the underlying GtkToggleButton.
func (v *ToggleButton) native() *C.GtkToggleButton {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toGtkToggleButton(ptr)
}

func marshalToggleButton(p uintptr) (interface{}, error) {
	c := C.g_value_get_object(C.toGValue(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapToggleButton(obj), nil
}

func wrapToggleButton(obj *glib.Object) *ToggleButton {
	button := wrapButton(obj)
	return &ToggleButton{*button}
}

// ToggleButtonNew is a wrapper around gtk_toggle_button_new().
func ToggleButtonNew() (*ToggleButton, error) {
	c := C.gtk_toggle_button_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapToggleButton(obj), nil
}

// ToggleButtonNewWithLabel is a wrapper around
// gtk_toggle_button_new_with_label().
func ToggleButtonNewWithLabel(label string) (*ToggleButton, error) {
	cstr := C.CString(label)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_toggle_button_new_with_label((*C.gchar)(cstr))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapToggleButton(obj), nil
}

// ToggleButtonNewWithMnemonic is a wrapper around
// gtk_toggle_button_new_with_mnemonic().
func ToggleButtonNewWithMnemonic(label string) (*ToggleButton, error) {
	cstr := C.CString(label)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_toggle_button_new_with_mnemonic((*C.gchar)(cstr))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapToggleButton(obj), nil
}

// GetActive is a wrapper around gtk_toggle_button_get_active().
func (v *ToggleButton) GetActive() bool {
	c := C.gtk_toggle_button_get_active(v.native())
	return gobool(c)
}

// SetActive is a wrapper around gtk_toggle_button_set_active().
func (v *ToggleButton) SetActive(isActive bool) {
	C.gtk_toggle_button_set_active(v.native(), gbool(isActive))
}

/*
 * GtkLinkButton
 */

// LinkButton is a representation of GTK's GtkLinkButton.
type LinkButton struct {
	Button
}

// native returns a pointer to the underlying GtkLinkButton.
func (v *LinkButton) native() *C.GtkLinkButton {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toGtkLinkButton(ptr)
}

func marshalLinkButton(p uintptr) (interface{}, error) {
	c := C.g_value_get_object(C.toGValue(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapLinkButton(obj), nil
}

func wrapLinkButton(obj *glib.Object) *LinkButton {
	button := wrapButton(obj)
	return &LinkButton{*button}
}

// LinkButtonNew is a wrapper around gtk_link_button_new().
func LinkButtonNew(label string) (*LinkButton, error) {
	cstr := C.CString(label)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_link_button_new((*C.gchar)(cstr))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapLinkButton(obj), nil
}

// LinkButtonNewWithLabel is a wrapper around gtk_link_button_new_with_label().
func LinkButtonNewWithLabel(uri, label string) (*LinkButton, error) {
	curi := C.CString(uri)
	defer C.free(unsafe.Pointer(curi))
	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	c := C.gtk_link_button_new_with_label((*C.gchar)(curi), (*C.gchar)(clabel))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapLinkButton(obj), nil
}

// GetUri is a wrapper around gtk_link_button_get_uri().
func (v *LinkButton) GetUri() string {
	c := C.gtk_link_button_get_uri(v.native())
	return goString(c)
}

// SetUri is a wrapper around gtk_link_button_set_uri().
func (v *LinkButton) SetUri(uri string) {
	cstr := C.CString(uri)
	C.gtk_link_button_set_uri(v.native(), (*C.gchar)(cstr))
}

/*
 * GtkMenuButton
 */

// MenuButton is a representation of GTK's GtkMenuButton.
type MenuButton struct {
	ToggleButton
}

// native returns a pointer to the underlying GtkMenuButton.
func (v *MenuButton) native() *C.GtkMenuButton {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toGtkMenuButton(ptr)
}

func marshalMenuButton(p uintptr) (interface{}, error) {
	c := C.g_value_get_object(C.toGValue(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapMenuButton(obj), nil
}

func wrapMenuButton(obj *glib.Object) *MenuButton {
	toggleButton := wrapToggleButton(obj)
	return &MenuButton{*toggleButton}
}

// MenuButtonNew is a wrapper around gtk_menu_button_new().
func MenuButtonNew() (*MenuButton, error) {
	c := C.gtk_menu_button_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapMenuButton(obj), nil
}

// SetPopup is a wrapper around gtk_menu_button_set_popup().
func (v *MenuButton) SetPopup(menu IMenu) {
	C.gtk_menu_button_set_popup(v.native(), menu.toWidget())
}

// GetPopup is a wrapper around gtk_menu_button_get_popup().
func (v *MenuButton) GetPopup() *Menu {
	c := C.gtk_menu_button_get_popup(v.native())
	if c == nil {
		return nil
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapMenu(obj)
}

// void
// gtk_menu_button_set_menu_model (GtkMenuButton *menu_button,
//                                 GMenuModel *menu_model);
func (v *MenuButton) SetMenuModel(model glib.IMenuModel) {
	C.gtk_menu_button_set_menu_model(v.native(),
		C.toGMenuModel(unsafe.Pointer(model.Native())))
}

// SetDirection is a wrapper around gtk_menu_button_set_direction().
func (v *MenuButton) SetDirection(direction ArrowType) {
	C.gtk_menu_button_set_direction(v.native(), C.GtkArrowType(direction))
}

// GetDirection is a wrapper around gtk_menu_button_get_direction().
func (v *MenuButton) GetDirection() ArrowType {
	c := C.gtk_menu_button_get_direction(v.native())
	return ArrowType(c)
}

// SetAlignWidget is a wrapper around gtk_menu_button_set_align_widget().
func (v *MenuButton) SetAlignWidget(alignWidget IWidget) {
	C.gtk_menu_button_set_align_widget(v.native(), alignWidget.toWidget())
}

// GetAlignWidget is a wrapper around gtk_menu_button_get_align_widget().
func (v *MenuButton) GetAlignWidget() *Widget {
	c := C.gtk_menu_button_get_align_widget(v.native())
	if c == nil {
		return nil
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapWidget(obj)
}

/*
 * GtkSwitch
 */

// Switch is a representation of GTK's GtkSwitch.
type Switch struct {
	Widget
	// Interfaces
	Actionable
}

// native returns a pointer to the underlying GtkSwitch.
func (v *Switch) native() *C.GtkSwitch {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toGtkSwitch(ptr)
}

func marshalSwitch(p uintptr) (interface{}, error) {
	c := C.g_value_get_object(C.toGValue(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapSwitch(obj), nil
}

func wrapSwitch(obj *glib.Object) *Switch {
	widget := wrapWidget(obj)
	actionable := wrapActionable(*glib.InterfaceFromObjectNew(obj))
	return &Switch{*widget, *actionable}
}

// SwitchNew is a wrapper around gtk_switch_new().
func SwitchNew() (*Switch, error) {
	c := C.gtk_switch_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapSwitch(obj), nil
}

// GetActive is a wrapper around gtk_switch_get_active().
func (v *Switch) GetActive() bool {
	c := C.gtk_switch_get_active(v.native())
	return gobool(c)
}

// SetActive is a wrapper around gtk_switch_set_active().
func (v *Switch) SetActive(isActive bool) {
	C.gtk_switch_set_active(v.native(), gbool(isActive))
}

/*
 * GtkScaleButton
 */

// ScaleButton is a representation of GTK's GtkScaleButton.
type ScaleButton struct {
	Button
}

// native() returns a pointer to the underlying GtkScaleButton.
func (v *ScaleButton) native() *C.GtkScaleButton {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toGtkScaleButton(ptr)
}

func marshalScaleButton(p uintptr) (interface{}, error) {
	c := C.g_value_get_object(C.toGValue(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapScaleButton(obj), nil
}

func wrapScaleButton(obj *glib.Object) *ScaleButton {
	button := wrapButton(obj)
	return &ScaleButton{*button}
}

// ScaleButtonNew() is a wrapper around gtk_scale_button_new().
func ScaleButtonNew(size IconSize, min, max, step float64, icons []string) (*ScaleButton, error) {
	cicons := make([]*C.gchar, len(icons))
	for i, icon := range icons {
		cicons[i] = (*C.gchar)(C.CString(icon))
		defer C.free(unsafe.Pointer(cicons[i]))
	}
	cicons = append(cicons, nil)

	c := C.gtk_scale_button_new(C.GtkIconSize(size),
		C.gdouble(min),
		C.gdouble(max),
		C.gdouble(step),
		&cicons[0])
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapScaleButton(obj), nil
}

// GetAdjustment() is a wrapper around gtk_scale_button_get_adjustment().
func (v *ScaleButton) GetAdjustment() (*Adjustment, error) {
	c := C.gtk_scale_button_get_adjustment(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapAdjustment(obj), nil
}

// GetPopup() is a wrapper around gtk_scale_button_get_popup().
func (v *ScaleButton) GetPopup() (*Widget, error) {
	c := C.gtk_scale_button_get_popup(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapWidget(obj), nil
}

// GetValue() is a wrapper around gtk_scale_button_get_value().
func (v *ScaleButton) GetValue() float64 {
	return float64(C.gtk_scale_button_get_value(v.native()))
}

// SetAdjustment() is a wrapper around gtk_scale_button_set_adjustment().
func (v *ScaleButton) SetAdjustment(adjustment *Adjustment) {
	C.gtk_scale_button_set_adjustment(v.native(), adjustment.native())
}

// SetValue() is a wrapper around gtk_scale_button_set_value().
func (v *ScaleButton) SetValue(value float64) {
	C.gtk_scale_button_set_value(v.native(), C.gdouble(value))
}

/*
 * GtkVolumeButton
 */

// VolumeButton is a representation of GTK's GtkVolumeButton.
type VolumeButton struct {
	ScaleButton
}

// native() returns a pointer to the underlying GtkVolumeButton.
func (v *VolumeButton) native() *C.GtkVolumeButton {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toGtkVolumeButton(ptr)
}

func marshalVolumeButton(p uintptr) (interface{}, error) {
	c := C.g_value_get_object(C.toGValue(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapVolumeButton(obj), nil
}

func wrapVolumeButton(obj *glib.Object) *VolumeButton {
	scaleButton := wrapScaleButton(obj)
	return &VolumeButton{*scaleButton}
}

// VolumeButtonNew() is a wrapper around gtk_button_new().
func VolumeButtonNew() (*VolumeButton, error) {
	c := C.gtk_volume_button_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapVolumeButton(obj), nil
}

// TODO: implement GtkLockButton and GtkModelButton
