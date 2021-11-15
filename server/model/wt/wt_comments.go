package wt

import (
	"goweb-gin-demo/global"
)

// WtComment 结构体
type WtComment struct {
      global.GLOBAL_MODEL
      ReportId  int `json:"reportId" form:"reportId" gorm:"column:report_id;comment:报告id;type:bigint"`
      UserName  string `json:"userName" form:"userName" gorm:"column:user_name;comment:评论的用户;type:varchar(100);"`
      Comment  string `json:"comment" form:"comment" gorm:"column:comment;comment:评论;type:text;"`
}


// TableName WtComment 表名
func (WtComment) TableName() string {
  return "wt_comments"
}

