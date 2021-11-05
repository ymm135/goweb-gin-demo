package response

import "goweb-gin-demo/global"

type WtRuleResult struct {
	global.GLOBAL_MODEL
	UserId    int
	Reporters []string
	StartTime string
	EndTime   string
}
