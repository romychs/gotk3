package libnotify

// #cgo pkg-config: libnotify
// #include <stdlib.h>
// #include <libnotify/notify.h>
// #include <libnotify.go.h>
import "C"
import (
	"errors"
	"unsafe"

	"github.com/d2r2/gotk3/gdk"
	"github.com/d2r2/gotk3/glib"
)

var (
	errNilPtr = errors.New("cgo returned unexpected nil pointer")
)

/*
 * Type conversions
 */

func gobool(b C.gboolean) bool {
	if b != 0 {
		return true
	}
	return false
}

func goString(cstr *C.gchar) string {
	return C.GoString((*C.char)(cstr))
}

// NotifyUrgency is a representation of Libnotify's NotifyUrgency.
type NotifyUrgency int

const (
	NOTIFY_URGENCY_LOW      NotifyUrgency = C.NOTIFY_URGENCY_LOW
	NOTIFY_URGENCY_NORMAL   NotifyUrgency = C.NOTIFY_URGENCY_NORMAL
	NOTIFY_URGENCY_CRITICAL NotifyUrgency = C.NOTIFY_URGENCY_CRITICAL
)

// Default timeouts
const (
	NOTIFY_EXPIRES_DEFAULT = -1
	NOTIFY_EXPIRES_NEVER   = 0
)

// NotifyNotification is a representation of NotifyNotification.
type NotifyNotification struct {
	*glib.Object
}

// native() returns a pointer to the underlying NotifyNotification.
func (v *NotifyNotification) native() *C.NotifyNotification {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toNotifyNotification(ptr)
}

func marshalNotifyNotification(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	return wrapNotifyNotification(glib.Take(unsafe.Pointer(c))), nil
}

func wrapNotifyNotification(obj *glib.Object) *NotifyNotification {
	return &NotifyNotification{obj}
}

// NotifyNotificationNew is a wrapper around notify_notification_new().
func NotifyNotificationNew(summary, body, icon string) (*NotifyNotification, error) {
	var cicon *C.char
	if icon != "" {
		cicon = C.CString(icon)
		defer C.free(unsafe.Pointer(cicon))
	}

	var cbody *C.char
	if body != "" {
		cbody = C.CString(body)
		defer C.free(unsafe.Pointer(cbody))
	}

	cstr := C.CString(summary)
	defer C.free(unsafe.Pointer(cstr))

	c := C.notify_notification_new((*C.gchar)(cstr), (*C.gchar)(cbody), (*C.gchar)(cicon))
	if c == nil {
		return nil, errNilPtr
	}
	return wrapNotifyNotification(glib.Take(unsafe.Pointer(c))), nil
}

// Update is a wrapper around notify_notification_update().
func (v *NotifyNotification) Update(summary, body, icon string) error {
	var cicon *C.char
	if icon != "" {
		cicon = C.CString(icon)
		defer C.free(unsafe.Pointer(cicon))
	}

	var cbody *C.char
	if body != "" {
		cbody = C.CString(body)
		defer C.free(unsafe.Pointer(cbody))
	}

	cstr := C.CString(summary)
	defer C.free(unsafe.Pointer(cstr))

	if !gobool(C.notify_notification_update(v.native(),
		(*C.gchar)(cstr), (*C.gchar)(cbody), (*C.gchar)(cicon))) {
		return errors.New("error updating NotifyNotification")
	}
	return nil
}

// Show is a wrapper around notify_notification_show().
func (v *NotifyNotification) Show() error {
	var err *C.GError
	c := C.notify_notification_show(v.native(), &err)
	if c == 0 {
		defer C.g_error_free(err)
		return errors.New(goString(err.message))
	}
	return nil
}

// SetAppName is a wrapper around notify_notification_set_app_name().
func (v *NotifyNotification) SetAppName(appName string) {
	cstr := C.CString(appName)
	defer C.free(unsafe.Pointer(cstr))

	C.notify_notification_set_app_name(v.native(), (*C.gchar)(cstr))
}

// SetTimeout is a wrapper around notify_notification_set_timeout().
// You can use NOTIFY_EXPIRES_DEFAULT or NOTIFY_EXPIRES_NEVER,
// either specific specific timeout amount in milliseconds.
func (v *NotifyNotification) SetTimeout(timeout int) {
	C.notify_notification_set_timeout(v.native(), C.int(timeout))
}

// SetCategory is a wrapper around notify_notification_set_category().
func (v *NotifyNotification) SetCategory(category string) {
	cstr := C.CString(category)
	defer C.free(unsafe.Pointer(cstr))

	C.notify_notification_set_category(v.native(), (*C.gchar)(cstr))
}

// SetUrgency is a wrapper around notify_notification_set_urgency().
func (v *NotifyNotification) SetUrgency(urgency NotifyUrgency) {
	C.notify_notification_set_urgency(v.native(), C.NotifyUrgency(urgency))
}

// SetImageFromPixbuf is a wrapper around notify_notification_set_image_from_pixbuf().
func (v *NotifyNotification) SetImageFromPixbuf(pixbuf *gdk.Pixbuf) {
	C.notify_notification_set_image_from_pixbuf(v.native(),
		C.toGdkPixbuf(unsafe.Pointer(pixbuf.Native())))
}

// Close is a wrapper around notify_notification_close().
func (v *NotifyNotification) Close() error {
	var err *C.GError
	c := C.notify_notification_close(v.native(), &err)
	if c == 0 {
		defer C.g_error_free(err)
		return errors.New(goString(err.message))
	}
	return nil
}

// Init is a wrapper around notify_init() and must be called before any
// other Libnotify calls and is used to initialize everything necessary.
func Init(appName string) error {
	cstr := C.CString(appName)
	defer C.free(unsafe.Pointer(cstr))

	if !gobool(C.notify_init((*C.gchar)(cstr))) {
		return errors.New("error initializing libnotify")
	}

	return nil
}

// Uninit is a wrapper around notify_uninit().
func Uninit() {
	C.notify_uninit()
}

// IsInitted is a wrapper around notify_is_initted().
func IsInitted() bool {
	c := C.notify_is_initted()
	return gobool(c)
}

// GetAppName is a wrapper around notify_get_app_name().
func GetAppName() string {
	c := C.notify_get_app_name()
	return goString(c)
}

// SetAppName is a wrapper around notify_set_app_name().
func SetAppName(appName string) {
	cstr := C.CString(appName)
	defer C.free(unsafe.Pointer(cstr))

	C.notify_set_app_name((*C.gchar)(cstr))
}
