package request

import (
	"goweb-gin-demo/model/common/request"
	"goweb-gin-demo/model/wt"
)

type WtReportsSearch struct{
	wt.WtReports
	request.PageInfo
}

//
type WtReportsVO struct{
	ID uint
	UserName string `json:"userName"`
	SendTo []wt.UserInfo `json:"sendTo"`
	Header string `json:"header"`
	Contents []wt.Contents `json:"contents"`
	Pictures []wt.UploadFileJson `json:"pictures"`
	Attachments []wt.UploadFileJson `json:"attachments"`
}

