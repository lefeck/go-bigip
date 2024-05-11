package ltm

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

type DataGroupInternalList struct {
	Items    []DataGroupInternal `json:"items"`
	Kind     string              `json:"kind"`
	SelfLink string              `json:"selfLink"`
}

type DataGroupInternal struct {
	Description string `json:"description"`
	FullPath    string `json:"fullPath"`
	Generation  int    `json:"generation"`
	Kind        string `json:"kind"`
	Name        string `json:"name"`
	Records     []struct {
		Data string `json:"data"`
		Name string `json:"name"`
	} `json:"records"`
	SelfLink  string `json:"selfLink"`
	Type      string `json:"type"`
	Partition string `json:"partition"`
}

const DataGroupInternalEndpoint = "/data-group/internal"

type DataGroupInternalResource struct {
	b *bigip.BigIP
}

func (dgir *DataGroupInternalResource) List() (*DataGroupInternalList, error) {
	var dgil DataGroupInternalList

	res, err := dgir.b.RestClient.Get().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
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
	res, err := dgir.b.RestClient.Get().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
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
	_, err = dgir.b.RestClient.Post().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
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
	_, err = dgir.b.RestClient.Put().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(DataGroupInternalEndpoint).ResourceInstance(fullPathName).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (dgir *DataGroupInternalResource) Delete(fullPathName string) error {
	_, err := dgir.b.RestClient.Delete().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(DataGroupInternalEndpoint).ResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
