package main

/*
#include "oidfed_wrap.h"
*/
import "C"
import (
	"runtime/cgo"
	"time"
	"unsafe"

	oidfed "github.com/go-oidfed/lib"
	"github.com/go-oidfed/lib/jwx"
)

//export oidfedFederationLeafCreate
func oidfedFederationLeafCreate(
	entityID *C.char,
	authorityHints **C.char,
	authorityHintsCount C.size_t,
	trustAnchors *C.struct_oidfed_trust_anchor,
	trustAnchorsCount C.size_t,
	federationSigner C.struct_oidfed_versatile_signer,
	oidcSigner C.struct_oidfed_versatile_signer,
	metadata C.struct_oidfed_metadata,
	errc *C.int,
) C.struct_oidfed_federation_leaf {
	goEntityID := C.GoString(entityID)

	goHints := goifyCArray(authorityHints, authorityHintsCount)
	goAnchors := unsafe.Slice(trustAnchors, uintptr(trustAnchorsCount))
	anchors := make(oidfed.TrustAnchors, trustAnchorsCount)
	for i, a := range goAnchors {
		anchors[i] = oidfed.TrustAnchor{
			EntityID: C.GoString(a.entity_id),
			JWKS:     *(*jwx.JWKS)(a.jwks),
		}
	}

	federationVs := cgo.Handle(federationSigner.impl).Value().(jwx.VersatileSigner)
	oidcVs := cgo.Handle(oidcSigner.impl).Value().(jwx.VersatileSigner)

	metadataImpl := cgo.Handle(metadata.impl).Value().(oidfed.Metadata)

	ess := jwx.NewEntityStatementSigner(federationVs)

	leaf, err := oidfed.NewFederationLeaf(
		goEntityID,
		goHints,
		anchors,
		&metadataImpl,
		ess,
		time.Hour*24,
		oidcVs,
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
