package gdk

// #cgo pkg-config: gdk-3.0
// #include <gdk/gdk.h>
import "C"
import (
	"unsafe"

	"github.com/d2r2/gotk3/glib"
)

type PixbufFormat struct {
	format *C.GdkPixbufFormat
}

// native returns a pointer to the underlying GdkPixbuf.
func (v *PixbufFormat) native() *C.GdkPixbufFormat {
	if v == nil {
		return nil
	}
	return v.format
}

// Native returns a pointer to the underlying GdkPixbuf.
func (v *PixbufFormat) Native() uintptr {
	return uintptr(unsafe.Pointer(v.native()))
}

func (f *PixbufFormat) GetName() string {
	c := C.gdk_pixbuf_format_get_name(f.native())
	defer C.g_free(C.gpointer(c))
	return goString(c)
}

func (f *PixbufFormat) GetDescription() string {
	c := C.gdk_pixbuf_format_get_description(f.native())
	defer C.g_free(C.gpointer(c))
	return goString(c)
}

func (f *PixbufFormat) GetLicense() string {
	c := C.gdk_pixbuf_format_get_license(f.native())
	defer C.g_free(C.gpointer(c))
	return goString(c)
}

func PixbufGetFormats() []*PixbufFormat {
	l := (*C.GSList)(C.gdk_pixbuf_get_formats())
	formats := glib.WrapSList(uintptr(unsafe.Pointer(l)))
	if formats == nil {
		return nil // no error. A nil list is considered to be empty.
	}

	// "The structures themselves are owned by GdkPixbuf". Free the list only.
	defer formats.Free()

	ret := make([]*PixbufFormat, 0, formats.Length())
	formats.Foreach(func(ptr unsafe.Pointer) {
		ret = append(ret, &PixbufFormat{(*C.GdkPixbufFormat)(ptr)})
	})

	return ret
}
