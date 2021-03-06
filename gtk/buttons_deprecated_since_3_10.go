// Copyright (c) 2013-2014 Conformal Systems <info@conformal.com>
//
// This file originated from: http://opensource.conformal.com/
//
// Permission to use, copy, modify, and distribute this software for any
// purpose with or without fee is hereby granted, provided that the above
// copyright notice and this permission notice appear in all copies.
//
// THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
// WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
// MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
// ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
// WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
// ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
// OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.

// This file includes wrapers for symbols deprecated beginning with GTK 3.10,
// and should only be included in a build targeted intended to target GTK
// 3.8 or earlier.  To target an earlier build build, use the build tag
// gtk_MAJOR_MINOR.  For example, to target GTK 3.8, run
// 'go build -tags gtk_3_8'.
// +build gtk_3_6 gtk_3_8

package gtk

// #cgo pkg-config: gtk+-3.0
// #include <stdlib.h>
// #include <gtk/gtk.h>
import "C"
import (
	"unsafe"

	"github.com/romychs/gotk3/glib"
)

/*
 * GtkButton
 */

// ButtonNewFromStock is a wrapper around gtk_button_new_from_stock().
func ButtonNewFromStock(stock Stock) (*Button, error) {
	cstr := C.CString(string(stock))
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_button_new_from_stock((*C.gchar)(cstr))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapButton(obj), nil
}

// SetUseStock is a wrapper around gtk_button_set_use_stock().
func (v *Button) SetUseStock(useStock bool) {
	C.gtk_button_set_use_stock(v.native(), gbool(useStock))
}

// GetUseStock is a wrapper around gtk_button_get_use_stock().
func (v *Button) GetUseStock() bool {
	c := C.gtk_button_get_use_stock(v.native())
	return gobool(c)
}
