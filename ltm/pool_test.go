package ltm

import (
	"github.com/lefeck/go-bigip"
	"testing"
)

func TestPoolResource(t *testing.T) {
	bigIP, err := bigip.NewSession("192.168.12.21", "admin", "1qaz@WSX3edc")
	if err != nil {
		t.Fatalf("connect to bigip failed: %v", err)
	}
	poolResource := PoolResource{
		b: bigIP,
	}

	// Create a new Pool
	pool := Pool{
		Name:              "test-pool",
		Partition:         "Common",
		LoadBalancingMode: "round-robin",
	}
	if err := poolResource.Create(pool); err != nil {
		t.Fatalf("Error creating Pool: %v", err)
	}

	// Get the Pool
	poolCheck, err := poolResource.Get("/Common/test-pool")
	if err != nil {
		t.Fatalf("Error getting Pool: %v", err)
	}

	// Validate some properties
	if poolCheck.Name != "test-pool" {
		t.Error("Name of Pool is not correct")
	}

	// Update properties of the Pool
	poolCheck.AllowNat = "yes"
	poolCheck.AllowSnat = "yes"
	if err := poolResource.Update(poolCheck.FullPath, *poolCheck); err != nil {
		t.Fatalf("Error updating Pool: %v", err)
	}

	// Validate that properties were updated
	poolUpdated, err := poolResource.Get("/Common/test-pool")
	if err != nil {
		t.Fatalf("Error getting updated Pool: %v", err)
	}
	if poolUpdated.AllowNat != "yes" || poolUpdated.AllowSnat != "yes" {
		t.Error("Failed to update AllowNat and AllowSnat of Pool")
	}

	// Delete the Pool
	if err := poolResource.Delete(poolCheck.FullPath); err != nil {
		t.Fatalf("Error deleting Pool: %v", err)
	}

	// Validate that Pool was deleted
	if _, err := poolResource.Get("/Common/test-pool"); err == nil {
		t.Error("Pool still exists after attempting deletion")
	}
}
