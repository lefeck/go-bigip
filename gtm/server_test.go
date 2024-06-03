package gtm

import (
	"github.com/lefeck/go-bigip"

	"testing"
)

func TestServerResource(t *testing.T) {
	bigIP, err := bigip.NewSession("192.168.12.21", "admin", "1qaz@WSX3edc")
	if err != nil {
		t.Fatalf("connect to bigip failed: %v", err)
	}

	// Create a Datacenter item
	datacenter := Datacenter{
		Name:      "test-datacenter",
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

	serverResource := ServerResource{
		b: bigIP,
	}

	// Create a new Server
	newServer := Server{
		Name:       "test-server",
		Datacenter: "/Common/test-datacenter",
		Enabled:    true,
		Addresses: []struct {
			DeviceName  string `json:"deviceName,omitempty"`
			Name        string `json:"name,omitempty"`
			Translation string `json:"translation,omitempty"`
		}{
			{DeviceName: "Device1", Name: "192.16.65.123", Translation: "192.168.53.1"},
		},
	}

	err = serverResource.Create(newServer)
	if err != nil {
		t.Fatalf("Error creating server: %v", err)
	}

	// Get the created Server
	server, err := serverResource.Get("/Common/test-server")
	if err != nil {
		t.Fatalf("Error getting server: %v", err)
	}

	// Validate properties
	if server.Name != "test-server" {
		t.Error("Server name is incorrect")
	}
	if server.Datacenter != "/Common/test-datacenter" {
		t.Error("Server datacenter is incorrect")
	}

	// Update the Server
	server.Enabled = false
	err = serverResource.Update("test-server", *server)
	if err != nil {
		t.Fatalf("Error updating server: %v", err)
	}

	// Validate updated properties
	updatedServer, err := serverResource.Get("/Common/test-server")
	if err != nil {
		t.Fatalf("Error getting updated server: %v", err)
	}
	if !updatedServer.Enabled {
		t.Error("Server Enabled property not updated")
	}

	// Delete the Server
	err = serverResource.Delete("test-server")
	if err != nil {
		t.Fatalf("Error deleting server: %v", err)
	}

	// Validate that the Server was deleted
	if _, err := serverResource.Get("/Common/test-server"); err == nil {
		t.Error("Server still exists after attempting deletion")
	}
}
