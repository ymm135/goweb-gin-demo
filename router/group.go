package router

import "goweb-gin-demo/router/web"

type RouterGroup struct {
	web.BaseRouter
	web.JwtRouter
	web.UserRouter
	web.MenuRouter
	web.SysRouter
	web.FileUploadAndDownloadRouter
}

var RouterGroupApp = new(RouterGroup)
