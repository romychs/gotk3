package gtk

// #cgo pkg-config: gtk+-3.0
// #include <gtk/gtk.h>
import "C"
import "errors"

type BuildVersion int

// Implement stringer interface
func (v BuildVersion) String() string {
	switch v {
	case GTK_3_6:
		return "3.6"
	case GTK_3_8:
		return "3.8"
	case GTK_3_10:
		return "3.10"
	case GTK_3_12:
		return "3.12"
	case GTK_3_14:
		return "3.14"
	case GTK_3_16:
		return "3.16"
	case GTK_3_18:
		return "3.18"
	case GTK_3_20:
		return "3.20"
	case GTK_3_22:
		return "3.22"
	default:
		return "< undefined >"
	}
}

const (
	GTK_UNDEF BuildVersion = iota
	GTK_3_6
	GTK_3_8
	GTK_3_10
	GTK_3_12
	GTK_3_14
	GTK_3_16
	GTK_3_18
	GTK_3_20
	GTK_3_22
)

// Save here version of GTK used to compile the library.
// Help to understand in runtime to what GTK API it correspond.
var buildVersion BuildVersion

func GetBuildVersion() BuildVersion {
	return buildVersion
}

func CheckVersion(major, minor, micro uint) error {
	errChar := C.gtk_check_version(C.guint(major), C.guint(minor), C.guint(micro))
	if errChar == nil {
		return nil
	}

	return errors.New(goString(errChar))
}

func GetMajorVersion() uint {
	v := C.gtk_get_major_version()
	return uint(v)
}

func GetMinorVersion() uint {
	v := C.gtk_get_minor_version()
	return uint(v)
}

func GetMicroVersion() uint {
	v := C.gtk_get_micro_version()
	return uint(v)
}
