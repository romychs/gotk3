//+build gtk_3_6 gtk_3_8 gtk_3_10 gtk_3_12 gtk_3_14 gtk_3_16 gtk_3_18

package gdk

// #cgo pkg-config: gdk-3.0
// #include <gdk/gdk.h>
import "C"
import (
	"runtime"
	"unsafe"

	"github.com/d2r2/gotk3/glib"
)

// Grab is a wrapper around gdk_device_grab().
func (v *Device) Grab(w *Window, ownership GrabOwnership, ownerEvents bool, eventMask EventMask,
	cursor *Cursor, time uint32) GrabStatus {

	ret := C.gdk_device_grab(
		v.native(),
		w.native(),
		C.GdkGrabOwnership(ownership),
		gbool(ownerEvents),
		C.GdkEventMask(eventMask),
		cursor.native(),
		C.guint32(time),
	)
	return GrabStatus(ret)
}

// GetClientPointer is a wrapper around gdk_device_manager_get_client_pointer().
func (v *DeviceManager) GetClientPointer() (*Device, error) {
	c := C.gdk_device_manager_get_client_pointer(v.native())
	if c == nil {
		return nil, nilPtrErr
	}

	return &Device{glib.Take(unsafe.Pointer(c))}, nil
}

// ListDevices is a wrapper around gdk_device_manager_list_devices().
// Returned list is wrapped to return *gdk.Device elements.
func (v *DeviceManager) ListDevices(tp DeviceType) *glib.List {
	clist := C.gdk_device_manager_list_devices(v.native(), C.GdkDeviceType(tp))

	glist := glib.WrapList(uintptr(unsafe.Pointer(clist)))
	glist.DataWrapper(func(ptr unsafe.Pointer) interface{} {
		d := wrapDevice(glib.Take(ptr))
		return d
	})

	if glist != nil {
		runtime.SetFinalizer(glist, func(glist *glib.List) {
			glist.Free()
		})
	}

	return glist
}

// Ungrab is a wrapper around gdk_device_ungrab().
func (v *Device) Ungrab(time uint32) {
	C.gdk_device_ungrab(v.native(), C.guint32(time))
}

// GetDeviceManager is a wrapper around gdk_display_get_device_manager().
func (v *Display) GetDeviceManager() (*DeviceManager, error) {
	c := C.gdk_display_get_device_manager(v.native())
	if c == nil {
		return nil, nilPtrErr
	}

	return &DeviceManager{glib.Take(unsafe.Pointer(c))}, nil
}

// GetScreen is a wrapper around gdk_display_get_screen().
func (v *Display) GetScreen(screenNum int) (*Screen, error) {
	c := C.gdk_display_get_screen(v.native(), C.gint(screenNum))
	if c == nil {
		return nil, nilPtrErr
	}

	return &Screen{glib.Take(unsafe.Pointer(c))}, nil
}
