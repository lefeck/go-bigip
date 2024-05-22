package profile

import "github.com/lefeck/go-bigip"

const (
	LtmManager = "ltm"
)

// ProfileEndpoint is a commonly used bigip.GetBaseResource(), providing a large number of api resource types
const ProfileEndpoint = "profile"

// 根据这个struct 帮我把下面的补齐
type ProfileResource struct {
	certificateAuthority CertificateAuthorityResource
	clientLDAP           ClientLDAPResource
	clientSSL            ClientSSLResource
	connector            ConnectorResource
	//dhcpv4                DHCPv4Resource
	//dhcpv6                DHCPv6Resource
	diameter DiameterResource
	dns      DNSResource
	//dns_logging              DNSLoggingResource
	//doh_proxy                DOHProxyResource
	//doh_server               DOHServerResource
	fasthttp FastHTTPResource
	fastL4   FastL4Resource
	fix      FIXResource
	ftp      FTPResource
	//georedundancy      GeoRedundancyResource
	gtp              GTPResource
	html             HTMLResource
	http             HTTPResource
	httpCompression  HTTPCompressionResource
	httpProxyConnect HTTPProxyConnectResource
	http2            HTTP2Resource
	http3            HTTP3Resource
	httprouter       HTTPRouterResource
	icap             ICAPResource
	imap             IMAPResource
	//ipother                  IPOtherResource
	//ipsecalg                 IPsecALGResource
	//mblb                     MBLBResource
	mqtt MQTTResource
	//mr_ratelimit             MRRateLimitResource
	netflow NetflowResource
	ntlm    NTLMResource
	ocsp    OCSPResource
	//ocsp_stapling_params     OCSPStaplingParamsResource
	oneConnect OneConnectResource
	pop3       POP3Resource
	pptp       PPTPResource
	qoe        QOEResource
	quic       QUICResource
	radius     RADIUSResource
	//request_adapt            RequestAdaptResource
	//request_log              RequestLogResource
	//response_adapt           ResponseAdaptResource
	rewrite    RewriteResource
	rtsp       RTSPResource
	sctp       SCTPResource
	serverLDAP ServerLDAPResource
	serverSSL  ServerSSLResource
	service    ServiceResource
	sip        SIPResource
	smtps      SMTPSResource
	socks      SocksResource
	//splitsessionclient       SplitSessionClientResource
	//splitsessionserver       SplitSessionServerResource
	statistics      StatisticsResource
	stream          StreamResource
	tcp             TCPResource
	tcpAnalytics    TCPAnalyticsResource
	tdr             TDRResource
	tftp            TFTPResource
	udp             UDPResource
	webAcceleration WebAccelerationResource
	websocket       WebSocketResource
	xml             XMLResource
}

func NewProfile(b *bigip.BigIP) ProfileResource {
	return ProfileResource{
		certificateAuthority: CertificateAuthorityResource{b: b},
		clientLDAP:           ClientLDAPResource{b: b},
		clientSSL:            ClientSSLResource{b: b},
		connector:            ConnectorResource{b: b},
		//dhcpv4:               DHCPv4Resource{b: b},
		//dhcpv6:               DHCPv6Resource{b: b},
		diameter: DiameterResource{b: b},
		dns:      DNSResource{b: b},
		//dns_logging:          DNSLoggingResource{b: b},
		//doh_proxy:            DOHProxyResource{b: b},
		//doh_server:           DOHServerResource{b: b},
		fasthttp: FastHTTPResource{b: b},
		fastL4:   FastL4Resource{b: b},
		fix:      FIXResource{b: b},
		ftp:      FTPResource{b: b},
		//georedundancy:        GeoRedundancyResource{b: b},
		gtp:              GTPResource{b: b},
		html:             HTMLResource{b: b},
		http:             HTTPResource{b: b},
		httpCompression:  HTTPCompressionResource{b: b},
		httpProxyConnect: HTTPProxyConnectResource{b: b},
		http2:            HTTP2Resource{b: b},
		http3:            HTTP3Resource{b: b},
		httprouter:       HTTPRouterResource{b: b},
		icap:             ICAPResource{b: b},
		imap:             IMAPResource{b: b},
		//ipother:              IPOtherResource{b: b},
		//ipsecalg:             IPsecALGResource{b: b},
		//mblb:                 MBLBResource{b: b},
		mqtt: MQTTResource{b: b},
		//mr_ratelimit:         MRRateLimitResource{b: b},
		netflow: NetflowResource{b: b},
		ntlm:    NTLMResource{b: b},
		ocsp:    OCSPResource{b: b},
		//ocsp_stapling_params: OCSPStaplingParamsResource{b: b},
		oneConnect: OneConnectResource{b: b},
		pop3:       POP3Resource{b: b},
		pptp:       PPTPResource{b: b},
		qoe:        QOEResource{b: b},
		quic:       QUICResource{b: b},
		radius:     RADIUSResource{b: b},
		//request_adapt:        RequestAdaptResource{b: b},
		//request_log:          RequestLogResource{b: b},
		//response_adapt:       ResponseAdaptResource{b: b},
		rewrite:    RewriteResource{b: b},
		rtsp:       RTSPResource{b: b},
		sctp:       SCTPResource{b: b},
		serverLDAP: ServerLDAPResource{b: b},
		serverSSL:  ServerSSLResource{b: b},
		service:    ServiceResource{b: b},
		sip:        SIPResource{b: b},
		smtps:      SMTPSResource{b: b},
		socks:      SocksResource{b: b},
		//splitsessionclient:   SplitSessionClientResource{b: b},
		//splitsessionserver:   SplitSessionServerResource{b: b},
		statistics:      StatisticsResource{b: b},
		stream:          StreamResource{b: b},
		tcp:             TCPResource{b: b},
		tcpAnalytics:    TCPAnalyticsResource{b: b},
		tdr:             TDRResource{b: b},
		tftp:            TFTPResource{b: b},
		udp:             UDPResource{b: b},
		webAcceleration: WebAccelerationResource{b: b},
		websocket:       WebSocketResource{b: b},
		xml:             XMLResource{b: b},
	}
}

func (p ProfileResource) CertificateAuthority() *CertificateAuthorityResource {
	return &p.certificateAuthority
}

func (p ProfileResource) ClientLDAP() *ClientLDAPResource { return &p.clientLDAP }

func (p ProfileResource) ClientSSL() *ClientSSLResource { return &p.clientSSL }

func (p ProfileResource) Connector() *ConnectorResource { return &p.connector }

//func (p ProfileResource) DHCPv4() *DHCPv4Resource { return &p.dhcpv4 }
//
//func (p ProfileResource) DHCPv6() *DHCPv6Resource { return &p.dhcpv6 }

func (p ProfileResource) Diameter() *DiameterResource { return &p.diameter }

func (p ProfileResource) DNS() *DNSResource { return &p.dns }

//func (p ProfileResource) DNSLogging() *DNSLoggingResource { return &p.dns_logging }
//
//func (p ProfileResource) DOHProxy() *DOHProxyResource { return &p.doh_proxy }
//
//func (p ProfileResource) DOHServer() *DOHServerResource { return &p.doh_server }

func (p ProfileResource) FastHTTP() *FastHTTPResource { return &p.fasthttp }

func (p ProfileResource) FastL4() *FastL4Resource { return &p.fastL4 }

func (p ProfileResource) FIX() *FIXResource { return &p.fix }

func (p ProfileResource) FTP() *FTPResource { return &p.ftp }

//func (p ProfileResource) GeoRedundancy() *GeoRedundancyResource { return &p.georedundancy }

func (p ProfileResource) GTP() *GTPResource { return &p.gtp }

func (p ProfileResource) HTML() *HTMLResource { return &p.html }

func (p ProfileResource) HTTP() *HTTPResource { return &p.http }

func (p ProfileResource) HTTPCompression() *HTTPCompressionResource { return &p.httpCompression }

func (p ProfileResource) HTTPProxyConnect() *HTTPProxyConnectResource { return &p.httpProxyConnect }

func (p ProfileResource) HTTP2() *HTTP2Resource { return &p.http2 }

func (p ProfileResource) HTTP3() *HTTP3Resource { return &p.http3 }

func (p ProfileResource) HTTPRouter() *HTTPRouterResource { return &p.httprouter }

func (p ProfileResource) ICAP() *ICAPResource { return &p.icap }

func (p ProfileResource) IMAP() *IMAPResource { return &p.imap }

//func (p ProfileResource) IPOther() *IPOtherResource { return &p.ipother }
//
//func (p ProfileResource) IPsecALG() *IPsecALGResource { return &p.ipsecalg }
//
//func (p ProfileResource) MBLB() *MBLBResource { return &p.mblb }

func (p ProfileResource) MQTT() *MQTTResource { return &p.mqtt }

//func (p ProfileResource) MRRateLimit() *MRRateLimitResource { return &p.mr_ratelimit }

func (p ProfileResource) Netflow() *NetflowResource { return &p.netflow }

func (p ProfileResource) NTLM() *NTLMResource { return &p.ntlm }

func (p ProfileResource) OCSP() *OCSPResource { return &p.ocsp }

//func (p ProfileResource) OCSPStaplingParams() *OCSPStaplingParamsResource { return &p.ocsp_stapling_params}

func (p ProfileResource) OneConnect() *OneConnectResource { return &p.oneConnect }

func (p ProfileResource) POP3() *POP3Resource { return &p.pop3 }

func (p ProfileResource) PPTP() *PPTPResource { return &p.pptp }

func (p ProfileResource) QOE() *QOEResource { return &p.qoe }

func (p ProfileResource) QUIC() *QUICResource { return &p.quic }

func (p ProfileResource) Radius() *RADIUSResource { return &p.radius }

//func (p ProfileResource) RequestAdapt() *RequestAdaptResource { return &p.request_adapt }
//
//func (p ProfileResource) RequestLog() *RequestLogResource { return &p.request_log }
//
//func (p ProfileResource) ResponseAdapt() *ResponseAdaptResource { return &p.response_adapt }

func (p ProfileResource) Rewrite() *RewriteResource { return &p.rewrite }

func (p ProfileResource) RTSP() *RTSPResource { return &p.rtsp }

func (p ProfileResource) SCTP() *SCTPResource { return &p.sctp }

func (p ProfileResource) ServerLDAP() *ServerLDAPResource { return &p.serverLDAP }

func (p ProfileResource) ServerSSL() *ServerSSLResource { return &p.serverSSL }

func (p ProfileResource) Service() *ServiceResource { return &p.service }

func (p ProfileResource) SIP() *SIPResource { return &p.sip }

func (p ProfileResource) SMTPS() *SMTPSResource { return &p.smtps }

func (p ProfileResource) SOCKS() *SocksResource { return &p.socks }

//func (p ProfileResource) SplitSessionClient() *SplitSessionClientResource { return &p.splitsessionclient }
//
//func (p ProfileResource) SplitSessionServer() *SplitSessionServerResource { return &p.splitsessionserver }

func (p ProfileResource) Statistics() *StatisticsResource { return &p.statistics }

func (p ProfileResource) Stream() *StreamResource { return &p.stream }

func (p ProfileResource) TCP() *TCPResource { return &p.tcp }

func (p ProfileResource) TCPAnalytics() *TCPAnalyticsResource { return &p.tcpAnalytics }

func (p ProfileResource) TDR() *TDRResource { return &p.tdr }

func (p ProfileResource) TFTP() *TFTPResource { return &p.tftp }

func (p ProfileResource) UDP() *UDPResource { return &p.udp }

func (p ProfileResource) WebAcceleration() *WebAccelerationResource { return &p.webAcceleration }

func (p ProfileResource) WebSocket() *WebSocketResource { return &p.websocket }

func (p ProfileResource) XML() *XMLResource { return &p.xml }
