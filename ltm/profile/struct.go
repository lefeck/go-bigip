package profile

//
//
//
//
//type connector struct {
//	Kind               string `json:"kind"`
//	Name               string `json:"name"`
//	Partition          string `json:"partition"`
//	FullPath           string `json:"fullPath"`
//	Generation         int    `json:"generation"`
//	SelfLink           string `json:"selfLink"`
//	AppService         string `json:"appService"`
//	ConnectOnData      string `json:"connectOnData"`
//	ConnectionTimeout  int    `json:"connectionTimeout"`
//	EntryVirtualServer string `json:"entryVirtualServer"`
//	ServiceDownAction  string `json:"serviceDownAction"`
//}
//
//
//type fasthttp struct {
//	Kind                        string `json:"kind"`
//	Name                        string `json:"name"`
//	Partition                   string `json:"partition"`
//	FullPath                    string `json:"fullPath"`
//	Generation                  int    `json:"generation"`
//	SelfLink                    string `json:"selfLink"`
//	AppService                  string `json:"appService"`
//	ClientCloseTimeout          int    `json:"clientCloseTimeout"`
//	ConnpoolIdleTimeoutOverride int    `json:"connpoolIdleTimeoutOverride"`
//	ConnpoolMaxReuse            int    `json:"connpoolMaxReuse"`
//	ConnpoolMaxSize             int    `json:"connpoolMaxSize"`
//	ConnpoolMinSize             int    `json:"connpoolMinSize"`
//	ConnpoolReplenish           string `json:"connpoolReplenish"`
//	ConnpoolStep                int    `json:"connpoolStep"`
//	DefaultsFrom                string `json:"defaultsFrom"`
//	Description                 string `json:"description"`
//	ForceHTTP10Response         string `json:"forceHttp_10Response"`
//	HardwareSynCookie           string `json:"hardwareSynCookie"`
//	HeaderInsert                string `json:"headerInsert"`
//	HTTP11CloseWorkarounds      string `json:"http_11CloseWorkarounds"`
//	IdleTimeout                 int    `json:"idleTimeout"`
//	InsertXforwardedFor         string `json:"insertXforwardedFor"`
//	Layer7                      string `json:"layer_7"`
//	MaxHeaderSize               int    `json:"maxHeaderSize"`
//	MaxRequests                 int    `json:"maxRequests"`
//	MssOverride                 int    `json:"mssOverride"`
//	ReceiveWindowSize           int    `json:"receiveWindowSize"`
//	ResetOnTimeout              string `json:"resetOnTimeout"`
//	ServerCloseTimeout          int    `json:"serverCloseTimeout"`
//	ServerSack                  string `json:"serverSack"`
//	ServerTimestamp             string `json:"serverTimestamp"`
//	UncleanShutdown             string `json:"uncleanShutdown"`
//}
//
//
//type fastl4 struct {
//	Kind                  string `json:"kind"`
//	Name                  string `json:"name"`
//	Partition             string `json:"partition"`
//	FullPath              string `json:"fullPath"`
//	Generation            int    `json:"generation"`
//	SelfLink              string `json:"selfLink"`
//	AppService            string `json:"appService"`
//	ClientTimeout         int    `json:"clientTimeout"`
//	DefaultsFrom          string `json:"defaultsFrom"`
//	DefaultsFromReference struct {
//		Link string `json:"link"`
//	} `json:"defaultsFromReference"`
//	Description                 string `json:"description"`
//	ExplicitFlowMigration       string `json:"explicitFlowMigration"`
//	HardwareSynCookie           string `json:"hardwareSynCookie"`
//	IdleTimeout                 string `json:"idleTimeout"`
//	IPDfMode                    string `json:"ipDfMode"`
//	IPTosToClient               string `json:"ipTosToClient"`
//	IPTosToServer               string `json:"ipTosToServer"`
//	IPTTLMode                   string `json:"ipTtlMode"`
//	IPTTLV4                     int    `json:"ipTtlV4"`
//	IPTTLV6                     int    `json:"ipTtlV6"`
//	KeepAliveInterval           string `json:"keepAliveInterval"`
//	LateBinding                 string `json:"lateBinding"`
//	LinkQosToClient             string `json:"linkQosToClient"`
//	LinkQosToServer             string `json:"linkQosToServer"`
//	LooseClose                  string `json:"looseClose"`
//	LooseInitialization         string `json:"looseInitialization"`
//	MssOverride                 int    `json:"mssOverride"`
//	OtherPvaClientpktsThreshold int    `json:"otherPvaClientpktsThreshold"`
//	OtherPvaOffloadDirection    string `json:"otherPvaOffloadDirection"`
//	OtherPvaServerpktsThreshold int    `json:"otherPvaServerpktsThreshold"`
//	OtherPvaWhentoOffload       string `json:"otherPvaWhentoOffload"`
//	PriorityToClient            string `json:"priorityToClient"`
//	PriorityToServer            string `json:"priorityToServer"`
//	PvaAcceleration             string `json:"pvaAcceleration"`
//	PvaDynamicClientPackets     int    `json:"pvaDynamicClientPackets"`
//	PvaDynamicServerPackets     int    `json:"pvaDynamicServerPackets"`
//	PvaFlowAging                string `json:"pvaFlowAging"`
//	PvaFlowEvict                string `json:"pvaFlowEvict"`
//	PvaOffloadDynamic           string `json:"pvaOffloadDynamic"`
//	PvaOffloadDynamicPriority   string `json:"pvaOffloadDynamicPriority"`
//	PvaOffloadInitialPriority   string `json:"pvaOffloadInitialPriority"`
//	PvaOffloadState             string `json:"pvaOffloadState"`
//	ReassembleFragments         string `json:"reassembleFragments"`
//	ReceiveWindowSize           int    `json:"receiveWindowSize"`
//	ResetOnTimeout              string `json:"resetOnTimeout"`
//	RttFromClient               string `json:"rttFromClient"`
//	RttFromServer               string `json:"rttFromServer"`
//	ServerSack                  string `json:"serverSack"`
//	ServerTimestamp             string `json:"serverTimestamp"`
//	SoftwareSynCookie           string `json:"softwareSynCookie"`
//	SynCookieDsrFlowResetBy     string `json:"synCookieDsrFlowResetBy"`
//	SynCookieEnable             string `json:"synCookieEnable"`
//	SynCookieMss                int    `json:"synCookieMss"`
//	SynCookieWhitelist          string `json:"synCookieWhitelist"`
//	TCPCloseTimeout             string `json:"tcpCloseTimeout"`
//	TCPGenerateIsn              string `json:"tcpGenerateIsn"`
//	TCPHandshakeTimeout         string `json:"tcpHandshakeTimeout"`
//	TCPPvaOffloadDirection      string `json:"tcpPvaOffloadDirection"`
//	TCPPvaWhentoOffload         string `json:"tcpPvaWhentoOffload"`
//	TCPStripSack                string `json:"tcpStripSack"`
//	TCPTimeWaitTimeout          int    `json:"tcpTimeWaitTimeout"`
//	TCPTimestampMode            string `json:"tcpTimestampMode"`
//	TCPWscaleMode               string `json:"tcpWscaleMode"`
//	TimeoutRecovery             string `json:"timeoutRecovery"`
//}
//
//
//type fix struct {
//	Kind                     string        `json:"kind"`
//	Name                     string        `json:"name"`
//	Partition                string        `json:"partition"`
//	FullPath                 string        `json:"fullPath"`
//	Generation               int           `json:"generation"`
//	SelfLink                 string        `json:"selfLink"`
//	AppService               string        `json:"appService"`
//	DefaultsFrom             string        `json:"defaultsFrom"`
//	Description              string        `json:"description"`
//	ErrorAction              string        `json:"errorAction"`
//	FullLogonParsing         string        `json:"fullLogonParsing"`
//	MessageLogPublisher      string        `json:"messageLogPublisher"`
//	QuickParsing             string        `json:"quickParsing"`
//	ReportLogPublisher       string        `json:"reportLogPublisher"`
//	ResponseParsing          string        `json:"responseParsing"`
//	SenderTagClass           []interface{} `json:"senderTagClass"`
//	StatisticsSampleInterval int           `json:"statisticsSampleInterval"`
//}
//
//
//type ftp struct {
//	Kind                   string `json:"kind"`
//	Name                   string `json:"name"`
//	Partition              string `json:"partition"`
//	FullPath               string `json:"fullPath"`
//	Generation             int    `json:"generation"`
//	SelfLink               string `json:"selfLink"`
//	AllowActiveMode        string `json:"allowActiveMode"`
//	AllowFtps              string `json:"allowFtps"`
//	AppService             string `json:"appService"`
//	DefaultsFrom           string `json:"defaultsFrom"`
//	Description            string `json:"description"`
//	EnforceTLSSessionReuse string `json:"enforceTlsSessionReuse"`
//	FtpsMode               string `json:"ftpsMode"`
//	InheritParentProfile   string `json:"inheritParentProfile"`
//	InheritVlanList        string `json:"inheritVlanList"`
//	LogProfile             string `json:"logProfile"`
//	LogPublisher           string `json:"logPublisher"`
//	Port                   int    `json:"port"`
//	Security               string `json:"security"`
//	TranslateExtended      string `json:"translateExtended"`
//}
//
//
//type html struct {
//		Kind             string   `json:"kind"`
//		Name             string   `json:"name"`
//		Partition        string   `json:"partition"`
//		FullPath         string   `json:"fullPath"`
//		Generation       int      `json:"generation"`
//		SelfLink         string   `json:"selfLink"`
//		AppService       string   `json:"appService"`
//		ContentDetection string   `json:"contentDetection"`
//		ContentSelection []string `json:"contentSelection"`
//		DefaultsFrom     string   `json:"defaultsFrom"`
//		Description      string   `json:"description"`
//	}
//
//
//
//	type http-compression struct {
//		Kind                  string        `json:"kind"`
//		Name                  string        `json:"name"`
//		Partition             string        `json:"partition"`
//		FullPath              string        `json:"fullPath"`
//		Generation            int           `json:"generation"`
//		SelfLink              string        `json:"selfLink"`
//		AllowHTTP10           string        `json:"allowHttp_10"`
//		AppService            string        `json:"appService"`
//		BrowserWorkarounds    string        `json:"browserWorkarounds"`
//		BufferSize            int           `json:"bufferSize"`
//		ContentTypeExclude    []interface{} `json:"contentTypeExclude"`
//		ContentTypeInclude    []string      `json:"contentTypeInclude"`
//		CPUSaver              string        `json:"cpuSaver"`
//		CPUSaverHigh          int           `json:"cpuSaverHigh"`
//		CPUSaverLow           int           `json:"cpuSaverLow"`
//		DefaultsFrom          string        `json:"defaultsFrom"`
//		DefaultsFromReference struct {
//			Link string `json:"link"`
//		} `json:"defaultsFromReference"`
//		Description        string        `json:"description"`
//		GzipLevel          int           `json:"gzipLevel"`
//		GzipMemoryLevel    int           `json:"gzipMemoryLevel"`
//		GzipWindowSize     int           `json:"gzipWindowSize"`
//		KeepAcceptEncoding string        `json:"keepAcceptEncoding"`
//		MethodPrefer       string        `json:"methodPrefer"`
//		MinSize            int           `json:"minSize"`
//		Selective          string        `json:"selective"`
//		URIExclude         []interface{} `json:"uriExclude"`
//		URIInclude         []string      `json:"uriInclude"`
//		VaryHeader         string        `json:"varyHeader"`
//	}
//
//
//
//	type http-proxy-connect struct {
//		Kind         string `json:"kind"`
//		Name         string `json:"name"`
//		Partition    string `json:"partition"`
//		FullPath     string `json:"fullPath"`
//		Generation   int    `json:"generation"`
//		SelfLink     string `json:"selfLink"`
//		AppService   string `json:"appService"`
//		DefaultState string `json:"defaultState"`
//		DefaultsFrom string `json:"defaultsFrom"`
//		Description  string `json:"description"`
//	}
//
//	type http2 struct {
//		Kind                           string   `json:"kind"`
//		Name                           string   `json:"name"`
//		Partition                      string   `json:"partition"`
//		FullPath                       string   `json:"fullPath"`
//		Generation                     int      `json:"generation"`
//		SelfLink                       string   `json:"selfLink"`
//		ActivationModes                []string `json:"activationModes"`
//		AppService                     string   `json:"appService"`
//		ConcurrentStreamsPerConnection int      `json:"concurrentStreamsPerConnection"`
//		ConnectionIdleTimeout          int      `json:"connectionIdleTimeout"`
//		DefaultsFrom                   string   `json:"defaultsFrom"`
//		Description                    string   `json:"description"`
//		EnforceTLSRequirements         string   `json:"enforceTlsRequirements"`
//		FrameSize                      int      `json:"frameSize"`
//		HeaderTableSize                int      `json:"headerTableSize"`
//		IncludeContentLength           string   `json:"includeContentLength"`
//		InsertHeader                   string   `json:"insertHeader"`
//		InsertHeaderName               string   `json:"insertHeaderName"`
//		ReceiveWindow                  int      `json:"receiveWindow"`
//		WriteSize                      int      `json:"writeSize"`
//	}
//
//
//
//type http3 struct {
//		Kind            string `json:"kind"`
//		Name            string `json:"name"`
//		Partition       string `json:"partition"`
//		FullPath        string `json:"fullPath"`
//		Generation      int    `json:"generation"`
//		SelfLink        string `json:"selfLink"`
//		AppService      string `json:"appService"`
//		DefaultsFrom    string `json:"defaultsFrom"`
//		Description     string `json:"description"`
//		HeaderTableSize int    `json:"headerTableSize"`
//	}
//
//
//	type httprouter struct {
//		Kind     string `json:"kind"`
//		SelfLink string `json:"selfLink"`
//		Items    []struct {
//			Kind         string `json:"kind"`
//			Name         string `json:"name"`
//			Partition    string `json:"partition"`
//			FullPath     string `json:"fullPath"`
//			Generation   int    `json:"generation"`
//			SelfLink     string `json:"selfLink"`
//			AppService   string `json:"appService"`
//			DefaultsFrom string `json:"defaultsFrom"`
//			Description  string `json:"description"`
//		} `json:"items"`
//	}
//
//
//
//	type icap struct {
//		Kind          string `json:"kind"`
//		Name          string `json:"name"`
//		Partition     string `json:"partition"`
//		FullPath      string `json:"fullPath"`
//		Generation    int    `json:"generation"`
//		SelfLink      string `json:"selfLink"`
//		AppService    string `json:"appService"`
//		DefaultsFrom  string `json:"defaultsFrom"`
//		HeaderFrom    string `json:"headerFrom"`
//		Host          string `json:"host"`
//		PreviewLength int    `json:"previewLength"`
//		Referer       string `json:"referer"`
//		URI           string `json:"uri"`
//		UserAgent     string `json:"userAgent"`
//	}
//
//
//
//	type imap struct {
//		Kind           string `json:"kind"`
//		Name           string `json:"name"`
//		Partition      string `json:"partition"`
//		FullPath       string `json:"fullPath"`
//		Generation     int    `json:"generation"`
//		SelfLink       string `json:"selfLink"`
//		ActivationMode string `json:"activationMode"`
//		AppService     string `json:"appService"`
//		DefaultsFrom   string `json:"defaultsFrom"`
//		Description    string `json:"description"`
//	}
//
//
//	type netflow struct {
//		Kind           string `json:"kind"`
//		Name           string `json:"name"`
//		Partition      string `json:"partition"`
//		FullPath       string `json:"fullPath"`
//		Generation     int    `json:"generation"`
//		SelfLink       string `json:"selfLink"`
//		AppService     string `json:"appService"`
//		DefaultsFrom   string `json:"defaultsFrom"`
//		Description    string `json:"description"`
//		NetflowVersion string `json:"netflowVersion"`
//		SamplingRate   int    `json:"samplingRate"`
//	}
//
//
//	type ntlm struct {
//		Kind                   string `json:"kind"`
//		Name                   string `json:"name"`
//		Partition              string `json:"partition"`
//		FullPath               string `json:"fullPath"`
//		Generation             int    `json:"generation"`
//		SelfLink               string `json:"selfLink"`
//		AppService             string `json:"appService"`
//		DefaultsFrom           string `json:"defaultsFrom"`
//		Description            string `json:"description"`
//		InsertCookieDomain     string `json:"insertCookieDomain"`
//		InsertCookieName       string `json:"insertCookieName"`
//		InsertCookiePassphrase string `json:"insertCookiePassphrase"`
//		KeyByCookie            string `json:"keyByCookie"`
//		KeyByCookieName        string `json:"keyByCookieName"`
//		KeyByDomain            string `json:"keyByDomain"`
//		KeyByIPAddress         string `json:"keyByIpAddress"`
//		KeyByTarget            string `json:"keyByTarget"`
//		KeyByUser              string `json:"keyByUser"`
//		KeyByWorkstation       string `json:"keyByWorkstation"`
//	}
//
//
//	type ocsp struct {
//		Kind         string `json:"kind"`
//		Name         string `json:"name"`
//		Partition    string `json:"partition"`
//		FullPath     string `json:"fullPath"`
//		Generation   int    `json:"generation"`
//		SelfLink     string `json:"selfLink"`
//		AppService   string `json:"appService"`
//		DefaultsFrom string `json:"defaultsFrom"`
//		MaxAge       int    `json:"maxAge"`
//		Nonce        string `json:"nonce"`
//	}
//
//
//	type one-connect struct {
//		Kind                string `json:"kind"`
//		Name                string `json:"name"`
//		Partition           string `json:"partition"`
//		FullPath            string `json:"fullPath"`
//		Generation          int    `json:"generation"`
//		SelfLink            string `json:"selfLink"`
//		AppService          string `json:"appService"`
//		DefaultsFrom        string `json:"defaultsFrom"`
//		Description         string `json:"description"`
//		IdleTimeoutOverride string `json:"idleTimeoutOverride"`
//		LimitType           string `json:"limitType"`
//		MaxAge              int    `json:"maxAge"`
//		MaxReuse            int    `json:"maxReuse"`
//		MaxSize             int    `json:"maxSize"`
//		SharePools          string `json:"sharePools"`
//		SourceMask          string `json:"sourceMask"`
//	}
//
//
//	type pop3 struct {
//		Kind           string `json:"kind"`
//		Name           string `json:"name"`
//		Partition      string `json:"partition"`
//		FullPath       string `json:"fullPath"`
//		Generation     int    `json:"generation"`
//		SelfLink       string `json:"selfLink"`
//		ActivationMode string `json:"activationMode"`
//		AppService     string `json:"appService"`
//		DefaultsFrom   string `json:"defaultsFrom"`
//		Description    string `json:"description"`
//	}
//
//
//	type rewrite struct {
//		Kind                  string        `json:"kind"`
//		Name                  string        `json:"name"`
//		Partition             string        `json:"partition"`
//		FullPath              string        `json:"fullPath"`
//		Generation            int           `json:"generation"`
//		SelfLink              string        `json:"selfLink"`
//		AppService            string        `json:"appService"`
//		BypassList            []interface{} `json:"bypassList"`
//		ClientCachingType     string        `json:"clientCachingType"`
//		DefaultsFrom          string        `json:"defaultsFrom"`
//		DefaultsFromReference struct {
//			Link string `json:"link"`
//		} `json:"defaultsFromReference"`
//		JavaCaFile          string `json:"javaCaFile"`
//		JavaCaFileReference struct {
//			Link string `json:"link"`
//		} `json:"javaCaFileReference"`
//		JavaCrl              string `json:"javaCrl"`
//		JavaSignKey          string `json:"javaSignKey"`
//		JavaSignKeyReference struct {
//			Link string `json:"link"`
//		} `json:"javaSignKeyReference"`
//		JavaSigner          string `json:"javaSigner"`
//		JavaSignerReference struct {
//			Link string `json:"link"`
//		} `json:"javaSignerReference"`
//		LocationSpecific string `json:"locationSpecific"`
//		Request          struct {
//			InsertXforwardedFor   string `json:"insertXforwardedFor"`
//			InsertXforwardedHost  string `json:"insertXforwardedHost"`
//			InsertXforwardedProto string `json:"insertXforwardedProto"`
//			RewriteHeaders        string `json:"rewriteHeaders"`
//		} `json:"request"`
//		Response struct {
//			RewriteContent string `json:"rewriteContent"`
//			RewriteHeaders string `json:"rewriteHeaders"`
//		} `json:"response"`
//		RewriteList       []interface{} `json:"rewriteList"`
//		RewriteMode       string        `json:"rewriteMode"`
//		SplitTunneling    string        `json:"splitTunneling"`
//		URIRulesReference struct {
//			Link            string `json:"link"`
//			IsSubcollection bool   `json:"isSubcollection"`
//		} `json:"uriRulesReference"`
//	}
//
//
//	type server-ssl struct {
//		Kind                       string        `json:"kind"`
//		Name                       string        `json:"name"`
//		Partition                  string        `json:"partition"`
//		FullPath                   string        `json:"fullPath"`
//		Generation                 int           `json:"generation"`
//		SelfLink                   string        `json:"selfLink"`
//		AlertTimeout               string        `json:"alertTimeout"`
//		AllowExpiredCrl            string        `json:"allowExpiredCrl"`
//		AppService                 string        `json:"appService"`
//		Authenticate               string        `json:"authenticate"`
//		AuthenticateDepth          int           `json:"authenticateDepth"`
//		AuthenticateName           string        `json:"authenticateName"`
//		BypassOnClientCertFail     string        `json:"bypassOnClientCertFail"`
//		BypassOnHandshakeAlert     string        `json:"bypassOnHandshakeAlert"`
//		C3DCaCert                  string        `json:"c3dCaCert"`
//		C3DCaKey                   string        `json:"c3dCaKey"`
//		C3DCertExtensionCustomOids []interface{} `json:"c3dCertExtensionCustomOids"`
//		C3DCertExtensionIncludes   []string      `json:"c3dCertExtensionIncludes"`
//		C3DCertLifespan            int           `json:"c3dCertLifespan"`
//		CaFile                     string        `json:"caFile"`
//		CacheSize                  int           `json:"cacheSize"`
//		CacheTimeout               int           `json:"cacheTimeout"`
//		Cert                       string        `json:"cert"`
//		CertReference              struct {
//			Link string `json:"link"`
//		} `json:"certReference"`
//		Chain                 string `json:"chain"`
//		CipherGroup           string `json:"cipherGroup"`
//		Ciphers               string `json:"ciphers"`
//		Crl                   string `json:"crl"`
//		CrlFile               string `json:"crlFile"`
//		Data0Rtt              string `json:"data_0rtt"`
//		DefaultsFrom          string `json:"defaultsFrom"`
//		DefaultsFromReference struct {
//			Link string `json:"link"`
//		} `json:"defaultsFromReference"`
//		Description               string `json:"description"`
//		ExpireCertResponseControl string `json:"expireCertResponseControl"`
//		GenericAlert              string `json:"genericAlert"`
//		HandshakeTimeout          string `json:"handshakeTimeout"`
//		Key                       string `json:"key"`
//		KeyReference              struct {
//			Link string `json:"link"`
//		} `json:"keyReference"`
//		LogPublisher          string `json:"logPublisher"`
//		LogPublisherReference struct {
//			Link string `json:"link"`
//		} `json:"logPublisherReference"`
//		MaxActiveHandshakes              string `json:"maxActiveHandshakes"`
//		ModSslMethods                    string `json:"modSslMethods"`
//		Mode                             string `json:"mode"`
//		Ocsp                             string `json:"ocsp"`
//		TmOptions                        string `json:"tmOptions"`
//		PeerCertMode                     string `json:"peerCertMode"`
//		ProxySsl                         string `json:"proxySsl"`
//		ProxySslPassthrough              string `json:"proxySslPassthrough"`
//		RenegotiatePeriod                string `json:"renegotiatePeriod"`
//		RenegotiateSize                  string `json:"renegotiateSize"`
//		Renegotiation                    string `json:"renegotiation"`
//		RetainCertificate                string `json:"retainCertificate"`
//		RevokedCertStatusResponseControl string `json:"revokedCertStatusResponseControl"`
//		SecureRenegotiation              string `json:"secureRenegotiation"`
//		ServerName                       string `json:"serverName"`
//		SessionMirroring                 string `json:"sessionMirroring"`
//		SessionTicket                    string `json:"sessionTicket"`
//		SniDefault                       string `json:"sniDefault"`
//		SniRequire                       string `json:"sniRequire"`
//		SslC3D                           string `json:"sslC3d"`
//		SslForwardProxy                  string `json:"sslForwardProxy"`
//		SslForwardProxyBypass            string `json:"sslForwardProxyBypass"`
//		SslForwardProxyVerifiedHandshake string `json:"sslForwardProxyVerifiedHandshake"`
//		SslSignHash                      string `json:"sslSignHash"`
//		StrictResume                     string `json:"strictResume"`
//		UncleanShutdown                  string `json:"uncleanShutdown"`
//		UnknownCertStatusResponseControl string `json:"unknownCertStatusResponseControl"`
//		UntrustedCertResponseControl     string `json:"untrustedCertResponseControl"`
//	}
//
//
//	type service struct {
//		Kind       string `json:"kind"`
//		Name       string `json:"name"`
//		Partition  string `json:"partition"`
//		FullPath   string `json:"fullPath"`
//		Generation int    `json:"generation"`
//		SelfLink   string `json:"selfLink"`
//		AppService string `json:"appService"`
//		Type       string `json:"type"`
//	}
//
//

//

//
//	type stream struct {
//		Kind         string `json:"kind"`
//		Name         string `json:"name"`
//		Partition    string `json:"partition"`
//		FullPath     string `json:"fullPath"`
//		Generation   int    `json:"generation"`
//		SelfLink     string `json:"selfLink"`
//		AppService   string `json:"appService"`
//		ChunkSize    int    `json:"chunkSize"`
//		Chunking     string `json:"chunking"`
//		DefaultsFrom string `json:"defaultsFrom"`
//		Description  string `json:"description"`
//		Source       string `json:"source"`
//		TmTarget     string `json:"tmTarget"`
//	}
//
//
//	type tftp struct {
//		Kind         string `json:"kind"`
//		Name         string `json:"name"`
//		Partition    string `json:"partition"`
//		FullPath     string `json:"fullPath"`
//		Generation   int    `json:"generation"`
//		SelfLink     string `json:"selfLink"`
//		AppService   string `json:"appService"`
//		DefaultsFrom string `json:"defaultsFrom"`
//		Description  string `json:"description"`
//		IdleTimeout  string `json:"idleTimeout"`
//		LogProfile   string `json:"logProfile"`
//		LogPublisher string `json:"logPublisher"`
//	}
//
//
//	type websocket struct {
//		Kind                   string `json:"kind"`
//		Name                   string `json:"name"`
//		Partition              string `json:"partition"`
//		FullPath               string `json:"fullPath"`
//		Generation             int    `json:"generation"`
//		SelfLink               string `json:"selfLink"`
//		AppService             string `json:"appService"`
//		CompressMode           string `json:"compressMode"`
//		Compression            string `json:"compression"`
//		DefaultsFrom           string `json:"defaultsFrom"`
//		Description            string `json:"description"`
//		Masking                string `json:"masking"`
//		NoDelay                string `json:"noDelay"`
//		PayloadProcessingMode  string `json:"payloadProcessingMode"`
//		PayloadProtocolProfile string `json:"payloadProtocolProfile"`
//		WindowBits             int    `json:"windowBits"`
//	}
//
//
//	type xml struct {
//		Kind                 string        `json:"kind"`
//		Name                 string        `json:"name"`
//		Partition            string        `json:"partition"`
//		FullPath             string        `json:"fullPath"`
//		Generation           int           `json:"generation"`
//		SelfLink             string        `json:"selfLink"`
//		AppService           string        `json:"appService"`
//		DefaultsFrom         string        `json:"defaultsFrom"`
//		Description          string        `json:"description"`
//		MultipleQueryMatches string        `json:"multipleQueryMatches"`
//		NamespaceMappings    []interface{} `json:"namespaceMappings"`
//		XpathQueries         []interface{} `json:"xpathQueries"`
//	}
