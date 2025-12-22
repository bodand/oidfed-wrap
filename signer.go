package main

/*
#include "oidfed_wrap.h"
*/
import "C"
import (
	"crypto/x509"
	"encoding/pem"
	"runtime/cgo"
)

//export oidfedSignerCreateFromPEM
func oidfedSignerCreateFromPEM(
	pemBytes *C.char,
	pemCount C.size_t,
	errc *C.int,
) C.struct_oidfed_signer {
	pemSlice := goifyCByteArray(pemBytes, pemCount)
	pemData, _ := pem.Decode(pemSlice)
	if pemData == nil {
		*errc = 1 // OidfedErrInvalidPEM
		return C.struct_oidfed_signer{0}
	}

	key, err := x509.ParsePKCS8PrivateKey(pemData.Bytes)
	if err != nil {
		*errc = 2 // OidfedErrInvalidPrivateKey
		return C.struct_oidfed_signer{0}
	}

	return C.struct_oidfed_signer{packageGoThing(key)}
}

//export oidfedSignerDestroy
func oidfedSignerDestroy(s *C.struct_oidfed_signer) {
	if s.impl != 0 {
		cgo.Handle(s.impl).Delete()
	}
}
