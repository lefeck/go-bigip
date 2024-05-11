package common

import "github.com/lefeck/go-bigip"

type CommonResource struct {
	B *bigip.BigIP
}

func NewCommon(b *bigip.BigIP) *CommonResource {
	return &CommonResource{
		B: b,
	}
}

//func tesyt()  {
//	s := CommonResource{b}
//	s.b.RestClient.Get().Prefix()
//}
