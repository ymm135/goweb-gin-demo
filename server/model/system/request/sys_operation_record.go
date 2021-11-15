package request

import (
	"goweb-gin-demo/model/common/request"
	"goweb-gin-demo/model/system"
)

type SysOperationRecordSearch struct {
	system.SysOperationRecord
	request.PageInfo
}
