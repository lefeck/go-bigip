package ltm

import (
	"github.com/lefeck/go-bigip"

	"testing"
)

func TestPoolMembersResource(t *testing.T) {
	bigIP, err := bigip.NewSession("192.168.12.21", "admin", "1qaz@WSX3edc")
	if err != nil {
		t.Fatalf("connect to bigip failed: %v", err)
	}
	poolResource := PoolResource{
		b: bigIP,
	}

	pool := Pool{
		Name:              "test-pool-members",
		Partition:         "Common",
		LoadBalancingMode: "round-robin",
	}

	err = poolResource.Create(pool)
	if err != nil {
		t.Fatalf("Error creating Pool: %v", err)
	}

	poolMemberResource := PoolMembersResource{
		b: bigIP,
	}

	// Add a Pool Member
	poolMember := PoolMembers{
		Name:    "192.168.1.1:80",
		Address: "192.168.1.1",
	}

	err = poolMemberResource.Create("test-pool-members", poolMember)
	if err != nil {
		t.Fatalf("Error creating Pool Member: %v", err)
	}

	// Get the Pool Member
	poolMemberCheck, err := poolMemberResource.Get("test-pool-members", poolMember.Name)
	if err != nil {
		t.Fatalf("Error getting Pool Member: %v", err)
	}

	// Validate properties
	if poolMemberCheck.Name != "192.168.1.1:80" {
		t.Error("Name of Pool Member is not correct")
	}

	// Update properties of the Pool Member
	poolMemberCheck.ConnectionLimit = 100
	err = poolMemberResource.Update("test-pool-members", poolMemberCheck.Name, *poolMemberCheck)
	if err != nil {
		t.Fatalf("Error updating Pool Member: %v", err)
	}

	// Validate that properties were updated
	updatedPoolMember, err := poolMemberResource.Get("test-pool-members", poolMemberCheck.Name)
	if err != nil {
		t.Fatalf("Error getting updated Pool Member: %v", err)
	}
	if updatedPoolMember.ConnectionLimit != 100 {
		t.Error("Failed to update ConnectionLimit of Pool Member")
	}

	// Delete the Pool Member
	err = poolMemberResource.Delete("test-pool-members", poolMember.Name)
	if err != nil {
		t.Fatalf("Error deleting Pool Member: %v", err)
	}

	// Validate that Pool Member was deleted
	if _, err := poolMemberResource.Get("test-pool-members", poolMemberCheck.Name); err == nil {
		t.Error("Pool Member still exists after attempting deletion")
	}
}
