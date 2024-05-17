# go-bigip

go-bigip provides a flexible Restfull API interface for users using bigip, users can flexibly bigip all features.

## Installation

```
go get -u github.com/lefeck/go-bigip
```


## Usage

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


## Features

- [x] Add support for HTTP Basic Authentication
- [ ] Add support for token based authentication
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
- [x] Add support for transaction

## Contributing

We appreciate any form of contribution (feature request, bug report,
pull request, ...). We have no special requirements for Pull Request,
just follow the standard [GitHub way](https://help.github.com/articles/using-pull-requests/).

