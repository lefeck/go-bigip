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
