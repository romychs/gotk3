// +build !glib_2_40

// See: https://developer.gnome.org/glib/2.42/api-index-2-42.html

package glib

// #cgo pkg-config: glib-2.0 gobject-2.0
// #include <gio/gio.h>
// #include <glib.h>
// #include <glib-object.h>
// #include "glib.go.h"
import "C"

// NotificationPriority is a representation of GLib's GNotificationPriority.
type NotificationPriority int

const (
	NOTIFICATION_PRIORITY_NORMAL NotificationPriority = C.G_NOTIFICATION_PRIORITY_NORMAL
	NOTIFICATION_PRIORITY_LOW    NotificationPriority = C.G_NOTIFICATION_PRIORITY_LOW
	NOTIFICATION_PRIORITY_HIGH   NotificationPriority = C.G_NOTIFICATION_PRIORITY_HIGH
	NOTIFICATION_PRIORITY_URGENT NotificationPriority = C.G_NOTIFICATION_PRIORITY_URGENT
)

// SetPriority is a wrapper around g_notification_set_priority().
func (v *Notification) SetPriority(prio NotificationPriority) {
	C.g_notification_set_priority(v.native(), C.GNotificationPriority(prio))
}
