package iup

/*
#cgo linux LDFLAGS: -liup
#cgo windows LDFLAGS: -liup -liupstub -luuid -lgdi32 -lole32 -lcomctl32 -lcomdlg32
#include <iup.h>
#include <stdlib.h>
*/
import "C"

import (
	"errors"
	"strings"
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

	// these are additional constants to give common IUP return values names

	// these are the STATUS values after using a file dialog
	NewFile                 = 1
	ExistingFileOrDirectory = 0
	FileOperationCancelled  = -1
)

// TODO convert os.Args to int* and char** and pass them to IupOpen
// then place the (possibly modified) args back into os.Args
func Open() error {
	result := int(C.IupOpen(nil, nil))
	if result == ERROR {
		return errors.New("IupOpen failed.")
	}
	return nil
}

func Close() { C.IupClose() }

//func ImageLibOpen() { C.IupImageLibOpen() } // TODO undefined ref
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

func LoadBuffer(buffer string) error {
	cBuffer := C.CString(buffer)
	defer C.free(unsafe.Pointer(cBuffer))
	result := C.GoString(C.IupLoadBuffer(cBuffer))
	if result == "" {
		return nil
	}
	return errors.New(result)
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

func Alarm1(title, msg, button string) int {
	cTitle := C.CString(title)
	defer C.free(unsafe.Pointer(cTitle))
	cMsg := C.CString(msg)
	defer C.free(unsafe.Pointer(cMsg))
	cButton := C.CString(button)
	defer C.free(unsafe.Pointer(cButton))
	return int(C.IupAlarm(cTitle, cMsg, cButton, nil, nil))
}

func Alarm2(title, msg, button1, button2 string) int {
	cTitle := C.CString(title)
	defer C.free(unsafe.Pointer(cTitle))
	cMsg := C.CString(msg)
	defer C.free(unsafe.Pointer(cMsg))
	cButton1 := C.CString(button1)
	defer C.free(unsafe.Pointer(cButton1))
	cButton2 := C.CString(button2)
	defer C.free(unsafe.Pointer(cButton2))
	return int(C.IupAlarm(cTitle, cMsg, cButton1, cButton2, nil))
}

func Alarm3(title, msg, button1, button2, button3 string) int {
	cTitle := C.CString(title)
	defer C.free(unsafe.Pointer(cTitle))
	cMsg := C.CString(msg)
	defer C.free(unsafe.Pointer(cMsg))
	cButton1 := C.CString(button1)
	defer C.free(unsafe.Pointer(cButton1))
	cButton2 := C.CString(button2)
	defer C.free(unsafe.Pointer(cButton2))
	cButton3 := C.CString(button3)
	defer C.free(unsafe.Pointer(cButton3))
	return int(C.IupAlarm(cTitle, cMsg, cButton1, cButton2, cButton3))
}

func GetFile(defaultFolderAndFilter string) (status int, filename string) {
	const MaxSize = 4096
	// TODO return error instead of panicing or document that this panics
	if len(defaultFolderAndFilter) > MaxSize {
		panic("string too long (see IUP documentation for IupGetFile)")
	}
	fillers := strings.Repeat(" ", MaxSize-1-len(defaultFolderAndFilter))
	var cStringDelimiter byte = 0
	cPath := C.CString(defaultFolderAndFilter + string(cStringDelimiter) + fillers)
	defer C.free(unsafe.Pointer(cPath))
	status = int(C.IupGetFile(cPath))
	filename = C.GoString(cPath)
	return
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
