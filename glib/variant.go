//GVariant : GVariant — strongly typed value datatype
// https://developer.gnome.org/glib/2.26/glib-GVariant.html

package glib

// #cgo pkg-config: glib-2.0 gobject-2.0
// #include "variant.go.h"
// #include "glib.go.h"
import "C"

import (
	"fmt"
	"runtime"
	"unsafe"
)

/*
 * GVariantClass
 */

type VariantClass int

const (
	VARIANT_CLASS_BOOLEAN     VariantClass = C.G_VARIANT_CLASS_BOOLEAN     //The GVariant is a boolean.
	VARIANT_CLASS_BYTE        VariantClass = C.G_VARIANT_CLASS_BYTE        //The GVariant is a byte.
	VARIANT_CLASS_INT16       VariantClass = C.G_VARIANT_CLASS_INT16       //The GVariant is a signed 16 bit integer.
	VARIANT_CLASS_UINT16      VariantClass = C.G_VARIANT_CLASS_UINT16      //The GVariant is an unsigned 16 bit integer.
	VARIANT_CLASS_INT32       VariantClass = C.G_VARIANT_CLASS_INT32       //The GVariant is a signed 32 bit integer.
	VARIANT_CLASS_UINT32      VariantClass = C.G_VARIANT_CLASS_UINT32      //The GVariant is an unsigned 32 bit integer.
	VARIANT_CLASS_INT64       VariantClass = C.G_VARIANT_CLASS_INT64       //The GVariant is a signed 64 bit integer.
	VARIANT_CLASS_UINT64      VariantClass = C.G_VARIANT_CLASS_UINT64      //The GVariant is an unsigned 64 bit integer.
	VARIANT_CLASS_HANDLE      VariantClass = C.G_VARIANT_CLASS_HANDLE      //The GVariant is a file handle index.
	VARIANT_CLASS_DOUBLE      VariantClass = C.G_VARIANT_CLASS_DOUBLE      //The GVariant is a double precision floating point value.
	VARIANT_CLASS_STRING      VariantClass = C.G_VARIANT_CLASS_STRING      //The GVariant is a normal string.
	VARIANT_CLASS_OBJECT_PATH VariantClass = C.G_VARIANT_CLASS_OBJECT_PATH //The GVariant is a D-Bus object path string.
	VARIANT_CLASS_SIGNATURE   VariantClass = C.G_VARIANT_CLASS_SIGNATURE   //The GVariant is a D-Bus signature string.
	VARIANT_CLASS_VARIANT     VariantClass = C.G_VARIANT_CLASS_VARIANT     //The GVariant is a variant.
	VARIANT_CLASS_MAYBE       VariantClass = C.G_VARIANT_CLASS_MAYBE       //The GVariant is a maybe-typed value.
	VARIANT_CLASS_ARRAY       VariantClass = C.G_VARIANT_CLASS_ARRAY       //The GVariant is an array.
	VARIANT_CLASS_TUPLE       VariantClass = C.G_VARIANT_CLASS_TUPLE       //The GVariant is a tuple.
	VARIANT_CLASS_DICT_ENTRY  VariantClass = C.G_VARIANT_CLASS_DICT_ENTRY  //The GVariant is a dictionary entry.
)

/*
 * GVariantType
 */

// A VariantType is a wrapper for the GVariantType, which encodes type
// information for GVariants.
type VariantType struct {
	gvariantType *C.GVariantType
}

func (v *VariantType) native() *C.GVariantType {
	if v == nil {
		return nil
	}
	return v.gvariantType
}

// String returns a copy of this VariantType's type string.
func (v *VariantType) String() string {
	c := C.g_variant_type_dup_string(v.native())
	defer C.g_free(C.gpointer(c))

	return goString(c)
}

func newVariantType(v *C.GVariantType) *VariantType {
	return &VariantType{v}
}

// Variant types for comparing between them.  Cannot be const because
// they are pointers.
var (
	VARIANT_TYPE_BOOLEAN           = newVariantType(C._G_VARIANT_TYPE_BOOLEAN)
	VARIANT_TYPE_BYTE              = newVariantType(C._G_VARIANT_TYPE_BYTE)
	VARIANT_TYPE_INT16             = newVariantType(C._G_VARIANT_TYPE_INT16)
	VARIANT_TYPE_UINT16            = newVariantType(C._G_VARIANT_TYPE_UINT16)
	VARIANT_TYPE_INT32             = newVariantType(C._G_VARIANT_TYPE_INT32)
	VARIANT_TYPE_UINT32            = newVariantType(C._G_VARIANT_TYPE_UINT32)
	VARIANT_TYPE_INT64             = newVariantType(C._G_VARIANT_TYPE_INT64)
	VARIANT_TYPE_UINT64            = newVariantType(C._G_VARIANT_TYPE_UINT64)
	VARIANT_TYPE_HANDLE            = newVariantType(C._G_VARIANT_TYPE_HANDLE)
	VARIANT_TYPE_DOUBLE            = newVariantType(C._G_VARIANT_TYPE_DOUBLE)
	VARIANT_TYPE_STRING            = newVariantType(C._G_VARIANT_TYPE_STRING)
	VARIANT_TYPE_ANY               = newVariantType(C._G_VARIANT_TYPE_ANY)
	VARIANT_TYPE_BASIC             = newVariantType(C._G_VARIANT_TYPE_BASIC)
	VARIANT_TYPE_TUPLE             = newVariantType(C._G_VARIANT_TYPE_TUPLE)
	VARIANT_TYPE_UNIT              = newVariantType(C._G_VARIANT_TYPE_UNIT)
	VARIANT_TYPE_DICTIONARY        = newVariantType(C._G_VARIANT_TYPE_DICTIONARY)
	VARIANT_TYPE_STRING_ARRAY      = newVariantType(C._G_VARIANT_TYPE_STRING_ARRAY)
	VARIANT_TYPE_OBJECT_PATH_ARRAY = newVariantType(C._G_VARIANT_TYPE_OBJECT_PATH_ARRAY)
	VARIANT_TYPE_BYTESTRING        = newVariantType(C._G_VARIANT_TYPE_BYTESTRING)
	VARIANT_TYPE_BYTESTRING_ARRAY  = newVariantType(C._G_VARIANT_TYPE_BYTESTRING_ARRAY)
	VARIANT_TYPE_VARDICT           = newVariantType(C._G_VARIANT_TYPE_VARDICT)
)

/*
 * GVariantDict
 */

// VariantDict is a representation of GLib's VariantDict.
type VariantDict struct {
	gvariantDict *C.GVariantDict
}

func (v *VariantDict) toVariantDict() *C.GVariantDict {
	return v.native()
}

// newVariantDict creates a new VariantDict from a GVariantDict pointer.
func newVariantDict(p *C.GVariantDict) *VariantDict {
	return &VariantDict{p}
}

// native returns a pointer to the underlying GVariantDict.
func (v *VariantDict) native() *C.GVariantDict {
	if v == nil {
		return nil
	}
	return v.gvariantDict
}

// Native returns a pointer to the underlying GVariantDict.
func (v *VariantDict) Native() uintptr {
	return uintptr(unsafe.Pointer(v.native()))
}

/*
 * GVariant
 */

// IVariant is an interface type implemented by Variant and all types which embed
// an Variant.  It is meant to be used as a type for function arguments which
// require GVariants or any subclasses thereof.
type IVariant interface {
	toVariant() *C.GVariant
}

// A Variant is a representation of GLib's GVariant.
type Variant struct {
	gvariant *C.GVariant
}

// ToGVariant exposes the underlying *C.GVariant type for this Variant,
// necessary to implement IVariant.
func (v *Variant) native() *C.GVariant {
	if v == nil {
		return nil
	}
	return v.gvariant
}

func (v *Variant) toVariant() *C.GVariant {
	return v.native()
}

// Native returns a pointer to the underlying GVariant.
func (v *Variant) Native() unsafe.Pointer {
	return unsafe.Pointer(v.native())
}

/*
// ToGVariant type converts an unsafe.Pointer as a native C GVariant.
// This function is exported for visibility in other gotk3 packages and
// is not meant to be used by applications.
func ToVariant(p unsafe.Pointer) *C.GVariant {
	return C.toVariant(p)
}
*/

func WrapVariant(ptr unsafe.Pointer) *Variant {
	vr := &Variant{C.toGVariant(ptr)}

	if vr.IsFloating() {
		vr.RefSink()
	} else {
		vr.Ref()
	}

	runtime.SetFinalizer(vr, (*Variant).Unref)
	return vr
}

// newVariant creates a new Variant from a GVariant pointer.
func newVariant(p *C.GVariant) *Variant {
	return &Variant{p}
}

// VariantFromUnsafePointer returns a Variant from an unsafe pointer.
// XXX: unnecessary footgun?
//func VariantFromUnsafePointer(p unsafe.Pointer) *Variant {
//	return &Variant{C.toGVariant(p)}
//}

// TypeString returns the g variant type string for this variant.
func (v *Variant) TypeString() string {
	// the string returned from this belongs to GVariant and must not be freed.
	c := C.g_variant_get_type_string(v.native())
	return goString(c)
}

// IsContainer returns true if the variant is a container and false otherwise.
func (v *Variant) IsContainer() bool {
	return gobool(C.g_variant_is_container(v.native()))
}

// GetStrv returns a slice of strings from this variant.  It wraps
// g_variant_get_strv, but returns copies of the strings instead.
func (v *Variant) GetStrv() []string {
	c := C.g_variant_get_strv(v.native(), nil)
	// we do not own the memory for these strings, so we must not use strfreev
	// but we must free the actual pointer we receive.
	defer C.g_free(C.gpointer(c))

	strs := goStringArray(c)
	return strs
}

// GetInt returns the int64 value of the variant if it is an integer type, and
// an error otherwise.  It wraps variouns `g_variant_get_*` functions dealing
// with integers of different sizes.
func (v *Variant) GetInt() (int64, error) {
	var i int64
	if v.IsOfType(VARIANT_TYPE_BYTE) {
		i = int64(v.GetByte())
	} else if v.IsOfType(VARIANT_TYPE_INT16) {
		i = int64(v.GetInt16())
	} else if v.IsOfType(VARIANT_TYPE_UINT16) {
		i = int64(v.GetUInt16())
	} else if v.IsOfType(VARIANT_TYPE_INT32) {
		i = int64(v.GetInt32())
	} else if v.IsOfType(VARIANT_TYPE_UINT32) {
		i = int64(v.GetUInt32())
	} else if v.IsOfType(VARIANT_TYPE_INT64) {
		i = int64(v.GetInt64())
	} else if v.IsOfType(VARIANT_TYPE_UINT64) {
		i = int64(v.GetUInt64())
	} else {
		return 0, fmt.Errorf("variant type %v not an integer type",
			v.GetType())
	}

	return i, nil
}

// GetType returns the VariantType for this variant.
func (v *Variant) GetType() *VariantType {
	return newVariantType(C.g_variant_get_type(v.native()))
}

// IsOfType returns true if the variant's type matches t.
func (v *Variant) IsOfType(t *VariantType) bool {
	return gobool(C.g_variant_is_of_type(v.native(), t.native()))
}

// String wraps g_variant_print().  It returns a string understood
// by g_variant_parse().
func (v *Variant) String() string {
	if v.native() != nil {
		c := C.g_variant_print(v.native(), gbool(false))
		defer C.g_free(C.gpointer(c))

		return goString(c)
	}
	return "nil"
}

// AnnotatedString wraps g_variant_print(), but returns a type-annotated
// string.
func (v *Variant) AnnotatedString() string {
	if v.native() != nil {
		c := C.g_variant_print(v.native(), gbool(true))
		defer C.g_free(C.gpointer(c))

		return goString(c)
	}
	return "nil"
}

// gboolean
// g_variant_is_floating (GVariant *value);
func (v *Variant) IsFloating() bool {
	return gobool(C.g_variant_is_floating(v.native()))
}

// void
// g_variant_unref (GVariant *value);
func (v *Variant) Unref() {
	C.g_variant_unref(v.native())
}

// GVariant *
// g_variant_ref (GVariant *value);
func (v *Variant) Ref() {
	C.g_variant_ref(v.native())
}

// GVariant *
// g_variant_ref_sink (GVariant *value);
func (v *Variant) RefSink() {
	C.g_variant_ref_sink(v.native())
}

// GVariant *
// g_variant_take_ref (GVariant *value);
func (v *Variant) TakeRef() {
	C.g_variant_take_ref(v.native())
}

// GVariant *
// g_variant_new_boolean (gboolean value);
func VariantBooleanNew(value bool) (*Variant, error) {
	c := C.g_variant_new_boolean(gbool(value))
	if c == nil {
		return nil, errNilPtr
	}

	return WrapVariant(unsafe.Pointer(c)), nil
}

// GVariant *
// g_variant_new_string (const gchar *string);
func VariantStringNew(value string) (*Variant, error) {
	cstr := C.CString(value)
	defer C.free(unsafe.Pointer(cstr))

	c := C.g_variant_new_string((*C.gchar)(cstr))
	if c == nil {
		return nil, errNilPtr
	}

	return WrapVariant(unsafe.Pointer(c)), nil
}

// GVariant *
// g_variant_new_byte (guchar value);
func VariantByteNew(value byte) (*Variant, error) {
	c := C.g_variant_new_byte(C.guchar(value))
	if c == nil {
		return nil, errNilPtr
	}

	return WrapVariant(unsafe.Pointer(c)), nil
}

// GVariant *
// g_variant_new_int16 (gint16 value);
func VariantInt16New(value int16) (*Variant, error) {
	c := C.g_variant_new_int16(C.gint16(value))
	if c == nil {
		return nil, errNilPtr
	}

	return WrapVariant(unsafe.Pointer(c)), nil
}

// GVariant *
// g_variant_new_uint16 (guint16 value);
func VariantUInt16New(value uint16) (*Variant, error) {
	c := C.g_variant_new_uint16(C.guint16(value))
	if c == nil {
		return nil, errNilPtr
	}

	return WrapVariant(unsafe.Pointer(c)), nil
}

// GVariant *
// g_variant_new_int32 (gint32 value);
func VariantInt32New(value int32) (*Variant, error) {
	c := C.g_variant_new_int32(C.gint32(value))
	if c == nil {
		return nil, errNilPtr
	}

	return WrapVariant(unsafe.Pointer(c)), nil
}

// GVariant *
// g_variant_new_uint32 (guint32 value);
func VariantUInt32New(value uint32) (*Variant, error) {
	c := C.g_variant_new_uint32(C.guint32(value))
	if c == nil {
		return nil, errNilPtr
	}

	return WrapVariant(unsafe.Pointer(c)), nil
}

// GVariant *
// g_variant_new_int64 (gint64 value);
func VariantInt64New(value int64) (*Variant, error) {
	c := C.g_variant_new_int64(C.gint64(value))
	if c == nil {
		return nil, errNilPtr
	}

	return WrapVariant(unsafe.Pointer(c)), nil
}

// GVariant *
// g_variant_new_uint64 (guint64 value);
func VariantUInt64New(value uint64) (*Variant, error) {
	c := C.g_variant_new_uint64(C.guint64(value))
	if c == nil {
		return nil, errNilPtr
	}

	return WrapVariant(unsafe.Pointer(c)), nil
}

// GetBoolean returns the bool value of this variant.
func (v *Variant) GetBoolean() bool {
	return gobool(C.g_variant_get_boolean(v.native()))
}

// GetString returns the string value of the variant.
func (v *Variant) GetString() string {
	if v.native() != nil {
		var len C.gsize
		gc := C.g_variant_get_string(v.native(), &len)
		defer C.g_free(C.gpointer(gc))
		str := C.GoStringN((*C.char)(gc), (C.int)(len))
		return str
	}
	return "nil"
}

// guchar
// g_variant_get_byte (GVariant *value);
func (v *Variant) GetByte() byte {
	return byte(C.g_variant_get_byte(v.native()))
}

// gint16
// g_variant_get_int16 (GVariant *value);
func (v *Variant) GetInt16() int16 {
	return int16(C.g_variant_get_int16(v.native()))
}

// guint16
// g_variant_get_uint16 (GVariant *value);
func (v *Variant) GetUInt16() uint16 {
	return uint16(C.g_variant_get_uint16(v.native()))
}

// gint32
// g_variant_get_int32 (GVariant *value);
func (v *Variant) GetInt32() int32 {
	return int32(C.g_variant_get_int32(v.native()))
}

// guint32
// g_variant_get_uint32 (GVariant *value);
func (v *Variant) GetUInt32() uint32 {
	return uint32(C.g_variant_get_uint32(v.native()))
}

// gint64
// g_variant_get_int64 (GVariant *value);
func (v *Variant) GetInt64() int64 {
	return int64(C.g_variant_get_int64(v.native()))
}

// guint64
// g_variant_get_uint64 (GVariant *value);
func (v *Variant) GetUInt64() uint64 {
	return uint64(C.g_variant_get_uint64(v.native()))
}

//gint	g_variant_compare ()
//GVariantClass	g_variant_classify ()
//gboolean	g_variant_check_format_string ()
//void	g_variant_get ()
//void	g_variant_get_va ()
//GVariant *	g_variant_new ()
//GVariant *	g_variant_new_va ()
//GVariant *	g_variant_new_handle ()
//GVariant *	g_variant_new_double ()
//GVariant *	g_variant_new_take_string ()
//GVariant *	g_variant_new_printf ()
//GVariant *	g_variant_new_object_path ()
//gboolean	g_variant_is_object_path ()
//GVariant *	g_variant_new_signature ()
//gboolean	g_variant_is_signature ()
//GVariant *	g_variant_new_variant ()
//GVariant *	g_variant_new_strv ()
//GVariant *	g_variant_new_objv ()
//GVariant *	g_variant_new_bytestring ()
//GVariant *	g_variant_new_bytestring_array ()
//guint16	g_variant_get_uint16 ()
//gint32	g_variant_get_int32 ()
//guint32	g_variant_get_uint32 ()
//gint64	g_variant_get_int64 ()
//guint64	g_variant_get_uint64 ()
//gint32	g_variant_get_handle ()
//gdouble	g_variant_get_double ()
//const gchar *	g_variant_get_string ()
//gchar *	g_variant_dup_string ()
//GVariant *	g_variant_get_variant ()
//const gchar **	g_variant_get_strv ()
//gchar **	g_variant_dup_strv ()
//const gchar **	g_variant_get_objv ()
//gchar **	g_variant_dup_objv ()
//const gchar *	g_variant_get_bytestring ()
//gchar *	g_variant_dup_bytestring ()
//const gchar **	g_variant_get_bytestring_array ()
//gchar **	g_variant_dup_bytestring_array ()
//GVariant *	g_variant_new_maybe ()
//GVariant *	g_variant_new_array ()
//GVariant *	g_variant_new_tuple ()
//GVariant *	g_variant_new_dict_entry ()
//GVariant *	g_variant_new_fixed_array ()
//GVariant *	g_variant_get_maybe ()
//gsize	g_variant_n_children ()
//GVariant *	g_variant_get_child_value ()
//void	g_variant_get_child ()
//GVariant *	g_variant_lookup_value ()
//gboolean	g_variant_lookup ()
//gconstpointer	g_variant_get_fixed_array ()
//gsize	g_variant_get_size ()
//gconstpointer	g_variant_get_data ()
//GBytes *	g_variant_get_data_as_bytes ()
//void	g_variant_store ()
//GVariant *	g_variant_new_from_data ()
//GVariant *	g_variant_new_from_bytes ()
//GVariant *	g_variant_byteswap ()
//GVariant *	g_variant_get_normal_form ()
//gboolean	g_variant_is_normal_form ()
//guint	g_variant_hash ()
//gboolean	g_variant_equal ()
//gchar *	g_variant_print ()
//GString *	g_variant_print_string ()
//GVariantIter *	g_variant_iter_copy ()
//void	g_variant_iter_free ()
//gsize	g_variant_iter_init ()
//gsize	g_variant_iter_n_children ()
//GVariantIter *	g_variant_iter_new ()
//GVariant *	g_variant_iter_next_value ()
//gboolean	g_variant_iter_next ()
//gboolean	g_variant_iter_loop ()
//void	g_variant_builder_unref ()
//GVariantBuilder *	g_variant_builder_ref ()
//GVariantBuilder *	g_variant_builder_new ()
//void	g_variant_builder_init ()
//void	g_variant_builder_clear ()
//void	g_variant_builder_add_value ()
//void	g_variant_builder_add ()
//void	g_variant_builder_add_parsed ()
//GVariant *	g_variant_builder_end ()
//void	g_variant_builder_open ()
//void	g_variant_builder_close ()
//void	g_variant_dict_unref ()
//GVariantDict *	g_variant_dict_ref ()
//GVariantDict *	g_variant_dict_new ()
//void	g_variant_dict_init ()
//void	g_variant_dict_clear ()
//gboolean	g_variant_dict_contains ()
//gboolean	g_variant_dict_lookup ()
//GVariant *	g_variant_dict_lookup_value ()
//void	g_variant_dict_insert ()
//void	g_variant_dict_insert_value ()
//gboolean	g_variant_dict_remove ()
//GVariant *	g_variant_dict_end ()
//#define	G_VARIANT_PARSE_ERROR
//GVariant *	g_variant_parse ()
//GVariant *	g_variant_new_parsed_va ()
//GVariant *	g_variant_new_parsed ()
//gchar *	g_variant_parse_error_print_context ()

/*
 * GVariantBuilder
 */

// VariantBuilder is a representation of GLib's VariantBuilder.
type VariantBuilder struct {
	gvariantBuilder *C.GVariantBuilder
}

func (v *VariantBuilder) toVariantBuilder() *C.GVariantBuilder {
	return v.native()
}

// newVariantBuilder creates a new VariantBuilder from a GVariantBuilder pointer.
func newVariantBuilder(p *C.GVariantBuilder) *VariantBuilder {
	return &VariantBuilder{p}
}

// native returns a pointer to the underlying GVariantBuilder.
func (v *VariantBuilder) native() *C.GVariantBuilder {
	if v == nil {
		return nil
	}
	return v.gvariantBuilder
}

// Native returns a pointer to the underlying GVariantBuilder.
func (v *VariantBuilder) Native() uintptr {
	return uintptr(unsafe.Pointer(v.native()))
}

/*
 * GVariantIter
 */

// VariantIter is a representation of GLib's GVariantIter.
type VariantIter struct {
	gvariantIter *C.GVariantIter
}

func (v *VariantIter) toVariantIter() *C.GVariantIter {
	return v.native()
}

// newVariantIter creates a new VariantIter from a GVariantIter pointer.
func newVariantIter(p *C.GVariantIter) *VariantIter {
	return &VariantIter{p}
}

// native returns a pointer to the underlying GVariantIter.
func (v *VariantIter) native() *C.GVariantIter {
	if v == nil {
		return nil
	}
	return v.gvariantIter
}

// Native returns a pointer to the underlying GVariantIter.
func (v *VariantIter) Native() uintptr {
	return uintptr(unsafe.Pointer(v.native()))
}
