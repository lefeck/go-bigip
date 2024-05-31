package ltm

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

type DataGroupInternalList struct {
	Items    []DataGroupInternal `json:"items,omitempty"`
	Kind     string              `json:"kind,omitempty"`
	SelfLink string              `json:"selfLink,omitempty"`
}

type DataGroupInternal struct {
	Description string `json:"description,omitempty"`
	FullPath    string `json:"fullPath,omitempty"`
	Generation  int    `json:"generation,omitempty"`
	Kind        string `json:"kind,omitempty"`
	Name        string `json:"name,omitempty"`
	Records     []struct {
		Data string `json:"data,omitempty"`
		Name string `json:"name,omitempty"`
	} `json:"records,omitempty"`
	SelfLink  string `json:"selfLink,omitempty"`
	Type      string `json:"type,omitempty"`
	Partition string `json:"partition,omitempty"`
}

const DataGroupInternalEndpoint = "/data-group/internal"

type DataGroupInternalResource struct {
	b *bigip.BigIP
}

func (dgir *DataGroupInternalResource) List() (*DataGroupInternalList, error) {
	var dgil DataGroupInternalList

	res, err := dgir.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(DataGroupInternalEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &dgil); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &dgil, nil
}

func (dgir *DataGroupInternalResource) Get(fullPathName string) (*DataGroupInternal, error) {
	var item DataGroupInternal
	res, err := dgir.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(DataGroupInternalEndpoint).ResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &item, nil
}

func (dgir *DataGroupInternalResource) Create(item DataGroupInternal) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = dgir.b.RestClient.Post().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(DataGroupInternalEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (dgir *DataGroupInternalResource) Update(fullPathName string, item DataGroupInternal) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = dgir.b.RestClient.Put().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(DataGroupInternalEndpoint).ResourceInstance(fullPathName).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (dgir *DataGroupInternalResource) Delete(fullPathName string) error {
	_, err := dgir.b.RestClient.Delete().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(DataGroupInternalEndpoint).ResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
