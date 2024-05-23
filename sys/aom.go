package sys

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

// AOMConfig holds the configuration of a single AOM.
type AOM struct {
	Kind             string `json:"kind"`
	SelfLink         string `json:"selfLink"`
	Ipmi             string `json:"ipmi"`
	MediaRedirection string `json:"mediaRedirection"`
	Readonly         string `json:"readonly"`
	Vkvm             string `json:"vkvm"`
	Webui            string `json:"webui"`
}

// AOMEndpoint represents the REST resource for managing AOM.
const AOMEndpoint = "aom"

// AOMResource provides an API to manage AOM configurations.
type AOMResource struct {
	b *bigip.BigIP
}

// List retrieves all AOM details.
func (r *AOMResource) Show() (*AOM, error) {
	var items AOM
	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(AOMEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &items); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &items, nil
}

// Update modifies the AOM item identified by the AOM name.
func (r *AOMResource) Update(name string, item AOM) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = r.b.RestClient.Put().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(AOMEndpoint).ResourceInstance(name).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
