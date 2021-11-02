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
}

var ServiceGroupApp = new(ServiceGroup)
