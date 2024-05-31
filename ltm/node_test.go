package ltm

import (
	"github.com/lefeck/go-bigip"

	"testing"
)

func TestNodeResource(t *testing.T) {
	bigIP, err := bigip.NewSession("192.168.12.21", "admin", "1qaz@WSX3edc")
	if err != nil {
		t.Fatalf("connect to bigip failed: %v", err)
	}
	nodeResource := NodeResource{
		b: bigIP,
	}
	name := "test-node90"
	partition := "Common"
	fullPathName := "/" + partition + "/" + name
	// Create a new Node
	node := Node{
		Name:      name,
		Partition: partition,
		Address:   "152.16.11.11",
	}
	if err := nodeResource.Create(node); err != nil {
		t.Fatalf("Error creating Node: %v", err)
	}

	// Get the Node
	nodeCheck, err := nodeResource.Get(fullPathName)
	if err != nil {
		t.Fatalf("Error getting Node: %v", err)
	}

	// Validate some properties
	if nodeCheck.Name != name {
		t.Error("Name of Node is not correct")
	}

	////Update properties of the Node
	//nodeCheck.Address = "12.8.3.12"
	//if err := nodeResource.Update(nodeCheck.FullPath, *nodeCheck); err != nil {
	//	t.Fatalf("Error updating Node: %v", err)
	//}
	//
	//// Validate that properties were updated
	//nodeUpdated, err := nodeResource.Get(fullPathName)
	//if err != nil {
	//	t.Fatalf("Error getting updated Node: %v", err)
	//}
	//if nodeUpdated.Address != "12.8.3.12" {
	//	t.Error("Failed to update ConnectionLimit of Node")
	//}

	// Force Node offline
	if err := nodeResource.ForceOffline(nodeCheck.Name); err != nil {
		t.Fatalf("Error forcing Node offline: %v", err)
	}

	// Enable the Node
	if err := nodeResource.Enable(nodeCheck.Name); err != nil {
		t.Fatalf("Error enabling Node: %v", err)
	}

	// Disable the Node
	if err := nodeResource.Disable(nodeCheck.Name); err != nil {
		t.Fatalf("Error disabling Node: %v", err)
	}

	// Delete the Node
	if err := nodeResource.Delete(nodeCheck.Name); err != nil {
		t.Fatalf("Error deleting Node: %v", err)
	}

	// Validate that Node was deleted
	if _, err := nodeResource.Get(fullPathName); err == nil {
		t.Error("Node still exists after attempting deletion")
	}
}
