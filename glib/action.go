package glib

// #cgo pkg-config: glib-2.0 gobject-2.0
// #include <gio/gio.h>
// #include <glib.h>
// #include <glib-object.h>
// #include "glib.go.h"
// #include "action.go.h"
import "C"
import (
	"errors"

	"unsafe"
)

// IAction is an interface type implemented by all structs
// embedding a Action.
type IAction interface {
	toAction() *C.GAction
}

// Action is a representation of GAction GInterface.
type Action struct {
	Interface
}

// Static cast to verify at compile time that type on the right side
// implement corresponding interface on the left.
var _ IAction = &Action{}

// native() returns a pointer to the underlying GAction.
func (v *Action) native() *C.GAction {
	return C.toGAction(unsafe.Pointer(v.Native()))
}

func (v *Action) toAction() *C.GAction {
	return v.native()
}

func marshalAction(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := Take(unsafe.Pointer(c))
	act := wrapAction(*InterfaceFromObjectNew(obj))
	return act, nil
}

func wrapAction(intf Interface) *Action {
	return &Action{intf}
}

// gboolean
// g_action_name_is_valid (const gchar *action_name);
func ActionNameIsValid(actionName string) bool {
	cstr := C.CString(actionName)
	defer C.free(unsafe.Pointer(cstr))

	c := C.g_action_name_is_valid((*C.gchar)(cstr))
	return gobool(c)
}

// const gchar *
// g_action_get_name (GAction *action);
func (v *Action) GetName() (string, error) {
	c := C.g_action_get_name(v.native())
	if c == nil {
		return "", errNilPtr
	}

	return goString(c), nil
}

// const GVariantType *
// g_action_get_parameter_type (GAction *action);
func (v *Action) GetParameterType() *VariantType {
	c := C.g_action_get_parameter_type(v.native())
	if c == nil {
		return nil
	}

	return newVariantType(c)
}

// const GVariantType *
// g_action_get_state_type (GAction *action);
func (v *Action) GetStateType() *VariantType {
	c := C.g_action_get_state_type(v.native())
	if c == nil {
		return nil
	}

	return newVariantType(c)
}

// GVariant *
// g_action_get_state_hint (GAction *action);
func (v *Action) GetStateHint() *Variant {
	c := C.g_action_get_state_hint(v.native())
	if c == nil {
		return nil
	}

	return WrapVariant(unsafe.Pointer(c))
}

// gboolean
// g_action_get_enabled (GAction *action);
func (v *Action) GetEnabled() bool {
	c := C.g_action_get_enabled(v.native())

	return gobool(c)
}

// GVariant *
// g_action_get_state (GAction *action);
func (v *Action) GetState() *Variant {
	c := C.g_action_get_state(v.native())
	if c == nil {
		return nil
	}

	return WrapVariant(unsafe.Pointer(c))
}

// void
// g_action_change_state (GAction *action,
//                        GVariant *value);
func (v *Action) ChangeState(value *Variant) {
	C.g_action_change_state(v.native(), value.native())
}

// void
// g_action_activate (GAction *action,
//                    GVariant *parameter);
func (v *Action) Activate(parameter *Variant) {
	C.g_action_activate(v.native(), parameter.native())
}

// gboolean
// g_action_parse_detailed_name (const gchar *detailed_name,
//                               gchar **action_name,
//                               GVariant **target_value,
//                               GError **error);
func ParseDetailedName(detailedName string) (actionName string, targetValue *Variant, e error) {
	cstr := C.CString(detailedName)
	defer C.free(unsafe.Pointer(cstr))

	var an *C.gchar
	var tv *C.GVariant
	var err *C.GError
	c := C.g_action_parse_detailed_name(cstr, &an, &tv, &err)
	if c == 0 {
		defer C.g_error_free(err)
		return "", nil, errors.New(goString(err.message))
	}

	defer C.g_free(C.gpointer(an))
	return goString(an), WrapVariant(unsafe.Pointer(tv)), nil
}

// gchar *
// g_action_print_detailed_name (const gchar *action_name,
//                               GVariant *target_value);
func PrintDetailedName(actionName string, targetValue *Variant) string {
	cstr := C.CString(actionName)
	defer C.free(unsafe.Pointer(cstr))

	c := C.g_action_print_detailed_name((*C.gchar)(cstr), targetValue.native())
	defer C.g_free(C.gpointer(c))

	return goString(c)
}

// SimpleAction is a representation of GSimpleAction.
type SimpleAction struct {
	// This must be a pointer so copies of the ref-sinked object
	// do not outlive the original object, causing an unref
	// finalizer to prematurely run.
	*Object
	// Interfaces
	Action
}

// Static cast to verify at compile time that type on the right side
// implement corresponding interface on the left.
var _ IAction = &SimpleAction{}

// native() returns a pointer to the underlying GSimpleAction.
func (v *SimpleAction) native() *C.GSimpleAction {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toGSimpleAction(ptr)
}

func marshalSimpleAction(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := Take(unsafe.Pointer(c))
	return wrapSimpleAction(obj), nil
}

func wrapSimpleAction(obj *Object) *SimpleAction {
	action := wrapAction(*InterfaceFromObjectNew(obj))
	return &SimpleAction{obj, *action}
}

// GSimpleAction *
// g_simple_action_new (const gchar *name,
//                      const GVariantType *parameter_type);
func SimpleActionNew(name string, parameterType *VariantType) (*SimpleAction, error) {
	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))

	var c *C.GSimpleAction
	c = C.g_simple_action_new((*C.gchar)(cstr), parameterType.native())
	if c == nil {
		return nil, errNilPtr
	}

	obj := Take(unsafe.Pointer(c))
	return wrapSimpleAction(obj), nil
}

// GSimpleAction *
// g_simple_action_new_stateful (const gchar *name,
//                               const GVariantType *parameter_type,
//                               GVariant *state);
func SimpleActionStatefullNew(name string, parameterType *VariantType,
	state *Variant) (*SimpleAction, error) {
	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))

	var c *C.GSimpleAction
	c = C.g_simple_action_new_stateful((*C.gchar)(cstr),
		parameterType.native(), state.native())
	if c == nil {
		return nil, errNilPtr
	}

	obj := Take(unsafe.Pointer(c))
	return wrapSimpleAction(obj), nil
}

// void
// g_simple_action_set_enabled (GSimpleAction *simple,
//                              gboolean enabled);
func (v *SimpleAction) SetEnabled(enabled bool) {
	C.g_simple_action_set_enabled(v.native(), gbool(enabled))
}

// void
// g_simple_action_set_state (GSimpleAction *simple,
//                            GVariant *value);
func (v *SimpleAction) SetState(value *Variant) {
	C.g_simple_action_set_state(v.native(), value.native())
}

// ActionMap is a representation of GActionMap.
type ActionMap struct {
	Interface
}

// native() returns a pointer to the underlying GAction.
func (v *ActionMap) native() *C.GActionMap {
	return C.toGActionMap(unsafe.Pointer(v.Native()))
}

func marshalActionMap(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := Take(unsafe.Pointer(c))
	act := wrapActionMap(*InterfaceFromObjectNew(obj))
	return act, nil
}

func wrapActionMap(intf Interface) *ActionMap {
	return &ActionMap{intf}
}

func ToActionMap(obj *Object) *ActionMap {
	v := wrapActionMap(*InterfaceFromObjectNew(obj))
	return v
}

// GAction *
// g_action_map_lookup_action (GActionMap *action_map,
//                             const gchar *action_name);
func (v *ActionMap) LookupAction(actionName string) *Action {
	cstr := C.CString(actionName)
	defer C.free(unsafe.Pointer(cstr))

	c := C.g_action_map_lookup_action(v.native(), (*C.gchar)(cstr))
	if c != nil {
		intf := InterfaceNew(unsafe.Pointer(c))
		return wrapAction(*intf)
	}
	return nil
}

// void
// g_action_map_add_action_entries (GActionMap *action_map,
//                                  const GActionEntry *entries,
//                                  gint n_entries,
//                                  gpointer user_data);

// void
// g_action_map_add_action (GActionMap *action_map,
//                          GAction *action);
func (v *ActionMap) AddAction(action IAction) {
	C.g_action_map_add_action(v.native(), action.toAction())
}

// void
// g_action_map_remove_action (GActionMap *action_map,
//                             const gchar *action_name);
func (v *ActionMap) RemoveAction(actionName string) {
	cstr := C.CString(actionName)
	defer C.free(unsafe.Pointer(cstr))

	C.g_action_map_remove_action(v.native(), (*C.gchar)(cstr))
}

func init() {
	tm := []TypeMarshaler{
		// Enums
		{Type(C.g_action_get_type()), marshalAction},
		{Type(C.g_simple_action_get_type()), marshalSimpleAction},
		{Type(C.g_action_map_get_type()), marshalActionMap},
	}
	RegisterGValueMarshalers(tm)
}
