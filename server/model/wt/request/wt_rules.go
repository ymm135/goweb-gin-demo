package request

import (
	"goweb-gin-demo/model/common"
	"goweb-gin-demo/model/common/request"
	"goweb-gin-demo/model/wt"
)

type WtRuleSearch struct {
	wt.WtRule
	request.PageInfo
}

type WtRuleRes struct {
	ID        uint              `json:"id" form:"id"`
	UserId    int               `json:"userId" form:"userId" gorm:"column:user_id;comment:用户id;type:bigint"`
	Reporters []common.UserInfo `json:"reporters" form:"reporters" gorm:"column:reporters;comment:需要提交报告的人;type:varchar(500);"`
	StartWeek int               `json:"startWeek" form:"startWeek" gorm:"column:start_week;comment:提交报告的开始在周几;type:int;"`
	StartHour int               `json:"startHour" form:"startHour" gorm:"column:start_hour;comment:提交周报的开始时间小时;type:int;"`
	EndWeek   int               `json:"endWeek" form:"endWeek" gorm:"column:end_week;comment:提交报告的结束在周几;type:int;"`
	EndHour   int               `json:"endHour" form:"endHour" gorm:"column:end_hour;comment:提交报告的结束时间小时;type:int;"`
}
