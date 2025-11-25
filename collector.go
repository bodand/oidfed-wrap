package main

/*
#include "oidfed_wrap.h"
*/
import "C"
import (
	"runtime/cgo"
	"unsafe"

	oidfed "github.com/go-oidfed/lib"
	"github.com/go-oidfed/lib/apimodel"
)

//export oidfedCollectorDestroy
func oidfedCollectorDestroy(collector *C.struct_oidfed_collector) {
	cgo.Handle(collector.impl).Delete()
}

//export oidfedCollectorCreateSimple
func oidfedCollectorCreateSimple() C.struct_oidfed_collector {
	c := &oidfed.SimpleEntityCollector{}
	handle := cgo.NewHandle(c)
	return C.struct_oidfed_collector{C.uintptr_t(handle)}
}

//export oidfedCollectorCreateSmart
func oidfedCollectorCreateSmart(anchors *C.struct_oidfed_trust_anchor, anchorsCount C.size_t) C.struct_oidfed_collector {
	nativeSlice := unsafe.Slice(anchors, uintptr(anchorsCount))
	taIds := trustAnchorIds(nativeSlice)
	o := oidfed.SmartRemoteEntityCollector{TrustAnchors: taIds}
	oh := cgo.NewHandle(&o)

	return C.struct_oidfed_collector{C.uintptr_t(oh)}
}

func trustAnchorIds(nativeSlice []C.struct_oidfed_trust_anchor) (ret []string) {
	for _, ta := range nativeSlice {
		ret = append(ret, C.GoString(ta.entity_id))
	}
	return
}

//export oidfedCollectorCollectVerifiedEntities
func oidfedCollectorCollectVerifiedEntities(ta C.struct_oidfed_trust_anchor,
	collector *C.struct_oidfed_collector,
	entities **C.struct_oidfed_collected_entity, entitiesCount *C.size_t) C.int {
	return oidfedCollectorCollectVerifiedEntitiesWithFilter(ta, collector, oidfedEmptyCollectionFilter(), entities, entitiesCount)
}

//export oidfedCollectorCollectVerifiedEntitiesWithFilter
func oidfedCollectorCollectVerifiedEntitiesWithFilter(ta C.struct_oidfed_trust_anchor,
	collector *C.struct_oidfed_collector, filters C.struct_oidfed_collection_filter,
	entities **C.struct_oidfed_collected_entity, entitiesCount *C.size_t) C.int {
	h := cgo.Handle(collector.impl)
	realCollector := h.Value().(oidfed.EntityCollector)
	fh := cgo.Handle(filters.impl)
	realFilters := fh.Value().(*[]oidfed.EntityCollectionFilter)

	collectorImpl := oidfed.FilterableVerifiedChainsEntityCollector{
		Collector: realCollector,
		Filters:   *realFilters,
	}
	collectionResp, err := collectorImpl.CollectEntities(
		apimodel.EntityCollectionRequest{TrustAnchor: C.GoString(ta.entity_id),
			EntityTypes: []string{"openid_provider"}},
	)
	if err != nil {
		return C.int(err.Status)
	}

	rawMemory := C.malloc(C.size_t(unsafe.Sizeof(**entities)) * C.size_t(len(collectionResp.FederationEntities)))
	*entities = (*C.struct_oidfed_collected_entity)(rawMemory)
	*entitiesCount = C.size_t(len(collectionResp.FederationEntities))
	for i, fe := range collectionResp.FederationEntities {
		next := (*C.struct_oidfed_collected_entity)(unsafe.Pointer(uintptr(unsafe.Pointer(*entities)) + uintptr(i)*unsafe.Sizeof(**entities)))
		next.entity_id = C.CString(fe.EntityID)
		h := cgo.NewHandle(fe)
		next.impl = C.uintptr_t(h)
	}
	return 0
}
