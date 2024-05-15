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

const FastHTTPEndpoint = "fasthttp"

type FastHTTPResource struct {
	b *bigip.BigIP
}

func (fr *FastHTTPResource) List() (*FastHTTPList, error) {
	var items FastHTTPList
	res, err := fr.b.RestClient.Get().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
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
	res, err := fr.b.RestClient.Get().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
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
	_, err = fr.b.RestClient.Post().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
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
	_, err = fr.b.RestClient.Put().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(FastHTTPEndpoint).SubResourceInstance(fullPathName).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (fr *FastHTTPResource) Delete(fullPathName string) error {
	_, err := fr.b.RestClient.Delete().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(FastHTTPEndpoint).SubResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
