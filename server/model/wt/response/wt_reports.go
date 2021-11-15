package response

import (
	"goweb-gin-demo/global"
	"goweb-gin-demo/model/common"
	"goweb-gin-demo/model/wt"
)

type WtReportsResult struct {
	global.GLOBAL_MODEL
	WtReportInfo
}

type WtReportInfo struct {
	UserName    string              `json:"userName"`
	UserId      int                 `json:"userId"`
	SendTo      []common.UserInfo   `json:"sendTo"`
	Header      string              `json:"header"`
	Contents    []wt.Contents       `json:"contents"`
	Pictures    []wt.UploadFileJson `json:"pictures"`
	Attachments []wt.UploadFileJson `json:"attachments"`
}

type WtReportsSearchBO struct {
	global.GLOBAL_MODEL
	wt.WtReports
	CommentCount uint `json:"commentCount"`
}

type WtReportsSearchResult struct {
	global.GLOBAL_MODEL
	WtReportInfo
	CommentCount uint `json:"commentCount"`
}
