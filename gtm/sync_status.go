package gtm

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
)

// SyncStatus holds the configuration of a single SyncStatus.
type SyncStatus struct {
	Big3dDateTime    string `json:"big3dDateTime,omitempty"`
	Generation       int    `json:"generation,omitempty"`
	LastUpdateMicros int    `json:"lastUpdateMicros,omitempty"`
	Syncing          bool   `json:"syncing,omitempty"`
}

// SyncStatusEndpoint represents the REST resource for managing SyncStatus.
const SyncStatusEndpoint = "sync-status"

// SyncStatusResource provides an API to manage SyncStatus configurations.
type SyncStatusResource struct {
	b *bigip.BigIP
}

// List retrieves all SyncStatus details.
func (r *SyncStatusResource) Show() (*SyncStatus, error) {
	var items SyncStatus
	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(SyncStatusEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &items); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &items, nil
}
