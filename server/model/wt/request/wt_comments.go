package request

import (
	"goweb-gin-demo/model/common/request"
	"goweb-gin-demo/model/wt"
)

type WtCommentSearch struct{
    wt.WtComment
    request.PageInfo
}