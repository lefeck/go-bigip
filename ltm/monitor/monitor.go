package monitor

import "github.com/lefeck/go-bigip"

// MonitorEndpoint is a commonly used basepath, providing a large number of api resource types
const MonitorEndpoint = "monitor"

type MonitorResource struct {
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

func NewMonitor(b *bigip.BigIP) MonitorResource {
	return MonitorResource{
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
	Diameter() *MonitorDiameterResource
	DNS() *MonitorDNSResource
	External() *MonitorExternalResource
	Firepass() *MonitorFirepassResource
	FTP() *MonitorFTPResource
	GatewayICMP() *MonitorGatewayICMPResource
	HTTP() *MonitorHTTPResource
	HTTPS() *MonitorHTTPSResource
	ICMP() *MonitorICMPResource
	IMAP() *MonitorIMAPResource
	Inband() *MonitorInbandResource
	LDAP() *MonitorLDAPResource
	ModuleScore() *MonitorModuleScoreResource
	MSSQL() *MonitorMSSQLResource
	MySQL() *MonitorMySQLResource
	NNTP() *MonitorNNTPResource
	Oracle() *MonitorOracleResource
	POP3() *MonitorPOP3Resource
	PostgreSQL() *MonitorPostgreSQLResource
	RadiusAccounting() *MonitorRadiusAccountingResource
	Radius() *MonitorRadiusResource
	RealServer() *MonitorRealServerResource
	RPC() *MonitorRPCResource
	SASP() *MonitorSASPResource
	Scripted() *MonitorScriptedResource
	SIP() *MonitorSIPResource
	SMB() *MonitorSMBResource
	SMTP() *MonitorSMTPResource
	SNMPDCABase() *MonitorSNMPDCABaseResource
	SNMPDCA() *MonitorSNMPDCAResource
	SOAP() *MonitorSOAPResource
	TCPEcho() *MonitorTCPEchoResource
	TCP() *MonitorTCPResource
	TCPHalfOpen() *MonitorTCPHalfOpenResource
	UDP() *MonitorUDPResource
	VirtualLocation() *MonitorVirtualLocationResource
	WAP() *MonitorWAPResource
	WMI() *MonitorWMIResource
}

var _ Monitor = MonitorResource{}

func (m MonitorResource) Diameter() *MonitorDiameterResource {
	return &m.monitorDiameter
}

func (m MonitorResource) DNS() *MonitorDNSResource {
	return &m.monitorDNS
}

func (m MonitorResource) External() *MonitorExternalResource {
	return &m.monitorExternal
}

func (m MonitorResource) Firepass() *MonitorFirepassResource {
	return &m.monitorFirepass
}

func (m MonitorResource) FTP() *MonitorFTPResource {
	return &m.monitorFTP
}

func (m MonitorResource) GatewayICMP() *MonitorGatewayICMPResource {
	return &m.monitorGatewayICMP
}

func (m MonitorResource) HTTP() *MonitorHTTPResource {
	return &m.monitorHTTP
}

func (m MonitorResource) HTTPS() *MonitorHTTPSResource {
	return &m.monitorHTTPS
}

func (m MonitorResource) ICMP() *MonitorICMPResource {
	return &m.monitorICMP
}

func (m MonitorResource) IMAP() *MonitorIMAPResource {
	return &m.monitorIMAP
}

func (m MonitorResource) Inband() *MonitorInbandResource {
	return &m.monitorInband
}

func (m MonitorResource) LDAP() *MonitorLDAPResource {
	return &m.monitorLDAP
}

func (m MonitorResource) ModuleScore() *MonitorModuleScoreResource {
	return &m.monitorModuleScore
}

func (m MonitorResource) MSSQL() *MonitorMSSQLResource {
	return &m.monitorMSSQL
}

func (m MonitorResource) MySQL() *MonitorMySQLResource {
	return &m.monitorMySQL
}

func (m MonitorResource) NNTP() *MonitorNNTPResource {
	return &m.monitorNNTP
}

func (m MonitorResource) Oracle() *MonitorOracleResource {
	return &m.monitorOracle
}

func (m MonitorResource) POP3() *MonitorPOP3Resource {
	return &m.monitorPOP3
}

func (m MonitorResource) PostgreSQL() *MonitorPostgreSQLResource {
	return &m.monitorPostgreSQL
}

func (m MonitorResource) RadiusAccounting() *MonitorRadiusAccountingResource {
	return &m.monitorRadiusAccounting
}

func (m MonitorResource) Radius() *MonitorRadiusResource {
	return &m.monitorRadius
}

func (m MonitorResource) RealServer() *MonitorRealServerResource {
	return &m.monitorRealServer
}

func (m MonitorResource) RPC() *MonitorRPCResource {
	return &m.monitorRPC
}

func (m MonitorResource) SASP() *MonitorSASPResource {
	return &m.monitorSASP
}

func (m MonitorResource) Scripted() *MonitorScriptedResource {
	return &m.monitorScripted
}

func (m MonitorResource) SIP() *MonitorSIPResource {
	return &m.monitorSIP
}

func (m MonitorResource) SMB() *MonitorSMBResource {
	return &m.monitorSMB
}

func (m MonitorResource) SMTP() *MonitorSMTPResource {
	return &m.monitorSMTP
}

func (m MonitorResource) SNMPDCABase() *MonitorSNMPDCABaseResource {
	return &m.monitorSNMPDCABase
}

func (m MonitorResource) SNMPDCA() *MonitorSNMPDCAResource {
	return &m.monitorSNMPDCA
}

func (m MonitorResource) SOAP() *MonitorSOAPResource {
	return &m.monitorSOAP
}

func (m MonitorResource) TCPEcho() *MonitorTCPEchoResource {
	return &m.monitorTCPEcho
}

func (m MonitorResource) TCP() *MonitorTCPResource {
	return &m.monitorTCP
}

func (m MonitorResource) TCPHalfOpen() *MonitorTCPHalfOpenResource {
	return &m.monitorTCPHalfOpen
}

func (m MonitorResource) UDP() *MonitorUDPResource {
	return &m.monitorUDP
}

func (m MonitorResource) VirtualLocation() *MonitorVirtualLocationResource {
	return &m.monitorVirtualLocation
}

func (m MonitorResource) WAP() *MonitorWAPResource {
	return &m.monitorWAP
}

func (m MonitorResource) WMI() *MonitorWMIResource {
	return &m.monitorWMI
}
