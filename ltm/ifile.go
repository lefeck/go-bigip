package ltm

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

type IFileList struct {
	Kind     string  `json:"kind,omitempty"`
	SelfLink string  `json:"selfLink,omitempty"`
	Items    []IFile `json:"items,omitempty"`
}

type IFile struct {
	AppService  string `json:"appService,omitempty"`
	Description string `json:"description,omitempty"`
	Name        string `json:"name,omitempty"`
	FileName    string `json:"fileName,omitempty"`
	TMPartition string `json:"tmPartition,omitempty"`
	Partition   string `json:"partition,omitempty"`
}

const IFileEndpoint = "ifile"

type IFileResource struct {
	b *bigip.BigIP
}

func (ifr *IFileResource) List() (*IFileList, error) {
	var ifl IFileList

	res, err := ifr.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(IFileEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &ifl); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &ifl, nil
}

func (ifr *IFileResource) Get(fullPathName string) (*IFile, error) {
	var item IFile

	res, err := ifr.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(IFileEndpoint).ResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &item, nil
}

func (ifr *IFileResource) Create(name, fileObject string) error {
	item := map[string]string{
		"name":      name,
		"file-name": fileObject,
	}
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = ifr.b.RestClient.Post().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(IFileEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (ifr *IFileResource) Edit(name, fileObject string) error {
	item := map[string]string{
		"name":      name,
		"file-name": fileObject,
	}
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = ifr.b.RestClient.Put().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(IFileEndpoint).ResourceInstance(name).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (ifr *IFileResource) Delete(name string) error {
	_, err := ifr.b.RestClient.Delete().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(IFileEndpoint).ResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
