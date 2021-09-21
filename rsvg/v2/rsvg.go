// Package rsvg loads svg files and draws them on cairo surface.
//
// https://developer-old.gnome.org/rsvg/2.40/
package rsvg

/*
#cgo pkg-config: librsvg-2.0
#include <librsvg/rsvg.h>
*/
import "C"
import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/cairo"
	"github.com/diamondburned/gotk4/pkg/core/gerror"
)

// RSVG is a representation of RSVG's RsvgHandle.
type RSVG struct{ ptr *C.RsvgHandle }

// NewFromFile loads the SVG specified by file_name.
//
// filename can be a URI if built with gnome-vfs.
func NewFromFile(filename string) (*RSVG, error) {
	cstr := (*C.char)(unsafe.Pointer(C.CString(filename)))
	defer C.free(unsafe.Pointer(cstr))

	var cerr *C.GError
	ptr := C.rsvg_handle_new_from_file(cstr, &cerr)

	if cerr != nil {
		return nil, gerror.Take(unsafe.Pointer(cerr))
	}
	return &RSVG{ptr}, nil
}

// Dimensions returns the SVG's size.
//
// Do not call from within the size_func callback, because an infinite loop will occur.
func (v *RSVG) Dimensions() (x, y int, em, ex float64) {
	var dim C.RsvgDimensionData
	C.rsvg_handle_get_dimensions(v.ptr, &dim)
	return int(dim.width), int(dim.height), float64(dim.em), float64(dim.ex)
}

// RenderCairo draws a SVG to a Cairo surface.
func (v *RSVG) RenderCairo(cr *cairo.Context) bool {
	context := (*C.cairo_t)(unsafe.Pointer(cr.Native()))
	cbool := C.rsvg_handle_render_cairo(v.ptr, context)
	return cbool > 0
}
