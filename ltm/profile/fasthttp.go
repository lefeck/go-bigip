package profile

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

type FasthttpList struct {
	Items    []Fasthttp `json:"items,omitempty"`
	Kind     string     `json:"kind,omitempty"`
	SelfLink string     `json:"selflink,omitempty"`
}

type Fasthttp struct {
	Kind                        string `json:"kind"`
	Name                        string `json:"name"`
	Partition                   string `json:"partition"`
	FullPath                    string `json:"fullPath"`
	Generation                  int    `json:"generation"`
	SelfLink                    string `json:"selfLink"`
	AppService                  string `json:"appService"`
	ClientCloseTimeout          int    `json:"clientCloseTimeout"`
	ConnpoolIdleTimeoutOverride int    `json:"connpoolIdleTimeoutOverride"`
	ConnpoolMaxReuse            int    `json:"connpoolMaxReuse"`
	ConnpoolMaxSize             int    `json:"connpoolMaxSize"`
	ConnpoolMinSize             int    `json:"connpoolMinSize"`
	ConnpoolReplenish           string `json:"connpoolReplenish"`
	ConnpoolStep                int    `json:"connpoolStep"`
	DefaultsFrom                string `json:"defaultsFrom"`
	Description                 string `json:"description"`
	ForceHTTP10Response         string `json:"forceHttp_10Response"`
	HardwareSynCookie           string `json:"hardwareSynCookie"`
	HeaderInsert                string `json:"headerInsert"`
	HTTP11CloseWorkarounds      string `json:"http_11CloseWorkarounds"`
	IdleTimeout                 int    `json:"idleTimeout"`
	InsertXforwardedFor         string `json:"insertXforwardedFor"`
	Layer7                      string `json:"layer_7"`
	MaxHeaderSize               int    `json:"maxHeaderSize"`
	MaxRequests                 int    `json:"maxRequests"`
	MssOverride                 int    `json:"mssOverride"`
	ReceiveWindowSize           int    `json:"receiveWindowSize"`
	ResetOnTimeout              string `json:"resetOnTimeout"`
	ServerCloseTimeout          int    `json:"serverCloseTimeout"`
	ServerSack                  string `json:"serverSack"`
	ServerTimestamp             string `json:"serverTimestamp"`
	UncleanShutdown             string `json:"uncleanShutdown"`
}

const FasthttpEndpoint = "fasthttp"

type FasthttpResource struct {
	b *bigip.BigIP
}

func (fr *FasthttpResource) List() (*FasthttpList, error) {
	var items FasthttpList
	res, err := fr.b.RestClient.Get().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(FasthttpEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &items); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &items, nil
}

func (fr *FasthttpResource) Get(fullPathName string) (*Fasthttp, error) {
	var item Fasthttp
	res, err := fr.b.RestClient.Get().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(FasthttpEndpoint).SubResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &item, nil
}

func (fr *FasthttpResource) Create(item Fasthttp) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = fr.b.RestClient.Post().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(FasthttpEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (fr *FasthttpResource) Update(fullPathName string, item Fasthttp) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = fr.b.RestClient.Put().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(FasthttpEndpoint).SubResourceInstance(fullPathName).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (fr *FasthttpResource) Delete(fullPathName string) error {
	_, err := fr.b.RestClient.Delete().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(FasthttpEndpoint).SubResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
