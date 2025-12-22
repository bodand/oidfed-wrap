package main

/*
#include "oidfed_wrap.h"
*/
import "C"
import (
	"runtime/cgo"
	"time"

	oidfed "github.com/go-oidfed/lib"
	"github.com/go-oidfed/lib/jwx"
)

//export oidfedFederationLeafCreateSimple
func oidfedFederationLeafCreateSimple(
	entityID *C.char,
	signer C.struct_oidfed_versatile_signer,
	errc *C.int,
) C.struct_oidfed_federation_leaf {
	goEntityID := C.GoString(entityID)
	vs := cgo.Handle(signer.impl).Value().(jwx.VersatileSigner)

	ess := jwx.NewEntityStatementSigner(vs)

	leaf, err := oidfed.NewFederationLeaf(
		goEntityID,
		nil,
		nil,
		nil,
		ess,
		time.Hour*24,
		vs,
		nil,
	)
	if err != nil {
		*errc = 1
		return C.struct_oidfed_federation_leaf{0}
	}
	return C.struct_oidfed_federation_leaf{packageGoThing(leaf)}
}

//export oidfedFederationLeafDestroy
func oidfedFederationLeafDestroy(leaf *C.struct_oidfed_federation_leaf) {
	if leaf.impl != 0 {
		cgo.Handle(leaf.impl).Delete()
	}
}

//export oidfedFederationLeafGetRequestObjectProducer
func oidfedFederationLeafGetRequestObjectProducer(
	leaf C.struct_oidfed_federation_leaf,
) C.struct_oidfed_request_producer {
	l := cgo.Handle(leaf.impl).Value().(*oidfed.FederationLeaf)
	rp := l.RequestObjectProducer()
	return C.struct_oidfed_request_producer{packageGoThing(rp)}
}
