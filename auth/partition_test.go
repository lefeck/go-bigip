package auth

import (
	"fmt"
	"github.com/lefeck/go-bigip"
	"testing"
)

func TestPartitionResource(t *testing.T) {
	bigIP, err := bigip.NewSession("192.168.13.91", "admin", "1qaz@WSX3edc")
	if err != nil {
		t.Fatalf("connect to bigip failed: %v", err)
	}

	pr := PartitionResource{b: bigIP}

	pl, err := pr.List()
	fmt.Println(pl)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(pl.Item) == 0 {
		t.Fatalf("expected at least 1 user, got %v", len(pl.Item))
	}

	err = pr.Create(Partition{Name: "new_test", Description: "New Test"})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	partition, err := pr.Get("new_test")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if partition.Name != "new_test" {
		t.Fatalf("expected partition with name 'new_test', got %v", partition.Name)
	}

	pa := Partition{DefaultRouteDomain: 101}
	err = pr.Update("new_test", pa)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	err = pr.Delete("new_test")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if partition.Name == "" {
		t.Fatalf("expected partition with name 'new_test', got %v", partition.Name)
	}

}
