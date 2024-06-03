package util

import (
	"github.com/lefeck/go-bigip"
)

// UtilManager is a commonly used tm, providing a large number of api resource types
const UtilManager = "util"

type Util struct {
	bash BashResource
}

func NewUtil(b *bigip.BigIP) Util {
	return Util{
		bash: BashResource{b: b},
	}
}

func (util Util) Bash() *BashResource {
	return &util.bash
}
