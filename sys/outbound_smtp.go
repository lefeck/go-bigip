package sys

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
)

// OutboundSMTP holds the configuration of a single OutboundSMTP.
type OutboundSMTP struct {
	Kind             string `json:"kind"`
	SelfLink         string `json:"selfLink"`
	FromLineOverride string `json:"fromLineOverride"`
	Mailhub          string `json:"mailhub"`
}

// OutboundSMTPEndpoint represents the REST resource for managing OutboundSMTP.
const OutboundSMTPEndpoint = "outboundsmtp"

// OutboundSMTPResource provides an API to manage OutboundSMTP configurations.
type OutboundSMTPResource struct {
	b *bigip.BigIP
}

// Show retrieves all OutboundSMTP details.
func (r *OutboundSMTPResource) Show() (*OutboundSMTP, error) {
	var item OutboundSMTP
	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(OutboundSMTPEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &item, nil
}
