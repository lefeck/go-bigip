package raid

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
)

type DiskList struct {
	Kind     string             `json:"kind"`
	SelfLink string             `json:"selfLink"`
	Entries  map[string]Entries `json:"entries"`
}

type Entries struct {
	HTTPSLocalhostMgmtTmSysRaidDiskHD1 struct {
		NestedStats struct {
			Kind     string `json:"kind"`
			SelfLink string `json:"selfLink"`
			Entries  struct {
				ArrayStatus struct {
					Description string `json:"description"`
				} `json:"arrayStatus"`
				IsArrayMember struct {
					Description string `json:"description"`
				} `json:"isArrayMember"`
				Model struct {
					Description string `json:"description"`
				} `json:"model"`
				TmName struct {
					Description string `json:"description"`
				} `json:"tmName"`
				SerialNumber struct {
					Description string `json:"description"`
				} `json:"serialNumber"`
			} `json:"entries"`
		} `json:"nestedStats"`
	} `json:"https://localhost/mgmt/tm/sys/raid/disk/HD1"`
}

const DiskEndpoint = "disk"

type DiskResource struct {
	b *bigip.BigIP
}

// Get a single logical disk details by the node name
func (r *DiskResource) Show() (*DiskList, error) {
	var item DiskList
	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(RAIDEndpoint).Resource(DiskEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &item, nil
}
