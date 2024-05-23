package sys

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

// GlobalSettingsEndpoint represents the REST resource for managing GlobalSettings.
const GlobalSettingsEndpoint = "global-settings"

// GlobalSettings holds the configuration of a single GlobalSettings.
type GlobalSettings struct {
	Kind                            string `json:"kind"`
	SelfLink                        string `json:"selfLink"`
	AwsAPIMaxConcurrency            int    `json:"awsApiMaxConcurrency"`
	ConsoleInactivityTimeout        int    `json:"consoleInactivityTimeout"`
	CustomAddr                      string `json:"customAddr"`
	FailsafeAction                  string `json:"failsafeAction"`
	FileBlacklistPathPrefix         string `json:"fileBlacklistPathPrefix"`
	FileBlacklistReadOnlyPathPrefix string `json:"fileBlacklistReadOnlyPathPrefix"`
	FileLocalPathPrefix             string `json:"fileLocalPathPrefix"`
	FileWhitelistPathPrefix         string `json:"fileWhitelistPathPrefix"`
	GuiAudit                        string `json:"guiAudit"`
	GuiExpiredCertAlert             string `json:"guiExpiredCertAlert"`
	GuiSecurityBanner               string `json:"guiSecurityBanner"`
	GuiSecurityBannerText           string `json:"guiSecurityBannerText"`
	GuiSetup                        string `json:"guiSetup"`
	HostAddrMode                    string `json:"hostAddrMode"`
	Hostname                        string `json:"hostname"`
	LcdDisplay                      string `json:"lcdDisplay"`
	LedLocator                      string `json:"ledLocator"`
	MgmtDhcp                        string `json:"mgmtDhcp"`
	NetReboot                       string `json:"netReboot"`
	PasswordPrompt                  string `json:"passwordPrompt"`
	QuietBoot                       string `json:"quietBoot"`
	SSHMaxSessionLimit              int    `json:"sshMaxSessionLimit"`
	SSHMaxSessionLimitPerUser       int    `json:"sshMaxSessionLimitPerUser"`
	SSHRootSessionLimit             string `json:"sshRootSessionLimit"`
	SSHSessionLimit                 string `json:"sshSessionLimit"`
	UsernamePrompt                  string `json:"usernamePrompt"`
}

// GlobalSettingsResource provides an API to manage GlobalSettings configurations.
type GlobalSettingsResource struct {
	b *bigip.BigIP
}

// about global-settting configuration link to https://clouddocs.f5.com/cli/tmsh-reference/v15/modules/sys/sys_global-settings.html
// List retrieves all GlobalSettings details.
func (r *GlobalSettingsResource) Show() (*GlobalSettings, error) {
	var items GlobalSettings
	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(GlobalSettingsEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &items); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &items, nil
}

// Update modifies the GlobalSettings item identified by the GlobalSettings name.
func (r *GlobalSettingsResource) Update(name string, item GlobalSettings) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = r.b.RestClient.Put().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(GlobalSettingsEndpoint).ResourceInstance(name).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
