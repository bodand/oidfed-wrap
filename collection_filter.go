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
func oidfedCollectionFilterAppend(cfs *C.struct_oidfed_collection_filter, filter unsafe.Pointer) {
	h := cgo.Handle(cfs.impl)
	defer h.Delete()

	realFilters := h.Value().(*[]oidfed.EntityCollectionFilter)
	newFilters := append(*realFilters, *(*oidfed.EntityCollectionFilter)(filter))
	reth := cgo.NewHandle(&newFilters)
	cfs.impl = C.uintptr_t(reth)
}

//export oidfedEntityCollectionFilterOPSupportsExplicitRegistration
func oidfedEntityCollectionFilterOPSupportsExplicitRegistration(taIds **C.char, taIdsCount C.size_t) unsafe.Pointer {
	taidsSlice := unsafe.Slice(taIds, uintptr(taIdsCount))
	var idSlice []string
	for _, cstr := range taidsSlice {
		taId := C.GoString(cstr)
		idSlice = append(idSlice, taId)
	}
	filter := oidfed.EntityCollectionFilterOPSupportsExplicitRegistration(idSlice)
	return unsafe.Pointer(&filter)
}

//export oidfedEntityCollectionFilterOPSupportsAutomaticRegistration
func oidfedEntityCollectionFilterOPSupportsAutomaticRegistration(taIds **C.char, taIdsCount C.size_t) unsafe.Pointer {
	taidsSlice := unsafe.Slice(taIds, uintptr(taIdsCount))
	var idSlice []string
	for _, cstr := range taidsSlice {
		taId := C.GoString(cstr)
		idSlice = append(idSlice, taId)
	}
	filter := oidfed.EntityCollectionFilterOPSupportsAutomaticRegistration(idSlice)
	return unsafe.Pointer(&filter)
}
