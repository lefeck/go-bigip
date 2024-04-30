package main

import (
	"fmt"
	"strings"
)

func validateCmdForBash(cmd string) string {
	cmd = strings.TrimSpace(cmd)
	if strings.HasPrefix(cmd, "-c ") {
		last := strings.TrimPrefix(cmd, "-c ")
		last = strings.TrimSpace(last)
		if !strings.HasPrefix(last, "'") || !strings.HasSuffix(last, "'") {
			last = "'" + last + "'"
		}
		cmd = "-c " + last
	}
	return cmd
}

func main() {
	cmd := "-c tmsh list  ltm  virtual"
	checkedCmd := validateCmdForBash(cmd)
	fmt.Println(checkedCmd)
}
