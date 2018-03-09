package gtk

/*
 #cgo pkg-config: gtk+-3.0
 #include <gtk/gtk.h>
*/
import "C"
import (
	"strings"
	"unsafe"

	"github.com/d2r2/gotk3/glib"
)

//export substring_match_equal_func
func substring_match_equal_func(model *C.GtkTreeModel,
	column C.gint,
	key *C.gchar,
	iter *C.GtkTreeIter,
	data C.gpointer) C.gboolean {

	obj := glib.Take(unsafe.Pointer(model))

	goModel := wrapTreeModel(*glib.InterfaceFromObjectNew(obj))
	goIter := wrapTreeIter(iter)

	value, err := goModel.GetValue(goIter, int(column))
	if err != nil {
		return gbool(true)
	}

	str, _ := value.GetString()
	if str == "" {
		return gbool(true)
	}

	subStr := goString(key)
	res := strings.Contains(str, subStr)
	return gbool(!res)
}

//export goBuilderConnect
func goBuilderConnect(builder *C.GtkBuilder,
	object *C.GObject,
	signal_name *C.gchar,
	handler_name *C.gchar,
	connect_object *C.GObject,
	flags C.GConnectFlags,
	user_data C.gpointer) {

	builderSignals.Lock()
	signals, ok := builderSignals.m[builder]
	builderSignals.Unlock()

	if !ok {
		panic("no signal mapping defined for this GtkBuilder")
	}

	h := goString(handler_name)
	s := goString(signal_name)

	handler, ok := signals[h]
	if !ok {
		return
	}

	if object == nil {
		panic("unexpected nil object from builder")
	}

	//TODO: figure out a better way to get a glib.Object from a *C.GObject
	gobj := glib.ToObject(unsafe.Pointer(object))
	gobj.Connect(s, handler)
}

//export goPageSetupDone
func goPageSetupDone(setup *C.GtkPageSetup,
	data C.gpointer) {

	id := int(uintptr(data))

	pageSetupDoneCallbackRegistry.Lock()
	r := pageSetupDoneCallbackRegistry.m[id]
	delete(pageSetupDoneCallbackRegistry.m, id)
	pageSetupDoneCallbackRegistry.Unlock()

	obj := glib.Take(unsafe.Pointer(setup))
	r.fn(wrapPageSetup(obj), r.data)

}

//export goPrintSettings
func goPrintSettings(key *C.gchar,
	value *C.gchar,
	userData C.gpointer) {

	id := int(uintptr(userData))

	printSettingsCallbackRegistry.Lock()
	r := printSettingsCallbackRegistry.m[id]
	// TODO: figure out a way to determine when we can clean up
	//delete(printSettingsCallbackRegistry.m, id)
	printSettingsCallbackRegistry.Unlock()

	r.fn(goString(key), goString(value), r.userData)

}
