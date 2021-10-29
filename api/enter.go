package api

import (
	"goweb-gin-demo/api/web"
)

type ApiGroup struct {
	ApiGroup   web.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
