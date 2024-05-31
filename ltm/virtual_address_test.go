package ltm

import (
	"github.com/lefeck/go-bigip"
	"testing"
)

func TestVirtualAddressResource(t *testing.T) {
	bigIP, err := bigip.NewSession("192.168.12.21", "admin", "1qaz@WSX3edc")
	if err != nil {
		t.Fatalf("connect to bigip failed: %v", err)
	}

	virtualAddressResource := VirtualAddressResource{
		b: bigIP,
	}

	// List all Virtual Addresses
	virtualAddressList, err := virtualAddressResource.List()
	if err != nil {
		t.Fatalf("Error getting VirtualAddress list: %v", err)
	}

	// Get the first Virtual Address
	if len(virtualAddressList.Items) > 0 {
		virtualAddress, err := virtualAddressResource.Get(virtualAddressList.Items[0].FullPath)
		if err != nil {
			t.Fatalf("Error getting Virtual Address: %v", err)
		}

		// Validate properties
		if virtualAddress.Name == "" {
			t.Error("Virtual Address name is empty")
		}

		// Update Virtual Address
		virtualAddress.Arp = "enabled"
		if err := virtualAddressResource.Update(virtualAddress.Name, *virtualAddress); err != nil {
			t.Fatalf("Error updating Virtual Address: %v", err)
		}

		// Check if ARP property was updated
		updatedVASingle, err := virtualAddressResource.Get(virtualAddress.Name)
		if err != nil {
			t.Fatalf("Error getting updated Virtual Address: %v", err)
		}
		if updatedVASingle.Arp != "enabled" {
			t.Error("Failed to update ARP property of Virtual Address")
		}

		// Disable Virtual Address
		if err := virtualAddressResource.Disable(virtualAddress.Name); err != nil {
			t.Fatalf("Error disabling Virtual Address: %v", err)
		}

		// Enable Virtual Address
		if err := virtualAddressResource.Enable(virtualAddress.Name); err != nil {
			t.Fatalf("Error enabling Virtual Address: %v", err)
		}
	}
}
