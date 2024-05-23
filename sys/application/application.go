package application

import "github.com/lefeck/go-bigip"

// ApplicationEndpoint represents the REST resource for managing Application.
const ApplicationEndpoint = "application"

const SysManager = "sys"

// ApplicationResource provides an API to manage Application configurations.
type ApplicationResource struct {
	//b       *bigip.BigIP
	service    ServiceResource
	aplScript  APLScriptResource
	customStat CustomStatResource
	template   TemplateResource
}

func NewApplication(b *bigip.BigIP) ApplicationResource {
	return ApplicationResource{
		service:    ServiceResource{b: b},
		aplScript:  APLScriptResource{b: b},
		customStat: CustomStatResource{b: b},
		template:   TemplateResource{b: b},
	}
}

func (app ApplicationResource) Service() *ServiceResource {
	return &app.service
}
func (app ApplicationResource) APLScript() *APLScriptResource {
	return &app.aplScript
}
func (app ApplicationResource) CustomStat() *CustomStatResource {
	return &app.customStat
}
func (app ApplicationResource) Template() *TemplateResource {
	return &app.template
}
