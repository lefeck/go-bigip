package main

import (
	"fmt"
	"github.com/lefeck/go-bigip"
	"github.com/lefeck/go-bigip/gtm"
	"log"
)

// this is a testing struct for bigip api
type bigipGTM struct {
	bigIP *bigip.BigIP
}

func (bs *bigipGTM) initGTM() {
	//b, err := bigip.NewSession("192.168.13.91", "admin", "MsTac@2001")
	//optionTimeout := bigip.WithTimeout(1200 * time.Second)

	b, err := bigip.NewToken("192.168.13.93", "admin", "MsTac@2001", "local")
	//b, err := bigip.NewToken("192.168.13.91", "admin", "MsTac@2001", "local", optionTimeout)
	if err != nil {
		panic(err)
	}
	bs.bigIP = b
}

func (bs *bigipGTM) ListDataCenter() {
	bg := gtm.New(bs.bigIP)
	dataList, _ := bg.Datacenter().List()
	//fmt.Println(addrList)

	for _, icmp := range dataList.Items {
		fullpath := icmp.FullPath
		item, err := bg.Datacenter().Get(fullpath)
		if err != nil {
			panic(err)
		}
		fmt.Println(item)
	}
}

func (bs *bigipGTM) CreateDataCenter() {
	dc := gtm.Datacenter{
		Name: "dc-mobile",
	}
	bg := gtm.New(bs.bigIP)
	_ = bg.Datacenter().Create(dc)
	//fmt.Println(addrList)
}

func (bs *bigipGTM) UpdateDataCenter() {
	name := "dc-mobile"
	dc := gtm.Datacenter{
		Location: "mobile",
		Enabled:  false,
	}
	bg := gtm.New(bs.bigIP)
	_ = bg.Datacenter().Update(name, dc)
	item, err := bg.Datacenter().Get(name)
	if err != nil {
		panic(err)
	}
	fmt.Println(item)
}

func (bs *bigipGTM) DeleteDataCenter() {
	name := "dc-mobile"

	bg := gtm.New(bs.bigIP)
	err := bg.Datacenter().Delete(name)
	if err != nil {
		log.Fatalf("error : %s\n", err)
	}
	item, err := bg.Datacenter().Get(name)
	if err != nil {
		log.Fatalf("error : %s\n", err)
	}
	fmt.Println(item)
}

func main() {
	bs := &bigipGTM{}
	bs.initGTM()
	bs.ListDataCenter()

	//bs.CreateDataCenter()
	//bs.UpdateDataCenter()
	//bs.DeleteDataCenter()

}
