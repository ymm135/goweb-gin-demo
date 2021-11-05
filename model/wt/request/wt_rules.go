package request

import (
	"goweb-gin-demo/model/common/request"
	"goweb-gin-demo/model/wt"
)

type WtRuleSearch struct{
    wt.WtRule
    request.PageInfo
}

type WtRuleRes struct {
	ID uint
	UserId  int
	Reporters  []string
	StartTime  string
	EndTime  string
}