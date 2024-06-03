package global_settings

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

// Metrics holds the configuration of a single Metrics.
type Metrics struct {
	DefaultProbeLimit             int      `json:"defaultProbeLimit,omitempty"`
	HopsPacketLength              int      `json:"hopsPacketLength,omitempty"`
	HopsSampleCount               int      `json:"hopsSampleCount,omitempty"`
	HopsTimeout                   int      `json:"hopsTimeout,omitempty"`
	HopsTTL                       int      `json:"hopsTtl,omitempty"`
	InactiveLdnsTTL               int      `json:"inactiveLdnsTtl,omitempty"`
	InactivePathsTTL              int      `json:"inactivePathsTtl,omitempty"`
	Kind                          string   `json:"kind,omitempty"`
	LdnsUpdateInterval            int      `json:"ldnsUpdateInterval,omitempty"`
	MaxSynchronousMonitorRequests int      `json:"maxSynchronousMonitorRequests,omitempty"`
	MetricsCaching                int      `json:"metricsCaching,omitempty"`
	MetricsCollectionProtocols    []string `json:"metricsCollectionProtocols,omitempty"`
	PathTTL                       int      `json:"pathTtl,omitempty"`
	PathsRetry                    int      `json:"pathsRetry,omitempty"`
	SelfLink                      string   `json:"selfLink,omitempty"`
}

// MetricsEndpoint represents the REST resource for managing Metrics.
const MetricsEndpoint = "metrics"

// MetricsResource provides an API to manage GlobalSettingsMetrics configurations.
type MetricsResource struct {
	b *bigip.BigIP
}

// List  lists all the Metrics configurations.
func (r *MetricsResource) List() (*Metrics, error) {
	var item Metrics
	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(GlobalSettingsEndpoint).SubResource(MetricsEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &item, nil
}

// Update a Metrics configuration.
func (r *MetricsResource) Update(item Metrics) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = r.b.RestClient.Post().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(GTMManager).
		Resource(GlobalSettingsEndpoint).SubResource(MetricsEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
