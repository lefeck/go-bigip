package gtm

import (
	"github.com/lefeck/go-bigip"
	global_settings "github.com/lefeck/go-bigip/gtm/global-settings"
	"github.com/lefeck/go-bigip/gtm/pool"
	"github.com/lefeck/go-bigip/gtm/wideip"
	"github.com/lefeck/go-bigip/ltm/monitor"
)

const GTMManager = "gtm"

type GTM struct {
	syncStatus       SyncStatusResource
	datacenter       DatacenterResource
	distributedApp   DistributedAppResource
	globalSettings   global_settings.GlobalSettingsResource
	link             LinkResource
	listener         ListenerResource
	listenerProfiles ListenerProfilesResource
	persist          PersistResource
	pool             pool.PoolResource
	proberPool       ProberPoolResource
	region           RegionResource
	rule             RuleResource
	server           ServerResource
	topology         TopologyResource
	wideip           wideip.WideipResource
	monitor          monitor.MonitorResource
}

// New creates a new GTM client.
func New(b *bigip.BigIP) GTM {
	return GTM{
		datacenter:       DatacenterResource{b: b},
		syncStatus:       SyncStatusResource{b: b},
		distributedApp:   DistributedAppResource{b: b},
		link:             LinkResource{b: b},
		listener:         ListenerResource{b: b},
		listenerProfiles: ListenerProfilesResource{b: b},
		persist:          PersistResource{b: b},
		proberPool:       ProberPoolResource{b: b},
		region:           RegionResource{b: b},
		rule:             RuleResource{b: b},
		server:           ServerResource{b: b},
		topology:         TopologyResource{b: b},

		globalSettings: global_settings.NewGlobalSettings(b),
		wideip:         wideip.NewWideip(b),
		pool:           pool.NewPoolResource(b),
		monitor:        monitor.NewMonitor(b),
	}
}

// Datacenter returns a configured DatacenterResource.
func (gtm GTM) SyncStatus() *SyncStatusResource {
	return &gtm.syncStatus
}

// Datacenter returns a configured DatacenterResource.
func (gtm GTM) Datacenter() *DatacenterResource {
	return &gtm.datacenter
}

// DistributedApp returns a configured DistributedAppResource.
func (gtm GTM) DistributedApp() *DistributedAppResource {
	return &gtm.distributedApp
}

// GlobalSettings returns a configured GlobalSettingsResource.
func (gtm GTM) GlobalSettingsGeneral() *global_settings.GlobalSettingsResource {
	return &gtm.globalSettings
}

// Persist returns a configured PersistResource.
func (gtm GTM) Persist() *PersistResource {
	return &gtm.persist
}

// ProberPool returns a configured ProberPoolResource.
func (gtm GTM) ProberPool() *ProberPoolResource {
	return &gtm.proberPool
}

// Region returns a configured RegionResource.
func (gtm GTM) Region() *RegionResource {
	return &gtm.region
}

// Rule returns a configured RuleResource.
func (gtm GTM) Rule() *RuleResource {
	return &gtm.rule
}

// Server returns a configured ServerResource.
func (gtm GTM) Server() *ServerResource {
	return &gtm.server
}

// Topology returns a configured TopologyResource.
func (gtm GTM) Topology() *TopologyResource {
	return &gtm.topology
}

// Wideip returns a configured WideipResource.
func (gtm GTM) WideipA() *wideip.WideipResource {
	return &gtm.wideip
}

// Pool returns a configured PoolResource.
func (gtm GTM) PoolA() *pool.PoolResource {
	return &gtm.pool
}

// Monitor returns a configured MonitorResource.
func (gtm GTM) Monitor() *monitor.MonitorResource {
	return &gtm.monitor
}
