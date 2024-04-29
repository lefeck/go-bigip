package ltm

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/bigip"
	"github.com/lefeck/bigip/rest"
	"strings"
)

type Persistence struct {
	Name      string `json:"name,omitempty"`
	Partition string `json:"partition,omitempty"`
	TMDefault string `json:"tmDefault,omitempty"`
}

// VirtualServerList holds a list of virtual server uration.
type VirtualServerList struct {
	Items    []VirtualServer `json:"items,omitempty"`
	Kind     string          `json:"kind,omitempty" pretty:",expanded"`
	SelfLink string          `json:"selfLink,omitempty" pretty:",expanded"`
}

// VirtualServer holds the uration of a single virtual server.
type VirtualServer struct {
	AddressStatus       string        `json:"addressStatus,omitempty"`
	AutoLasthop         string        `json:"autoLasthop,omitempty"`
	CmpEnabled          string        `json:"cmpEnabled,omitempty"`
	ConnectionLimit     int64         `json:"connectionLimit,omitempty"`
	Description         string        `json:"description,omitempty"`
	Destination         string        `json:"destination,omitempty"`
	Enabled             bool          `json:"enabled,omitempty"`
	FallbackPersistence string        `json:"fallbackPersistence,omitempty"`
	FullPath            string        `json:"fullPath,omitempty" pretty:",expanded"`
	FwEnforcedPolicy    string        `json:"fwEnforcedPolicy,omitempty"`
	Generation          int64         `json:"generation,omitempty" pretty:",expanded"`
	GtmScore            int64         `json:"gtmScore,omitempty" pretty:",expanded"`
	IPProtocol          string        `json:"ipProtocol,omitempty"`
	Kind                string        `json:"kind,omitempty" pretty:",expanded"`
	Mask                string        `json:"mask,omitempty"`
	Mirror              string        `json:"mirror,omitempty"`
	MobileAppTunnel     string        `json:"mobileAppTunnel,omitempty" pretty:",expanded"`
	Name                string        `json:"name,omitempty"`
	Nat64               string        `json:"nat64,omitempty" pretty:",expanded"`
	Partition           string        `json:"partition,omitempty"`
	Persistences        []Persistence `json:"persist,omitempty"`
	PoliciesReference   struct {
		IsSubcollection bool   `json:"isSubcollection,omitempty"`
		Link            string `json:"link,omitempty"`
	} `json:"policiesReference,omitempty"`
	Pool              string   `json:"pool,omitempty"`
	Profiles          []string `json:"profiles,omitempty"` // only used to link existing profiles a creation or update
	ProfilesReference struct {
		IsSubcollection bool      `json:"isSubcollection,omitempty"`
		Link            string    `json:"link,omitempty"`
		Profiles        []Profile `json:"items,omitempty"`
	} `json:"profilesReference,omitempty"`
	RateLimit                string                   `json:"rateLimit,omitempty" pretty:",expanded"`
	RateLimitDstMask         int64                    `json:"rateLimitDstMask,omitempty" pretty:",expanded"`
	RateLimitMode            string                   `json:"rateLimitMode,omitempty" pretty:",expanded"`
	RateLimitSrcMask         int64                    `json:"rateLimitSrcMask,omitempty" pretty:",expanded"`
	Rules                    []string                 `json:"rules,omitempty"`
	SelfLink                 string                   `json:"selfLink,omitempty" pretty:",expanded"`
	SecurityLogProfiles      []string                 `json:"securityLogProfiles,omitempty" pretty:",expanded"`
	Source                   string                   `json:"source,omitempty"`
	SourceAddressTranslation SourceAddressTranslation `json:"sourceAddressTranslation,omitempty"`
	SourcePort               string                   `json:"sourcePort,omitempty"`
	SynCookieStatus          string                   `json:"synCookieStatus,omitempty"`
	TranslateAddress         string                   `json:"translateAddress,omitempty"`
	TranslatePort            string                   `json:"translatePort,omitempty"`
	Vlans                    []string                 `json:"vlans,omitempty"`
	VlansDisabled            bool                     `json:"vlansDisabled,omitempty"`
	VlansEnabled             bool                     `json:"vlansEnabled,omitempty"`
	VsIndex                  int64                    `json:"vsIndex,omitempty" pretty:",expanded"`
}

type SourceAddressTranslation struct {
	Type string `json:"type,omitempty"`
	Pool string `json:"pool,omitempty"`
}

type Profile struct {
	Name    string `json:"name,omitempty"`
	Context string `json:"context,omitempty"`
}

// VirtualResource provides an API to manage virtual server urations.
type VirtualResource struct {
	b *bigip.BigIP
}

// ListAll lists all the virtual server urations.
func (vr *VirtualResource) List() (*VirtualServerList, error) {
	res, err := vr.b.RestClient.Get().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LTMManager).
		Resource(VirtualEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	var vsl VirtualServerList
	if err := json.Unmarshal(res, &vsl); err != nil {
		panic(err)
	}
	return &vsl, nil
}

//
//// ListAllWithParams lists all the virtual server urations.
//func (vr *VirtualResource) ListAllWithParams(v url.Values) (*VirtualServerList, error) {
//	params := v.Encode()
//
//	resp, err := vr.doRequest("GET", "?"+params, nil)
//	if err != nil {
//		return nil, err
//	}
//	defer resp.Body.Close()
//	if err := vr.readError(resp); err != nil {
//		return nil, err
//	}
//	var vsc VirtualServerList
//	dec := json.NewDecoder(resp.Body)
//	if err := dec.Decode(&vsc); err != nil {
//		return nil, err
//	}
//	return &vsc, nil
//}

// Get a single virtual server uration identified by id.
func (vr *VirtualResource) Get(fullPathName string) (*VirtualServer, error) {
	res, err := vr.b.RestClient.Get().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LTMManager).
		Resource(VirtualEndpoint).Suffix(Suffix).ResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	result := rest.Result{Body: res}

	var vs VirtualServer
	if err := json.Unmarshal(result.Body, &vs); err != nil {
		panic(err)
	}
	return &vs, nil
}

// Create a new virtual server instance
func (vr *VirtualResource) Create(item VirtualServer) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	res, err := vr.b.RestClient.Post().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LTMManager).
		Resource(VirtualEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	result := rest.Result{Body: res}

	var vs VirtualServer
	if err := json.Unmarshal(result.Body, &vs); err != nil {
		panic(err)
	}
	return nil
}

// update the virtual server identified by the virtual server name.
func (vr *VirtualResource) Update(name string, item VirtualServer) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	res, err := vr.b.RestClient.Put().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LTMManager).
		Resource(VirtualEndpoint).Suffix(Suffix).ResourceInstance(name).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}

	var vs VirtualServer
	if err := json.Unmarshal(res, &vs); err != nil {
		panic(err)
	}
	return nil
}

// Enabling a virtual server item identified by the virtual server name.
func (vr *VirtualResource) Enable(name string) error {
	item := VirtualServer{Enabled: true}
	res, err := vr.b.RestClient.Patch().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LTMManager).
		Resource(VirtualEndpoint).ResourceInstance(name).Body(item).DoRaw(context.Background())
	if err != nil {
		return err
	}

	var vs VirtualServer
	if err := json.Unmarshal(res, &vs); err != nil {
		panic(err)
	}
	return nil
}

// Delete a single server identified by the virtual server name. if it is not exist return error
// for example: https://192.168.13.91/mgmt/tm/ltm/virtual/~Common~go-test
func (vr *VirtualResource) Delete(name string) error {
	_, err := vr.b.RestClient.Delete().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LTMManager).
		Resource(VirtualEndpoint).Suffix(Suffix).ResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

//
//// RemoveRule removes a single  iRule from the virtual server identified by id.
//func (vr *VirtualResource) RemoveRule(vsID, ruleID string) error {
//	res, err := vr.b.RestClient.Delete().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LTMManager).
//		Resource(VirtualEndpoint).ResourceInstance(name).Body(name).DoRaw(context.Background())
//	if err != nil {
//		return err
//	}
//
//	var vs VirtualServer
//	if err := json.Unmarshal(res, &vs); err != nil {
//		panic(err)
//	}
//	return nil
//}

//// Rules gets the iRules uration for a virtual server identified by id.
//func (vr *VirtualResource) Rules(id string) ([]Rule, error) {
//	resp, err := vr.doRequest("GET", id+"/rule", nil)
//	if err != nil {
//		return nil, err
//	}
//	defer resp.Body.Close()
//	if err := vr.readError(resp); err != nil {
//		return nil, err
//	}
//	var rules []Rule
//	dec := json.NewDecoder(resp.Body)
//	if err := dec.Decode(&rules); err != nil {
//		return nil, err
//	}
//	return rules, nil
//}
//
//// AddRule adds an iRule to the virtual server identified by id.
//func (vr *VirtualResource) AddRule(id string, rule Rule) error {
//	resp, err := vr.doRequest("POST", id+"/rule", rule)
//	if err != nil {
//		return err
//	}
//	defer resp.Body.Close()
//	if err := vr.readError(resp); err != nil {
//		return err
//	}
//	return nil
//}

// RemoveRule removes a single  iRule from the virtual server identified by id.

//func (vr *VirtualResource) RemoveRule(vsID, ruleID string) error {
//	resp, err := vr.doRequest("DELETE", vsID+"/rule/"+ruleID, nil)
//	if err != nil {
//		return err
//	}
//	defer resp.Body.Close()
//	if err := vr.readError(resp); err != nil {
//		return err
//	}
//	return nil
//}
