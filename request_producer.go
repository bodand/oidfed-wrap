package main

/*
#include "oidfed_wrap.h"
*/
import "C"
import (
	"runtime/cgo"
	"time"

	oidfed "github.com/go-oidfed/lib"
	"github.com/go-oidfed/lib/jwx"
	"github.com/lestrrat-go/jwx/v3/jws"
)

//export oidfedRequestProducerCreate
func oidfedRequestProducerCreate(
	entityId *C.char,
	duration C.int64_t,
	signer C.struct_oidfed_versatile_signer,
) C.struct_oidfed_request_producer {
	goEntityId := C.GoString(entityId)
	vs := cgo.Handle(signer.impl).Value().(jwx.VersatileSigner)
	prod := oidfed.NewRequestObjectProducer(goEntityId,
		vs,
		time.Duration(int64(duration)))
	return C.struct_oidfed_request_producer{packageGoThing(prod)}
}

//export oidfedRequestProducerDestroy
func oidfedRequestProducerDestroy(producer *C.struct_oidfed_request_producer) {
	cgo.Handle(producer.impl).Delete()
}

//export oidfedRequestProducerProduceObject
func oidfedRequestProducerProduceObject(
	producer C.struct_oidfed_request_producer,
	requestValues C.struct_oidfed_map,
	headers C.struct_oidfed_jws_headers,
	algorithms **C.char,
	algorithmsCount C.size_t,
	errc *C.int,
) C.struct_oidfed_signed_bytes {
	rop := cgo.Handle(producer.impl).Value().(*oidfed.RequestObjectProducer)
	rv := cgo.Handle(requestValues.impl).Value().(map[string]any)

	var h jws.Headers
	if headers.impl != 0 {
		h = cgo.Handle(headers.impl).Value().(jws.Headers)
	}

	algs := goifyCArray(algorithms, algorithmsCount)
	bytes, err := rop.RequestObject(rv, h, algs...)
	if err != nil {
		*errc = 1
		return C.struct_oidfed_signed_bytes{}
	}
	return C.struct_oidfed_signed_bytes{packageGoThing(bytes)}
}

//export oidfedRequestProducerClientAssertion
func oidfedRequestProducerClientAssertion(
	producer C.struct_oidfed_request_producer,
	audience *C.char,
	algorithms **C.char,
	algorithmsCount C.size_t,
	errc *C.int,
) C.struct_oidfed_signed_bytes {
	rop := cgo.Handle(producer.impl).Value().(*oidfed.RequestObjectProducer)
	goAudience := C.GoString(audience)
	algs := goifyCArray(algorithms, algorithmsCount)
	bytes, err := rop.ClientAssertion(goAudience, algs...)
	if err != nil {
		*errc = 1
		return C.struct_oidfed_signed_bytes{}
	}
	return C.struct_oidfed_signed_bytes{packageGoThing(bytes)}
}

//export oidfedSignedBytesDestroy
func oidfedSignedBytesDestroy(sb *C.struct_oidfed_signed_bytes) {
	if sb.impl != 0 {
		cgo.Handle(sb.impl).Delete()
	}
}

//export oidfedSignedBytesGetData
func oidfedSignedBytesGetData(sb C.struct_oidfed_signed_bytes, count *C.size_t) *C.char {
	bytes := cgo.Handle(sb.impl).Value().([]byte)
	*count = C.size_t(len(bytes))
	return (*C.char)(C.CBytes(bytes))
}
