package application

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/lefeck/go-bigip"
	"strings"
)

// APLScriptList holds a list of APLScript configurations.
type APLScriptList struct {
	Items    []APLScript `json:"items"`
	Kind     string      `json:"kind"`
	SelfLink string      `json:"selflink"`
}

// APLScript holds the configuration of a single APLScript.
type APLScript struct {
	Kind                string `json:"kind"`
	Name                string `json:"name"`
	Partition           string `json:"partition"`
	FullPath            string `json:"fullPath"`
	Generation          int    `json:"generation"`
	SelfLink            string `json:"selfLink"`
	APIAnonymous        string `json:"apiAnonymous"`
	AplSignature        string `json:"aplSignature"`
	IgnoreVerification  string `json:"ignoreVerification"`
	SigningKey          string `json:"signingKey"`
	SigningKeyReference struct {
		Link string `json:"link"`
	} `json:"signingKeyReference"`
	TotalSigningStatus string `json:"totalSigningStatus"`
	VerificationStatus string `json:"verificationStatus"`
}

// APLScriptEndpoint represents the REST resource for managing APLScript.
const APLScriptEndpoint = "apl-script"

// APLScriptResource provides an API to manage APLScript configurations.
type APLScriptResource struct {
	b *bigip.BigIP
}

// List retrieves all APLScript details.
func (r *APLScriptResource) List() (*APLScriptList, error) {
	var items APLScriptList
	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(ApplicationEndpoint).SubResource(APLScriptEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &items); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &items, nil
}

// Get retrieves the details of a single APLScript by node name.
func (r *APLScriptResource) Get(name string) (*APLScript, error) {
	var item APLScript
	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(ApplicationEndpoint).SubResource(APLScriptEndpoint).SubResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &item, nil
}

// Create creates a new APLScript item.
func (r *APLScriptResource) Create(item APLScript) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = r.b.RestClient.Post().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(ApplicationEndpoint).SubResource(APLScriptEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Update modifies the APLScript item identified by the APLScript name.
func (r *APLScriptResource) Update(name string, item APLScript) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = r.b.RestClient.Put().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(ApplicationEndpoint).SubResource(APLScriptEndpoint).SubResourceInstance(name).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Delete a single APLScript identified by the APLScript name. if it is not exist return error
func (r *APLScriptResource) Delete(name string) error {
	_, err := r.b.RestClient.Delete().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(ApplicationEndpoint).SubResource(APLScriptEndpoint).SubResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return err
	}

	return nil
}
