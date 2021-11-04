package service

import "goweb-gin-demo/service/wt"

type ServiceGroup struct {
	UserService
	JwtService
	CasbinService
	OperationRecordService
	MenuService
	BaseMenuService
	SystemConfigService
	FileUploadAndDownloadService
	AuthorityService

	wt.WtReportsService
	wt.WtTemplateService
}

var ServiceGroupApp = new(ServiceGroup)
