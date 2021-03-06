package iup

/*
#include <iup.h>
#include <stdlib.h>

static inline char** allocateStringArrayOfSize(int size) {
	return (char**) malloc(size * sizeof(char**));
}

static inline char* getStringElementAtIndex(char** array, int index) {
	return array[index];
}
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

func (h *Handle) Update()             { C.IupUpdate(h.cptr()) }
func (h *Handle) UpdateChildren()     { C.IupUpdateChildren(h.cptr()) }
func (h *Handle) Redraw(children int) { C.IupRedraw(h.cptr(), C.int(children)) }
func (h *Handle) Refresh()            { C.IupRefresh(h.cptr()) }
func (h *Handle) RefreshChildren()    { C.IupRefreshChildren(h.cptr()) }
func (h *Handle) SetLanguagePack()    { C.IupSetLanguagePack(h.cptr()) }
func (h *Handle) Destroy()            { C.IupDestroy(h.cptr()) }
func (h *Handle) Detach()             { C.IupDetach(h.cptr()) }

func (h *Handle) Append(child *Handle) *Handle {
	return (*Handle)(C.IupAppend(h.cptr(), child.cptr()))
}

func (h *Handle) Insert(refChild, child *Handle) *Handle {
	return (*Handle)(C.IupInsert(h.cptr(), refChild.cptr(), child.cptr()))
}

func (h *Handle) GetChild(position int) *Handle {
	return (*Handle)(C.IupGetChild(h.cptr(), C.int(position)))
}

func (h *Handle) GetChildPos(child *Handle) int {
	return int(C.IupGetChildPos(h.cptr(), child.cptr()))
}

func (h *Handle) GetChildCount() int {
	return int(C.IupGetChildCount(h.cptr()))
}

func (h *Handle) GetNextChild(child *Handle) *Handle {
	return (*Handle)(C.IupGetNextChild(h.cptr(), child.cptr()))
}

func (h *Handle) GetBrother() *Handle {
	return (*Handle)(C.IupGetBrother(h.cptr()))
}

func (h *Handle) GetParent() *Handle {
	return (*Handle)(C.IupGetParent(h.cptr()))
}

func (h *Handle) GetDialog() *Handle {
	return (*Handle)(C.IupGetDialog(h.cptr()))
}

func (h *Handle) GetDialogChild(name string) *Handle {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	return (*Handle)(C.IupGetDialogChild(h.cptr(), cName))
}

func (h *Handle) Reparent(newParent, refChild *Handle) int {
	return int(C.IupReparent(h.cptr(), newParent.cptr(), refChild.cptr()))
}

func (h *Handle) Popup(x, y int) int {
	return int(C.IupPopup(h.cptr(), C.int(x), C.int(y)))
}

func (h *Handle) Show() int {
	return int(C.IupShow(h.cptr()))
}

func (h *Handle) ShowXY(x, y int) int {
	return int(C.IupShowXY(h.cptr(), C.int(x), C.int(y)))
}

func (h *Handle) Hide() int { return int(C.IupHide(h.cptr())) }
func (h *Handle) Map() int  { return int(C.IupMap(h.cptr())) }
func (h *Handle) Unmap()    { C.IupUnmap(h.cptr()) }

func (h *Handle) ResetAttribute(name string) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	C.IupResetAttribute(h.cptr(), cName)
}

func (h *Handle) GetAllAttributes() (names []string) {
	maxNumAttributes := int(C.IupGetAllAttributes(h.cptr(), nil, 0))
	cNames := C.allocateStringArrayOfSize(C.int(maxNumAttributes))
	defer C.free(unsafe.Pointer(cNames))
	numAttributes := int(C.IupGetAllAttributes(h.cptr(), cNames, C.int(maxNumAttributes)))
	names = make([]string, numAttributes)
	for i := range names {
		names[i] = C.GoString(C.getStringElementAtIndex(cNames, C.int(i)))
	}
	return
}

func (h *Handle) SetAttributes(str string) {
	cStr := C.CString(str)
	defer C.free(unsafe.Pointer(cStr))
	C.IupSetAttributes(h.cptr(), cStr)
}

func (h *Handle) GetAttributes() string {
	return C.GoString(C.IupGetAttributes(h.cptr()))
}

func (h *Handle) SetAttribute(name, value string) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	cValue := C.CString(value)
	defer C.free(unsafe.Pointer(cValue))
	C.IupSetStrAttribute(h.cptr(), cName, cValue)
}

func (h *Handle) SetInt(name string, value int) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	C.IupSetInt(h.cptr(), cName, C.int(value))
}

func (h *Handle) SetFloat(name string, value float32) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	C.IupSetFloat(h.cptr(), cName, C.float(value))
}

func (h *Handle) SetDouble(name string, value float64) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	C.IupSetDouble(h.cptr(), cName, C.double(value))
}

func (h *Handle) SetRGB(name string, r, g, b byte) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	C.IupSetRGB(h.cptr(), cName, C.uchar(r), C.uchar(g), C.uchar(b))
}

func (h *Handle) GetAttribute(name string) string {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	return C.GoString(C.IupGetAttribute(h.cptr(), cName))
}

func (h *Handle) GetInt(name string) int {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	return int(C.IupGetInt(h.cptr(), cName))
}

func (h *Handle) GetInt2(name string) int {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	return int(C.IupGetInt2(h.cptr(), cName))
}

func (h *Handle) GetIntInt(name string) (val1, val2, numVals int) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	var i1, i2 C.int
	numVals = int(C.IupGetIntInt(
		h.cptr(),
		cName,
		(*C.int)(unsafe.Pointer(&i1)),
		(*C.int)(unsafe.Pointer(&i2))))
	val1 = int(i1)
	val2 = int(i2)
	return
}

func (h *Handle) GetFloat(name string) float32 {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	return float32(C.IupGetFloat(h.cptr(), cName))
}

func (h *Handle) GetDouble(name string) float64 {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	return float64(C.IupGetDouble(h.cptr(), cName))
}

func (h *Handle) GetRGB(name string) (r, g, b byte) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	var cR, cG, cB C.uchar
	C.IupGetRGB(
		h.cptr(),
		cName,
		(*C.uchar)(unsafe.Pointer(&cR)),
		(*C.uchar)(unsafe.Pointer(&cG)),
		(*C.uchar)(unsafe.Pointer(&cB)))
	return byte(cR), byte(cG), byte(cB)
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

func (h *Handle) SetFocus() *Handle {
	return (*Handle)(C.IupSetFocus(h.cptr()))
}

func (h *Handle) SetAttributeId(name string, id int, value string) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	cValue := C.CString(value)
	defer C.free(unsafe.Pointer(cValue))
	C.IupSetStrAttributeId(h.cptr(), cName, C.int(id), cValue)
}

func (h *Handle) SetIntId(name string, id, value int) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	C.IupSetIntId(h.cptr(), cName, C.int(id), C.int(value))
}

func (h *Handle) SetFloatId(name string, id int, value float32) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	C.IupSetFloatId(h.cptr(), cName, C.int(id), C.float(value))
}

func (h *Handle) SetDoubleId(name string, id int, value float64) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	C.IupSetDoubleId(h.cptr(), cName, C.int(id), C.double(value))
}

func (h *Handle) SetRGBId(name string, id int, r, g, b byte) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	C.IupSetRGBId(h.cptr(), cName, C.int(id), C.uchar(r), C.uchar(g), C.uchar(b))
}

func (h *Handle) GetAttributeId(name string, id int) string {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	return C.GoString(C.IupGetAttributeId(h.cptr(), cName, C.int(id)))
}

func (h *Handle) GetIntId(name string, id int) int {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	return int(C.IupGetIntId(h.cptr(), cName, C.int(id)))
}

func (h *Handle) GetFloatId(name string, id int) float32 {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	return float32(C.IupGetFloatId(h.cptr(), cName, C.int(id)))
}

func (h *Handle) GetDoubleId(name string, id int) float64 {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	return float64(C.IupGetDoubleId(h.cptr(), cName, C.int(id)))
}

func (h *Handle) GetRGBId(name string, id int) (r, g, b byte) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	var cR, cG, cB C.uchar
	C.IupGetRGBId(
		h.cptr(),
		cName,
		C.int(id),
		(*C.uchar)(unsafe.Pointer(&cR)),
		(*C.uchar)(unsafe.Pointer(&cG)),
		(*C.uchar)(unsafe.Pointer(&cB)))
	return byte(cR), byte(cG), byte(cB)
}
