package ltm

import (
	"github.com/lefeck/go-bigip"
	"testing"
)

func TestSnatTranslationstateResource(t *testing.T) {
	bigIP, err := bigip.NewSession("192.168.12.21", "admin", "1qaz@WSX3edc")
	if err != nil {
		t.Fatalf("connect to bigip failed: %v", err)
	}

	snatTranslationResource := SnatTranslationResource{
		b: bigIP,
	}

	fullPathName := "/Common/test_snat_translation_90"

	// Create a new SnatTranslationstate
	snatTranslationstate := SnatTranslation{
		Name:            "test_snat_translation_90",
		Partition:       "Common",
		Address:         "152.16.11.11",
		ConnectionLimit: 10,
	}
	if err := snatTranslationResource.Create(snatTranslationstate); err != nil {
		t.Fatalf("Error creating SnatTranslationstate: %v", err)
	}

	// Get the SnatTranslationstate
	snatTranslationstateCheck, err := snatTranslationResource.Get(fullPathName)
	if err != nil {
		t.Fatalf("Error getting SnatTranslationstate: %v", err)
	}

	// Validate some properties
	if snatTranslationstateCheck.Name != "test_snat_translation_90" {
		t.Error("Name of SnatTranslationstate is not correct")
	}

	// Update the SnatTranslationstate
	snatTranslationstateCheck.ConnectionLimit = 123
	if err := snatTranslationResource.Update(fullPathName, *snatTranslationstateCheck); err != nil {
		t.Fatalf("Error updating SnatTranslationstate: %v", err)
	}

	// Validate that properties were updated
	snatTranslationstateUpdated, err := snatTranslationResource.Get(fullPathName)
	if err != nil {
		t.Fatalf("Error getting updated SnatTranslationstate: %v", err)
	}
	if snatTranslationstateUpdated.ConnectionLimit != 123 {
		t.Error("Failed to update ConnectionLimit  of SnatTranslationstate")
	}

	// Disable the SnatTranslation
	if err := snatTranslationResource.Disable(fullPathName); err != nil {
		t.Fatalf("Error disabling SnatTranslation: %v", err)
	}

	// Validate that SnatTranslation was disabled
	snatTranslationStat, err := snatTranslationResource.Get(fullPathName)
	if err != nil {
		t.Error("SnatTranslation still exists after attempting deletion")
	}

	if snatTranslationStat.Disabled != true {
		t.Error("Failed to update Disabling status of SnatTranslation")
	}

	// Enable the SnatTranslation
	if err := snatTranslationResource.Enable(fullPathName); err != nil {
		t.Fatalf("Error Enabling SnatTranslation: %v", err)
	}

	// Validate that SnatTranslation was enbled
	snatTranslationStat, err = snatTranslationResource.Get(fullPathName)
	if err != nil {
		t.Error("SnatTranslation still exists after attempting deletion")
	}

	if snatTranslationStat.Enabled != true {
		t.Error("Failed to update Enabling status of SnatTranslation")
	}

	// Delete the SnatTranslationstate
	if err := snatTranslationResource.Delete(fullPathName); err != nil {
		t.Fatalf("Error deleting SnatTranslationstate: %v", err)
	}

	// Validate that SnatTranslationstate was deleted
	if _, err := snatTranslationResource.Get(fullPathName); err == nil {
		t.Error("SnatTranslationstate still exists after attempting deletion")
	}
}
