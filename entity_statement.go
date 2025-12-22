package main

/*
#include "oidfed_wrap.h"
*/
import "C"
import (
	"runtime/cgo"

	oidfed "github.com/go-oidfed/lib"
)

//export oidfedEntityStatementParse
func oidfedEntityStatementParse(
	jwt *C.char,
	jwtLen C.size_t,
	errc *C.int,
) C.struct_oidfed_entity_statement {
	goJwt := goifyCByteArray(jwt, jwtLen)
	stmt, err := oidfed.ParseEntityStatement(goJwt)
	if err != nil {
		*errc = 1
		return C.struct_oidfed_entity_statement{0}
	}
	return C.struct_oidfed_entity_statement{packageGoThing(stmt)}
}

//export oidfedEntityStatementDestroy
func oidfedEntityStatementDestroy(stmt *C.struct_oidfed_entity_statement) {
	if stmt.impl != 0 {
		cgo.Handle(stmt.impl).Delete()
	}
}

//export oidfedGetEntityConfiguration
func oidfedGetEntityConfiguration(
	entityID *C.char,
	errc *C.int,
) C.struct_oidfed_entity_statement {
	goEntityID := C.GoString(entityID)
	stmt, err := oidfed.GetEntityConfiguration(goEntityID)
	if err != nil {
		*errc = 1
		return C.struct_oidfed_entity_statement{0}
	}
	return C.struct_oidfed_entity_statement{packageGoThing(stmt)}
}
