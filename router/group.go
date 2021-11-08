package router

import (
	"goweb-gin-demo/router/web"
	"goweb-gin-demo/router/wt"
)

type RouterGroup struct {
	web.BaseRouter
	web.JwtRouter
	web.UserRouter
	web.MenuRouter
	web.SysRouter
	web.FileUploadAndDownloadRouter
	web.AuthorityRouter

	wt.WtReportsRouter
	wt.WtRuleRouter
	wt.WtOutputRouter
}

var RouterGroupApp = new(RouterGroup)
