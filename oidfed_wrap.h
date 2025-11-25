#ifndef OIDFED_WRAP_H
#define OIDFED_WRAP_H

#include <stdbool.h>
#include <stdint.h>
#include <stdlib.h>

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

#endif
