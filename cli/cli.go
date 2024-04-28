package cli

import "github.com/lefeck/bigip"

type Cli struct {
	b       *bigip.BigIP
	version VersionStatsResoure
}

func NewCli(b *bigip.BigIP) Cli {
	return Cli{
		b:       b,
		version: VersionStatsResoure{b: b},
	}
}

func (c Cli) Version() *VersionStatsResoure {
	return &c.version
}
