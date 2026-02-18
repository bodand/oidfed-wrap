package main

/*
#include "oidfed_wrap.h"
*/
import "C"
import (
	"runtime/cgo"

	"github.com/lestrrat-go/jwx/v3/jws"
)

//export oidfedJWSHeadersCreate
func oidfedJWSHeadersCreate() C.struct_oidfed_jws_headers {
	h := jws.NewHeaders()
	return C.struct_oidfed_jws_headers{packageGoThing(h)}
}

//export oidfedJWSHeadersDestroy
func oidfedJWSHeadersDestroy(h *C.struct_oidfed_jws_headers) {
	if h.impl != 0 {
		cgo.Handle(h.impl).Delete()
	}
}

//export oidfedJWSHeadersSet
func oidfedJWSHeadersSet(h C.struct_oidfed_jws_headers, key *C.char, value *C.char) C.int {
	headers := cgo.Handle(h.impl).Value().(jws.Headers)
	goKey := C.GoString(key)
	goValue := C.GoString(value)
	if err := headers.Set(goKey, goValue); err != nil {
		return 1
	}
	return 0
}
