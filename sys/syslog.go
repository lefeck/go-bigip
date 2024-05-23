// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

// SyslogConfig holds the configuration of a single Syslog.
type Syslog struct {
	AuthPrivFrom         string         `json:"authPrivFrom,omitempty"`
	AuthPrivTo           string         `json:"authPrivTo,omitempty"`
	ClusteredHostSlot    string         `json:"clusteredHostSlot,omitempty"`
	ClusteredMessageSlot string         `json:"clusteredMessageSlot,omitempty"`
	ConsoleLog           string         `json:"consoleLog,omitempty"`
	CronFrom             string         `json:"cronFrom,omitempty"`
	CronTo               string         `json:"cronTo,omitempty"`
	DaemonFrom           string         `json:"daemonFrom,omitempty"`
	DaemonTo             string         `json:"daemonTo,omitempty"`
	IsoDate              string         `json:"isoDate,omitempty"`
	KernFrom             string         `json:"kernFrom,omitempty"`
	KernTo               string         `json:"kernTo,omitempty"`
	Kind                 string         `json:"kind,omitempty"`
	Local6From           string         `json:"local6From,omitempty"`
	Local6To             string         `json:"local6To,omitempty"`
	MailFrom             string         `json:"mailFrom,omitempty"`
	MailTo               string         `json:"mailTo,omitempty"`
	MessagesFrom         string         `json:"messagesFrom,omitempty"`
	MessagesTo           string         `json:"messagesTo,omitempty"`
	RemoteServers        []RemoteServer `json:"remoteServers,omitempty"`
	SelfLink             string         `json:"selfLink,omitempty"`
	UserLogFrom          string         `json:"userLogFrom,omitempty"`
	UserLogTo            string         `json:"userLogTo,omitempty"`
}

type RemoteServer struct {
	Host       string `json:"host,omitempty"`
	LocalIP    string `json:"localIp,omitempty"`
	Name       string `json:"name,omitempty"`
	RemotePort int    `json:"remotePort,omitempty"`
}

func (r *SyslogResource) AddRemoteServers(rs ...RemoteServer) error {
	if len(rs) == 0 {
		return nil
	}
	var item Syslog
	_, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(SyslogEndpoint).DoRaw(context.Background())
	if err != nil {
		return err
	}

	item.RemoteServers = append(item.RemoteServers, rs...)

	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = r.b.RestClient.Put().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(SyslogEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// SyslogEndpoint represents the REST resource for managing Syslog.
const SyslogEndpoint = "syslog"

// SyslogResource provides an API to manage Syslog configurations.
type SyslogResource struct {
	b *bigip.BigIP
}

// List all syslog details
func (r *SyslogResource) Get() (*Syslog, error) {
	var item Syslog
	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(SyslogEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &item, nil
}

// Update the syslog item identified by the syslog name, otherwise an error will be reported.
func (r *SyslogResource) Update(item Syslog) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = r.b.RestClient.Put().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(SyslogEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Delete a single syslog identified by the syslog name. if it is not exist return error
func (r *SyslogResource) Delete(name string) error {
	_, err := r.b.RestClient.Delete().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(SyslogEndpoint).ResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return err
	}

	return nil
}
