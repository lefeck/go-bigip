package ltm

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
	"time"
)

type Persistence struct {
	Name      string `json:"name,omitempty"`
	Partition string `json:"partition,omitempty"`
	TMDefault string `json:"tmDefault,omitempty"`
}

// VirtualServerList contains a list of virtual server.
type VirtualServerList struct {
	Items    []VirtualServer `json:"items,omitempty"`
	Kind     string          `json:"kind,omitempty" pretty:",expanded"`
	SelfLink string          `json:"selfLink,omitempty" pretty:",expanded"`
}

// VirtualServer contains only a single virtual server.
type VirtualServer struct {
	Kind                             string                   `json:"kind"`
	Name                             string                   `json:"name"`
	Partition                        string                   `json:"partition"`
	FullPath                         string                   `json:"fullPath"`
	Generation                       int64                    `json:"generation"`
	SelfLink                         string                   `json:"selfLink"`
	AddressStatus                    string                   `json:"addressStatus"`
	AutoLasthop                      string                   `json:"autoLasthop"`
	CmpEnabled                       string                   `json:"cmpEnabled"`
	ConnectionLimit                  int                      `json:"connectionLimit"`
	CreationTime                     time.Time                `json:"creationTime"`
	Description                      string                   `json:"description"`
	Destination                      string                   `json:"destination"`
	Enabled                          bool                     `json:"enabled"`
	Disabled                         bool                     `json:"disabled"`
	EvictionProtected                string                   `json:"evictionProtected"`
	FallbackPersistence              string                   `json:"fallbackPersistence"`
	GtmScore                         int64                    `json:"gtmScore"`
	IPProtocol                       string                   `json:"ipProtocol"`
	LastModifiedTime                 time.Time                `json:"lastModifiedTime"`
	Mask                             string                   `json:"mask"`
	Mirror                           string                   `json:"mirror"`
	MobileAppTunnel                  string                   `json:"mobileAppTunnel"`
	Nat64                            string                   `json:"nat64"`
	Pool                             string                   `json:"pool"`
	Profiles                         []string                 `json:"profiles"`
	RateLimit                        string                   `json:"rateLimit"`
	RateLimitDstMask                 int64                    `json:"rateLimitDstMask"`
	RateLimitMode                    string                   `json:"rateLimitMode"`
	RateLimitSrcMask                 int64                    `json:"rateLimitSrcMask"`
	ReselectTries                    int64                    `json:"reselectTries"`
	ServersslUseSni                  string                   `json:"serversslUseSni"`
	ServiceDownAction                string                   `json:"serviceDownAction"`
	ServiceDownImmediateAction       string                   `json:"serviceDownImmediateAction"`
	Source                           string                   `json:"source"`
	SourceAddressTranslation         SourceAddressTranslation `json:"sourceAddressTranslation,omitempty"`
	SourcePort                       string                   `json:"sourcePort"`
	Rules                            []string                 `json:"rules,omitempty"`
	SlowRampTime                     int                      `json:"slowRampTime"`
	SynCookieStatus                  string                   `json:"synCookieStatus"`
	TrafficMatchingCriteria          string                   `json:"trafficMatchingCriteria"`
	TrafficMatchingCriteriaReference struct {
		Link string `json:"link"`
	} `json:"trafficMatchingCriteriaReference"`
	TranslateAddress string `json:"translateAddress"`
	TranslatePort    string `json:"translatePort"`
	VlansDisabled    bool   `json:"vlansDisabled"`
	PoolReference    struct {
		Link string `json:"link"`
	} `json:"poolReference"`
	Persistences      []Persistence `json:"persist,omitempty"`
	PoliciesReference struct {
		Link            string `json:"link"`
		IsSubcollection bool   `json:"isSubcollection"`
	} `json:"policiesReference"`
	ProfilesReference struct {
		Link            string    `json:"link"`
		IsSubcollection bool      `json:"isSubcollection"`
		Profiles        []Profile `json:"items,omitempty"`
	} `json:"profilesReference"`
}

type SourceAddressTranslation struct {
	Type string `json:"type,omitempty"`
	Pool string `json:"pool,omitempty"`
}

type Profile struct {
	Name    string `json:"name,omitempty"`
	Context string `json:"context,omitempty"`
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

// Update the virtual server identified by the virtual server name,
// Source and Mask fields must be specified, otherwise an error will be reported. For example:
/*
   item := VirtualServer{
       ......
       Source:    "0.0.0.0/32",
       Mask:      "255.255.255.255",
       ......
   }
*/
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
