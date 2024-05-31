package ltm

import (
	"testing"

	"github.com/lefeck/go-bigip"
)

func TestTrafficMatchingCriteriaResource(t *testing.T) {
	bigIP, err := bigip.NewSession("192.168.13.91", "admin", "MsTac@2001")
	if err != nil {
		t.Fatalf("connect to bigip failed: %v", err)
	}

	tcResource := &TrafficMatchingCriteriaResource{
		b: bigIP,
	}

	// Create a TrafficMatchingCriteria
	tc := TrafficMatchingCriteria{
		Name:                     "test_traffic_matching_criteria",
		Partition:                "Common",
		Protocol:                 "tcp",
		DestinationAddressInline: "192.168.1.1",

		DestinationPortInline: "0",
		DestinationPortList:   "/Common/_sys_self_allow_udp_defaults",

		RouteDomain:         "any",
		SourceAddressInline: "0.0.0.0",
		SourceAddressList:   "/Common/pp2",
		SourcePortInline:    0,
	}
	if err := tcResource.Create(tc); err != nil {
		t.Fatalf("Error creating TrafficMatchingCriteria: %v", err)
	}

	// Get the TrafficMatchingCriteria
	tcCheck, err := tcResource.Get("/Common/test_traffic_matching_criteria")
	if err != nil {
		t.Fatalf("Error getting TrafficMatchingCriteria: %v", err)
	}

	// Validate some properties
	if tcCheck.Name != "test_traffic_matching_criteria" {
		t.Error("Name of TrafficMatchingCriteria is not correct")
	}

	// Update properties of the TrafficMatchingCriteria
	tcCheck.DestinationPortInline = "0"
	if err := tcResource.Update(tcCheck.FullPath, *tcCheck); err != nil {
		t.Fatalf("Error updating TrafficMatchingCriteria: %v", err)
	}

	// Validate that properties were updated
	tcUpdated, err := tcResource.Get("/Common/test_traffic_matching_criteria")
	if err != nil {
		t.Fatalf("Error getting updated TrafficMatchingCriteria: %v", err)
	}
	if tcUpdated.DestinationPortInline != "0" {
		t.Error("Failed to update DestinationPortInline of TrafficMatchingCriteria")
	}

	// Delete the TrafficMatchingCriteria
	if err := tcResource.Delete("/Common/test_traffic_matching_criteria"); err != nil {
		t.Fatalf("Error deleting TrafficMatchingCriteria: %v", err)
	}

	// Validate that TrafficMatchingCriteria was deleted
	if _, err := tcResource.Get("/Common/test_traffic_matching_criteria"); err == nil {
		t.Error("TrafficMatchingCriteria still exists after attempting deletion")
	}
}
