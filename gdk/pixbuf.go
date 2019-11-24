package gdk

// #cgo pkg-config: gdk-3.0
// #include <gdk/gdk.h>
// #include "gdk.go.h"
import "C"

import (
	"errors"
	"reflect"
	"runtime"
	"strconv"
	"unsafe"

	"github.com/d2r2/gotk3/glib"
)

// PixbufAlphaMode is a representation of GDK's GdkPixbufAlphaMode.
type PixbufAlphaMode int

const (
	GDK_PIXBUF_ALPHA_BILEVEL PixbufAlphaMode = C.GDK_PIXBUF_ALPHA_BILEVEL
	GDK_PIXBUF_ALPHA_FULL    PixbufAlphaMode = C.GDK_PIXBUF_ALPHA_FULL
)

func marshalPixbufAlphaMode(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return PixbufAlphaMode(c), nil
}

/*
 * GdkPixbufFormat
 */

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

/*
 * GdkPixbuf
 */

// Pixbuf is a representation of GDK's GdkPixbuf.
type Pixbuf struct {
	*glib.Object
}

// native returns a pointer to the underlying GdkPixbuf.
func (v *Pixbuf) native() *C.GdkPixbuf {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toGdkPixbuf(ptr)
}

func marshalPixbuf(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.ToObject(unsafe.Pointer(c))
	return &Pixbuf{obj}, nil
}

// GetColorspace is a wrapper around gdk_pixbuf_get_colorspace().
func (v *Pixbuf) GetColorspace() Colorspace {
	c := C.gdk_pixbuf_get_colorspace(v.native())
	return Colorspace(c)
}

// GetNChannels is a wrapper around gdk_pixbuf_get_n_channels().
func (v *Pixbuf) GetNChannels() int {
	c := C.gdk_pixbuf_get_n_channels(v.native())
	return int(c)
}

// GetHasAlpha is a wrapper around gdk_pixbuf_get_has_alpha().
func (v *Pixbuf) GetHasAlpha() bool {
	c := C.gdk_pixbuf_get_has_alpha(v.native())
	return gobool(c)
}

// GetBitsPerSample is a wrapper around gdk_pixbuf_get_bits_per_sample().
func (v *Pixbuf) GetBitsPerSample() int {
	c := C.gdk_pixbuf_get_bits_per_sample(v.native())
	return int(c)
}

// GetPixels is a wrapper around gdk_pixbuf_get_pixels_with_length().
// A Go slice is used to represent the underlying Pixbuf data array, one
// byte per channel.
func (v *Pixbuf) GetPixels() (channels []byte) {
	var length C.guint
	c := C.gdk_pixbuf_get_pixels_with_length(v.native(), &length)
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&channels))
	sliceHeader.Data = uintptr(unsafe.Pointer(c))
	sliceHeader.Len = int(length)
	sliceHeader.Cap = int(length)

	// To make sure the slice doesn't outlive the Pixbuf, add a reference
	v.Ref()
	runtime.SetFinalizer(&channels, func(_ *[]byte) {
		v.Unref()
	})
	return
}

// GetWidth is a wrapper around gdk_pixbuf_get_width().
func (v *Pixbuf) GetWidth() int {
	c := C.gdk_pixbuf_get_width(v.native())
	return int(c)
}

// GetHeight is a wrapper around gdk_pixbuf_get_height().
func (v *Pixbuf) GetHeight() int {
	c := C.gdk_pixbuf_get_height(v.native())
	return int(c)
}

// GetRowstride is a wrapper around gdk_pixbuf_get_rowstride().
func (v *Pixbuf) GetRowstride() int {
	c := C.gdk_pixbuf_get_rowstride(v.native())
	return int(c)
}

// GetByteLength is a wrapper around gdk_pixbuf_get_byte_length().
func (v *Pixbuf) GetByteLength() int {
	c := C.gdk_pixbuf_get_byte_length(v.native())
	return int(c)
}

// GetOption is a wrapper around gdk_pixbuf_get_option().  ok is true if
// the key has an associated value.
func (v *Pixbuf) GetOption(key string) (value string, ok bool) {
	cstr := C.CString(key)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gdk_pixbuf_get_option(v.native(), (*C.gchar)(cstr))
	if c == nil {
		return "", false
	}
	return goString(c), true
}

// PixbufNew is a wrapper around gdk_pixbuf_new().
func PixbufNew(colorspace Colorspace, hasAlpha bool, bitsPerSample, width, height int) (*Pixbuf, error) {
	c := C.gdk_pixbuf_new(C.GdkColorspace(colorspace), gbool(hasAlpha),
		C.int(bitsPerSample), C.int(width), C.int(height))
	if c == nil {
		return nil, nilPtrErr
	}

	return &Pixbuf{glib.Take(unsafe.Pointer(c))}, nil
}

// PixbufCopy is a wrapper around gdk_pixbuf_copy().
func PixbufCopy(v *Pixbuf) (*Pixbuf, error) {
	c := C.gdk_pixbuf_copy(v.native())
	if c == nil {
		return nil, nilPtrErr
	}

	return &Pixbuf{glib.Take(unsafe.Pointer(c))}, nil
}

// PixbufNewFromFile is a wrapper around gdk_pixbuf_new_from_file().
func PixbufNewFromFile(filename string) (*Pixbuf, error) {
	cstr := C.CString(filename)
	defer C.free(unsafe.Pointer(cstr))

	var err *C.GError
	res := C.gdk_pixbuf_new_from_file(cstr, &err)
	if res == nil {
		defer C.g_error_free(err)
		return nil, errors.New(goString(err.message))
	}

	return &Pixbuf{glib.Take(unsafe.Pointer(res))}, nil
}

// PixbufNewFromStream is a wrapper around gdk_pixbuf_new_from_stream().
func PixbufNewFromStream(stream *glib.InputStream, cancellable *glib.Cancellable) (*Pixbuf, error) {
	var err *C.GError
	var cancell *C.GCancellable
	if cancellable != nil {
		cancell = C.toGCancellable(unsafe.Pointer(cancellable.Native()))
	}
	res := C.gdk_pixbuf_new_from_stream(C.toGInputStream(unsafe.Pointer(stream.Native())),
		cancell, &err)
	if res == nil {
		defer C.g_error_free(err)
		return nil, errors.New(goString(err.message))
	}

	return &Pixbuf{glib.Take(unsafe.Pointer(res))}, nil
}

// PixbufNewFromFileAtSize is a wrapper around gdk_pixbuf_new_from_file_at_size().
func PixbufNewFromFileAtSize(filename string, width, height int) (*Pixbuf, error) {
	cstr := C.CString(filename)
	defer C.free(unsafe.Pointer(cstr))

	var err *C.GError
	res := C.gdk_pixbuf_new_from_file_at_size(cstr, C.int(width), C.int(height), &err)
	if err != nil {
		defer C.g_error_free(err)
		return nil, errors.New(goString(err.message))
	}

	if res == nil {
		return nil, nilPtrErr
	}

	return &Pixbuf{glib.Take(unsafe.Pointer(res))}, nil
}

// PixbufNewFromFileAtScale is a wrapper around gdk_pixbuf_new_from_file_at_scale().
func PixbufNewFromFileAtScale(filename string, width, height int, preserveAspectRatio bool) (*Pixbuf, error) {
	cstr := C.CString(filename)
	defer C.free(unsafe.Pointer(cstr))

	var err *C.GError
	res := C.gdk_pixbuf_new_from_file_at_scale(cstr, C.int(width), C.int(height),
		gbool(preserveAspectRatio), &err)
	if err != nil {
		defer C.g_error_free(err)
		return nil, errors.New(goString(err.message))
	}

	if res == nil {
		return nil, nilPtrErr
	}

	return &Pixbuf{glib.Take(unsafe.Pointer(res))}, nil
}

// ScaleSimple is a wrapper around gdk_pixbuf_scale_simple().
func (v *Pixbuf) ScaleSimple(destWidth, destHeight int, interpType InterpType) (*Pixbuf, error) {
	c := C.gdk_pixbuf_scale_simple(v.native(), C.int(destWidth),
		C.int(destHeight), C.GdkInterpType(interpType))
	if c == nil {
		return nil, nilPtrErr
	}

	return &Pixbuf{glib.Take(unsafe.Pointer(c))}, nil
}

// RotateSimple is a wrapper around gdk_pixbuf_rotate_simple().
func (v *Pixbuf) RotateSimple(angle PixbufRotation) (*Pixbuf, error) {
	c := C.gdk_pixbuf_rotate_simple(v.native(), C.GdkPixbufRotation(angle))
	if c == nil {
		return nil, nilPtrErr
	}

	return &Pixbuf{glib.Take(unsafe.Pointer(c))}, nil
}

// ApplyEmbeddedOrientation is a wrapper around gdk_pixbuf_apply_embedded_orientation().
func (v *Pixbuf) ApplyEmbeddedOrientation() (*Pixbuf, error) {
	c := C.gdk_pixbuf_apply_embedded_orientation(v.native())
	if c == nil {
		return nil, nilPtrErr
	}

	return &Pixbuf{glib.Take(unsafe.Pointer(c))}, nil
}

// Flip is a wrapper around gdk_pixbuf_flip().
func (v *Pixbuf) Flip(horizontal bool) (*Pixbuf, error) {
	c := C.gdk_pixbuf_flip(v.native(), gbool(horizontal))
	if c == nil {
		return nil, nilPtrErr
	}

	return &Pixbuf{glib.Take(unsafe.Pointer(c))}, nil
}

// SaveJPEG is a wrapper around gdk_pixbuf_save().
// Quality is a number between 0...100
func (v *Pixbuf) SaveJPEG(path string, quality int) error {
	cpath := C.CString(path)
	cquality := C.CString(strconv.Itoa(quality))
	defer C.free(unsafe.Pointer(cpath))
	defer C.free(unsafe.Pointer(cquality))

	var err *C.GError
	c := C._gdk_pixbuf_save_jpeg(v.native(), cpath, &err, cquality)
	if !gobool(c) {
		defer C.g_error_free(err)
		return errors.New(goString(err.message))
	}

	return nil
}

// SavePNG is a wrapper around gdk_pixbuf_save().
// Compression is a number between 0...9
func (v *Pixbuf) SavePNG(path string, compression int) error {
	cpath := C.CString(path)
	ccompression := C.CString(strconv.Itoa(compression))
	defer C.free(unsafe.Pointer(cpath))
	defer C.free(unsafe.Pointer(ccompression))

	var err *C.GError
	c := C._gdk_pixbuf_save_png(v.native(), cpath, &err, ccompression)
	if !gobool(c) {
		defer C.g_error_free(err)
		return errors.New(goString(err.message))
	}
	return nil
}

// PixbufGetFileInfo is a wrapper around gdk_pixbuf_get_file_info().
func PixbufGetFileInfo(filename string) (*PixbufFormat, int, int, error) {
	cstr := C.CString(filename)
	defer C.free(unsafe.Pointer(cstr))
	var cw, ch C.gint
	c := C.gdk_pixbuf_get_file_info((*C.gchar)(cstr), &cw, &ch)
	if c == nil {
		return nil, 0, 0, nilPtrErr
	}

	return &PixbufFormat{c}, int(cw), int(ch), nil
}

/*
 * GdkPixbufAnimation
 */

// PixbufAnimation is a representation of GDK's GdkPixbufAnimation.
type PixbufAnimation struct {
	*glib.Object
}

// native returns a pointer to the underlying GdkPixbuf.
func (v *PixbufAnimation) native() *C.GdkPixbufAnimation {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toGdkPixbufAnimation(ptr)
}

func marshalPixbufAnimation(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.ToObject(unsafe.Pointer(c))
	return &Pixbuf{obj}, nil
}

// PixbufAnimationNewFromStream is a wrapper around gdk_pixbuf_animation_new_from_stream().
func PixbufAnimationNewFromStream(stream *glib.InputStream, cancellable *glib.Cancellable) (*PixbufAnimation, error) {
	var err *C.GError
	var cancell *C.GCancellable
	if cancellable != nil {
		cancell = C.toGCancellable(unsafe.Pointer(cancellable.Native()))
	}
	res := C.gdk_pixbuf_animation_new_from_stream(C.toGInputStream(unsafe.Pointer(stream.Native())),
		cancell, &err)
	if res == nil {
		defer C.g_error_free(err)
		return nil, errors.New(goString(err.message))
	}

	return &PixbufAnimation{glib.Take(unsafe.Pointer(res))}, nil
}

/*
 * GdkPixbufLoader
 */

// PixbufLoader is a representation of GDK's GdkPixbufLoader.
// Users of PixbufLoader are expected to call Close() when they are finished.
type PixbufLoader struct {
	*glib.Object
}

// native() returns a pointer to the underlying GdkPixbufLoader.
func (v *PixbufLoader) native() *C.GdkPixbufLoader {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toGdkPixbufLoader(ptr)
}

// PixbufLoaderNew is a wrapper around gdk_pixbuf_loader_new().
func PixbufLoaderNew() (*PixbufLoader, error) {
	c := C.gdk_pixbuf_loader_new()
	if c == nil {
		return nil, nilPtrErr
	}

	p := &PixbufLoader{glib.Take(unsafe.Pointer(c))}
	return p, nil
}

// PixbufLoaderNewWithType is a wrapper around gdk_pixbuf_loader_new_with_type().
func PixbufLoaderNewWithType(t string) (*PixbufLoader, error) {
	var err *C.GError

	cstr := C.CString(t)
	defer C.free(unsafe.Pointer(cstr))

	c := C.gdk_pixbuf_loader_new_with_type(cstr, &err)
	if err != nil {
		defer C.g_error_free(err)
		return nil, errors.New(goString(err.message))
	}

	if c == nil {
		return nil, nilPtrErr
	}

	return &PixbufLoader{glib.Take(unsafe.Pointer(c))}, nil
}

// Write is a wrapper around gdk_pixbuf_loader_write().  The
// function signature differs from the C equivalent to satisfy the
// io.Writer interface.
func (v *PixbufLoader) Write(data []byte) (int, error) {
	// n is set to 0 on error, and set to len(data) otherwise.
	// This is a tiny hacky to satisfy io.Writer and io.WriteCloser,
	// which would allow access to all io and ioutil goodies,
	// and play along nice with go environment.

	if len(data) == 0 {
		return 0, nil
	}

	var err *C.GError
	c := C.gdk_pixbuf_loader_write(v.native(),
		(*C.guchar)(unsafe.Pointer(&data[0])), C.gsize(len(data)),
		&err)

	if !gobool(c) {
		defer C.g_error_free(err)
		return 0, errors.New(goString(err.message))
	}

	return len(data), nil
}

// WriteBytes is a wrapper around gdk_pixbuf_loader_write_bytes().
func (v *PixbufLoader) WriteBytes(bytes *glib.Bytes) error {
	var err *C.GError
	c := C.gdk_pixbuf_loader_write_bytes(v.native(),
		(*C.GBytes)(unsafe.Pointer(bytes.Native())),
		&err)

	if !gobool(c) {
		defer C.g_error_free(err)
		return errors.New(goString(err.message))
	}

	return nil
}

// Close is a wrapper around gdk_pixbuf_loader_close().  An error is
// returned instead of a bool like the native C function to support the
// io.Closer interface.
func (v *PixbufLoader) Close() error {
	var err *C.GError

	if ok := gobool(C.gdk_pixbuf_loader_close(v.native(), &err)); !ok {
		defer C.g_error_free(err)
		return errors.New(goString(err.message))
	}
	return nil
}

// SetSize is a wrapper around gdk_pixbuf_loader_set_size().
func (v *PixbufLoader) SetSize(width, height int) {
	C.gdk_pixbuf_loader_set_size(v.native(), C.int(width), C.int(height))
}

// GetPixbuf is a wrapper around gdk_pixbuf_loader_get_pixbuf().
func (v *PixbufLoader) GetPixbuf() (*Pixbuf, error) {
	c := C.gdk_pixbuf_loader_get_pixbuf(v.native())
	if c == nil {
		return nil, nilPtrErr
	}

	return &Pixbuf{glib.Take(unsafe.Pointer(c))}, nil
}

// GetPixbufAnimation is a wrapper around gdk_pixbuf_loader_get_animation().
func (v *PixbufLoader) GetPixbufAnimation() (*PixbufAnimation, error) {
	c := C.gdk_pixbuf_loader_get_animation(v.native())
	if c == nil {
		return nil, nilPtrErr
	}

	return &PixbufAnimation{glib.Take(unsafe.Pointer(c))}, nil
}
