package request

import (
	"goweb-gin-demo/model/common/request"
	"goweb-gin-demo/model/wt"
)

//query参数要用: form, 而不是json
type WtReportsSearch struct{
	UserId uint `form:"userId"`
	StartTime string `form:"startTime" example:"2021-11-04 12:36:34"`
	EndTime string `form:"endTime"`
	Content string `form:"content" example:"xx项目"`
	request.PageInfo
}

type WtReportsVO struct{
	ID uint
	UserName string `json:"userName"`
	SendTo []wt.UserInfo `json:"sendTo"`
	Header string `json:"header"`
	Contents []wt.Contents `json:"contents"`
	Pictures []wt.UploadFileJson `json:"pictures"`
	Attachments []wt.UploadFileJson `json:"attachments"`
}

