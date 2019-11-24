package glib

// #cgo pkg-config: glib-2.0 gobject-2.0
// #include <gio/gio.h>
// #include <glib.h>
// #include <glib-object.h>
// #include "glib.go.h"
import "C"
import "errors"

type BuildVersion int

// Implement stringer interface
func (v BuildVersion) String() string {
	switch v {
	case GLIB_2_40:
		return "2.40"
	case GLIB_2_42:
		return "2.42"
	case GLIB_2_44:
		return "2.44"
	case GLIB_2_46:
		return "2.46"
	case GLIB_2_48:
		return "2.48"
	case GLIB_2_50:
		return "2.50"
	case GLIB_2_52:
		return "2.52"
	case GLIB_2_54:
		return "2.54"
	case GLIB_2_56:
		return "2.56"
	case GLIB_2_58:
		return "2.58"
	case GLIB_2_60:
		return "2.60"
	case GLIB_2_62:
		return "2.62"
	default:
		return "< undefined >"
	}
}

const (
	GLIB_UNDEF BuildVersion = iota
	GLIB_2_40
	GLIB_2_42
	GLIB_2_44
	GLIB_2_46
	GLIB_2_48
	GLIB_2_50
	GLIB_2_52
	GLIB_2_54
	GLIB_2_56
	GLIB_2_58 // released in September 2018
	GLIB_2_60 // released in March 2019
	GLIB_2_62 // released in September 2019
)

// Save here version of GTK used to compile the library.
// Help to understand in runtime to what GTK API it correspond.
var buildVersion BuildVersion

func GetBuildVersion() BuildVersion {
	return buildVersion
}

func CheckVersion(major, minor, micro uint) error {
	errChar := C.glib_check_version(C.guint(major), C.guint(minor), C.guint(micro))
	if errChar == nil {
		return nil
	}

	return errors.New(goString(errChar))
}

func GetMajorVersion() uint {
	v := C.glib_major_version
	return uint(v)
}

func GetMinorVersion() uint {
	v := C.glib_minor_version
	return uint(v)
}

func GetMicroVersion() uint {
	v := C.glib_micro_version
	return uint(v)
}
