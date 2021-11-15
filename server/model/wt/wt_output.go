package wt

type StatDataSearch struct {
	StartTime string `json:"startTime" form:"startTime"`
	EndTime   string `json:"endTime" form:"endTime"`
	UserIds   []int  `json:"userIds" form:"userIds"`
}
