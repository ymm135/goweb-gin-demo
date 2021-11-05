package wt

import (
	"goweb-gin-demo/service"
)

type ApiWtGroup struct {
	WtReportsApi
	WtTemplateApi
	WtCommentApi
	WtRuleApi
}

var wtReportsService = service.ServiceGroupApp.WtServiceGroup.WtReportsService
var wtTemplatesService = service.ServiceGroupApp.WtServiceGroup.WtTemplateService
var wtCommentService = service.ServiceGroupApp.WtServiceGroup.WtCommentService
var wtRuleService = service.ServiceGroupApp.WtServiceGroup.WtRuleService
