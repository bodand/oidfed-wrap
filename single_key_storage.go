package main

/*
#include "oidfed_wrap.h"
*/
import "C"
import (
	"crypto"
	"runtime/cgo"

	"github.com/go-oidfed/lib/jwx"
	"github.com/lestrrat-go/jwx/v3/jwa"
)

//export oidfedSingleKeyStorageCreate
func oidfedSingleKeyStorageCreate(
	signer C.struct_oidfed_signer,
	alg C.struct_oidfed_signature_algorithm,
) C.struct_oidfed_single_key_storage {
	signerImpl := cgo.Handle(signer.impl).Value().(crypto.Signer)
	signAlg := cgo.Handle(alg.impl).Value().(jwa.SignatureAlgorithm)
	sks := jwx.NewSingleKeyVersatileSigner(signerImpl, signAlg)
	return C.struct_oidfed_single_key_storage{
		packageGoThing(sks),
	}
}

//export oidfedSingleKeyStorageDestroy
func oidfedSingleKeyStorageDestroy(sks *C.struct_oidfed_single_key_storage) {
	cgo.Handle(sks.impl).Delete()
}
