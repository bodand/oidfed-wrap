#include <stdio.h>
#include <assert.h>

#include "oidfed_wrap_lib.h"

int
main() {
    bool success = false;
    struct oidfed_signature_algorithm es384 = oidfedSignatureAlgorithmGet("ES384", &success);
    assert(success);

    struct oidfed_trust_anchor ta = oidfedTrustAnchorCreate("https://ta.oidf-pilot.edugain.org");
    struct oidfed_collector collector = oidfedCollectorCreateSmart(&ta, 1);
    struct oidfed_collection_filter filter = oidfedEmptyCollectionFilter();
    oidfedCollectionFilterAppend(&filter, oidfedEntityCollectionFilterOPs());
    char* op_uri = "https://ta.oidf-pilot.edugain.org";
    oidfedCollectionFilterAppend(&filter, oidfedEntityCollectionFilterOPSupportsAutomaticRegistration(&op_uri, 1));

    int errc = 0;
    struct oidfed_entity_statement stmt = oidfedGetEntityConfiguration(op_uri, &errc);
    if (errc == 0) {
        printf("Got entity configuration for %s\n", op_uri);
        oidfedEntityStatementDestroy(&stmt);
    }

    struct oidfed_trust_resolver resolver = oidfedTrustResolverCreate(op_uri, &ta, 1);
    struct oidfed_trust_chains chains = oidfedTrustResolverResolveToValidChains(resolver);
    size_t chains_count = oidfedTrustChainsCount(chains);
    printf("Found %zu trust chains\n", chains_count);
    for (size_t i = 0; i < chains_count; i++) {
        struct oidfed_trust_chain chain = oidfedTrustChainsGet(chains, i);
        struct oidfed_metadata metadata = oidfedTrustChainGetMetadata(chain, &errc);
        if (errc == 0) {
            printf("  - Chain %zu metadata resolved\n", i);
            oidfedMetadataDestroy(&metadata);
        }
        oidfedTrustChainDestroy(&chain);
    }
    oidfedTrustChainsDestroy(&chains);
    oidfedTrustResolverDestroy(&resolver);

    struct oidfed_collected_entity* entities = 0;
    size_t entities_len = 0;
    oidfedCollectorCollectVerifiedEntitiesWithFilter(ta, &collector, filter, &entities, &entities_len);
    printf("%zu\n", entities_len);
    for (size_t i = 0; i < entities_len; i++) {
        printf("%s\n", entities[i].entity_id);
        struct oidfed_collected_entity_ui_enumerator enumerator =
            oidfedCollectedEntityEnumerateUi(entities[i]);
        while (oidfedCollectedEntityNextUi(&enumerator)) {
            struct oidfed_ui_info ui = oidfedCollectedEntityGetUiValue(&enumerator);
            printf("\t- %s\n", ui.display_name);
            oidfedUiInfoDestroy(&ui);
        }
        oidfedCollectedEntityFinishUi(&enumerator);
        oidfedCollectedEntityDestroy(&entities[i]);
    }

    oidfedCollectionFilterDestroy(&filter);
    oidfedCollectorDestroy(&collector);
    oidfedTrustAnchorDestroy(&ta);
    free(entities);
    oidfedSignatureAlgorithmDestroy(&es384);
	return 0;
}
