package util

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

type Bash struct {
	Command       string `json:"command"`
	UtilCmdArgs   string `json:"utilCmdArgs"`
	CommandResult string `json:"commandResult"`
}

type BashResource struct {
	b *bigip.BigIP
}

// BashEndpoint is the base path of the util API.
const BashEndpoint = "bash"

/*
Provides a way to manipulate bigip via commands，for example:

	"tmsh list ltm virtual"
	"uptime"
*/
func (br *BashResource) Run(item Bash) (*Bash, error) {

	var bash Bash

	if item.Command == "" {
		item.Command = "run"
	}

	if item.Command != "run" {
		return nil, errors.New("The command you entered is not supported, only the run parameter is supported by default.")
	}

	if len(item.UtilCmdArgs) == 0 {
		return nil, errors.New("The input execution command cannot be empty")
	}

	newitem := validateCmdArgs(item)

	jsonData, err := json.Marshal(newitem)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	res, err := br.b.RestClient.Post().Prefix(bigip.BasePath).ResourceCategory(bigip.TMResource).
		ManagerName(UtilManager).Resource(BashEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &bash); err != nil {
		return nil, fmt.Errorf("failed to marshal JSON data: %w", err)
	}

	return &bash, nil
}

func validateCmdArgs(item Bash) Bash {
	cmdArgs := strings.TrimSpace(item.UtilCmdArgs)
	// 	for example : tmsh list ltm virtual",
	if !strings.HasPrefix(cmdArgs, "-c") {
		item.UtilCmdArgs = fmt.Sprintf("-c '%s'", cmdArgs)
		return item
	}
	// 	for example:  "-c 'tmsh list  ltm  virtual'",
	item.UtilCmdArgs = validateCmdArgsForBash(cmdArgs)

	return item
}

func validateCmdArgsForBash(cmd string) string {
	last := strings.TrimPrefix(cmd, "-c")
	last = strings.TrimSpace(last)
	if !strings.HasPrefix(last, "'") || !strings.HasSuffix(last, "'") {
		last = "'" + last + "'"
	}
	cmd = "-c " + last
	return cmd
}
