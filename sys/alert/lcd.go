package alert

import "github.com/lefeck/go-bigip"

// AlertLCDConfigList holds a list of AlertLCD configuration.
type AlertLCDConfigList struct {
	Items    []AlertLCDConfig `json:"items"`
	Kind     string           `json:"kind"`
	SelfLink string           `json:"selflink"`
}

// AlertLCDConfig holds the configuration of a single AlertLCD.
type AlertLCDConfig struct {
}

// AlertLCDEndpoint represents the REST resource for managing AlertLCD.
const LCDEndpoint = "alert"

// AlertLCDResource provides an API to manage AlertLCD configurations.
type AlertLCDResource struct {
	b *bigip.BigIP
}
