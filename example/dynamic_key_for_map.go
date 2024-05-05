package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

/*
在 Golang 中，映射类型 map 是用于表示键值对集合的数据结构。当我们知道 JSON 对象的具体键时，通常使用结构体字段；
但是，当键是动态的或不确定，或者只关心值的处理而不关心特定的键，则使用映射类型 map。在这个例子中，entries 的内容具有动态字符串键，
所以选择使用 map[string]Entry。
*/

type VirtualAddressStats struct {
	Kind       string           `json:"kind"`
	Generation int              `json:"generation"`
	SelfLink   string           `json:"selfLink"`
	Entries    map[string]Entry `json:"entries"`
}

type Entry struct {
	NestedStats NestedStats `json:"nestedStats"`
}

type NestedStats struct {
	Entries NestedEntries `json:"entries"`
}

type NestedEntries struct {
	Addr                    DescriptionOrValue `json:"addr"`
	ClientsideBitsIn        DescriptionOrValue `json:"clientside.bitsIn"`
	ClientsideBitsOut       DescriptionOrValue `json:"clientside.bitsOut"`
	ClientsideCurConns      DescriptionOrValue `json:"clientside.curConns"`
	ClientsideMaxConns      DescriptionOrValue `json:"clientside.maxConns"`
	ClientsidePktsIn        DescriptionOrValue `json:"clientside.pktsIn"`
	ClientsidePktsOut       DescriptionOrValue `json:"clientside.pktsOut"`
	ClientsideTotConns      DescriptionOrValue `json:"clientside.totConns"`
	TmName                  DescriptionOrValue `json:"tmName"`
	StatusAvailabilityState DescriptionOrValue `json:"status.availabilityState"`
	StatusEnabledState      DescriptionOrValue `json:"status.enabledState"`
	StatusStatusReason      DescriptionOrValue `json:"status.statusReason"`
}

type DescriptionOrValue struct {
	Description string `json:"description,omitempty"`
	Value       int    `json:"value,omitempty"`
}

func main() {
	jsonBlob, err := os.ReadFile("example/dynamic_key_for_map.json")
	if err != nil {
		log.Fatalf("read file failed : %v", err)
	}
	var stats VirtualAddressStats
	if err := json.Unmarshal(jsonBlob, &stats); err != nil {
		log.Fatalf("json Unmarshal failed : %v", err)
	}

	fmt.Printf("%+v\n", stats)
}

/*
result:
在原始 JSON 中，entries 是一个对象，它包含一个键值对，键是字符串（在这个例子中是 https://localhost/mgmt/tm/ltm/virtual-address/~Common~10.100.131.91/stats），值是（嵌套的结构体 Entry）。因此，我们在 Golang 中使用 map[string]Entry 来表示这种类似字典的对象结构。
{Kind:tm:ltm:virtual-address:virtual-addressstats Generation:1 SelfLink:https://localhost/mgmt/tm/ltm/virtual-address/~Common~10.100.131.91/stats?ver=17.0.0.1 Entries:map[https://localhost/mgmt/tm/ltm/virtual-address/~Common~10.100.131.91/stats:{NestedStats:{Entries:{Addr:{Description:10.100.131.91 Value:0} ClientsideBitsIn:{Description: Value:0} ClientsideBitsOut:{Description: Value:0} ClientsideCurConns:{Description: Value:0} ClientsideMaxConns:{Description: Value:0} ClientsidePktsIn:{Description: Value:0} ClientsidePktsOut:{Description: Value:0} ClientsideTotConns:{Description: Value:0} TmName:{Description:/Common/10.100.131.91 Value:0} StatusAvailabilityState:{Description:available Value:0} StatusEnabledState:{Description:enabled Value:0} StatusStatusReason:{Description:The virtual address is available Value:0}}}}]}

*/
