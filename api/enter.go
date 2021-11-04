package api

import (
	"goweb-gin-demo/api/web"
	"goweb-gin-demo/api/wt"
)

type ApiGroup struct {
	ApiGroup   web.ApiGroup
	WtServiceGroup wt.ApiWtGroup
}

var ApiGroupApp = new(ApiGroup)
