package cli

import "github.com/lefeck/go-bigip"

// CliManager is a commonly used bigip.GetBaseResource(), providing a large number of api resource types
const CliManager = "cli"

type Cli struct {
	version VersionStatsResoure
}

func NewCli(b *bigip.BigIP) Cli {
	return Cli{
		version: VersionStatsResoure{b: b},
	}
}

func (c Cli) Version() *VersionStatsResoure {
	return &c.version
}
