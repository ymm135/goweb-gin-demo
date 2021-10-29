package web

import (
	"goweb-gin-demo/service"
)

type ApiGroup struct {
	BaseApi
	JwtApi
	AuthorityMenuApi
}

var userService = service.ServiceGroupApp.UserService
var jwtService =  service.ServiceGroupApp.JwtService
var menuService = service.ServiceGroupApp.MenuService
var baseMenuService = service.ServiceGroupApp.BaseMenuService

