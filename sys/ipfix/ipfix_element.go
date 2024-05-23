package ipfix

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
)

// IPFixElementList holds a list of IPFixElement configurations.
type IPFixElementList struct {
	Items    []IPFixElement `json:"items"`
	Kind     string         `json:"kind"`
	SelfLink string         `json:"selflink"`
}

// IPFixElementConfig holds the configuration of a single IPFixElement.
type IPFixElement struct {
	DataType     string `json:"dataType"`
	EnterpriseID int    `json:"enterpriseId"`
	FullPath     string `json:"fullPath"`
	Generation   int    `json:"generation"`
	ID           int    `json:"id"`
	Kind         string `json:"kind"`
	Name         string `json:"name"`
	Partition    string `json:"partition"`
	SelfLink     string `json:"selfLink"`
	Size         int    `json:"size"`
}

// IPFixElementEndpoint represents the REST resource for managing IPFixElement.
const IPFixElementEndpoint = "element"

// IPFixElementResource provides an API to manage IPFixElement configurations.
type IPFixElementResource struct {
	b *bigip.BigIP
}

// List retrieves all IPFixElement details.
func (r *IPFixElementResource) List() (*IPFixElementList, error) {
	var items IPFixElementList
	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(IPFixEndpoint).SubResource(IPFixElementEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &items); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &items, nil
}

// Get retrieves the details of a single IPFixElement by node name.
func (r *IPFixElementResource) Get(name string) (*IPFixElement, error) {
	var item IPFixElement
	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(IPFixEndpoint).SubResource(IPFixElementEndpoint).ResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &item, nil
}
