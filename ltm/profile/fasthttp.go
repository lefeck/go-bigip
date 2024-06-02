package profile

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

type FastHTTPList struct {
	Items    []FastHTTP `json:"items,omitempty"`
	Kind     string     `json:"kind,omitempty"`
	SelfLink string     `json:"selflink,omitempty"`
}

type FastHTTP struct {
	Kind                        string `json:"kind,omitempty"`
	Name                        string `json:"name,omitempty"`
	Partition                   string `json:"partition,omitempty"`
	FullPath                    string `json:"fullPath,omitempty"`
	Generation                  int    `json:"generation,omitempty"`
	SelfLink                    string `json:"selfLink,omitempty"`
	AppService                  string `json:"appService,omitempty"`
	ClientCloseTimeout          int    `json:"clientCloseTimeout,omitempty"`
	ConnpoolIdleTimeoutOverride int    `json:"connpoolIdleTimeoutOverride,omitempty"`
	ConnpoolMaxReuse            int    `json:"connpoolMaxReuse,omitempty"`
	ConnpoolMaxSize             int    `json:"connpoolMaxSize,omitempty"`
	ConnpoolMinSize             int    `json:"connpoolMinSize,omitempty"`
	ConnpoolReplenish           string `json:"connpoolReplenish,omitempty"`
	ConnpoolStep                int    `json:"connpoolStep,omitempty"`
	DefaultsFrom                string `json:"defaultsFrom,omitempty"`
	Description                 string `json:"description,omitempty"`
	ForceHTTP10Response         string `json:"forceHttp_10Response,omitempty"`
	HardwareSynCookie           string `json:"hardwareSynCookie,omitempty"`
	HeaderInsert                string `json:"headerInsert,omitempty"`
	HTTP11CloseWorkarounds      string `json:"http_11CloseWorkarounds,omitempty"`
	IdleTimeout                 int    `json:"idleTimeout,omitempty"`
	InsertXforwardedFor         string `json:"insertXforwardedFor,omitempty"`
	Layer7                      string `json:"layer_7,omitempty"`
	MaxHeaderSize               int    `json:"maxHeaderSize,omitempty"`
	MaxRequests                 int    `json:"maxRequests,omitempty"`
	MssOverride                 int    `json:"mssOverride,omitempty"`
	ReceiveWindowSize           int    `json:"receiveWindowSize,omitempty"`
	ResetOnTimeout              string `json:"resetOnTimeout,omitempty"`
	ServerCloseTimeout          int    `json:"serverCloseTimeout,omitempty"`
	ServerSack                  string `json:"serverSack,omitempty"`
	ServerTimestamp             string `json:"serverTimestamp,omitempty"`
	UncleanShutdown             string `json:"uncleanShutdown,omitempty"`
}

const FastHTTPEndpoint = "fasthttp"

type FastHTTPResource struct {
	b *bigip.BigIP
}

func (fr *FastHTTPResource) List() (*FastHTTPList, error) {
	var items FastHTTPList
	res, err := fr.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(FastHTTPEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &items); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &items, nil
}

func (fr *FastHTTPResource) Get(fullPathName string) (*FastHTTP, error) {
	var item FastHTTP
	res, err := fr.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(FastHTTPEndpoint).SubResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &item, nil
}

func (fr *FastHTTPResource) Create(item FastHTTP) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = fr.b.RestClient.Post().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(FastHTTPEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (fr *FastHTTPResource) Update(fullPathName string, item FastHTTP) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = fr.b.RestClient.Put().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(FastHTTPEndpoint).SubResourceInstance(fullPathName).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (fr *FastHTTPResource) Delete(fullPathName string) error {
	_, err := fr.b.RestClient.Delete().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(FastHTTPEndpoint).SubResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
