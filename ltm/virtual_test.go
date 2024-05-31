package ltm

import (
	"github.com/lefeck/go-bigip"
	"testing"
)

func TestVirtualResource(t *testing.T) {
	bigIP, err := bigip.NewSession("192.168.12.21", "admin", "1qaz@WSX3edc")
	if err != nil {
		t.Fatalf("connect to bigip failed: %v", err)
	}

	vResource := &VirtualResource{
		b: bigIP,
	}

	// Create a VirtualServer
	vs := VirtualServer{
		Name:        "test_virtual_server",
		Partition:   "Common",
		Destination: "10.0.0.1:80",
		Mask:        "255.255.255.255",
		SourceAddressTranslation: SourceAddressTranslation{
			Type: "automap",
		},
	}
	if err := vResource.Create(vs); err != nil {
		t.Fatalf("Error creating VirtualServer: %v", err)
	}

	// Get the VirtualServer
	vsCheck, err := vResource.Get("/Common/test_virtual_server")
	if err != nil {
		t.Fatalf("Error getting VirtualServer: %v", err)
	}

	// Validate some properties
	if vsCheck.Name != "test_virtual_server" {
		t.Error("Name of VirtualServer is not correct")
	}

	// Update properties of the VirtualServer
	vsCheck.Description = "Test VirtualServer"
	if err := vResource.Update(vsCheck.FullPath, *vsCheck); err != nil {
		t.Fatalf("Error updating VirtualServer: %v", err)
	}

	// Validate that properties were updated
	vsUpdated, err := vResource.Get("/Common/test_virtual_server")
	if err != nil {
		t.Fatalf("Error getting updated VirtualServer: %v", err)
	}
	if vsUpdated.Description != "Test VirtualServer" {
		t.Error("Failed to update Description of VirtualServer")
	}

	// Disable the VirtualServer
	if err := vResource.Disable(vsCheck.FullPath); err != nil {
		t.Fatalf("Error disabling VirtualServer: %v", err)
	}

	// Enable the VirtualServer
	if err := vResource.Enable(vsCheck.FullPath); err != nil {
		t.Fatalf("Error enabling VirtualServer: %v", err)
	}

	// Delete the VirtualServer
	if err := vResource.Delete("/Common/test_virtual_server"); err != nil {
		t.Fatalf("Error deleting VirtualServer: %v", err)
	}

	// Validate that VirtualServer was deleted
	if _, err := vResource.Get("/Common/test_virtual_server"); err == nil {
		t.Error("VirtualServer still exists after attempting deletion")
	}
}
