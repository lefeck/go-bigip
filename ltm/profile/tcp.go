package profile

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

type TCPList struct {
	Items    []TCP  `json:"items,omitempty"`
	Kind     string `json:"kind,omitempty"`
	SelfLink string `json:"selflink,omitempty"`
}

type TCP struct {
	Kind                  string `json:"kind,omitempty"`
	Name                  string `json:"name,omitempty"`
	Partition             string `json:"partition,omitempty"`
	FullPath              string `json:"fullPath,omitempty"`
	Generation            int    `json:"generation,omitempty"`
	SelfLink              string `json:"selfLink,omitempty"`
	Abc                   string `json:"abc,omitempty"`
	AckOnPush             string `json:"ackOnPush,omitempty"`
	AppService            string `json:"appService,omitempty"`
	AutoProxyBufferSize   string `json:"autoProxyBufferSize,omitempty"`
	AutoReceiveWindowSize string `json:"autoReceiveWindowSize,omitempty"`
	AutoSendBufferSize    string `json:"autoSendBufferSize,omitempty"`
	CloseWaitTimeout      int    `json:"closeWaitTimeout,omitempty"`
	CmetricsCache         string `json:"cmetricsCache,omitempty"`
	CmetricsCacheTimeout  int    `json:"cmetricsCacheTimeout,omitempty"`
	CongestionControl     string `json:"congestionControl,omitempty"`
	DefaultsFrom          string `json:"defaultsFrom,omitempty"`
	DefaultsFromReference struct {
		Link string `json:"link,omitempty"`
	} `json:"defaultsFromReference,omitempty"`
	DeferredAccept           string `json:"deferredAccept,omitempty"`
	DelayWindowControl       string `json:"delayWindowControl,omitempty"`
	DelayedAcks              string `json:"delayedAcks,omitempty"`
	Description              string `json:"description,omitempty"`
	Dsack                    string `json:"dsack,omitempty"`
	EarlyRetransmit          string `json:"earlyRetransmit,omitempty"`
	Ecn                      string `json:"ecn,omitempty"`
	EnhancedLossRecovery     string `json:"enhancedLossRecovery,omitempty"`
	FastOpen                 string `json:"fastOpen,omitempty"`
	FastOpenCookieExpiration int    `json:"fastOpenCookieExpiration,omitempty"`
	FinWait2Timeout          int    `json:"finWait_2Timeout,omitempty"`
	FinWaitTimeout           int    `json:"finWaitTimeout,omitempty"`
	HardwareSynCookie        string `json:"hardwareSynCookie,omitempty"`
	IdleTimeout              int    `json:"idleTimeout,omitempty"`
	InitCwnd                 int    `json:"initCwnd,omitempty"`
	InitRwnd                 int    `json:"initRwnd,omitempty"`
	IPDfMode                 string `json:"ipDfMode,omitempty"`
	IPTosToClient            string `json:"ipTosToClient,omitempty"`
	IPTTLMode                string `json:"ipTtlMode,omitempty"`
	IPTTLV4                  int    `json:"ipTtlV4,omitempty"`
	IPTTLV6                  int    `json:"ipTtlV6,omitempty"`
	KeepAliveInterval        int    `json:"keepAliveInterval,omitempty"`
	LimitedTransmit          string `json:"limitedTransmit,omitempty"`
	LinkQosToClient          string `json:"linkQosToClient,omitempty"`
	MaxRetrans               int    `json:"maxRetrans,omitempty"`
	MaxSegmentSize           int    `json:"maxSegmentSize,omitempty"`
	Md5Signature             string `json:"md5Signature,omitempty"`
	MinimumRto               int    `json:"minimumRto,omitempty"`
	Mptcp                    string `json:"mptcp,omitempty"`
	MptcpCsum                string `json:"mptcpCsum,omitempty"`
	MptcpCsumVerify          string `json:"mptcpCsumVerify,omitempty"`
	MptcpDebug               string `json:"mptcpDebug,omitempty"`
	MptcpFallback            string `json:"mptcpFallback,omitempty"`
	MptcpFastjoin            string `json:"mptcpFastjoin,omitempty"`
	MptcpIdleTimeout         int    `json:"mptcpIdleTimeout,omitempty"`
	MptcpJoinMax             int    `json:"mptcpJoinMax,omitempty"`
	MptcpMakeafterbreak      string `json:"mptcpMakeafterbreak,omitempty"`
	MptcpNojoindssack        string `json:"mptcpNojoindssack,omitempty"`
	MptcpRtomax              int    `json:"mptcpRtomax,omitempty"`
	MptcpRxmitmin            int    `json:"mptcpRxmitmin,omitempty"`
	MptcpSubflowmax          int    `json:"mptcpSubflowmax,omitempty"`
	MptcpTimeout             int    `json:"mptcpTimeout,omitempty"`
	Nagle                    string `json:"nagle,omitempty"`
	PktLossIgnoreBurst       int    `json:"pktLossIgnoreBurst,omitempty"`
	PktLossIgnoreRate        int    `json:"pktLossIgnoreRate,omitempty"`
	ProxyBufferHigh          int    `json:"proxyBufferHigh,omitempty"`
	ProxyBufferLow           int    `json:"proxyBufferLow,omitempty"`
	ProxyMss                 string `json:"proxyMss,omitempty"`
	ProxyOptions             string `json:"proxyOptions,omitempty"`
	PushFlag                 string `json:"pushFlag,omitempty"`
	RatePace                 string `json:"ratePace,omitempty"`
	RatePaceMaxRate          int    `json:"ratePaceMaxRate,omitempty"`
	ReceiveWindowSize        int    `json:"receiveWindowSize,omitempty"`
	ResetOnTimeout           string `json:"resetOnTimeout,omitempty"`
	RexmtThresh              int    `json:"rexmtThresh,omitempty"`
	SelectiveAcks            string `json:"selectiveAcks,omitempty"`
	SelectiveNack            string `json:"selectiveNack,omitempty"`
	SendBufferSize           int    `json:"sendBufferSize,omitempty"`
	SlowStart                string `json:"slowStart,omitempty"`
	SynCookieEnable          string `json:"synCookieEnable,omitempty"`
	SynCookieWhitelist       string `json:"synCookieWhitelist,omitempty"`
	SynMaxRetrans            int    `json:"synMaxRetrans,omitempty"`
	SynRtoBase               int    `json:"synRtoBase,omitempty"`
	TailLossProbe            string `json:"tailLossProbe,omitempty"`
	TCPOptions               string `json:"tcpOptions,omitempty"`
	TimeWaitRecycle          string `json:"timeWaitRecycle,omitempty"`
	TimeWaitTimeout          string `json:"timeWaitTimeout,omitempty"`
	Timestamps               string `json:"timestamps,omitempty"`
	VerifiedAccept           string `json:"verifiedAccept,omitempty"`
	ZeroWindowTimeout        int    `json:"zeroWindowTimeout,omitempty"`
}

const TCPEndpoint = "tcp"

type TCPResource struct {
	b *bigip.BigIP
}

// List retrieves a list of TCP resources.
func (cr *TCPResource) List() (*TCPList, error) {
	var items TCPList
	// Perform a GET request to retrieve a list of TCP resource objects
	res, err := cr.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(TCPEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	// Unmarshal JSON response data into TCPList struct
	if err := json.Unmarshal(res, &items); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &items, nil
}

// Get retrieves a TCP resource by its full path name.
func (cr *TCPResource) Get(fullPathName string) (*TCP, error) {
	var item TCP
	// Perform a GET request to retrieve a specific TCP resource by its full path name
	res, err := cr.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(TCPEndpoint).SubResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	// Unmarshal JSON response data into TCP struct
	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &item, nil
}

// Create adds a new TCP resource using the provided TCP item.
func (cr *TCPResource) Create(item TCP) error {
	// Marshal the TCP struct into JSON data
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)

	// Perform a POST request to create a new TCP resource using the JSON data
	_, err = cr.b.RestClient.Post().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(TCPEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Update modifies a TCP resource identified by its full path name using the provided TCP item.
func (cr *TCPResource) Update(fullPathName string, item TCP) error {
	// Marshal the TCP struct into JSON data
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)

	// Perform a PUT request to update the specified TCP resource with the JSON data
	_, err = cr.b.RestClient.Put().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(TCPEndpoint).SubResourceInstance(fullPathName).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Delete removes a TCP resource by its full path name.
func (cr *TCPResource) Delete(fullPathName string) error {
	// Perform a DELETE request to delete the specified TCP resource
	_, err := cr.b.RestClient.Delete().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(TCPEndpoint).SubResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
