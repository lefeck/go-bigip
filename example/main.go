package main

import (
	"fmt"
	"github.com/lefeck/bigip"
	"github.com/lefeck/bigip/ltm"
)

/*
  for example:
   https://192.168.13.91/mgmt/tm/ltm/virtual-address/

   https://192.168.13.91/mgmt/tm/ltm/virtual-address/~Common~1.1.1.1
*/

func main() {
	virtualserver()
}

func virtualaddress() {
	b, err := bigip.NewSession("192.168.13.91", "admin", "MsTac@2001")
	if err != nil {
		panic(err)
	}
	bg := ltm.New(b)
	val, _ := bg.VirtualAddress().List()
	fmt.Println(val)
	for _, va := range val.Items {
		name := va.FullPath
		address, err := bg.VirtualAddress().GetAddressByVirtualServerName(name)
		if err != nil {
			panic(err)
		}
		fmt.Println(address)
	}
}

func virtualserver() {
	b, err := bigip.NewSession("192.168.13.91", "admin", "MsTac@2001")
	if err != nil {
		panic(err)
	}
	bg := ltm.New(b)
	vs, _ := bg.Virtual().List()
	fmt.Println(vs)
	for _, va := range vs.Items {
		name := va.FullPath
		address, err := bg.Virtual().Get(name)
		if err != nil {
			panic(err)
		}
		fmt.Println(address.Name)
	}
}
