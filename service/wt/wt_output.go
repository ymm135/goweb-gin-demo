package wt

import (
	"goweb-gin-demo/model/common"
	"goweb-gin-demo/model/common/request"
	"goweb-gin-demo/model/wt"
	wtReq "goweb-gin-demo/model/wt/request"
	wtRes "goweb-gin-demo/model/wt/response"
	"goweb-gin-demo/utils"
)

type WtOutputService struct {
}


// GetWtRule 根据id获取WtRule记录
func (wtOutputService *WtOutputService) GetStatResult(idInfo request.GetByUserID) (err error, wtRule wtRes.StatResult) {
	//首先查询要统计的人
	ruleInfo := wtReq.WtRuleSearch{}
	ruleInfo.UserId = int(idInfo.UserId)
	var WtServiceGroup WtServiceGroup
	err, wtRuleResultList, _ := WtServiceGroup.WtRuleService.GetWtRuleInfoList(ruleInfo)
	if err != nil {
		return err, wtRes.StatResult{}
	}

	wtRuleResult := wtRuleResultList[0]

	//计算起始时间 起始时间是 5-0900 7-1000

	startTime := utils.GetTimeFromWeek(wtRuleResult.StartWeek, wtRuleResult.StartHour)
	endTime := utils.GetTimeFromWeek(wtRuleResult.EndWeek, wtRuleResult.EndHour)

	//在周报表中进行统计
	reports := wtRuleResult.Reporters
	var reportSearch wtReq.WtReportsSearch
	reportSearch.StartTime = startTime
	reportSearch.EndTime = endTime

	var userIds []int

	for _, userInfo := range reports {
		userIds = append(userIds, int(userInfo.ID))
	}

	var statData wt.StatDataSearch
	statData.StartTime = startTime
	statData.EndTime = endTime
	statData.UserIds = userIds

	err, commitedUserIds := WtServiceGroup.WtReportsService.getWtReportListForStat(statData)
	if err != nil {

	}
	var commitedList []common.UserInfo
	var uncommitedList []common.UserInfo

	// 比较那些人提交了，那些人没有提交
	for _, needCommitPeoples := range reports {
		isCommit := false
		for _, commitedUserId := range commitedUserIds {
			if int(needCommitPeoples.ID) == commitedUserId {
				isCommit = true
			}
		}

		if isCommit {
			commitedList = append(commitedList, needCommitPeoples)
		}else {
			uncommitedList = append(uncommitedList, needCommitPeoples)
		}
	}

	var statResult wtRes.StatResult
	statResult.CommitCount = len(commitedList)
	statResult.CommitPeoples = commitedList
	statResult.UncommitCount = len(uncommitedList)
	statResult.UncommitPeoples = uncommitedList

	return err, statResult
}
