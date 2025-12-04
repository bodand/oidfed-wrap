package main

/*
#include "oidfed_wrap.h"
*/
import "C"
import "runtime/cgo"

//export oidfCreateMap
func oidfCreateMap() C.struct_oidfed_map {
	m := make(map[string]any)
	return C.struct_oidfed_map{packageGoThing(m)}
}

//export oidfMapHasKey
func oidfMapHasKey(m C.struct_oidfed_map, key *C.char) C.bool {
	impl := cgo.Handle(m.impl).Value().(map[string]any)
	_, ok := impl[C.GoString(key)]
	return C.bool(ok)
}

//export oidfMapSetString
func oidfMapSetString(m C.struct_oidfed_map, key *C.char, value *C.char) {
	impl := cgo.Handle(m.impl).Value().(map[string]any)
	impl[C.GoString(key)] = C.GoString(value)
}

//export oidfMapSetInt64
func oidfMapSetInt64(m C.struct_oidfed_map, key *C.char, value C.int64_t) {
	impl := cgo.Handle(m.impl).Value().(map[string]any)
	impl[C.GoString(key)] = int64(value)
}
