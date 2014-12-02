package iup

/*
#include <iup.h>
#include <stdlib.h>
*/
import "C"

import (
	"unsafe"
)

func Menu(children ...*Handle) *Handle {
	return (*Handle)(C.IupMenuv(makeHandlePointerArray(children)))
}

func Submenu(title string, menu *Handle) *Handle {
	cTitle := C.CString(title)
	defer C.free(unsafe.Pointer(cTitle))
	return (*Handle)(C.IupSubmenu(cTitle, menu.cptr()))
}

func Item(title, action string) *Handle {
	cTitle := C.CString(title)
	defer C.free(unsafe.Pointer(cTitle))
	cAction := C.CString(action)
	defer C.free(unsafe.Pointer(cAction))
	return (*Handle)(C.IupItem(cTitle, cAction))
}
