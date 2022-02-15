package wt

import "goweb-gin-demo/global"

// WtReports 结构体
type WtReports struct {
	global.GLOBAL_MODEL
	UserName    string `json:"userName" form:"userName" gorm:"column:user_name;comment:用户名;type:varchar(100);"`
	UserId      int    `json:"userId" form:"userId" gorm:"column:user_id;comment:用户id;type:bigint"`
	SendTo      string `json:"sendTo" form:"sendTo" gorm:"column:send_to;comment:发送给谁;type:varchar(100);"`
	Header      string `json:"header" form:"header" gorm:"column:header;comment:报告标题名;type:varchar(100);"`
	Contents    string `json:"contents" form:"contents" gorm:"column:contents;comment:报告内容;type:mediumtext;"`
	Pictures    string `json:"pictures" form:"pictures" gorm:"column:pictures;comment:图片;type:text;"`
	Attachments string `json:"attachments" form:"attachments" gorm:"column:attachments;comment:附件;type:text;"`
}

// TableName WtReports 表名
func (WtReports) TableName() string {
	return "wt_reports"
}

// 报告内容
type Contents struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

// 上传文件
type UploadFileJson struct {
	Key  string `json:"key"`
	Name string `json:"name"`
}

type PageResult struct {
	List     interface{} `json:"list"`
	Ids      []int       `json:"ids"`
	Total    int64       `json:"total"`
	Page     int         `json:"page"`
	PageSize int         `json:"pageSize"`
}
