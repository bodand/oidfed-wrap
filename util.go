package main

/*
#include "stdint.h"
*/
import "C"
import (
	"runtime/cgo"
	"unsafe"
)

func goifyCArray(cStrings **C.char, cStringCount C.size_t) []string {
	grantsSlice := unsafe.Slice(cStrings, uintptr(cStringCount))
	var goStrings []string
	for _, cstr := range grantsSlice {
		gt := C.GoString(cstr)
		goStrings = append(goStrings, gt)
	}
	return goStrings
}

func packageGoThing(v any) C.uintptr_t {
	h := cgo.NewHandle(v)
	return C.uintptr_t(h)
}
