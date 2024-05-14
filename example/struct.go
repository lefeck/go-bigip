package main

import (
	"fmt"
	"github.com/lefeck/go-bigip"
)

func main() {
	nrn := bigip.NewResourceName()
	name := nrn.BaseResourceNameString()
	fmt.Println(name)
}
