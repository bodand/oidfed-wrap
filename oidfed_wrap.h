#ifndef OIDFED_WRAP_H
#define OIDFED_WRAP_H

#include <stdbool.h>
#include <stdint.h>
#include <stdlib.h>

struct oidfed_map {
    uintptr_t impl;
};

struct oidfed_trust_mark {
	void* impl;
};

struct oidfed_collector {
    uintptr_t impl;
};

struct oidfed_collection_filter {
	uintptr_t impl;
};

struct oidfed_trust_anchor {
    const char* entity_id;
    void* jwks;
};

struct oidfed_collected_entity {
    const char* entity_id;
    uintptr_t impl;
};

struct oidfed_collected_entity_ui_enumerator {
    uintptr_t impl;
};

struct oidfed_ui_info {
    const char* display_name;
    uintptr_t impl;
};

struct oidfed_signer {
    uintptr_t impl;
};

struct oidfed_signature_algorithm {
    const char* name;
    bool deprecated;
    bool symmetric;
    uintptr_t impl;
};

struct oidfed_versatile_signer {
    uintptr_t impl;
};

struct oidfed_single_key_storage {
    uintptr_t impl;
};

inline struct oidfed_versatile_signer
oidfedSingleKeyStorageAsVersatileSigner(struct oidfed_single_key_storage* storage) {
    return (struct oidfed_versatile_signer){ .impl = storage->impl };
}

struct oidfed_request_producer {
    uintptr_t impl;
};

struct oidfed_request_object {
    uintptr_t impl;
};

struct oidfed_signed_bytes {
    uintptr_t impl;
};

struct oidfed_entity_statement {
    uintptr_t impl;
};

struct oidfed_trust_resolver {
    uintptr_t impl;
};

struct oidfed_trust_chain {
    uintptr_t impl;
};

struct oidfed_trust_chains {
    uintptr_t impl;
};

struct oidfed_metadata {
    uintptr_t impl;
};

struct oidfed_openid_relying_party_metadata {
    uintptr_t impl;
};

struct oidfed_federation_entity_metadata {
    uintptr_t impl;
};

struct oidfed_federation_leaf {
    uintptr_t impl;
};

#endif
