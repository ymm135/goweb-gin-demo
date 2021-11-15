package wt

import (
	"goweb-gin-demo/global"
)

// WtTemplate 结构体
// 如果含有time.Time 请自行import time包
type WtTemplate struct {
      global.GLOBAL_MODEL
      UserName  string `json:"userName" form:"userName" gorm:"column:user_name;comment:用户名;type:varchar(100);"`
      Header  string `json:"header" form:"header" gorm:"column:header;comment:报告标题名;type:varchar(100);"`
      Contents  string `json:"contents" form:"contents" gorm:"column:contents;comment:报告内容;type:mediumtext;"`
}


// TableName WtTemplate 表名
func (WtTemplate) TableName() string {
  return "wt_templates"
}

