package web

import (
	"goweb-gin-demo/service"
)

type ApiGroup struct {
	BaseApi
	JwtApi
	AuthorityMenuApi
	AuthorityApi
	FileUploadAndDownloadApi
	SystemApi
}

var userService = service.ServiceGroupApp.SystemServiceGroup.UserService
var jwtService =  service.ServiceGroupApp.SystemServiceGroup.JwtService
var menuService = service.ServiceGroupApp.SystemServiceGroup.MenuService
var baseMenuService = service.ServiceGroupApp.SystemServiceGroup.BaseMenuService
var systemConfigService = service.ServiceGroupApp.SystemServiceGroup.SystemConfigService
var fileUploadAndDownloadService = service.ServiceGroupApp.SystemServiceGroup.FileUploadAndDownloadService
var authorityService = service.ServiceGroupApp.SystemServiceGroup.AuthorityService
var casbinService = service.ServiceGroupApp.SystemServiceGroup.CasbinService

