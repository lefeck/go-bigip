package monitor

import "github.com/lefeck/go-bigip"

const LtmManager = "ltm"

// Endpoint is a commonly used bigip.GetBaseResource(), providing a large number of api resource types
const MonitorEndpoint = "monitor"

type MonitorResource struct {
	diameter         DiameterResource
	dns              DNSResource
	external         ExternalResource
	firepass         FirepassResource
	ftp              FTPResource
	gatewayicmp      GatewayICMPResource
	http             HTTPResource
	https            HTTPSResource
	icmp             ICMPResource
	imap             IMAPResource
	inband           InbandResource
	ldap             LDAPResource
	modulescore      ModuleScoreResource
	mssql            MSSQLResource
	mysql            MySQLResource
	nntp             NNTPResource
	oracle           OracleResource
	pop3             POP3Resource
	postgresql       PostgreSQLResource
	radiusaccounting RadiusAccountingResource
	radius           RadiusResource
	realserver       RealServerResource
	rpc              RPCResource
	sasp             SASPResource
	scripted         ScriptedResource
	sip              SIPResource
	smb              SMBResource
	smtp             SMTPResource
	snmpdcabase      SNMPDCABaseResource
	snmpdca          SNMPDCAResource
	soap             SOAPResource
	tcpecho          TCPEchoResource
	tcp              TCPResource
	tcphalfopen      TCPHalfOpenResource
	udp              UDPResource
	virtuallocation  VirtualLocationResource
	wap              WAPResource
	wmi              WMIResource
}

func NewMonitor(b *bigip.BigIP) MonitorResource {
	return MonitorResource{
		diameter:         DiameterResource{b: b},
		dns:              DNSResource{b: b},
		external:         ExternalResource{b: b},
		firepass:         FirepassResource{b: b},
		ftp:              FTPResource{b: b},
		gatewayicmp:      GatewayICMPResource{b: b},
		http:             HTTPResource{b: b},
		https:            HTTPSResource{b: b},
		icmp:             ICMPResource{b: b},
		imap:             IMAPResource{b: b},
		inband:           InbandResource{b: b},
		ldap:             LDAPResource{b: b},
		modulescore:      ModuleScoreResource{b: b},
		mssql:            MSSQLResource{b: b},
		mysql:            MySQLResource{b: b},
		nntp:             NNTPResource{b: b},
		oracle:           OracleResource{b: b},
		pop3:             POP3Resource{b: b},
		postgresql:       PostgreSQLResource{b: b},
		radiusaccounting: RadiusAccountingResource{b: b},
		radius:           RadiusResource{b: b},
		realserver:       RealServerResource{b: b},
		rpc:              RPCResource{b: b},
		sasp:             SASPResource{b: b},
		scripted:         ScriptedResource{b: b},
		sip:              SIPResource{b: b},
		smb:              SMBResource{b: b},
		smtp:             SMTPResource{b: b},
		snmpdcabase:      SNMPDCABaseResource{b: b},
		snmpdca:          SNMPDCAResource{b: b},
		soap:             SOAPResource{b: b},
		tcpecho:          TCPEchoResource{b: b},
		tcp:              TCPResource{b: b},
		tcphalfopen:      TCPHalfOpenResource{b: b},
		udp:              UDPResource{b: b},
		virtuallocation:  VirtualLocationResource{b: b},
		wap:              WAPResource{b: b},
		wmi:              WMIResource{b: b},
	}
}

type Monitor interface {
	Diameter() *DiameterResource
	DNS() *DNSResource
	External() *ExternalResource
	Firepass() *FirepassResource
	FTP() *FTPResource
	GatewayICMP() *GatewayICMPResource
	HTTP() *HTTPResource
	HTTPS() *HTTPSResource
	ICMP() *ICMPResource
	IMAP() *IMAPResource
	Inband() *InbandResource
	LDAP() *LDAPResource
	ModuleScore() *ModuleScoreResource
	MSSQL() *MSSQLResource
	MySQL() *MySQLResource
	NNTP() *NNTPResource
	Oracle() *OracleResource
	POP3() *POP3Resource
	PostgreSQL() *PostgreSQLResource
	RadiusAccounting() *RadiusAccountingResource
	Radius() *RadiusResource
	RealServer() *RealServerResource
	RPC() *RPCResource
	SASP() *SASPResource
	Scripted() *ScriptedResource
	SIP() *SIPResource
	SMB() *SMBResource
	SMTP() *SMTPResource
	SNMPDCABase() *SNMPDCABaseResource
	SNMPDCA() *SNMPDCAResource
	SOAP() *SOAPResource
	TCPEcho() *TCPEchoResource
	TCP() *TCPResource
	TCPHalfOpen() *TCPHalfOpenResource
	UDP() *UDPResource
	VirtualLocation() *VirtualLocationResource
	WAP() *WAPResource
	WMI() *WMIResource
}

var _ Monitor = MonitorResource{}

func (m MonitorResource) Diameter() *DiameterResource {
	return &m.diameter
}

func (m MonitorResource) DNS() *DNSResource {
	return &m.dns
}

func (m MonitorResource) External() *ExternalResource {
	return &m.external
}

func (m MonitorResource) Firepass() *FirepassResource {
	return &m.firepass
}

func (m MonitorResource) FTP() *FTPResource {
	return &m.ftp
}

func (m MonitorResource) GatewayICMP() *GatewayICMPResource {
	return &m.gatewayicmp
}

func (m MonitorResource) HTTP() *HTTPResource {
	return &m.http
}

func (m MonitorResource) HTTPS() *HTTPSResource {
	return &m.https
}

func (m MonitorResource) ICMP() *ICMPResource {
	return &m.icmp
}

func (m MonitorResource) IMAP() *IMAPResource {
	return &m.imap
}

func (m MonitorResource) Inband() *InbandResource {
	return &m.inband
}

func (m MonitorResource) LDAP() *LDAPResource {
	return &m.ldap
}

func (m MonitorResource) ModuleScore() *ModuleScoreResource {
	return &m.modulescore
}

func (m MonitorResource) MSSQL() *MSSQLResource {
	return &m.mssql
}

func (m MonitorResource) MySQL() *MySQLResource {
	return &m.mysql
}

func (m MonitorResource) NNTP() *NNTPResource {
	return &m.nntp
}

func (m MonitorResource) Oracle() *OracleResource {
	return &m.oracle
}

func (m MonitorResource) POP3() *POP3Resource {
	return &m.pop3
}

func (m MonitorResource) PostgreSQL() *PostgreSQLResource {
	return &m.postgresql
}

func (m MonitorResource) RadiusAccounting() *RadiusAccountingResource {
	return &m.radiusaccounting
}

func (m MonitorResource) Radius() *RadiusResource {
	return &m.radius
}

func (m MonitorResource) RealServer() *RealServerResource {
	return &m.realserver
}

func (m MonitorResource) RPC() *RPCResource {
	return &m.rpc
}

func (m MonitorResource) SASP() *SASPResource {
	return &m.sasp
}

func (m MonitorResource) Scripted() *ScriptedResource {
	return &m.scripted
}

func (m MonitorResource) SIP() *SIPResource {
	return &m.sip
}

func (m MonitorResource) SMB() *SMBResource {
	return &m.smb
}

func (m MonitorResource) SMTP() *SMTPResource {
	return &m.smtp
}

func (m MonitorResource) SNMPDCABase() *SNMPDCABaseResource {
	return &m.snmpdcabase
}

func (m MonitorResource) SNMPDCA() *SNMPDCAResource {
	return &m.snmpdca
}

func (m MonitorResource) SOAP() *SOAPResource {
	return &m.soap
}

func (m MonitorResource) TCPEcho() *TCPEchoResource {
	return &m.tcpecho
}

func (m MonitorResource) TCP() *TCPResource {
	return &m.tcp
}

func (m MonitorResource) TCPHalfOpen() *TCPHalfOpenResource {
	return &m.tcphalfopen
}

func (m MonitorResource) UDP() *UDPResource {
	return &m.udp
}

func (m MonitorResource) VirtualLocation() *VirtualLocationResource {
	return &m.virtuallocation
}

func (m MonitorResource) WAP() *WAPResource {
	return &m.wap
}

func (m MonitorResource) WMI() *WMIResource {
	return &m.wmi
}
