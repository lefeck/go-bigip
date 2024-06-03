package global_settings

import "github.com/lefeck/go-bigip"

// GlobalSettingsEndpoint represents the REST resource for managing GlobalSettings.
const GlobalSettingsEndpoint = "global-settings"
const GTMManager = "gtm"

// GlobalSettingsResource provides an API to manage GlobalSettings configurations.
type GlobalSettingsResource struct {
	general       GeneralResource
	loadBalancing LoadBalancingResource
	metrics       MetricsResource
}

func NewGlobalSettings(b *bigip.BigIP) GlobalSettingsResource {
	return GlobalSettingsResource{
		general:       GeneralResource{b: b},
		loadBalancing: LoadBalancingResource{b: b},
		metrics:       MetricsResource{b: b},
	}
}

func (gs *GlobalSettingsResource) General() *GeneralResource {
	return &gs.general
}

func (gs *GlobalSettingsResource) LoadBalancing() *LoadBalancingResource {
	return &gs.loadBalancing
}

func (gs *GlobalSettingsResource) Metrics() *MetricsResource {
	return &gs.metrics
}
