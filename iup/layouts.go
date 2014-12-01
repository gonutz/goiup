package iup

/*
#cgo linux LDFLAGS: -liup
#include <iup.h>
*/
import "C"

func Cbox(children ...*Handle) *Handle {
	return (*Handle)(C.IupCboxv(makeHandlePointerArray(children)))
}

func Fill() *Handle {
	return (*Handle)(C.IupFill())
}

func GridBox(children ...*Handle) *Handle {
	return (*Handle)(C.IupGridBoxv(makeHandlePointerArray(children)))
}

func Hbox(children ...*Handle) *Handle {
	return (*Handle)(C.IupHboxv(makeHandlePointerArray(children)))
}

func Vbox(children ...*Handle) *Handle {
	return (*Handle)(C.IupVboxv(makeHandlePointerArray(children)))
}

func Zbox(children ...*Handle) *Handle {
	return (*Handle)(C.IupZboxv(makeHandlePointerArray(children)))
}

func Radio(child *Handle) *Handle {
	return (*Handle)(C.IupRadio(child.cptr()))
}

func Normalizer(list ...*Handle) *Handle {
	return (*Handle)(C.IupNormalizerv(makeHandlePointerArray(list)))
}

func BackgroundBox(child *Handle) *Handle {
	return (*Handle)(C.IupBackgroundBox(child.cptr()))
}

func DetachBox(child *Handle) *Handle {
	return (*Handle)(C.IupDetachBox(child.cptr()))
}

func Expander(child *Handle) *Handle {
	return (*Handle)(C.IupExpander(child.cptr()))
}

func Sbox(child *Handle) *Handle {
	return (*Handle)(C.IupSbox(child.cptr()))
}

func ScrollBox(child *Handle) *Handle {
	return (*Handle)(C.IupScrollBox(child.cptr()))
}

func Split(topOrLeft, bottomOrRight *Handle) *Handle {
	return (*Handle)(C.IupSplit(topOrLeft.cptr(), bottomOrRight.cptr()))
}
