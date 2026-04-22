package main

/*
#include "oidfed_wrap.h"
*/
import "C"
import (
	"runtime/cgo"

	oidfed "github.com/go-oidfed/lib"
	"github.com/go-oidfed/lib/jwx"
)

//export oidfedOpenIDRelyingPartyMetadataCreate
func oidfedOpenIDRelyingPartyMetadataCreate() C.struct_oidfed_openid_relying_party_metadata {
	rp := &oidfed.OpenIDRelyingPartyMetadata{}
	return C.struct_oidfed_openid_relying_party_metadata{packageGoThing(rp)}
}

//export oidfedOpenIDRelyingPartyMetadataGetClientID
func oidfedOpenIDRelyingPartyMetadataGetClientID(rp C.struct_oidfed_openid_relying_party_metadata) *C.char {
	metadata := cgo.Handle(rp.impl).Value().(*oidfed.OpenIDRelyingPartyMetadata)
	return C.CString(metadata.ClientID)
}

//export oidfedOpenIDRelyingPartyMetadataSetClientID
func oidfedOpenIDRelyingPartyMetadataSetClientID(rp C.struct_oidfed_openid_relying_party_metadata, clientID *C.char) {
	metadata := cgo.Handle(rp.impl).Value().(*oidfed.OpenIDRelyingPartyMetadata)
	metadata.ClientID = C.GoString(clientID)
}

//export oidfedOpenIDRelyingPartyMetadataGetClientName
func oidfedOpenIDRelyingPartyMetadataGetClientName(rp C.struct_oidfed_openid_relying_party_metadata) *C.char {
	metadata := cgo.Handle(rp.impl).Value().(*oidfed.OpenIDRelyingPartyMetadata)
	return C.CString(metadata.ClientName)
}

//export oidfedOpenIDRelyingPartyMetadataSetClientName
func oidfedOpenIDRelyingPartyMetadataSetClientName(rp C.struct_oidfed_openid_relying_party_metadata, clientName *C.char) {
	metadata := cgo.Handle(rp.impl).Value().(*oidfed.OpenIDRelyingPartyMetadata)
	metadata.ClientName = C.GoString(clientName)
}

//export oidfedOpenIDRelyingPartyMetadataSetRedirectUris
func oidfedOpenIDRelyingPartyMetadataSetRedirectUris(rp C.struct_oidfed_openid_relying_party_metadata, uris **C.char, count C.size_t) {
	metadata := cgo.Handle(rp.impl).Value().(*oidfed.OpenIDRelyingPartyMetadata)
	metadata.RedirectURIS = goifyCArray(uris, count)
}

//export oidfedOpenIDRelyingPartyMetadataSetResponseTypes
func oidfedOpenIDRelyingPartyMetadataSetResponseTypes(rp C.struct_oidfed_openid_relying_party_metadata, types **C.char, count C.size_t) {
	metadata := cgo.Handle(rp.impl).Value().(*oidfed.OpenIDRelyingPartyMetadata)
	metadata.ResponseTypes = goifyCArray(types, count)
}

//export oidfedOpenIDRelyingPartyMetadataSetGrantTypes
func oidfedOpenIDRelyingPartyMetadataSetGrantTypes(rp C.struct_oidfed_openid_relying_party_metadata, types **C.char, count C.size_t) {
	metadata := cgo.Handle(rp.impl).Value().(*oidfed.OpenIDRelyingPartyMetadata)
	metadata.GrantTypes = goifyCArray(types, count)
}

//export oidfedOpenIDRelyingPartyMetadataGetApplicationType
func oidfedOpenIDRelyingPartyMetadataGetApplicationType(rp C.struct_oidfed_openid_relying_party_metadata) *C.char {
	metadata := cgo.Handle(rp.impl).Value().(*oidfed.OpenIDRelyingPartyMetadata)
	return C.CString(metadata.ApplicationType)
}

//export oidfedOpenIDRelyingPartyMetadataSetApplicationType
func oidfedOpenIDRelyingPartyMetadataSetApplicationType(rp C.struct_oidfed_openid_relying_party_metadata, applicationType *C.char) {
	metadata := cgo.Handle(rp.impl).Value().(*oidfed.OpenIDRelyingPartyMetadata)
	metadata.ApplicationType = C.GoString(applicationType)
}

//export oidfedOpenIDRelyingPartyMetadataGetLogoURI
func oidfedOpenIDRelyingPartyMetadataGetLogoURI(rp C.struct_oidfed_openid_relying_party_metadata) *C.char {
	metadata := cgo.Handle(rp.impl).Value().(*oidfed.OpenIDRelyingPartyMetadata)
	return C.CString(metadata.LogoURI)
}

//export oidfedOpenIDRelyingPartyMetadataSetLogoURI
func oidfedOpenIDRelyingPartyMetadataSetLogoURI(rp C.struct_oidfed_openid_relying_party_metadata, logoURI *C.char) {
	metadata := cgo.Handle(rp.impl).Value().(*oidfed.OpenIDRelyingPartyMetadata)
	metadata.LogoURI = C.GoString(logoURI)
}

//export oidfedOpenIDRelyingPartyMetadataGetJWKS
func oidfedOpenIDRelyingPartyMetadataGetJWKS(rp C.struct_oidfed_openid_relying_party_metadata) C.uintptr_t {
	metadata := cgo.Handle(rp.impl).Value().(*oidfed.OpenIDRelyingPartyMetadata)
	if metadata.JWKS == nil {
		return 0
	}
	return packageGoThing(metadata.JWKS)
}

//export oidfedOpenIDRelyingPartyMetadataSetJWKS
func oidfedOpenIDRelyingPartyMetadataSetJWKS(rp C.struct_oidfed_openid_relying_party_metadata, jwks C.uintptr_t) {
	metadata := cgo.Handle(rp.impl).Value().(*oidfed.OpenIDRelyingPartyMetadata)
	metadata.JWKS = cgo.Handle(jwks).Value().(*jwx.JWKS)
}

//export oidfedOpenIDRelyingPartyMetadataSetJWKSFromKeyStorage
func oidfedOpenIDRelyingPartyMetadataSetJWKSFromKeyStorage(
	rp C.struct_oidfed_openid_relying_party_metadata,
	storage C.struct_oidfed_single_key_storage,
) {
	goStorage := cgo.Handle(storage.impl).Value().(jwx.SingleKeySigner)
	metadata := cgo.Handle(rp.impl).Value().(*oidfed.OpenIDRelyingPartyMetadata)

	jwks, err := goStorage.JWKS()
	if err == nil {
		metadata.JWKS = &jwks
	}
}

//export oidfedOpenIDRelyingPartyMetadataGetOrganizationName
func oidfedOpenIDRelyingPartyMetadataGetOrganizationName(rp C.struct_oidfed_openid_relying_party_metadata) *C.char {
	metadata := cgo.Handle(rp.impl).Value().(*oidfed.OpenIDRelyingPartyMetadata)
	return C.CString(metadata.OrganizationName)
}

//export oidfedOpenIDRelyingPartyMetadataSetOrganizationName
func oidfedOpenIDRelyingPartyMetadataSetOrganizationName(rp C.struct_oidfed_openid_relying_party_metadata, organizationName *C.char) {
	metadata := cgo.Handle(rp.impl).Value().(*oidfed.OpenIDRelyingPartyMetadata)
	metadata.OrganizationName = C.GoString(organizationName)
}

//export oidfedOpenIDRelyingPartyMetadataSetClientRegistrationTypes
func oidfedOpenIDRelyingPartyMetadataSetClientRegistrationTypes(rp C.struct_oidfed_openid_relying_party_metadata, types **C.char, count C.size_t) {
	metadata := cgo.Handle(rp.impl).Value().(*oidfed.OpenIDRelyingPartyMetadata)
	metadata.ClientRegistrationTypes = goifyCArray(types, count)
}

//export oidfedOpenIDRelyingPartyMetadataDestroy
func oidfedOpenIDRelyingPartyMetadataDestroy(rp *C.struct_oidfed_openid_relying_party_metadata) {
	if rp.impl != 0 {
		cgo.Handle(rp.impl).Delete()
	}
}
