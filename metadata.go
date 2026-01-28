package main

/*
#include "oidfed_wrap.h"
*/
import "C"
import (
	"runtime/cgo"

	oidfed "github.com/go-oidfed/lib"
)

//export oidfedMetadataCreate
func oidfedMetadataCreate() C.struct_oidfed_metadata {
	m := &oidfed.Metadata{}
	return C.struct_oidfed_metadata{packageGoThing(m)}
}

//export oidfedMetadataGetOPMetadata
func oidfedMetadataGetOPMetadata(m C.struct_oidfed_metadata) C.uintptr_t {
	metadata := cgo.Handle(m.impl).Value().(*oidfed.Metadata)
	if metadata.OpenIDProvider == nil {
		return 0
	}
	return packageGoThing(metadata.OpenIDProvider)
}

//export oidfedMetadataSetOPMetadata
func oidfedMetadataSetOPMetadata(m C.struct_oidfed_metadata, op C.uintptr_t) {
	metadata := cgo.Handle(m.impl).Value().(*oidfed.Metadata)
	metadata.OpenIDProvider = cgo.Handle(op).Value().(*oidfed.OpenIDProviderMetadata)
}

//export oidfedMetadataGetRPMetadata
func oidfedMetadataGetRPMetadata(m C.struct_oidfed_metadata) C.uintptr_t {
	metadata := cgo.Handle(m.impl).Value().(*oidfed.Metadata)
	if metadata.RelyingParty == nil {
		return 0
	}
	return packageGoThing(metadata.RelyingParty)
}

//export oidfedMetadataSetRPMetadata
func oidfedMetadataSetRPMetadata(m C.struct_oidfed_metadata, rp C.uintptr_t) {
	metadata := cgo.Handle(m.impl).Value().(*oidfed.Metadata)
	metadata.RelyingParty = cgo.Handle(rp).Value().(*oidfed.OpenIDRelyingPartyMetadata)
}

//export oidfedMetadataGetFederationEntityMetadata
func oidfedMetadataGetFederationEntityMetadata(m C.struct_oidfed_metadata) C.uintptr_t {
	metadata := cgo.Handle(m.impl).Value().(*oidfed.Metadata)
	if metadata.FederationEntity == nil {
		return 0
	}
	return packageGoThing(metadata.FederationEntity)
}

//export oidfedMetadataSetFederationEntityMetadata
func oidfedMetadataSetFederationEntityMetadata(m C.struct_oidfed_metadata, fe C.uintptr_t) {
	metadata := cgo.Handle(m.impl).Value().(*oidfed.Metadata)
	metadata.FederationEntity = cgo.Handle(fe).Value().(*oidfed.FederationEntityMetadata)
}
