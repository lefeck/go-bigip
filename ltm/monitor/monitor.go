package monitor

import (
	"github.com/lefeck/go-bigip"
)

// MonitorEndpoint is a commonly used basepath, providing a large number of api resource types
const MonitorEndpoint = "monitor"

type MonitorResoucre struct {
	b *bigip.BigIP

	monitorDiameter         MonitorDiameterResource
	monitorDNS              MonitorDNSResource
	monitorExternal         MonitorExternalResource
	monitorFirepass         MonitorFirepassResource
	monitorFTP              MonitorFTPResource
	monitorGatewayICMP      MonitorGatewayICMPResource
	monitorHTTP             MonitorHTTPResource
	monitorHTTPS            MonitorHTTPSResource
	monitorICMP             MonitorICMPResource
	monitorIMAP             MonitorIMAPResource
	monitorInband           MonitorInbandResource
	monitorLDAP             MonitorLDAPResource
	monitorModuleScore      MonitorModuleScoreResource
	monitorMSSQL            MonitorMSSQLResource
	monitorMySQL            MonitorMySQLResource
	monitorNNTP             MonitorNNTPResource
	monitorOracle           MonitorOracleResource
	monitorPOP3             MonitorPOP3Resource
	monitorPostgreSQL       MonitorPostgreSQLResource
	monitorRadiusAccounting MonitorRadiusAccountingResource
	monitorRadius           MonitorRadiusResource
	monitorRealServer       MonitorRealServerResource
	monitorRPC              MonitorRPCResource
	monitorSASP             MonitorSASPResource
	monitorScripted         MonitorScriptedResource
	monitorSIP              MonitorSIPResource
	monitorSMB              MonitorSMBResource
	monitorSMTP             MonitorSMTPResource
	monitorSNMPDCABase      MonitorSNMPDCABaseResource
	monitorSNMPDCA          MonitorSNMPDCAResource
	monitorSOAP             MonitorSOAPResource
	monitorTCPEcho          MonitorTCPEchoResource
	monitorTCP              MonitorTCPResource
	monitorTCPHalfOpen      MonitorTCPHalfOpenResource
	monitorUDP              MonitorUDPResource
	monitorVirtualLocation  MonitorVirtualLocationResource
	monitorWAP              MonitorWAPResource
	monitorWMI              MonitorWMIResource
}

func NewMonitor(b *bigip.BigIP) MonitorResoucre {
	return MonitorResoucre{
		b:                       b,
		monitorDiameter:         MonitorDiameterResource{b: b},
		monitorDNS:              MonitorDNSResource{b: b},
		monitorExternal:         MonitorExternalResource{b: b},
		monitorFirepass:         MonitorFirepassResource{b: b},
		monitorFTP:              MonitorFTPResource{b: b},
		monitorGatewayICMP:      MonitorGatewayICMPResource{b: b},
		monitorHTTP:             MonitorHTTPResource{b: b},
		monitorHTTPS:            MonitorHTTPSResource{b: b},
		monitorICMP:             MonitorICMPResource{b: b},
		monitorIMAP:             MonitorIMAPResource{b: b},
		monitorInband:           MonitorInbandResource{b: b},
		monitorLDAP:             MonitorLDAPResource{b: b},
		monitorModuleScore:      MonitorModuleScoreResource{b: b},
		monitorMSSQL:            MonitorMSSQLResource{b: b},
		monitorMySQL:            MonitorMySQLResource{b: b},
		monitorNNTP:             MonitorNNTPResource{b: b},
		monitorOracle:           MonitorOracleResource{b: b},
		monitorPOP3:             MonitorPOP3Resource{b: b},
		monitorPostgreSQL:       MonitorPostgreSQLResource{b: b},
		monitorRadiusAccounting: MonitorRadiusAccountingResource{b: b},
		monitorRadius:           MonitorRadiusResource{b: b},
		monitorRealServer:       MonitorRealServerResource{b: b},
		monitorRPC:              MonitorRPCResource{b: b},
		monitorSASP:             MonitorSASPResource{b: b},
		monitorScripted:         MonitorScriptedResource{b: b},
		monitorSIP:              MonitorSIPResource{b: b},
		monitorSMB:              MonitorSMBResource{b: b},
		monitorSMTP:             MonitorSMTPResource{b: b},
		monitorSNMPDCABase:      MonitorSNMPDCABaseResource{b: b},
		monitorSNMPDCA:          MonitorSNMPDCAResource{b: b},
		monitorSOAP:             MonitorSOAPResource{b: b},
		monitorTCPEcho:          MonitorTCPEchoResource{b: b},
		monitorTCP:              MonitorTCPResource{b: b},
		monitorTCPHalfOpen:      MonitorTCPHalfOpenResource{b: b},
		monitorUDP:              MonitorUDPResource{b: b},
		monitorVirtualLocation:  MonitorVirtualLocationResource{b: b},
		monitorWAP:              MonitorWAPResource{b: b},
		monitorWMI:              MonitorWMIResource{b: b},
	}
}

type Monitor interface {
	MonitorDiameter() *MonitorDiameterResource
	MonitorDNS() *MonitorDNSResource
	MonitorExternal() *MonitorExternalResource
	MonitorFirepass() *MonitorFirepassResource
	MonitorFTP() *MonitorFTPResource
	MonitorGatewayICMP() *MonitorGatewayICMPResource
	MonitorHTTP() *MonitorHTTPResource
	MonitorHTTPS() *MonitorHTTPSResource
	MonitorICMP() *MonitorICMPResource
	MonitorIMAP() *MonitorIMAPResource
	MonitorInband() *MonitorInbandResource
	MonitorLDAP() *MonitorLDAPResource
	MonitorModuleScore() *MonitorModuleScoreResource
	MonitorMSSQL() *MonitorMSSQLResource
	MonitorMySQL() *MonitorMySQLResource
	MonitorNNTP() *MonitorNNTPResource
	MonitorOracle() *MonitorOracleResource
	MonitorPOP3() *MonitorPOP3Resource
	MonitorPostgreSQL() *MonitorPostgreSQLResource
	MonitorRadiusAccounting() *MonitorRadiusAccountingResource
	MonitorRadius() *MonitorRadiusResource
	MonitorRealServer() *MonitorRealServerResource
	MonitorRPC() *MonitorRPCResource
	MonitorSASP() *MonitorSASPResource
	MonitorScripted() *MonitorScriptedResource
	MonitorSIP() *MonitorSIPResource
	MonitorSMB() *MonitorSMBResource
	MonitorSMTP() *MonitorSMTPResource
	MonitorSNMPDCABase() *MonitorSNMPDCABaseResource
	MonitorSNMPDCA() *MonitorSNMPDCAResource
	MonitorSOAP() *MonitorSOAPResource
	MonitorTCPEcho() *MonitorTCPEchoResource
	MonitorTCP() *MonitorTCPResource
	MonitorTCPHalfOpen() *MonitorTCPHalfOpenResource
	MonitorUDP() *MonitorUDPResource
	MonitorVirtualLocation() *MonitorVirtualLocationResource
	MonitorWAP() *MonitorWAPResource
	MonitorWMI() *MonitorWMIResource
}

//var _ Monitor = MonitorResoucre{}

func (m MonitorResoucre) MonitorDiameter() *MonitorDiameterResource {
	return &m.monitorDiameter
}

func (m MonitorResoucre) MonitorDNS() *MonitorDNSResource {
	return &m.monitorDNS
}

func (m MonitorResoucre) MonitorExternal() *MonitorExternalResource {
	return &m.monitorExternal
}

func (m MonitorResoucre) MonitorFirepass() *MonitorFirepassResource {
	return &m.monitorFirepass
}

func (m MonitorResoucre) MonitorFTP() *MonitorFTPResource {
	return &m.monitorFTP
}

func (m MonitorResoucre) MonitorGatewayICMP() *MonitorGatewayICMPResource {
	return &m.monitorGatewayICMP
}

func (m MonitorResoucre) MonitorHTTP() *MonitorHTTPResource {
	return &m.monitorHTTP
}

func (m MonitorResoucre) MonitorHTTPS() *MonitorHTTPSResource {
	return &m.monitorHTTPS
}

func (m MonitorResoucre) MonitorICMP() *MonitorICMPResource {
	return &m.monitorICMP
}

func (m MonitorResoucre) MonitorIMAP() *MonitorIMAPResource {
	return &m.monitorIMAP
}

func (m MonitorResoucre) MonitorInband() *MonitorInbandResource {
	return &m.monitorInband
}

func (m MonitorResoucre) MonitorLDAP() *MonitorLDAPResource {
	return &m.monitorLDAP
}

func (m MonitorResoucre) MonitorModuleScore() *MonitorModuleScoreResource {
	return &m.monitorModuleScore
}

func (m MonitorResoucre) MonitorMSSQL() *MonitorMSSQLResource {
	return &m.monitorMSSQL
}

func (m MonitorResoucre) MonitorMySQL() *MonitorMySQLResource {
	return &m.monitorMySQL
}

func (m MonitorResoucre) MonitorNNTP() *MonitorNNTPResource {
	return &m.monitorNNTP
}

func (m MonitorResoucre) MonitorOracle() *MonitorOracleResource {
	return &m.monitorOracle
}

func (m MonitorResoucre) MonitorPOP3() *MonitorPOP3Resource {
	return &m.monitorPOP3
}

func (m MonitorResoucre) MonitorPostgreSQL() *MonitorPostgreSQLResource {
	return &m.monitorPostgreSQL
}

func (m MonitorResoucre) MonitorRadiusAccounting() *MonitorRadiusAccountingResource {
	return &m.monitorRadiusAccounting
}

func (m MonitorResoucre) MonitorRadius() *MonitorRadiusResource {
	return &m.monitorRadius
}

func (m MonitorResoucre) MonitorRealServer() *MonitorRealServerResource {
	return &m.monitorRealServer
}

func (m MonitorResoucre) MonitorRPC() *MonitorRPCResource {
	return &m.monitorRPC
}

func (m MonitorResoucre) MonitorSASP() *MonitorSASPResource {
	return &m.monitorSASP
}

func (m MonitorResoucre) MonitorScripted() *MonitorScriptedResource {
	return &m.monitorScripted
}

func (m MonitorResoucre) MonitorSIP() *MonitorSIPResource {
	return &m.monitorSIP
}

func (m MonitorResoucre) MonitorSMB() *MonitorSMBResource {
	return &m.monitorSMB
}

func (m MonitorResoucre) MonitorSMTP() *MonitorSMTPResource {
	return &m.monitorSMTP
}

func (m MonitorResoucre) MonitorSNMPDCABase() *MonitorSNMPDCABaseResource {
	return &m.monitorSNMPDCABase
}

func (m MonitorResoucre) MonitorSNMPDCA() *MonitorSNMPDCAResource {
	return &m.monitorSNMPDCA
}

func (m MonitorResoucre) MonitorSOAP() *MonitorSOAPResource {
	return &m.monitorSOAP
}

func (m MonitorResoucre) MonitorTCPEcho() *MonitorTCPEchoResource {
	return &m.monitorTCPEcho
}

func (m MonitorResoucre) MonitorTCP() *MonitorTCPResource {
	return &m.monitorTCP
}

func (m MonitorResoucre) MonitorTCPHalfOpen() *MonitorTCPHalfOpenResource {
	return &m.monitorTCPHalfOpen
}

func (m MonitorResoucre) MonitorUDP() *MonitorUDPResource {
	return &m.monitorUDP
}

func (m MonitorResoucre) MonitorVirtualLocation() *MonitorVirtualLocationResource {
	return &m.monitorVirtualLocation
}

func (m MonitorResoucre) MonitorWAP() *MonitorWAPResource {
	return &m.monitorWAP
}

func (m MonitorResoucre) MonitorWMI() *MonitorWMIResource {
	return &m.monitorWMI
}
