package iup

/*
#cgo linux LDFLAGS: -liup
#cgo windows LDFLAGS: -liup -liupstub -luuid -lgdi32 -lole32 -lcomctl32 -lcomdlg32
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

	CENTER       = C.IUP_CENTER
	LEFT         = C.IUP_LEFT
	TOP          = C.IUP_TOP
	BOTTOM       = C.IUP_BOTTOM
	RIGHT        = C.IUP_RIGHT
	MOUSEPOS     = C.IUP_MOUSEPOS
	CENTERPARENT = C.IUP_CENTERPARENT
	CURRENT      = C.IUP_CURRENT
)

// TODO convert os.Args to int* and char** and pass them to IupOpen
// then place the (possibly modified) args back into os.Args
func Open() int { return int(C.IupOpen(nil, nil)) }
func Close()    { C.IupClose() }

//func ImageLibOpen()      { C.IupImageLibOpen() } // TODO undefined ref
func MainLoop() int      { return int(C.IupMainLoop()) }
func LoopStep() int      { return int(C.IupLoopStep()) }
func LoopStepWait() int  { return int(C.IupLoopStepWait()) }
func MainLoopLevel() int { return int(C.IupMainLoopLevel()) }
func Flush()             { C.IupFlush() }
func ExitLoop()          { C.IupExitLoop() }

func RecordInput(filename string, mode int) int {
	cFilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cFilename))
	return int(C.IupRecordInput(cFilename, C.int(mode)))
}

func PlayInput(filename string) int {
	cFilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cFilename))
	return int(C.IupPlayInput(cFilename))
}

func Help(url string) int {
	cUrl := C.CString(url)
	defer C.free(unsafe.Pointer(cUrl))
	return int(C.IupHelp(cUrl))
}

func Load(filename string) string {
	cFilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cFilename))
	return C.GoString(C.IupLoad(cFilename))
}

func LoadBuffer(buffer string) string {
	cBuffer := C.CString(buffer)
	defer C.free(unsafe.Pointer(cBuffer))
	return C.GoString(C.IupLoadBuffer(cBuffer))
}

func Version() string     { return C.GoString(C.IupVersion()) }
func VersionDate() string { return C.GoString(C.IupVersionDate()) }
func VersionNumber() int  { return int(C.IupVersionNumber()) }

func SetLanguage(lang string) {
	cLang := C.CString(lang)
	defer C.free(unsafe.Pointer(cLang))
	C.IupSetLanguage(cLang)
}

func GetLanguage() string { return C.GoString(C.IupGetLanguage()) }

func SetLanguageString(name, str string) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	cStr := C.CString(str)
	defer C.free(unsafe.Pointer(cStr))
	C.IupSetLanguageString(cName, cStr)
}

func StoreLanguageString(name, str string) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	cStr := C.CString(str)
	defer C.free(unsafe.Pointer(cStr))
	C.IupStoreLanguageString(cName, cStr)
}

func GetLanguageString(name string) string {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	return C.GoString(C.IupGetLanguageString(cName))
}

func Message(title, msg string) {
	cTitle := C.CString(title)
	defer C.free(unsafe.Pointer(cTitle))
	cMsg := C.CString(msg)
	defer C.free(unsafe.Pointer(cMsg))
	C.IupMessage(cTitle, cMsg)
}

func GetGlobal(name string) string {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	return C.GoString(C.IupGetGlobal(cName))
}

func SetGlobal(name, value string) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	cValue := C.CString(value)
	defer C.free(unsafe.Pointer(cValue))
	C.IupSetStrGlobal(cName, cValue)
}

// SetHandle associates a name with an interface element and returns the previously
// associated element.
func SetHandle(name string, h *Handle) *Handle {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	return (*Handle)(C.IupSetHandle(cName, h.cptr()))
}

// GetHandle returns the interface element associated with the given name (per
// SetHandle).
func GetHandle(name string) *Handle {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	return (*Handle)(C.IupGetHandle(cName))
}

func GetFocus() *Handle {
	return (*Handle)(C.IupGetFocus())
}
