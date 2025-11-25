package main

import (
	"time"
	"unsafe"

	"github.com/go-oidfed/lib"
	"github.com/zachmann/go-utils/duration"
)

/*
#include "oidfed_wrap.h"
*/
import "C"

//export oidfedTrustMarkCreate
func oidfedTrustMarkCreate() C.struct_oidfed_trust_mark {
	opaque := new(oidfed.EntityConfigurationTrustMarkConfig)
	return C.struct_oidfed_trust_mark{unsafe.Pointer(opaque)}
}

//export oidfedTrustMarkSetType
func oidfedTrustMarkSetType(tm C.struct_oidfed_trust_mark, typ *C.char) {
	x := (*oidfed.EntityConfigurationTrustMarkConfig)(tm.impl)
	x.TrustMarkType = C.GoString(typ)
}

//export oidfedTrustMarkGetType
func oidfedTrustMarkGetType(tm C.struct_oidfed_trust_mark) *C.char {
	x := (*oidfed.EntityConfigurationTrustMarkConfig)(tm.impl)
	return C.CString(x.TrustMarkType)
}

//export oidfedTrustMarkSetIssuer
func oidfedTrustMarkSetIssuer(tm C.struct_oidfed_trust_mark, issuer *C.char) {
	x := (*oidfed.EntityConfigurationTrustMarkConfig)(tm.impl)
	x.TrustMarkIssuer = C.GoString(issuer)
}

//export oidfedTrustMarkGetIssuer
func oidfedTrustMarkGetIssuer(tm C.struct_oidfed_trust_mark) *C.char {
	x := (*oidfed.EntityConfigurationTrustMarkConfig)(tm.impl)
	return C.CString(x.TrustMarkIssuer)
}

//export oidfedTrustMarkSetSelfIssued
func oidfedTrustMarkSetSelfIssued(tm C.struct_oidfed_trust_mark, selfIssued C.bool) {
	x := (*oidfed.EntityConfigurationTrustMarkConfig)(tm.impl)
	x.SelfIssued = bool(selfIssued)
}

//export oidfedTrustMarkIsSelfIssued
func oidfedTrustMarkIsSelfIssued(tm C.struct_oidfed_trust_mark) C.bool {
	x := (*oidfed.EntityConfigurationTrustMarkConfig)(tm.impl)
	return C.bool(x.SelfIssued)
}

//export oidfedTrustMarkSetJwt
func oidfedTrustMarkSetJwt(tm C.struct_oidfed_trust_mark, jwt *C.char) {
	x := (*oidfed.EntityConfigurationTrustMarkConfig)(tm.impl)
	x.JWT = C.GoString(jwt)
}

//export oidfedTrustMarkGetJwt
func oidfedTrustMarkGetJwt(tm C.struct_oidfed_trust_mark) *C.char {
	x := (*oidfed.EntityConfigurationTrustMarkConfig)(tm.impl)
	return C.CString(x.JWT)
}

//export oidfedTrustMarkSetRefresh
func oidfedTrustMarkSetRefresh(tm C.struct_oidfed_trust_mark, refresh C.bool) {
	x := (*oidfed.EntityConfigurationTrustMarkConfig)(tm.impl)
	x.Refresh = bool(refresh)
}

//export oidfedTrustMarkIsRefresh
func oidfedTrustMarkIsRefresh(tm C.struct_oidfed_trust_mark) C.bool {
	x := (*oidfed.EntityConfigurationTrustMarkConfig)(tm.impl)
	return C.bool(x.Refresh)
}

//export oidfedTrustMarkSetMinLifetimeSeconds
func oidfedTrustMarkSetMinLifetimeSeconds(tm C.struct_oidfed_trust_mark, minLifetime C.uint64_t) {
	x := (*oidfed.EntityConfigurationTrustMarkConfig)(tm.impl)
	x.MinLifetime = duration.DurationOption(time.Duration(minLifetime))
}

//export oidfedTrustMarkGetMinLifetimeSeconds
func oidfedTrustMarkGetMinLifetimeSeconds(tm C.struct_oidfed_trust_mark) C.uint64_t {
	x := (*oidfed.EntityConfigurationTrustMarkConfig)(tm.impl)
	return C.uint64_t(x.MinLifetime.Duration().Seconds())
}

//export oidfedTrustMarkSetRefreshGracePeriod
func oidfedTrustMarkSetRefreshGracePeriod(tm C.struct_oidfed_trust_mark, refreshGracePeriod C.uint64_t) {
	x := (*oidfed.EntityConfigurationTrustMarkConfig)(tm.impl)
	x.RefreshGracePeriod = duration.DurationOption(time.Duration(refreshGracePeriod))
}

//export oidfedTrustMarkGetRefreshGracePeriod
func oidfedTrustMarkGetRefreshGracePeriod(tm C.struct_oidfed_trust_mark) C.uint64_t {
	x := (*oidfed.EntityConfigurationTrustMarkConfig)(tm.impl)
	return C.uint64_t(x.RefreshGracePeriod.Duration().Seconds())
}
