package wt

import "goweb-gin-demo/service"

type ApiWtGroup struct {
	WtReportsApi
	WtTemplateApi
}

var wtReportsService = service.ServiceGroupApp.WtReportsService
var wtTemplatesService = service.ServiceGroupApp.WtTemplateService

