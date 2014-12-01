package iup

/*
#cgo linux LDFLAGS: -liup
#cgo windows LDFLAGS: -liup -luuid -lgdi32 -lole32 -lcomctl32 -lcomdlg32
#include <iup.h>
#include <stdlib.h>
*/
import "C"

// not working
// #cgo windows CFLAGS: -fno-stack-check -fno-stack-protector -mno-stack-arg-probe
// -lmingwex -lmingw32

import (
	"unsafe"
)

const (
	OPENED  = C.IUP_OPENED
	ERROR   = C.IUP_ERROR
	NOERROR = C.IUP_NOERROR

	DEFAULT  = C.IUP_DEFAULT
	CLOSE    = C.IUP_CLOSE
	IGNORE   = C.IUP_IGNORE
	CONTINUE = C.IUP_CONTINUE
)

// TODO convert os.Args to int* and char** and pass them to IupOpen
// then place the (possibly modified) args back into os.Args
func Open() int {
	return int(C.IupOpen(nil, nil))
}

func Close() {
	C.IupClose()
}

func MainLoop() int {
	return int(C.IupMainLoop())
}

func Message(title, msg string) {
	cTitle := C.CString(title)
	defer C.free(unsafe.Pointer(cTitle))
	cMsg := C.CString(msg)
	defer C.free(unsafe.Pointer(cMsg))
	C.IupMessage(cTitle, cMsg)
}

func SetLanguage(lang string) {
	cLang := C.CString(lang)
	defer C.free(unsafe.Pointer(cLang))
	C.IupSetLanguage(cLang)
}

func GetGlobal(name string) string {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	return C.GoString(C.IupGetGlobal(cName))
}
