package wt

import "goweb-gin-demo/service"

type ApiWtGroup struct {
	WtReportsApi
}

var wtReportsService = service.ServiceGroupApp.WtReportsService

