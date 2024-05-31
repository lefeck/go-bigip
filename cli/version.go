package cli

import (
	"context"
	"encoding/json"
	"github.com/lefeck/go-bigip"
)

type VersionStats struct {
	Kind     string           `json:"kind,omitempty"`
	SelfLink string           `json:"selfLink,omitempty"`
	Entries  map[string]Entry `json:"entries,omitempty"`
}

type Entry struct {
	NestedStats NestedStats `json:"nestedStats,omitempty"`
}

type NestedStats struct {
	EntriesMenu EntriesMenu `json:"entries,omitempty"`
}

type EntriesMenu struct {
	Active    Description `json:"active,omitempty"`
	Latest    Description `json:"latest,omitempty"`
	Supported Description `json:"supported,omitempty"`
}

type Description struct {
	Description string `json:"description,omitempty"`
}

type VersionStatsResoure struct {
	b *bigip.BigIP
}

// VersionEndpoint is the base path of the TM API.
const VersionEndpoint = "version"

// Show bigip device version
func (vsr *VersionStatsResoure) Show() (*VersionStats, error) {
	var vs *VersionStats
	res, err := vsr.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(CliManager).
		Resource(VersionEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &vs); err != nil {
		panic(err)
	}

	return vs, nil
}
