package main

import (
	"encoding/json"
	"fmt"
	"github.com/lefeck/bigip"
	"github.com/lefeck/bigip/auth"
	"github.com/lefeck/bigip/cli"
	"github.com/lefeck/bigip/ltm"
	"github.com/lefeck/bigip/util"
	"log"
	"os"
)

/*
  for example:

   https://192.168.13.91/mgmt/tm/ltm/virtual-address/

   <IP>/<api-prefix>/<resource-category>/<manager>/<resource-type>/<resource-instance>

   https://192.168.13.91/mgmt/tm/ltm/virtual-address/~Common~1.1.1.1
*/

func main() {
	//virtualaddress()
	//virtualserver()
	//virtualtoken()
	//createVS()
	//modifyVS()
	//UserTest()
	//VersionTest()
	//getSinglevVS()

	modifyVsStateToEnable()
	//utilTest()
}

func virtualtoken() {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWluIiwicGFzc3dvcmQiOiJNc1RhY0AyMDAxIiwic3ViIjoiZGVtbyIsImlhdCI6MTcxMzQ5NDYzNSwibmJmIjoxNzEzNDk0NjM1LCJleHAiOjE3MTM1ODEwMzV9.j1vnb6LonkCDxs7bbfDovjbFHSjRk7vCZAns5Bwiqf8"
	b, err := bigip.NewToken("192.168.13.91", token)
	if err != nil {
		panic(err)
	}
	bg := ltm.New(b)
	val, _ := bg.VirtualAddress().List()
	fmt.Println(val)
	for _, va := range val.Items {
		name := va.FullPath
		_, err := bg.VirtualAddress().GetAddressByVirtualServerName(name)
		if err != nil {
			panic(err)
		}
		//fmt.Println(address)
	}
}

func virtualaddress() {
	b, err := bigip.NewSession("192.168.13.91", "admin", "MsTac@2001")
	if err != nil {
		panic(err)
	}
	bg := ltm.New(b)
	val, _ := bg.VirtualAddress().List()
	fmt.Println(val)
	//fmt.Println(val.Items)
	for _, va := range val.Items {
		name := va.FullPath
		va, err := bg.VirtualAddress().GetAddressByVirtualServerName(name)
		if err != nil {
			panic(err)
		}
		fmt.Println(va)
	}
}

func VersionTest() {
	b, err := bigip.NewSession("192.168.13.91", "admin", "MsTac@2001")
	if err != nil {
		panic(err)
	}
	bga := cli.NewCli(b)

	version, _ := bga.Version().Get()
	fmt.Println(version.Entries.HTTPSLocalhostMgmtTmCliVersion0.NestedStats.EntriesMenu.Supported)

	//bt, _ := json.Marshal(version)
	//
	//fmt.Println(string(bt))
}

func utilTest() {
	b, err := bigip.NewSession("192.168.10.101", "admin", "admin")
	//b, err := bigip.NewSession("192.168.13.91", "admin", "MsTac@2001")
	if err != nil {
		panic(err)
	}
	bga := util.NewUtil(b)

	item := util.Bash{
		Command: "run",
		//UtilCmdArgs: "uptime",
		UtilCmdArgs: " -c  tmsh list ltm virtual  ",
	}

	bashr, _ := bga.Bash().Run(item)
	fmt.Println(bashr)
	//bt, err := json.Marshal(bashr)
	//if err != nil {
	//	panic(err)
	//}
	//os.Stdout.Write(bt)
}

func UserTest() {
	//b, _ := Newbigip()
	b, err := bigip.NewSession("192.168.13.91", "admin", "MsTac@2001")
	if err != nil {
		panic(err)
	}
	bga := auth.NewAuth(b)
	user, _ := bga.Users().List()
	fmt.Println(user)

}

func virtualserverlist() {
	b, err := bigip.NewSession("192.168.13.91", "admin", "MsTac@2001")
	if err != nil {
		panic(err)
	}
	bg := ltm.New(b)
	vs, _ := bg.Virtual().List()
	fmt.Println(vs.Items)
	for _, va := range vs.Items {
		name := va.FullPath
		address, err := bg.Virtual().Get(name)
		if err != nil {
			panic(err)
		}
		fmt.Println(address.Name)
	}
}

func modifyVS() {
	b, err := bigip.NewSession("192.168.13.91", "admin", "MsTac@2001")
	if err != nil {
		panic(err)
	}
	bg := ltm.New(b)
	name := "/Common/hello-vs1"
	item := ltm.VirtualServer{
		IPProtocol:               "tcp",
		Source:                   "0.0.0.0/32",
		Destination:              "192.168.83.26:9090",
		Mask:                     "255.255.255.255",
		SourceAddressTranslation: ltm.SourceAddressTranslation{Type: "automap"},
		ConnectionLimit:          1000,
	}

	if err := bg.Virtual().Update(name, item); err != nil {
		log.Fatalf("update virtual server is failed %v", err)
	}
}

func modifyVsStateToDisable() {
	b, err := bigip.NewSession("192.168.13.91", "admin", "MsTac@2001")
	if err != nil {
		panic(err)
	}
	bg := ltm.New(b)
	name := "/Common/hello-vs1"
	if err := bg.Virtual().Disable(name); err != nil {
		log.Fatalf("disable virtual server is failed %v", err)
	}
}

func modifyVsStateToEnable() {
	b, err := bigip.NewSession("192.168.13.91", "admin", "MsTac@2001")
	if err != nil {
		panic(err)
	}
	bg := ltm.New(b)
	name := "/Common/hello-vs1"
	if err := bg.Virtual().Enable(name); err != nil {
		log.Fatalf("enabled virtual server is failed %v", err)
	}

}

func createVS() {
	b, err := bigip.NewSession("192.168.13.91", "admin", "MsTac@2001")
	if err != nil {
		panic(err)
	}
	bg := ltm.New(b)

	item := ltm.VirtualServer{
		Name:                     "hello-vs1",
		Destination:              "192.168.83.23:90",
		Mask:                     "255.255.255.255",
		SourceAddressTranslation: ltm.SourceAddressTranslation{Type: "automap"},
	}

	if err := bg.Virtual().Create(item); err != nil {
		log.Fatalf("create virtual server is failed %v", err)
	}
}

func getSinglevVS() {
	b, err := bigip.NewSession("192.168.13.91", "admin", "MsTac@2001")
	if err != nil {
		panic(err)
	}
	bg := ltm.New(b)
	name := "/Common/go-test"

	vs, err := bg.Virtual().Get(name)
	if err != nil {
		log.Fatalf("get virtual server failed  %v", err)
	}

	bt, _ := json.Marshal(vs)
	os.Stdout.Write(bt)
}

func deletevs() {
	b, err := bigip.NewSession("192.168.13.91", "admin", "MsTac@2001")
	if err != nil {
		panic(err)
	}
	bg := ltm.New(b)
	name := "/Common/go-test"

	if err := bg.Virtual().Delete(name); err != nil {
		log.Fatalf("delete virtual server failed  %v", err)
	}
}
