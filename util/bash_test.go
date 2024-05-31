package util

import (
	"github.com/lefeck/go-bigip"
	"testing"
)

func TestBashResource_Run(t *testing.T) {
	bigIP, err := bigip.NewSession("192.168.12.21", "admin", "1qaz@WSX3edc")
	if err != nil {
		t.Fatalf("connect to bigip failed: %v", err)
	}
	bashResource := BashResource{
		b: bigIP,
	}

	// Run a command on the BIG-IP device
	command := "tmsh list ltm virtual"
	bashCmd := Bash{
		UtilCmdArgs: command,
	}

	result, err := bashResource.Run(bashCmd)
	if err != nil {
		t.Fatalf("Error running command on BIG-IP device: %v", err)
	}

	// Validate command result
	if result.CommandResult == "" {
		t.Error("Empty command result")
	}
}
