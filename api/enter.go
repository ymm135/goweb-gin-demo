package api

import (
	"github.com/ymm135/goweb-gin-demo/api/web"
)

type ApiGroup struct {
	ApiGroup   web.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
