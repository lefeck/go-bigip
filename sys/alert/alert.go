package alert

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
)

// AlertConfigList holds a list of Alert configuration.
type AlertConfigList struct {
	Items    []AlertConfig `json:"items"`
	Kind     string        `json:"kind"`
	SelfLink string        `json:"selflink"`
}

// AlertConfig holds the configuration of a single Alert.
type AlertConfig struct {
}

// AlertEndpoint represents the REST resource for managing Alert.
const AlertEndpoint = "alert"
const SysManager = "sys"

// AlertResource provides an API to manage Alert configurations.
type AlertResource struct {
	b *bigip.BigIP
}

// ListAll  lists all the Alert configurations.
func (r *AlertResource) List() (*AlertConfigList, error) {
	var items AlertConfigList
	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(AlertEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &items); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &items, nil
}
