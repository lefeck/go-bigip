package ltm

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
	"time"
)

// VirtualServerList contains a list of virtual server.
type VirtualServerList struct {
	Items    []VirtualServer `json:"items,omitempty"`
	Kind     string          `json:"kind,omitempty" pretty:",expanded"`
	SelfLink string          `json:"selfLink,omitempty" pretty:",expanded"`
}

// VirtualServer contains only a single virtual server.
type VirtualServer struct {
	Kind                             string                   `json:"kind,omitempty"`
	Name                             string                   `json:"name,omitempty"`
	Partition                        string                   `json:"partition,omitempty"`
	FullPath                         string                   `json:"fullPath,omitempty"`
	Generation                       int64                    `json:"generation,omitempty"`
	SelfLink                         string                   `json:"selfLink,omitempty"`
	AddressStatus                    string                   `json:"addressStatus,omitempty"`
	AutoLasthop                      string                   `json:"autoLasthop,omitempty"`
	CmpEnabled                       string                   `json:"cmpEnabled,omitempty"`
	ConnectionLimit                  int                      `json:"connectionLimit,omitempty"`
	CreationTime                     time.Time                `json:"creationTime,omitempty"`
	Description                      string                   `json:"description,omitempty"`
	Destination                      string                   `json:"destination,omitempty"`
	Enabled                          bool                     `json:"enabled,omitempty"`
	Disabled                         bool                     `json:"disabled,omitempty"`
	EvictionProtected                string                   `json:"evictionProtected,omitempty"`
	FallbackPersistence              string                   `json:"fallbackPersistence,omitempty"`
	GtmScore                         int64                    `json:"gtmScore,omitempty"`
	IPProtocol                       string                   `json:"ipProtocol,omitempty"`
	LastModifiedTime                 time.Time                `json:"lastModifiedTime,omitempty"`
	Mask                             string                   `json:"mask,omitempty"`
	Mirror                           string                   `json:"mirror,omitempty"`
	MobileAppTunnel                  string                   `json:"mobileAppTunnel,omitempty"`
	Nat64                            string                   `json:"nat64,omitempty"`
	Pool                             string                   `json:"pool,omitempty"`
	Profiles                         []string                 `json:"profiles,omitempty"`
	RateLimit                        string                   `json:"rateLimit,omitempty"`
	RateLimitDstMask                 int64                    `json:"rateLimitDstMask,omitempty"`
	RateLimitMode                    string                   `json:"rateLimitMode,omitempty"`
	RateLimitSrcMask                 int64                    `json:"rateLimitSrcMask,omitempty"`
	ReselectTries                    int64                    `json:"reselectTries,omitempty"`
	ServersslUseSni                  string                   `json:"serversslUseSni,omitempty"`
	ServiceDownAction                string                   `json:"serviceDownAction,omitempty"`
	ServiceDownImmediateAction       string                   `json:"serviceDownImmediateAction,omitempty"`
	Source                           string                   `json:"source,omitempty"`
	SourceAddressTranslation         SourceAddressTranslation `json:"sourceAddressTranslation,omitempty"`
	SourcePort                       string                   `json:"sourcePort,omitempty"`
	Rules                            []string                 `json:"rules,omitempty"`
	SlowRampTime                     int                      `json:"slowRampTime,omitempty"`
	SynCookieStatus                  string                   `json:"synCookieStatus,omitempty"`
	TrafficMatchingCriteria          string                   `json:"trafficMatchingCriteria,omitempty"`
	TrafficMatchingCriteriaReference struct {
		Link string `json:"link,omitempty"`
	} `json:"trafficMatchingCriteriaReference,omitempty"`
	TranslateAddress string   `json:"translateAddress,omitempty"`
	TranslatePort    string   `json:"translatePort,omitempty"`
	Vlans            []string `json:"vlans,omitempty"`
	VlansEnabled     bool     `json:"vlansEnabled,omitempty"`
	VlansDisabled    bool     `json:"vlansDisabled,omitempty"`
	PoolReference    struct {
		Link string `json:"link,omitempty"`
	} `json:"poolReference,omitempty"`
	Persistences      []Persistence `json:"persist,omitempty"`
	PoliciesReference struct {
		Link            string `json:"link,omitempty"`
		IsSubcollection bool   `json:"isSubcollection,omitempty"`
	} `json:"policiesReference,omitempty"`
	ProfilesReference struct {
		Link            string    `json:"link,omitempty"`
		IsSubcollection bool      `json:"isSubcollection,omitempty"`
		Profiles        []Profile `json:"items,omitempty"`
	} `json:"profilesReference,omitempty"`
}

type SourceAddressTranslation struct {
	Type string `json:"type,omitempty"`
	Pool string `json:"pool,omitempty"`
}

type Profile struct {
	Name    string `json:"name,omitempty"`
	Context string `json:"context,omitempty"`
}

type Persistence struct {
	Name      string `json:"name,omitempty"`
	Partition string `json:"partition,omitempty"`
	TMDefault string `json:"tmDefault,omitempty"`
}

// VirtualEndpoint is the base path of the ltm API.
const VirtualEndpoint = "virtual"

// VirtualResource provides an API to manage virtual server.
type VirtualResource struct {
	b *bigip.BigIP
}

// List all virtual server items
func (vr *VirtualResource) List() (*VirtualServerList, error) {

	res, err := vr.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(VirtualEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	var vsl VirtualServerList
	if err := json.Unmarshal(res, &vsl); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &vsl, nil
}

// List all the details of the virtual server, including: profile, policy, etc.
func (vr *VirtualResource) ListDetail() (*VirtualServerList, error) {
	res, err := vr.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(VirtualEndpoint).SetParams("expandSubcollections", "true").DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	var vsl VirtualServerList
	if err := json.Unmarshal(res, &vsl); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &vsl, nil
}

// ListVirtualServerName get all virtual server names
func (vr *VirtualResource) ListVirtualServerName() ([]string, error) {
	vsl, err := vr.List()
	if err != nil {
		return nil, err
	}
	var items []string
	for _, vs := range vsl.Items {
		fullPathName := vs.FullPath
		items = append(items, fullPathName)
	}
	return items, nil
}

// Get a single virtual server identified by name.
func (vr *VirtualResource) Get(fullPathName string) (*VirtualServer, error) {
	res, err := vr.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(VirtualEndpoint).ResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	var vs VirtualServer
	if err := json.Unmarshal(res, &vs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &vs, nil
}

// Create a new virtual server item
func (vr *VirtualResource) Create(item VirtualServer) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = vr.b.RestClient.Post().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(VirtualEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Update the virtual server identified by the virtual server name
func (vr *VirtualResource) Update(name string, item VirtualServer) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = vr.b.RestClient.Put().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(VirtualEndpoint).ResourceInstance(name).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Delete a single virtual server identified by the virtual server name. if it is not exist return error
func (vr *VirtualResource) Delete(name string) error {
	_, err := vr.b.RestClient.Delete().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(VirtualEndpoint).ResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Enabling a virtual server item identified by the virtual server name.
func (vr *VirtualResource) Enable(name string) error {
	item := VirtualServer{Enabled: true}
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = vr.b.RestClient.Patch().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(VirtualEndpoint).ResourceInstance(name).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Disabling a virtual server item identified by the virtual server name.
func (vr *VirtualResource) Disable(name string) error {
	item := VirtualServer{Disabled: true}

	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = vr.b.RestClient.Patch().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(VirtualEndpoint).ResourceInstance(name).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// removes a single iRule from the virtual server identified by virtual server name.
func (vr *VirtualResource) RemoveRuleForVirtualServer(vsName, ruleName string) error {
	item := VirtualServer{
		Rules: []string{
			ruleName,
		},
	}
	res, err := vr.b.RestClient.Delete().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(VirtualEndpoint).ResourceInstance(vsName).Body(item).DoRaw(context.Background())
	if err != nil {
		return err
	}
	var vs VirtualServer
	if err := json.Unmarshal(res, &vs); err != nil {
		panic(err)
	}
	return nil
}

// gets the iRules for a virtual server identified by name.
func (vr *VirtualResource) GetRulesByVirtualServer(name string) ([]Rule, error) {
	res, err := vr.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(VirtualEndpoint).ResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	var rules []Rule
	if err := json.Unmarshal(res, &rules); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}

	return rules, nil
}

// adds an iRule to the virtual server identified by name.
func (vr *VirtualResource) AddRuleForVirtualServer(vsName string, rule Rule) error {
	jsonData, err := json.Marshal(rule)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = vr.b.RestClient.Post().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(VirtualEndpoint).SubResource(RuleEndpoint).ResourceInstance(vsName).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
