package ltm

import (
	"github.com/lefeck/go-bigip"

	"testing"
)

func TestRuleResource(t *testing.T) {
	bigIP, err := bigip.NewSession("192.168.12.21", "admin", "1qaz@WSX3edc")
	if err != nil {
		t.Fatalf("connect to bigip failed: %v", err)
	}
	ruleResource := RuleResource{
		b: bigIP,
	}

	ruleScript := "when HTTP_REQUEST {HTTP::respond 200 content \"Hello, iRule!\"}"
	// Create an iRule
	rule := Rule{
		Name:         "test-rule",
		Partition:    "Common",
		ApiAnonymous: ruleScript,
	}

	err = ruleResource.Create(rule)
	if err != nil {
		t.Fatalf("Error creating iRule: %v", err)
	}

	// Get the iRule
	ruleCheck, err := ruleResource.Get("/Common/test-rule")
	if err != nil {
		t.Fatalf("Error getting iRule: %v", err)
	}

	// Validate properties
	if ruleCheck.Name != "test-rule" {
		t.Error("Name of iRule is not correct")
	}

	updateruleScript := "when HTTP_REQUEST {HTTP::respond 200 content \"Hello, updated iRule!\"}"
	// Update properties of the iRule
	ruleCheck.ApiAnonymous = updateruleScript
	err = ruleResource.Update("test-rule", *ruleCheck)
	if err != nil {
		t.Fatalf("Error updating iRule: %v", err)
	}

	// Validate that properties were updated
	updatedRule, err := ruleResource.Get("/Common/test-rule")
	if err != nil {
		t.Fatalf("Error getting updated iRule: %v", err)
	}
	if updatedRule.ApiAnonymous != updateruleScript {
		t.Error("Failed to update ApiAnonymous of iRule")
	}

	// Delete the iRule
	err = ruleResource.Delete("test-rule")
	if err != nil {
		t.Fatalf("Error deleting iRule: %v", err)
	}

	// Validate that iRule was deleted
	if _, err := ruleResource.Get("/Common/test-rule"); err == nil {
		t.Error("iRule still exists after attempting deletion")
	}
}
