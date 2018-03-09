package glib

/*
// #cgo pkg-config: glib-2.0 gobject-2.0
// #include <gio/gio.h>
// #include <glib.h>
// #include <glib-object.h>
import "C"
import "unsafe"
import "errors"
import "fmt"
//import "log"
import "runtime"



// Abstract respresentation for GTK+ GBoxed object.
type Boxed struct {
	GBoxed C.gpointer
}

func BoxedNew(ptr unsafe.Pointer) Boxed {
	v := Boxed{(C.gpointer)(ptr)}
	return v
}

func BoxedFromObjectNew(obj *Object) Boxed {
	p := unsafe.Pointer(obj.GObject)
	return BoxedNew(p)
}

func (v *Boxed) ToPtr() unsafe.Pointer {
	if v == nil {
		return nil
	}
	return unsafe.Pointer(v.GBoxed)
}

func (v *Boxed) ToUintptr() uintptr {
	if v == nil {
		return 0
	}
	return uintptr(unsafe.Pointer(v.GBoxed))
}

type Boxed2 struct {
	GBoxed uintptr
}

func ToBoxed2(ptr unsafe.Pointer) *Boxed2 {
	boxed := &Boxed2{uintptr(ptr)}
	return boxed
}

func (v *Boxed2) ToPtr() unsafe.Pointer {
	if v == nil {
		return nil
	}
	return unsafe.Pointer(v.GBoxed)
}

func (v *Boxed2) nativePtr() C.gpointer {
	if v == nil {
		return nil
	}
	return C.gpointer(v.GBoxed)
}

func (v *Boxed2) nativeConstPtr() C.gconstpointer {
	if v == nil {
		return nil
	}
	return C.gconstpointer(v.GBoxed)
}

func (v *Boxed2) Copy() (*Boxed2, error) {
	val := &Value{(*C.GValue)(unsafe.Pointer(v.GBoxed))}
	t1, t2, err := val.Type()
	if err != nil {
		return nil, err
	}
	if t2 != TYPE_BOXED {
		return nil, errors.New(fmt.Sprintf("fundamental type is not TYPE_BOXED: $q", t1.Name()))
	}
	c := C.g_boxed_copy(C.GType(t1), v.nativeConstPtr())
	if c == nil {
		return nil, errNilPtr
	}

	boxed := &Boxed2{uintptr(c)}
	runtime.SetFinalizer(boxed, (*Boxed2).Free)

	return boxed, nil
}

func (v* Boxed2) Free() error {
	//log.Println("object of type is freeeing")
	val := &Value{(*C.GValue)(unsafe.Pointer(v.GBoxed))}
	t1, t2, err := val.Type()
	if err != nil {
		return err
	}
	//log.Println(fmt.Sprintf("object of type $q is freeeing", t1.Name()))
	if t2 != TYPE_BOXED {
		return errors.New(fmt.Sprintf("fundamental type is not TYPE_BOXED: $q", t1.Name()))
	}
	C.g_boxed_free(C.GType(t1), v.nativePtr())
	//log.Println(fmt.Sprintf("object of type $q is freed", t1.Name()))
	return nil
}
*/
