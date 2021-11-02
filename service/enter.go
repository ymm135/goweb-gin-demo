package service

type ServiceGroup struct {
	UserService
	JwtService
	CasbinService
	OperationRecordService
	MenuService
	BaseMenuService
	SystemConfigService
}

var ServiceGroupApp = new(ServiceGroup)
