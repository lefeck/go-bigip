package profile

import "github.com/lefeck/go-bigip"

type TCPList struct {
	Items    []TCP  `json:"items,omitempty"`
	Kind     string `json:"kind,omitempty"`
	SelfLink string `json:"selflink,omitempty"`
}

type TCP struct {
	Kind                  string `json:"kind"`
	Name                  string `json:"name"`
	Partition             string `json:"partition"`
	FullPath              string `json:"fullPath"`
	Generation            int    `json:"generation"`
	SelfLink              string `json:"selfLink"`
	Abc                   string `json:"abc"`
	AckOnPush             string `json:"ackOnPush"`
	AppService            string `json:"appService"`
	AutoProxyBufferSize   string `json:"autoProxyBufferSize"`
	AutoReceiveWindowSize string `json:"autoReceiveWindowSize"`
	AutoSendBufferSize    string `json:"autoSendBufferSize"`
	CloseWaitTimeout      int    `json:"closeWaitTimeout"`
	CmetricsCache         string `json:"cmetricsCache"`
	CmetricsCacheTimeout  int    `json:"cmetricsCacheTimeout"`
	CongestionControl     string `json:"congestionControl"`
	DefaultsFrom          string `json:"defaultsFrom"`
	DefaultsFromReference struct {
		Link string `json:"link"`
	} `json:"defaultsFromReference,omitempty"`
	DeferredAccept           string `json:"deferredAccept"`
	DelayWindowControl       string `json:"delayWindowControl"`
	DelayedAcks              string `json:"delayedAcks"`
	Description              string `json:"description"`
	Dsack                    string `json:"dsack"`
	EarlyRetransmit          string `json:"earlyRetransmit"`
	Ecn                      string `json:"ecn"`
	EnhancedLossRecovery     string `json:"enhancedLossRecovery"`
	FastOpen                 string `json:"fastOpen"`
	FastOpenCookieExpiration int    `json:"fastOpenCookieExpiration"`
	FinWait2Timeout          int    `json:"finWait_2Timeout"`
	FinWaitTimeout           int    `json:"finWaitTimeout"`
	HardwareSynCookie        string `json:"hardwareSynCookie"`
	IdleTimeout              int    `json:"idleTimeout"`
	InitCwnd                 int    `json:"initCwnd"`
	InitRwnd                 int    `json:"initRwnd"`
	IPDfMode                 string `json:"ipDfMode"`
	IPTosToClient            string `json:"ipTosToClient"`
	IPTTLMode                string `json:"ipTtlMode"`
	IPTTLV4                  int    `json:"ipTtlV4"`
	IPTTLV6                  int    `json:"ipTtlV6"`
	KeepAliveInterval        int    `json:"keepAliveInterval"`
	LimitedTransmit          string `json:"limitedTransmit"`
	LinkQosToClient          string `json:"linkQosToClient"`
	MaxRetrans               int    `json:"maxRetrans"`
	MaxSegmentSize           int    `json:"maxSegmentSize"`
	Md5Signature             string `json:"md5Signature"`
	MinimumRto               int    `json:"minimumRto"`
	Mptcp                    string `json:"mptcp"`
	MptcpCsum                string `json:"mptcpCsum"`
	MptcpCsumVerify          string `json:"mptcpCsumVerify"`
	MptcpDebug               string `json:"mptcpDebug"`
	MptcpFallback            string `json:"mptcpFallback"`
	MptcpFastjoin            string `json:"mptcpFastjoin"`
	MptcpIdleTimeout         int    `json:"mptcpIdleTimeout"`
	MptcpJoinMax             int    `json:"mptcpJoinMax"`
	MptcpMakeafterbreak      string `json:"mptcpMakeafterbreak"`
	MptcpNojoindssack        string `json:"mptcpNojoindssack"`
	MptcpRtomax              int    `json:"mptcpRtomax"`
	MptcpRxmitmin            int    `json:"mptcpRxmitmin"`
	MptcpSubflowmax          int    `json:"mptcpSubflowmax"`
	MptcpTimeout             int    `json:"mptcpTimeout"`
	Nagle                    string `json:"nagle"`
	PktLossIgnoreBurst       int    `json:"pktLossIgnoreBurst"`
	PktLossIgnoreRate        int    `json:"pktLossIgnoreRate"`
	ProxyBufferHigh          int    `json:"proxyBufferHigh"`
	ProxyBufferLow           int    `json:"proxyBufferLow"`
	ProxyMss                 string `json:"proxyMss"`
	ProxyOptions             string `json:"proxyOptions"`
	PushFlag                 string `json:"pushFlag"`
	RatePace                 string `json:"ratePace"`
	RatePaceMaxRate          int    `json:"ratePaceMaxRate"`
	ReceiveWindowSize        int    `json:"receiveWindowSize"`
	ResetOnTimeout           string `json:"resetOnTimeout"`
	RexmtThresh              int    `json:"rexmtThresh"`
	SelectiveAcks            string `json:"selectiveAcks"`
	SelectiveNack            string `json:"selectiveNack"`
	SendBufferSize           int    `json:"sendBufferSize"`
	SlowStart                string `json:"slowStart"`
	SynCookieEnable          string `json:"synCookieEnable"`
	SynCookieWhitelist       string `json:"synCookieWhitelist"`
	SynMaxRetrans            int    `json:"synMaxRetrans"`
	SynRtoBase               int    `json:"synRtoBase"`
	TailLossProbe            string `json:"tailLossProbe"`
	TCPOptions               string `json:"tcpOptions"`
	TimeWaitRecycle          string `json:"timeWaitRecycle"`
	TimeWaitTimeout          string `json:"timeWaitTimeout"`
	Timestamps               string `json:"timestamps"`
	VerifiedAccept           string `json:"verifiedAccept"`
	ZeroWindowTimeout        int    `json:"zeroWindowTimeout"`
}

const TCPEndpoint = "tcp"

type TCPResource struct {
	b *bigip.BigIP
}
