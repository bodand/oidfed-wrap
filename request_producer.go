package main

/*
#include "oidfed_wrap.h"
*/
import "C"
import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"runtime/cgo"
	"time"

	oidfed "github.com/go-oidfed/lib"
	"github.com/go-oidfed/lib/jwx"
	"github.com/lestrrat-go/jwx/v3/jws"
	"github.com/lestrrat-go/jwx/v3/jwt"
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
		log.Printf("aaaa: %s", err)
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

//export oidfedRequestProducerExchangeCode
func oidfedRequestProducerExchangeCode(
	producer C.struct_oidfed_request_producer,
	tokenEndpoint *C.char,
	code *C.char,
	redirectURI *C.char,
	errc *C.int,
) *C.char {
	rop := cgo.Handle(producer.impl).Value().(*oidfed.RequestObjectProducer)
	goTokenEndpoint := C.GoString(tokenEndpoint)
	goCode := C.GoString(code)
	goRedirectURI := C.GoString(redirectURI)

	clientAssertion, err := rop.ClientAssertion(goTokenEndpoint)
	if err != nil {
		log.Printf("failed to produce client assertion: %v", err)
		*errc = 1
		return nil
	}

	params := url.Values{}
	params.Set("grant_type", "authorization_code")
	params.Set("code", goCode)
	params.Set("redirect_uri", goRedirectURI)
	params.Set("client_id", rop.EntityID)
	params.Set("client_assertion_type", "urn:ietf:params:oauth:client-assertion-type:jwt-bearer")
	params.Set("client_assertion", string(clientAssertion))

	res, err := http.PostForm(goTokenEndpoint, params)
	if err != nil {
		log.Printf("failed to post form to token endpoint: %v", err)
		*errc = 1
		return nil
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Printf("failed to read response body: %v", err)
		*errc = 1
		return nil
	}

	return C.CString(string(body))
}

//export oidfedExtractSubjectFromTokenResponse
func oidfedExtractSubjectFromTokenResponse(tokenResponse *C.char) *C.char {
	goTokenResponse := C.GoString(tokenResponse)
	var data map[string]interface{}
	if err := json.Unmarshal([]byte(goTokenResponse), &data); err != nil {
		log.Printf("Failed to unmarshal token response: %v", err)
		return nil
	}

	idToken, ok := data["id_token"].(string)
	if !ok {
		if sub, ok := data["sub"].(string); ok {
			return C.CString(sub)
		}
		return nil
	}

	t, err := jwt.ParseString(idToken, jwt.WithVerify(false))
	if err != nil {
		log.Printf("Failed to parse ID token: %v", err)
		return nil
	}

	sub, ok := t.Subject()
	if !ok {
		log.Printf("ID token does not contain a subject")
		return nil
	}

	return C.CString(sub)
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
