package gtm

import (
	"github.com/lefeck/go-bigip"
)

const GTMManager = "gtm"

type GTM struct {
	//syncStatus                  SyncStatusResource
	datacenter DatacenterResource
	//distributedApp              DistributedAppResource
	//globalSettingsGeneral       GlobalSettingsGeneralResource
	//globalSettingsLoadBalancing GlobalSettingsLoadBalancingResource
	//globalSettingsMetrics       GlobalSettingsMetricsResource
	//link                        LinkResource
	//listener                    ListenerResource
	//listenerProfiles            ListenerProfilesResource
	//persist                     PersistResource
	//poolA                       PoolAResource
	//poolAAAA                    PoolAAAAResource
	//poolCNAME                   PoolCNAMEResource
	//poolMX                      PoolMXResource
	//poolNAPTR                   PoolNAPTRResource
	//poolSRV                     PoolSRVResource
	//proberPool                  ProberPoolResource
	//region                      RegionResource
	//rule                        RuleResource
	//server                      ServerResource
	//topology                    TopologyResource
	//wideipA                     WideipAResource
	//wideipAAAA                  WideipAAAAResource
	//wideipCname                 WideipCnameResource
	//wideipMx                    WideipMxResource
	//wideipNaptr                 WideipNaptrResource
	//wideipSrv                   WideipSrvResource
}

// New creates a new GTM client.
func New(b *bigip.BigIP) GTM {
	return GTM{
		datacenter: DatacenterResource{b: b},
		//syncStatus:                  SyncStatusResource{c: c},
		//distributedApp:              DistributedAppResource{c: c},
		//globalSettingsGeneral:       GlobalSettingsGeneralResource{c: c},
		//globalSettingsLoadBalancing: GlobalSettingsLoadBalancingResource{c: c},
		//globalSettingsMetrics:       GlobalSettingsMetricsResource{c: c},
		//link:                        LinkResource{c: c},
		//listener:                    ListenerResource{c: c},
		//listenerProfiles:            ListenerProfilesResource{c: c},
		//
		//persist:     PersistResource{c: c},
		//poolA:       PoolAResource{c: c},
		//poolAAAA:    PoolAAAAResource{c: c},
		//poolCNAME:   PoolCNAMEResource{c: c},
		//poolMX:      PoolMXResource{c: c},
		//poolNAPTR:   PoolNAPTRResource{c: c},
		//poolSRV:     PoolSRVResource{c: c},
		//proberPool:  ProberPoolResource{c: c},
		//region:      RegionResource{c: c},
		//rule:        RuleResource{c: c},
		//server:      ServerResource{c: c},
		//topology:    TopologyResource{c: c},
		//wideipA:     WideipAResource{c: c},
		//wideipAAAA:  WideipAAAAResource{c: c},
		//wideipCname: WideipCnameResource{c: c},
		//wideipMx:    WideipMxResource{c: c},
		//wideipNaptr: WideipNaptrResource{c: c},
		//wideipSrv:   WideipSrvResource{c: c},
	}
}

//// Datacenter returns a configured DatacenterResource.
//func (gtm GTM) SyncStatus() *SyncStatusResource {
//	return &gtm.syncStatus
//}

// Datacenter returns a configured DatacenterResource.
func (gtm GTM) Datacenter() *DatacenterResource {
	return &gtm.datacenter
}

//// DistributedApp returns a configured DistributedAppResource.
//func (gtm GTM) DistributedApp() *DistributedAppResource {
//	return &gtm.distributedApp
//}
//
//// GlobalSettingsGeneral returns a configured GlobalSettingsGeneralResource.
//func (gtm GTM) GlobalSettingsGeneral() *GlobalSettingsGeneralResource {
//	return &gtm.globalSettingsGeneral
//}
//
//// GlobalSettingsLoadBalancing returns a configured GlobalSettingsLoadBalancingResource.
//func (gtm GTM) GlobalSettingsLoadBalancing() *GlobalSettingsLoadBalancingResource {
//	return &gtm.globalSettingsLoadBalancing
//}
//
//// GlobalSettingsMetrics returns a configured GlobalSettingsMetricsResource.
//func (gtm GTM) GlobalSettingsMetrics() *GlobalSettingsMetricsResource {
//	return &gtm.globalSettingsMetrics
//}
//
//// Persist returns a configured PersistResource.
//func (gtm GTM) Persist() *PersistResource {
//	return &gtm.persist
//}
//
//// PoolA returns a configured PoolAResource.
//func (gtm GTM) PoolA() *PoolAResource {
//	return &gtm.poolA
//}
//
//// PoolAAAA returns a configured PoolAAAAResource.
//func (gtm GTM) PoolAAAA() *PoolAAAAResource {
//	return &gtm.poolAAAA
//}
//
//// PoolCNAME returns a configured PoolCNAMEResource.
//func (gtm GTM) PoolCNAME() *PoolCNAMEResource {
//	return &gtm.poolCNAME
//}
//
//// PoolMX returns a configured PoolMXResource.
//func (gtm GTM) PoolMX() *PoolMXResource {
//	return &gtm.poolMX
//}
//
//// PoolNAPTR returns a configured PoolNAPTRResource.
//func (gtm GTM) PoolNAPTR() *PoolNAPTRResource {
//	return &gtm.poolNAPTR
//}
//
//// PoolSRV returns a configured PoolSRVResource.
//func (gtm GTM) PoolSRV() *PoolSRVResource {
//	return &gtm.poolSRV
//}
//
//// ProberPool returns a configured ProberPoolResource.
//func (gtm GTM) ProberPool() *ProberPoolResource {
//	return &gtm.proberPool
//}
//
//// Region returns a configured RegionResource.
//func (gtm GTM) Region() *RegionResource {
//	return &gtm.region
//}
//
//// Rule returns a configured RuleResource.
//func (gtm GTM) Rule() *RuleResource {
//	return &gtm.rule
//}
//
//// Server returns a configured ServerResource.
//func (gtm GTM) Server() *ServerResource {
//	return &gtm.server
//}
//
//// Topology returns a configured TopologyResource.
//func (gtm GTM) Topology() *TopologyResource {
//	return &gtm.topology
//}
//
//// WideipA returns a configured WideipAResource.
//func (gtm GTM) WideipA() *WideipAResource {
//	return &gtm.wideipA
//}
//
//// WideipAAAA returns a configured WideipAAAAResource.
//func (gtm GTM) WideipAAAA() *WideipAAAAResource {
//	return &gtm.wideipAAAA
//}
//
//// WideipCname returns a configured WideipCnameResource.
//func (gtm GTM) WideipCname() *WideipCnameResource {
//	return &gtm.wideipCname
//}
//
//// WideipMx returns a configured WideipMxResource.
//func (gtm GTM) WideipMx() *WideipMxResource {
//	return &gtm.wideipMx
//}
//
//// WideipNaptr returns a configured WideipNaptrResource.
//func (gtm GTM) WideipNaptr() *WideipNaptrResource {
//	return &gtm.wideipNaptr
//}
//
//// WideipSrv returns a configured WideipSrvResource.
//func (gtm GTM) WideipSrv() *WideipSrvResource {
//	return &gtm.wideipSrv
//}
