package cli

import (
	"github.com/lefeck/go-bigip"
	"testing"
)

func TestVersionStatsResource(t *testing.T) {
	bigIP, err := bigip.NewSession("192.168.12.21", "admin", "1qaz@WSX3edc")
	if err != nil {
		t.Fatalf("connect to bigip failed: %v", err)
	}
	versionStatsResource := VersionStatsResoure{
		b: bigIP,
	}

	// Show BIG-IP device version
	versionStats, err := versionStatsResource.Show()
	if err != nil {
		t.Fatalf("Error getting BIG-IP device version: %v", err)
	}

	// Validate properties
	entries := versionStats.Entries
	if entries == nil {
		t.Error("Entries of VersionStats is nil")
	}

	for _, entry := range entries {
		if entry.NestedStats.EntriesMenu.Active.Description == "" {
			t.Error("Active version description is empty")
		}
		if entry.NestedStats.EntriesMenu.Latest.Description == "" {
			t.Error("Latest version description is empty")
		}
		if entry.NestedStats.EntriesMenu.Supported.Description == "" {
			t.Error("Supported version description is empty")
		}
	}
}
