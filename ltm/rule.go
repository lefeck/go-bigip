package ltm

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

// A RuleList contains a list of iRule configurations.
type RuleList struct {
	Items    []Rule `json:"items,omitempty"`
	Kind     string `json:"kind,omitempty" pretty:",expanded"`
	SelfLink string `json:"selfLink,omitempty" pretty:",expanded"`
}

// Rule contains only the iRule configuration.
type Rule struct {
	Action              string `json:"action,omitempty"`
	AppService          string `json:"appService,omitempty"`
	DefinitionChecksum  string `json:"definitionChecksum,omitempty"`
	DefinitionSignature string `json:"definitionSignature,omitempty"`
	Hidden              bool   `json:"hidden,omitempty"`
	IgnoreVerification  string `json:"ignoreVerification,omitempty"`
	NoDelete            bool   `json:"noDelete,omitempty"`
	NoWrite             bool   `json:"noWrite,omitempty"`
	TMPartition         string `json:"tmPartition,omitempty"`
	Plugin              string `json:"plugin,omitempty"`
	PublicCert          string `json:"publicCert,omitempty"`
	SigningKey          string `json:"signingKey,omitempty"`

	Name         string `json:"name,omitempty"`
	Partition    string `json:"partition,omitempty"`
	FullPath     string `json:"fullPath,omitempty"`
	SelfLink     string `json:"selfLink,omitempty"`
	ApiAnonymous string `json:"apiAnonymous,omitempty"`
}

// RuleEndpoint represents the REST resource for managing iRule configurations.
const RuleEndpoint = "rule"

// A RuleResource provides an API to manage iRule configurations.
type RuleResource struct {
	b *bigip.BigIP
}

func (rr *RuleResource) List() (*RuleList, error) {
	res, err := rr.b.RestClient.Get().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(RuleEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	var rl RuleList
	if err := json.Unmarshal(res, &rl); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &rl, nil
}

func (rr *RuleResource) Get(name string) (*Rule, error) {
	res, err := rr.b.RestClient.Get().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(RuleEndpoint).ResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	var rule Rule
	if err := json.Unmarshal(res, &rule); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &rule, nil
}

func (rr *RuleResource) Create(item Rule) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = rr.b.RestClient.Post().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(RuleEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (rr *RuleResource) Update(name string, item Rule) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = rr.b.RestClient.Put().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(RuleEndpoint).ResourceInstance(name).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (rr *RuleResource) Delete(name string) error {
	_, err := rr.b.RestClient.Delete().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(RuleEndpoint).ResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
