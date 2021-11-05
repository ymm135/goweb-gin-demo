package wt

import (
	"goweb-gin-demo/global"
)

// WtRule 结构体
type WtRule struct {
      global.GLOBAL_MODEL
      UserId  int `json:"userId" form:"userId" gorm:"column:user_id;comment:用户id;type:bigint"`
      Reporters  string `json:"reporters" form:"reporters" gorm:"column:reporters;comment:需要提交报告的人;type:varchar(500);"`
      StartTime  string `json:"startTime" form:"startTime" gorm:"column:start_time;comment:提交报告的起始时间5-0900 代表周五9点;type:varchar(50);"`
      EndTime  string `json:"endTime" form:"endTime" gorm:"column:end_time;comment:提交报告的结束时间8-0900 代表次周一9点截止;type:varchar(50);"`
}


// TableName WtRule 表名
func (WtRule) TableName() string {
  return "wt_rules"
}

