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

//export oidfedEmptyCollectionFilter
func oidfedEmptyCollectionFilter() C.struct_oidfed_collection_filter {
	ret := &[]oidfed.EntityCollectionFilter{}
	reth := cgo.NewHandle(ret)
	return C.struct_oidfed_collection_filter{C.uintptr_t(reth)}
}

//export oidfedCollectionFilterAppend
func oidfedCollectionFilterAppend(
	cfs *C.struct_oidfed_collection_filter,
	filter C.uintptr_t,
) {
	h := cgo.Handle(cfs.impl)
	defer h.Delete()

	realFilters := h.Value().(*[]oidfed.EntityCollectionFilter)
	// resolve incoming filter handle
	fh := cgo.Handle(filter)
	filterVal := fh.Value().(oidfed.EntityCollectionFilter)
	newFilters := append(*realFilters, filterVal)
	reth := cgo.NewHandle(&newFilters)
	cfs.impl = C.uintptr_t(reth)
}

//export oidfedEntityCollectionFilterOPSupportsExplicitRegistration
func oidfedEntityCollectionFilterOPSupportsExplicitRegistration(
	taIds **C.char,
	taIdsCount C.size_t,
) C.uintptr_t {
	taidsSlice := unsafe.Slice(taIds, uintptr(taIdsCount))
	var idSlice []string
	for _, cstr := range taidsSlice {
		taId := C.GoString(cstr)
		idSlice = append(idSlice, taId)
	}
	filter := oidfed.EntityCollectionFilterOPSupportsExplicitRegistration(idSlice)
	h := cgo.NewHandle(filter)
	return C.uintptr_t(h)
}

//export oidfedEntityCollectionFilterOPSupportsAutomaticRegistration
func oidfedEntityCollectionFilterOPSupportsAutomaticRegistration(
	taIds **C.char,
	taIdsCount C.size_t,
) C.uintptr_t {
	taidsSlice := unsafe.Slice(taIds, uintptr(taIdsCount))
	var idSlice []string
	for _, cstr := range taidsSlice {
		taId := C.GoString(cstr)
		idSlice = append(idSlice, taId)
	}
	filter := oidfed.EntityCollectionFilterOPSupportsAutomaticRegistration(idSlice)
	h := cgo.NewHandle(filter)
	return C.uintptr_t(h)
}

//export oidfedEntityCollectionFilterOPSupportedGrantTypesIncludes
func oidfedEntityCollectionFilterOPSupportedGrantTypesIncludes(
	taIds **C.char,
	taIdsCount C.size_t,
	grantTypes **C.char,
	grantTypesCount C.size_t,
) C.uintptr_t {
	taidsSlice := unsafe.Slice(taIds, uintptr(taIdsCount))
	var idSlice []string
	for _, cstr := range taidsSlice {
		taId := C.GoString(cstr)
		idSlice = append(idSlice, taId)
	}

	grantsSlice := unsafe.Slice(grantTypes, uintptr(grantTypesCount))
	var neededGrantTypes []string
	for _, cstr := range grantsSlice {
		gt := C.GoString(cstr)
		neededGrantTypes = append(neededGrantTypes, gt)
	}

	filter := oidfed.EntityCollectionFilterOPSupportedGrantTypesIncludes(idSlice, neededGrantTypes...)
	h := cgo.NewHandle(filter)
	return C.uintptr_t(h)
}

//export oidfedEntityCollectionFilterOPSupportedScopesIncludes
func oidfedEntityCollectionFilterOPSupportedScopesIncludes(
	taIds **C.char,
	taIdsCount C.size_t,
	scopes **C.char,
	scopesCount C.size_t,
) C.uintptr_t {
	taidsSlice := unsafe.Slice(taIds, uintptr(taIdsCount))
	var idSlice []string
	for _, cstr := range taidsSlice {
		taId := C.GoString(cstr)
		idSlice = append(idSlice, taId)
	}

	scopesSlice := unsafe.Slice(scopes, uintptr(scopesCount))
	var neededScopes []string
	for _, cstr := range scopesSlice {
		sc := C.GoString(cstr)
		neededScopes = append(neededScopes, sc)
	}

	filter := oidfed.EntityCollectionFilterOPSupportedScopesIncludes(idSlice, neededScopes...)
	h := cgo.NewHandle(filter)
	return C.uintptr_t(h)
}

//export oidfedEntityCollectionFilterOPs
func oidfedEntityCollectionFilterOPs() C.uintptr_t {
	filter := oidfed.EntityCollectionFilterOPs()
	h := cgo.NewHandle(filter)
	return C.uintptr_t(h)
}
