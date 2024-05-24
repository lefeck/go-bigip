package crypto

import "github.com/lefeck/go-bigip"

// CryptoConfigList holds a list of Crypto configuration.

// CryptoEndpoint represents the REST resource for managing Crypto.
const CryptoEndpoint = "crypto"
const SysManager = "sys"

// CryptoResource provides an API to manage Crypto configurations.
type CryptoResource struct {
	cert   CertResource
	client ClientResource
	crl    CrlResource
	csr    CsrResource
	key    KeyResource
	server ServerResource
}

func NewCrypto(b *bigip.BigIP) CryptoResource {
	return CryptoResource{
		cert:   CertResource{b: b},
		client: ClientResource{b: b},
		crl:    CrlResource{b: b},
		csr:    CsrResource{b: b},
		key:    KeyResource{b: b},
		server: ServerResource{b: b},
	}
}

func (crypto *CryptoResource) Cert() CertResource {
	return crypto.cert
}

func (crypto *CryptoResource) Client() ClientResource {
	return crypto.client
}

func (crypto *CryptoResource) CRL() CrlResource {
	return crypto.crl
}
func (crypto *CryptoResource) CSR() CsrResource {
	return crypto.csr
}
func (crypto *CryptoResource) Key() KeyResource {
	return crypto.key
}

func (crypto *CryptoResource) Server() ServerResource {
	return crypto.server
}
