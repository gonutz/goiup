package iup

/*
#cgo linux LDFLAGS: -liup
#include <iup.h>
#include <stdlib.h>
*/
import "C"

import (
	"fmt"
	"unsafe"
)

func Create(classname string) *Handle {
	cClassname := C.CString(classname)
	defer C.free(unsafe.Pointer(cClassname))
	return (*Handle)(C.IupCreate(cClassname))
}

func Button(title, action string) *Handle {
	cTitle := C.CString(title)
	defer C.free(unsafe.Pointer(cTitle))
	cAction := C.CString(action)
	defer C.free(unsafe.Pointer(cAction))
	return (*Handle)(C.IupButton(cTitle, cAction))
}

func Canvas(action string) *Handle {
	cAction := C.CString(action)
	defer C.free(unsafe.Pointer(cAction))
	return (*Handle)(C.IupCanvas(cAction))
}

func Frame(child *Handle) *Handle {
	return (*Handle)(C.IupFrame(child.cptr()))
}

func Label(title string) *Handle {
	cTitle := C.CString(title)
	defer C.free(unsafe.Pointer(cTitle))
	return (*Handle)(C.IupLabel(cTitle))
}

func Link(url, title string) *Handle {
	cUrl := C.CString(url)
	defer C.free(unsafe.Pointer(cUrl))
	cTitle := C.CString(title)
	defer C.free(unsafe.Pointer(cTitle))
	return (*Handle)(C.IupLink(cUrl, cTitle))
}

func List(action string) *Handle {
	cAction := C.CString(action)
	defer C.free(unsafe.Pointer(cAction))
	return (*Handle)(C.IupList(cAction))
}

func MultiLine(action string) *Handle {
	cAction := C.CString(action)
	defer C.free(unsafe.Pointer(cAction))
	return (*Handle)(C.IupMultiLine(cAction))
}

func ProgressBar() *Handle {
	return (*Handle)(C.IupProgressBar())
}

func Spin() *Handle {
	return (*Handle)(C.IupSpin())
}

func Tabs(children ...*Handle) *Handle {
	return (*Handle)(C.IupTabsv(makeHandlePointerArray(children)))
}

func makeHandlePointerArray(handles []*Handle) **C.Ihandle {
	const maxHandles = 63
	if len(handles) > maxHandles {
		panic(fmt.Sprintf("too many handles for array, max is %v but were %v",
			maxHandles, len(handles)))
	}
	var cHandles [maxHandles + 1]*C.Ihandle
	for i, h := range handles {
		cHandles[i] = h.cptr()
	}
	cHandles[len(handles)] = nil
	return (**C.Ihandle)(unsafe.Pointer(&cHandles))
}

func Text(action string) *Handle {
	cAction := C.CString(action)
	defer C.free(unsafe.Pointer(cAction))
	return (*Handle)(C.IupText(cAction))
}

func Toggle(title, action string) *Handle {
	cTitle := C.CString(title)
	defer C.free(unsafe.Pointer(cTitle))
	cAction := C.CString(action)
	defer C.free(unsafe.Pointer(cAction))
	return (*Handle)(C.IupToggle(cTitle, cAction))
}

func Tree() *Handle {
	return (*Handle)(C.IupTree())
}

func Val(orientation string) *Handle {
	cOrientation := C.CString(orientation)
	defer C.free(unsafe.Pointer(cOrientation))
	return (*Handle)(C.IupVal(cOrientation))
}

func Dialog(child *Handle) *Handle {
	return (*Handle)(C.IupDialog(child.cptr()))
}
