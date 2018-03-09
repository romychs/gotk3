// Same copyright and license as the rest of the files in this project
// This file contains style related functions and structures

// This file includes wrapers for symbols included since GTK 3.12, and
// and should not be included in a build intended to target any older GTK
// versions.  To target an older build, such as 3.10, use
// 'go build -tags gtk_3_10'.  Otherwise, if no build tags are used, GTK 3.12
// is assumed and this file is built.
// +build !gtk_3_6,!gtk_3_8,!gtk_3_10

package gtk

// #include <gtk/gtk.h>
// #include "gtk.go.h"
import "C"

/*
 * GtkMenuButton
 */

// void
// gtk_menu_button_set_use_popover (GtkMenuButton *menu_button,
//                                  gboolean use_popover);
func (v *MenuButton) SetUsePopover(usePopover bool) {
	C.gtk_menu_button_set_use_popover(v.native(), gbool(usePopover))
}
