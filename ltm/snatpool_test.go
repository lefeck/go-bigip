package ltm

import (
	"github.com/lefeck/go-bigip"
	"testing"
)

func TestSnatPoolResource(t *testing.T) {
	bigIP, err := bigip.NewSession("192.168.12.21", "admin", "1qaz@WSX3edc")
	if err != nil {
		t.Fatalf("connect to bigip failed: %v", err)
	}

	snatPoolResource := SnatPoolResource{
		b: bigIP,
	}

	fullPathName := "/Common/test_snat_pool_90"

	// Create a new SnatPool
	snatPool := SnatPool{
		Name:      "test_snat_pool_90",
		Partition: "Common",

		Members: []string{"192.168.1.10"},
	}
	if err := snatPoolResource.Create(snatPool); err != nil {
		t.Fatalf("Error creating SnatPool: %v", err)
	}

	// Get the SnatPool
	snatPoolCheck, err := snatPoolResource.Get(fullPathName)
	if err != nil {
		t.Fatalf("Error getting SnatPool: %v", err)
	}

	// Validate some properties
	if snatPoolCheck.Name != "test_snat_pool_90" {
		t.Error("Name of SnatPool is not correct")
	}

	// Update the SnatPool
	snatPoolCheck.Members = []string{"192.168.1.20"}
	if err := snatPoolResource.Update(fullPathName, *snatPoolCheck); err != nil {
		t.Fatalf("Error updating SnatPool: %v", err)
	}

	// Validate that properties were updated
	snatPoolUpdated, err := snatPoolResource.Get(fullPathName)
	if err != nil {
		t.Fatalf("Error getting updated SnatPool: %v", err)
	}
	if len(snatPoolUpdated.Members) != 1 || snatPoolUpdated.Members[0] != "192.168.1.20" {
		t.Error("Failed to update Members of SnatPool")
	}

	// Delete the SnatPool
	if err := snatPoolResource.Delete(fullPathName); err != nil {
		t.Fatalf("Error deleting SnatPool: %v", err)
	}

	// Validate that SnatPool was deleted
	if _, err := snatPoolResource.Get(fullPathName); err == nil {
		t.Error("SnatPool still exists after attempting deletion")
	}
}
