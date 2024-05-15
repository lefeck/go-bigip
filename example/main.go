package main

import (
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"github.com/lefeck/go-bigip/auth"
	"github.com/lefeck/go-bigip/cli"
	"github.com/lefeck/go-bigip/ltm"
	"github.com/lefeck/go-bigip/ltm/monitor"
	"github.com/lefeck/go-bigip/util"
	"log"
	"os"
)

//	func virtualtoken() {
//		token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWluIiwicGFzc3dvcmQiOiJNc1RhY0AyMDAxIiwic3ViIjoiZGVtbyIsImlhdCI6MTcxMzQ5NDYzNSwibmJmIjoxNzEzNDk0NjM1LCJleHAiOjE3MTM1ODEwMzV9.j1vnb6LonkCDxs7bbfDovjbFHSjRk7vCZAns5Bwiqf8"
//		b, err := bigip.Token("192.168.13.91", token)
//		if err != nil {
//			panic(err)
//		}
//		bg := (b)
//		val, _ := bg.VirtualAddress().List()
//		fmt.Println(val)
//		for _, va := range val.Items {
//			name := va.FullPath
//			_, err := bg.VirtualAddress().GetAddressByVirtualServerName(name)
//			if err != nil {
//				panic(err)
//			}
//			//fmt.Println(address)
//		}
//	}
//

/*
  for example:
    https://192.168.13.91/mgmt/tm/ltm/virtual-address/~Common~1.1.1.1
    Schema://<ip>/<api-prefix>/<resource-category>/<manager>/<resource-type>/<resource-instance>
*/

func main() {
	bs := &bigipTest{}
	bs.init()
	//bs.listRules()
	//bs.listPool()
	//bs.createPool()
	//bs.updatePool()
	//bs.deletePool()
	//bs.listPoolMembers()
	//bs.createPoolMembers()
	//bs.updatePoolMembers()
	//bs.deletePoolMembers()
	//bs.listPoolStats()
	//bs.getSinglePoolStats()
	//bs.getSingleMemberStats()
	//bs.getPoolAllMemberStats()
	//bs.listVirtualServerStats()
	//bs.getSingleVirtualServerStats()
	//bs.listVirtualAddressStats()
	//bs.getSingleVirtualAddressStats()
	//bs.listVirtualServerDetail()
	//bs.listSnatPool()

	bs.ListICMP()
	//bs.CreateICMP()
	//bs.UpdateICMP()
	//bs.DeleteICMP()
	// profile
	bs.ListProfileFastHttp()
}

// this is a testing struct for bigip api
type bigipTest struct {
	bigIP *bigip.BigIP
}

func (bs *bigipTest) init() {
	b, err := bigip.NewSession("192.168.13.91", "admin", "MsTac@2001")
	if err != nil {
		panic(err)
	}
	bs.bigIP = b
}

func (bs *bigipTest) ListICMP() {
	bg := ltm.New(bs.bigIP)
	icmpList, _ := bg.Monitor().ICMP().List()
	fmt.Println(icmpList)

	for _, icmp := range icmpList.Items {
		fullpath := icmp.FullPath
		item, err := bg.Monitor().ICMP().Get(fullpath)
		if err != nil {
			panic(err)
		}
		fmt.Println(item)
	}
}

func (bs *bigipTest) ListProfileFastHttp() {
	bg := ltm.New(bs.bigIP)
	fasthttp, _ := bg.Profile().FastHTTP().List()

	fmt.Println(fasthttp)

	for _, icmp := range fasthttp.Items {
		fullpath := icmp.FullPath
		item, err := bg.Profile().FastHTTP().Get(fullpath)
		if err != nil {
			panic(err)
		}
		fmt.Println(item)
	}
}

func (bs *bigipTest) CreateICMP() {
	bg := ltm.New(bs.bigIP)
	item := monitor.ICMP{
		Name:          "hello-icmp-m1",
		Interval:      20,
		Timeout:       50,
		AdaptiveLimit: 100,
	}
	if err := bg.Monitor().ICMP().Create(item); err != nil {
		panic(err)
	}
}

func (bs *bigipTest) UpdateICMP() {
	bg := ltm.New(bs.bigIP)
	fullPathname := "/Common/hello-icmp-m1"

	item := monitor.ICMP{
		Interval:      10,
		Timeout:       80,
		AdaptiveLimit: 100,
	}
	if err := bg.Monitor().ICMP().Update(fullPathname, item); err != nil {
		panic(err)
	}
}

func (bs *bigipTest) DeleteICMP() {
	bg := ltm.New(bs.bigIP)
	name := "/Common/hello-icmp-m1"

	if err := bg.Monitor().ICMP().Delete(name); err != nil {
		log.Fatalf("delete  is failed %v", err)
	}
}

func (bs *bigipTest) ListProfileSSLClient() {
	bg := ltm.New(bs.bigIP)
	// 	bg.Profile().SSLClient().List()
	fmt.Println(bg)
}

func (bs *bigipTest) listPoolStats() {
	bg := ltm.New(bs.bigIP)
	ps, _ := bg.PoolStats().List()
	fmt.Println(ps)
	//for key, va := range ps.Entries {
	//	name := va.NestedPoolStats.Entries
	//	va, err := bg.PoolStats().GetPoolStats(name)
	//	if err != nil {
	//		panic(err)
	//	}
	//	fmt.Println(va)
	//}
}

func (bs *bigipTest) getSinglePoolStats() {
	bg := ltm.New(bs.bigIP)
	name := "/Common/hello-pool"
	ps, err := bg.PoolStats().GetPoolStats(name)
	if err != nil {
		panic(err)
	}
	fmt.Println(ps)
}

func (bs *bigipTest) listVirtualServerStats() {
	bg := ltm.New(bs.bigIP)
	vs, err := bg.VirtualStats().List()
	if err != nil {
		panic(err)
	}
	fmt.Println(vs)
}

func (bs *bigipTest) getSingleVirtualServerStats() {
	bg := ltm.New(bs.bigIP)
	vsName := "/Common/10.100.131.91"
	ps, err := bg.VirtualStats().Get(vsName)
	if err != nil {
		panic(err)
	}
	bt, _ := json.Marshal(ps)
	fmt.Println(string(bt))
	//fmt.Println(ps)
}

func (bs *bigipTest) listVirtualAddressStats() {
	bg := ltm.New(bs.bigIP)
	vs, err := bg.VirtualAddressStats().List()
	if err != nil {
		panic(err)
	}
	fmt.Println(vs)
}

func (bs *bigipTest) getSingleVirtualAddressStats() {
	bg := ltm.New(bs.bigIP)
	vsName := "/Common/10.100.131.91"
	ps, err := bg.VirtualAddressStats().Get(vsName)
	if err != nil {
		panic(err)
	}
	bt, _ := json.Marshal(ps)
	fmt.Println(string(bt))
	//fmt.Println(ps)
}

func (bs *bigipTest) getSingleMemberStats() {
	bg := ltm.New(bs.bigIP)
	poolName := "/Common/hello-pool"
	memberName := "/Common/142.10.3.2:4523"
	ps, err := bg.PoolStats().GetMemberStats(poolName, memberName)
	if err != nil {
		panic(err)
	}
	fmt.Println(ps)
}

func (bs *bigipTest) getPoolAllMemberStats() {
	bg := ltm.New(bs.bigIP)
	poolName := "/Common/hello-pool"
	//memberName := "/Common/142.10.3.2:4523"
	ps, err := bg.PoolStats().GetPoolAllMemberStats(poolName)
	if err != nil {
		panic(err)
	}
	fmt.Println(ps)
}

func (bs *bigipTest) virtualAddressList() {
	bg := ltm.New(bs.bigIP)
	val, _ := bg.VirtualAddress().List()
	fmt.Println(val)
	for _, va := range val.Items {
		name := va.FullPath
		va, err := bg.VirtualAddress().GetAddressByVirtualServerName(name)
		if err != nil {
			panic(err)
		}
		fmt.Println(va)
	}
}

func (bs *bigipTest) getVersion() {

	bga := cli.NewCli(bs.bigIP)
	version, _ := bga.Version().Get()
	fmt.Println(version.Entries.HTTPSLocalhostMgmtTmCliVersion0.NestedStats.EntriesMenu.Supported)

	//bt, _ := json.Marshal(version)
	//fmt.Println(string(bt))
}

func (bs *bigipTest) useUtil() {
	bga := util.NewUtil(bs.bigIP)

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

func (bs *bigipTest) listUser() {
	bga := auth.NewAuth(bs.bigIP)
	user, _ := bga.Users().List()
	fmt.Println(user)
}

func (bs *bigipTest) listVirtualServer() {
	bg := ltm.New(bs.bigIP)
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

func (bs *bigipTest) listVirtualServerDetail() {
	bg := ltm.New(bs.bigIP)
	vs, _ := bg.Virtual().ListDetail()
	fmt.Println(vs.Items)
	//for _, va := range vs.Items {
	//	name := va.FullPath
	//	address, err := bg.Virtual().Get(name)
	//	if err != nil {
	//		panic(err)
	//	}
	//	fmt.Println(address.Name)
	//}
}

func (bs *bigipTest) listSnatPool() {
	bg := ltm.New(bs.bigIP)
	spl, _ := bg.SnatPool().List()
	fmt.Println(spl.Items)
	for _, sp := range spl.Items {
		name := sp.FullPath
		address, err := bg.SnatPool().Get(name)
		if err != nil {
			panic(err)
		}
		fmt.Println(address.Name)
	}
}

func (bs *bigipTest) listPool() {
	bg := ltm.New(bs.bigIP)
	pools, _ := bg.Pool().List()
	//fmt.Println(pools)
	for _, pool := range pools.Items {
		name := pool.FullPath
		fmt.Println(name)
		pl, err := bg.Pool().Get(name)
		if err != nil {
			panic(err)
		}
		fmt.Println(pl)
	}
}

func (bs *bigipTest) createPool() {
	bg := ltm.New(bs.bigIP)
	item := ltm.Pool{
		Name:              "hello-pool-bg3",
		LoadBalancingMode: "round-robin",
		Members:           []string{"192.13.23.1:90", "128.3.2.53:90"},
		Monitor:           "http",
	}

	if err := bg.Pool().Create(item); err != nil {
		log.Fatalf("create pool is failed %v", err)
	}
}

func (bs *bigipTest) updatePool() {
	bg := ltm.New(bs.bigIP)
	name := "/Common/hello-pool"
	item := ltm.Pool{
		LoadBalancingMode: "fastest-node",
	}

	if err := bg.Pool().Update(name, item); err != nil {
		log.Fatalf("update pool is failed %v", err)
	}
}

func (bs *bigipTest) deletePool() {
	bg := ltm.New(bs.bigIP)
	name := "/Common/hello-pool"

	if err := bg.Pool().Delete(name); err != nil {
		log.Fatalf("delete pool is failed %v", err)
	}
}

func (bs *bigipTest) updateVirtualServer() {
	bg := ltm.New(bs.bigIP)
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

func (bs *bigipTest) updateVSStateToDisable() {
	bg := ltm.New(bs.bigIP)
	name := "/Common/hello-vs1"
	if err := bg.Virtual().Disable(name); err != nil {
		log.Fatalf("disable virtual server is failed %v", err)
	}
}

func (bs *bigipTest) updateVSStateToEnable() {
	bg := ltm.New(bs.bigIP)
	name := "/Common/hello-vs1"
	if err := bg.Virtual().Enable(name); err != nil {
		log.Fatalf("enabled virtual server is failed %v", err)
	}

}

func (bs *bigipTest) createVirtualServer() {
	bg := ltm.New(bs.bigIP)

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

func (bs *bigipTest) getVirtualServer() {
	bg := ltm.New(bs.bigIP)
	name := "/Common/go-test"

	vs, err := bg.Virtual().Get(name)
	if err != nil {
		log.Fatalf("get virtual server failed  %v", err)
	}

	bt, _ := json.Marshal(vs)
	os.Stdout.Write(bt)
}

func (bs *bigipTest) deleteVirtualServer() {
	bg := ltm.New(bs.bigIP)
	name := "/Common/go-test"

	if err := bg.Virtual().Delete(name); err != nil {
		log.Fatalf("delete virtual server failed  %v", err)
	}
}

func (bs *bigipTest) listRules() {
	bg := ltm.New(bs.bigIP)
	rulelist, err := bg.Rule().List()
	if err != nil {
		panic(err)
	}
	//fmt.Println(rulelist)
	for _, rule := range rulelist.Items {
		fullpath := rule.FullPath
		rl, err := bg.Rule().Get(fullpath)
		if err != nil {
			log.Fatalf("get rule failed  %v", err)
		}
		fmt.Println(rl)
		//bt, _ := json.Marshal(rl)
		//os.Stdout.Write(bt)
	}
}

func (bs *bigipTest) createRule() {

	bg := ltm.New(bs.bigIP)

	data, err := os.ReadFile("example/irule.tcl")
	if err != nil {
		log.Fatal(err)
	}

	rule := ltm.Rule{
		Name:         "test_rule",
		ApiAnonymous: string(data),
	}

	err = bg.Rule().Create(rule)
	if err != nil {
		log.Fatalf("create rule failed %v\n", err)
	}
}

func (bs *bigipTest) updateRule() {
	bg := ltm.New(bs.bigIP)

	data, err := os.ReadFile("example/irule-test.tcl")
	if err != nil {
		log.Fatal(err)
	}

	name := "/Common/test_rule"
	rule := ltm.Rule{
		ApiAnonymous: string(data),
	}

	err = bg.Rule().Update(name, rule)
	if err != nil {
		log.Fatalf("update rule failed %v\n", err)
	}
}

func (bs *bigipTest) deleteRule() {
	bg := ltm.New(bs.bigIP)
	ruleName := "/Common/test_rule"

	err := bg.Rule().Delete(ruleName)
	if err != nil {
		log.Fatalf("delete rule failed %v\n", err)
	}
}

func (bs *bigipTest) listPoolMembers() {
	bg := ltm.New(bs.bigIP)
	poolName := "/Project_abdc90059f18487d847d439167b01928/Project_ef79a1d0-a437-4f70-8f08-9a43d8998c25"
	poolMembers, _ := bg.PoolMembers().List(poolName)
	fmt.Println(poolMembers)
	for _, pool := range poolMembers.Items {
		memberName := pool.FullPath
		//fmt.Println(name)
		pl, err := bg.PoolMembers().Get(poolName, memberName)
		if err != nil {
			panic(err)
		}
		fmt.Println(pl)
	}
}

func (bs *bigipTest) createPoolMembers() {
	bg := ltm.New(bs.bigIP)
	poolName := "/Common/hello-pool"
	item := ltm.PoolMembers{
		//Name: "142.10.3.2:4523",
		Name: "142.10.3.3:4523",
	}
	if err := bg.PoolMembers().Create(poolName, item); err != nil {
		log.Fatalf("create pool member is failed %v", err)
	}
}

/*
Enabled --->
Forced --> "state": "unchecked"
*/
func (bs *bigipTest) updatePoolMembers() {
	bg := ltm.New(bs.bigIP)
	poolName := "/Common/hello-pool"
	memberName := "/Common/142.10.3.2:4523"
	item := ltm.PoolMembers{
		ConnectionLimit: 1000,
		Ratio:           10,
		//:         "enable",
	}

	if err := bg.PoolMembers().Update(poolName, memberName, item); err != nil {
		log.Fatalf("update pool member is failed %v", err)
	}
}

func (bs *bigipTest) deletePoolMembers() {
	bg := ltm.New(bs.bigIP)
	poolName := "/Common/hello-pool"
	memberName := "/Common/142.10.3.2:4523"

	if err := bg.PoolMembers().Delete(poolName, memberName); err != nil {
		log.Fatalf("delete pool member is failed %v", err)
	}
}
