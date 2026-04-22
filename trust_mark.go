package main

import (
	"runtime/cgo"
	"time"
	"unsafe"

	oidfed "github.com/go-oidfed/lib"
	"github.com/zachmann/go-utils/duration"
)

/*
#include "oidfed_wrap.h"
*/
import "C"

//export oidfedTrustMarkCreate
func oidfedTrustMarkCreate() C.struct_oidfed_trust_mark {
	opaque := &oidfed.EntityConfigurationTrustMarkConfig{}
	h := cgo.NewHandle(opaque)
	return C.struct_oidfed_trust_mark{unsafe.Pointer(uintptr(h))}
}

//export oidfedTrustMarkSetType
func oidfedTrustMarkSetType(tm C.struct_oidfed_trust_mark, typ *C.char) {
	x := cgo.Handle(uintptr(tm.impl)).Value().(*oidfed.EntityConfigurationTrustMarkConfig)
	x.TrustMarkType = C.GoString(typ)
}

//export oidfedTrustMarkGetType
func oidfedTrustMarkGetType(tm C.struct_oidfed_trust_mark) *C.char {
	x := cgo.Handle(uintptr(tm.impl)).Value().(*oidfed.EntityConfigurationTrustMarkConfig)
	return C.CString(x.TrustMarkType)
}

//export oidfedTrustMarkSetIssuer
func oidfedTrustMarkSetIssuer(tm C.struct_oidfed_trust_mark, issuer *C.char) {
	x := cgo.Handle(uintptr(tm.impl)).Value().(*oidfed.EntityConfigurationTrustMarkConfig)
	x.TrustMarkIssuer = C.GoString(issuer)
}

//export oidfedTrustMarkGetIssuer
func oidfedTrustMarkGetIssuer(tm C.struct_oidfed_trust_mark) *C.char {
	x := cgo.Handle(uintptr(tm.impl)).Value().(*oidfed.EntityConfigurationTrustMarkConfig)
	return C.CString(x.TrustMarkIssuer)
}

//export oidfedTrustMarkSetJwt
func oidfedTrustMarkSetJwt(tm C.struct_oidfed_trust_mark, jwt *C.char) {
	x := cgo.Handle(uintptr(tm.impl)).Value().(*oidfed.EntityConfigurationTrustMarkConfig)
	x.JWT = C.GoString(jwt)
}

//export oidfedTrustMarkGetJwt
func oidfedTrustMarkGetJwt(tm C.struct_oidfed_trust_mark) *C.char {
	x := cgo.Handle(uintptr(tm.impl)).Value().(*oidfed.EntityConfigurationTrustMarkConfig)
	return C.CString(x.JWT)
}

//export oidfedTrustMarkSetRefresh
func oidfedTrustMarkSetRefresh(tm C.struct_oidfed_trust_mark, refresh C.bool) {
	x := cgo.Handle(uintptr(tm.impl)).Value().(*oidfed.EntityConfigurationTrustMarkConfig)
	x.Refresh = bool(refresh)
}

//export oidfedTrustMarkIsRefresh
func oidfedTrustMarkIsRefresh(tm C.struct_oidfed_trust_mark) C.bool {
	x := cgo.Handle(uintptr(tm.impl)).Value().(*oidfed.EntityConfigurationTrustMarkConfig)
	return C.bool(x.Refresh)
}

//export oidfedTrustMarkSetMinLifetimeSeconds
func oidfedTrustMarkSetMinLifetimeSeconds(tm C.struct_oidfed_trust_mark, minLifetime C.uint64_t) {
	x := cgo.Handle(uintptr(tm.impl)).Value().(*oidfed.EntityConfigurationTrustMarkConfig)
	x.MinLifetime = duration.DurationOption(time.Duration(minLifetime) * time.Second)
}

//export oidfedTrustMarkGetMinLifetimeSeconds
func oidfedTrustMarkGetMinLifetimeSeconds(tm C.struct_oidfed_trust_mark) C.uint64_t {
	x := cgo.Handle(uintptr(tm.impl)).Value().(*oidfed.EntityConfigurationTrustMarkConfig)
	return C.uint64_t(x.MinLifetime.Duration().Seconds())
}

//export oidfedTrustMarkSetRefreshGracePeriod
func oidfedTrustMarkSetRefreshGracePeriod(tm C.struct_oidfed_trust_mark, refreshGracePeriod C.uint64_t) {
	x := cgo.Handle(uintptr(tm.impl)).Value().(*oidfed.EntityConfigurationTrustMarkConfig)
	x.RefreshGracePeriod = duration.DurationOption(time.Duration(refreshGracePeriod) * time.Second)
}

//export oidfedTrustMarkGetRefreshGracePeriod
func oidfedTrustMarkGetRefreshGracePeriod(tm C.struct_oidfed_trust_mark) C.uint64_t {
	x := cgo.Handle(uintptr(tm.impl)).Value().(*oidfed.EntityConfigurationTrustMarkConfig)
	return C.uint64_t(x.RefreshGracePeriod.Duration().Seconds())
}

//export oidfedTrustMarkDestroy
func oidfedTrustMarkDestroy(tm *C.struct_oidfed_trust_mark) {
	if tm.impl != nil {
		cgo.Handle(uintptr(tm.impl)).Delete()
	}
}
