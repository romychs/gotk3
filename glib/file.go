package glib

// #cgo pkg-config: glib-2.0 gobject-2.0
// #include <gio/gio.h>
// #include <glib.h>
// #include <glib-object.h>
// #include "glib.go.h"
import "C"
import (
	"errors"
	"runtime"
	"unsafe"
)

// Cancellable is a representation of GCancellable.
type Cancellable struct {
	*Object
}

// native() returns a pointer to the underlying GCancellable.
func (v *Cancellable) native() *C.GCancellable {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toGCancellable(ptr)
}

func marshalCancellable(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := Take(unsafe.Pointer(c))
	return wrapCancellable(obj), nil
}

func wrapCancellable(obj *Object) *Cancellable {
	return &Cancellable{obj}
}

// CancellableNew is a wrapper around g_cancellable_new().
func CancellableNew() (*Cancellable, error) {
	c := C.g_cancellable_new()

	if c == nil {
		return nil, errNilPtr
	}
	obj := Take(unsafe.Pointer(c))
	return wrapCancellable(obj), nil
}

// gboolean	g_cancellable_is_cancelled ()
func (v *Cancellable) IsCancelled() bool {
	c := C.g_cancellable_is_cancelled(v.native())
	return gobool(c)
}

// void	g_cancellable_reset ()
func (v *Cancellable) Reset() {
	C.g_cancellable_reset(v.native())
}

// void	g_cancellable_cancel ()
func (v *Cancellable) Cancel() {
	C.g_cancellable_cancel(v.native())
}

// FileInfo is a representation of GFileInfo.
type FileInfo struct {
	*Object
}

// native() returns a pointer to the underlying GInputStream.
func (v *FileInfo) native() *C.GFileInfo {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toGFileInfo(ptr)
}

func marshalFileInfo(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := Take(unsafe.Pointer(c))
	return wrapFileInfo(obj), nil
}

func wrapFileInfo(obj *Object) *FileInfo {
	return &FileInfo{obj}
}

// InputStream is a representation of GInputStream.
type InputStream struct {
	*Object
}

// native() returns a pointer to the underlying GInputStream.
func (v *InputStream) native() *C.GInputStream {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toGInputStream(ptr)
}

func marshalInputStream(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := Take(unsafe.Pointer(c))
	return wrapInputStream(obj), nil
}

func wrapInputStream(obj *Object) *InputStream {
	return &InputStream{obj}
}

// gssize	g_input_stream_read ()
func (v *InputStream) Read(b []byte, cancel *Cancellable) (bytesRead int, e error) {
	var err *C.GError
	c := C.g_input_stream_read(v.native(), unsafe.Pointer(&b[0]),
		C.gsize(len(b)), cancel.native(), &err)
	if err != nil {
		defer C.g_error_free(err)
		return int(c), errors.New(goString(err.message))
	}
	return int(c), nil
}

// gboolean	g_input_stream_read_all ()
func (v *InputStream) ReadAll(b []byte, cancel *Cancellable) (bytesRead int, e error) {
	var err *C.GError
	var br C.gsize
	c := C.g_input_stream_read_all(v.native(), unsafe.Pointer(&b[0]),
		C.gsize(len(b)), &br, cancel.native(), &err)
	if c == 0 {
		defer C.g_error_free(err)
		return int(br), errors.New(goString(err.message))
	}
	return int(br), nil
}

// gssize	g_input_stream_skip ()
func (v *InputStream) Skip(count int, cancel *Cancellable) (bytesSkipped int, e error) {
	var err *C.GError
	c := C.g_input_stream_skip(v.native(), C.gsize(count), cancel.native(), &err)
	if err != nil {
		defer C.g_error_free(err)
		return 0, errors.New(goString(err.message))
	}
	return int(c), nil
}

// gboolean	g_input_stream_close ()
func (v *InputStream) Close(cancel *Cancellable) error {
	var err *C.GError
	c := C.g_input_stream_close(v.native(), cancel.native(), &err)
	if c == 0 {
		defer C.g_error_free(err)
		return errors.New(goString(err.message))
	}
	return nil
}

// gboolean	g_input_stream_is_closed ()
func (v *InputStream) IsClosed() bool {
	c := C.g_input_stream_is_closed(v.native())
	return gobool(c)
}

// InputStream is a representation of GInputStream.
type FileInputStream struct {
	InputStream
}

// native() returns a pointer to the underlying GFileInputStream.
func (v *FileInputStream) native() *C.GFileInputStream {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toGFileInputStream(ptr)
}

func marshalFileInputStream(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := Take(unsafe.Pointer(c))
	return wrapFileInputStream(obj), nil
}

func wrapFileInputStream(obj *Object) *FileInputStream {
	return &FileInputStream{InputStream{obj}}
}

// OutputStreamSpliceFlags is a representation of GLib's GOutputStreamSpliceFlags.

type OutputStreamSpliceFlags int

const (
	OUTPUT_STREAM_SPLICE_NONE         OutputStreamSpliceFlags = C.G_OUTPUT_STREAM_SPLICE_NONE
	OUTPUT_STREAM_SPLICE_CLOSE_SOURCE OutputStreamSpliceFlags = C.G_OUTPUT_STREAM_SPLICE_CLOSE_SOURCE
	OUTPUT_STREAM_SPLICE_CLOSE_TARGET OutputStreamSpliceFlags = C.G_OUTPUT_STREAM_SPLICE_CLOSE_TARGET
)

// OutputStream is a representation of GOutputStream.
type OutputStream struct {
	*Object
}

// native() returns a pointer to the underlying GOutputStream.
func (v *OutputStream) native() *C.GOutputStream {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toGOutputStream(ptr)
}

func marshalOutputStream(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := Take(unsafe.Pointer(c))
	return wrapOutputStream(obj), nil
}

func wrapOutputStream(obj *Object) *OutputStream {
	return &OutputStream{obj}
}

// gssize	g_output_stream_write ()
func (v *OutputStream) Write(b []byte, cancel *Cancellable) (bytesWritten int, e error) {
	var err *C.GError
	c := C.g_output_stream_write(v.native(), unsafe.Pointer(&b[0]),
		C.gsize(len(b)), cancel.native(), &err)
	if err != nil {
		defer C.g_error_free(err)
		return int(c), errors.New(goString(err.message))
	}
	return int(c), nil
}

// gssize	g_output_stream_write_all ()
func (v *OutputStream) WriteAll(b []byte, cancel *Cancellable) (bytesWritten int, e error) {
	var err *C.GError
	var bw C.gsize
	c := C.g_output_stream_write_all(v.native(), unsafe.Pointer(&b[0]),
		C.gsize(len(b)), &bw, cancel.native(), &err)
	if c == 0 {
		defer C.g_error_free(err)
		return int(bw), errors.New(goString(err.message))
	}
	return int(bw), nil
}

// gssize	g_output_stream_splice ()
func (v *OutputStream) Splice(source *InputStream, flags OutputStreamSpliceFlags,
	cancel *Cancellable) (dataSpliced int, e error) {
	var err *C.GError
	c := C.g_output_stream_splice(v.native(), source.native(),
		C.GOutputStreamSpliceFlags(flags), cancel.native(), &err)
	if err != nil {
		defer C.g_error_free(err)
		return int(c), errors.New(goString(err.message))
	}
	return int(c), nil
}

// gboolean	g_output_stream_flush ()
func (v *OutputStream) Flush(cancel *Cancellable) error {
	var err *C.GError
	c := C.g_output_stream_flush(v.native(), cancel.native(), &err)
	if c == 0 {
		defer C.g_error_free(err)
		return errors.New(goString(err.message))
	}
	return nil
}

// gboolean	g_output_stream_close ()
func (v *OutputStream) Close(cancel *Cancellable) error {
	var err *C.GError
	c := C.g_output_stream_close(v.native(), cancel.native(), &err)
	if c == 0 {
		defer C.g_error_free(err)
		return errors.New(goString(err.message))
	}
	return nil
}

// gboolean	g_output_stream_is_closing ()
func (v *OutputStream) IsClosing() bool {
	c := C.g_output_stream_is_closing(v.native())
	return gobool(c)
}

// gboolean	g_output_stream_is_closed ()
func (v *OutputStream) IsClosed() bool {
	c := C.g_output_stream_is_closed(v.native())
	return gobool(c)
}

// FileOutputStream is a representation of GFileOutputStream.
type FileOutputStream struct {
	OutputStream
}

// native() returns a pointer to the underlying GFileInputStream.
func (v *FileOutputStream) native() *C.GFileOutputStream {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toGFileOutputStream(ptr)
}

func marshalFileOutputStream(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := Take(unsafe.Pointer(c))
	return wrapFileOutputStream(obj), nil
}

func wrapFileOutputStream(obj *Object) *FileOutputStream {
	return &FileOutputStream{OutputStream{obj}}
}

// FileEnumerator is a representation of GFileEnumerator.
type FileEnumerator struct {
	*Object
}

// native() returns a pointer to the underlying GFileInputStream.
func (v *FileEnumerator) native() *C.GFileEnumerator {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toGFileEnumerator(ptr)
}

func marshalFileEnumerator(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := Take(unsafe.Pointer(c))
	return wrapFileEnumerator(obj), nil
}

func wrapFileEnumerator(obj *Object) *FileEnumerator {
	return &FileEnumerator{obj}
}

/*
// gboolean	g_file_enumerator_iterate ()
func (v *FileEnumerator) Iterate(cancel *Cancellable) (*FileInfo, *File, error) {
	var err *C.GError
	var info *C.GFileInfo
	var child *C.GFile
	c := C.g_file_enumerator_iterate(v.native(), &info, &child, cancel.native(), &err)
	if c == 0 {
		defer C.g_error_free(err)
		return nil, nil, errors.New(goString(err.message))
	}
	var info2 *FileInfo
	var child2 *File
	if info != nil {
		obj := Take(unsafe.Pointer(info))
		info2 = wrapFileInfo(obj)
	}
	if child != nil {
		intf := SetFinOnInterface(unsafe.Pointer(child))
		child2 = wrapFile(intf)
	}
	return info2, child2, nil
}
*/

// GFileInfo *	g_file_enumerator_next_file ()
func (v *FileEnumerator) NextFile(cancel *Cancellable) (*FileInfo, error) {
	var err *C.GError
	c := C.g_file_enumerator_next_file(v.native(), cancel.native(), &err)
	if err != nil {
		defer C.g_error_free(err)
		return nil, errors.New(goString(err.message))
	}
	obj := Take(unsafe.Pointer(c))
	return wrapFileInfo(obj), nil
}

// gboolean	g_file_enumerator_close ()
func (v *FileEnumerator) Close(cancel *Cancellable) error {
	var err *C.GError
	c := C.g_file_enumerator_close(v.native(), cancel.native(), &err)
	if c == 0 {
		defer C.g_error_free(err)
		return errors.New(goString(err.message))
	}
	return nil
}

// gboolean	g_file_enumerator_is_closed ()
func (v *FileEnumerator) IsClosed() bool {
	c := C.g_file_enumerator_is_closed(v.native())
	return gobool(c)
}

// gboolean	g_file_enumerator_has_pending ()
func (v *FileEnumerator) HasPending() bool {
	c := C.g_file_enumerator_has_pending(v.native())
	return gobool(c)
}

// void	g_file_enumerator_set_pending ()
func (v *FileEnumerator) SetPending(pending bool) {
	C.g_file_enumerator_set_pending(v.native(), gbool(pending))
}

// GFile *	g_file_enumerator_get_container ()
func (v *FileEnumerator) GetContainer() *File {
	c := C.g_file_enumerator_get_container(v.native())
	intf := SetFinOnInterface(unsafe.Pointer(c))
	return wrapFile(intf)
}

// GFile *	g_file_enumerator_get_child ()
func (v *FileEnumerator) GetChild(info *FileInfo) *File {
	c := C.g_file_enumerator_get_child(v.native(), info.native())
	intf := SetFinOnInterface(unsafe.Pointer(c))
	return wrapFile(intf)
}

// IOStream is a representation of GIOStream.
type IOStream struct {
	*Object
}

// native() returns a pointer to the underlying GFileInputStream.
func (v *IOStream) native() *C.GIOStream {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toGIOStream(ptr)
}

func marshalIOStream(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := Take(unsafe.Pointer(c))
	return wrapIOStream(obj), nil
}

func wrapIOStream(obj *Object) *IOStream {
	return &IOStream{obj}
}

// GInputStream *	g_io_stream_get_input_stream ()
func (v *IOStream) GetInputStream() (*InputStream, error) {
	c := C.g_io_stream_get_input_stream(v.native())
	if c == nil {
		return nil, errNilPtr
	}
	// a GInputStream, owned by the GIOStream. Do not free.
	obj := wrapObject(unsafe.Pointer(c))
	return wrapInputStream(obj), nil
}

// GOutputStream *	g_io_stream_get_output_stream ()
func (v *IOStream) GetOutputStream() (*OutputStream, error) {
	c := C.g_io_stream_get_output_stream(v.native())
	if c == nil {
		return nil, errNilPtr
	}
	// a GOutputStream, owned by the GIOStream. Do not free.
	obj := wrapObject(unsafe.Pointer(c))
	return wrapOutputStream(obj), nil
}

// gboolean	g_io_stream_close ()
func (v *IOStream) Close(cancel *Cancellable) error {
	var err *C.GError
	c := C.g_io_stream_close(v.native(), cancel.native(), &err)
	if c == 0 {
		defer C.g_error_free(err)
		return errors.New(goString(err.message))
	}
	return nil
}

// gboolean	g_io_stream_is_closed ()
func (v *IOStream) IsClosed() bool {
	c := C.g_io_stream_is_closed(v.native())
	return gobool(c)
}

// gboolean	g_io_stream_has_pending ()
func (v *IOStream) HasPending() bool {
	c := C.g_io_stream_has_pending(v.native())
	return gobool(c)
}

// gboolean	g_io_stream_set_pending ()
func (v *IOStream) SetPending() error {
	var err *C.GError
	c := C.g_io_stream_set_pending(v.native(), &err)
	if c == 0 {
		defer C.g_error_free(err)
		return errors.New(goString(err.message))
	}
	return nil
}

// void	g_io_stream_clear_pending ()
func (v *IOStream) ClearPending() {
	C.g_io_stream_clear_pending(v.native())
}

// SeekType is a representation of GLib's GSeekType.
type SeekType int

const (
	SEEK_CUR SeekType = C.G_SEEK_CUR
	SEEK_SET SeekType = C.G_SEEK_SET
	SEEK_END SeekType = C.G_SEEK_END
)

// Seekable is a representation of GSeekable.
type Seekable struct {
	Interface
}

// native() returns a pointer to the underlying GAction.
func (v *Seekable) native() *C.GSeekable {
	return C.toGSeekable(unsafe.Pointer(v.Native()))
}

func marshalSeekable(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := Take(unsafe.Pointer(c))
	act := wrapSeekable(*InterfaceFromObjectNew(obj))
	return act, nil
}

func wrapSeekable(intf Interface) *Seekable {
	return &Seekable{intf}
}

// goffset	g_seekable_tell ()
func (v *Seekable) Tell() int64 {
	c := C.g_seekable_tell(v.native())
	return int64(c)
}

// gboolean	g_seekable_can_seek ()
func (v *Seekable) CanSeek() bool {
	c := C.g_seekable_can_seek(v.native())
	return gobool(c)
}

// gboolean	g_seekable_seek ()
func (v *Seekable) Seek(offset int64, seekType SeekType, cancel *Cancellable) error {
	var err *C.GError
	c := C.g_seekable_seek(v.native(), C.goffset(offset),
		C.GSeekType(seekType), cancel.native(), &err)
	if c == 0 {
		defer C.g_error_free(err)
		return errors.New(goString(err.message))
	}
	return nil
}

// gboolean	g_seekable_can_truncate ()
func (v *Seekable) CanTruncate() bool {
	c := C.g_seekable_can_truncate(v.native())
	return gobool(c)
}

// gboolean	g_seekable_truncate ()
func (v *Seekable) Truncate(offset int64, cancel *Cancellable) error {
	var err *C.GError
	c := C.g_seekable_truncate(v.native(), C.goffset(offset),
		cancel.native(), &err)
	if c == 0 {
		defer C.g_error_free(err)
		return errors.New(goString(err.message))
	}
	return nil
}

// FileIOStream is a representation of GFileIOStream.
type FileIOStream struct {
	IOStream
	// Interfaces
	Seekable
}

// native() returns a pointer to the underlying GFileInputStream.
func (v *FileIOStream) native() *C.GFileIOStream {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toGFileIOStream(ptr)
}

func marshalFileIOStream(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := Take(unsafe.Pointer(c))
	return wrapFileIOStream(obj), nil
}

func wrapFileIOStream(obj *Object) *FileIOStream {
	seekable := wrapSeekable(*InterfaceFromObjectNew(obj))
	return &FileIOStream{IOStream{obj}, *seekable}
}

// FileCreateFlags is a representation of GLib's GFileCreateFlags.

type FileCreateFlags int

const (
	FILE_CREATE_NONE                FileCreateFlags = C.G_FILE_CREATE_NONE
	FILE_CREATE_PRIVATE             FileCreateFlags = C.G_FILE_CREATE_PRIVATE
	FILE_CREATE_REPLACE_DESTINATION FileCreateFlags = C.G_FILE_CREATE_REPLACE_DESTINATION
)

// FileQueryInfoFlags is a representation of GLib's GFileQueryInfoFlags.

type FileQueryInfoFlags int

const (
	FILE_QUERY_INFO_NONE              FileQueryInfoFlags = C.G_FILE_QUERY_INFO_NONE
	FILE_QUERY_INFO_NOFOLLOW_SYMLINKS FileQueryInfoFlags = C.G_FILE_QUERY_INFO_NOFOLLOW_SYMLINKS
)

// FileType is a representation of GLib's GFileType.

type FileType int

const (
	FILE_TYPE_UNKNOWN       FileType = C.G_FILE_TYPE_UNKNOWN
	FILE_TYPE_REGULAR       FileType = C.G_FILE_TYPE_REGULAR
	FILE_TYPE_DIRECTORY     FileType = C.G_FILE_TYPE_DIRECTORY
	FILE_TYPE_SYMBOLIC_LINK FileType = C.G_FILE_TYPE_SYMBOLIC_LINK
	FILE_TYPE_SPECIAL       FileType = C.G_FILE_TYPE_SPECIAL
	FILE_TYPE_SHORTCUT      FileType = C.G_FILE_TYPE_SHORTCUT
	FILE_TYPE_MOUNTABLE     FileType = C.G_FILE_TYPE_MOUNTABLE
)

// File is a representation of GFile.
type File struct {
	// Since GFile is based on GInterface, but require
	// freed approach same as GObject has, use reference
	// to Interface, instead of instance.
	// This must be a pointer so copies of the ref-sinked object
	// do not outlive the original object, causing an unref
	// finalizer to prematurely run.
	*Interface
}

// native() returns a pointer to the underlying GThemedIcon.
func (v *File) native() *C.GFile {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Native())
	return C.toGFile(ptr)
}

func (v *File) Unref() {
	// wrapper around g_object_unref().
	C.g_object_unref(v.Interface.ginterface)
}

func SetFinOnInterface(ptr unsafe.Pointer) *Interface {
	intf := InterfaceNew(ptr)

	runtime.SetFinalizer(intf, func(intf *Interface) {
		// wrapper around g_object_unref().
		C.g_object_unref(intf.ginterface)
	})
	return intf
}

func marshalFile(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	intf := SetFinOnInterface(unsafe.Pointer(c))
	return wrapFile(intf), nil
}

func wrapFile(intf *Interface) *File {
	return &File{intf}
}

// FileForPathNew is a wrapper around g_file_new_for_path().
func FileForPathNew(path string) (*File, error) {
	cstr := C.CString(path)
	defer C.free(unsafe.Pointer(cstr))

	c := C.g_file_new_for_path(cstr)
	if c == nil {
		return nil, errNilPtr
	}

	intf := SetFinOnInterface(unsafe.Pointer(c))
	return wrapFile(intf), nil
}

// FileForUriNew is a wrapper around g_file_new_for_uri().
func FileForUriNew(path string) (*File, error) {
	cstr := C.CString(path)
	defer C.free(unsafe.Pointer(cstr))

	c := C.g_file_new_for_uri(cstr)
	if c == nil {
		return nil, errNilPtr
	}

	intf := SetFinOnInterface(unsafe.Pointer(c))
	return wrapFile(intf), nil
}

// GFile *	g_file_new_tmp ()
func (v *File) NewTmp(template string) (*File, *FileIOStream, error) {
	var cstr *C.char
	if template != "" {
		cstr = C.CString(template)
		defer C.free(unsafe.Pointer(cstr))
	}

	var err *C.GError
	var iostream *C.GFileIOStream
	c := C.g_file_new_tmp(cstr, &iostream, &err)
	if err != nil {
		defer C.g_error_free(err)
		return nil, nil, errors.New(goString(err.message))
	}

	intf := SetFinOnInterface(unsafe.Pointer(c))
	obj := Take(unsafe.Pointer(c))
	return wrapFile(intf), wrapFileIOStream(obj), nil
}

// GFile *	g_file_parse_name ()
func ParseName(parseName string) (*File, error) {
	cstr := C.CString(parseName)
	defer C.free(unsafe.Pointer(cstr))

	c := C.g_file_parse_name(cstr)
	if c == nil {
		return nil, errNilPtr
	}

	intf := SetFinOnInterface(unsafe.Pointer(c))
	return wrapFile(intf), nil
}

// GFile *	g_file_dup ()
func (v *File) Dup() (*File, error) {
	c := C.g_file_dup(v.native())
	if c == nil {
		return nil, errNilPtr
	}

	intf := SetFinOnInterface(unsafe.Pointer(c))
	return wrapFile(intf), nil
}

// guint	g_file_hash ()
func (v *File) Hash() int {
	c := C.g_file_hash(C.gconstpointer(v.native()))
	return int(c)
}

// gboolean	g_file_equal ()
func (v *File) Equal(file *File) bool {
	c := C.g_file_equal(v.native(), file.native())
	return gobool(c)
}

// char *	g_file_get_basename ()
func (v *File) GetBasename() string {
	c := C.g_file_get_basename(v.native())
	defer C.g_free(C.gpointer(c))
	return goString((*C.gchar)(c))
}

// char *	g_file_get_path ()
func (v *File) GetPath() string {
	c := C.g_file_get_path(v.native())
	defer C.g_free(C.gpointer(c))
	return goString((*C.gchar)(c))
}

// char *	g_file_get_uri ()
func (v *File) GetUri() string {
	c := C.g_file_get_uri(v.native())
	defer C.g_free(C.gpointer(c))
	return goString((*C.gchar)(c))
}

// GFile *	g_file_get_parent ()
func (v *File) GetParent() (*File, error) {
	c := C.g_file_get_parent(v.native())
	if c == nil {
		return nil, errNilPtr
	}

	intf := SetFinOnInterface(unsafe.Pointer(c))
	return wrapFile(intf), nil
}

// gboolean	g_file_has_parent ()
func (v *File) HasParent(parent *File) bool {
	c := C.g_file_has_parent(v.native(), parent.native())
	return gobool(c)
}

// GFile *	g_file_get_child ()
func (v *File) GetChild(name string) (*File, error) {
	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))

	c := C.g_file_get_child(v.native(), cstr)
	if c == nil {
		return nil, errNilPtr
	}

	intf := SetFinOnInterface(unsafe.Pointer(c))
	return wrapFile(intf), nil
}

// GFile *	g_file_get_child_for_display_name ()
func (v *File) GetChildForDisplayName(displayName string) (*File, error) {
	cstr := C.CString(displayName)
	defer C.free(unsafe.Pointer(cstr))

	var err *C.GError
	c := C.g_file_get_child_for_display_name(v.native(), cstr, &err)
	if err != nil {
		defer C.g_error_free(err)
		return nil, errors.New(goString(err.message))
	}

	intf := SetFinOnInterface(unsafe.Pointer(c))
	return wrapFile(intf), nil
}

// gboolean	g_file_has_prefix ()
func (v *File) HasPrefix(prefix *File) bool {
	c := C.g_file_has_prefix(v.native(), prefix.native())
	return gobool(c)
}

// char *	g_file_get_relative_path ()
func (v *File) GetRelativePath(descendant *File) string {
	c := C.g_file_get_relative_path(v.native(), descendant.native())
	defer C.g_free(C.gpointer(c))
	return goString((*C.gchar)(c))
}

// gboolean	g_file_is_native ()
func (v *File) IsNative() bool {
	c := C.g_file_is_native(v.native())
	return gobool(c)
}

// gboolean	g_file_has_uri_scheme ()
func (v *File) HasUriScheme(uriScheme string) bool {
	cstr := C.CString(uriScheme)
	defer C.free(unsafe.Pointer(cstr))

	c := C.g_file_has_uri_scheme(v.native(), cstr)
	return gobool(c)
}

// char *	g_file_get_uri_scheme ()
func (v *File) GetUriScheme() string {
	c := C.g_file_get_uri_scheme(v.native())
	defer C.g_free(C.gpointer(c))
	return goString((*C.gchar)(c))
}

// GFileInputStream *	g_file_read ()
func (v *File) Read(cancel *Cancellable) (*FileInputStream, error) {
	var err *C.GError
	c := C.g_file_read(v.native(), cancel.native(), &err)
	if err != nil {
		defer C.g_error_free(err)
		return nil, errors.New(goString(err.message))
	}
	obj := Take(unsafe.Pointer(c))
	return wrapFileInputStream(obj), nil
}

// GFileOutputStream *	g_file_append_to ()
func (v *File) AppendTo(flags FileCreateFlags, cancel *Cancellable) (*FileOutputStream, error) {
	var err *C.GError
	c := C.g_file_append_to(v.native(), C.GFileCreateFlags(flags), cancel.native(), &err)
	if err != nil {
		defer C.g_error_free(err)
		return nil, errors.New(goString(err.message))
	}
	obj := Take(unsafe.Pointer(c))
	return wrapFileOutputStream(obj), nil
}

// GFileOutputStream *	g_file_create ()
func (v *File) Create(flags FileCreateFlags, cancel *Cancellable) (*FileOutputStream, error) {
	var err *C.GError
	c := C.g_file_create(v.native(), C.GFileCreateFlags(flags), cancel.native(), &err)
	if err != nil {
		defer C.g_error_free(err)
		return nil, errors.New(goString(err.message))
	}
	obj := Take(unsafe.Pointer(c))
	return wrapFileOutputStream(obj), nil
}

// GFileInfo *	g_file_query_info ()
func (v *File) QueryInfo(attributes string, flags FileQueryInfoFlags, cancel *Cancellable) (*FileInfo, error) {
	cstr := C.CString(attributes)
	defer C.free(unsafe.Pointer(cstr))

	var err *C.GError
	c := C.g_file_query_info(v.native(), cstr, C.GFileQueryInfoFlags(flags), cancel.native(), &err)
	if err != nil {
		defer C.g_error_free(err)
		return nil, errors.New(goString(err.message))
	}
	obj := Take(unsafe.Pointer(c))
	return wrapFileInfo(obj), nil
}

// gboolean	g_file_query_exists ()
func (v *File) QueryExists(cancel *Cancellable) bool {
	c := C.g_file_query_exists(v.native(), cancel.native())
	return gobool(c)
}

// GFileType	g_file_query_file_type ()
func (v *File) FileType(flags FileQueryInfoFlags, cancel *Cancellable) FileType {
	c := C.g_file_query_file_type(v.native(), C.GFileQueryInfoFlags(flags), cancel.native())
	return FileType(c)
}

// GFileInfo *	g_file_query_filesystem_info ()
func (v *File) QueryFileSystemInfo(attributes string, cancel *Cancellable) (*FileInfo, error) {
	cstr := C.CString(attributes)
	defer C.free(unsafe.Pointer(cstr))

	var err *C.GError
	c := C.g_file_query_filesystem_info(v.native(), cstr, cancel.native(), &err)
	if err != nil {
		defer C.g_error_free(err)
		return nil, errors.New(goString(err.message))
	}
	obj := Take(unsafe.Pointer(c))
	return wrapFileInfo(obj), nil
}

// GFileEnumerator *	g_file_enumerate_children ()
func (v *File) EnumerateChildren(attributes string, flags FileQueryInfoFlags, cancel *Cancellable) (*FileEnumerator, error) {
	cstr := C.CString(attributes)
	defer C.free(unsafe.Pointer(cstr))

	var err *C.GError
	c := C.g_file_enumerate_children(v.native(), cstr, C.GFileQueryInfoFlags(flags), cancel.native(), &err)
	if err != nil {
		defer C.g_error_free(err)
		return nil, errors.New(goString(err.message))
	}
	obj := Take(unsafe.Pointer(c))
	return wrapFileEnumerator(obj), nil
}

// GFile *	g_file_set_display_name ()
func (v *File) SetDisplayName(displayName string, cancel *Cancellable) (*File, error) {
	cstr := C.CString(displayName)
	defer C.free(unsafe.Pointer(cstr))

	var err *C.GError
	c := C.g_file_set_display_name(v.native(), cstr, cancel.native(), &err)
	if err != nil {
		defer C.g_error_free(err)
		return nil, errors.New(goString(err.message))
	}
	intf := SetFinOnInterface(unsafe.Pointer(c))
	return wrapFile(intf), nil
}

// gboolean	g_file_delete ()
func (v *File) Delete(cancel *Cancellable) error {
	var err *C.GError
	c := C.g_file_delete(v.native(), cancel.native(), &err)
	if c == 0 {
		defer C.g_error_free(err)
		return errors.New(goString(err.message))
	}
	return nil
}

// gboolean	g_file_trash ()
func (v *File) Trash(cancel *Cancellable) error {
	var err *C.GError
	c := C.g_file_trash(v.native(), cancel.native(), &err)
	if c == 0 {
		defer C.g_error_free(err)
		return errors.New(goString(err.message))
	}
	return nil
}

// gboolean	g_file_make_directory ()
func (v *File) MakeDirectory(cancel *Cancellable) error {
	var err *C.GError
	c := C.g_file_make_directory(v.native(), cancel.native(), &err)
	if c == 0 {
		defer C.g_error_free(err)
		return errors.New(goString(err.message))
	}
	return nil
}

// gboolean	g_file_make_directory_with_parents ()
func (v *File) MakeDirectoryWithParents(cancel *Cancellable) error {
	var err *C.GError
	c := C.g_file_make_directory_with_parents(v.native(), cancel.native(), &err)
	if c == 0 {
		defer C.g_error_free(err)
		return errors.New(goString(err.message))
	}
	return nil
}

// gboolean	g_file_make_symbolic_link ()
func (v *File) MakeSymbolicLink(symlinkValue string, cancel *Cancellable) error {
	cstr := C.CString(symlinkValue)
	defer C.free(unsafe.Pointer(cstr))

	var err *C.GError
	c := C.g_file_make_symbolic_link(v.native(), cstr, cancel.native(), &err)
	if c == 0 {
		defer C.g_error_free(err)
		return errors.New(goString(err.message))
	}
	return nil
}

// GFileIOStream *	g_file_open_readwrite ()
func (v *File) OpenReadWrite(cancel *Cancellable) (*FileIOStream, error) {
	var err *C.GError
	c := C.g_file_open_readwrite(v.native(), cancel.native(), &err)
	if err != nil {
		defer C.g_error_free(err)
		return nil, errors.New(goString(err.message))
	}
	obj := Take(unsafe.Pointer(c))
	return wrapFileIOStream(obj), nil
}

// GFile *	g_file_new_for_commandline_arg ()
// GFile *	g_file_new_for_commandline_arg_and_cwd ()

// char *	g_file_get_parse_name ()

// GFile *	g_file_resolve_relative_path ()
// void	g_file_read_async ()
// GFileInputStream *	g_file_read_finish ()
// GFileOutputStream *	g_file_replace ()
// void	g_file_append_to_async ()
// GFileOutputStream *	g_file_append_to_finish ()
// void	g_file_create_async ()
// GFileOutputStream *	g_file_create_finish ()
// void	g_file_replace_async ()
// GFileOutputStream *	g_file_replace_finish ()
// void	g_file_query_info_async ()
// GFileInfo *	g_file_query_info_finish ()
// void	g_file_query_filesystem_info_async ()
// GFileInfo *	g_file_query_filesystem_info_finish ()
// GAppInfo *	g_file_query_default_handler ()
// gboolean	g_file_measure_disk_usage ()
// void	g_file_measure_disk_usage_async ()
// gboolean	g_file_measure_disk_usage_finish ()
// GMount *	g_file_find_enclosing_mount ()
// void	g_file_find_enclosing_mount_async ()
// GMount *	g_file_find_enclosing_mount_finish ()

// void	g_file_enumerate_children_async ()
// GFileEnumerator *	g_file_enumerate_children_finish ()
// void	g_file_set_display_name_async ()
// GFile *	g_file_set_display_name_finish ()
// void	g_file_delete_async ()
// gboolean	g_file_delete_finish ()
// void	g_file_trash_async ()
// gboolean	g_file_trash_finish ()
// gboolean	g_file_copy ()
// void	g_file_copy_async ()
// gboolean	g_file_copy_finish ()
// gboolean	g_file_move ()
// void	g_file_make_directory_async ()
// gboolean	g_file_make_directory_finish ()
// GFileAttributeInfoList *	g_file_query_settable_attributes ()
// GFileAttributeInfoList *	g_file_query_writable_namespaces ()
// gboolean	g_file_set_attribute ()
// gboolean	g_file_set_attributes_from_info ()
// void	g_file_set_attributes_async ()
// gboolean	g_file_set_attributes_finish ()
// gboolean	g_file_set_attribute_string ()
// gboolean	g_file_set_attribute_byte_string ()
// gboolean	g_file_set_attribute_uint32 ()
// gboolean	g_file_set_attribute_int32 ()
// gboolean	g_file_set_attribute_uint64 ()
// gboolean	g_file_set_attribute_int64 ()
// void	g_file_mount_mountable ()
// GFile *	g_file_mount_mountable_finish ()
// void	g_file_unmount_mountable ()
// gboolean	g_file_unmount_mountable_finish ()
// void	g_file_unmount_mountable_with_operation ()
// gboolean	g_file_unmount_mountable_with_operation_finish ()
// void	g_file_eject_mountable ()
// gboolean	g_file_eject_mountable_finish ()
// void	g_file_eject_mountable_with_operation ()
// gboolean	g_file_eject_mountable_with_operation_finish ()
// void	g_file_start_mountable ()
// gboolean	g_file_start_mountable_finish ()
// void	g_file_stop_mountable ()
// gboolean	g_file_stop_mountable_finish ()
// void	g_file_poll_mountable ()
// gboolean	g_file_poll_mountable_finish ()
// void	g_file_mount_enclosing_volume ()
// gboolean	g_file_mount_enclosing_volume_finish ()
// GFileMonitor *	g_file_monitor_directory ()
// GFileMonitor *	g_file_monitor_file ()
// GFileMonitor *	g_file_monitor ()
// gboolean	g_file_load_contents ()
// void	g_file_load_contents_async ()
// gboolean	g_file_load_contents_finish ()
// void	g_file_load_partial_contents_async ()
// gboolean	g_file_load_partial_contents_finish ()
// gboolean	g_file_replace_contents ()
// void	g_file_replace_contents_async ()
// void	g_file_replace_contents_bytes_async ()
// gboolean	g_file_replace_contents_finish ()
// gboolean	g_file_copy_attributes ()
// GFileIOStream *	g_file_create_readwrite ()
// void	g_file_create_readwrite_async ()
// GFileIOStream *	g_file_create_readwrite_finish ()
// void	g_file_open_readwrite_async ()
// GFileIOStream *	g_file_open_readwrite_finish ()
// GFileIOStream *	g_file_replace_readwrite ()
// void	g_file_replace_readwrite_async ()
// GFileIOStream *	g_file_replace_readwrite_finish ()
// gboolean	g_file_supports_thread_contexts ()
