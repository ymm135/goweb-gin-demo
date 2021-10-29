package service

type ServiceGroup struct {
	UserService
	JwtService
	CasbinService
	OperationRecordService
	MenuService
	BaseMenuService
}

var ServiceGroupApp = new(ServiceGroup)
