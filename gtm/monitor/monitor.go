package monitor

import "github.com/lefeck/go-bigip"

// MonitorEndpoint represents the REST resource for managing monitor.
const MonitorEndpoint = "monitor"
const GTMManager = "gtm"

// MonitorResource struct is a container for all the monitoring resources
type MonitorResource struct {
	bigip            BigIPResource
	bigIPLink        BigIPLinkResource
	external         ExternalResource
	ftp              FTPResource
	firepass         FirepassResource
	gtp              GTPResource
	http             HTTPResource
	https            HTTPSResource
	icmp             ICMPResource
	imap             IMAPResource
	ldap             LDAPResource
	mssql            MSSQLResource
	mysql            MySQLResource
	nntp             NNTPResource
	none             NoneResource
	oracle           OracleResource
	pop3             POP3Resource
	postgreSQL       PostgreSQLResource
	radius           RadiusResource
	radiusAccounting RadiusAccountingResource
	realServer       RealServerResource
	sip              SIPResource
	smtp             SMTPResource
	snmp             SNMPResource
	snmpLink         SNMPLinkResource
	soap             SOAPResource
	tcp              TCPResource
	tcpHalf          TCPHalfResource
	udp              UDPResource
	wap              WAPResource
	wmi              WMIResource
	scripted         ScriptedResource
}

// NewMonitor constructs a new instance of MonitorResource with a given bigip.BigIP instance
func NewMonitor(b *bigip.BigIP) MonitorResource {
	return MonitorResource{
		bigip:            BigIPResource{b: b},
		bigIPLink:        BigIPLinkResource{b: b},
		external:         ExternalResource{b: b},
		ftp:              FTPResource{b: b},
		firepass:         FirepassResource{b: b},
		gtp:              GTPResource{b: b},
		http:             HTTPResource{b: b},
		https:            HTTPSResource{b: b},
		icmp:             ICMPResource{b: b},
		imap:             IMAPResource{b: b},
		ldap:             LDAPResource{b: b},
		mssql:            MSSQLResource{b: b},
		mysql:            MySQLResource{b: b},
		nntp:             NNTPResource{b: b},
		none:             NoneResource{b: b},
		oracle:           OracleResource{b: b},
		pop3:             POP3Resource{b: b},
		postgreSQL:       PostgreSQLResource{b: b},
		radius:           RadiusResource{b: b},
		radiusAccounting: RadiusAccountingResource{b: b},
		realServer:       RealServerResource{b: b},
		sip:              SIPResource{b: b},
		smtp:             SMTPResource{b: b},
		snmp:             SNMPResource{b: b},
		snmpLink:         SNMPLinkResource{b: b},
		soap:             SOAPResource{b: b},
		tcp:              TCPResource{b: b},
		tcpHalf:          TCPHalfResource{b: b},
		udp:              UDPResource{b: b},
		wap:              WAPResource{b: b},
		wmi:              WMIResource{b: b},
		scripted:         ScriptedResource{b: b},
	}
}

// BIGIP returns a reference to the BigIPResource instance
func (m *MonitorResource) BIGIP() *BigIPResource {
	return &m.bigip
}

// Getter functions for other resource types

// BigIPLink returns a reference to the BigIPLinkResource instance
func (m *MonitorResource) BigIPLink() *BigIPLinkResource {
	return &m.bigIPLink
}

// External returns a reference to the ExternalResource instance
func (m *MonitorResource) External() *ExternalResource {
	return &m.external
}

// FTP returns a reference to the FTPResource instance
func (m *MonitorResource) FTP() *FTPResource {
	return &m.ftp
}

// Firepass returns a reference to the FirepassResource instance
func (m *MonitorResource) Firepass() *FirepassResource {
	return &m.firepass
}

// GTP returns a reference to the GTPResource instance
func (m *MonitorResource) GTP() *GTPResource {
	return &m.gtp
}

// HTTP returns a reference to the HTTPResource instance
func (m *MonitorResource) HTTP() *HTTPResource {
	return &m.http
}

// HTTPS returns a reference to the HTTPSResource instance
func (m *MonitorResource) HTTPS() *HTTPSResource {
	return &m.https
}

// ICMP returns a reference to the ICMPResource instance
func (m *MonitorResource) ICMP() *ICMPResource {
	return &m.icmp
}

// IMAP returns a reference to the IMAPResource instance
func (m *MonitorResource) IMAP() *IMAPResource {
	return &m.imap
}

// LDAP returns a reference to the LDAPResource instance
func (m *MonitorResource) LDAP() *LDAPResource {
	return &m.ldap
}

// MSSQL returns a reference to the MSSQLResource instance
func (m *MonitorResource) MSSQL() *MSSQLResource {
	return &m.mssql
}

// MySQL returns a reference to the MySQLResource instance
func (m *MonitorResource) MySQL() *MySQLResource {
	return &m.mysql
}

// NNTP returns a reference to the NNTPResource instance
func (m *MonitorResource) NNTP() *NNTPResource {
	return &m.nntp
}

// None returns a reference to the NoneResource instance
func (m *MonitorResource) None() *NoneResource {
	return &m.none
}

// Oracle returns a reference to the OracleResource instance
func (m *MonitorResource) Oracle() *OracleResource {
	return &m.oracle
}

// POP3 returns a reference to the POP3Resource instance
func (m *MonitorResource) POP3() *POP3Resource {
	return &m.pop3
}

// PostgreSQL returns a reference to the PostgreSQLResource instance
func (m *MonitorResource) PostgreSQL() *PostgreSQLResource {
	return &m.postgreSQL
}

// Radius returns a reference to the RadiusResource instance
func (m *MonitorResource) Radius() *RadiusResource {
	return &m.radius
}

// RadiusAccounting returns a reference to the RadiusAccountingResource instance
func (m *MonitorResource) RadiusAccounting() *RadiusAccountingResource {
	return &m.radiusAccounting
}

// RealServer returns a reference to the RealServerResource instance
func (m *MonitorResource) RealServer() *RealServerResource {
	return &m.realServer
}

// SIP returns a reference to the SIPResource instance
func (m *MonitorResource) SIP() *SIPResource {
	return &m.sip
}

// SMTP returns a reference to the SMTPResource instance
func (m *MonitorResource) SMTP() *SMTPResource {
	return &m.smtp
}

// SNMP returns a reference to the SNMPResource instance
func (m *MonitorResource) SNMP() *SNMPResource {
	return &m.snmp
}

// SNMPLink returns a reference to the SNMPLinkResource instance
func (m *MonitorResource) SNMPLink() *SNMPLinkResource {
	return &m.snmpLink
}

// SOAP returns a reference to the SOAPResource instance
func (m *MonitorResource) SOAP() *SOAPResource {
	return &m.soap
}

// TCP returns a reference to the TCPResource instance
func (m *MonitorResource) TCP() *TCPResource {
	return &m.tcp
}

// TCPHalf returns a reference to the TCPHalfResource instance
func (m *MonitorResource) TCPHalf() *TCPHalfResource {
	return &m.tcpHalf
}

// UDP returns a reference to the UDPResource instance
func (m *MonitorResource) UDP() *UDPResource {
	return &m.udp
}

// WAP returns a reference to the WAPResource instance
func (m *MonitorResource) WAP() *WAPResource {
	return &m.wap
}

// WMI returns a reference to the WMIResource instance
func (m *MonitorResource) WMI() *WMIResource {
	return &m.wmi
}

// Scripted returns a reference to the ScriptedResource instance
func (m *MonitorResource) Scripted() *ScriptedResource {
	return &m.scripted
}
