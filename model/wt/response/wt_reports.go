package response

import (
	"goweb-gin-demo/global"
	"goweb-gin-demo/model/wt"
)


type WtReportsResult struct{
	global.GLOBAL_MODEL
	UserName string `json:"userName"`
	SendTo []wt.UserInfo `json:"sendTo"`
	Header string `json:"header"`
	Contents []wt.Contents `json:"contents"`
	Pictures []wt.UploadFileJson `json:"pictures"`
	Attachments []wt.UploadFileJson `json:"attachments"`
}
