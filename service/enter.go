package service

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
}

var ServiceGroupApp = new(ServiceGroup)
