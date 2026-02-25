package main

/*
#include "oidfed_wrap.h"
*/
import "C"
import (
	"runtime/cgo"

	oidfed "github.com/go-oidfed/lib"
)

//export oidfedOpenIDProviderMetadataGetAuthorizationEndpoint
func oidfedOpenIDProviderMetadataGetAuthorizationEndpoint(m C.struct_oidfed_openid_provider_metadata) *C.char {
	metadata := cgo.Handle(m.impl).Value().(*oidfed.OpenIDProviderMetadata)
	return C.CString(metadata.AuthorizationEndpoint)
}

//export oidfedOpenIDProviderMetadataGetIssuer
func oidfedOpenIDProviderMetadataGetIssuer(m C.struct_oidfed_openid_provider_metadata) *C.char {
	metadata := cgo.Handle(m.impl).Value().(*oidfed.OpenIDProviderMetadata)
	return C.CString(metadata.Issuer)
}

//export oidfedOpenIDProviderMetadataDestroy
func oidfedOpenIDProviderMetadataDestroy(m *C.struct_oidfed_openid_provider_metadata) {
	if m == nil {
		return
	}
	h := cgo.Handle(m.impl)
	h.Delete()
}
