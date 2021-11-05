package system

type SystemServiceGroup struct {
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
