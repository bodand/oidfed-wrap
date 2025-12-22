package main

/*
#include "oidfed_wrap.h"
*/
import "C"
import (
	"runtime/cgo"
	"unsafe"

	oidfed "github.com/go-oidfed/lib"
)

//export oidfedTrustResolverCreate
func oidfedTrustResolverCreate(
	subID *C.char,
	anchors *C.struct_oidfed_trust_anchor,
	anchorsCount C.size_t,
) C.struct_oidfed_trust_resolver {
	goSubID := C.GoString(subID)
	nativeSlice := unsafe.Slice(anchors, uintptr(anchorsCount))
	var goAnchors oidfed.TrustAnchors
	for _, ta := range nativeSlice {
		goAnchors = append(goAnchors, oidfed.TrustAnchor{
			EntityID: C.GoString(ta.entity_id),
		})
	}
	resolver := &oidfed.TrustResolver{
		StartingEntity: goSubID,
		TrustAnchors:   goAnchors,
	}
	return C.struct_oidfed_trust_resolver{packageGoThing(resolver)}
}

//export oidfedTrustResolverDestroy
func oidfedTrustResolverDestroy(r *C.struct_oidfed_trust_resolver) {
	if r.impl != 0 {
		cgo.Handle(r.impl).Delete()
	}
}

//export oidfedTrustResolverResolveToValidChains
func oidfedTrustResolverResolveToValidChains(
	r C.struct_oidfed_trust_resolver,
) C.struct_oidfed_trust_chains {
	resolver := cgo.Handle(r.impl).Value().(*oidfed.TrustResolver)
	chains := resolver.ResolveToValidChains()
	return C.struct_oidfed_trust_chains{packageGoThing(chains)}
}

//export oidfedTrustChainsCount
func oidfedTrustChainsCount(chains C.struct_oidfed_trust_chains) C.size_t {
	c := cgo.Handle(chains.impl).Value().(oidfed.TrustChains)
	return C.size_t(len(c))
}

//export oidfedTrustChainsGet
func oidfedTrustChainsGet(
	chains C.struct_oidfed_trust_chains,
	index C.size_t,
) C.struct_oidfed_trust_chain {
	c := cgo.Handle(chains.impl).Value().(oidfed.TrustChains)
	if int(index) >= len(c) {
		return C.struct_oidfed_trust_chain{0}
	}
	return C.struct_oidfed_trust_chain{packageGoThing(c[index])}
}

//export oidfedTrustChainsDestroy
func oidfedTrustChainsDestroy(chains *C.struct_oidfed_trust_chains) {
	if chains.impl != 0 {
		cgo.Handle(chains.impl).Delete()
	}
}

//export oidfedTrustChainDestroy
func oidfedTrustChainDestroy(chain *C.struct_oidfed_trust_chain) {
	if chain.impl != 0 {
		cgo.Handle(chain.impl).Delete()
	}
}

//export oidfedTrustChainGetMetadata
func oidfedTrustChainGetMetadata(
	chain C.struct_oidfed_trust_chain,
	errc *C.int,
) C.struct_oidfed_metadata {
	c := cgo.Handle(chain.impl).Value().(oidfed.TrustChain)
	m, err := c.Metadata()
	if err != nil {
		*errc = 1
		return C.struct_oidfed_metadata{0}
	}
	return C.struct_oidfed_metadata{packageGoThing(m)}
}

//export oidfedMetadataDestroy
func oidfedMetadataDestroy(m *C.struct_oidfed_metadata) {
	if m.impl != 0 {
		cgo.Handle(m.impl).Delete()
	}
}
