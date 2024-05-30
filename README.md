# go-bigip

go-bigip offers a flexible REST API interface for users working with bigip, making it easy for them to address various problems according to their specific needs.
## Installation

```
go get -u github.com/lefeck/go-bigip
```


## Usage

### Basic Authentication
```go
package main

import (
	"log"
	"github.com/lefeck/go-bigip"
	"github.com/lefeck/go-bigip/ltm"
)

func main() {
	// setup F5 BigIP client
	client, err := bigip.NewSession("192.168.13.91", "admin", "MsTac@2001")
	if err != nil {
		log.Fatal(err)
	}

	// setup client for the LTM API
	ltmClient := ltm.New(client)

	// query the /ltm/virtual API
	vsl, err := ltmClient.Virtual().List()
	if err != nil {
		log.Fatal(err)
	}
	log.Print(vsl)
}
```

### Token Authentication
```go
package main

import (
	"log"
	"github.com/lefeck/go-bigip"
	"github.com/lefeck/go-bigip/ltm"
)

func main() {
	// setup F5 BigIP client
	// default timeout value is 60s
	optionTimeout := bigip.WithTimeout(1200*time.Second)
	client, err := bigip.NewToken("192.168.13.91", "admin", "MsTac@2001", "local", optionTimeout)
	if err != nil {
		log.Fatal(err)
	}

	// setup client for the LTM API
	ltmClient := ltm.New(client)

	// query the /ltm/virtual API
	vsl, err := ltmClient.Virtual().List()
	if err != nil {
		log.Fatal(err)
	}
	log.Print(vsl)
}
```

## Features

- [x] Add support for HTTP Basic Authentication
- [x] Add support for token based authentication
- [ ] Add support for authentication through external providers
- [x] Manage Virtual Server, pool, node, irules, monitors (/ltm)
- [x] Manage Cluster Management (/cm)
- [x] Manage interfaces, vlan, trunk, self ip, route, route domains (/net)
- [x] Manage system related stuffs (/sys)
- [ ] Manage firewall, WAF and DOS profiles (/security)
- [ ] Manage virtualization features (/vcmp)
- [ ] Manage access policies (/apm)
- [x] Manage DNS and global load balancing servers (/gtm)
- [ ] Add support for analytics read-only API (/analytics)
- [ ] Add support for results pagination
