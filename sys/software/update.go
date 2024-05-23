package software

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
)

// Update holds the configuration of a single Update.
type Update struct {
	Kind          string `json:"kind"`
	SelfLink      string `json:"selfLink"`
	AutoCheck     string `json:"autoCheck"`
	AutoPhonehome string `json:"autoPhonehome"`
	CheckStatus   string `json:"checkStatus"`
	Errors        int    `json:"errors"`
	Frequency     string `json:"frequency"`
}

// UpdateEndpoint represents the REST resource for managing Update.
const UpdateEndpoint = "update"

// UpdateResource provides an API to manage Update configurations.
type UpdateResource struct {
	b *bigip.BigIP
}

// Show retrieves the details of a single Update
func (r *UpdateStatusResource) Show() (*Update, error) {
	var item Update
	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(SoftwareEndpoint).SubResource(UpdateEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &item, nil
}
