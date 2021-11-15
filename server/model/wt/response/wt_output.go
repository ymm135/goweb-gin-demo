package response

import "goweb-gin-demo/model/common"

// 周报统计结果
type StatResult struct {
	CommitCount int `json:"commitCount"`
	UncommitCount int `json:"uncommitCount"`
	CommitPeoples []common.UserInfo `json:"commitPeoples"`
	UncommitPeoples []common.UserInfo `json:"uncommitPeoples"`
}
