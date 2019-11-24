package glib

// #cgo pkg-config: glib-2.0 gobject-2.0
// #include <gio/gio.h>
// #include <glib.h>
// #include <glib-object.h>
// #include "glib.go.h"
import "C"
import (
	"unsafe"
)

// Settings is a representation of GSettings.
type Settings struct {
	*Object
}

// native() returns a pointer to the underlying GSettings.
func (v *Settings) native() *C.GSettings {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toGSettings(ptr)
}

func marshalSettings(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	return wrapSettings(wrapObject(unsafe.Pointer(c))), nil
}

func wrapSettings(obj *Object) *Settings {
	return &Settings{obj}
}

// SettingsNew is a wrapper around g_settings_new().
func SettingsNew(schemaID string) (*Settings, error) {
	cstr := C.CString(schemaID)
	defer C.free(unsafe.Pointer(cstr))

	c := C.g_settings_new((*C.gchar)(cstr))
	if c == nil {
		return nil, errNilPtr
	}
	return wrapSettings(wrapObject(unsafe.Pointer(c))), nil
}

// SettingsNewWithPath is a wrapper around g_settings_new_with_path().
func SettingsNewWithPath(schemaID, path string) (*Settings, error) {
	cstr1 := C.CString(schemaID)
	defer C.free(unsafe.Pointer(cstr1))

	cstr2 := C.CString(path)
	defer C.free(unsafe.Pointer(cstr2))

	c := C.g_settings_new_with_path((*C.gchar)(cstr1), (*C.gchar)(cstr2))
	if c == nil {
		return nil, errNilPtr
	}
	return wrapSettings(wrapObject(unsafe.Pointer(c))), nil
}

// SettingsNewWithBackend is a wrapper around g_settings_new_with_backend().
func SettingsNewWithBackend(schemaID string, backend *SettingsBackend) (*Settings, error) {
	cstr := C.CString(schemaID)
	defer C.free(unsafe.Pointer(cstr))

	c := C.g_settings_new_with_backend((*C.gchar)(cstr), backend.native())
	if c == nil {
		return nil, errNilPtr
	}
	return wrapSettings(wrapObject(unsafe.Pointer(c))), nil
}

// SettingsNewWithBackendAndPath is a wrapper around g_settings_new_with_backend_and_path().
func SettingsNewWithBackendAndPath(schemaID string, backend *SettingsBackend, path string) (*Settings, error) {
	cstr1 := C.CString(schemaID)
	defer C.free(unsafe.Pointer(cstr1))

	cstr2 := C.CString(path)
	defer C.free(unsafe.Pointer(cstr2))

	c := C.g_settings_new_with_backend_and_path((*C.gchar)(cstr1),
		backend.native(), (*C.gchar)(cstr2))
	if c == nil {
		return nil, errNilPtr
	}
	return wrapSettings(wrapObject(unsafe.Pointer(c))), nil
}

// SettingsNewFull is a wrapper around g_settings_new_full().
func SettingsNewFull(schema *SettingsSchema, backend *SettingsBackend, path string) (*Settings, error) {
	cstr := C.CString(path)
	defer C.free(unsafe.Pointer(cstr))

	c := C.g_settings_new_full(schema.native(), backend.native(), (*C.gchar)(cstr))
	if c == nil {
		return nil, errNilPtr
	}
	return wrapSettings(wrapObject(unsafe.Pointer(c))), nil
}

// SettingsSync is a wrapper around g_settings_sync().
func SettingsSync() {
	C.g_settings_sync()
}

// IsWritable is a wrapper around g_settings_is_writable().
func (v *Settings) IsWritable(name string) bool {
	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))

	return gobool(C.g_settings_is_writable(v.native(), (*C.gchar)(cstr)))
}

// Delay is a wrapper around g_settings_delay().
func (v *Settings) Delay() {
	C.g_settings_delay(v.native())
}

// Apply is a wrapper around g_settings_apply().
func (v *Settings) Apply() {
	C.g_settings_apply(v.native())
}

// Revert is a wrapper around g_settings_revert().
func (v *Settings) Revert() {
	C.g_settings_revert(v.native())
}

// GetHasUnapplied is a wrapper around g_settings_get_has_unapplied().
func (v *Settings) GetHasUnapplied() bool {
	return gobool(C.g_settings_get_has_unapplied(v.native()))
}

// GetChild is a wrapper around g_settings_get_child().
func (v *Settings) GetChild(name string) (*Settings, error) {
	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))

	c := C.g_settings_get_child(v.native(), (*C.gchar)(cstr))
	if c == nil {
		return nil, errNilPtr
	}
	return wrapSettings(wrapObject(unsafe.Pointer(c))), nil
}

// Reset is a wrapper around g_settings_reset().
func (v *Settings) Reset(key string) {
	cstr := C.CString(key)
	defer C.free(unsafe.Pointer(cstr))

	C.g_settings_reset(v.native(), (*C.gchar)(cstr))
}

// ListChildren is a wrapper around g_settings_list_children().
func (v *Settings) ListChildren() []string {
	c := C.g_settings_list_children(v.native())
	// both pointer array and strings should be freed.
	defer C.g_strfreev(c)

	strs := goStringArray(c)
	return strs
}

// GetBoolean is a wrapper around g_settings_get_boolean().
func (v *Settings) GetBoolean(key string) bool {
	cstr := C.CString(key)
	defer C.free(unsafe.Pointer(cstr))

	return gobool(C.g_settings_get_boolean(v.native(), (*C.gchar)(cstr)))
}

// SetBoolean is a wrapper around g_settings_set_boolean().
func (v *Settings) SetBoolean(key string, value bool) bool {
	cstr := C.CString(key)
	defer C.free(unsafe.Pointer(cstr))

	return gobool(C.g_settings_set_boolean(v.native(), (*C.gchar)(cstr), gbool(value)))
}

// GetInt is a wrapper around g_settings_get_int().
func (v *Settings) GetInt(key string) int {
	cstr := C.CString(key)
	defer C.free(unsafe.Pointer(cstr))

	return int(C.g_settings_get_int(v.native(), (*C.gchar)(cstr)))
}

// SetInt is a wrapper around g_settings_set_int().
func (v *Settings) SetInt(key string, value int) bool {
	cstr := C.CString(key)
	defer C.free(unsafe.Pointer(cstr))

	return gobool(C.g_settings_set_int(v.native(), (*C.gchar)(cstr), C.gint(value)))
}

// GetUInt is a wrapper around g_settings_get_uint().
func (v *Settings) GetUInt(key string) uint {
	cstr := C.CString(key)
	defer C.free(unsafe.Pointer(cstr))

	return uint(C.g_settings_get_uint(v.native(), (*C.gchar)(cstr)))
}

// SetUInt is a wrapper around g_settings_set_uint().
func (v *Settings) SetUInt(key string, value uint) bool {
	cstr := C.CString(key)
	defer C.free(unsafe.Pointer(cstr))

	return gobool(C.g_settings_set_uint(v.native(), (*C.gchar)(cstr), C.guint(value)))
}

// GetDouble is a wrapper around g_settings_get_double().
func (v *Settings) GetDouble(key string) float64 {
	cstr := C.CString(key)
	defer C.free(unsafe.Pointer(cstr))

	return float64(C.g_settings_get_double(v.native(), (*C.gchar)(cstr)))
}

// SetDouble is a wrapper around g_settings_set_double().
func (v *Settings) SetDouble(key string, value float64) bool {
	cstr := C.CString(key)
	defer C.free(unsafe.Pointer(cstr))

	return gobool(C.g_settings_set_double(v.native(), (*C.gchar)(cstr), C.gdouble(value)))
}

// GetString is a wrapper around g_settings_get_string().
func (v *Settings) GetString(key string) string {
	cstr := C.CString(key)
	defer C.free(unsafe.Pointer(cstr))

	c := C.g_settings_get_string(v.native(), (*C.gchar)(cstr))
	defer C.g_free(C.gpointer(c))

	return goString(c)
}

// SetString is a wrapper around g_settings_set_string().
func (v *Settings) SetString(key string, value string) bool {
	cstr1 := C.CString(key)
	defer C.free(unsafe.Pointer(cstr1))

	cstr2 := C.CString(value)
	defer C.free(unsafe.Pointer(cstr2))

	return gobool(C.g_settings_set_string(v.native(), (*C.gchar)(cstr1), (*C.gchar)(cstr2)))
}

// GetEnum is a wrapper around g_settings_get_enum().
func (v *Settings) GetEnum(key string) int {
	cstr := C.CString(key)
	defer C.free(unsafe.Pointer(cstr))

	return int(C.g_settings_get_enum(v.native(), (*C.gchar)(cstr)))
}

// SetEnum is a wrapper around g_settings_set_enum().
func (v *Settings) SetEnum(key string, value int) bool {
	cstr := C.CString(key)
	defer C.free(unsafe.Pointer(cstr))

	return gobool(C.g_settings_set_enum(v.native(), (*C.gchar)(cstr), C.gint(value)))
}

// GetFlags is a wrapper around g_settings_get_flags().
func (v *Settings) GetFlags(key string) uint {
	cstr := C.CString(key)
	defer C.free(unsafe.Pointer(cstr))

	return uint(C.g_settings_get_flags(v.native(), (*C.gchar)(cstr)))
}

// SetFlags is a wrapper around g_settings_set_flags().
func (v *Settings) SetFlags(key string, value uint) bool {
	cstr := C.CString(key)
	defer C.free(unsafe.Pointer(cstr))

	return gobool(C.g_settings_set_flags(v.native(), (*C.gchar)(cstr), C.guint(value)))
}

// GVariant * 	g_settings_get_value ()
func (v *Settings) GetValue(key string) *Variant {
	cstr := C.CString(key)
	defer C.free(unsafe.Pointer(cstr))

	c := C.g_settings_get_value(v.native(), (*C.gchar)(cstr))
	return WrapVariant(unsafe.Pointer(c))
}

// gboolean 	g_settings_set_value ()
func (v *Settings) SetValue(key string, value *Variant) bool {
	cstr := C.CString(key)
	defer C.free(unsafe.Pointer(cstr))

	return gobool(C.g_settings_set_value(v.native(), (*C.gchar)(cstr), value.native()))
}

// SettingsBindFlags is a representation of GLib's GSettingsBindFlags.
type SettingsBindFlags uint

const (
	SETTINGS_BIND_DEFAULT        SettingsBindFlags = C.G_SETTINGS_BIND_DEFAULT
	SETTINGS_BIND_GET            SettingsBindFlags = C.G_SETTINGS_BIND_GET
	SETTINGS_BIND_SET            SettingsBindFlags = C.G_SETTINGS_BIND_SET
	SETTINGS_BIND_NO_SENSITIVITY SettingsBindFlags = C.G_SETTINGS_BIND_NO_SENSITIVITY
	SETTINGS_BIND_GET_NO_CHANGES SettingsBindFlags = C.G_SETTINGS_BIND_GET_NO_CHANGES
	SETTINGS_BIND_INVERT_BOOLEAN SettingsBindFlags = C.G_SETTINGS_BIND_INVERT_BOOLEAN
)

// void 	g_settings_bind ()
func (v *Settings) Bind(key string, object IObject, property string, flags SettingsBindFlags) {
	ckey := C.CString(key)
	defer C.free(unsafe.Pointer(ckey))
	cproperty := C.CString(property)
	defer C.free(unsafe.Pointer(cproperty))

	C.g_settings_bind(v.native(), (*C.gchar)(ckey),
		C.gpointer(unsafe.Pointer(object.toObject().native())),
		(*C.gchar)(cproperty), C.GSettingsBindFlags(flags))
}

// void 	g_settings_unbind ()
func (v *Settings) Unbind(object IObject, property string) {
	cproperty := C.CString(property)
	defer C.free(unsafe.Pointer(cproperty))

	C.g_settings_unbind(C.gpointer(unsafe.Pointer(object.toObject().native())),
		(*C.gchar)(cproperty))
}

// gchar ** 	g_settings_get_strv ()
func (v *Settings) GetStrv(key string) []string {
	cstr := C.CString(key)
	defer C.free(unsafe.Pointer(cstr))

	c := C.g_settings_get_strv(v.native(), cstr)
	// we do not own the memory for these strings, so we must not use strfreev
	// but we must free the actual pointer we receive.
	defer C.g_free(C.gpointer(c))

	strs := goStringArray(c)
	return strs
}

// gboolean 	g_settings_set_strv ()
func (v *Settings) SetStrv(key string, value []string) bool {
	cstr := C.CString(key)
	defer C.free(unsafe.Pointer(cstr))

	count := C.int(len(value))
	cvalue := C.make_strings(count + 1)
	defer C.destroy_strings(cvalue)

	for i, str := range value {
		cval := C.CString(str)
		defer C.free(unsafe.Pointer(cval))
		C.set_string(cvalue, C.int(i), (*C.gchar)(cval))
	}
	C.set_string(cvalue, C.int(len(value)), nil)

	c := C.g_settings_set_strv(v.native(), cstr, cvalue)
	return gobool(c)
}

// gchar ** 	g_settings_list_keys ()
// GVariant * 	g_settings_get_user_value ()
// GVariant * 	g_settings_get_default_value ()
// const gchar * const * 	g_settings_list_schemas ()
// const gchar * const * 	g_settings_list_relocatable_schemas ()
// GVariant * 	g_settings_get_range ()
// gboolean 	g_settings_range_check ()
// void 	g_settings_get ()
// gboolean 	g_settings_set ()
// gpointer 	g_settings_get_mapped ()
// void 	g_settings_bind_with_mapping ()
// void 	g_settings_bind_writable ()
// gaction * 	g_settings_create_action ()

// SettingsSchema is a representation of GSettingsSchema.
type SettingsSchema struct {
	schema *C.GSettingsSchema
}

func wrapSettingsSchema(obj *C.GSettingsSchema) *SettingsSchema {
	return &SettingsSchema{obj}
}

func (v *SettingsSchema) Native() uintptr {
	return uintptr(unsafe.Pointer(v.native()))
}

func (v *SettingsSchema) native() *C.GSettingsSchema {
	if v == nil {
		return nil
	}
	return v.schema
}

func marshalSettingsSchema(p uintptr) (interface{}, error) {
	c := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return &SettingsSchema{(*C.GSettingsSchema)(unsafe.Pointer(c))}, nil
}

// Ref is a wrapper around g_settings_schema_ref().
func (v *SettingsSchema) Ref() *SettingsSchema {
	return wrapSettingsSchema(C.g_settings_schema_ref(v.native()))
}

// Unref is a wrapper around g_settings_schema_unref().
func (v *SettingsSchema) Unref() {
	C.g_settings_schema_unref(v.native())
}

// GetID is a wrapper around g_settings_schema_get_id().
func (v *SettingsSchema) GetID() string {
	c := C.g_settings_schema_get_id(v.native())
	return goString(c)
}

// GetPath is a wrapper around g_settings_schema_get_path().
func (v *SettingsSchema) GetPath() *string {
	c := C.g_settings_schema_get_path(v.native())
	if c == nil {
		return nil
	}
	str := goString(c)
	return &str
}

// HasKey is a wrapper around g_settings_schema_has_key().
func (v *SettingsSchema) HasKey(v1 string) bool {
	cstr := C.CString(v1)
	defer C.free(unsafe.Pointer(cstr))

	return gobool(C.g_settings_schema_has_key(v.native(), (*C.gchar)(cstr)))
}

// // ListChildren() is a wrapper around g_settings_schema_list_children().
// func (v *SettingsSchema) ListChildren() []string {
// 	return toGoStringArray(C.g_settings_schema_list_children(v.native()))
// }

// // ListKeys() is a wrapper around g_settings_schema_list_keys().
// func (v *SettingsSchema) ListKeys() []string {
// 	return toGoStringArray(C.g_settings_schema_list_keys(v.native()))
// }

// const GVariantType * 	g_settings_schema_key_get_value_type ()
// GVariant * 	g_settings_schema_key_get_default_value ()
// GVariant * 	g_settings_schema_key_get_range ()
// gboolean 	g_settings_schema_key_range_check ()
// const gchar * 	g_settings_schema_key_get_name ()
// const gchar * 	g_settings_schema_key_get_summary ()
// const gchar * 	g_settings_schema_key_get_description ()

// GSettingsSchemaKey * 	g_settings_schema_get_key ()
// GSettingsSchemaKey * 	g_settings_schema_key_ref ()
// void 	g_settings_schema_key_unref ()

// SettingsSchemaSource is a representation of GSettingsSchemaSource.
type SettingsSchemaSource struct {
	source *C.GSettingsSchemaSource
}

func wrapSettingsSchemaSource(obj *C.GSettingsSchemaSource) *SettingsSchemaSource {
	return &SettingsSchemaSource{obj}
}

func (v *SettingsSchemaSource) Native() uintptr {
	return uintptr(unsafe.Pointer(v.native()))
}

func (v *SettingsSchemaSource) native() *C.GSettingsSchemaSource {
	if v == nil {
		return nil
	}
	return v.source
}

// SettingsSchemaSourceGetDefault is a wrapper around g_settings_schema_source_get_default().
func SettingsSchemaSourceGetDefault() *SettingsSchemaSource {
	c := C.g_settings_schema_source_get_default()
	// Null return is allowed.
	if c == nil {
		return nil
	}
	return wrapSettingsSchemaSource(c)
}

// Ref is a wrapper around g_settings_schema_source_ref().
func (v *SettingsSchemaSource) Ref() *SettingsSchemaSource {
	return wrapSettingsSchemaSource(C.g_settings_schema_source_ref(v.native()))
}

// Unref is a wrapper around g_settings_schema_source_unref().
func (v *SettingsSchemaSource) Unref() {
	C.g_settings_schema_source_unref(v.native())
}

// SettingsSchemaSourceNewFromDirectory is a wrapper around g_settings_schema_source_new_from_directory().
func SettingsSchemaSourceNewFromDirectory(dir string, parent *SettingsSchemaSource, trusted bool) (*SettingsSchemaSource, error) {
	cstr := C.CString(dir)
	defer C.free(unsafe.Pointer(cstr))

	c := C.g_settings_schema_source_new_from_directory((*C.gchar)(cstr), parent.native(), gbool(trusted), nil)
	if c == nil {
		return nil, errNilPtr
	}
	return wrapSettingsSchemaSource(c), nil
}

// Lookup is a wrapper around g_settings_schema_source_lookup().
func (v *SettingsSchemaSource) Lookup(schema string, recursive bool) *SettingsSchema {
	cstr := C.CString(schema)
	defer C.free(unsafe.Pointer(cstr))

	c := C.g_settings_schema_source_lookup(v.native(), (*C.gchar)(cstr), gbool(recursive))
	// Null return is allowed.
	if c == nil {
		return nil
	}

	return wrapSettingsSchema(c)
}

// ListSchemas is a wrapper around 	g_settings_schema_source_list_schemas().
func (v *SettingsSchemaSource) ListSchemas(recursive bool) (nonReolcatable, relocatable []string) {
	var nonRel, rel **C.gchar
	C.g_settings_schema_source_list_schemas(v.native(), gbool(recursive), &nonRel, &rel)
	// both pointer array and strings should be freed.
	defer C.g_strfreev(nonRel)
	defer C.g_strfreev(rel)

	nonRelStrs := goStringArray(nonRel)
	relStrs := goStringArray(rel)
	return nonRelStrs, relStrs
}

// SettingsBackend is a representation of GSettingsBackend.
type SettingsBackend struct {
	*Object
}

// native returns a pointer to the underlying GSettingsBackend.
func (v *SettingsBackend) native() *C.GSettingsBackend {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toGSettingsBackend(ptr)
}

func marshalSettingsBackend(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	return wrapSettingsBackend(wrapObject(unsafe.Pointer(c))), nil
}

func wrapSettingsBackend(obj *Object) *SettingsBackend {
	return &SettingsBackend{obj}
}

// SettingsBackendGetDefault is a wrapper around g_settings_backend_get_default().
func SettingsBackendGetDefault() *SettingsBackend {
	return wrapSettingsBackend(wrapObject(unsafe.Pointer(C.g_settings_backend_get_default())))
}

// KeyfileSettingsBackendNew is a wrapper around g_keyfile_settings_backend_new().
func KeyfileSettingsBackendNew(filename, rootPath, rootGroup string) (*SettingsBackend, error) {
	cstr1 := C.CString(filename)
	defer C.free(unsafe.Pointer(cstr1))

	cstr2 := C.CString(rootPath)
	defer C.free(unsafe.Pointer(cstr2))

	cstr3 := C.CString(rootGroup)
	defer C.free(unsafe.Pointer(cstr3))

	c := C.g_keyfile_settings_backend_new((*C.gchar)(cstr1), (*C.gchar)(cstr2), (*C.gchar)(cstr3))
	if c == nil {
		return nil, errNilPtr
	}

	return wrapSettingsBackend(wrapObject(unsafe.Pointer(c))), nil
}

// MemorySettingsBackendNew is a wrapper around g_memory_settings_backend_new().
func MemorySettingsBackendNew() (*SettingsBackend, error) {
	c := C.g_memory_settings_backend_new()
	if c == nil {
		return nil, errNilPtr
	}

	return wrapSettingsBackend(wrapObject(unsafe.Pointer(c))), nil
}

// NullSettingsBackendNew is a wrapper around g_null_settings_backend_new().
func NullSettingsBackendNew() (*SettingsBackend, error) {
	c := C.g_null_settings_backend_new()
	if c == nil {
		return nil, errNilPtr
	}

	return wrapSettingsBackend(wrapObject(unsafe.Pointer(c))), nil
}

// void 	g_settings_backend_changed ()
// void 	g_settings_backend_path_changed ()
// void 	g_settings_backend_keys_changed ()
// void 	g_settings_backend_path_writable_changed ()
// void 	g_settings_backend_writable_changed ()
// void 	g_settings_backend_changed_tree ()
// void 	g_settings_backend_flatten_tree ()

func init() {
	tm := []TypeMarshaler{
		{Type(C.g_settings_schema_get_type()), marshalSettingsSchema},
	}
	RegisterGValueMarshalers(tm)
}
