package main

/*
#include "oidfed_wrap.h"
*/
import "C"
import (
	"runtime/cgo"
	"unsafe"

	"github.com/lestrrat-go/jwx/v3/jwa"
)

//export oidfedSignatureAlgorithmCreateEmpty
func oidfedSignatureAlgorithmCreateEmpty() C.struct_oidfed_signature_algorithm {
	return C.struct_oidfed_signature_algorithm{nil, false, false, C.uintptr_t(0)}
}

//export oidfedSignatureAlgorithmDestroy
func oidfedSignatureAlgorithmDestroy(sa *C.struct_oidfed_signature_algorithm) {
	C.free(unsafe.Pointer(sa.name))
	cgo.Handle(sa.impl).Delete()
}

//export oidfedSignatureAlgorithmGet
func oidfedSignatureAlgorithmGet(name *C.char, succ *C.bool) C.struct_oidfed_signature_algorithm {
	goName := C.GoString(name)
	signAlg, ok := jwa.LookupSignatureAlgorithm(goName)
	*succ = C.bool(ok)

	if !ok {
		return oidfedSignatureAlgorithmCreateEmpty()
	}

	cname := C.CString(signAlg.String())
	deprecation := C.bool(signAlg.IsDeprecated())
	symmetric := C.bool(signAlg.IsSymmetric())
	return C.struct_oidfed_signature_algorithm{
		cname,
		deprecation,
		symmetric,
		packageGoThing(signAlg),
	}
}
