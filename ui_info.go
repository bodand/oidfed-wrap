package main

/*
#include "oidfed_wrap.h"
*/
import "C"
import (
	"runtime/cgo"
	"unsafe"
)

//export oidfedUiInfoDestroy
func oidfedUiInfoDestroy(ui *C.struct_oidfed_ui_info) {
	C.free(unsafe.Pointer(ui.display_name))
	cgo.Handle(ui.impl).Delete()
}
