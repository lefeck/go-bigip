package main

import (
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"github.com/lefeck/go-bigip/auth"
	"github.com/lefeck/go-bigip/cli"
	"github.com/lefeck/go-bigip/ltm"
	"github.com/lefeck/go-bigip/ltm/monitor"
	bgnet "github.com/lefeck/go-bigip/net"
	"github.com/lefeck/go-bigip/sys"
	"github.com/lefeck/go-bigip/util"
	"log"
	"net"
	"os"
	"regexp"
	"strings"
)

/*
  for example:
    https://192.168.13.91/mgmt/tm/ltm/virtual-address/~Common~1.1.1.1
    Schema://<ip>/<api-prefix>/<resource-category>/<manager>/<resource-type>/<resource-instance>
*/

func main() {
	bs := &bigipTest{}
	bs.init()

	bs.getVersion()
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

	//bs.createVirtualServer()

	//bs.listVirtualServer()

	//bs.ListNetRouteDomain()
	//
	//bs.virtualAddressList()
	//bs.ListICMP()
	//bs.ListNetAddressList()
	//bs.CreateICMP()
	//bs.UpdateICMP()
	//bs.DeleteICMP()
	// profile
	//bs.ListProfileFastHttp()

	// system resource about
	//bs.ListSysServiceList()

	//bs.ListTrafficMatchingCriteria()
	//bs.ListTrafficMatchingCriteriaName()
	//bs.CreateTrafficMatchingCriteria()

	//bs.ListNetPortList()

	//bs.createNetPortList()
}

// this is a testing struct for bigip api
type bigipTest struct {
	bigIP *bigip.BigIP
}

func (bs *bigipTest) init() {
	//b, err := bigip.NewSession("192.168.13.91", "admin", "MsTac@2001")
	//optionTimeout := bigip.WithTimeout(1200)

	b, err := bigip.NewToken("192.168.13.91", "admin", "MsTac@2001", "local")
	//b, err := bigip.NewToken("192.168.13.91", "admin", "MsTac@2001", "local", optionTimeout)
	if err != nil {
		panic(err)
	}
	bs.bigIP = b
}

func (bs *bigipTest) ListTrafficMatchingCriteria() {
	bg := ltm.New(bs.bigIP)
	routeDomainList, _ := bg.TrafficMatchingCriteria().List()
	fmt.Println(routeDomainList)

	for _, icmp := range routeDomainList.Items {
		fullpath := icmp.FullPath
		item, err := bg.TrafficMatchingCriteria().Get(fullpath)
		if err != nil {
			log.Fatalf("route domain get failed %v\n", err)
		}
		bt, _ := json.Marshal(item)
		fmt.Println(string(bt))
	}
}

func (bs *bigipTest) ListTrafficMatchingCriteriaName() {
	bg := ltm.New(bs.bigIP)
	routeDomainList, _ := bg.TrafficMatchingCriteria().ListName()
	fmt.Println(routeDomainList)
}

func (bs *bigipTest) CreateTrafficMatchingCriteria() {
	bg := ltm.New(bs.bigIP)
	item := ltm.TrafficMatchingCriteria{
		Name:                     "vs-testdeom",
		DestinationAddressInline: "0.0.0.0",
		DestinationAddressList:   "/Common/pp1",

		DestinationPortInline: "0",
		DestinationPortList:   "/Common/_sys_self_allow_udp_defaults",

		Protocol:            "udp",
		RouteDomain:         "any",
		SourceAddressInline: "0.0.0.0",
		SourceAddressList:   "/Common/pp2",
		SourcePortInline:    0,
	}

	err := bg.TrafficMatchingCriteria().Create(item)
	if err != nil {
		log.Fatal(err)
	}
}

func (bs *bigipTest) UpdateTrafficMatchingCriteria() {
	bg := ltm.New(bs.bigIP)
	name := "/Common/vs-testdeom"
	item := ltm.TrafficMatchingCriteria{
		Name:                     "vs-testdeom",
		DestinationAddressInline: "0.0.0.0",
		DestinationAddressList:   "/Common/pp1",

		DestinationPortInline: "0",
		DestinationPortList:   "/Common/_sys_self_allow_udp_defaults",

		Protocol:            "udp",
		RouteDomain:         "any",
		SourceAddressInline: "0.0.0.0",
		SourceAddressList:   "/Common/pp2",
		SourcePortInline:    2,
	}

	err := bg.TrafficMatchingCriteria().Update(name, item)
	if err != nil {
		log.Fatal(err)
	}
}

func (bs *bigipTest) ListNetPortList() {
	bg := bgnet.New(bs.bigIP)
	routeDomainList, _ := bg.PortList().List()
	fmt.Println(routeDomainList)

	for _, icmp := range routeDomainList.Items {
		fullpath := icmp.FullPath
		item, err := bg.PortList().Get(fullpath)
		if err != nil {
			log.Fatalf("route domain get failed %v\n", err)
		}
		bt, _ := json.Marshal(item)
		fmt.Println(string(bt))
	}
}

func (bs *bigipTest) createNetPortList() {
	bg := bgnet.New(bs.bigIP)
	item := bgnet.Port{
		Name: "hello-portlist",
		Ports: []bgnet.PortMember{
			{
				Name: "89",
			},
			{
				Name: "8009",
			},
		},
	}
	_ = bg.PortList().Create(item)
}

func (bs *bigipTest) ListNetRouteDomain() {
	bg := bgnet.New(bs.bigIP)
	routeDomainList, _ := bg.RouteDomain().List()
	//fmt.Println(routeDomainList)

	for _, icmp := range routeDomainList.Items {
		fullpath := icmp.FullPath
		item, err := bg.RouteDomain().Get(fullpath)
		if err != nil {
			log.Fatalf("route domain get failed %v\n", err)
		}
		bt, _ := json.Marshal(item)
		fmt.Println(string(bt))
	}
}

func (bs *bigipTest) listVirtualServer() {
	bg := ltm.New(bs.bigIP)
	vs, _ := bg.Virtual().List()
	fmt.Println(vs.Items)
	for _, va := range vs.Items {
		name := va.FullPath
		vs, err := bg.Virtual().Get(name)
		if err != nil {
			panic(err)
		}
		// destination中地址， 判断vs地址 ipv4地址还是ipv6地址， 如果是ipv4地址， 就去除最后的:port, 获取vs的virtualAddress获取真的地址， 如果是ipv6地址就
		//
		fmt.Printf("virtual server name : %s, %s\n", vs.FullPath, vs.Destination)
		des := removePort(vs.Destination)
		addrs, _ := extractAddressWithoutPort(vs.Destination)
		fmt.Println(addrs)
		addr, err := bg.VirtualAddress().GetAddressByVirtualServerName(des)
		if err != nil {
			panic(err)
		}
		fmt.Printf("virtual address name : %s\n", addr)
	}
}

func isIPv4(ipStr string) bool {
	ip := net.ParseIP(ipStr)
	return ip != nil && ip.To4() != nil
}

func isIPv6(ipStr string) bool {
	ip := net.ParseIP(ipStr)
	return ip != nil && ip.To4() == nil && ip.To16() != nil
}

func extractAddressWithoutPort(vsAddress string) (string, error) {
	re := regexp.MustCompile(`/(?P<Partition>[^/]+)/((?P<Address>[0-9a-fA-F:]+)(?P<RouteDomain>%\d+))\.(?P<Port>\d+)$`)
	matches := re.FindStringSubmatch(vsAddress)

	if matches == nil {
		return "", fmt.Errorf("无法匹配给定的虚拟服务器地址")
	}

	result := make(map[string]string)
	for i, name := range re.SubexpNames() {
		if i != 0 && name != "" {
			result[name] = matches[i]
		}
	}

	return "/" + result["Partition"] + "/" + result["Address"] + result["RouteDomain"], nil
}

func (bs *bigipTest) ListNetAddressList() {
	bg := bgnet.New(bs.bigIP)
	addrList, _ := bg.AddressList().List()
	fmt.Println(addrList)

	for _, icmp := range addrList.Items {
		fullpath := icmp.FullPath
		item, err := bg.AddressList().Get(fullpath)
		if err != nil {
			panic(err)
		}
		fmt.Println(item)
	}
}

func (bs *bigipTest) ListSysServiceList() {
	bg := sys.New(bs.bigIP)
	addrList, _ := bg.Service().List()
	//fmt.Println(addrList)

	for _, service := range addrList.Items {
		fullpath := service.FullPath
		item, err := bg.Service().Get(fullpath)
		if err != nil {
			panic(err)
		}
		fmt.Println(item)
	}
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
	version, _ := bga.Version().Show()
	//fmt.Println(version.NestedStats.EntriesMenu.Supported)

	bt, _ := json.Marshal(version)
	fmt.Println(string(bt))
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
func removePort(address string) string {
	if index := strings.LastIndex(address, ":"); index != -1 {
		return address[:index]
	}
	return address
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
		Name:                     "hello-vs13",
		Source:                   "0.0.0.0/32",
		Destination:              "/Common/10.50.11.91:90",
		Mask:                     "255.255.255.255",
		SourceAddressTranslation: ltm.SourceAddressTranslation{Type: "automap"},
		//AddressStatus:            "yes",
		//AutoLasthop:              "default",
		//CmpEnabled:               "yes",
		//EvictionProtected:        "disabled",
		//IPProtocol:               "tcp",
		//Mirror:                   "disabled",
		//MobileAppTunnel:          "disabled",
		//Nat64:                    "disabled",
		//RateLimit:                "disabled",
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
Enabled           ---> "state": "up"
Disabled          ---> "state": "unchecked"
Forced Offline    ---> "state": "user-down"
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
