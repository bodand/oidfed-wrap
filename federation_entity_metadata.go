package main

/*
#include "oidfed_wrap.h"
*/
import "C"
import (
	"runtime/cgo"

	oidfed "github.com/go-oidfed/lib"
)

//export oidfedFederationEntityMetadataCreate
func oidfedFederationEntityMetadataCreate() C.struct_oidfed_federation_entity_metadata {
	m := &oidfed.FederationEntityMetadata{}
	return C.struct_oidfed_federation_entity_metadata{packageGoThing(m)}
}

//export oidfedFederationEntityMetadataGetFederationFetchEndpoint
func oidfedFederationEntityMetadataGetFederationFetchEndpoint(m C.struct_oidfed_federation_entity_metadata) *C.char {
	metadata := cgo.Handle(m.impl).Value().(*oidfed.FederationEntityMetadata)
	return C.CString(metadata.FederationFetchEndpoint)
}

//export oidfedFederationEntityMetadataSetFederationFetchEndpoint
func oidfedFederationEntityMetadataSetFederationFetchEndpoint(m C.struct_oidfed_federation_entity_metadata, endpoint *C.char) {
	metadata := cgo.Handle(m.impl).Value().(*oidfed.FederationEntityMetadata)
	metadata.FederationFetchEndpoint = C.GoString(endpoint)
}

//export oidfedFederationEntityMetadataGetFederationListEndpoint
func oidfedFederationEntityMetadataGetFederationListEndpoint(m C.struct_oidfed_federation_entity_metadata) *C.char {
	metadata := cgo.Handle(m.impl).Value().(*oidfed.FederationEntityMetadata)
	return C.CString(metadata.FederationListEndpoint)
}

//export oidfedFederationEntityMetadataSetFederationListEndpoint
func oidfedFederationEntityMetadataSetFederationListEndpoint(m C.struct_oidfed_federation_entity_metadata, endpoint *C.char) {
	metadata := cgo.Handle(m.impl).Value().(*oidfed.FederationEntityMetadata)
	metadata.FederationListEndpoint = C.GoString(endpoint)
}

//export oidfedFederationEntityMetadataGetFederationResolveEndpoint
func oidfedFederationEntityMetadataGetFederationResolveEndpoint(m C.struct_oidfed_federation_entity_metadata) *C.char {
	metadata := cgo.Handle(m.impl).Value().(*oidfed.FederationEntityMetadata)
	return C.CString(metadata.FederationResolveEndpoint)
}

//export oidfedFederationEntityMetadataSetFederationResolveEndpoint
func oidfedFederationEntityMetadataSetFederationResolveEndpoint(m C.struct_oidfed_federation_entity_metadata, endpoint *C.char) {
	metadata := cgo.Handle(m.impl).Value().(*oidfed.FederationEntityMetadata)
	metadata.FederationResolveEndpoint = C.GoString(endpoint)
}

//export oidfedFederationEntityMetadataGetFederationTrustMarkStatusEndpoint
func oidfedFederationEntityMetadataGetFederationTrustMarkStatusEndpoint(m C.struct_oidfed_federation_entity_metadata) *C.char {
	metadata := cgo.Handle(m.impl).Value().(*oidfed.FederationEntityMetadata)
	return C.CString(metadata.FederationTrustMarkStatusEndpoint)
}

//export oidfedFederationEntityMetadataSetFederationTrustMarkStatusEndpoint
func oidfedFederationEntityMetadataSetFederationTrustMarkStatusEndpoint(m C.struct_oidfed_federation_entity_metadata, endpoint *C.char) {
	metadata := cgo.Handle(m.impl).Value().(*oidfed.FederationEntityMetadata)
	metadata.FederationTrustMarkStatusEndpoint = C.GoString(endpoint)
}

//export oidfedFederationEntityMetadataGetFederationTrustMarkListEndpoint
func oidfedFederationEntityMetadataGetFederationTrustMarkListEndpoint(m C.struct_oidfed_federation_entity_metadata) *C.char {
	metadata := cgo.Handle(m.impl).Value().(*oidfed.FederationEntityMetadata)
	return C.CString(metadata.FederationTrustMarkListEndpoint)
}

//export oidfedFederationEntityMetadataSetFederationTrustMarkListEndpoint
func oidfedFederationEntityMetadataSetFederationTrustMarkListEndpoint(m C.struct_oidfed_federation_entity_metadata, endpoint *C.char) {
	metadata := cgo.Handle(m.impl).Value().(*oidfed.FederationEntityMetadata)
	metadata.FederationTrustMarkListEndpoint = C.GoString(endpoint)
}

//export oidfedFederationEntityMetadataGetFederationTrustMarkEndpoint
func oidfedFederationEntityMetadataGetFederationTrustMarkEndpoint(m C.struct_oidfed_federation_entity_metadata) *C.char {
	metadata := cgo.Handle(m.impl).Value().(*oidfed.FederationEntityMetadata)
	return C.CString(metadata.FederationTrustMarkEndpoint)
}

//export oidfedFederationEntityMetadataSetFederationTrustMarkEndpoint
func oidfedFederationEntityMetadataSetFederationTrustMarkEndpoint(m C.struct_oidfed_federation_entity_metadata, endpoint *C.char) {
	metadata := cgo.Handle(m.impl).Value().(*oidfed.FederationEntityMetadata)
	metadata.FederationTrustMarkEndpoint = C.GoString(endpoint)
}

//export oidfedFederationEntityMetadataGetFederationHistoricalKeysEndpoint
func oidfedFederationEntityMetadataGetFederationHistoricalKeysEndpoint(m C.struct_oidfed_federation_entity_metadata) *C.char {
	metadata := cgo.Handle(m.impl).Value().(*oidfed.FederationEntityMetadata)
	return C.CString(metadata.FederationHistoricalLKeysEndpoint)
}

//export oidfedFederationEntityMetadataSetFederationHistoricalKeysEndpoint
func oidfedFederationEntityMetadataSetFederationHistoricalKeysEndpoint(m C.struct_oidfed_federation_entity_metadata, endpoint *C.char) {
	metadata := cgo.Handle(m.impl).Value().(*oidfed.FederationEntityMetadata)
	metadata.FederationHistoricalLKeysEndpoint = C.GoString(endpoint)
}

//export oidfedFederationEntityMetadataGetLogoURI
func oidfedFederationEntityMetadataGetLogoURI(m C.struct_oidfed_federation_entity_metadata) *C.char {
	metadata := cgo.Handle(m.impl).Value().(*oidfed.FederationEntityMetadata)
	return C.CString(metadata.LogoURI)
}

//export oidfedFederationEntityMetadataSetLogoURI
func oidfedFederationEntityMetadataSetLogoURI(m C.struct_oidfed_federation_entity_metadata, logoURI *C.char) {
	metadata := cgo.Handle(m.impl).Value().(*oidfed.FederationEntityMetadata)
	metadata.LogoURI = C.GoString(logoURI)
}

//export oidfedFederationEntityMetadataGetPolicyURI
func oidfedFederationEntityMetadataGetPolicyURI(m C.struct_oidfed_federation_entity_metadata) *C.char {
	metadata := cgo.Handle(m.impl).Value().(*oidfed.FederationEntityMetadata)
	return C.CString(metadata.PolicyURI)
}

//export oidfedFederationEntityMetadataSetPolicyURI
func oidfedFederationEntityMetadataSetPolicyURI(m C.struct_oidfed_federation_entity_metadata, policyURI *C.char) {
	metadata := cgo.Handle(m.impl).Value().(*oidfed.FederationEntityMetadata)
	metadata.PolicyURI = C.GoString(policyURI)
}

//export oidfedFederationEntityMetadataGetInformationURI
func oidfedFederationEntityMetadataGetInformationURI(m C.struct_oidfed_federation_entity_metadata) *C.char {
	metadata := cgo.Handle(m.impl).Value().(*oidfed.FederationEntityMetadata)
	return C.CString(metadata.InformationURI)
}

//export oidfedFederationEntityMetadataSetInformationURI
func oidfedFederationEntityMetadataSetInformationURI(m C.struct_oidfed_federation_entity_metadata, informationURI *C.char) {
	metadata := cgo.Handle(m.impl).Value().(*oidfed.FederationEntityMetadata)
	metadata.InformationURI = C.GoString(informationURI)
}

//export oidfedFederationEntityMetadataGetOrganizationName
func oidfedFederationEntityMetadataGetOrganizationName(m C.struct_oidfed_federation_entity_metadata) *C.char {
	metadata := cgo.Handle(m.impl).Value().(*oidfed.FederationEntityMetadata)
	return C.CString(metadata.OrganizationName)
}

//export oidfedFederationEntityMetadataSetOrganizationName
func oidfedFederationEntityMetadataSetOrganizationName(m C.struct_oidfed_federation_entity_metadata, organizationName *C.char) {
	metadata := cgo.Handle(m.impl).Value().(*oidfed.FederationEntityMetadata)
	metadata.OrganizationName = C.GoString(organizationName)
}

//export oidfedFederationEntityMetadataGetOrganizationURL
func oidfedFederationEntityMetadataGetOrganizationURL(m C.struct_oidfed_federation_entity_metadata) *C.char {
	metadata := cgo.Handle(m.impl).Value().(*oidfed.FederationEntityMetadata)
	return C.CString(metadata.OrganizationURI)
}

//export oidfedFederationEntityMetadataSetOrganizationURL
func oidfedFederationEntityMetadataSetOrganizationURL(m C.struct_oidfed_federation_entity_metadata, organizationURL *C.char) {
	metadata := cgo.Handle(m.impl).Value().(*oidfed.FederationEntityMetadata)
	metadata.OrganizationURI = C.GoString(organizationURL)
}

//export oidfedFederationEntityMetadataDestroy
func oidfedFederationEntityMetadataDestroy(m *C.struct_oidfed_federation_entity_metadata) {
	if m.impl != 0 {
		cgo.Handle(m.impl).Delete()
	}
}
