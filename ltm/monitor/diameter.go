package monitor

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

type DiameterList struct {
	Items    []Diameter `json:"items,omitempty"`
	Kind     string     `json:"kind,omitempty"`
	SelfLink string     `json:"selflink,omitempty"`
}

type Diameter struct {
	AcctApplicationId               string `json:"acctApplicationId,omitempty"`
	AppService                      string `json:"appService,omitempty"`
	AuthApplicationId               string `json:"authApplicationId,omitempty"`
	DefaultsFrom                    string `json:"defaultsFrom,omitempty"`
	Description                     string `json:"description,omitempty"`
	Destination                     string `json:"destination,omitempty"`
	FullPath                        string `json:"fullPath,omitempty"`
	Generation                      int    `json:"generation,omitempty"`
	HostIpAddress                   string `json:"hostIpAddress,omitempty"`
	Interval                        int    `json:"interval,omitempty"`
	Kind                            string `json:"kind,omitempty"`
	ManualResume                    string `json:"manualResume,omitempty"`
	Name                            string `json:"name,omitempty"`
	OriginHost                      string `json:"originHost,omitempty"`
	OriginRealm                     string `json:"originRealm,omitempty"`
	Partition                       string `json:"partition,omitempty"`
	ProductName                     string `json:"productName,omitempty"`
	SelfLink                        string `json:"selfLink,omitempty"`
	TimeUntilUp                     int    `json:"timeUntilUp,omitempty"`
	Timeout                         int    `json:"timeout,omitempty"`
	UpInterval                      int    `json:"upInterval,omitempty"`
	VendorId                        string `json:"vendorId,omitempty"`
	VendorSpecificAcctApplicationId string `json:"vendorSpecificAcctApplicationId,omitempty"`
	VendorSpecificAuthApplicationId string `json:"vendorSpecificAuthApplicationId,omitempty"`
	VendorSpecificVendorId          string `json:"vendorSpecificVendorId,omitempty"`
}

const DiameterEndpoint = "diameter"

type DiameterResource struct {
	b *bigip.BigIP
}

func (mdr *DiameterResource) List() (*DiameterList, error) {
	var mdcl DiameterList
	res, err := mdr.b.RestClient.Get().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(MonitorEndpoint).SubResource(DiameterEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &mdcl); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &mdcl, nil
}

func (mdr *DiameterResource) Get(fullPathName string) (*Diameter, error) {
	var mdc Diameter
	res, err := mdr.b.RestClient.Get().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(MonitorEndpoint).SubResource(DiameterEndpoint).SubResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &mdc); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &mdc, nil
}

func (mdr *DiameterResource) Create(item Diameter) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = mdr.b.RestClient.Post().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(MonitorEndpoint).SubResource(DiameterEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (mdr *DiameterResource) Update(fullPathName string, item Diameter) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = mdr.b.RestClient.Put().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(MonitorEndpoint).SubResource(DiameterEndpoint).SubResourceInstance(fullPathName).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (mdr *DiameterResource) Delete(fullPathName string) error {
	_, err := mdr.b.RestClient.Delete().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(MonitorEndpoint).SubResource(DiameterEndpoint).SubResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
