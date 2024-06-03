package gtm

import (
	"github.com/lefeck/go-bigip"
	"testing"
)

func TestCRUDDatacenter(t *testing.T) {
	// Replace these variables with your F5 device details
	user := "admin"
	pass := "1qaz@WSX3edc"
	deviceIP := "192.168.12.21"

	// Connect to the device
	bigIP, err := bigip.NewSession(deviceIP, user, pass)
	if err != nil {
		t.Error(err)
	}

	// Create a Datacenter item
	datacenter := Datacenter{
		Name:      "Test_Datacenter",
		Partition: "Common",
		Contact:   "test@test.com",
		Enabled:   true,
		Location:  "Test location",
		//ProberFallback:   "any-available-member",
		//ProberPreference: "prefers-tgip-null",
	}

	gtm := New(bigIP)

	err = gtm.Datacenter().Create(datacenter)
	if err != nil {
		t.Error(err)
	}

	// List all Datacenters to ensure the Test_Datacenter is created.
	datacenterList, err := gtm.Datacenter().List()
	if err != nil {
		t.Error(err)
	}

	found := false
	for _, dc := range datacenterList.Items {
		if dc.Name == "Test_Datacenter" {
			found = true
			break
		}
	}
	if !found {
		t.Error("Test_Datacenter not found in the list of datacenters")
	}

	// Get Test_Datacenter details
	createdDatacenter, err := gtm.Datacenter().Get("Test_Datacenter")
	if err != nil {
		t.Error(err)
	}
	if datacenter.Name != createdDatacenter.Name {
		t.Errorf("Datacenter name mismatch: expected %s, got %s", datacenter.Name, createdDatacenter.Name)
	}

	// Update the Datacenter
	datacenter.Contact = "test2@test.com"
	err = gtm.Datacenter().Update(datacenter.Name, datacenter)
	if err != nil {
		t.Error(err)
	}

	// Check if the update was successful
	updatedDatacenter, err := gtm.Datacenter().Get("Test_Datacenter")
	if err != nil {
		t.Error(err)
	}
	if datacenter.Contact != updatedDatacenter.Contact {
		t.Errorf("Datacenter contact mismatch: expected %s, got %s", datacenter.Contact, updatedDatacenter.Contact)
	}

	// Delete the Test_Datacenter
	err = gtm.Datacenter().Delete(datacenter.Name)
	if err != nil {
		t.Error(err)
	}

	// Check if the Test_Datacenter was deleted
	deletedDatacenter, err := gtm.Datacenter().Get("Test_Datacenter")
	if err == nil {
		t.Error("Test_Datacenter is expected to be deleted but still exists")
	}
	if deletedDatacenter != nil {
		t.Error("Test_Datacenter is expected to be nil but it is not")
	}
}
