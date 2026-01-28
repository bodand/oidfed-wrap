package main

/*
#include "oidfed_wrap.h"
*/
import "C"
import (
	"unsafe"

	"github.com/go-oidfed/lib/jwx"
)

//export oidfedTrustAnchorDestroy
func oidfedTrustAnchorDestroy(ta *C.struct_oidfed_trust_anchor) {
	C.free(ta.jwks)
}

//export oidfedTrustAnchorCreate
func oidfedTrustAnchorCreate(id *C.char) C.struct_oidfed_trust_anchor {
	jwks := C.malloc(C.size_t(unsafe.Sizeof(jwx.JWKS{})))
	*(*jwx.JWKS)(jwks) = jwx.JWKS{}
	return C.struct_oidfed_trust_anchor{
		entity_id: id,
		jwks:      unsafe.Pointer(jwks),
	}
}
