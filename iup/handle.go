package iup

/*
#cgo linux LDFLAGS: -liup
#include <iup.h>
#include <stdlib.h>
*/
import "C"

import (
	"github.com/gonutz/goiup/gen"
	"unsafe"
)

type Handle C.Ihandle

func (h *Handle) cptr() *C.Ihandle {
	return (*C.Ihandle)(h)
}

func (h *Handle) Map() int {
	return int(C.IupMap(h.cptr()))
}

func (h *Handle) Unmap() {
	C.IupUnmap(h.cptr())
}

func (h *Handle) Destroy() {
	C.IupDestroy(h.cptr())
}

func (h *Handle) Show() int {
	return int(C.IupShow(h.cptr()))
}

func (h *Handle) SetAttribute(name, value string) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	cValue := C.CString(value)
	defer C.free(unsafe.Pointer(cValue))
	C.IupSetAttribute(h.cptr(), cName, cValue)
}

func (h *Handle) GetAttribute(name string) string {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	return C.GoString(C.IupGetAttribute(h.cptr(), cName))
}

func (h *Handle) SetCallback(name string, cb interface{}) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	C.IupSetCallback(h.cptr(), cName, gen.GetNextCallback(cb))
}

func (h *Handle) GetClassName() string {
	return C.GoString(C.IupGetClassName(h.cptr()))
}

func (h *Handle) GetClassType() string {
	return C.GoString(C.IupGetClassType(h.cptr()))
}
