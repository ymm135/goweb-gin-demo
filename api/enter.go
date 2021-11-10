package api

import (
	"goweb-gin-demo/api/system"
	"goweb-gin-demo/api/wt"
)

type ApiGroup struct {
	ApiGroup       system.ApiGroup
	WtServiceGroup wt.ApiWtGroup
}

var ApiGroupApp = new(ApiGroup)
