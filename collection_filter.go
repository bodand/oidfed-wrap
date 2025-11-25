package main

/*
#include "oidfed_wrap.h"
*/
import "C"
import (
	"runtime/cgo"

	oidfed "github.com/go-oidfed/lib"
)

//export oidfedEmptyCollectionFilter
func oidfedEmptyCollectionFilter() C.struct_oidfed_collection_filter {
	ret := &[]oidfed.EntityCollectionFilter{}
	val := packageGoThing(ret)
	return C.struct_oidfed_collection_filter{val}
}

//export oidfedCollectionFilterAppend
func oidfedCollectionFilterAppend(
	cfs *C.struct_oidfed_collection_filter,
	filter C.uintptr_t,
) {
	h := cgo.Handle(cfs.impl)
	defer h.Delete()
	realFilters := h.Value().(*[]oidfed.EntityCollectionFilter)

	fh := cgo.Handle(filter)
	filterVal := fh.Value().(oidfed.EntityCollectionFilter)

	newFilters := append(*realFilters, filterVal)
	cfs.impl = packageGoThing(&newFilters)
}

//export oidfedEntityCollectionFilterOPSupportsExplicitRegistration
func oidfedEntityCollectionFilterOPSupportsExplicitRegistration(
	taIds **C.char,
	taIdsCount C.size_t,
) C.uintptr_t {
	idSlice := goifyCArray(taIds, taIdsCount)
	filter := oidfed.EntityCollectionFilterOPSupportsExplicitRegistration(idSlice)
	return packageGoThing(filter)
}

//export oidfedEntityCollectionFilterOPSupportsAutomaticRegistration
func oidfedEntityCollectionFilterOPSupportsAutomaticRegistration(
	taIds **C.char,
	taIdsCount C.size_t,
) C.uintptr_t {
	idSlice := goifyCArray(taIds, taIdsCount)
	filter := oidfed.EntityCollectionFilterOPSupportsAutomaticRegistration(idSlice)
	return packageGoThing(filter)
}

//export oidfedEntityCollectionFilterOPSupportedGrantTypesIncludes
func oidfedEntityCollectionFilterOPSupportedGrantTypesIncludes(
	taIds **C.char,
	taIdsCount C.size_t,
	grantTypes **C.char,
	grantTypesCount C.size_t,
) C.uintptr_t {
	idSlice := goifyCArray(taIds, taIdsCount)
	neededGrantTypes := goifyCArray(grantTypes, grantTypesCount)

	filter := oidfed.EntityCollectionFilterOPSupportedGrantTypesIncludes(idSlice, neededGrantTypes...)
	return packageGoThing(filter)
}

//export oidfedEntityCollectionFilterOPSupportedScopesIncludes
func oidfedEntityCollectionFilterOPSupportedScopesIncludes(
	taIds **C.char,
	taIdsCount C.size_t,
	scopes **C.char,
	scopesCount C.size_t,
) C.uintptr_t {
	idSlice := goifyCArray(taIds, taIdsCount)
	neededScopes := goifyCArray(scopes, scopesCount)

	filter := oidfed.EntityCollectionFilterOPSupportedScopesIncludes(idSlice, neededScopes...)
	return packageGoThing(filter)
}

//export oidfedEntityCollectionFilterOPs
func oidfedEntityCollectionFilterOPs() C.uintptr_t {
	filter := oidfed.EntityCollectionFilterOPs()
	return packageGoThing(filter)
}
