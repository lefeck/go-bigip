package ltm

import (
	"github.com/lefeck/go-bigip"

	"testing"
)

func TestIFileResource(t *testing.T) {
	bigIP, err := bigip.NewSession("192.168.12.21", "admin", "1qaz@WSX3edc")
	if err != nil {
		t.Fatalf("connect to bigip failed: %v", err)
	}
	iFileResource := IFileResource{
		b: bigIP,
	}

	// Create an iFile
	fileObject := "./file.txt"
	err = iFileResource.Create("test-ifile", fileObject)
	if err != nil {
		t.Fatalf("Error creating iFile: %v", err)
	}

	// Get the iFile
	iFileCheck, err := iFileResource.Get("/Common/test-ifile")
	if err != nil {
		t.Fatalf("Error getting iFile: %v", err)
	}

	// Validate properties
	if iFileCheck.Name != "test-ifile" {
		t.Error("Name of iFile is not correct")
	}

	// Update properties of the iFile
	fileObjectUpdated := "./updated-file.txt"
	err = iFileResource.Edit("test-ifile", fileObjectUpdated)
	if err != nil {
		t.Fatalf("Error updating iFile: %v", err)
	}

	// Validate that properties were updated
	updatedIFile, err := iFileResource.Get("/Common/test-ifile")
	if err != nil {
		t.Fatalf("Error getting updated iFile: %v", err)
	}
	if updatedIFile.FileName != fileObjectUpdated {
		t.Error("Failed to update the filename of iFile")
	}

	// Delete the iFile
	err = iFileResource.Delete("test-ifile")
	if err != nil {
		t.Fatalf("Error deleting iFile: %v", err)
	}

	// Validate that iFile was deleted
	if _, err := iFileResource.Get("/Common/test-ifile"); err == nil {
		t.Error("iFile still exists after attempting deletion")
	}
}
