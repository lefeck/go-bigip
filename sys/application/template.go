package application

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

// TemplateList holds a list of Template configurations.
type TemplateList struct {
	Items    []Template `json:"items"`
	Kind     string     `json:"kind"`
	SelfLink string     `json:"selflink"`
}

// Template holds the configuration of a single Template.
type Template struct {
	Kind                    string   `json:"kind"`
	Name                    string   `json:"name"`
	Partition               string   `json:"partition"`
	FullPath                string   `json:"fullPath"`
	Generation              int      `json:"generation"`
	SelfLink                string   `json:"selfLink"`
	IgnoreVerification      string   `json:"ignoreVerification"`
	RequiresBigipVersionMax string   `json:"requiresBigipVersionMax,omitempty"`
	RequiresBigipVersionMin string   `json:"requiresBigipVersionMin"`
	RequiresModules         []string `json:"requiresModules,omitempty"`
	SigningKey              string   `json:"signingKey"`
	SigningKeyReference     struct {
		Link string `json:"link"`
	} `json:"signingKeyReference"`
	TmplSignature      string `json:"tmplSignature"`
	TotalSigningStatus string `json:"totalSigningStatus"`
	VerificationStatus string `json:"verificationStatus"`
	ActionsReference   struct {
		Link            string `json:"link"`
		IsSubcollection bool   `json:"isSubcollection"`
	} `json:"actionsReference"`
}

// TemplateEndpoint represents the REST resource for managing Template.
const TemplateEndpoint = "template"

// TemplateResource provides an API to manage Template configurations.
type TemplateResource struct {
	b *bigip.BigIP
}

// List retrieves all Template details.
func (r *TemplateResource) List() (*TemplateList, error) {
	var items TemplateList
	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(ApplicationEndpoint).SubResource(TemplateEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &items); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &items, nil
}

// Get retrieves the details of a single Template by node name.
func (r *TemplateResource) Get(name string) (*Template, error) {
	var item Template
	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(ApplicationEndpoint).SubResource(TemplateEndpoint).SubResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &item, nil
}

// Create creates a new Template item.
func (r *TemplateResource) Create(item Template) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = r.b.RestClient.Post().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(ApplicationEndpoint).SubResource(TemplateEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Update modifies the Template item identified by the Template name.
func (r *TemplateResource) Update(name string, item Template) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = r.b.RestClient.Put().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(ApplicationEndpoint).SubResource(TemplateEndpoint).SubResourceInstance(name).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Delete a single Template identified by the Template name. if it is not exist return error
func (r *TemplateResource) Delete(name string) error {
	_, err := r.b.RestClient.Delete().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(ApplicationEndpoint).SubResource(TemplateEndpoint).SubResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return err
	}

	return nil
}
