package global_settings

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

// LoadBalancing holds the configuration of a single LoadBalancing.
type LoadBalancing struct {
	FailureRcode              string `json:"failureRcode,omitempty"`
	FailureRcodeResponse      string `json:"failureRcodeResponse,omitempty"`
	FailureRcodeTTL           int    `json:"failureRcodeTtl,omitempty"`
	IgnorePathTTL             string `json:"ignorePathTtl,omitempty"`
	Kind                      string `json:"kind,omitempty"`
	QosFactorBps              int    `json:"qosFactorBps,omitempty"`
	QosFactorHitRatio         int    `json:"qosFactorHitRatio,omitempty"`
	QosFactorHops             int    `json:"qosFactorHops,omitempty"`
	QosFactorLinkCapacity     int    `json:"qosFactorLinkCapacity,omitempty"`
	QosFactorPacketRate       int    `json:"qosFactorPacketRate,omitempty"`
	QosFactorRtt              int    `json:"qosFactorRtt,omitempty"`
	QosFactorTopology         int    `json:"qosFactorTopology,omitempty"`
	QosFactorVsCapacity       int    `json:"qosFactorVsCapacity,omitempty"`
	QosFactorVsScore          int    `json:"qosFactorVsScore,omitempty"`
	RespectFallbackDependency string `json:"respectFallbackDependency,omitempty"`
	SelfLink                  string `json:"selfLink,omitempty"`
	TopologyAllowZeroScores   string `json:"topologyAllowZeroScores,omitempty"`
	TopologyLongestMatch      string `json:"topologyLongestMatch,omitempty"`
	VerifyVsAvailability      string `json:"verifyVsAvailability,omitempty"`
}

// LoadBalancingEndpoint represents the REST resource for managing LoadBalancing.
const LoadBalancingEndpoint = "load-balancing"

// LoadBalancingResource provides an API to manage LoadBalancing configurations.
type LoadBalancingResource struct {
	b *bigip.BigIP
}

// List  lists all the LoadBalancing configurations.
func (r *LoadBalancingResource) List() (*LoadBalancing, error) {
	var item LoadBalancing
	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(GlobalSettingsEndpoint).SubResource(LoadBalancingEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &item, nil
}

// Update a LoadBalancing configuration.
func (r *LoadBalancingResource) Update(item LoadBalancing) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = r.b.RestClient.Post().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(GlobalSettingsEndpoint).SubResource(LoadBalancingEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
