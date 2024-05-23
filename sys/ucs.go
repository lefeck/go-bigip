package sys

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
	"time"
)

// UCSList holds a list of UCS configuration.
type UCSList struct {
	Items    []UCS  `json:"items"`
	Kind     string `json:"kind"`
	SelfLink string `json:"selflink"`
}

// UCS holds the configuration of a single UCS.
type UCS struct {
	Kind         string `json:"kind"`
	Generation   int    `json:"generation"`
	APIRawValues struct {
		BaseBuild       string    `json:"base_build"`
		Build           string    `json:"build"`
		Built           string    `json:"built"`
		Changelist      string    `json:"changelist"`
		Edition         string    `json:"edition"`
		Encrypted       string    `json:"encrypted"`
		FileCreatedDate time.Time `json:"file_created_date"`
		FileSize        string    `json:"file_size"`
		Filename        string    `json:"filename"`
		InstallDate     string    `json:"install_date"`
		JobID           string    `json:"job_id"`
		Product         string    `json:"product"`
		Sequence        string    `json:"sequence"`
		Version         string    `json:"version"`
	} `json:"apiRawValues"`
}

// UCSEndpoint represents the REST resource for managing UCS.
const UCSEndpoint = "ucs"

// UCSResource provides an API to manage UCS configurations.
type UCSResource struct {
	b *bigip.BigIP
}

// List all ucs details
func (r *UCSResource) List() (*UCSList, error) {
	var items UCSList
	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(UCSEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &items); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &items, nil
}

// Get a single ucs details by the node name
func (r *UCSResource) Get(name string) (*UCS, error) {
	var item UCS
	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(UCSEndpoint).ResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &item, nil
}

// Create a new ucs item
func (r *UCSResource) Create(item UCS) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = r.b.RestClient.Post().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(UCSEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Update the ucs item identified by the ucs name, otherwise an error will be reported.
func (r *UCSResource) Update(name string, item UCS) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = r.b.RestClient.Put().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(UCSEndpoint).ResourceInstance(name).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Delete a single ucs identified by the ucs name. if it is not exist return error
func (r *UCSResource) Delete(name string) error {
	_, err := r.b.RestClient.Delete().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(UCSEndpoint).ResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return err
	}

	return nil
}
