package request

import (
	"goweb-gin-demo/model/common/request"
	"goweb-gin-demo/model/wt"
)

type WtTemplateSearch struct{
	wt.WtReports
	request.PageInfo
}

type WtTemplateRes struct{
	ID uint
	UserName string `json:"userName"`
	Header string `json:"header"`
	Contents []wt.Contents `json:"contents"`
}
