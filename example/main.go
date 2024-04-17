package main

import (
	"fmt"
	"github.com/lefeck/bigip"
	"github.com/lefeck/bigip/ltm"
)

func main() {
	b, err := bigip.NewSession("192.168.13.91", "admin", "MsTac@2001")
	if err != nil {
		panic(err)
	}
	bg := ltm.New(b)
	val, _ := bg.VirtualAddress().List()
	fmt.Println(val)
}

//https://192.168.13.91/mgmt/tm/ltm/virtual-address/~Common~1.1.1.1
//https://192.168.13.91/mgmt/tm/ltm/virtual-address/~Common~1.1.1.1
//                     /mgmt/tm/ltm/virtual-address
