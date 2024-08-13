package auth

import (
	"fmt"
	"testing"

	"github.com/lefeck/go-bigip"
)

func TestUsersResource(t *testing.T) {
	bigIP, err := bigip.NewSession("192.168.13.91", "admin", "1qaz@WSX3edc")
	if err != nil {
		t.Fatalf("connect to bigip failed: %v", err)
	}

	ur := UsersResource{b: bigIP}

	ul, err := ur.List()
	fmt.Println(ul)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(ul.Item) == 0 {
		t.Fatalf("expected at least 1 user, got %v", len(ul.Item))
	}

	err = ur.Create(User{Name: "newuser", Description: "New User"})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	user, err := ur.Get("newuser")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if user.Name != "newuser" {
		t.Fatalf("expected user with name 'testuser', got %v", user.Name)
	}

	pa := []PartitionAccess{{Name: "all-partitions", Role: "admin"}}
	err = ur.Update("newuser", User{Shell: "tmsh", Password: "ComplexPassword123!", Description: "Updated User", PartitionAccess: pa})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	err = ur.Delete("newuser")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if user.Name == "" {
		t.Fatalf("expected user with name 'testuser', got %v", user.Name)
	}

}
