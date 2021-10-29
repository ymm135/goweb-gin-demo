package router

import "goweb-gin-demo/router/web"

type RouterGroup struct {
	web.BaseRouter
	web.JwtRouter
	web.UserRouter
	web.MenuRouter
}

var RouterGroupApp = new(RouterGroup)
