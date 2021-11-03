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

var userService = service.ServiceGroupApp.UserService
var jwtService =  service.ServiceGroupApp.JwtService
var menuService = service.ServiceGroupApp.MenuService
var baseMenuService = service.ServiceGroupApp.BaseMenuService
var systemConfigService = service.ServiceGroupApp.SystemConfigService
var fileUploadAndDownloadService = service.ServiceGroupApp.FileUploadAndDownloadService
var authorityService = service.ServiceGroupApp.AuthorityService
var casbinService = service.ServiceGroupApp.CasbinService

