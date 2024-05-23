// Copyright 2016 e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import (
	"github.com/lefeck/go-bigip"
	"github.com/lefeck/go-bigip/sys/application"
	"github.com/lefeck/go-bigip/sys/disk"
	"github.com/lefeck/go-bigip/sys/ipfix"
	"github.com/lefeck/go-bigip/sys/raid"
	"github.com/lefeck/go-bigip/sys/software"
)

const SysManager = "sys"

type Sys struct {
	aOM AOMResource
	//alert    AlertResource
	//alertLCD AlertLCDResource

	application    application.ApplicationResource
	autoscaleGroup AutoscaleGroupResource
	//classificationSignature ClassificationSignatureResource
	clock      ClockResource
	cluster    ClusterResource
	connection ConnectionResource
	console    ConsoleResource
	cpuStats   CPUStatsResource
	//crypto                              CryptoResource
	//cryptoCRL                           CryptoCRLResource
	//cryptoCSR                           CryptoCSRResource
	//cryptoCert                          CryptoCertResource
	//cryptoCheckCert                     CryptoCheckCertResource
	//cryptoClient                        CryptoClientResource
	//cryptoKey                           CryptoKeyResource
	//cryptoPKCS12                        CryptoPKCS12Resource
	//cryptoServer                        CryptoServerResource
	dB       DBResource
	dNS      DNSResource
	daemonHA DaemonHAResource
	//daemonLogSettings                   DaemonLogSettingsResource
	//daemonLogSettingsClusterd           DaemonLogSettingsClusterdResource
	//daemonLogSettingsCsyncd             DaemonLogSettingsCsyncdResource
	//daemonLogSettingsICRD               DaemonLogSettingsICRDResource
	//daemonLogSettingsLind               DaemonLogSettingsLindResource
	//daemonLogSettingsMCPD               DaemonLogSettingsMCPDResource
	//daemonLogSettingsTMM                DaemonLogSettingsTMMResource
	dataStor DataStorResource
	disk     disk.DiskResource
	//eCM                                 ECMResource
	//eCMCloudProvider                    ECMCloudProviderResource
	//eCMConfig                           ECMConfigResource
	//fPGA                                FPGAResource
	//fPGAFirmwareConfig                  FPGAFirmwareConfigResource
	//fPGAInfo                            FPGAInfoResource
	//failover      FailoverResource
	featureModule FeatureModuleResource
	//fileApacheSSLCert                   FileApacheSSLCertResource
	//fileApacheSSLCertBundleCertificates FileApacheSSLCertBundleCertificatesResource
	//fileBrowserCapabilitiesDB           FileBrowserCapabilitiesDBResource
	//fileDashboardViewset                FileDashboardViewsetResource
	//fileDataGroup                       FileDataGroupResource
	//fileDeviceCapabilitiesDB            FileDeviceCapabilitiesDBResource
	//fileExternalMonitor                 FileExternalMonitorResource
	//fileIFile                           FileIFileResource
	//fileLWTunnelTBL                     FileLWTunnelTBLResource
	//fileSSLCRL                          FileSSLCRLResource
	//fileSSLCSR                          FileSSLCSRResource
	//fileSSLCert                         FileSSLCertResource
	//fileSSLCertBundleCertificates       FileSSLCertBundleCertificatesResource
	//fileSSLKey                          FileSSLKeyResource
	//fileSystemSSLCert                   FileSystemSSLCertResource
	//fileSystemSSLCertBundleCertificates FileSystemSSLCertBundleCertificatesResource
	//fileSystemSSLKey                    FileSystemSSLKeyResource
	fixConnection FixConnectionResource
	//folder         FolderResource
	//globalSettings GlobalSettingsResource
	hAGroup HAGroupResource
	hTTPD   HTTPDResource
	//iCall                               ICallResource
	//iCallEvent                          ICallEventResource
	//iCallIStatsTrigger                  ICallIStatsTriggerResource
	//iCallPublisher                      ICallPublisherResource
	//iCallScript                         ICallScriptResource
	//iControlSOAP IControlSOAPResource
	iPFix ipfix.IPFixResource
	//iPFixElement                        IPFixElementResource
	//lTCFGClass                          LTCFGClassResource
	//lTCFGClassFields                    LTCFGClassFieldsResource
	//lTCFGInstance                       LTCFGInstanceResource
	//lTCFGInstanceFields                 LTCFGInstanceFieldsResource
	license LicenseResource
	//log                                 LogResource
	//logConfig                           LogConfigResource
	//logConfigFilter                     LogConfigFilterResource
	//logConfigPublisher                  LogConfigPublisherResource
	//logRotate                           LogRotateResource
	managementDHCP  ManagementDHCPResource
	managementIP    ManagementIPResource
	managementOVSDB ManagementOVSDBResource
	managementRoute ManagementRouteResource
	memoryStats     MemoryStatsResource
	nTP             NTPResource
	outboundSMTP    OutboundSMTPResource
	//pFMan                               PFManResource
	//pFManConsumer                       PFManConsumerResource
	//pFManDevice                         PFManDeviceResource
	pPTPCallInfo PPTPCallInfoResource
	//performance      PerformanceResource
	provision        ProvisionResource
	rAID             raid.RAIDResource
	restrictedModule RestrictedModuleResource
	//sFlow                               SFlowResource
	//sFlowReceiver                       SFlowReceiverResource
	sMTPServer     SMTPServerResource
	sNMP           SNMPResource
	sSHD           SSHDResource
	scriptd        ScriptdResource
	service        ServiceResource
	software       software.SoftwareResource
	stateMirroring StateMirroringResource
	syncSysFiles   SyncSysFilesResource
	syslog         SyslogResource
	//uRLDB                               URLDBResource
	//uRLDBDownloadResult                 URLDBDownloadResultResource
	//uRLDBDownloadSchedule               URLDBDownloadScheduleResource
	//uRLDBURLCategory                    URLDBURLCategoryResource
}

func New(b *bigip.BigIP) Sys {
	return Sys{
		aOM: AOMResource{b: b},
		//alert:                               AlertResource{c: c},
		//alertLCD:                            AlertLCDResource{c: c},
		application:    application.NewApplication(b),
		autoscaleGroup: AutoscaleGroupResource{b: b},
		//classificationSignature: ClassificationSignatureResource{c: b},
		clock: ClockResource{b: b},
		//cluster:                 ClusterResource{c: b},
		connection: ConnectionResource{b: b},
		console:    ConsoleResource{b: b},
		cpuStats:   CPUStatsResource{b: b},
		//crypto:                              CryptoResource{c: c},
		//cryptoCRL:                           CryptoCRLResource{c: c},
		//cryptoCSR:                           CryptoCSRResource{c: c},
		//cryptoCert:                          CryptoCertResource{c: c},
		//cryptoCheckCert:                     CryptoCheckCertResource{c: c},
		//cryptoClient:                        CryptoClientResource{c: c},
		//cryptoKey:                           CryptoKeyResource{c: c},
		//cryptoPKCS12:                        CryptoPKCS12Resource{c: c},
		//cryptoServer:                        CryptoServerResource{c: c},
		dB:       DBResource{b: b},
		dNS:      DNSResource{b: b},
		daemonHA: DaemonHAResource{b: b},
		//daemonLogSettings:                   DaemonLogSettingsResource{c: c},
		//daemonLogSettingsClusterd:           DaemonLogSettingsClusterdResource{c: c},
		//daemonLogSettingsCsyncd:             DaemonLogSettingsCsyncdResource{c: c},
		//daemonLogSettingsICRD:               DaemonLogSettingsICRDResource{c: c},
		//daemonLogSettingsLind:               DaemonLogSettingsLindResource{c: c},
		//daemonLogSettingsMCPD:               DaemonLogSettingsMCPDResource{c: c},
		//daemonLogSettingsTMM:                DaemonLogSettingsTMMResource{c: c},
		dataStor: DataStorResource{b: b},
		disk:     disk.NewDisk(b),
		//eCM:                                 ECMResource{c: c},
		//eCMCloudProvider:                    ECMCloudProviderResource{c: c},
		//eCMConfig:                           ECMConfigResource{c: c},
		//fPGA:                                FPGAResource{c: c},
		//fPGAFirmwareConfig:                  FPGAFirmwareConfigResource{c: c},
		//fPGAInfo:                            FPGAInfoResource{c: c},
		//failover: FailoverResource{c: b},
		//featureModule: FeatureModuleResource{c: b},
		//fileApacheSSLCert:                   FileApacheSSLCertResource{c: c},
		//fileApacheSSLCertBundleCertificates: FileApacheSSLCertBundleCertificatesResource{c: c},
		//fileBrowserCapabilitiesDB:           FileBrowserCapabilitiesDBResource{c: c},
		//fileDashboardViewset:                FileDashboardViewsetResource{c: c},
		//fileDataGroup:                       FileDataGroupResource{c: c},
		//fileDeviceCapabilitiesDB:            FileDeviceCapabilitiesDBResource{c: c},
		//fileExternalMonitor:                 FileExternalMonitorResource{c: c},
		//fileIFile:                           FileIFileResource{c: c},
		//fileLWTunnelTBL:                     FileLWTunnelTBLResource{c: c},
		//fileSSLCRL:                          FileSSLCRLResource{c: c},
		//fileSSLCSR:                          FileSSLCSRResource{c: c},
		//fileSSLCert:                         FileSSLCertResource{c: c},
		//fileSSLCertBundleCertificates:       FileSSLCertBundleCertificatesResource{c: c},
		//fileSSLKey:                          FileSSLKeyResource{c: c},
		//fileSystemSSLCert:                   FileSystemSSLCertResource{c: c},
		//fileSystemSSLCertBundleCertificates: FileSystemSSLCertBundleCertificatesResource{c: c},
		//fileSystemSSLKey:                    FileSystemSSLKeyResource{c: c},
		//fixConnection:  FixConnectionResource{c: b},
		//folder:         FolderResource{c: b},
		//globalSettings: GlobalSettingsResource{c: b},
		hAGroup: HAGroupResource{b: b},
		hTTPD:   HTTPDResource{b: b},
		//iCall:                               ICallResource{c: c},
		//iCallEvent:                          ICallEventResource{c: c},
		//iCallIStatsTrigger:                  ICallIStatsTriggerResource{c: c},
		//iCallPublisher:                      ICallPublisherResource{c: c},
		//iCallScript:                         ICallScriptResource{c: c},
		//iControlSOAP: IControlSOAPResource{c: b},
		iPFix: ipfix.NewIPFix(b),
		//lTCFGClass:            LTCFGClassResource{c: c},
		//lTCFGClassFields:      LTCFGClassFieldsResource{c: c},
		//lTCFGInstance:         LTCFGInstanceResource{c: c},
		//lTCFGInstanceFields:   LTCFGInstanceFieldsResource{c: c},
		license: LicenseResource{b: b},
		//log:                   LogResource{c: c},
		//logConfig:             LogConfigResource{c: c},
		//logConfigFilter:       LogConfigFilterResource{c: c},
		//logConfigPublisher:    LogConfigPublisherResource{c: c},
		//logRotate:             LogRotateResource{c: c},
		managementDHCP:  ManagementDHCPResource{b: b},
		managementIP:    ManagementIPResource{b: b},
		managementOVSDB: ManagementOVSDBResource{b: b},
		managementRoute: ManagementRouteResource{b: b},
		memoryStats:     MemoryStatsResource{b: b},
		nTP:             NTPResource{b: b},
		outboundSMTP:    OutboundSMTPResource{b: b},
		//pFMan:                 PFManResource{c: c},
		//pFManConsumer:         PFManConsumerResource{c: c},
		//pFManDevice:           PFManDeviceResource{b: b},
		pPTPCallInfo: PPTPCallInfoResource{b: b},
		//performance:      PerformanceResource{c: b},
		provision:        ProvisionResource{b: b},
		rAID:             raid.NewRAID(b),
		restrictedModule: RestrictedModuleResource{b: b},
		//sFlow:                 SFlowResource{c: c},
		//sFlowReceiver:         SFlowReceiverResource{c: c},
		sMTPServer:     SMTPServerResource{b: b},
		sNMP:           SNMPResource{b: b},
		sSHD:           SSHDResource{b: b},
		scriptd:        ScriptdResource{b: b},
		service:        ServiceResource{b: b},
		software:       software.NewSoftware(b),
		stateMirroring: StateMirroringResource{b: b},
		syncSysFiles:   SyncSysFilesResource{b: b},
		syslog:         SyslogResource{b: b},
		//uRLDB:                 URLDBResource{c: c},
		//uRLDBDownloadResult:   URLDBDownloadResultResource{c: c},
		//uRLDBDownloadSchedule: URLDBDownloadScheduleResource{c: c},
		//uRLDBURLCategory:      URLDBURLCategoryResource{c: c},
	}
}

// aOM returns a configured AOMResource.
func (sys Sys) AOM() *AOMResource {
	return &sys.aOM
}

//// alert returns a configured AlertResource.
//func (sys Sys) Alert() *AlertResource {
//	return &sys.alert
//}
//
//// alertLCD returns a configured AlertLCDResource.
//func (sys Sys) AlertLCD() *AlertLCDResource {
//	return &sys.alertLCD
//}

// application returns a configured ApplicationResource.
func (sys Sys) Application() *application.ApplicationResource {
	return &sys.application
}

// autoscaleGroup returns a configured AutoscaleGroupResource.
func (sys Sys) AutoscaleGroup() *AutoscaleGroupResource {
	return &sys.autoscaleGroup
}

//// classificationSignature returns a configured ClassificationSignatureResource.
//func (sys Sys) ClassificationSignature() *ClassificationSignatureResource {
//	return &sys.classificationSignature
//}

// clock returns a configured ClockResource.
func (sys Sys) Clock() *ClockResource {
	return &sys.clock
}

// cluster returns a configured ClusterResource.
func (sys Sys) Cluster() *ClusterResource {
	return &sys.cluster
}

// connection returns a configured ConnectionResource.
func (sys Sys) Connection() *ConnectionResource {
	return &sys.connection
}

// console returns a configured ConsoleResource.
func (sys Sys) Console() *ConsoleResource {
	return &sys.console
}

// console returns a configured ConsoleResource.
func (sys Sys) CPUStats() *CPUStatsResource {
	return &sys.cpuStats
}

//// crypto returns a configured CryptoResource.
//func (sys Sys) Crypto() *CryptoResource {
//	return &sys.crypto
//}
//
//// cryptoCRL returns a configured CryptoCRLResource.
//func (sys Sys) CryptoCRL() *CryptoCRLResource {
//	return &sys.cryptoCRL
//}
//
//// cryptoCSR returns a configured CryptoCSRResource.
//func (sys Sys) CryptoCSR() *CryptoCSRResource {
//	return &sys.cryptoCSR
//}
//
//// cryptoCert returns a configured CryptoCertResource.
//func (sys Sys) CryptoCert() *CryptoCertResource {
//	return &sys.cryptoCert
//}
//
//// cryptoCheckCert returns a configured CryptoCheckCertResource.
//func (sys Sys) CryptoCheckCert() *CryptoCheckCertResource {
//	return &sys.cryptoCheckCert
//}
//
//// cryptoClient returns a configured CryptoClientResource.
//func (sys Sys) CryptoClient() *CryptoClientResource {
//	return &sys.cryptoClient
//}
//
//// cryptoKey returns a configured CryptoKeyResource.
//func (sys Sys) CryptoKey() *CryptoKeyResource {
//	return &sys.cryptoKey
//}
//
//// cryptoPKCS12 returns a configured CryptoPKCS12Resource.
//func (sys Sys) CryptoPKCS12() *CryptoPKCS12Resource {
//	return &sys.cryptoPKCS12
//}
//
//// cryptoServer returns a configured CryptoServerResource.
//func (sys Sys) CryptoServer() *CryptoServerResource {
//	return &sys.cryptoServer
//}

// dB returns a configured DBResource.
func (sys Sys) DB() *DBResource {
	return &sys.dB
}

// dNS returns a configured DNSResource.
func (sys Sys) DNS() *DNSResource {
	return &sys.dNS
}

// daemonHA returns a configured DaemonHAResource.
func (sys Sys) DaemonHA() *DaemonHAResource {
	return &sys.daemonHA
}

//// daemonLogSettings returns a configured DaemonLogSettingsResource.
//func (sys Sys) DaemonLogSettings() *DaemonLogSettingsResource {
//	return &sys.daemonLogSettings
//}
//
//// daemonLogSettingsClusterd returns a configured DaemonLogSettingsClusterdResource.
//func (sys Sys) DaemonLogSettingsClusterd() *DaemonLogSettingsClusterdResource {
//	return &sys.daemonLogSettingsClusterd
//}
//
//// daemonLogSettingsCsyncd returns a configured DaemonLogSettingsCsyncdResource.
//func (sys Sys) DaemonLogSettingsCsyncd() *DaemonLogSettingsCsyncdResource {
//	return &sys.daemonLogSettingsCsyncd
//}
//
//// daemonLogSettingsICRD returns a configured DaemonLogSettingsICRDResource.
//func (sys Sys) DaemonLogSettingsICRD() *DaemonLogSettingsICRDResource {
//	return &sys.daemonLogSettingsICRD
//}
//
//// daemonLogSettingsLind returns a configured DaemonLogSettingsLindResource.
//func (sys Sys) DaemonLogSettingsLind() *DaemonLogSettingsLindResource {
//	return &sys.daemonLogSettingsLind
//}
//
//// daemonLogSettingsMCPD returns a configured DaemonLogSettingsMCPDResource.
//func (sys Sys) DaemonLogSettingsMCPD() *DaemonLogSettingsMCPDResource {
//	return &sys.daemonLogSettingsMCPD
//}
//
//// daemonLogSettingsTMM returns a configured DaemonLogSettingsTMMResource.
//func (sys Sys) DaemonLogSettingsTMM() *DaemonLogSettingsTMMResource {
//	return &sys.daemonLogSettingsTMM
//}

// dataStor returns a configured DataStorResource.
func (sys Sys) DataStor() *DataStorResource {
	return &sys.dataStor
}

// disk returns a configured DiskResource.
func (sys Sys) Disk() *disk.DiskResource {
	return &sys.disk
}

//
//
//// eCM returns a configured ECMResource.
//func (sys Sys) ECM() *ECMResource {
//	return &sys.eCM
//}
//
//// eCMCloudProvider returns a configured ECMCloudProviderResource.
//func (sys Sys) ECMCloudProvider() *ECMCloudProviderResource {
//	return &sys.eCMCloudProvider
//}
//
//// eCMConfig returns a configured ECMConfigResource.
//func (sys Sys) ECMConfig() *ECMConfigResource {
//	return &sys.eCMConfig
//}
//
//// fPGA returns a configured FPGAResource.
//func (sys Sys) FPGA() *FPGAResource {
//	return &sys.fPGA
//}
//
//// fPGAFirmwareConfig returns a configured FPGAFirmwareConfigResource.
//func (sys Sys) FPGAFirmwareConfig() *FPGAFirmwareConfigResource {
//	return &sys.fPGAFirmwareConfig
//}
//
//// fPGAInfo returns a configured FPGAInfoResource.
//func (sys Sys) FPGAInfo() *FPGAInfoResource {
//	return &sys.fPGAInfo
//}
//
//// failover returns a configured FailoverResource.
//func (sys Sys) Failover() *FailoverResource {
//	return &sys.failover
//}

// featureModule returns a configured FeatureModuleResource.
func (sys Sys) FeatureModule() *FeatureModuleResource {
	return &sys.featureModule
}

//
//// fileApacheSSLCert returns a configured FileApacheSSLCertResource.
//func (sys Sys) FileApacheSSLCert() *FileApacheSSLCertResource {
//	return &sys.fileApacheSSLCert
//}
//
//// fileApacheSSLCertBundleCertificates returns a configured FileApacheSSLCertBundleCertificatesResource.
//func (sys Sys) FileApacheSSLCertBundleCertificates() *FileApacheSSLCertBundleCertificatesResource {
//	return &sys.fileApacheSSLCertBundleCertificates
//}
//
//// fileBrowserCapabilitiesDB returns a configured FileBrowserCapabilitiesDBResource.
//func (sys Sys) FileBrowserCapabilitiesDB() *FileBrowserCapabilitiesDBResource {
//	return &sys.fileBrowserCapabilitiesDB
//}
//
//// fileDashboardViewset returns a configured FileDashboardViewsetResource.
//func (sys Sys) FileDashboardViewset() *FileDashboardViewsetResource {
//	return &sys.fileDashboardViewset
//}
//
//// fileDataGroup returns a configured FileDataGroupResource.
//func (sys Sys) FileDataGroup() *FileDataGroupResource {
//	return &sys.fileDataGroup
//}
//
//// fileDeviceCapabilitiesDB returns a configured FileDeviceCapabilitiesDBResource.
//func (sys Sys) FileDeviceCapabilitiesDB() *FileDeviceCapabilitiesDBResource {
//	return &sys.fileDeviceCapabilitiesDB
//}
//
//// fileExternalMonitor returns a configured FileExternalMonitorResource.
//func (sys Sys) FileExternalMonitor() *FileExternalMonitorResource {
//	return &sys.fileExternalMonitor
//}
//
//// fileIFile returns a configured FileIFileResource.
//func (sys Sys) FileIFile() *FileIFileResource {
//	return &sys.fileIFile
//}
//
//// fileLWTunnelTBL returns a configured FileLWTunnelTBLResource.
//func (sys Sys) FileLWTunnelTBL() *FileLWTunnelTBLResource {
//	return &sys.fileLWTunnelTBL
//}
//
//// fileSSLCRL returns a configured FileSSLCRLResource.
//func (sys Sys) FileSSLCRL() *FileSSLCRLResource {
//	return &sys.fileSSLCRL
//}
//
//// fileSSLCSR returns a configured FileSSLCSRResource.
//func (sys Sys) FileSSLCSR() *FileSSLCSRResource {
//	return &sys.fileSSLCSR
//}
//
//// fileSSLCert returns a configured FileSSLCertResource.
//func (sys Sys) FileSSLCert() *FileSSLCertResource {
//	return &sys.fileSSLCert
//}
//
//// fileSSLCertBundleCertificates returns a configured FileSSLCertBundleCertificatesResource.
//func (sys Sys) FileSSLCertBundleCertificates() *FileSSLCertBundleCertificatesResource {
//	return &sys.fileSSLCertBundleCertificates
//}
//
//// fileSSLKey returns a configured FileSSLKeyResource.
//func (sys Sys) FileSSLKey() *FileSSLKeyResource {
//	return &sys.fileSSLKey
//}
//
//// fileSystemSSLCert returns a configured FileSystemSSLCertResource.
//func (sys Sys) FileSystemSSLCert() *FileSystemSSLCertResource {
//	return &sys.fileSystemSSLCert
//}
//
//// fileSystemSSLCertBundleCertificates returns a configured FileSystemSSLCertBundleCertificatesResource.
//func (sys Sys) FileSystemSSLCertBundleCertificates() *FileSystemSSLCertBundleCertificatesResource {
//	return &sys.fileSystemSSLCertBundleCertificates
//}
//
//// fileSystemSSLKey returns a configured FileSystemSSLKeyResource.
//func (sys Sys) FileSystemSSLKey() *FileSystemSSLKeyResource {
//	return &sys.fileSystemSSLKey
//}

// fixConnection returns a configured FixConnectionResource.
func (sys Sys) FixConnection() *FixConnectionResource {
	return &sys.fixConnection
}

//// folder returns a configured FolderResource.
//func (sys Sys) Folder() *FolderResource {
//	return &sys.folder
//}
//
//// globalSettings returns a configured GlobalSettingsResource.
//func (sys Sys) GlobalSettings() *GlobalSettingsResource {
//	return &sys.globalSettings
//}

// hAGroup returns a configured HAGroupResource.
func (sys Sys) HAGroup() *HAGroupResource {
	return &sys.hAGroup
}

// hTTPD returns a configured HTTPDResource.
func (sys Sys) HTTPD() *HTTPDResource {
	return &sys.hTTPD
}

//
//// iCall returns a configured ICallResource.
//func (sys Sys) ICall() *ICallResource {
//	return &sys.iCall
//}
//
//// iCallEvent returns a configured ICallEventResource.
//func (sys Sys) ICallEvent() *ICallEventResource {
//	return &sys.iCallEvent
//}
//
//// iCallIStatsTrigger returns a configured ICallIStatsTriggerResource.
//func (sys Sys) ICallIStatsTrigger() *ICallIStatsTriggerResource {
//	return &sys.iCallIStatsTrigger
//}
//
//// iCallPublisher returns a configured ICallPublisherResource.
//func (sys Sys) ICallPublisher() *ICallPublisherResource {
//	return &sys.iCallPublisher
//}
//
//// iCallScript returns a configured ICallScriptResource.
//func (sys Sys) ICallScript() *ICallScriptResource {
//	return &sys.iCallScript
//}

//// iControlSOAP returns a configured IControlSOAPResource.
//func (sys Sys) IControlSOAP() *IControlSOAPResource {
//	return &sys.iControlSOAP
//}

// iPFix returns a configured IPFixResource.
func (sys Sys) IPFix() *ipfix.IPFixResource {
	return &sys.iPFix
}

// iPFixElement returns a configured IPFixElementResource.
func (sys Sys) IPFixElement() *ipfix.IPFixResource {
	return &sys.iPFix
}

//
//// lTCFGClass returns a configured LTCFGClassResource.
//func (sys Sys) LTCFGClass() *LTCFGClassResource {
//	return &sys.lTCFGClass
//}
//
//// lTCFGClassFields returns a configured LTCFGClassFieldsResource.
//func (sys Sys) LTCFGClassFields() *LTCFGClassFieldsResource {
//	return &sys.lTCFGClassFields
//}
//
//// lTCFGInstance returns a configured LTCFGInstanceResource.
//func (sys Sys) LTCFGInstance() *LTCFGInstanceResource {
//	return &sys.lTCFGInstance
//}
//
//// lTCFGInstanceFields returns a configured LTCFGInstanceFieldsResource.
//func (sys Sys) LTCFGInstanceFields() *LTCFGInstanceFieldsResource {
//	return &sys.lTCFGInstanceFields
//}

// license returns a configured LicenseResource.
func (sys Sys) License() *LicenseResource {
	return &sys.license
}

//// log returns a configured LogResource.
//func (sys Sys) Log() *LogResource {
//	return &sys.log
//}
//
//// logConfig returns a configured LogConfigResource.
//func (sys Sys) LogConfig() *LogConfigResource {
//	return &sys.logConfig
//}
//
//// logConfigFilter returns a configured LogConfigFilterResource.
//func (sys Sys) LogConfigFilter() *LogConfigFilterResource {
//	return &sys.logConfigFilter
//}
//
//// logConfigPublisher returns a configured LogConfigPublisherResource.
//func (sys Sys) LogConfigPublisher() *LogConfigPublisherResource {
//	return &sys.logConfigPublisher
//}
//
//// logRotate returns a configured LogRotateResource.
//func (sys Sys) LogRotate() *LogRotateResource {
//	return &sys.logRotate
//}

// managementDHCP returns a configured ManagementDHCPResource.
func (sys Sys) ManagementDHCP() *ManagementDHCPResource {
	return &sys.managementDHCP
}

// managementIP returns a configured ManagementIPResource.
func (sys Sys) ManagementIP() *ManagementIPResource {
	return &sys.managementIP
}

// managementOVSDB returns a configured ManagementOVSDBResource.
func (sys Sys) ManagementOVSDB() *ManagementOVSDBResource {
	return &sys.managementOVSDB
}

// managementRoute returns a configured ManagementRouteResource.
func (sys Sys) ManagementRoute() *ManagementRouteResource {
	return &sys.managementRoute
}

func (sys Sys) MemoryStats() *MemoryStatsResource {
	return &sys.memoryStats
}

// nTP returns a configured NTPResource.
func (sys Sys) NTP() *NTPResource {
	return &sys.nTP
}

// outboundSMTP returns a configured OutboundSMTPResource.
func (sys Sys) OutboundSMTP() *OutboundSMTPResource {
	return &sys.outboundSMTP
}

//
//// pFMan returns a configured PFManResource.
//func (sys Sys) PFMan() *PFManResource {
//	return &sys.pFMan
//}
//
//// pFManConsumer returns a configured PFManConsumerResource.
//func (sys Sys) PFManConsumer() *PFManConsumerResource {
//	return &sys.pFManConsumer
//}
//
//// pFManDevice returns a configured PFManDeviceResource.
//func (sys Sys) PFManDevice() *PFManDeviceResource {
//	return &sys.pFManDevice
//}

// pPTPCallInfo returns a configured PPTPCallInfoResource.
func (sys Sys) PPTPCallInfo() *PPTPCallInfoResource {
	return &sys.pPTPCallInfo
}

//// performance returns a configured PerformanceResource.
//func (sys Sys) Performance() *PerformanceResource {
//	return &sys.performance
//}

// provision returns a configured ProvisionResource.
func (sys Sys) Provision() *ProvisionResource {
	return &sys.provision
}

// rAID returns a configured RAIDResource.
func (sys Sys) RAID() *raid.RAIDResource {
	return &sys.rAID
}

// restrictedModule returns a configured RestrictedModuleResource.
func (sys Sys) RestrictedModule() *RestrictedModuleResource {
	return &sys.restrictedModule
}

//
//// sFlow returns a configured SFlowResource.
//func (sys Sys) SFlow() *SFlowResource {
//	return &sys.sFlow
//}
//
//// sFlowReceiver returns a configured SFlowReceiverResource.
//func (sys Sys) SFlowReceiver() *SFlowReceiverResource {
//	return &sys.sFlowReceiver
//}

// sMTPServer returns a configured SMTPServerResource.
func (sys Sys) SMTPServer() *SMTPServerResource {
	return &sys.sMTPServer
}

// sNMP returns a configured SNMPResource.
func (sys Sys) SNMP() *SNMPResource {
	return &sys.sNMP
}

// sSHD returns a configured SSHDResource.
func (sys Sys) SSHD() *SSHDResource {
	return &sys.sSHD
}

// scriptd returns a configured ScriptdResource.
func (sys Sys) Scriptd() *ScriptdResource {
	return &sys.scriptd
}

// server returns a configured ServerResource.
func (sys Sys) Service() *ServiceResource {
	return &sys.service
}

// software returns a configured SoftwareResource.
func (sys Sys) Software() *software.SoftwareResource {
	return &sys.software
}

// stateMirroring returns a configured StateMirroringResource.
func (sys Sys) StateMirroring() *StateMirroringResource {
	return &sys.stateMirroring
}

// syncSysFiles returns a configured SyncSysFilesResource.
func (sys Sys) SyncSysFiles() *SyncSysFilesResource {
	return &sys.syncSysFiles
}

// syslog returns a configured SyslogResource.
func (sys Sys) Syslog() *SyslogResource {
	return &sys.syslog
}

//// uRLDB returns a configured URLDBResource.
//func (sys Sys) URLDB() *URLDBResource {
//	return &sys.uRLDB
//}
//
//// uRLDBDownloadResult returns a configured URLDBDownloadResultResource.
//func (sys Sys) URLDBDownloadResult() *URLDBDownloadResultResource {
//	return &sys.uRLDBDownloadResult
//}
//
//// uRLDBDownloadSchedule returns a configured URLDBDownloadScheduleResource.
//func (sys Sys) URLDBDownloadSchedule() *URLDBDownloadScheduleResource {
//	return &sys.uRLDBDownloadSchedule
//}
//
//// uRLDBURLCategory returns a configured URLDBURLCategoryResource.
//func (sys Sys) URLDBURLCategory() *URLDBURLCategoryResource {
//	return &sys.uRLDBURLCategory
//}
