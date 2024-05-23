package sys

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
)

/*
The iControl SOAP interface was deprecated in version 11.6. The interface is still functional, but is not being
actively developed. See K98821303 on AskF5 for details.
*/
// IControlSOAP holds the configuration of a single IControlSOAP.
type IControlSOAP struct {
	Kind     string   `json:"kind"`
	SelfLink string   `json:"selfLink"`
	Allow    []string `json:"allow"`
}

// IControlSOAPEndpoint represents the REST resource for managing IControlSOAP.
const IControlSOAPEndpoint = "/icontrol-soap"

// IControlSOAPResource provides an API to manage IControlSOAP configurations.
type IControlSOAPResource struct {
	b *bigip.BigIP
}

// Show retrieves IControlSOAP details.
func (r *IControlSOAPResource) Show() (*IControlSOAP, error) {
	var items IControlSOAP
	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(IControlSOAPEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &items); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &items, nil
}
