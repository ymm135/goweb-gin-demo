package request

import (
	"goweb-gin-demo/model/common/request"
	"goweb-gin-demo/model/web"
)

type SysOperationRecordSearch struct {
	web.SysOperationRecord
	request.PageInfo
}
