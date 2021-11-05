package service

import (
	"goweb-gin-demo/service/system"
	"goweb-gin-demo/service/wt"
)

type ServiceGroup struct {
	SystemServiceGroup system.SystemServiceGroup
	WtServiceGroup     wt.WtServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
