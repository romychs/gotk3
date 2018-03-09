package main

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/d2r2/gotk3/glib"
	"github.com/d2r2/gotk3/gtk"
	"github.com/davecgh/go-spew/spew"
)

// ===================================================================================
// ************************* UTILITIES SECTION START *********************************
// ===================================================================================
/*	Copy this section to separate file in real application as utilities functions to simplify
	creation of GTK+ 3 components and controls, including menus, dialog boxes, messages,
	app settings save and restore and so on...
*/

func setupHeader(title, subtitle string, showCloseButton bool) (*gtk.HeaderBar, error) {
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

func setupMenuItemWithIcon(label, detailedAction string, icon *glib.Icon) (*glib.MenuItem, error) {
	mi, err := glib.MenuItemNew(label, detailedAction)
	if err != nil {
		return nil, err
	}
	//mi.SetAttributeValue("verb-icon", iconNameVar)
	mi.SetIcon(icon)
	return mi, nil
}

func setupMenuItemWithThemedIcon(label, detailedAction, iconName string) (*glib.MenuItem, error) {
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

func setupToolButton(themedIconName, label string) (*gtk.ToolButton, error) {
	var btn *gtk.ToolButton
	var img *gtk.Image
	var err error
	if themedIconName != "" {
		img, err = gtk.ImageNew()
		if err != nil {
			return nil, err
		}
		img.SetFromIconName(themedIconName, gtk.ICON_SIZE_BUTTON)
	}

	btn, err = gtk.ToolButtonNew(img, label)
	if err != nil {
		return nil, err
	}
	return btn, nil
}

func setupButtonWithThemedImage(themedIconName string) (*gtk.Button, error) {
	img, err := gtk.ImageNew()
	if err != nil {
		return nil, err
	}
	img.SetFromIconName(themedIconName, gtk.ICON_SIZE_BUTTON)

	btn, err := gtk.ButtonNew()
	if err != nil {
		return nil, err
	}

	btn.Add(img)

	return btn, nil
}

func setupMenuButtonWithThemedImage(themedIconName string) (*gtk.MenuButton, error) {
	img, err := gtk.ImageNew()
	if err != nil {
		return nil, err
	}
	img.SetFromIconName(themedIconName, gtk.ICON_SIZE_BUTTON)

	btn, err := gtk.MenuButtonNew()
	if err != nil {
		return nil, err
	}

	btn.Add(img)

	return btn, nil
}

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

type DialogButton struct {
	Text      string
	Response  gtk.ResponseType
	Default   bool
	Customize func(button *gtk.Button) error
}

func getActiveWindow(win *gtk.Window) (*gtk.Window, error) {
	app, err := win.GetApplication()
	if err != nil {
		return nil, err
	}
	return app.GetActiveWindow(), nil
}

func isResponseYes(response gtk.ResponseType) bool {
	return response == gtk.RESPONSE_YES
}

func isResponseNo(response gtk.ResponseType) bool {
	return response == gtk.RESPONSE_NO
}

func isResponseNone(response gtk.ResponseType) bool {
	return response == gtk.RESPONSE_NONE
}

func isResponseOk(response gtk.ResponseType) bool {
	return response == gtk.RESPONSE_OK
}

func isResponseCancel(response gtk.ResponseType) bool {
	return response == gtk.RESPONSE_CANCEL
}

func isResponseReject(response gtk.ResponseType) bool {
	return response == gtk.RESPONSE_REJECT
}

func isResponseClose(response gtk.ResponseType) bool {
	return response == gtk.RESPONSE_CLOSE
}

func isResponseDeleteEvent(response gtk.ResponseType) bool {
	return response == gtk.RESPONSE_DELETE_EVENT
}

func printResponse(response gtk.ResponseType) {
	if isResponseNo(response) {
		log.Println("Dialog result = NO")
	} else if isResponseYes(response) {
		log.Println("Dialog result = YES")
	} else if isResponseNone(response) {
		log.Println("Dialog result = NONE")
	} else if isResponseOk(response) {
		log.Println("Dialog result = OK")
	} else if isResponseReject(response) {
		log.Println("Dialog result = REJECT")
	} else if isResponseCancel(response) {
		log.Println("Dialog result = CANCEL")
	} else if isResponseClose(response) {
		log.Println("Dialog result = CLOSE")
	} else if isResponseDeleteEvent(response) {
		log.Println("Dialog result = DELETE_EVENT")
	}
}

func setupMessageDialog(parent *gtk.Window, markupTitle string, messages []string,
	addButtons []DialogButton, addExtraControls func(area *gtk.Box) error) (*gtk.MessageDialog, error) {
	active, err := getActiveWindow(parent)
	if err != nil {
		return nil, err
	}

	dlg, err := gtk.MessageDialogNew(active, /*gtk.DIALOG_MODAL|*/
		gtk.DIALOG_USE_HEADER_BAR, gtk.MESSAGE_WARNING, gtk.BUTTONS_NONE, nil, nil)
	if err != nil {
		return nil, err
	}
	dlg.SetTransientFor(active)
	dlg.SetMarkup(markupTitle)

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

	for _, msg := range messages {
		lbl, err := gtk.LabelNew(msg)
		if err != nil {
			return nil, err
		}
		lbl.SetHAlign(gtk.ALIGN_START)
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

	return dlg, nil
}

func runMessageDialog(parent *gtk.Window, markupTitle string, messages []string, ignoreCloseBox bool,
	addButtons []DialogButton, addExtraControls func(area *gtk.Box) error) (gtk.ResponseType, error) {
	dlg, err := setupMessageDialog(parent, markupTitle, messages, addButtons, addExtraControls)
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

func setupDialog(parent *gtk.Window, messageType gtk.MessageType, userHeaderbar bool,
	title string, text []string, textAlign gtk.Align, addButtons []DialogButton,
	addExtraControls func(area *gtk.Box) error) (*gtk.Dialog, error) {

	active, err := getActiveWindow(parent)
	if err != nil {
		return nil, err
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

	/*
		style, err := btn.GetStyleContext()
		if err != nil {
			return
		}
		//style.AddClass("suggested-action")
		style.AddClass("destructive-action")
		//style.AddClass("flat")
		//style.AddClass("circular")
		//style.AddClass("text-button")
		style.AddClass("image-button")
		style.RemoveClass("destructive-action")
	*/

	col := 1
	row := 0

	for _, msg := range text {
		lbl, err := gtk.LabelNew(msg)
		if err != nil {
			return nil, err
		}
		lbl.SetHAlign(textAlign)
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

	/*
		if userHeaderbar {
			hdr, err := dlg.GetHeaderBar()
			if err != nil {
				return nil, err
			}

			_, w := hdr.GetPreferredWidth()
			dlg.Resize(w, 100)
		}
	*/

	_, w := dlg.GetPreferredWidth()
	_, h := dlg.GetPreferredHeight()
	dlg.Resize(w, h)

	return dlg, nil
}

func runDialog(parent *gtk.Window, messageType gtk.MessageType, userHeaderbar bool,
	title string, text []string, textAlign gtk.Align, ignoreCloseBox bool, addButtons []DialogButton,
	addExtraControls func(area *gtk.Box) error) (gtk.ResponseType, error) {
	dlg, err := setupDialog(parent, messageType, userHeaderbar, title,
		text, textAlign, addButtons, addExtraControls)
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

func getActionNameAndState(act *glib.SimpleAction) (string, *glib.Variant, error) {
	name, err := act.GetName()
	if err != nil {
		return "", nil, err
	}
	state := act.GetState()
	return name, state, nil
}

type Binding struct {
	Key      string
	Object   glib.IObject
	Property string
	Flags    glib.SettingsBindFlags
}

/**

	Code taken from https://github.com/gnunn1/tilix project

* Bookkeeping class that keps track of objects which are
* binded to a GSettings object so they can be unbinded later. it
* also supports the concept of deferred bindings where a binding
* can be added but is not actually attached to a Settings object
* until one is set.
*/
type BindingHelper struct {
	bindings []Binding
	settings *glib.Settings
}

func BindingHelperNew(settings *glib.Settings) *BindingHelper {
	bh := &BindingHelper{settings: settings}
	return bh
}

/**
 * Setting a new GSettings object will cause this class to unbind
 * previously set bindings and re-bind to the new settings automatically.
 */
func (v *BindingHelper) SetSettings(value *glib.Settings) {
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
			v.settings.Bind(b.Key, b.Object, b.Property, b.Flags)
		}
	}
}

/**
 * Adds a binding to the list
 */
func (v *BindingHelper) addBind(key string, object glib.IObject, property string, flags glib.SettingsBindFlags) {
	v.bindings = append(v.bindings, Binding{key, object, property, flags})
}

/**
 * Add a binding to list and binds to Settings if it is set.
 */
func (v *BindingHelper) Bind(key string, object glib.IObject, property string, flags glib.SettingsBindFlags) {
	v.addBind(key, object, property, flags)
	if v.settings != nil {
		v.settings.Bind(key, object, property, flags)
	}
}

/**
 * Unbinds all added binds from settings object
 */
func (v *BindingHelper) Unbind() {
	for _, b := range v.bindings {
		v.settings.Unbind(b.Object, b.Property)
	}
}

/**
 * Unbinds all bindings and clears list of bindings.
 */
func (v *BindingHelper) Clear() {
	v.Unbind()
	v.bindings = nil
}

/**
 * Sets margins of a widget to the passed values
 */
func SetMargins(widget gtk.IWidget, left int, top int, right int, bottom int) {
	w := widget.GetWidget()
	w.SetMarginStart(left)
	w.SetMarginTop(top)
	w.SetMarginEnd(right)
	w.SetMarginBottom(bottom)
}

/**
 * Sets all margins of a widget to the same value
 */
func SetAllMargins(widget gtk.IWidget, margin int) {
	SetMargins(widget, margin, margin, margin, margin)
}

/**
 * Appends multiple values to a row in a list store
 */
func AppendValues(ls *gtk.ListStore, values ...string) *gtk.TreeIter {
	iter := ls.Append()
	for i := 0; i < len(values); i++ {
		ls.SetValue(iter, i, values[i])
	}
	return iter
}

/**
 * Creates a combobox that holds a set of name/value pairs
 * where the name is displayed.
 */
func CreateNameValueCombo(keyValues []struct{ value, key string }) (*gtk.ComboBox, error) {
	ls, err := gtk.ListStoreNew(glib.TYPE_STRING, glib.TYPE_STRING)
	if err != nil {
		return nil, err
	}

	for _, item := range keyValues {
		AppendValues(ls, item.value, item.key)
	}

	cb, err := gtk.ComboBoxNewWithModel(ls)
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

func GetGtkVersion() (magor, minor, micro uint) {
	magor = gtk.GetMajorVersion()
	minor = gtk.GetMinorVersion()
	micro = gtk.GetMicroVersion()
	return
}

// ===================================================================================
// ************************* UTILITIES SECTION END *********************************
// ===================================================================================

// String constants used as titles/identifiers
var (
	APP_ID string = "org.d2r2.gotk3.cool_app_1"

	APP_TITLE string = "Cool app (GTK+ 3 UI adaptation for golang based on imporved GOTK3)"

	//Preference Constants
	SETTINGS_ID string = APP_ID + ".Settings"

	SETTINGS_AUTO_HIDE_MOUSE_KEY                   string = "auto-hide-mouse"
	SETTINGS_PROMPT_ON_NEW_SESSION_KEY             string = "prompt-on-new-session"
	SETTINGS_ENABLE_TRANSPARENCY_KEY               string = "enable-transparency"
	SETTINGS_CLOSE_WITH_LAST_SESSION_KEY           string = "close-with-last-session"
	SETTINGS_APP_TITLE_KEY                         string = "app-title"
	SETTINGS_CONTROL_CLICK_TITLE_KEY               string = "control-click-titlebar"
	SETTINGS_INHERIT_WINDOW_STATE_KEY              string = "new-window-inherit-state"
	SETTINGS_USE_OVERLAY_SCROLLBAR_KEY             string = "use-overlay-scrollbar"
	SETTINGS_TERMINAL_FOCUS_FOLLOWS_MOUSE_KEY      string = "focus-follow-mouse"
	SETTINGS_MIDDLE_CLICK_CLOSE_KEY                string = "middle-click-close"
	SETTINGS_CONTROL_SCROLL_ZOOM_KEY               string = "control-scroll-zoom"
	SETTINGS_WINDOW_SAVE_STATE_KEY                 string = "window-save-state"
	SETTINGS_NOTIFY_ON_PROCESS_COMPLETE_KEY        string = "notify-on-process-complete"
	SETTINGS_PASTE_ADVANCED_DEFAULT_KEY            string = "paste-advanced-default"
	SETTINGS_UNSAFE_PASTE_ALERT_KEY                string = "unsafe-paste-alert"
	SETTINGS_STRIP_FIRST_COMMENT_CHAR_ON_PASTE_KEY string = "paste-strip-first-char"
	SETTINGS_COPY_ON_SELECT_KEY                    string = "copy-on-select"
	SETTINGS_WINDOW_STYLE_KEY                      string = "window-style"

	SETTINGS_TERMINAL_TITLE_STYLE_KEY          string = "terminal-title-style"
	SETTINGS_TERMINAL_TITLE_STYLE_VALUE_NORMAL string = "normal"
	SETTINGS_TERMINAL_TITLE_STYLE_VALUE_SMALL  string = "small"
	SETTINGS_TERMINAL_TITLE_STYLE_VALUE_NONE   string = "none"

	SETTINGS_TAB_POSITION_KEY string = "tab-position"

	// Theme Settings
	SETTINGS_THEME_VARIANT_KEY          string = "theme-variant"
	SETTINGS_THEME_VARIANT_SYSTEM_VALUE string = "system"
	SETTINGS_THEME_VARIANT_LIGHT_VALUE  string = "light"
	SETTINGS_THEME_VARIANT_DARK_VALUE   string = "dark"

	SETTINGS_BACKGROUND_IMAGE_KEY                string = "background-image"
	SETTINGS_BACKGROUND_IMAGE_MODE_KEY           string = "background-image-mode"
	SETTINGS_BACKGROUND_IMAGE_MODE_SCALE_VALUE   string = "scale"
	SETTINGS_BACKGROUND_IMAGE_MODE_TILE_VALUE    string = "tile"
	SETTINGS_BACKGROUND_IMAGE_MODE_CENTER_VALUE  string = "center"
	SETTINGS_BACKGROUND_IMAGE_MODE_STRETCH_VALUE string = "stretch"

	SETTINGS_SIDEBAR_RIGHT string = "sidebar-on-right"

	SETTINGS_TERMINAL_TITLE_SHOW_WHEN_SINGLE_KEY string = "terminal-title-show-when-single"

	SETTINGS_USE_TABS_KEY string = "use-tabs"
)

type GUI_Data struct {
	EventBox             *gtk.EventBox
	Revealer             *gtk.Revealer
	FullscreeAction      *glib.Action
	RightTopMenuButton   *gtk.MenuButton
	InFullscreenEventBox bool
	Timer                *time.Timer
}

// Deomnstration of Popover menu creation with different kind of menus:
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
	item, err = setupMenuItemWithThemedIcon("About", "win.AboutAction", "help-about-symbolic")
	if err != nil {
		return nil, err
	}
	section.AppendItem(item)
	item, err = setupMenuItemWithThemedIcon("Preference", "win.PreferenceAction", "preferences-other-symbolic")
	if err != nil {
		return nil, err
	}
	section.AppendItem(item)
	item, err = setupMenuItemWithThemedIcon("Fullscreen", "win.FullscreenAction", "view-fullscreen-symbolic")
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
	item, err = setupMenuItemWithThemedIcon("Quit application", "win.QuitAction", "application-exit-symbolic")
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
	item, err = setupMenuItemWithIcon("Preference dialog demo", "win.PreferenceAction", &icon.Icon)
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
	icon2, err := glib.ThemedIconNew("application-exit-symbolic")
	if err != nil {
		return nil, err
	}
	item, err = setupMenuItemWithIcon("Quit application", "win.QuitAction", &icon2.Icon)
	if err != nil {
		return nil, err
	}
	section.AppendItem(item)
	main.AppendSection("", section)

	return rootMenu, nil
}

/*
func createToolbar() (*gtk.Toolbar, error) {
	tbx, err := gtk.ToolbarNew()
	if err != nil {
		return nil, err
	}

	dirName := "./icons"
	files, err := ioutil.ReadDir(dirName)
	if err != nil {
		return nil, err
	}

	for _, f := range files {
		fullName, err := filepath.Abs(path.Join(dirName, f.Name()))
		if err != nil {
			return nil, err
		}
		if strings.HasSuffix(fullName, "32x32.gif") {
			img, err := gtk.ImageNewFromFile(fullName)
			if err != nil {
				return nil, err
			}
			titem, err := gtk.ToolButtonNew(img, "")
			if err != nil {
				return nil, err
			}
			tbx.Add(titem)
		}
	}
	return tbx, nil
}
*/

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

func createCheckBoxAction() (glib.IAction, error) {
	v, err := glib.VariantBooleanNew(true)
	if err != nil {
		return nil, err
	}
	act, err := glib.SimpleActionStatefullNew("CheckBoxAction", nil, v)
	if err != nil {
		log.Fatal(err)
	}
	if act == nil {
		log.Fatal(errors.New("error"))
	}

	act.Connect("activate", func(action *glib.SimpleAction, param *glib.Variant) {
		name, state, err := getActionNameAndState(action)
		if err != nil {
			return
		}
		log.Println(spew.Sprintf("%v action activated with current state %v and args %v",
			name, state, param))

		if state != nil && state.IsOfType(glib.VARIANT_TYPE_BOOLEAN) {
			state, err = glib.VariantBooleanNew(!state.GetBoolean())
			if err != nil {
				return
			}
			action.ChangeState(state)
		}
	})

	return act, nil
}

func createChooseAction() (glib.IAction, error) {
	v, err := glib.VariantStringNew("green")
	if err != nil {
		return nil, err
	}
	act, err := glib.SimpleActionStatefullNew("ChooseColor", glib.VARIANT_TYPE_STRING, v)
	if err != nil {
		log.Fatal(err)
	}
	if act == nil {
		log.Fatal(errors.New("error"))
	}

	act.Connect("activate", func(action *glib.SimpleAction, param *glib.Variant) {
		name, state, err := getActionNameAndState(action)
		if err != nil {
			return
		}
		log.Println(spew.Sprintf("%v action activated with current state %v and args %v",
			name, state, param))

		act.ChangeState(param)
	})

	return act, nil
}

func createFullscreenAction(win *gtk.Window, data *GUI_Data) (glib.IAction, error) {
	v, err := glib.VariantBooleanNew(false)
	if err != nil {
		return nil, err
	}
	act, err := glib.SimpleActionStatefullNew("FullscreenAction", nil, v)
	if err != nil {
		log.Fatal(err)
	}
	data.FullscreeAction = &act.Action

	act.Connect("activate", func(action *glib.SimpleAction, param *glib.Variant) {
		name, state, err := getActionNameAndState(action)
		if err != nil {
			return
		}
		log.Println(spew.Sprintf("%v action activated with current state %v and args %v",
			name, state, param))

		if state != nil && state.IsOfType(glib.VARIANT_TYPE_BOOLEAN) {
			state, err = glib.VariantBooleanNew(!state.GetBoolean())
			if err != nil {
				return
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

	return act, nil
}

func createQuitAction(win *gtk.Window) (glib.IAction, error) {
	act, err := glib.SimpleActionNew("QuitAction", nil)
	if err != nil {
		return nil, err
	}

	act.Connect("activate", func(action *glib.SimpleAction, param *glib.Variant) {
		name, state, err := getActionNameAndState(action)
		if err != nil {
			return
		}
		log.Println(spew.Sprintf("%v action activated with current state %v and args %v",
			name, state, param))

		application, err := win.GetApplication()
		if err != nil {
			return
		}
		application.Quit()
	})

	return act, nil
}

func createAboutAction(win *gtk.Window) (glib.IAction, error) {
	act, err := glib.SimpleActionNew("AboutAction", nil)
	if err != nil {
		return nil, err
	}

	act.Connect("activate", func(action *glib.SimpleAction, param *glib.Variant) {
		name, state, err := getActionNameAndState(action)
		if err != nil {
			return
		}
		log.Println(spew.Sprintf("%v action activated with current state %v and args %v",
			name, state, param))

		dlg, err := gtk.AboutDialogNew()
		if err != nil {
			return
		}

		dlg.SetAuthors([]string{"Denis Dyakov <denis.dyakov@gmail.com>"})
		dlg.SetProgramName("Cool app")
		dlg.SetLogoIconName("face-cool-symbolic")
		dlg.SetVersion("0.1")

		var buf bytes.Buffer
		major, minor, micro := GetGtkVersion()
		buf.WriteString(fmt.Sprintln(fmt.Sprintf("GTK+ version %d.%d.%d", major, minor, micro)))

		buf.WriteString(fmt.Sprintln("This application built for education purpose and compose"))
		buf.WriteString(fmt.Sprintln("modern practices around GTK3 user interface composing."))
		buf.WriteString(fmt.Sprintln("TODO: write infromation about interfaces app demostrate."))
		dlg.SetComments(buf.String())

		dlg.SetTransientFor(win)
		dlg.SetModal(true)
		dlg.ShowNow()
	})

	return act, nil
}

func createDialogAction1(win *gtk.Window) (glib.IAction, error) {
	act, err := glib.SimpleActionNew("DialogAction1", nil)
	if err != nil {
		return nil, err
	}

	act.Connect("activate", func(action *glib.SimpleAction, param *glib.Variant) {
		name, state, err := getActionNameAndState(action)
		if err != nil {
			return
		}
		log.Println(spew.Sprintf("%v action activated with current state %v and args %v",
			name, state, param))

		title := "<span weight='bold' size='larger'>Configuration Issue Detected (demonstration)</span>"
		text := []string{
			"Message dialog (old style) demonstration.",
			"This message based on GtkMessageDialog functionality.\n" +
				"GtkMessageDialog doesn't support setting of image\n" +
				"via SetImage as message type specification since 3.22.\n" +
				"So, that's why there is not warning icon either something similar here.",
			"There appears to be an issue with the configuration of the application.\n" +
				"This issue is not serious, but correcting it will improve your experience.",
		}

		buttons := []DialogButton{
			{"_OK", gtk.RESPONSE_OK, false, nil},
		}

		response, err := runMessageDialog(win, title, text, false, buttons,
			func(area *gtk.Box) error {
				lbl, err := gtk.LabelNew("Click the link below for more information:")
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

		printResponse(response)
	})

	return act, nil
}

func createDialogAction2(win *gtk.Window) (glib.IAction, error) {
	act, err := glib.SimpleActionNew("DialogAction2", nil)
	if err != nil {
		return nil, err
	}

	act.Connect("activate", func(action *glib.SimpleAction, param *glib.Variant) {
		name, state, err := getActionNameAndState(action)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(spew.Sprintf("%v action activated with current state %v and args %v",
			name, state, param))

		title := "Configuration Issue Detected (demonstration)"
		text := []string{
			"Message dialog (new style) demonstration.",
			"This message based on GtkDialog functionality.",
			"There appears to be an issue with the configuration of the application.\n" +
				"This issue is not serious, but correcting it will improve your experience.",
		}
		buttons := []DialogButton{
			{"_OK", gtk.RESPONSE_OK, false, nil},
		}

		response, err := runDialog(win, gtk.MESSAGE_WARNING, true, title, text, gtk.ALIGN_START, false, buttons,
			func(area *gtk.Box) error {
				lbl, err := gtk.LabelNew("Click the link below for more information:")
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

		printResponse(response)
	})

	return act, nil
}

func createDialogAction3(win *gtk.Window) (glib.IAction, error) {
	act, err := glib.SimpleActionNew("DialogAction3", nil)
	if err != nil {
		return nil, err
	}

	act.Connect("activate", func(action *glib.SimpleAction, param *glib.Variant) {
		name, state, err := getActionNameAndState(action)
		if err != nil {
			return
		}
		log.Println(spew.Sprintf("%v action activated with current state %v and args %v",
			name, state, param))

		title := "<span weight='bold' size='larger'>Configuration Issue Detected (demonstration)</span>"
		text := []string{
			"Message dialog (old style) demonstration.",
			"This message based on GtkMessageDialog functionality.\n" +
				"GtkMessageDialog doesn't support setting of image\n" +
				"via SetImage as message type specification since 3.22.\n" +
				"So, that's why there is not warning icon either something similar here.",
			"There appears to be an issue with the configuration of the application.\n" +
				"This issue is not serious, but correcting it will improve your experience.",
		}

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

		response, err := runMessageDialog(win, title, text, true, buttons,
			func(area *gtk.Box) error {
				lbl, err := gtk.LabelNew("Click the link below for more information:")
				if err != nil {
					return err
				}
				lbl.SetHAlign(gtk.ALIGN_START)
				area.Add(lbl)
				/*
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
				*/
				return nil
			})
		if err != nil {
			log.Fatal(err)
		}

		printResponse(response)
	})

	return act, nil
}

func createDialogAction4(win *gtk.Window) (glib.IAction, error) {
	act, err := glib.SimpleActionNew("DialogAction4", nil)
	if err != nil {
		return nil, err
	}

	act.Connect("activate", func(action *glib.SimpleAction, param *glib.Variant) {
		name, state, err := getActionNameAndState(action)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(spew.Sprintf("%v action activated with current state %v and args %v",
			name, state, param))

		title := "Choose option (demonstration)"
		text := []string{
			"Press Yes to start processing.",
			"Note: processing might takes significant amount of time.",
		}
		buttons := []DialogButton{
			{"_Yes", gtk.RESPONSE_YES, true, func(btn *gtk.Button) error {
				style, err := btn.GetStyleContext()
				if err != nil {
					return err
				}
				style.AddClass("suggested-action")
				//style.AddClass("destructive-action")
				return nil
			}},
			{"_No", gtk.RESPONSE_NO, false, nil},
		}

		response, err := runDialog(win, gtk.MESSAGE_QUESTION, true, title, text, gtk.ALIGN_CENTER, true, buttons,
			nil)
		if err != nil {
			log.Fatal(err)
		}

		printResponse(response)
	})

	return act, nil
}

func createDialogAction5(win *gtk.Window) (glib.IAction, error) {
	act, err := glib.SimpleActionNew("DialogAction5", nil)
	if err != nil {
		return nil, err
	}

	act.Connect("activate", func(action *glib.SimpleAction, param *glib.Variant) {
		name, state, err := getActionNameAndState(action)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(spew.Sprintf("%v action activated with current state %v and args %v",
			name, state, param))

		title := "Choose option (demonstration)"
		text := []string{
			"Press Yes to start processing.",
			"Note: processing might takes significant amount of time.",
		}
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

		response, err := runDialog(win, gtk.MESSAGE_QUESTION, false, title, text, gtk.ALIGN_CENTER, true, buttons,
			nil)
		if err != nil {
			log.Fatal(err)
		}

		printResponse(response)
	})

	return act, nil
}

func createPreferenceAction(win *gtk.Window) (glib.IAction, error) {
	act, err := glib.SimpleActionNew("PreferenceAction", nil)
	if err != nil {
		return nil, err
	}

	act.Connect("activate", func(action *glib.SimpleAction, param *glib.Variant) {
		name, state, err := getActionNameAndState(action)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(spew.Sprintf("%v action activated with current state %v and args %v",
			name, state, param))

		app, err := win.GetApplication()
		if err != nil {
			log.Fatal(err)
		}

		found, err := checkSchemaSettingsIsInstalled(app)
		if err != nil {
			log.Fatal(err)
		}

		if found {

			win, err := createPreferenceDialog(app)
			if err != nil {
				log.Fatal(err)
			}

			win.ShowAll()
			win.Show()

			win.Connect("destroy", func(window *gtk.ApplicationWindow) {
				/*
					application, err := window.GetApplication()
					if err != nil {
						return
					}
				*/
				window.Destroy()
			})
		}

	})

	return act, nil
}

func createRevealer(data *GUI_Data) (*gtk.Revealer, error) {
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

	btn, err := setupMenuButtonWithThemedImage("open-menu-symbolic")
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

	btn3, err := setupToolButton("view-restore-symbolic", "Leave ssdfsdf")
	if err != nil {
		return nil, err
	}
	//	btn2.SetTitle("asdasdas")
	btn3.SetActionName("win.FullscreenAction")

	hdr, err := setupHeader(APP_TITLE, "(fullscreen mode)", false)
	if err != nil {
		log.Fatal(err)
	}

	hdr.PackEnd(btn3)
	hdr.PackEnd(btn2)
	hdr.PackEnd(btn)

	box.PackStart(hdr, false, false, 0)

	return rev, nil

}

func manageFullscreenHeaderHideTimer(data *GUI_Data, dur time.Duration, start bool) {
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

func showFullscreenHeader(data *GUI_Data, show bool) {
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

func createEventBox(data *GUI_Data) (*gtk.EventBox, error) {
	eb, err := gtk.EventBoxNew()
	if err != nil {
		return nil, err
	}
	dur := time.Millisecond * 1500
	eb.Connect("enter-notify-event", func() {
		log.Println("EventBox enter notify")

		manageFullscreenHeaderHideTimer(data, dur, false)
		showFullscreenHeader(data, true)
	})
	eb.Connect("leave-notify-event", func() {
		glib.IdleAdd(func() {
			log.Println("EventBox leave notify")

			manageFullscreenHeaderHideTimer(data, dur, true)
		})
	})

	return eb, nil
}

type GenericPreferenceRow struct {
	gtk.ListBoxRow
	Name  string
	Title string
}

func GenericPreferenceRowNew(name, title string) (*GenericPreferenceRow, error) {
	lbr, err := gtk.ListBoxRowNew()
	if err != nil {
		return nil, err
	}

	v := &GenericPreferenceRow{ListBoxRow: *lbr, Name: name, Title: title}
	lbl, err := gtk.LabelNew(name)
	if err != nil {
		return nil, err
	}
	lbl.SetHAlign(gtk.ALIGN_START)
	SetAllMargins(lbl, 6)
	v.Add(lbl)

	return v, nil
}

/**
 * Global preferences page *
 */
type GlobalPreferences struct {
	gtk.Box
}

func GlobalPreferencesNew(gsSettings *glib.Settings) (*GlobalPreferences, error) {
	box, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 6)
	if err != nil {
		return nil, err
	}

	SetAllMargins(box, 18)

	bh := BindingHelperNew(gsSettings)

	lblBehavior, err := gtk.LabelNew(fmt.Sprintf("<b>%s</b>", "Behavior"))
	if err != nil {
		return nil, err
	}
	lblBehavior.SetUseMarkup(true)
	lblBehavior.SetHAlign(gtk.ALIGN_START)
	box.Add(lblBehavior)

	//Prompt on new session
	cbPrompt, err := gtk.CheckButtonNewWithLabel("Prompt when creating a new session")
	if err != nil {
		return nil, err
	}
	bh.Bind(SETTINGS_PROMPT_ON_NEW_SESSION_KEY, cbPrompt, "active", glib.SETTINGS_BIND_DEFAULT)
	box.Add(cbPrompt)

	//Focus follows the mouse
	cbFocusMouse, err := gtk.CheckButtonNewWithLabel("Focus a terminal when the mouse moves over it")
	if err != nil {
		return nil, err
	}
	bh.Bind(SETTINGS_TERMINAL_FOCUS_FOLLOWS_MOUSE_KEY, cbFocusMouse, "active", glib.SETTINGS_BIND_DEFAULT)
	box.Add(cbFocusMouse)

	//Auto hide the mouse
	cbAutoHideMouse, err := gtk.CheckButtonNewWithLabel("Autohide the mouse pointer when typing")
	if err != nil {
		return nil, err
	}
	bh.Bind(SETTINGS_AUTO_HIDE_MOUSE_KEY, cbAutoHideMouse, "active", glib.SETTINGS_BIND_DEFAULT)
	box.Add(cbAutoHideMouse)

	//middle click closes the terminal
	cbMiddleClickClose, err := gtk.CheckButtonNewWithLabel("Close terminal by clicking middle mouse button on title")
	if err != nil {
		return nil, err
	}
	bh.Bind(SETTINGS_MIDDLE_CLICK_CLOSE_KEY, cbMiddleClickClose, "active", glib.SETTINGS_BIND_DEFAULT)
	box.Add(cbMiddleClickClose)

	//zoom in/out terminal with scroll wheel
	cbControlScrollZoom, err := gtk.CheckButtonNewWithLabel("Zoom the terminal using <Control> and scroll wheel")
	if err != nil {
		return nil, err
	}
	bh.Bind(SETTINGS_CONTROL_SCROLL_ZOOM_KEY, cbControlScrollZoom, "active", glib.SETTINGS_BIND_DEFAULT)
	box.Add(cbControlScrollZoom)

	//require control modifier when clicking title
	cbControlClickTitle, err := gtk.CheckButtonNewWithLabel("Require the <Control> modifier to edit title on click")
	if err != nil {
		return nil, err
	}
	bh.Bind(SETTINGS_CONTROL_CLICK_TITLE_KEY, cbControlClickTitle, "active", glib.SETTINGS_BIND_DEFAULT)
	box.Add(cbControlClickTitle)

	//Closing of last session closes window
	cbCloseWithLastSession, err := gtk.CheckButtonNewWithLabel("Close window when last session is closed")
	if err != nil {
		return nil, err
	}
	bh.Bind(SETTINGS_CLOSE_WITH_LAST_SESSION_KEY, cbCloseWithLastSession, "active", glib.SETTINGS_BIND_DEFAULT)
	box.Add(cbCloseWithLastSession)

	cbNewWindowInheritState, err := gtk.CheckButtonNewWithLabel("New window inherits directory and profile from active terminal")
	if err != nil {
		return nil, err
	}
	bh.Bind(SETTINGS_INHERIT_WINDOW_STATE_KEY, cbNewWindowInheritState, "active", glib.SETTINGS_BIND_DEFAULT)
	box.Add(cbNewWindowInheritState)

	// Save window state (maximized, minimized, fullscreen) between invocations
	cbWindowSaveState, err := gtk.CheckButtonNewWithLabel("Save and restore window state")
	if err != nil {
		return nil, err
	}
	bh.Bind(SETTINGS_WINDOW_SAVE_STATE_KEY, cbWindowSaveState, "active", glib.SETTINGS_BIND_DEFAULT)
	box.Add(cbWindowSaveState)

	/*
		//Show Notifications, only show option if notifications are supported
		if checkVTEFeature(TerminalFeature.EVENT_NOTIFICATION) {
			cbNotify, err := gtk.CheckButtonNewWithLabel("Send desktop notification on process complete")
			if err != nil {
				return nil, err
			}
			bh.Bind(SETTINGS_NOTIFY_ON_PROCESS_COMPLETE_KEY, cbNotify, "active", glib.SETTINGS_BIND_DEFAULT)
			box.Add(cbNotify)
		}

		   //New Instance Options
		   Box bNewInstance = new Box(Orientation.HORIZONTAL, 6);

		   Label lblNewInstance = new Label(_("On new instance"));
		   lblNewInstance.setHalign(Align.END);
		   bNewInstance.add(lblNewInstance);
		   ComboBox cbNewInstance = createNameValueCombo([_("New Window"), _("New Session"), _("Split Right"), _("Split Down"),
		   	_("Focus Window")], SETTINGS_NEW_INSTANCE_MODE_VALUES);
		   bh.bind(SETTINGS_NEW_INSTANCE_MODE_KEY, cbNewInstance, "active-id", glib.SETTINGS_BIND_DEFAULT);
		   bNewInstance.add(cbNewInstance);
		   add(bNewInstance);
	*/

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
	bh.Bind(SETTINGS_PASTE_ADVANCED_DEFAULT_KEY, cbAdvDefault, "active", glib.SETTINGS_BIND_DEFAULT)
	box.Add(cbAdvDefault)

	//Unsafe Paste Warning
	cbUnsafe, err := gtk.CheckButtonNewWithLabel("Warn when attempting unsafe paste")
	if err != nil {
		return nil, err
	}
	bh.Bind(SETTINGS_UNSAFE_PASTE_ALERT_KEY, cbUnsafe, "active", glib.SETTINGS_BIND_DEFAULT)
	box.Add(cbUnsafe)

	//Strip Paste
	cbStrip, err := gtk.CheckButtonNewWithLabel("Strip first character of paste if comment or variable declaration")
	if err != nil {
		return nil, err
	}
	bh.Bind(SETTINGS_STRIP_FIRST_COMMENT_CHAR_ON_PASTE_KEY, cbStrip, "active", glib.SETTINGS_BIND_DEFAULT)
	box.Add(cbStrip)

	//Copy on Select
	cbCopyOnSelect, err := gtk.CheckButtonNewWithLabel("Automatically copy text to clipboard when selecting")
	if err != nil {
		return nil, err
	}
	bh.Bind(SETTINGS_COPY_ON_SELECT_KEY, cbCopyOnSelect, "active", glib.SETTINGS_BIND_DEFAULT)
	box.Add(cbCopyOnSelect)

	box.Connect("destroy", func(b *gtk.Box) {
		bh.Unbind()
	})

	v := &GlobalPreferences{Box: *box}

	return v, nil
}

/**
 * Appearance preferences page
 */
type AppearancePreferences struct {
	gtk.Box
}

func AppearancePreferencesNew(gsSettings *glib.Settings) (*AppearancePreferences, error) {
	box, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 6)
	if err != nil {
		return nil, err
	}

	SetAllMargins(box, 18)

	bh := BindingHelperNew(gsSettings)

	grid, err := gtk.GridNew()
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
	bh.Bind(SETTINGS_WINDOW_STYLE_KEY, cbWindowStyle, "active-id", glib.SETTINGS_BIND_DEFAULT)
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
		{"Normal", SETTINGS_TERMINAL_TITLE_STYLE_VALUE_NORMAL},
		{"Small", SETTINGS_TERMINAL_TITLE_STYLE_VALUE_SMALL},
		{"None", SETTINGS_TERMINAL_TITLE_STYLE_VALUE_NONE},
	}
	cbTitleStyle, err := CreateNameValueCombo(values)
	bh.Bind(SETTINGS_TERMINAL_TITLE_STYLE_KEY, cbTitleStyle, "active-id", glib.SETTINGS_BIND_DEFAULT)
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
	bh.Bind(SETTINGS_TAB_POSITION_KEY, cbTabPosition, "active-id", glib.SETTINGS_BIND_DEFAULT)
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
		{"Default", SETTINGS_THEME_VARIANT_SYSTEM_VALUE},
		{"Light", SETTINGS_THEME_VARIANT_LIGHT_VALUE},
		{"Dark", SETTINGS_THEME_VARIANT_DARK_VALUE},
	}
	cbThemeVariant, err := CreateNameValueCombo(values)
	bh.Bind(SETTINGS_THEME_VARIANT_KEY, cbThemeVariant, "active-id", glib.SETTINGS_BIND_DEFAULT)
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
	filename := gsSettings.GetString(SETTINGS_BACKGROUND_IMAGE_KEY)
	if _, err := os.Stat(filename); !os.IsNotExist(err) {
		// log.Println(spew.Sprintf("File %q found", filename))
		fcbImage.SetFilename(filename)
	}
	fcbImage.Connect("file-set", func(fcb *gtk.FileChooserButton) {
		selectedFilename := fcb.GetFilename()
		if _, err := os.Stat(selectedFilename); !os.IsNotExist(err) {
			gsSettings.SetString(SETTINGS_BACKGROUND_IMAGE_KEY, selectedFilename)
		}
	})

	btnReset, err := setupButtonWithThemedImage("edit-delete-symbolic")
	if err != nil {
		return nil, err
	}
	btnReset.SetTooltipText("Reset background image")
	btnReset.Connect("clicked", func(btn *gtk.Button) {
		fcbImage.UnselectAll()
		gsSettings.Reset(SETTINGS_BACKGROUND_IMAGE_KEY)
	})

	values = []struct{ value, key string }{
		{"Scale", SETTINGS_BACKGROUND_IMAGE_MODE_SCALE_VALUE},
		{"Tile", SETTINGS_BACKGROUND_IMAGE_MODE_TILE_VALUE},
		{"Center", SETTINGS_BACKGROUND_IMAGE_MODE_CENTER_VALUE},
		{"Stretch", SETTINGS_BACKGROUND_IMAGE_MODE_STRETCH_VALUE},
	}
	cbImageMode, err := CreateNameValueCombo(values)
	if err != nil {
		return nil, err
	}
	bh.Bind(SETTINGS_BACKGROUND_IMAGE_MODE_KEY, cbImageMode, "active-id", glib.SETTINGS_BIND_DEFAULT)

	// Background image settings only enabled if transparency is enabled
	bh.Bind(SETTINGS_ENABLE_TRANSPARENCY_KEY, fcbImage, "sensitive", glib.SETTINGS_BIND_DEFAULT)
	bh.Bind(SETTINGS_ENABLE_TRANSPARENCY_KEY, btnReset, "sensitive", glib.SETTINGS_BIND_DEFAULT)
	bh.Bind(SETTINGS_ENABLE_TRANSPARENCY_KEY, cbImageMode, "sensitive", glib.SETTINGS_BIND_DEFAULT)

	bChooser, err := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 2)
	bChooser.Add(fcbImage)
	bChooser.Add(btnReset)

	bImage, err := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 6)
	bImage.Add(bChooser)
	bImage.Add(cbImageMode)
	grid.Attach(bImage, 1, row, 1, 1)
	row++
	/*
	   //Session Name
	   Label lblSessionName = new Label(_("Default session name"));
	   lblSessionName.setHalign(Align.END);
	   grid.attach(lblSessionName, 0, row, 1, 1);

	   Entry eSessionName = new Entry();
	   eSessionName.setHexpand(true);
	   bh.bind(SETTINGS_SESSION_NAME_KEY, eSessionName, "text", GSettingsBindFlags.DEFAULT);
	   if (Version.checkVersion(3, 16, 0).length == 0) {
	       grid.attach(createTitleEditHelper(eSessionName, TitleEditScope.SESSION), 1, row, 1, 1);
	   } else {
	       grid.attach(eSessionName, 1, row, 1, 1);
	   }
	   row++;

	   //Application Title
	   Label lblAppTitle = new Label(_("Application title"));
	   lblAppTitle.setHalign(Align.END);
	   grid.attach(lblAppTitle, 0, row, 1, 1);

	   Entry eAppTitle = new Entry();
	   eAppTitle.setHexpand(true);
	   bh.bind(SETTINGS_APP_TITLE_KEY, eAppTitle, "text", GSettingsBindFlags.DEFAULT);
	   if (Version.checkVersion(3, 16, 0).length == 0) {
	       grid.attach(createTitleEditHelper(eAppTitle, TitleEditScope.WINDOW), 1, row, 1, 1);
	   } else {
	       grid.attach(eAppTitle, 1, row, 1, 1);
	   }
	   row++;
	*/
	box.Add(grid)
	/*
	   //Enable Transparency, only enabled if less then 3.18
	   if (Version.getMajorVersion() <= 3 && Version.getMinorVersion() < 18) {
	       CheckButton cbTransparent = new CheckButton(_("Enable transparency, requires re-start"));
	       bh.bind(SETTINGS_ENABLE_TRANSPARENCY_KEY, cbTransparent, "active", GSettingsBindFlags.DEFAULT);
	       add(cbTransparent);
	   }

	   if (Version.checkVersion(3, 16, 0).length == 0) {
	       CheckButton cbWideHandle = new CheckButton(_("Use a wide handle for splitters"));
	       bh.bind(SETTINGS_ENABLE_WIDE_HANDLE_KEY, cbWideHandle, "active", GSettingsBindFlags.DEFAULT);
	       add(cbWideHandle);
	   }
	*/

	cbRightSidebar, err := gtk.CheckButtonNewWithLabel("Place the sidebar on the right")
	if err != nil {
		return nil, err
	}
	bh.Bind(SETTINGS_SIDEBAR_RIGHT, cbRightSidebar, "active", glib.SETTINGS_BIND_DEFAULT)
	box.Add(cbRightSidebar)

	cbTitleShowWhenSingle, err := gtk.CheckButtonNewWithLabel("Show the terminal title even if it's the only terminal")
	if err != nil {
		return nil, err
	}
	bh.Bind(SETTINGS_TERMINAL_TITLE_SHOW_WHEN_SINGLE_KEY, cbTitleShowWhenSingle, "active", glib.SETTINGS_BIND_DEFAULT)
	box.Add(cbTitleShowWhenSingle)
	/*
	   if (Version.checkVersion(3, GTK_SCROLLEDWINDOW_VERSION, 0).length == 0 && environment.get("GTK_OVERLAY_SCROLLING","1") == "1") {
	       CheckButton cbOverlay = new CheckButton(_("Use overlay scrollbars (Application restart required)"));
	       bh.bind(SETTINGS_USE_OVERLAY_SCROLLBAR_KEY, cbOverlay, "active", GSettingsBindFlags.DEFAULT);
	       add(cbOverlay);
	   }
	*/

	cbUseTabs, err := gtk.CheckButtonNewWithLabel("Use tabs instead of sidebar (Application restart required)")
	bh.Bind(SETTINGS_USE_TABS_KEY, cbUseTabs, "active", glib.SETTINGS_BIND_DEFAULT)
	box.Add(cbUseTabs)

	box.Connect("destroy", func(b *gtk.Box) {
		bh.Unbind()
	})

	v := &AppearancePreferences{Box: *box}

	return v, nil
}

func checkSchemaSettingsIsInstalled(app *gtk.Application) (bool, error) {
	parentWin := app.GetActiveWindow()
	// Verify that GSettingsSchema is installed
	schemaSource := glib.SettingsSchemaSourceGetDefault()
	if schemaSource == nil {
		title := "<span weight='bold' size='larger'>Schema settings configuration error</span>"
		text := []string{
			"No one GTK+ schema settings is found.",
			"Please install xml schema and repeat operation.",
		}

		buttons := []DialogButton{
			{"_OK", gtk.RESPONSE_OK, false, nil},
		}

		_, err := runMessageDialog(parentWin, title, text, false, buttons, nil)
		if err != nil {
			return false, err
		}
		return false, nil
	}
	schema := schemaSource.Lookup(SETTINGS_ID, false)
	if schema == nil {
		title := "<span weight='bold' size='larger'>Schema settings configuration error</span>"
		text := []string{
			fmt.Sprintf("GTK+ schema %q is not found.", SETTINGS_ID),
			"Please install xml schema and repeat operation.",
		}

		buttons := []DialogButton{
			{"_OK", gtk.RESPONSE_OK, false, nil},
		}

		_, err := runMessageDialog(parentWin, title, text, false, buttons, nil)
		if err != nil {
			return false, err
		}
		return false, nil
	}
	return true, nil
}

func createPreferenceDialog(app *gtk.Application) (*gtk.ApplicationWindow, error) {
	parentWin := app.GetActiveWindow()
	win, err := gtk.ApplicationWindowNew(app)
	if err != nil {
		return nil, err
	}

	// Settings
	win.SetTitle("Preferences")
	//win.setTypeHint(WindowTypeHint.DIALOG);
	win.SetTransientFor(parentWin)
	win.SetDestroyWithParent(true)
	win.SetShowMenubar(false)
	gsSettings, err := glib.SettingsNew(SETTINGS_ID)
	if err != nil {
		return nil, err
	}

	// Create window header

	hbMain, err := setupHeader("", "", true)
	if err != nil {
		return nil, err
	}
	hbMain.SetHExpand(true)

	searchButton, err := gtk.ToggleButtonNew()
	img, err := gtk.ImageNew()
	if err != nil {
		return nil, err
	}
	img.SetFromIconName("system-search-symbolic", gtk.ICON_SIZE_MENU)
	searchButton.SetNoShowAll(true)
	hbMain.PackEnd(searchButton)

	hbSide, err := setupHeader("Preferences", "", false)
	if err != nil {
		return nil, err
	}
	hbSide.SetHExpand(false)

	bTitle, err := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0)
	bTitle.Add(hbSide)
	sTitle, err := gtk.SeparatorNew(gtk.ORIENTATION_VERTICAL)
	//sTitle.getStyleContext().addClass("tilix-title-separator");
	bTitle.Add(sTitle)
	bTitle.Add(hbMain)

	win.SetTitlebar(bTitle)
	/*
	   this.addOnNotify(delegate(ParamSpec, ObjectG) {
	       onDecorationLayout();
	   }, "gtk-decoration-layout");
	   onDecorationLayout();
	*/

	var list = make(map[uintptr]interface{})

	//Create Stack and boxes
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
	lbSide.Connect("row-selected", func(lb *gtk.ListBox, row *gtk.ListBoxRow) {
		r := list[row.Native()]
		if r1, ok := r.(*GenericPreferenceRow); ok {
			//log.Println(spew.Sprintf("%+v", r1))
			pages.SetVisibleChildName(r1.Name)
			hbMain.SetTitle(r1.Title)
		}
	})

	gp, err := GlobalPreferencesNew(gsSettings)
	if err != nil {
		return nil, err
	}
	pages.AddTitled(gp, "Global", "Global")
	row, err := GenericPreferenceRowNew("Global", "Global")
	if err != nil {
		return nil, err
	}
	list[row.Native()] = row
	lbSide.Add(row)

	ap, err := AppearancePreferencesNew(gsSettings)
	if err != nil {
		return nil, err
	}
	pages.AddTitled(ap, "Appearance", "Appearance")
	row, err = GenericPreferenceRowNew("Appearance", "Appearance")
	if err != nil {
		return nil, err
	}
	list[row.Native()] = row
	lbSide.Add(row)

	/*
	   QuakePreferences qp = new QuakePreferences(gsSettings, _wayland);
	   pages.addTitled(qp, N_("Quake"), _("Quake"));
	   addNonProfileRow(new GenericPreferenceRow(N_("Quake"), _("Quake")));

	   bmEditor = new GlobalBookmarkEditor();
	   pages.addTitled(bmEditor, N_("Bookmarks"), _("Bookmarks"));
	   addNonProfileRow(new GenericPreferenceRow(N_("Bookmarks"), _("Bookmarks")));

	   ShortcutPreferences sp = new ShortcutPreferences(gsSettings);
	   searchButton.addOnToggled(delegate(ToggleButton button) {
	       sp.toggleShortcutsFind();
	   });
	   pages.addTitled(sp, N_("Shortcuts"), _("Shortcuts"));
	   addNonProfileRow(new GenericPreferenceRow(N_("Shortcuts"), _("Shortcuts")));
	*/

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
	/*
		sep, err := gtk.SeparatorNew(gtk.ORIENTATION_HORIZONTAL)
		if err != nil {
			return nil, err
		}
		bSide.Add(sep)
		bSide.Add(bButtons)
	*/

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

	//Set initial title
	hbMain.SetTitle("Global")

	return win, nil
}

func main() {

	data := &GUI_Data{}

	gtk.Init(nil)
	app, err := gtk.ApplicationNew(APP_ID, glib.APPLICATION_FLAGS_NONE)
	if err != nil {
		log.Fatal(err)
	}

	app.Application.Connect("activate", func(application *gtk.Application) {
		win, err := gtk.ApplicationWindowNew(application)
		if err != nil {
			log.Fatal(err)
		}
		win.SetTitle("Example")
		win.SetDefaultSize(800, 600)

		win.Connect("destroy", func(window *gtk.ApplicationWindow) {
			application, err := window.GetApplication()
			if err != nil {
				return
			}
			application.Quit()
			// gtk.MainQuit()
		})

		var act glib.IAction

		act, err = createQuitAction(&win.Window)
		if err != nil {
			log.Fatal(err)
		}
		win.AddAction(act)

		act, err = createAboutAction(&win.Window)
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

		act, err = createPreferenceAction(&win.Window)
		if err != nil {
			log.Fatal(err)
		}
		win.AddAction(act)

		menu, err := createMenuModelForPopover()
		if err != nil {
			log.Fatal(err)
		}
		btn, err := setupMenuButtonWithThemedImage("open-menu-symbolic")
		if err != nil {
			log.Fatal(err)
		}
		btn.SetUsePopover(true)
		btn.SetMenuModel(menu)

		hdr, err := setupHeader(APP_TITLE, "", true)
		if err != nil {
			log.Fatal(err)
		}

		hdr.PackEnd(btn)

		win.SetTitlebar(hdr)

		//bindModel(po)

		box, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)
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

		win.Add(box)

		act, err = createFullscreenAction(&win.Window, data)
		if err != nil {
			log.Fatal(err)
		}
		win.AddAction(act)

		win.ShowAll()
		//win.SetFocus(btn)
		//win.GrabFocus()

	})

	// thm, err := gtk.IconThemeGetDefault()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// pbf, err := thm.LoadIcon("mail-send-receive-symbolic", gtk.ICON_SIZE_BUTTON, 0)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// Use objcopy utility and debug/elf to embedd and read resources from executable file

	// img, err := gtk.ImageNewFromFile("./ajax-loader-gears_32x32.gif")
	// img, err := gtk.ImageNewFromFile("./dotdot32x32.gif")

	app.Run([]string{})
	// app.Application.Unref()
	// gtk.Main()
}
