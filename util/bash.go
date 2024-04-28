package util

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/lefeck/bigip"
	"github.com/lefeck/bigip/ltm"
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

// https://192.168.13.91/mgmt/tm/util/bash

/*
curl -k -u admin:MsTac@2001 -H "Content-Type: application/json" -X POST -d \
'{"command": "run", "utilCmdArgs": " -c 'ls -l'"}' https://192.168.13.91/mgmt/tm/util/bash

curl -k -u admin:MsTac@2001 -H "Content-Type: application/json" -X POST -d \
'{"command": "run", "utilCmdArgs": "-c uptime"}' https://192.168.13.91/mgmt/tm/util/bash
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
		return nil, errors.New("Enter the query command please")
	}

	newitem := ValidateCmdArgs(item)

	jsonData, err := json.Marshal(newitem)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	res, err := br.b.RestClient.Post().Prefix(ltm.BasePath).ResourceCategory(ltm.TMResource).
		ManagerName(ltm.UTILManager).Resource(ltm.BashEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &bash); err != nil {
		return nil, fmt.Errorf("failed to marshal JSON data: %w", err)
	}

	return &bash, nil
}

func ValidateCmdArgs(item Bash) Bash {
	cmdArgs := strings.TrimSpace(item.UtilCmdArgs)
	if !strings.HasPrefix(cmdArgs, "-c") {
		item.UtilCmdArgs = fmt.Sprintf("-c '%s'", cmdArgs)
	}
	return item
}
