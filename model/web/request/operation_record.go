package request

import (
	"github.com/ymm135/goweb-gin-demo/model/common/request"
	"github.com/ymm135/goweb-gin-demo/model/web"
)

type SysOperationRecordSearch struct {
	web.SysOperationRecord
	request.PageInfo
}
