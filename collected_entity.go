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

//export oidfedCollectedEntityDestroy
func oidfedCollectedEntityDestroy(ce *C.struct_oidfed_collected_entity) {
	C.free(unsafe.Pointer(ce.entity_id))
	cgo.Handle(ce.impl).Delete()
}

type uiEnumerationValue struct {
	entity  *oidfed.CollectedEntity
	keys    []string
	current int
}

//export oidfedCollectedEntityEnumerateUi
func oidfedCollectedEntityEnumerateUi(ce C.struct_oidfed_collected_entity) C.struct_oidfed_collected_entity_ui_enumerator {
	entity := cgo.Handle(ce.impl).Value().(*oidfed.CollectedEntity)
	uiEnum := &uiEnumerationValue{entity: entity, keys: []string{}, current: -1}
	for key := range entity.UIInfos {
		uiEnum.keys = append(uiEnum.keys, key)
	}
	opaque := cgo.NewHandle(uiEnum)
	return C.struct_oidfed_collected_entity_ui_enumerator{C.uintptr_t(opaque)}
}

//export oidfedCollectedEntityNextUi
func oidfedCollectedEntityNextUi(enumer *C.struct_oidfed_collected_entity_ui_enumerator) C.bool {
	val := cgo.Handle(enumer.impl).Value().(*uiEnumerationValue)
	if val == nil {
		return C.bool(false)
	}
	if val.current+1 >= len(val.keys) {
		return C.bool(false)
	}
	val.current++
	return C.bool(true)
}

//export oidfedCollectedEntityFinishUi
func oidfedCollectedEntityFinishUi(enumer *C.struct_oidfed_collected_entity_ui_enumerator) {
	cgo.Handle(enumer.impl).Delete()
}

//export oidfedCollectedEntityGetUiValue
func oidfedCollectedEntityGetUiValue(enumer *C.struct_oidfed_collected_entity_ui_enumerator) C.struct_oidfed_ui_info {
	val := cgo.Handle(enumer.impl).Value().(*uiEnumerationValue)
	if val == nil {
		return C.struct_oidfed_ui_info{}
	}
	ui := val.entity.UIInfos[val.keys[val.current]]
	uiH := cgo.NewHandle(ui)

	cname := C.CString(ui.DisplayName)
	return C.struct_oidfed_ui_info{display_name: cname, impl: C.uintptr_t(uiH)}
}
