package cli

import "github.com/lefeck/go-bigip"

// CliManager is a commonly used basepath, providing a large number of api resource types
const CliManager = "cli"

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
