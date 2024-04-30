package util

import (
	"github.com/lefeck/bigip"
)

type Util struct {
	b    *bigip.BigIP
	bash BashResource
}

func NewUtil(b *bigip.BigIP) Util {
	return Util{
		b:    b,
		bash: BashResource{b: b},
	}
}

func (util Util) Bash() *BashResource {
	return &util.bash
}

// UtilManager is a commonly used basepath, providing a large number of api resource types
const UtilManager = "util"
