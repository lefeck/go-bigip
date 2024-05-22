package profile

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

type FTPList struct {
	Items    []FTP  `json:"items,omitempty"`
	Kind     string `json:"kind,omitempty"`
	SelfLink string `json:"selflink,omitempty"`
}

type FTP struct {
	Kind                   string `json:"kind"`
	Name                   string `json:"name"`
	Partition              string `json:"partition"`
	FullPath               string `json:"fullPath"`
	Generation             int    `json:"generation"`
	SelfLink               string `json:"selfLink"`
	AllowActiveMode        string `json:"allowActiveMode"`
	AllowFtps              string `json:"allowFtps"`
	AppService             string `json:"appService"`
	DefaultsFrom           string `json:"defaultsFrom"`
	Description            string `json:"description"`
	EnforceTLSSessionReuse string `json:"enforceTlsSessionReuse"`
	FtpsMode               string `json:"ftpsMode"`
	InheritParentProfile   string `json:"inheritParentProfile"`
	InheritVlanList        string `json:"inheritVlanList"`
	LogProfile             string `json:"logProfile"`
	LogPublisher           string `json:"logPublisher"`
	Port                   int    `json:"port"`
	Security               string `json:"security"`
	TranslateExtended      string `json:"translateExtended"`
}

const FTPEndpoint = "ftp"

type FTPResource struct {
	b *bigip.BigIP
}

func (cr *FTPResource) List() (*FTPList, error) {
	var items FTPList
	res, err := cr.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(FTPEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &items); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &items, nil
}

func (cr *FTPResource) Get(fullPathName string) (*FTP, error) {
	var item FTP
	res, err := cr.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(FTPEndpoint).SubResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &item, nil
}

func (cr *FTPResource) Create(item FTP) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = cr.b.RestClient.Post().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(FTPEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (cr *FTPResource) Update(fullPathName string, item FTP) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = cr.b.RestClient.Put().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(FTPEndpoint).SubResourceInstance(fullPathName).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (cr *FTPResource) Delete(fullPathName string) error {
	_, err := cr.b.RestClient.Delete().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(FTPEndpoint).SubResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
