// This file includes wrappers for symbols included since GTK 3.12, and
// and should not be included in a build intended to target any older GTK
// versions.  To target an older build, such as 3.10, use
// 'go build -tags gtk_3_10'.  Otherwise, if no build tags are used, GTK 3.12
// is assumed and this file is built.
// +build !gtk_3_6,!gtk_3_8,!gtk_3_10

package main

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
	"sync"
	"time"

	"github.com/romychs/gotk3/gdk"
	"github.com/romychs/gotk3/glib"
	"github.com/romychs/gotk3/gtk"
	"github.com/romychs/gotk3/pango"
	"github.com/davecgh/go-spew/spew"
)

// ========================================================================================
// ************************* GTK+ UI UTILITIES SECTION START ******************************
// ========================================================================================
//	In real application use this code section as utilities to simplify creation
//	of GLIB/GTK+ components and widgets, including menus, dialog boxes, messages,
//	application settings and so on...

// SetupLabelJustifyRight create GtkLabel with justification to the right by default.
func SetupLabelJustifyRight(caption string) (*gtk.Label, error) {
	lbl, err := gtk.LabelNew(caption)
	if err != nil {
		return nil, err
	}
	lbl.SetHAlign(gtk.ALIGN_END)
	lbl.SetJustify(gtk.JUSTIFY_RIGHT)
	return lbl, nil
}

// SetupLabelJustifyLeft create GtkLabel with justification to the left by default.
func SetupLabelJustifyLeft(caption string) (*gtk.Label, error) {
	lbl, err := gtk.LabelNew(caption)
	if err != nil {
		return nil, err
	}
	lbl.SetHAlign(gtk.ALIGN_START)
	lbl.SetJustify(gtk.JUSTIFY_LEFT)
	return lbl, nil
}

// SetupLabelJustifyCenter create GtkLabel with justification to the center by default.
func SetupLabelJustifyCenter(caption string) (*gtk.Label, error) {
	lbl, err := gtk.LabelNew(caption)
	if err != nil {
		return nil, err
	}
	lbl.SetHAlign(gtk.ALIGN_CENTER)
	lbl.SetJustify(gtk.JUSTIFY_CENTER)
	return lbl, nil
}

// SetupHeader construct Header widget with standard initialization.
func SetupHeader(title, subtitle string, showCloseButton bool) (*gtk.HeaderBar, error) {
	hdr, err := gtk.HeaderBarNew()
	if err != nil {
		return nil, err
	}
	hdr.SetShowCloseButton(showCloseButton)
	hdr.SetTitle(title)
	if subtitle != "" {
		hdr.SetSubtitle(subtitle)
	}
	return hdr, nil
}

// SetupMenuItemWithIcon construct MenuItem widget with icon image.
func SetupMenuItemWithIcon(label, detailedAction string, icon *glib.Icon) (*glib.MenuItem, error) {
	mi, err := glib.MenuItemNew(label, detailedAction)
	if err != nil {
		return nil, err
	}
	//mi.SetAttributeValue("verb-icon", iconNameVar)
	mi.SetIcon(icon)
	return mi, nil
}

// SetupMenuItemWithThemedIcon construct MenuItem widget with image
// taken by iconName from themed icons image lib.
func SetupMenuItemWithThemedIcon(label, detailedAction, iconName string) (*glib.MenuItem, error) {
	iconNameVar, err := glib.VariantStringNew(iconName)
	if err != nil {
		return nil, err
	}
	mi, err := glib.MenuItemNew(label, detailedAction)
	if err != nil {
		return nil, err
	}
	mi.SetAttributeValue("verb-icon", iconNameVar)
	return mi, nil
}

// SetupToolButton construct ToolButton widget with standart initialization.
func SetupToolButton(themedIconName, label string) (*gtk.ToolButton, error) {
	var btn *gtk.ToolButton
	var img *gtk.Image
	var err error
	if themedIconName != "" {
		img, err = gtk.ImageNewFromIconName(themedIconName, gtk.ICON_SIZE_BUTTON)
		if err != nil {
			return nil, err
		}
	}

	btn, err = gtk.ToolButtonNew(img, label)
	if err != nil {
		return nil, err
	}
	return btn, nil
}

// SetupButtonWithThemedImage construct Button widget with image
// taken by themedIconName from themed icons image lib.
func SetupButtonWithThemedImage(themedIconName string) (*gtk.Button, error) {
	img, err := gtk.ImageNewFromIconName(themedIconName, gtk.ICON_SIZE_BUTTON)
	if err != nil {
		return nil, err
	}

	btn, err := gtk.ButtonNew()
	if err != nil {
		return nil, err
	}

	btn.Add(img)

	return btn, nil
}

// getPixbufFromBytes create gdk.PixBuf loaded from raw bytes buffer
func getPixbufFromBytes(bytes []byte) (*gdk.Pixbuf, error) {
	b2, err := glib.BytesNew(bytes)
	if err != nil {
		return nil, err
	}
	pbl, err := gdk.PixbufLoaderNew()
	if err != nil {
		return nil, err
	}
	err = pbl.WriteBytes(b2)
	if err != nil {
		return nil, err
	}
	err = pbl.Close()
	if err != nil {
		return nil, err
	}

	pb, err := pbl.GetPixbuf()
	if err != nil {
		return nil, err
	}
	return pb, nil
}

// getPixbufFromBytesWithResize create gdk.PixBuf loaded from raw bytes buffer, applying resize
func getPixbufFromBytesWithResize(bytes []byte, resizeToWidth, resizeToHeight int) (*gdk.Pixbuf, error) {
	b2, err := glib.BytesNew(bytes)
	if err != nil {
		return nil, err
	}
	pbl, err := gdk.PixbufLoaderNew()
	if err != nil {
		return nil, err
	}
	pbl.SetSize(resizeToWidth, resizeToHeight)
	err = pbl.WriteBytes(b2)
	if err != nil {
		return nil, err
	}
	err = pbl.Close()
	if err != nil {
		return nil, err
	}

	pb, err := pbl.GetPixbuf()
	if err != nil {
		return nil, err
	}
	return pb, nil
}

// getPixbufFromBytes create gdk.PixbufAnimation loaded from raw bytes buffer
func getPixbufAnimationFromBytes(bytes []byte) (*gdk.PixbufAnimation, error) {
	b2, err := glib.BytesNew(bytes)
	if err != nil {
		return nil, err
	}
	pbl, err := gdk.PixbufLoaderNew()
	if err != nil {
		return nil, err
	}
	err = pbl.WriteBytes(b2)
	if err != nil {
		return nil, err
	}
	err = pbl.Close()
	if err != nil {
		return nil, err
	}

	pba, err := pbl.GetPixbufAnimation()
	if err != nil {
		return nil, err
	}
	return pba, nil
}

// getPixbufFromBytesWithResize create gdk.PixbufAnimation loaded from raw bytes buffer, applying resize
func getPixbufAnimationFromBytesWithResize(bytes []byte, resizeToWidth,
	resizeToHeight int) (*gdk.PixbufAnimation, error) {

	b2, err := glib.BytesNew(bytes)
	if err != nil {
		return nil, err
	}
	pbl, err := gdk.PixbufLoaderNew()
	if err != nil {
		return nil, err
	}
	pbl.SetSize(resizeToWidth, resizeToHeight)
	err = pbl.WriteBytes(b2)
	if err != nil {
		return nil, err
	}
	err = pbl.Close()
	if err != nil {
		return nil, err
	}

	pba, err := pbl.GetPixbufAnimation()
	if err != nil {
		return nil, err
	}
	return pba, nil
}

// SetupMenuButtonWithThemedImage construct gtk.MenuButton widget with image
// taken by themedIconName from themed icons image lib.
func SetupMenuButtonWithThemedImage(themedIconName string) (*gtk.MenuButton, error) {
	img, err := gtk.ImageNewFromIconName(themedIconName, gtk.ICON_SIZE_BUTTON)
	if err != nil {
		return nil, err
	}

	btn, err := gtk.MenuButtonNew()
	if err != nil {
		return nil, err
	}

	btn.Add(img)

	return btn, nil
}

// AppendSectionAsHorzButtons used for gtk.Popover widget menu
// as a hint to display items as a horizontal buttons.
func AppendSectionAsHorzButtons(main, section *glib.Menu) error {
	val1, err := glib.VariantStringNew("horizontal-buttons")
	if err != nil {
		return err
	}
	mi1, err := glib.MenuItemNew("", "")
	if err != nil {
		return err
	}
	mi1.SetSection(section)
	mi1.SetAttributeValue("display-hint", val1)
	main.AppendItem(mi1)
	//section.AppendItem(mi1)
	return nil
}

// DialogButton simplify dialog window initialization.
// Keep all necessary information about how attached
// dialog button should look and act.
type DialogButton struct {
	Text      string
	Response  gtk.ResponseType
	Default   bool
	Customize func(button *gtk.Button) error
}

// GetActiveWindow find real active window in application running.
func GetActiveWindow(win *gtk.Window) (*gtk.Window, error) {
	app, err := win.GetApplication()
	if err != nil {
		return nil, err
	}
	return app.GetActiveWindow(), nil
}

// IsResponseYes gives true if dialog window
// responded with gtk.RESPONSE_YES.
func IsResponseYes(response gtk.ResponseType) bool {
	return response == gtk.RESPONSE_YES
}

// IsResponseNo gives true if dialog window
// responded with gtk.RESPONSE_NO.
func IsResponseNo(response gtk.ResponseType) bool {
	return response == gtk.RESPONSE_NO
}

// IsResponseNone gives true if dialog window
// responded with gtk.RESPONSE_NONE.
func IsResponseNone(response gtk.ResponseType) bool {
	return response == gtk.RESPONSE_NONE
}

// IsResponseOk gives true if dialog window
// responded with gtk.RESPONSE_OK.
func IsResponseOk(response gtk.ResponseType) bool {
	return response == gtk.RESPONSE_OK
}

// IsResponseCancel gives true if dialog window
// responded with gtk.RESPONSE_CANCEL.
func IsResponseCancel(response gtk.ResponseType) bool {
	return response == gtk.RESPONSE_CANCEL
}

// IsResponseReject gives true if dialog window
// responded with gtk.RESPONSE_REJECT.
func IsResponseReject(response gtk.ResponseType) bool {
	return response == gtk.RESPONSE_REJECT
}

// IsResponseClose gives true if dialog window
// responded with gtk.RESPONSE_CLOSE.
func IsResponseClose(response gtk.ResponseType) bool {
	return response == gtk.RESPONSE_CLOSE
}

// IsResponseDeleteEvent gives true if dialog window
// responded with gtk.RESPONSE_DELETE_EVENT.
func IsResponseDeleteEvent(response gtk.ResponseType) bool {
	return response == gtk.RESPONSE_DELETE_EVENT
}

// DialogParagraph is an object which keep text paragraph added
// to dialog window, complemented with all necessary format options.
type DialogParagraph struct {
	Text          string
	Markup        bool
	HorizAlign    gtk.Align
	Justify       gtk.Justification
	Ellipsize     pango.EllipsizeMode
	MaxWidthChars int
}

// NewDialogParagraph create new text paragraph instance,
// with default align, justification and so on.
func NewDialogParagraph(text string) *DialogParagraph {
	v := &DialogParagraph{Text: text, HorizAlign: gtk.Align(-1), Justify: gtk.Justification(-1),
		Ellipsize: pango.EllipsizeMode(-1), MaxWidthChars: -1}
	return v
}

// SetMarkup update Markup flag.
func (v *DialogParagraph) SetMarkup(markup bool) *DialogParagraph {
	v.Markup = markup
	return v
}

// SetHorizAlign set horizontal alignment of text paragraph.
func (v *DialogParagraph) SetHorizAlign(align gtk.Align) *DialogParagraph {
	v.HorizAlign = align
	return v
}

// SetJustify set text justification.
func (v *DialogParagraph) SetJustify(justify gtk.Justification) *DialogParagraph {
	v.Justify = justify
	return v
}

// SetEllipsize set text ellipsis mode.
func (v *DialogParagraph) SetEllipsize(ellipsize pango.EllipsizeMode) *DialogParagraph {
	v.Ellipsize = ellipsize
	return v
}

// SetMaxWidthChars set maximum number of chars in one line.
func (v *DialogParagraph) SetMaxWidthChars(maxWidthChars int) *DialogParagraph {
	v.MaxWidthChars = maxWidthChars
	return v
}

// createLabel create gtk.Label widget to put paragraph text in.
func (v *DialogParagraph) createLabel() (*gtk.Label, error) {
	lbl, err := gtk.LabelNew("")
	if err != nil {
		return nil, err
	}
	if v.Markup {
		lbl.SetMarkup(v.Text)
	} else {
		lbl.SetText(v.Text)
	}
	if v.HorizAlign != gtk.Align(-1) {
		lbl.SetHAlign(v.HorizAlign)
	}
	if v.Justify != gtk.Justification(-1) {
		lbl.SetJustify(v.Justify)
	}
	if v.Ellipsize != pango.EllipsizeMode(-1) {
		lbl.SetEllipsize(v.Ellipsize)
	}
	if v.MaxWidthChars != -1 {
		lbl.SetMaxWidthChars(v.MaxWidthChars)
	}
	return lbl, nil
}

// TextToDialogParagraphs multi-line text to DialogParagraph instance.
func TextToDialogParagraphs(lines []string) []*DialogParagraph {
	var msgs []*DialogParagraph
	for _, line := range lines {
		msgs = append(msgs, NewDialogParagraph(line))
	}
	return msgs
}

// TextToMarkupDialogParagraphs multi-line markup text to DialogParagraph instance.
func TextToMarkupDialogParagraphs(makrupLines []string) []*DialogParagraph {
	var msgs []*DialogParagraph
	for _, markupLine := range makrupLines {
		msgs = append(msgs, NewDialogParagraph(markupLine).SetMarkup(true))
	}
	return msgs
}

type MessageDialog struct {
	dialog *gtk.MessageDialog
}

// SetupMessageDialog construct MessageDialog widget with customized settings.
func SetupMessageDialog(parent *gtk.Window, markupTitle string, secondaryMarkupTitle string,
	paragraphs []*DialogParagraph, addButtons []DialogButton,
	addExtraControls func(area *gtk.Box) error) (*MessageDialog, error) {

	var active *gtk.Window
	var err error

	if parent != nil {
		active, err = GetActiveWindow(parent)
		if err != nil {
			return nil, err
		}
	}

	dlg, err := gtk.MessageDialogNew(active, /*gtk.DIALOG_MODAL|*/
		gtk.DIALOG_USE_HEADER_BAR, gtk.MESSAGE_WARNING, gtk.BUTTONS_NONE, nil, nil)
	if err != nil {
		return nil, err
	}
	if active != nil {
		dlg.SetTransientFor(active)
	}
	dlg.SetMarkup(markupTitle)
	if secondaryMarkupTitle != "" {
		dlg.FormatSecondaryMarkup(secondaryMarkupTitle)
	}

	for _, button := range addButtons {
		btn, err := dlg.AddButton(button.Text, button.Response)
		if err != nil {
			return nil, err
		}

		if button.Default {
			dlg.SetDefaultResponse(button.Response)
		}

		if button.Customize != nil {
			err := button.Customize(btn)
			if err != nil {
				return nil, err
			}
		}
	}

	grid, err := gtk.GridNew()
	if err != nil {
		return nil, err
	}

	grid.SetRowSpacing(6)
	grid.SetHAlign(gtk.ALIGN_CENTER)

	box, err := dlg.GetMessageArea()
	if err != nil {
		return nil, err
	}
	box.Add(grid)

	col := 1
	row := 0

	// add empty line after title
	paragraphs = append([]*DialogParagraph{NewDialogParagraph("")}, paragraphs...)

	for _, paragraph := range paragraphs {
		lbl, err := paragraph.createLabel()
		if err != nil {
			return nil, err
		}
		grid.Attach(lbl, col, row, 1, 1)
		row++
	}

	if addExtraControls != nil {
		box1, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 6)
		if err != nil {
			return nil, err
		}
		grid.Attach(box1, col, row, 1, 1)

		err = addExtraControls(box1)
		if err != nil {
			return nil, err
		}
	}

	box.ShowAll()

	v := &MessageDialog{dialog: dlg}
	return v, nil
}

// Run run MessageDialog widget with customized settings.
func (v *MessageDialog) Run(ignoreCloseBox bool) gtk.ResponseType {

	defer v.dialog.Destroy()

	v.dialog.ShowAll()
	var res gtk.ResponseType
	res = v.dialog.Run()
	for gtk.ResponseType(res) == gtk.RESPONSE_NONE || gtk.ResponseType(res) == gtk.RESPONSE_DELETE_EVENT && ignoreCloseBox {
		res = v.dialog.Run()
	}
	return gtk.ResponseType(res)
}

// SetupDialog construct Dialog widget with customized settings.
func SetupDialog(parent *gtk.Window, messageType gtk.MessageType, userHeaderbar bool,
	title string, paragraphs []*DialogParagraph, addButtons []DialogButton,
	addExtraControls func(area *gtk.Box) error) (*gtk.Dialog, error) {

	var active *gtk.Window
	var err error

	if parent != nil {
		active, err = GetActiveWindow(parent)
		if err != nil {
			return nil, err
		}
	}

	flags := gtk.DIALOG_MODAL
	if userHeaderbar {
		flags |= gtk.DIALOG_USE_HEADER_BAR
	}
	dlg, err := gtk.DialogWithFlagsNew(title, active, flags)
	if err != nil {
		return nil, err
	}

	dlg.SetDefaultSize(100, 100)
	dlg.SetTransientFor(active)
	dlg.SetDeletable(false)

	var img *gtk.Image
	size := gtk.ICON_SIZE_DIALOG
	if userHeaderbar {
		size = gtk.ICON_SIZE_LARGE_TOOLBAR
	}
	var iconName string
	switch messageType {
	case gtk.MESSAGE_WARNING:
		iconName = "dialog-warning"
	case gtk.MESSAGE_ERROR:
		iconName = "dialog-error"
	case gtk.MESSAGE_INFO:
		iconName = "dialog-information"
	case gtk.MESSAGE_QUESTION:
		iconName = "dialog-question"
	}

	if iconName != "" {
		img, err = gtk.ImageNewFromIconName(iconName, size)
		if err != nil {
			return nil, err
		}
	}

	grid, err := gtk.GridNew()
	if err != nil {
		return nil, err
	}

	grid.SetBorderWidth(10)
	grid.SetColumnSpacing(10)
	grid.SetRowSpacing(6)
	grid.SetHAlign(gtk.ALIGN_CENTER)

	box, err := dlg.GetContentArea()
	if err != nil {
		return nil, err
	}
	box.Add(grid)

	if img != nil {
		if userHeaderbar {
			hdr, err := dlg.GetHeaderBar()
			if err != nil {
				return nil, err
			}

			hdr.PackStart(img)
		} else {
			grid.Attach(img, 0, 0, 1, 1)
		}
	}

	for _, button := range addButtons {
		btn, err := dlg.AddButton(button.Text, button.Response)
		if err != nil {
			return nil, err
		}

		if button.Default {
			dlg.SetDefaultResponse(button.Response)
		}

		if button.Customize != nil {
			err := button.Customize(btn)
			if err != nil {
				return nil, err
			}
		}
	}

	col := 1
	row := 0

	for _, paragraph := range paragraphs {
		lbl, err := paragraph.createLabel()
		if err != nil {
			return nil, err
		}
		grid.Attach(lbl, col, row, 1, 1)
		row++
	}

	if addExtraControls != nil {
		box1, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 6)
		if err != nil {
			return nil, err
		}
		grid.Attach(box1, col, row, 1, 1)

		err = addExtraControls(box1)
		if err != nil {
			return nil, err
		}
	}

	_, w := dlg.GetPreferredWidth()
	_, h := dlg.GetPreferredHeight()
	dlg.Resize(w, h)

	return dlg, nil
}

// RunDialog construct and run Dialog widget with customized settings.
func RunDialog(parent *gtk.Window, messageType gtk.MessageType, userHeaderbar bool,
	title string, paragraphs []*DialogParagraph, ignoreCloseBox bool, addButtons []DialogButton,
	addExtraControls func(area *gtk.Box) error) (gtk.ResponseType, error) {

	dlg, err := SetupDialog(parent, messageType, userHeaderbar, title,
		paragraphs, addButtons, addExtraControls)
	if err != nil {
		return 0, err
	}
	defer dlg.Destroy()

	//dlg.ShowAll()
	dlg.ShowAll()
	res := dlg.Run()
	for gtk.ResponseType(res) == gtk.RESPONSE_NONE || gtk.ResponseType(res) == gtk.RESPONSE_DELETE_EVENT && ignoreCloseBox {
		res = dlg.Run()
	}
	return gtk.ResponseType(res), nil
}

// ErrorMessage build and run error message dialog.
func ErrorMessage(parent *gtk.Window, titleMarkup string, text []*DialogParagraph) error {
	buttons := []DialogButton{
		{"_OK", gtk.RESPONSE_OK, false, nil},
	}
	dialog, err := SetupMessageDialog(parent, titleMarkup, "", text, buttons, nil)
	if err != nil {
		return err
	}
	dialog.Run(false)
	return nil
}

// QuestionDialog build and run question message dialog with Yes/No choice.
func QuestionDialog(parent *gtk.Window, title string,
	messages []*DialogParagraph, defaultYes bool) (bool, error) {

	title2 := spew.Sprintf("%s", title)
	buttons := []DialogButton{
		{"_YES", gtk.RESPONSE_YES, defaultYes, nil},
		{"_NO", gtk.RESPONSE_NO, !defaultYes, nil},
	}
	response, err := RunDialog(parent, gtk.MESSAGE_QUESTION, true, title2,
		messages, false, buttons, nil)
	if err != nil {
		return false, err
	}
	return IsResponseYes(response), nil
}

// GetActionNameAndState display status of action-with-state, which used in
// menu-with-state behavior. Convenient for debug purpose.
func GetActionNameAndState(act *glib.SimpleAction) (string, *glib.Variant, error) {
	name, err := act.GetName()
	if err != nil {
		return "", nil, err
	}
	state := act.GetState()
	return name, state, nil
}

// SetMargins set margins of a widget to the passed values,
// replacing 4 calls with only one.
func SetMargins(widget gtk.IWidget, left int, top int, right int, bottom int) {
	w := widget.GetWidget()
	w.SetMarginStart(left)
	w.SetMarginTop(top)
	w.SetMarginEnd(right)
	w.SetMarginBottom(bottom)
}

// SetAllMargins set all margins of a widget to the same value.
func SetAllMargins(widget gtk.IWidget, margin int) {
	SetMargins(widget, margin, margin, margin, margin)
}

// AppendValues append multiple values to a row in a list store.
func AppendValues(ls *gtk.ListStore, values ...interface{}) (*gtk.TreeIter, error) {
	iter := ls.Append()
	for i := 0; i < len(values); i++ {
		err := ls.SetValue(iter, i, values[i])
		if err != nil {
			return nil, err
		}
	}
	return iter, nil
}

// CreateNameValueCombo create a GtkComboBox that holds
// a set of name/value pairs where the name is displayed.
func CreateNameValueCombo(keyValues []struct{ value, key string }) (*gtk.ComboBox, error) {
	ls, err := gtk.ListStoreNew(glib.TYPE_STRING, glib.TYPE_STRING)
	if err != nil {
		return nil, err
	}

	for _, item := range keyValues {
		_, err = AppendValues(ls, item.value, item.key)
		if err != nil {
			return nil, err
		}
	}

	cb, err := gtk.ComboBoxNew()
	if err != nil {
		return nil, err
	}
	err = UpdateNameValueCombo(cb, keyValues)
	if err != nil {
		return nil, err
	}
	cb.SetFocusOnClick(false)
	cb.SetIDColumn(1)
	cell, err := gtk.CellRendererTextNew()
	if err != nil {
		return nil, err
	}
	cell.SetAlignment(0, 0)
	cb.PackStart(cell, false)
	cb.AddAttribute(cell, "text", 0)

	return cb, nil
}

// UpdateNameValueCombo update GtkComboBox list of name/value pairs.
func UpdateNameValueCombo(cb *gtk.ComboBox, keyValues []struct{ value, key string }) error {
	ls, err := gtk.ListStoreNew(glib.TYPE_STRING, glib.TYPE_STRING)
	if err != nil {
		return err
	}

	for _, item := range keyValues {
		_, err = AppendValues(ls, item.value, item.key)
		if err != nil {
			return err
		}
	}

	cb.SetModel(ls)
	return nil
}

// GetComboValue return GtkComboBox selected value from specific column.
func GetComboValue(cb *gtk.ComboBox, columnID int) (*glib.Value, error) {
	ti, err := cb.GetActiveIter()
	if err != nil {
		return nil, err
	}
	tm, err := cb.GetModel()
	if err != nil {
		return nil, err
	}
	val, err := tm.GetValue(ti, 0)
	if err != nil {
		return nil, err
	}
	return val, nil
}

// GetGtkVersion return actually installed GTK+ version.
func GetGtkVersion() (magor, minor, micro uint) {
	magor = gtk.GetMajorVersion()
	minor = gtk.GetMinorVersion()
	micro = gtk.GetMicroVersion()
	return
}

// GetGlibVersion return actually installed GLIB version.
func GetGlibVersion() (magor, minor, micro uint) {
	magor = glib.GetMajorVersion()
	minor = glib.GetMinorVersion()
	micro = glib.GetMicroVersion()
	return
}

// GetGdkVersion return actually installed GDK version.
func GetGdkVersion() (magor, minor, micro uint) {
	magor = gdk.GetMajorVersion()
	minor = gdk.GetMinorVersion()
	micro = gdk.GetMicroVersion()
	return
}

// ApplyStyleCSS apply custom CSS to specific widget.
func ApplyStyleCSS(widget *gtk.Widget, css string) error {
	//	provider, err := gtk.CssProviderNew()
	provider, err := gtk.CssProviderNew()
	if err != nil {
		return err
	}
	err = provider.LoadFromData(css)
	if err != nil {
		return err
	}
	sc, err := widget.GetStyleContext()
	if err != nil {
		return err
	}
	sc.AddProvider(provider, gtk.STYLE_PROVIDER_PRIORITY_USER)
	return nil
}

// AddStyleClasses apply specific CSS style classes to the widget.
func AddStyleClasses(widget *gtk.Widget, cssClasses []string) error {
	sc, err := widget.GetStyleContext()
	if err != nil {
		return err
	}
	for _, className := range cssClasses {
		sc.AddClass(className)
	}
	return nil
}

// AddStyleClass apply specific CSS style class to the widget.
func AddStyleClass(widget *gtk.Widget, cssClass string) error {
	sc, err := widget.GetStyleContext()
	if err != nil {
		return err
	}
	sc.AddClass(cssClass)
	return nil
}

// RemoveStyleClass remove specific CSS style class from the widget.
func RemoveStyleClass(widget *gtk.Widget, cssClass string) error {
	sc, err := widget.GetStyleContext()
	if err != nil {
		return err
	}
	sc.RemoveClass(cssClass)
	return nil
}

// RemoveStyleClasses remove specific CSS style classes from the widget.
func RemoveStyleClasses(widget *gtk.Widget, cssClasses []string) error {
	sc, err := widget.GetStyleContext()
	if err != nil {
		return err
	}
	for _, className := range cssClasses {
		sc.RemoveClass(className)
	}
	return nil
}

// RemoveStyleClassesAll remove all style classes from the widget.
func RemoveStyleClassesAll(widget *gtk.Widget) error {
	sc, err := widget.GetStyleContext()
	if err != nil {
		return err
	}
	list := sc.ListClasses()
	list.Foreach(func(item interface{}) {
		cssClass := item.(string)
		sc.RemoveClass(cssClass)
	})
	return nil
}

// ========================================================================================
// ************************* GTK+ UI UTILITIES SECTION END ********************************
// ========================================================================================

// ==========================================================================================
// ************************* GLIB SETTINGS UTILITIES SECTION START **************************
// ==========================================================================================
//	In real application use this code section as utilities to simplify creation
//	of GLIB/GTK+ components and widgets, including menus, dialog boxes, messages,
//	application settings and so on...

// SettingsStore simplify work with glib.Settings.
type SettingsStore struct {
	settings *glib.Settings
	schemaID string
	path     string
}

// removeExcessSlashChars normalize path and remove excess path divider in glib.Settings schema path.
func removeExcessSlashChars(path string) string {
	var buf bytes.Buffer
	lastCharIsSlash := false
	for _, ch := range path {
		if ch == '/' {
			if lastCharIsSlash {
				continue
			}
			lastCharIsSlash = true
		} else {
			lastCharIsSlash = false
		}
		buf.WriteRune(ch)
	}

	path = buf.String()

	return path
}

// NewSettingsStore create new SettingsStore object - wrapper on glib.Settings.
func NewSettingsStore(schemaID string, path string, changed func()) (*SettingsStore, error) {
	path = removeExcessSlashChars(path)
	gs, err := glib.SettingsNewWithPath(schemaID, path)
	if err != nil {
		return nil, err
	}
	_, err = gs.Connect("changed", func() {
		if changed != nil {
			changed()
		}
	})
	if err != nil {
		return nil, err
	}
	v := &SettingsStore{settings: gs, schemaID: schemaID, path: path}
	return v, nil
}

// GetChildSettingsStore generate child glib.Settings object to manipulate with nested scheme.
func (v *SettingsStore) GetChildSettingsStore(suffixSchemaID string, suffixPath string,
	changed func()) (*SettingsStore, error) {

	newSchemaID := v.schemaID + "." + suffixSchemaID
	newPath := v.path + "/" + suffixPath + "/"
	settings, err := NewSettingsStore(newSchemaID, newPath, changed)
	return settings, err
}

// GetSchema obtains glib.SettingsSchema from glib.Settings.
func (v *SettingsStore) GetSchema() (*glib.SettingsSchema, error) {
	val, err := v.settings.GetProperty("settings-schema")
	if err != nil {
		return nil, err
	}
	if schema, ok := val.(*glib.SettingsSchema); ok {
		return schema, nil
	} else {
		return nil, errors.New("GLib settings-schema property is not convertible to SettingsSchema")
	}
}

// SettingsArray is a way how to create multiple (indexed) GLib setting's group
// based on single schema. For instance, multiple backup profiles with identical
// settings inside of each profile.
type SettingsArray struct {
	store   *SettingsStore
	arrayID string
}

// NewSettingsArray creates new SettingsArray, to keep/add/delete new
// indexed glib.Settings object based on single schema.
func (v *SettingsStore) NewSettingsArray(arrayID string) *SettingsArray {
	sa := &SettingsArray{store: v, arrayID: arrayID}
	return sa
}

// DeleteNode delete specific indexed glib.Settings defined by nodeID.
func (v *SettingsArray) DeleteNode(childStore *SettingsStore, nodeID string) error {
	// Delete/reset whole child settings object.
	schema, err := childStore.GetSchema()
	if err != nil {
		return err
	}
	keys := schema.ListKeys()
	for _, key := range keys {
		childStore.settings.Reset(key)
	}

	// Delete index from the array, which identify
	// child object settings.
	original := v.store.settings.GetStrv(v.arrayID)
	var updated []string
	for _, id := range original {
		if id != nodeID {
			updated = append(updated, id)
		}
	}
	v.store.settings.SetStrv(v.arrayID, updated)
	return nil
}

// AddNode add specific indexed glib.Settings identified by returned nodeID.
func (v *SettingsArray) AddNode() (nodeID string, err error) {
	list := v.store.settings.GetStrv(v.arrayID)
	// Append index to the end of array, which reference to the list
	// of child settings based on single settings schema.
	var ni int
	if len(list) > 0 {
		ni, err = strconv.Atoi(list[len(list)-1])
		if err != nil {
			return "", err
		}
		ni++
	}
	list = append(list, strconv.Itoa(ni))
	v.store.settings.SetStrv(v.arrayID, list)
	return list[len(list)-1], nil
}

// GetArrayIDs return identifiers of glib.Settings with common schema,
// which can be accessed using id from the list.
func (v *SettingsArray) GetArrayIDs() []string {
	list := v.store.settings.GetStrv(v.arrayID)
	return list
}

// Binding cache link between Key string identifier and GLIB object property.
// Code partially taken from https://github.com/gnunn1/tilix project.
type Binding struct {
	Key      string
	Object   glib.IObject
	Property string
	Flags    glib.SettingsBindFlags
}

// BindingHelper is a bookkeeping class that keeps track of objects which are
// binded to a GSettings object so they can be unbinded later. it
// also supports the concept of deferred bindings where a binding
// can be added but is not actually attached to a Settings object
// until one is set.
type BindingHelper struct {
	bindings []Binding
	settings *SettingsStore
}

// NewBindingHelper creates new BindingHelper object.
func (v *SettingsStore) NewBindingHelper() *BindingHelper {
	bh := &BindingHelper{settings: v}
	return bh
}

// SetSettings will replace underlying GLIB Settings object to unbind
// previously set bindings and re-bind to the new settings automatically.
func (v *BindingHelper) SetSettings(value *SettingsStore) {
	if value != v.settings {
		if v.settings != nil {
			v.Unbind()
		}
		v.settings = value
		if v.settings != nil {
			v.bindAll()
		}
	}
}

func (v *BindingHelper) bindAll() {
	if v.settings != nil {
		for _, b := range v.bindings {
			v.settings.settings.Bind(b.Key, b.Object, b.Property, b.Flags)
		}
	}
}

// addBind add a binding to the list
func (v *BindingHelper) addBind(key string, object glib.IObject, property string, flags glib.SettingsBindFlags) {
	v.bindings = append(v.bindings, Binding{key, object, property, flags})
}

// Bind add a binding to list and binds to Settings if it is set.
func (v *BindingHelper) Bind(key string, object glib.IObject, property string, flags glib.SettingsBindFlags) {
	v.addBind(key, object, property, flags)
	if v.settings != nil {
		v.settings.settings.Bind(key, object, property, flags)
	}
}

// Unbind all added binds from settings object.
func (v *BindingHelper) Unbind() {
	for _, b := range v.bindings {
		v.settings.settings.Unbind(b.Object, b.Property)
	}
}

// Clear unbind all bindings and clears list of bindings.
func (v *BindingHelper) Clear() {
	v.Unbind()
	v.bindings = nil
}

// ==========================================================================================
// ************************* GLIB SETTINGS UTILITIES SECTION END ****************************
// ==========================================================================================

// String constants used as titles/identifiers
var (
	APP_TITLE = "Cool App (GTK+ 3 UI adaptation for golang based on imporved GOTK3)"

	//Preference Constants
	APP_SCHEMA_ID              = "org.d2r2.gotk3.cool_app_1"
	SETTINGS_SCHEMA_ID         = APP_SCHEMA_ID + "." + "Settings"
	SETTINGS_SCHEMA_PATH       = "/org/gtk/gotk3/cool_app_1/"
	PROFILE_SCHEMA_SUFFIX_ID   = "Profile"
	PROFILE_SCHEMA_SUFFIX_PATH = "profiles/%s"

	CFG_AUTO_HIDE_MOUSE_KEY                   = "auto-hide-mouse"
	CFG_DONT_SHOW_ABOUT_ON_STARTUP_KEY        = "dont-show-about-dialog-on-startup"
	CFG_PROMPT_ON_NEW_SESSION_KEY             = "prompt-on-new-session"
	CFG_ENABLE_TRANSPARENCY_KEY               = "enable-transparency"
	CFG_CLOSE_WITH_LAST_SESSION_KEY           = "close-with-last-session"
	CFG_APP_TITLE_KEY                         = "app-title"
	CFG_CONTROL_CLICK_TITLE_KEY               = "control-click-titlebar"
	CFG_INHERIT_WINDOW_STATE_KEY              = "new-window-inherit-state"
	CFG_USE_OVERLAY_SCROLLBAR_KEY             = "use-overlay-scrollbar"
	CFG_TERMINAL_FOCUS_FOLLOWS_MOUSE_KEY      = "focus-follow-mouse"
	CFG_MIDDLE_CLICK_CLOSE_KEY                = "middle-click-close"
	CFG_CONTROL_SCROLL_ZOOM_KEY               = "control-scroll-zoom"
	CFG_WINDOW_SAVE_STATE_KEY                 = "window-save-state"
	CFG_NOTIFY_ON_PROCESS_COMPLETE_KEY        = "notify-on-process-complete"
	CFG_PASTE_ADVANCED_DEFAULT_KEY            = "paste-advanced-default"
	CFG_UNSAFE_PASTE_ALERT_KEY                = "unsafe-paste-alert"
	CFG_STRIP_FIRST_COMMENT_CHAR_ON_PASTE_KEY = "paste-strip-first-char"
	CFG_COPY_ON_SELECT_KEY                    = "copy-on-select"
	CFG_WINDOW_STYLE_KEY                      = "window-style"

	CFG_TERMINAL_TITLE_STYLE_KEY          = "terminal-title-style"
	CFG_TERMINAL_TITLE_STYLE_VALUE_NORMAL = "normal"
	CFG_TERMINAL_TITLE_STYLE_VALUE_SMALL  = "small"
	CFG_TERMINAL_TITLE_STYLE_VALUE_NONE   = "none"

	CFG_TAB_POSITION_KEY = "tab-position"

	// Theme Settings
	CFG_THEME_VARIANT_KEY          = "theme-variant"
	CFG_THEME_VARIANT_SYSTEM_VALUE = "system"
	CFG_THEME_VARIANT_LIGHT_VALUE  = "light"
	CFG_THEME_VARIANT_DARK_VALUE   = "dark"

	CFG_BACKGROUND_IMAGE_KEY                = "background-image"
	CFG_BACKGROUND_IMAGE_MODE_KEY           = "background-image-mode"
	CFG_BACKGROUND_IMAGE_MODE_SCALE_VALUE   = "scale"
	CFG_BACKGROUND_IMAGE_MODE_TILE_VALUE    = "tile"
	CFG_BACKGROUND_IMAGE_MODE_CENTER_VALUE  = "center"
	CFG_BACKGROUND_IMAGE_MODE_STRETCH_VALUE = "stretch"

	CFG_SIDEBAR_RIGHT = "sidebar-on-right"

	CFG_TERMINAL_TITLE_SHOW_WHEN_SINGLE_KEY = "terminal-title-show-when-single"

	CFG_USE_TABS_KEY = "use-tabs"

	CFG_BACKUP_LIST         = "profile-list"
	CFG_PROFILE_NAME        = "profile-name"
	CFG_PROFILE_DESCRIPTION = "profile-description"
)

type AppRunMode int

const (
	AppRegularRun AppRunMode = iota
	AppRunReload
)

var appRunMode AppRunMode

// Keeps here global fullscreen data, which available throughout the application.
// TODO: add thread-safe protection.
type FullscreenGlobalData struct {
	EventBox             *gtk.EventBox
	Revealer             *gtk.Revealer
	FullscreeAction      *glib.Action
	RightTopMenuButton   *gtk.MenuButton
	InFullscreenEventBox bool
	Timer                *time.Timer
}

// Demonstration of Popover menu creation with different kind of menus:
//	1) Regular menus (with text)
//	2) Button menus with themed icons (native GTK+ icons) and loaded file-base icons
//	3) Submenus
func createMenuModelForPopover() (glib.IMenuModel, error) {

	main, err := glib.MenuNew()
	if err != nil {
		return nil, err
	}

	var section *glib.Menu
	var item *glib.MenuItem

	// New menu section (with buttons)
	section, err = glib.MenuNew()
	if err != nil {
		return nil, err
	}
	item, err = SetupMenuItemWithThemedIcon("About", "win.AboutAction", "help-about-symbolic")
	if err != nil {
		return nil, err
	}
	section.AppendItem(item)
	item, err = SetupMenuItemWithThemedIcon("Preference", "win.PreferenceAction", "preferences-other-symbolic")
	if err != nil {
		return nil, err
	}
	section.AppendItem(item)
	item, err = SetupMenuItemWithThemedIcon("Fullscreen", "win.FullscreenAction", "view-fullscreen-symbolic")
	if err != nil {
		return nil, err
	}
	section.AppendItem(item)
	err = AppendSectionAsHorzButtons(main, section)
	if err != nil {
		return nil, err
	}

	// New menu section (automatically separeted by dividers)
	section, err = glib.MenuNew()
	if err != nil {
		return nil, err
	}
	section.Append("Enable/Disable option", "win.CheckBoxAction")
	main.AppendSection("", section)

	// New menu section (automatically separeted by dividers)
	section, err = glib.MenuNew()
	if err != nil {
		return nil, err
	}
	section.Append("Red", "win.ChooseColor('red')")
	section.Append("Green", "win.ChooseColor('green')")
	section.Append("Blue", "win.ChooseColor('blue')")
	main.AppendSection("", section)

	// New menu section (automatically separeted by dividers)
	section, err = glib.MenuNew()
	if err != nil {
		return nil, err
	}
	section.Append("Dialog demo 1", "win.DialogAction1")
	section.Append("Dialog demo 2", "win.DialogAction2")
	section.Append("Dialog demo 3", "win.DialogAction3")
	section.Append("Dialog demo 4", "win.DialogAction4")
	section.Append("Dialog demo 5", "win.DialogAction5")
	main.AppendSection("", section)

	// New menu section with submenu (automatically separeted by dividers)
	section, err = glib.MenuNew()
	if err != nil {
		return nil, err
	}
	subMenu, err := glib.MenuNew()
	if err != nil {
		return nil, err
	}
	subMenu.Append("Fullscreen mode", "win.FullscreenAction")
	section.AppendSubmenu("Submenu", subMenu)
	main.AppendSection("", section)

	// New menu section with submenu (with buttons)
	section, err = glib.MenuNew()
	if err != nil {
		return nil, err
	}
	item, err = SetupMenuItemWithThemedIcon("Quit application", "win.QuitAction", "application-exit-symbolic")
	if err != nil {
		return nil, err
	}
	section.AppendItem(item)
	err = AppendSectionAsHorzButtons(main, section)
	if err != nil {
		return nil, err
	}

	return main, nil
}

// Creates GLIB's MenuModel interface for using in GTK MenuBar.
func createMenuModelForMenuBar() (glib.IMenuModel, error) {
	rootMenu, err := glib.MenuNew()
	if err != nil {
		return nil, err
	}
	main, err := glib.MenuNew()
	if err != nil {
		return nil, err
	}
	rootMenu.AppendSubmenu("Main menu", main)

	var section *glib.Menu
	var item *glib.MenuItem

	// New menu section (automatically separeted by dividers)
	section, err = glib.MenuNew()
	if err != nil {
		return nil, err
	}
	main.AppendSection("", section)
	section.Append("Fullscreen mode", "win.FullscreenAction")

	// New menu section (automatically separeted by dividers)
	section, err = glib.MenuNew()
	if err != nil {
		return nil, err
	}
	section.Append("Enable/Disable option", "win.CheckBoxAction")
	main.AppendSection("", section)

	// New menu section (automatically separeted by dividers)
	section, err = glib.MenuNew()
	if err != nil {
		return nil, err
	}
	main.AppendSection("", section)
	section.Append("Red", "win.ChooseColor('red')")
	section.Append("Green", "win.ChooseColor('green')")
	section.Append("Blue", "win.ChooseColor('blue')")

	// New menu section with submeny (automatically separeted by dividers)
	section, err = glib.MenuNew()
	if err != nil {
		return nil, err
	}
	main.AppendSection("", section)
	subMenu, err := glib.MenuNew()
	if err != nil {
		return nil, err
	}
	subMenu.Append("Fullscreen", "win.FullscreenAction")
	section.AppendSubmenu("Submenu", subMenu)

	// New menu section (automatically separeted by dividers)
	section, err = glib.MenuNew()
	if err != nil {
		return nil, err
	}
	section.Append("Dialog demo 1", "win.DialogAction1")
	section.Append("Dialog demo 2", "win.DialogAction2")
	section.Append("Dialog demo 3", "win.DialogAction3")
	section.Append("Dialog demo 4", "win.DialogAction4")
	section.Append("Dialog demo 5", "win.DialogAction5")
	main.AppendSection("", section)

	// New menu section with submeny (automatically separeted by dividers)
	section, err = glib.MenuNew()
	if err != nil {
		return nil, err
	}
	file, err := glib.FileForPathNew("./icons/ajax-loader-gears_32x32.gif")
	//file, err := glib.FileForPathNew("./icons/com.gexperts.Tilix.png")
	if err != nil {
		return nil, err
	}
	icon, err := glib.FileIconNew(file)
	if err != nil {
		return nil, err
	}
	item, err = SetupMenuItemWithIcon("Preference dialog demo", "win.PreferenceAction", &icon.Icon)
	if err != nil {
		return nil, err
	}
	section.AppendItem(item)
	icon2, err := glib.ThemedIconNew("help-about-symbolic")
	if err != nil {
		return nil, err
	}
	item, err = SetupMenuItemWithIcon("About demo", "win.AboutAction", &icon2.Icon)
	if err != nil {
		return nil, err
	}
	section.AppendItem(item)
	main.AppendSection("", section)

	// New menu section with submeny (automatically separeted by dividers)
	section, err = glib.MenuNew()
	if err != nil {
		return nil, err
	}
	icon3, err := glib.ThemedIconNew("application-exit-symbolic")
	if err != nil {
		return nil, err
	}
	item, err = SetupMenuItemWithIcon("Quit application", "win.QuitAction", &icon3.Icon)
	if err != nil {
		return nil, err
	}
	section.AppendItem(item)
	main.AppendSection("", section)

	return rootMenu, nil
}

// Creates GTK Toolbar actions attached to buttons.
func createToolbar() (*gtk.Toolbar, error) {
	tbx, err := gtk.ToolbarNew()
	if err != nil {
		return nil, err
	}
	tbx.SetStyle(gtk.TOOLBAR_BOTH_HORIZ)

	var tbtn *gtk.ToolButton
	var tdvd *gtk.SeparatorToolItem
	var img *gtk.Image

	img, err = gtk.ImageNew()
	if err != nil {
		return nil, err
	}
	img.SetFromIconName("application-exit-symbolic", gtk.ICON_SIZE_BUTTON)
	tbtn, err = gtk.ToolButtonNew(img, "")
	if err != nil {
		return nil, err
	}
	tbtn.SetActionName("win.QuitAction")
	tbx.Add(tbtn)

	img, err = gtk.ImageNew()
	if err != nil {
		return nil, err
	}
	img.SetFromIconName("view-fullscreen-symbolic", gtk.ICON_SIZE_BUTTON)
	tbtn, err = gtk.ToolButtonNew(img, "")
	if err != nil {
		return nil, err
	}
	tbtn.SetActionName("win.FullscreenAction")
	tbx.Add(tbtn)

	tdvd, err = gtk.SeparatorToolItemNew()
	if err != nil {
		return nil, err
	}
	tbx.Add(tdvd)

	tbtn, err = gtk.ToolButtonNew(nil, "Demo 1")
	if err != nil {
		return nil, err
	}
	tbtn.SetActionName("win.DialogAction1")
	tbx.Add(tbtn)

	tbtn, err = gtk.ToolButtonNew(nil, "Demo 2")
	if err != nil {
		return nil, err
	}
	tbtn.SetActionName("win.DialogAction2")
	tbx.Add(tbtn)

	tbtn, err = gtk.ToolButtonNew(nil, "Demo 3")
	if err != nil {
		return nil, err
	}
	tbtn.SetActionName("win.DialogAction3")
	tbx.Add(tbtn)

	tbtn, err = gtk.ToolButtonNew(nil, "Demo 4")
	if err != nil {
		return nil, err
	}
	tbtn.SetActionName("win.DialogAction4")
	tbx.Add(tbtn)

	tbtn, err = gtk.ToolButtonNew(nil, "Demo 5")
	if err != nil {
		return nil, err
	}
	tbtn.SetActionName("win.DialogAction5")
	tbx.Add(tbtn)

	tdvd, err = gtk.SeparatorToolItemNew()
	if err != nil {
		return nil, err
	}
	tbx.Add(tdvd)

	img, err = gtk.ImageNewFromFile("./icons/ajax-loader-gears_32x32.gif")
	if err != nil {
		return nil, err
	}
	tbtn, err = gtk.ToolButtonNew(img, "")
	if err != nil {
		return nil, err
	}
	tbtn.SetActionName("win.PreferenceAction")
	tbx.Add(tbtn)

	img, err = gtk.ImageNew()
	if err != nil {
		return nil, err
	}
	img.SetFromIconName("help-about-symbolic", gtk.ICON_SIZE_BUTTON)
	tbtn, err = gtk.ToolButtonNew(img, "")
	if err != nil {
		return nil, err
	}
	tbtn.SetActionName("win.AboutAction")
	tbx.Add(tbtn)

	return tbx, nil
}

// Create demonstration action-with-state with boolean logic (on/of).
// Action trigger is included.
func createCheckBoxAction() (glib.IAction, error) {
	v, err := glib.VariantBooleanNew(true)
	if err != nil {
		return nil, err
	}
	act, err := glib.SimpleActionStatefullNew("CheckBoxAction", nil, v)
	if err != nil {
		return nil, err
	}
	if act == nil {
		return nil, errors.New("error")
	}

	_, err = act.Connect("activate", func(action *glib.SimpleAction, param *glib.Variant) {
		name, state, err := GetActionNameAndState(action)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(spew.Sprintf("%v action activated with current state %v and args %v",
			name, state, param))

		if state != nil && state.IsOfType(glib.VARIANT_TYPE_BOOLEAN) {
			state, err = glib.VariantBooleanNew(!state.GetBoolean())
			if err != nil {
				log.Fatal(err)
			}
			action.ChangeState(state)
		}
	})
	if err != nil {
		return nil, err
	}

	return act, nil
}

// Create demonstration action-with-state with list logic (opt1/opt2/.../optN).
// Action trigger is included.
func createChooseAction() (glib.IAction, error) {
	v, err := glib.VariantStringNew("green")
	if err != nil {
		return nil, err
	}
	act, err := glib.SimpleActionStatefullNew("ChooseColor", glib.VARIANT_TYPE_STRING, v)
	if err != nil {
		return nil, err
	}
	if act == nil {
		return nil, errors.New("error")
	}

	_, err = act.Connect("activate", func(action *glib.SimpleAction, param *glib.Variant) {
		name, state, err := GetActionNameAndState(action)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(spew.Sprintf("%v action activated with current state %v and args %v",
			name, state, param))

		act.ChangeState(param)
	})
	if err != nil {
		return nil, err
	}

	return act, nil
}

// Create fullscreen action with boolean in the base (fullscreen on/off)
// Action trigger is included.
func createFullscreenAction(win *gtk.Window, data *FullscreenGlobalData) (glib.IAction, error) {
	v, err := glib.VariantBooleanNew(false)
	if err != nil {
		return nil, err
	}
	act, err := glib.SimpleActionStatefullNew("FullscreenAction", nil, v)
	if err != nil {
		return nil, err
	}
	data.FullscreeAction = &act.Action

	_, err = act.Connect("activate", func(action *glib.SimpleAction, param *glib.Variant) {
		name, state, err := GetActionNameAndState(action)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(spew.Sprintf("%v action activated with current state %v and args %v",
			name, state, param))

		if state != nil && state.IsOfType(glib.VARIANT_TYPE_BOOLEAN) {
			state, err = glib.VariantBooleanNew(!state.GetBoolean())
			if err != nil {
				log.Fatal(err)
			}
			action.ChangeState(state)

			if state.GetBoolean() {
				win.Fullscreen()
				data.EventBox.ShowAll()
				//data.Revealer.SetRevealChild(false)
				showFullscreenHeader(data, true)
				manageFullscreenHeaderHideTimer(data, time.Millisecond*2500, true)
			} else {
				data.EventBox.Hide()
				//data.Revealer.SetRevealChild(false)
				win.Unfullscreen()
			}
		}
	})
	if err != nil {
		return nil, err
	}

	return act, nil
}

// Create regular exit app action.
// Action trigger is included.
func createQuitAction(win *gtk.Window) (glib.IAction, error) {
	act, err := glib.SimpleActionNew("QuitAction", nil)
	if err != nil {
		return nil, err
	}

	_, err = act.Connect("activate", func(action *glib.SimpleAction, param *glib.Variant) {
		name, state, err := GetActionNameAndState(action)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(spew.Sprintf("%v action activated with current state %v and args %v",
			name, state, param))

		app, err := win.GetApplication()
		if err != nil {
			log.Fatal(err)
		}
		// loop through and close all windows currently opened
		for {
			win2 := app.GetActiveWindow()
			if win2 != nil {
				win2.Close()
				for gtk.EventsPending() {
					gtk.MainIterationDo(true)
				}
			} else {
				break
			}
		}
		// quit application
		app.Quit()
	})
	if err != nil {
		return nil, err
	}

	return act, nil
}

// Create regular "about dialog" action.
// Action trigger is included.
func createAboutAction(win *gtk.Window, appSettings *SettingsStore) (glib.IAction, error) {
	act, err := glib.SimpleActionNew("AboutAction", nil)
	if err != nil {
		return nil, err
	}

	_, err = act.Connect("activate", func(action *glib.SimpleAction, param *glib.Variant) {
		name, state, err := GetActionNameAndState(action)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(spew.Sprintf("%v action activated with current state %v and args %v",
			name, state, param))

		dlg, err := gtk.AboutDialogNew()
		if err != nil {
			log.Fatal(err)
		}

		dlg.SetAuthors([]string{"Written by Denis Dyakov <denis.dyakov@gmail.com>"})
		dlg.SetProgramName("Cool App")
		dlg.SetLogoIconName("face-cool-symbolic")
		dlg.SetVersion("v0.3")

		bh := appSettings.NewBindingHelper()
		// Show about dialog on application startup
		cbAboutInfo, err := gtk.CheckButtonNewWithLabel("Do not show about information on app startup")
		if err != nil {
			log.Fatal(err)
		}
		bh.Bind(CFG_DONT_SHOW_ABOUT_ON_STARTUP_KEY, cbAboutInfo, "active", glib.SETTINGS_BIND_DEFAULT)

		content, err := dlg.GetContentArea()
		if err != nil {
			log.Fatal(err)
		}
		content.Add(cbAboutInfo)
		content.ShowAll()

		var buf bytes.Buffer
		glibMajor, glibMinor, glibMicro := GetGlibVersion()
		glibBuildVersion := glib.GetBuildVersion()
		gtkMajor, gtkMinor, gtkMicro := GetGtkVersion()
		gtkBuildVersion := gtk.GetBuildVersion()
		buf.WriteString(fmt.Sprintln("This application built for education purpose and compose"))
		buf.WriteString(fmt.Sprintln("code patterns to write GTK+3 user interface in Go language."))
		buf.WriteString(fmt.Sprintln())
		buf.WriteString(fmt.Sprintln("Environment:"))
		buf.WriteString(fmt.Sprintln(fmt.Sprintf("GLIB compiled version %s, detected version %d.%d.%d.",
			glibBuildVersion, glibMajor, glibMinor, glibMicro)))
		buf.WriteString(fmt.Sprintln(fmt.Sprintf("GTK+ compiled version %s, detected version %d.%d.%d.",
			gtkBuildVersion, gtkMajor, gtkMinor, gtkMicro)))

		display, err := gdk.DisplayGetDefault()
		if err != nil {
			log.Fatal(err)
		}
		if gdk.IsWaylandDisplay(display) {
			buf.WriteString("WAYLAND display detected.")
			buf.WriteString(fmt.Sprintln())
		} else if gdk.IsX11Display(display) {
			buf.WriteString("X11 display detected.")
			buf.WriteString(fmt.Sprintln())
		}

		buf.WriteString(fmt.Sprintln(fmt.Sprintf("Application compiled with %s %s.",
			runtime.Version(), runtime.GOARCH)))
		buf.WriteString(fmt.Sprintln())
		buf.WriteString(fmt.Sprintln("Features:"))
		buf.WriteString(fmt.Sprintln("- Actions as code entry points with states and stateless."))
		buf.WriteString(fmt.Sprintln("- Fullscreen mode code pattern out-of-the-box."))
		buf.WriteString(fmt.Sprintln("- Preference dialog demo with save/restore functionality out-of-the-box."))
		buf.WriteString(fmt.Sprintln("  Multi-profile settings option implemented."))
		buf.WriteString(fmt.Sprintln("- Modern popover menu functionality (right upper corner button)."))
		buf.WriteString(fmt.Sprintln("- Various dialog's windows demonstations."))
		buf.WriteString(fmt.Sprintln())
		buf.WriteString(fmt.Sprint("Follow my golang projects on GitHub:"))
		dlg.SetComments(buf.String())

		dlg.SetWebsite("https://github.com/romychs/")

		dlg.SetTransientFor(win)
		dlg.SetModal(true)
		dlg.ShowNow()
	})
	if err != nil {
		return nil, err
	}

	return act, nil
}

// Create demonstration GTK MessageDialog action.
// Action trigger is included.
func createDialogAction1(win *gtk.Window) (glib.IAction, error) {
	act, err := glib.SimpleActionNew("DialogAction1", nil)
	if err != nil {
		return nil, err
	}

	_, err = act.Connect("activate", func(action *glib.SimpleAction, param *glib.Variant) {
		name, state, err := GetActionNameAndState(action)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(spew.Sprintf("%v action activated with current state %v and args %v",
			name, state, param))

		title := "<span weight='bold' size='larger'>Configuration Issue Detected (demonstration)</span>"
		paragraphs := []*DialogParagraph{NewDialogParagraph("Message dialog based on GTK+ MessageDialog.").
			SetJustify(gtk.JUSTIFY_LEFT).SetHorizAlign(gtk.ALIGN_START)}
		paragraphs = append(paragraphs, NewDialogParagraph("This message based on GtkMessageDialog functionality.\n"+
			"GtkMessageDialog doesn't support setting of image\n"+
			"via SetImage as message type specification since 3.22.\n"+
			"So, that's why there is not warning icon either something similar here.").
			SetJustify(gtk.JUSTIFY_LEFT).SetHorizAlign(gtk.ALIGN_START))
		paragraphs = append(paragraphs, NewDialogParagraph("There appears to be an issue with the configuration of the application.\n"+
			"This issue is not serious, but correcting it will improve your experience.").
			SetJustify(gtk.JUSTIFY_LEFT).SetHorizAlign(gtk.ALIGN_START))

		buttons := []DialogButton{
			{"_OK", gtk.RESPONSE_OK, false, nil},
		}

		dialog, err := SetupMessageDialog(win, title, "", paragraphs, buttons,
			func(area *gtk.Box) error {
				lbl, err := gtk.LabelNew("Click the link below for more information, if still not clear:")
				if err != nil {
					return err
				}
				lbl.SetHAlign(gtk.ALIGN_START)
				area.Add(lbl)

				link, err := gtk.LinkButtonNew("https://www.youtube.com/watch?v=dQw4w9WgXcQ")
				if err != nil {
					return err
				}
				area.Add(link)

				cb, err := gtk.CheckButtonNewWithLabel("Do not show this message again")
				if err != nil {
					return err
				}
				area.Add(cb)

				return nil
			})
		if err != nil {
			log.Fatal(err)
		}
		dialog.Run(false)
	})
	if err != nil {
		return nil, err
	}

	return act, nil
}

// Create demonstration GTK Dialog action.
// Action trigger is included.
func createDialogAction2(win *gtk.Window) (glib.IAction, error) {
	act, err := glib.SimpleActionNew("DialogAction2", nil)
	if err != nil {
		return nil, err
	}

	_, err = act.Connect("activate", func(action *glib.SimpleAction, param *glib.Variant) {
		name, state, err := GetActionNameAndState(action)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(spew.Sprintf("%v action activated with current state %v and args %v",
			name, state, param))

		title := "Configuration Issue Detected (demonstration)"
		paragraphs := []*DialogParagraph{NewDialogParagraph("Message dialog based on GTK+ Dialog.").
			SetJustify(gtk.JUSTIFY_CENTER).SetHorizAlign(gtk.ALIGN_CENTER)}
		paragraphs = append(paragraphs, NewDialogParagraph("This message based on GtkDialog functionality.").
			SetJustify(gtk.JUSTIFY_CENTER).SetHorizAlign(gtk.ALIGN_CENTER))
		paragraphs = append(paragraphs, NewDialogParagraph("There appears to be an issue with the configuration of the application.\n"+
			"This issue is not serious, but correcting it will improve your experience.").
			SetJustify(gtk.JUSTIFY_CENTER).SetHorizAlign(gtk.ALIGN_CENTER))
		buttons := []DialogButton{
			{"_OK", gtk.RESPONSE_OK, false, nil},
		}

		_, err = RunDialog(win, gtk.MESSAGE_WARNING, true, title, paragraphs, false, buttons,
			func(area *gtk.Box) error {
				lbl, err := gtk.LabelNew("Click the link below for more information, if still not clear:")
				if err != nil {
					return err
				}
				lbl.SetHAlign(gtk.ALIGN_START)
				area.Add(lbl)

				link, err := gtk.LinkButtonNew("https://www.youtube.com/watch?v=dQw4w9WgXcQ")
				if err != nil {
					return err
				}
				area.Add(link)

				cb, err := gtk.CheckButtonNewWithLabel("Do not show this message again")
				if err != nil {
					return err
				}
				area.Add(cb)

				return nil
			})
		if err != nil {
			log.Fatal(err)
		}
	})
	if err != nil {
		return nil, err
	}

	return act, nil
}

// Create demonstration GTK MessageDialog action.
// Action trigger is included.
func createDialogAction3(win *gtk.Window) (glib.IAction, error) {
	act, err := glib.SimpleActionNew("DialogAction3", nil)
	if err != nil {
		return nil, err
	}

	_, err = act.Connect("activate", func(action *glib.SimpleAction, param *glib.Variant) {
		name, state, err := GetActionNameAndState(action)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(spew.Sprintf("%v action activated with current state %v and args %v",
			name, state, param))

		title := "<span weight='bold' size='larger'>Configuration Issue Detected (demonstration)</span>"
		paragraphs := []*DialogParagraph{NewDialogParagraph("Message dialog based on GTK+ MessageDialog.").
			SetJustify(gtk.JUSTIFY_CENTER).SetHorizAlign(gtk.ALIGN_CENTER)}
		paragraphs = append(paragraphs, NewDialogParagraph("This message based on GtkMessageDialog functionality.\n"+
			"GtkMessageDialog doesn't support setting of image\n"+
			"via SetImage as message type specification since 3.22.\n"+
			"So, that's why there is not warning icon either something similar here.").
			SetJustify(gtk.JUSTIFY_CENTER).SetHorizAlign(gtk.ALIGN_CENTER))
		paragraphs = append(paragraphs, NewDialogParagraph("There appears to be an issue with the configuration of the application.\n"+
			"This issue is not serious, but correcting it will improve your experience.").
			SetJustify(gtk.JUSTIFY_CENTER).SetHorizAlign(gtk.ALIGN_CENTER))

		buttons := []DialogButton{
			{"_Yes", gtk.RESPONSE_YES, false, func(btn *gtk.Button) error {
				style, err := btn.GetStyleContext()
				if err != nil {
					return err
				}
				style.AddClass("suggested-action")
				//style.AddClass("destructive-action")
				return nil
			}},
			{"_No", gtk.RESPONSE_NO, true, nil},
		}

		dialog, err := SetupMessageDialog(win, title, "", paragraphs, buttons,
			func(area *gtk.Box) error {
				lbl, err := gtk.LabelNew("Click the link below for more information, if still not clear:")
				if err != nil {
					return err
				}
				lbl.SetHAlign(gtk.ALIGN_START)
				area.Add(lbl)

				link, err := gtk.LinkButtonNew("https://www.youtube.com/watch?v=dQw4w9WgXcQ")
				if err != nil {
					return err
				}
				area.Add(link)

				return nil
			})
		if err != nil {
			log.Fatal(err)
		}
		dialog.Run(true)
	})
	if err != nil {
		return nil, err
	}

	return act, nil
}

// Create demonstration GTK Dialog action.
// Action trigger is included.
func createDialogAction4(win *gtk.Window) (glib.IAction, error) {
	act, err := glib.SimpleActionNew("DialogAction4", nil)
	if err != nil {
		return nil, err
	}

	_, err = act.Connect("activate", func(action *glib.SimpleAction, param *glib.Variant) {
		name, state, err := GetActionNameAndState(action)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(spew.Sprintf("%v action activated with current state %v and args %v",
			name, state, param))

		title := "Choose option (demonstration)"
		paragraphs := []*DialogParagraph{NewDialogParagraph("Press Yes to start processing.").
			SetJustify(gtk.JUSTIFY_CENTER).SetHorizAlign(gtk.ALIGN_CENTER)}
		paragraphs = append(paragraphs, NewDialogParagraph("Note: processing might takes significant amount of time.").
			SetJustify(gtk.JUSTIFY_CENTER).SetHorizAlign(gtk.ALIGN_CENTER))

		buttons := []DialogButton{
			{"_Yes", gtk.RESPONSE_YES, true, func(btn *gtk.Button) error {
				/*
					style, err := btn.GetStyleContext()
					if err != nil {
						return err
					}
					//style.AddClass("suggested-action")
					style.RemoveClass("suggested-action")

					style.AddClass("destructive-action")
				*/
				return nil
			}},
			{"_No", gtk.RESPONSE_NO, false, nil},
		}

		_, err = RunDialog(win, gtk.MESSAGE_QUESTION, true, title, paragraphs, true, buttons,
			nil)
		if err != nil {
			log.Fatal(err)
		}
	})
	if err != nil {
		return nil, err
	}

	return act, nil
}

// Create demonstration GTK Dialog action.
// Action trigger is included.
func createDialogAction5(win *gtk.Window) (glib.IAction, error) {
	act, err := glib.SimpleActionNew("DialogAction5", nil)
	if err != nil {
		return nil, err
	}

	_, err = act.Connect("activate", func(action *glib.SimpleAction, param *glib.Variant) {
		name, state, err := GetActionNameAndState(action)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(spew.Sprintf("%v action activated with current state %v and args %v",
			name, state, param))

		title := "Choose option (demonstration)"
		paragraphs := []*DialogParagraph{NewDialogParagraph("Press Yes to start processing.").
			SetJustify(gtk.JUSTIFY_CENTER).SetHorizAlign(gtk.ALIGN_CENTER)}
		paragraphs = append(paragraphs, NewDialogParagraph("Note: processing might takes significant amount of time.").
			SetJustify(gtk.JUSTIFY_CENTER).SetHorizAlign(gtk.ALIGN_CENTER))

		buttons := []DialogButton{
			{"_Yes", gtk.RESPONSE_YES, false, func(btn *gtk.Button) error {
				style, err := btn.GetStyleContext()
				if err != nil {
					return err
				}
				//style.AddClass("suggested-action")
				style.AddClass("destructive-action")
				return nil
			}},
			{"_No", gtk.RESPONSE_NO, true, nil},
		}

		_, err = RunDialog(win, gtk.MESSAGE_QUESTION, false, title, paragraphs, true, buttons,
			nil)
		if err != nil {
			log.Fatal(err)
		}
	})
	if err != nil {
		return nil, err
	}

	return act, nil
}

// Create sophisticated multi-page preference dialog
// with save/restore functionality to/from the GLib Setting object.
// Action activation require to have GLib Setting Schema
// preliminary installed, otherwise will not work raising message.
// Installation bash script from app folder must be performed in advance.
func createPreferenceAction(mainWin *gtk.ApplicationWindow) (glib.IAction, error) {
	act, err := glib.SimpleActionNew("PreferenceAction", nil)
	if err != nil {
		return nil, err
	}

	_, err = act.Connect("activate", func(action *glib.SimpleAction, param *glib.Variant) {
		name, state, err := GetActionNameAndState(action)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(spew.Sprintf("%v action activated with current state %v and args %v",
			name, state, param))

		app, err := mainWin.GetApplication()
		if err != nil {
			log.Fatal(err)
		}

		found, err := checkSchemaSettingsIsInstalled(app)
		if err != nil {
			log.Fatal(err)
		}

		if found {

			win, err := createPreferenceDialog(SETTINGS_SCHEMA_ID, SETTINGS_SCHEMA_PATH, mainWin)
			if err != nil {
				log.Fatal(err)
			}

			win.ShowAll()
			win.Show()

			_, err = win.Connect("destroy", func(window *gtk.ApplicationWindow) {
				window.Destroy()
				log.Println("Destroy window")
			})
			if err != nil {
				log.Fatal(err)
			}
		}

	})
	if err != nil {
		return nil, err
	}

	return act, nil
}

// Automatically hide/unhide panel used when fullscreen mode is on.
func createRevealer(data *FullscreenGlobalData) (*gtk.Revealer, error) {
	rev, err := gtk.RevealerNew()
	if err != nil {
		return nil, err
	}
	box, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)
	if err != nil {
		return nil, err
	}
	rev.Add(box)

	menu, err := createMenuModelForPopover()
	if err != nil {
		return nil, err
	}

	btn, err := SetupMenuButtonWithThemedImage("open-menu-symbolic")
	if err != nil {
		return nil, err
	}
	data.RightTopMenuButton = btn

	btn.SetUsePopover(true)
	btn.SetMenuModel(menu)

	btn2, err := gtk.SeparatorToolItemNew()
	if err != nil {
		return nil, err
	}

	btn3, err := SetupToolButton("view-restore-symbolic", "Leave ssdfsdf")
	if err != nil {
		return nil, err
	}
	//	btn2.SetTitle("asdasdas")
	btn3.SetActionName("win.FullscreenAction")

	hdr, err := SetupHeader(APP_TITLE, "(fullscreen mode)", false)
	if err != nil {
		log.Fatal(err)
	}

	hdr.PackEnd(btn3)
	hdr.PackEnd(btn2)
	hdr.PackEnd(btn)

	box.PackStart(hdr, false, false, 0)

	return rev, nil

}

// Define timer timeouts to hide fullscreen header in case of inactivity.
func manageFullscreenHeaderHideTimer(data *FullscreenGlobalData, dur time.Duration, start bool) {
	if start {
		if data.Timer != nil {
			if !data.Timer.Stop() {
			}
			data.Timer.Reset(dur)
		} else {
			data.Timer = time.AfterFunc(dur, func() {
				showFullscreenHeader(data, false)
			})
		}
	} else {
		if data.Timer != nil {
			if !data.Timer.Stop() {
			}
		}
	}
}

// Show/hide fullscreen header.
func showFullscreenHeader(data *FullscreenGlobalData, show bool) {
	if show {
		data.InFullscreenEventBox = true

		v := data.FullscreeAction.GetState()
		if v != nil && v.GetBoolean() {
			data.Revealer.SetRevealChild(true)
		}
	} else {
		data.InFullscreenEventBox = false
		v := data.FullscreeAction.GetState()
		if v != nil && v.GetBoolean() && !data.RightTopMenuButton.GetActive() {
			data.Revealer.SetRevealChild(false)
		}
	}
}

// Create EventBox widget to catch mouse move events over hidden fullscreen panel.
func createEventBox(data *FullscreenGlobalData) (*gtk.EventBox, error) {
	eb, err := gtk.EventBoxNew()
	if err != nil {
		return nil, err
	}
	dur := time.Millisecond * 1500
	_, err = eb.Connect("enter-notify-event", func() {
		log.Println("EventBox enter notify")

		manageFullscreenHeaderHideTimer(data, dur, false)
		showFullscreenHeader(data, true)
	})
	if err != nil {
		return nil, err
	}
	_, err = eb.Connect("leave-notify-event", func() {
		_, err2 := glib.IdleAdd(func() {
			log.Println("EventBox leave notify")

			manageFullscreenHeaderHideTimer(data, dur, true)
		})
		if err2 != nil {
			log.Fatal(err2)
		}
	})
	if err != nil {
		return nil, err
	}

	return eb, nil
}

// PreferenceRow keeps here extra data for each page of multi-page preference dialog.
type PreferenceRow struct {
	sync.RWMutex
	ID        string
	name      string
	Title     string
	Row       *gtk.ListBoxRow
	Container *gtk.Box
	Label     *gtk.Label
	Icon      *gtk.Image
	Page      *gtk.Container
	Profile   bool
}

// PreferenceRowNew instantiate new PreferenceRow object.
func PreferenceRowNew(id, title string, page *gtk.Container,
	profile bool) (*PreferenceRow, error) {

	box, err := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0)
	if err != nil {
		return nil, err
	}
	SetAllMargins(box, 6)
	box.SetSpacing(6)

	lbl, err := gtk.LabelNew("")
	if err != nil {
		return nil, err
	}
	lbl.SetHAlign(gtk.ALIGN_START)
	box.PackStart(lbl, false, true, 0)

	row, err := gtk.ListBoxRowNew()
	if err != nil {
		return nil, err
	}
	row.Add(box)

	pr := &PreferenceRow{ID: id, Title: title, Row: row,
		Container: box, Label: lbl, Page: page,
		Profile: profile}

	pr.SetName(title)

	return pr, nil
}

// SetName set profile name as a template "Profile(<name>)"
func (v *PreferenceRow) SetName(name string) {
	v.Lock()
	defer v.Unlock()

	v.name = name
	if v.Profile {
		publicName := fmt.Sprintf("Profile (%s)", name)
		v.Label.SetText(publicName)
	} else {
		v.Label.SetText(name)
	}
}

// GetName get name.
func (v *PreferenceRow) GetName() string {
	v.RLock()
	defer v.RUnlock()

	return v.name
}

// PreferenceRowList keeps a link between GtkListBoxRow
// and specific PreferenceRow object.
type PreferenceRowList struct {
	m      map[uintptr]*PreferenceRow
	sorted []uintptr
}

func PreferenceRowListNew() *PreferenceRowList {
	var m = make(map[uintptr]*PreferenceRow)
	v := &PreferenceRowList{m: m}
	return v
}

func (v *PreferenceRowList) Append(row *PreferenceRow) {
	v.m[row.Row.Native()] = row
	v.sorted = append(v.sorted, row.Row.Native())
}

func (v *PreferenceRowList) Delete(rowID uintptr) {
	delete(v.m, rowID)
	for ind, val := range v.sorted {
		if val == rowID {
			v.sorted = append(v.sorted[:ind], v.sorted[ind+1:]...)
			break
		}
	}
}

func (v *PreferenceRowList) Get(rowID uintptr) *PreferenceRow {
	return v.m[rowID]
}

func (v *PreferenceRowList) GetLastProfileListIndex() int {
	lastIndex := -1
	for _, rowID := range v.sorted {
		if v.m[rowID].Profile && v.m[rowID].Row.GetIndex() > lastIndex {
			lastIndex = v.m[rowID].Row.GetIndex()
		}
	}
	return lastIndex
}

func (v *PreferenceRowList) GetProfileCount() int {
	count := 0
	for _, rowID := range v.sorted {
		if v.m[rowID].Profile {
			count++
		}
	}
	return count
}

func (v *PreferenceRowList) GetProfiles() []*PreferenceRow {
	var rows []*PreferenceRow
	for _, rowID := range v.sorted {
		if v.m[rowID].Profile {
			rows = append(rows, v.m[rowID])
		}
	}
	return rows
}

/*
// Keeps here extra data for each page of multi-page preference dialog.
type PreferenceRow struct {
	Name  string
	Title string
}

func PreferenceRowNew(name, title string) (*PreferenceRow, *gtk.ListBoxRow, error) {
	pr := &PreferenceRow{Name: name, Title: title}
	lbl, err := gtk.LabelNew(name)
	if err != nil {
		return nil, nil, err
	}

	lbr, err := gtk.ListBoxRowNew()
	if err != nil {
		return nil, nil, err
	}
	lbl.SetHAlign(gtk.ALIGN_START)
	SetAllMargins(lbl, 6)
	lbr.Add(lbl)

	return pr, lbr, nil
}
*/

// Create preference dialog with "Global" page, where controls
// being bound to GLib Setting object to save/restore functionality.
func GlobalPreferencesNew(appSettings *SettingsStore, actions *glib.ActionMap) (*gtk.Container, error) {
	box, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 6)
	if err != nil {
		return nil, err
	}

	SetAllMargins(box, 18)

	restartBox, err := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 6)
	if err != nil {
		return nil, err
	}
	css := `
		box {
			background-color: shade(@theme_bg_color, 0.8);
			padding: 10px;
		}
	`
	err = ApplyStyleCSS(&restartBox.Widget, css)
	if err != nil {
		return nil, err
	}
	rvl, err := gtk.RevealerNew()
	if err != nil {
		return nil, err
	}
	rvl.Add(restartBox)
	box.Add(rvl)

	img, err := gtk.ImageNew()
	if err != nil {
		return nil, err
	}
	css = `
		image {
			color: @theme_selected_bg_color;
		}
	`
	err = ApplyStyleCSS(&img.Widget, css)
	if err != nil {
		return nil, err
	}
	img.SetFromIconName("emblem-important-symbolic", gtk.ICON_SIZE_BUTTON)
	restartBox.Add(img)
	lblRestart, err := SetupLabelJustifyLeft("")
	if err != nil {
		return nil, err
	}
	lblRestart.SetMarkup("Application reload required. Click to <a href=\"restart_uri\">restart</a> application.")
	_, err = lblRestart.Connect("activate-link", func(v *gtk.Label, href string) {
		if href == "restart_uri" {
			appRunMode = AppRunReload

			actionName := "QuitAction"
			action := actions.LookupAction(actionName)
			if action != nil {
				action.Activate(nil)
			}
		}
	})
	if err != nil {
		return nil, err
	}
	restartBox.Add(lblRestart)
	box.Add(restartBox)

	bh := appSettings.NewBindingHelper()

	lblBehavior, err := gtk.LabelNew(fmt.Sprintf("<b>%s</b>", "Behavior (with app reload demo)"))
	if err != nil {
		return nil, err
	}
	lblBehavior.SetUseMarkup(true)
	lblBehavior.SetHAlign(gtk.ALIGN_START)
	box.Add(lblBehavior)

	fnActivateRestartService := func(v *gtk.CheckButton, initialValue bool) {
		activate := initialValue != v.GetActive()
		// Show "restart app" panel only when language has changed
		// from original setting. Otherwise - hide panel.
		rvl.SetRevealChild(activate)
		if err != nil {
			log.Fatal(err)
		}
	}

	// Show about dialog on application startup
	cbAboutInfo, err := gtk.CheckButtonNewWithLabel("Do not show about information on app startup")
	if err != nil {
		return nil, err
	}
	bh.Bind(CFG_DONT_SHOW_ABOUT_ON_STARTUP_KEY, cbAboutInfo, "active", glib.SETTINGS_BIND_DEFAULT)
	_, err = cbAboutInfo.Connect("clicked", fnActivateRestartService, cbAboutInfo.GetActive())
	if err != nil {
		return nil, err
	}
	box.Add(cbAboutInfo)

	//Prompt on new session
	cbPrompt, err := gtk.CheckButtonNewWithLabel("Prompt when creating a new session")
	if err != nil {
		return nil, err
	}
	bh.Bind(CFG_PROMPT_ON_NEW_SESSION_KEY, cbPrompt, "active", glib.SETTINGS_BIND_DEFAULT)
	_, err = cbPrompt.Connect("clicked", fnActivateRestartService, cbPrompt.GetActive())
	if err != nil {
		return nil, err
	}
	box.Add(cbPrompt)

	//Focus follows the mouse
	cbFocusMouse, err := gtk.CheckButtonNewWithLabel("Focus a terminal when the mouse moves over it")
	if err != nil {
		return nil, err
	}
	bh.Bind(CFG_TERMINAL_FOCUS_FOLLOWS_MOUSE_KEY, cbFocusMouse, "active", glib.SETTINGS_BIND_DEFAULT)
	_, err = cbFocusMouse.Connect("clicked", fnActivateRestartService, cbFocusMouse.GetActive())
	if err != nil {
		return nil, err
	}
	box.Add(cbFocusMouse)

	//Auto hide the mouse
	cbAutoHideMouse, err := gtk.CheckButtonNewWithLabel("Autohide the mouse pointer when typing")
	if err != nil {
		return nil, err
	}
	bh.Bind(CFG_AUTO_HIDE_MOUSE_KEY, cbAutoHideMouse, "active", glib.SETTINGS_BIND_DEFAULT)
	_, err = cbAutoHideMouse.Connect("clicked", fnActivateRestartService, cbAutoHideMouse.GetActive())
	if err != nil {
		return nil, err
	}
	box.Add(cbAutoHideMouse)

	//middle click closes the terminal
	cbMiddleClickClose, err := gtk.CheckButtonNewWithLabel("Close terminal by clicking middle mouse button on title")
	if err != nil {
		return nil, err
	}
	bh.Bind(CFG_MIDDLE_CLICK_CLOSE_KEY, cbMiddleClickClose, "active", glib.SETTINGS_BIND_DEFAULT)
	_, err = cbMiddleClickClose.Connect("clicked", fnActivateRestartService, cbMiddleClickClose.GetActive())
	if err != nil {
		return nil, err
	}
	box.Add(cbMiddleClickClose)

	//zoom in/out terminal with scroll wheel
	cbControlScrollZoom, err := gtk.CheckButtonNewWithLabel("Zoom the terminal using <Control> and scroll wheel")
	if err != nil {
		return nil, err
	}
	bh.Bind(CFG_CONTROL_SCROLL_ZOOM_KEY, cbControlScrollZoom, "active", glib.SETTINGS_BIND_DEFAULT)
	_, err = cbControlScrollZoom.Connect("clicked", fnActivateRestartService, cbControlScrollZoom.GetActive())
	if err != nil {
		return nil, err
	}
	box.Add(cbControlScrollZoom)

	//require control modifier when clicking title
	cbControlClickTitle, err := gtk.CheckButtonNewWithLabel("Require the <Control> modifier to edit title on click")
	if err != nil {
		return nil, err
	}
	bh.Bind(CFG_CONTROL_CLICK_TITLE_KEY, cbControlClickTitle, "active", glib.SETTINGS_BIND_DEFAULT)
	_, err = cbControlClickTitle.Connect("clicked", fnActivateRestartService, cbControlClickTitle.GetActive())
	if err != nil {
		return nil, err
	}
	box.Add(cbControlClickTitle)

	//Closing of last session closes window
	cbCloseWithLastSession, err := gtk.CheckButtonNewWithLabel("Close window when last session is closed")
	if err != nil {
		return nil, err
	}
	bh.Bind(CFG_CLOSE_WITH_LAST_SESSION_KEY, cbCloseWithLastSession, "active", glib.SETTINGS_BIND_DEFAULT)
	_, err = cbCloseWithLastSession.Connect("clicked", fnActivateRestartService, cbCloseWithLastSession.GetActive())
	if err != nil {
		return nil, err
	}
	box.Add(cbCloseWithLastSession)

	cbNewWindowInheritState, err := gtk.CheckButtonNewWithLabel("New window inherits directory and profile from active terminal")
	if err != nil {
		return nil, err
	}
	bh.Bind(CFG_INHERIT_WINDOW_STATE_KEY, cbNewWindowInheritState, "active", glib.SETTINGS_BIND_DEFAULT)
	_, err = cbNewWindowInheritState.Connect("clicked", fnActivateRestartService, cbNewWindowInheritState.GetActive())
	if err != nil {
		return nil, err
	}
	box.Add(cbNewWindowInheritState)

	// Save window state (maximized, minimized, fullscreen) between invocations
	cbWindowSaveState, err := gtk.CheckButtonNewWithLabel("Save and restore window state")
	if err != nil {
		return nil, err
	}
	bh.Bind(CFG_WINDOW_SAVE_STATE_KEY, cbWindowSaveState, "active", glib.SETTINGS_BIND_DEFAULT)
	_, err = cbWindowSaveState.Connect("clicked", fnActivateRestartService, cbWindowSaveState.GetActive())
	if err != nil {
		return nil, err
	}
	box.Add(cbWindowSaveState)

	// *********** Clipboard Options
	lblClipboard, err := gtk.LabelNew(fmt.Sprintf("<b>%s</b>", "Clipboard"))
	if err != nil {
		return nil, err
	}
	lblClipboard.SetUseMarkup(true)
	lblClipboard.SetHAlign(gtk.ALIGN_START)
	box.Add(lblClipboard)

	//Advacned paste is default
	cbAdvDefault, err := gtk.CheckButtonNewWithLabel("Always use advanced paste dialog")
	if err != nil {
		return nil, err
	}
	bh.Bind(CFG_PASTE_ADVANCED_DEFAULT_KEY, cbAdvDefault, "active", glib.SETTINGS_BIND_DEFAULT)
	box.Add(cbAdvDefault)

	//Unsafe Paste Warning
	cbUnsafe, err := gtk.CheckButtonNewWithLabel("Warn when attempting unsafe paste")
	if err != nil {
		return nil, err
	}
	bh.Bind(CFG_UNSAFE_PASTE_ALERT_KEY, cbUnsafe, "active", glib.SETTINGS_BIND_DEFAULT)
	box.Add(cbUnsafe)

	//Strip Paste
	cbStrip, err := gtk.CheckButtonNewWithLabel("Strip first character of paste if comment or variable declaration")
	if err != nil {
		return nil, err
	}
	bh.Bind(CFG_STRIP_FIRST_COMMENT_CHAR_ON_PASTE_KEY, cbStrip, "active", glib.SETTINGS_BIND_DEFAULT)
	box.Add(cbStrip)

	//Copy on Select
	cbCopyOnSelect, err := gtk.CheckButtonNewWithLabel("Automatically copy text to clipboard when selecting")
	if err != nil {
		return nil, err
	}
	bh.Bind(CFG_COPY_ON_SELECT_KEY, cbCopyOnSelect, "active", glib.SETTINGS_BIND_DEFAULT)
	box.Add(cbCopyOnSelect)

	// Disclaimer
	wdg, err := gtk.SeparatorNew(gtk.ORIENTATION_HORIZONTAL)
	if err != nil {
		return nil, err
	}
	box.Add(wdg)
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintln("******************************************************"))
	buf.WriteString(fmt.Sprintln("Design of this preference page taken from " +
		"<a href=\"https://github.com/gnunn1/tilix\">Tilix</a> project."))
	buf.WriteString(fmt.Sprintln("Settings here mainly for demonstration purpose and"))
	buf.WriteString(fmt.Sprintln("have minimal impact to application."))
	buf.WriteString(fmt.Sprintln("******************************************************"))
	lbl, err := gtk.LabelNew("")
	if err != nil {
		return nil, err
	}
	lbl.SetMarkup(buf.String())
	lbl.SetJustify(gtk.JUSTIFY_CENTER)
	box.PackEnd(lbl, true, true, 0)

	_, err = box.Connect("destroy", func(b *gtk.Box) {
		bh.Unbind()
		log.Println("Destroy box")
	})
	if err != nil {
		return nil, err
	}

	return &box.Container, nil
}

// Create preference dialog with "Appearance" page, where controls
// being bound to GLib Setting object to save/restore functionality.
func AppearancePreferencesNew(appSettings *SettingsStore) (*gtk.Container, error) {
	box, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 6)
	if err != nil {
		return nil, err
	}

	SetAllMargins(box, 18)

	bh := appSettings.NewBindingHelper()

	grid, err := gtk.GridNew()
	if err != nil {
		return nil, err
	}
	grid.SetColumnSpacing(12)
	grid.SetRowSpacing(6)
	row := 0

	//Window style
	lbl, err := gtk.LabelNew("Window style")
	if err != nil {
		return nil, err
	}
	lbl.SetHAlign(gtk.ALIGN_END)
	grid.Attach(lbl, 0, row, 1, 1)
	bWindowStyle, err := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 6)
	if err != nil {
		return nil, err
	}
	values := []struct{ value, key string }{
		{"Normal", "normal"},
		{"Disable CSD", "disable-csd"},
		{"Disable CSD, hide toolbar", "disable-csd-hide-toolbar"},
		{"Borderless", "borderless"},
	}
	cbWindowStyle, err := CreateNameValueCombo(values)
	if err != nil {
		return nil, err
	}
	bh.Bind(CFG_WINDOW_STYLE_KEY, cbWindowStyle, "active-id", glib.SETTINGS_BIND_DEFAULT)
	bWindowStyle.Add(cbWindowStyle)

	lblRestart, err := gtk.LabelNew("Window restart required")
	if err != nil {
		return nil, err
	}
	lblRestart.SetHAlign(gtk.ALIGN_START)
	lblRestart.SetSensitive(false)
	bWindowStyle.Add(lblRestart)

	grid.Attach(bWindowStyle, 1, row, 1, 1)
	row++

	//Render terminal titlebars smaller then default
	lbl, err = gtk.LabelNew("Terminal title style")
	if err != nil {
		return nil, err
	}
	lbl.SetHAlign(gtk.ALIGN_END)
	grid.Attach(lbl, 0, row, 1, 1)
	values = []struct{ value, key string }{
		{"Normal", CFG_TERMINAL_TITLE_STYLE_VALUE_NORMAL},
		{"Small", CFG_TERMINAL_TITLE_STYLE_VALUE_SMALL},
		{"None", CFG_TERMINAL_TITLE_STYLE_VALUE_NONE},
	}
	cbTitleStyle, err := CreateNameValueCombo(values)
	if err != nil {
		return nil, err
	}
	bh.Bind(CFG_TERMINAL_TITLE_STYLE_KEY, cbTitleStyle, "active-id", glib.SETTINGS_BIND_DEFAULT)
	grid.Attach(cbTitleStyle, 1, row, 1, 1)
	row++

	lbl, err = gtk.LabelNew("Tab position")
	if err != nil {
		return nil, err
	}
	lbl.SetHAlign(gtk.ALIGN_END)
	grid.Attach(lbl, 0, row, 1, 1)
	values = []struct{ value, key string }{
		{"Left", "left"},
		{"Right", "right"},
		{"Top", "top"},
		{"Bottom", "bottom"},
	}
	cbTabPosition, err := CreateNameValueCombo(values)
	if err != nil {
		return nil, err
	}
	bh.Bind(CFG_TAB_POSITION_KEY, cbTabPosition, "active-id", glib.SETTINGS_BIND_DEFAULT)
	grid.Attach(cbTabPosition, 1, row, 1, 1)
	row++

	//Dark Theme
	lbl, err = gtk.LabelNew("Theme variant")
	if err != nil {
		return nil, err
	}
	lbl.SetHAlign(gtk.ALIGN_END)
	grid.Attach(lbl, 0, row, 1, 1)
	values = []struct{ value, key string }{
		{"Default", CFG_THEME_VARIANT_SYSTEM_VALUE},
		{"Light", CFG_THEME_VARIANT_LIGHT_VALUE},
		{"Dark", CFG_THEME_VARIANT_DARK_VALUE},
	}
	cbThemeVariant, err := CreateNameValueCombo(values)
	if err != nil {
		return nil, err
	}
	bh.Bind(CFG_THEME_VARIANT_KEY, cbThemeVariant, "active-id", glib.SETTINGS_BIND_DEFAULT)
	grid.Attach(cbThemeVariant, 1, row, 1, 1)
	row++

	//Background Image
	lbl, err = gtk.LabelNew("Background image")
	if err != nil {
		return nil, err
	}
	grid.Attach(lbl, 0, row, 1, 1)

	fcbImage, err := gtk.FileChooserButtonNew("Select Image", gtk.FILE_CHOOSER_ACTION_OPEN)
	if err != nil {
		return nil, err
	}
	fcbImage.SetHExpand(true)
	ff, err := gtk.FileFilterNew()
	if err != nil {
		return nil, err
	}
	ff.SetName("All Image Files")
	ff.AddMimeType("image/jpeg")
	ff.AddMimeType("image/png")
	ff.AddMimeType("image/bmp")
	fcbImage.AddFilter(ff)
	ff, err = gtk.FileFilterNew()
	if err != nil {
		return nil, err
	}
	ff.AddPattern("*")
	ff.SetName("All Files")
	fcbImage.AddFilter(ff)
	filename := appSettings.settings.GetString(CFG_BACKGROUND_IMAGE_KEY)
	if _, err := os.Stat(filename); !os.IsNotExist(err) {
		// log.Println(spew.Sprintf("File %q found", filename))
		fcbImage.SetFilename(filename)
	}
	_, err = fcbImage.Connect("file-set", func(fcb *gtk.FileChooserButton) {
		selectedFilename := fcb.GetFilename()
		if _, err := os.Stat(selectedFilename); !os.IsNotExist(err) {
			appSettings.settings.SetString(CFG_BACKGROUND_IMAGE_KEY, selectedFilename)
		}
	})
	if err != nil {
		return nil, err
	}

	btnReset, err := SetupButtonWithThemedImage("edit-delete-symbolic")
	if err != nil {
		return nil, err
	}
	btnReset.SetTooltipText("Reset background image")
	_, err = btnReset.Connect("clicked", func(btn *gtk.Button) {
		fcbImage.UnselectAll()
		appSettings.settings.Reset(CFG_BACKGROUND_IMAGE_KEY)
	})
	if err != nil {
		return nil, err
	}

	values = []struct{ value, key string }{
		{"Scale", CFG_BACKGROUND_IMAGE_MODE_SCALE_VALUE},
		{"Tile", CFG_BACKGROUND_IMAGE_MODE_TILE_VALUE},
		{"Center", CFG_BACKGROUND_IMAGE_MODE_CENTER_VALUE},
		{"Stretch", CFG_BACKGROUND_IMAGE_MODE_STRETCH_VALUE},
	}
	cbImageMode, err := CreateNameValueCombo(values)
	if err != nil {
		return nil, err
	}
	bh.Bind(CFG_BACKGROUND_IMAGE_MODE_KEY, cbImageMode, "active-id", glib.SETTINGS_BIND_DEFAULT)

	// Background image settings only enabled if transparency is enabled
	bh.Bind(CFG_ENABLE_TRANSPARENCY_KEY, fcbImage, "sensitive", glib.SETTINGS_BIND_DEFAULT)
	bh.Bind(CFG_ENABLE_TRANSPARENCY_KEY, btnReset, "sensitive", glib.SETTINGS_BIND_DEFAULT)
	bh.Bind(CFG_ENABLE_TRANSPARENCY_KEY, cbImageMode, "sensitive", glib.SETTINGS_BIND_DEFAULT)

	bChooser, err := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 2)
	if err != nil {
		return nil, err
	}
	bChooser.Add(fcbImage)
	bChooser.Add(btnReset)

	bImage, err := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 6)
	if err != nil {
		return nil, err
	}
	bImage.Add(bChooser)
	bImage.Add(cbImageMode)
	grid.Attach(bImage, 1, row, 1, 1)
	row++

	box.Add(grid)

	cbRightSidebar, err := gtk.CheckButtonNewWithLabel("Place the sidebar on the right")
	if err != nil {
		return nil, err
	}
	bh.Bind(CFG_SIDEBAR_RIGHT, cbRightSidebar, "active", glib.SETTINGS_BIND_DEFAULT)
	box.Add(cbRightSidebar)

	cbTitleShowWhenSingle, err := gtk.CheckButtonNewWithLabel("Show the terminal title even if it's the only terminal")
	if err != nil {
		return nil, err
	}
	bh.Bind(CFG_TERMINAL_TITLE_SHOW_WHEN_SINGLE_KEY, cbTitleShowWhenSingle, "active", glib.SETTINGS_BIND_DEFAULT)
	box.Add(cbTitleShowWhenSingle)

	cbUseTabs, err := gtk.CheckButtonNewWithLabel("Use tabs instead of sidebar (Application restart required)")
	if err != nil {
		return nil, err
	}
	bh.Bind(CFG_USE_TABS_KEY, cbUseTabs, "active", glib.SETTINGS_BIND_DEFAULT)
	box.Add(cbUseTabs)

	// Disclaimer
	wdg, err := gtk.SeparatorNew(gtk.ORIENTATION_HORIZONTAL)
	if err != nil {
		return nil, err
	}
	box.Add(wdg)
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintln("******************************************************"))
	buf.WriteString(fmt.Sprintln("Design of this preference page taken from " +
		"<a href=\"https://github.com/gnunn1/tilix\">Tilix</a> project."))
	buf.WriteString(fmt.Sprintln("Settings here mainly for demonstration purpose and"))
	buf.WriteString(fmt.Sprintln("have minimal impact to application."))
	buf.WriteString(fmt.Sprintln("******************************************************"))
	lbl, err = gtk.LabelNew("")
	if err != nil {
		return nil, err
	}
	lbl.SetMarkup(buf.String())
	lbl.SetJustify(gtk.JUSTIFY_CENTER)
	box.PackEnd(lbl, true, true, 0)

	_, err = box.Connect("destroy", func(b *gtk.Box) {
		bh.Unbind()
	})
	if err != nil {
		return nil, err
	}

	return &box.Container, nil
}

// getProfileSettings create GlibSettings object with change event
// connected to specific indexed profile[profileID].
func getProfileSettings(appStore *SettingsStore, profileID string, changed func()) (*SettingsStore, error) {
	pathSuffix := fmt.Sprintf(PROFILE_SCHEMA_SUFFIX_PATH, profileID)
	store, err := appStore.GetChildSettingsStore(PROFILE_SCHEMA_SUFFIX_ID, pathSuffix, changed)
	if err != nil {
		return nil, err
	}
	return store, nil
}

// addProfilePage build UI on the top of profile taken from GlibSettings.
func addProfilePage(win *gtk.ApplicationWindow, profileID string, initProfileName *string, appSettings *SettingsStore,
	list *PreferenceRowList, lbSide *gtk.ListBox, pages *gtk.Stack, selectNew bool, profileChanged func()) error {

	prefRow, err := PreferenceRowNew(profileID,
		"Profile", nil, true)
	if err != nil {
		return err
	}
	page, profileName, err := ProfilePreferencesNew(win, appSettings, list,
		profileID, prefRow, initProfileName, profileChanged)
	if err != nil {
		return err
	}
	prefRow.SetName(profileName)
	prefRow.Page = page
	pages.AddTitled(page, profileID, "Profile")
	list.Append(prefRow)
	index := list.GetLastProfileListIndex()
	lbSide.Insert(prefRow.Row, index+1)
	lbSide.ShowAll()
	pages.ShowAll()
	if selectNew {
		lbSide.SelectRow(prefRow.Row)
	}
	return nil
}

// Create preference dialog with "Sources" page, where controls
// being bound to GLib Setting object to save/restore functionality.
func ProfilePreferencesNew(win *gtk.ApplicationWindow, appSettings *SettingsStore,
	list *PreferenceRowList, profileID string, prefRow *PreferenceRow,
	initProfileName *string, profileChanged func()) (*gtk.Container, string, error) {

	sw, err := gtk.ScrolledWindowNew(nil, nil)
	if err != nil {
		return nil, "", err
	}
	sw.SetPolicy(gtk.POLICY_NEVER, gtk.POLICY_AUTOMATIC)
	//SetScrolledWindowPropogatedHeight(sw, true)

	profileSettings, err := getProfileSettings(appSettings, profileID, profileChanged)
	if err != nil {
		return nil, "", err
	}

	grid, err := gtk.GridNew()
	if err != nil {
		return nil, "", err
	}
	grid.SetColumnSpacing(12)
	grid.SetRowSpacing(6)
	grid.SetHAlign(gtk.ALIGN_FILL)
	row := 0

	var lbl *gtk.Label

	appBH := appSettings.NewBindingHelper()
	profileBH := profileSettings.NewBindingHelper()

	// Profile name
	lbl, err = SetupLabelJustifyLeft("Profile name")
	if err != nil {
		return nil, "", err
	}
	grid.Attach(lbl, 0, row, 1, 1)

	edProfileName, err := gtk.EntryNew()
	if err != nil {
		return nil, "", err
	}
	edProfileName.SetHExpand(true)
	edProfileName.SetHAlign(gtk.ALIGN_FILL)
	profileBH.Bind(CFG_PROFILE_NAME, edProfileName, "text", glib.SETTINGS_BIND_DEFAULT)
	timer := time.AfterFunc(time.Millisecond*500, func() {
		_, err := glib.IdleAdd(func() {
			name, err := edProfileName.GetText()
			if err != nil {
				log.Fatal(err)
			}
			prefRow.SetName(name)
		})
		if err != nil {
			log.Fatal(err)
		}
	})
	_, err = edProfileName.Connect("changed", func(v *gtk.Entry, tmr *time.Timer) {
		tmr.Stop()
		tmr.Reset(time.Millisecond * 500)
	}, timer)
	if err != nil {
		return nil, "", err
	}
	if initProfileName != nil {
		edProfileName.SetText(*initProfileName)
	}
	grid.Attach(edProfileName, 1, row, 1, 1)
	row++

	// Profile name
	lbl, err = SetupLabelJustifyLeft("Profile description")
	if err != nil {
		return nil, "", err
	}
	grid.Attach(lbl, 0, row, 1, 1)

	edProfileDescription, err := gtk.EntryNew()
	if err != nil {
		return nil, "", err
	}
	edProfileDescription.SetHExpand(true)
	edProfileDescription.SetHAlign(gtk.ALIGN_FILL)
	profileBH.Bind(CFG_PROFILE_DESCRIPTION, edProfileDescription, "text", glib.SETTINGS_BIND_DEFAULT)
	grid.Attach(edProfileDescription, 1, row, 1, 1)
	row++

	box2, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 6)
	if err != nil {
		return nil, "", err
	}
	SetAllMargins(box2, 18)
	box2.Add(grid)
	vp, err := gtk.ViewportNew(nil, nil)
	if err != nil {
		return nil, "", err
	}
	vp.Add(box2)

	sw.Add(vp)
	_, err = sw.Connect("destroy", func(b gtk.IWidget) {
		appBH.Unbind()
		profileBH.Unbind()
	})
	if err != nil {
		return nil, "", err
	}

	// Disclaimer
	wdg, err := gtk.SeparatorNew(gtk.ORIENTATION_HORIZONTAL)
	if err != nil {
		return nil, "", err
	}
	box2.Add(wdg)
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintln("******************************************************"))
	buf.WriteString(fmt.Sprintln("Design of this preference page demonstrate multi-profile"))
	buf.WriteString(fmt.Sprintln("settings dialog. Application allows to create any number"))
	buf.WriteString(fmt.Sprintln("of profiles (based on same schema) and stores it in"))
	buf.WriteString(fmt.Sprintln("glib.GSettings indexed with unique identifier."))
	buf.WriteString(fmt.Sprintln("******************************************************"))
	lbl, err = gtk.LabelNew("")
	if err != nil {
		return nil, "", err
	}
	lbl.SetMarkup(buf.String())
	lbl.SetJustify(gtk.JUSTIFY_CENTER)
	box2.PackEnd(lbl, true, true, 0)

	_, err = box2.Connect("destroy", func(b *gtk.Box) {
		profileBH.Unbind()
		log.Println("Destroy box")
	})
	if err != nil {
		return nil, "", err
	}

	name := profileSettings.settings.GetString(CFG_PROFILE_NAME)
	return &sw.Container, name, nil
}

// Verify, that GLib Setting's schema is installed, otherwise return false.
func checkSchemaSettingsIsInstalled(app *gtk.Application) (bool, error) {
	parentWin := app.GetActiveWindow()
	// Verify that GSettingsSchema is installed
	schemaSource := glib.SettingsSchemaSourceGetDefault()
	if schemaSource == nil {
		title := "<span weight='bold' size='larger'>Schema settings configuration error</span>"
		paragraphs := []*DialogParagraph{NewDialogParagraph("No one GTK+ schema settings is found.").
			SetJustify(gtk.JUSTIFY_CENTER).SetHorizAlign(gtk.ALIGN_CENTER)}
		paragraphs = append(paragraphs, NewDialogParagraph("Please install xml schema and repeat operation.").
			SetJustify(gtk.JUSTIFY_CENTER).SetHorizAlign(gtk.ALIGN_CENTER))

		buttons := []DialogButton{
			{"_OK", gtk.RESPONSE_OK, false, nil},
		}

		dialog, err := SetupMessageDialog(parentWin, title, "", paragraphs, buttons, nil)
		if err != nil {
			return false, err
		}
		dialog.Run(false)
		return false, nil
	}
	schema := schemaSource.Lookup(SETTINGS_SCHEMA_ID, false)
	if schema == nil {
		title := "<span weight='bold' size='larger'>Schema settings configuration error</span>"
		paragraphs := []*DialogParagraph{NewDialogParagraph(fmt.Sprintf("GTK+ schema %q is not found.", SETTINGS_SCHEMA_ID)).
			SetJustify(gtk.JUSTIFY_CENTER).SetHorizAlign(gtk.ALIGN_CENTER)}
		paragraphs = append(paragraphs, NewDialogParagraph("Please install xml schema and repeat operation.").
			SetJustify(gtk.JUSTIFY_CENTER).SetHorizAlign(gtk.ALIGN_CENTER))

		buttons := []DialogButton{
			{"_OK", gtk.RESPONSE_OK, false, nil},
		}

		dialog, err := SetupMessageDialog(parentWin, title, "", paragraphs, buttons, nil)
		if err != nil {
			return false, err
		}
		dialog.Run(false)
		return false, nil
	}
	return true, nil
}

// Create sophisticated multi-page preference dialog
// with save/restore functionallity to/from the GLib Setting object.
func createPreferenceDialog(settingsID, settingsPath string,
	mainWin *gtk.ApplicationWindow) (*gtk.ApplicationWindow, error) {

	app, err := mainWin.GetApplication()
	if err != nil {
		return nil, err
	}
	win, err := gtk.ApplicationWindowNew(app)
	if err != nil {
		return nil, err
	}

	// Settings
	win.SetTitle("Preferences")
	win.SetTransientFor(mainWin)
	win.SetDestroyWithParent(true)
	win.SetShowMenubar(false)
	appSettings, err := NewSettingsStore(settingsID, settingsPath, nil)
	if err != nil {
		return nil, err
	}

	// Create window header
	hbMain, err := SetupHeader("", "", true)
	if err != nil {
		return nil, err
	}
	hbMain.SetHExpand(true)

	hbSide, err := SetupHeader("Preferences", "", false)
	if err != nil {
		return nil, err
	}
	hbSide.SetHExpand(false)

	bTitle, err := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0)
	if err != nil {
		return nil, err
	}
	bTitle.Add(hbSide)
	sTitle, err := gtk.SeparatorNew(gtk.ORIENTATION_VERTICAL)
	if err != nil {
		return nil, err
	}
	bTitle.Add(sTitle)
	bTitle.Add(hbMain)

	win.SetTitlebar(bTitle)

	var list = PreferenceRowListNew()

	// Create Stack and boxes
	pages, err := gtk.StackNew()
	if err != nil {
		return nil, err
	}
	pages.SetHExpand(true)
	pages.SetVExpand(true)

	// Create ListBox
	lbSide, err := gtk.ListBoxNew()
	if err != nil {
		return nil, err
	}
	lbSide.SetCanFocus(true)
	lbSide.SetSelectionMode(gtk.SELECTION_BROWSE)
	lbSide.SetVExpand(true)

	profileChanged := make(chan struct{})
	var once sync.Once
	changedFunc := func() {
		once.Do(func() {
			close(profileChanged)
		})
	}

	var pr *PreferenceRow

	profileSettingsArray := appSettings.NewSettingsArray(CFG_BACKUP_LIST)
	profileList := profileSettingsArray.GetArrayIDs()
	if len(profileList) == 0 {
		profileID, err := profileSettingsArray.AddNode()
		if err != nil {
			return nil, err
		}
		profileName := profileID
		if i, err := strconv.Atoi(profileID); err == nil {
			profileName = strconv.Itoa(i + 1)
		}
		err = addProfilePage(win, profileID, &profileName, appSettings, list,
			lbSide, pages, false, changedFunc)
		if err != nil {
			return nil, err
		}
	} else {
		for _, profileID := range profileList {
			err = addProfilePage(win, profileID, nil, appSettings, list,
				lbSide, pages, false, changedFunc)
			if err != nil {
				return nil, err
			}
		}
	}

	gp, err := GlobalPreferencesNew(appSettings, &mainWin.ActionMap)
	if err != nil {
		return nil, err
	}
	pages.AddTitled(gp, "Global", "Global")
	pr, err = PreferenceRowNew("Global", "Global", gp, false)
	if err != nil {
		return nil, err
	}
	list.Append(pr)
	lbSide.Add(pr.Row)

	ap, err := AppearancePreferencesNew(appSettings)
	if err != nil {
		return nil, err
	}
	pages.AddTitled(ap, "Appearance", "Appearance")
	pr, err = PreferenceRowNew("Appearance", "Appearance", gp, false)
	if err != nil {
		return nil, err
	}
	list.Append(pr)
	lbSide.Add(pr.Row)

	sw, err := gtk.ScrolledWindowNew(nil, nil)
	if err != nil {
		return nil, err
	}
	sw.Add(lbSide)
	sw.SetPolicy(gtk.POLICY_NEVER, gtk.POLICY_AUTOMATIC)
	sw.SetShadowType(gtk.SHADOW_NONE)
	sw.SetSizeRequest(220, -1)

	bSide, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)
	if err != nil {
		return nil, err
	}
	bSide.Add(sw)

	bButtons, err := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0)
	if err != nil {
		return nil, err
	}
	SetAllMargins(bButtons, 6)
	bSide.Add(bButtons)
	btnAddProfile, err := SetupButtonWithThemedImage("list-add-symbolic")
	if err != nil {
		return nil, err
	}
	btnAddProfile.SetTooltipText("Add profile")
	_, err = btnAddProfile.Connect("clicked", func() {
		profileID, err := profileSettingsArray.AddNode()
		if err != nil {
			log.Fatal(err)
		}

		profileName := profileID
		if i, err := strconv.Atoi(profileID); err == nil {
			profileName = strconv.Itoa(i + 1)
		}
		err = addProfilePage(win, profileID, &profileName, appSettings, list,
			lbSide, pages, true, changedFunc)
		if err != nil {
			log.Fatal(err)
		}
	})
	if err != nil {
		return nil, err
	}
	bButtons.PackStart(btnAddProfile, false, false, 0)

	updateBtnDeleteProfileSensitive := func(deleteBtn *gtk.Button, row *gtk.ListBoxRow) {
		var pr *PreferenceRow
		if row != nil {
			pr = list.Get(row.Native())
			pages.SetVisibleChildName(pr.ID)
			hbMain.SetTitle(pr.Title)
		}
		deleteBtn.SetSensitive(pr != nil && pr.Profile && list.GetProfileCount() > 1)
	}

	btnDeleteProfile, err := SetupButtonWithThemedImage("list-remove-symbolic")
	if err != nil {
		return nil, err
	}
	btnDeleteProfile.SetTooltipText("Delete profile")
	_, err = btnDeleteProfile.Connect("clicked", func() {
		responseYes, err := QuestionDialog(&win.Window, "Delete selected profile?",
			[]*DialogParagraph{NewDialogParagraph("Press YES to delete profile.")}, true)
		if err != nil {
			log.Fatal(err)
		}

		if responseYes {
			sr := lbSide.GetSelectedRow()
			sri := sr.GetIndex()
			pr := list.Get(sr.Native())
			if pr.Profile {
				profileID := pr.ID
				profileSettings, err := getProfileSettings(appSettings, profileID, changedFunc)
				if err != nil {
					log.Fatal(err)
				}
				err = profileSettingsArray.DeleteNode(profileSettings, profileID)
				if err != nil {
					log.Fatal(err)
				}
				nsr := lbSide.GetRowAtIndex(sri + 1)
				lbSide.SelectRow(nsr)
				pages.Remove(pr.Page)
				list.Delete(sr.Native())
				pr.Page.Destroy()
				sr.Destroy()
				updateBtnDeleteProfileSensitive(btnDeleteProfile, lbSide.GetSelectedRow())
			}
		}
	})
	if err != nil {
		return nil, err
	}
	bButtons.PackStart(btnDeleteProfile, false, false, 0)

	_, err = lbSide.Connect("row-selected", func(lb *gtk.ListBox, row *gtk.ListBoxRow) {
		updateBtnDeleteProfileSensitive(btnDeleteProfile, row)
	})
	if err != nil {
		return nil, err
	}

	box, err := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0)
	if err != nil {
		return nil, err
	}
	box.Add(bSide)
	sep, err := gtk.SeparatorNew(gtk.ORIENTATION_HORIZONTAL)
	if err != nil {
		return nil, err
	}
	box.Add(sep)
	box.Add(pages)

	win.Add(box)

	sgSide, err := gtk.SizeGroupNew(gtk.SIZE_GROUP_HORIZONTAL)
	if err != nil {
		return nil, err
	}
	sgSide.AddWidget(hbSide)
	sgSide.AddWidget(bSide)

	sgMain, err := gtk.SizeGroupNew(gtk.SIZE_GROUP_HORIZONTAL)
	if err != nil {
		return nil, err
	}
	sgMain.AddWidget(hbMain)
	sgMain.AddWidget(pages)

	// Set initial title
	hbMain.SetTitle("Global")

	return win, nil
}

func createApp(data *FullscreenGlobalData) (*gtk.Application, error) {
	app, err := gtk.ApplicationNew(APP_SCHEMA_ID, glib.APPLICATION_FLAGS_NONE)
	if err != nil {
		return nil, err
	}

	_, err = app.Application.Connect("activate", func(application *gtk.Application) {
		win, err := gtk.ApplicationWindowNew(application)
		if err != nil {
			log.Fatal(err)
		}
		win.SetTitle("Example")
		win.SetDefaultSize(900, 600)

		_, err2 := win.Connect("destroy", func(window *gtk.ApplicationWindow) {
			application, err := window.GetApplication()
			if err != nil {
				log.Fatal(err)
			}
			application.Quit()
		})
		if err2 != nil {
			log.Fatal(err2)
		}

		var act glib.IAction

		act, err = createQuitAction(&win.Window)
		if err != nil {
			log.Fatal(err)
		}
		win.AddAction(act)

		appSettings, err := NewSettingsStore(SETTINGS_SCHEMA_ID, SETTINGS_SCHEMA_PATH, nil)
		if err != nil {
			log.Fatal(err)
		}
		act, err = createAboutAction(&win.Window, appSettings)
		if err != nil {
			log.Fatal(err)
		}
		win.AddAction(act)

		act, err = createDialogAction1(&win.Window)
		if err != nil {
			log.Fatal(err)
		}
		win.AddAction(act)

		act, err = createDialogAction2(&win.Window)
		if err != nil {
			log.Fatal(err)
		}
		win.AddAction(act)

		act, err = createDialogAction3(&win.Window)
		if err != nil {
			log.Fatal(err)
		}
		win.AddAction(act)

		act, err = createDialogAction4(&win.Window)
		if err != nil {
			log.Fatal(err)
		}
		win.AddAction(act)

		act, err = createDialogAction5(&win.Window)
		if err != nil {
			log.Fatal(err)
		}
		win.AddAction(act)

		act, err = createCheckBoxAction()
		if err != nil {
			log.Fatal(err)
		}
		win.AddAction(act)

		act, err = createChooseAction()
		if err != nil {
			log.Fatal(err)
		}
		win.AddAction(act)

		act, err = createPreferenceAction(win)
		if err != nil {
			log.Fatal(err)
		}
		win.AddAction(act)

		menu, err := createMenuModelForPopover()
		if err != nil {
			log.Fatal(err)
		}
		btn, err := SetupMenuButtonWithThemedImage("open-menu-symbolic")
		if err != nil {
			log.Fatal(err)
		}
		btn.SetUsePopover(true)
		btn.SetMenuModel(menu)

		hdr, err := SetupHeader(APP_TITLE, "", true)
		if err != nil {
			log.Fatal(err)
		}
		hdr.PackEnd(btn)

		win.SetTitlebar(hdr)

		box, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)
		if err != nil {
			log.Fatal(err)
		}

		_, err = box.Connect("destroy", func() {
			log.Println("Destroy box")
		})
		if err != nil {
			log.Fatal(err)
		}

		eb, err := createEventBox(data)
		if err != nil {
			log.Fatal(err)
		}
		data.EventBox = eb
		rev, err := createRevealer(data)
		if err != nil {
			log.Fatal(err)
		}
		data.Revealer = rev
		eb.Add(rev)
		eb.SetSizeRequest(-1, 1)
		eb.Hide()

		box.PackStart(eb, false, false, 0)

		menu2, err := createMenuModelForMenuBar()
		if err != nil {
			log.Fatal(err)
		}

		menuBar, err := gtk.MenuBarFromModelNew(menu2)
		if err != nil {
			log.Fatal(err)
		}

		box.PackStart(menuBar, false, false, 0)

		tbx, err := createToolbar()
		if err != nil {
			log.Fatal(err)
		}

		box.PackStart(tbx, false, false, 0)

		div, err := gtk.SeparatorNew(gtk.ORIENTATION_HORIZONTAL)
		if err != nil {
			log.Fatal(err)
		}
		box.PackStart(div, false, false, 0)

		sw, err := gtk.ScrolledWindowNew(nil, nil)
		// sw, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)
		// sw, err := gtk.FrameNew("sdfsdfsdf")
		// sw, err := gtk.ButtonNew()
		if err != nil {
			log.Fatal(err)
		}

		vp, err := gtk.ViewportNew(nil, nil)
		if err != nil {
			log.Fatal(err)
		}
		box2, err := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0)
		if err != nil {
			log.Fatal(err)
		}
		_, err = box2.Connect("destroy", func() {
			log.Println("Destroy box2")
		})
		if err != nil {
			log.Fatal(err)
		}

		vp.Add(box2)
		sw.Add(vp)
		_, err = sw.Connect("destroy", func() {
			log.Println("Destroy sw")
		})
		if err != nil {
			log.Fatal(err)
		}

		box.Add(sw)

		win.Add(box)

		act, err = createFullscreenAction(&win.Window, data)
		if err != nil {
			log.Fatal(err)
		}
		win.AddAction(act)

		win.ShowAll()

		// Run code, when app message queue becomes empty.
		if !appSettings.settings.GetBoolean(CFG_DONT_SHOW_ABOUT_ON_STARTUP_KEY) {
			_, err2 := glib.IdleAdd(func() {
				action := win.LookupAction("AboutAction")
				if action != nil {
					action.Activate(nil)
				}
			})
			if err2 != nil {
				log.Fatal(err2)
			}
		}
	})
	if err != nil {
		return nil, err
	}
	return app, nil
}

func main() {

	gtk.Init(nil)

	appRunMode = AppRegularRun
	for {
		data := &FullscreenGlobalData{}

		app, err := createApp(data)
		if err != nil {
			log.Fatal(err)
		}

		// Run application.
		app.Run([]string{})

		// If request was made to reload app, then we re-run app
		// without exiting (can be used for changing app UI language).
		if appRunMode == AppRegularRun {
			break
		} else if appRunMode == AppRunReload {
			appRunMode = AppRegularRun
		}
	}
}
