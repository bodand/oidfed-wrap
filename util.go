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
	cStringSlice := unsafe.Slice(cStrings, uintptr(cStringCount))
	goStrings := make([]string, 0, cStringCount)
	for _, cstr := range cStringSlice {
		gt := C.GoString(cstr)
		goStrings = append(goStrings, gt)
	}
	return goStrings
}

func goifyCByteArray(bytes *C.char, bytesCount C.size_t) []byte {
	bytesSlice := unsafe.Slice(bytes, uintptr(bytesCount))
	goStrings := make([]byte, 0, bytesCount)
	for _, byt := range bytesSlice {
		goStrings = append(goStrings, byt.(byte))
	}
	return goStrings
}

func packageGoThing(v any) C.uintptr_t {
	h := cgo.NewHandle(v)
	return C.uintptr_t(h)
}
