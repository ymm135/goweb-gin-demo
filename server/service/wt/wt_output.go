package wt

import (
	"errors"
	"fmt"
	"github.com/xuri/excelize/v2"
	"goweb-gin-demo/model/common"
	"goweb-gin-demo/model/common/request"
	"goweb-gin-demo/model/wt"
	wtReq "goweb-gin-demo/model/wt/request"
	wtRes "goweb-gin-demo/model/wt/response"
	"goweb-gin-demo/utils"
	"jaytaylor.com/html2text"
)

type WtOutputService struct {
}

// GetWtRule 根据id获取WtRule记录
func (wtOutputService *WtOutputService) GetStatResult(idInfo request.GetByUserID) (err error, wtRule wtRes.StatResult) {
	//首先查询要统计的人

	var WtServiceGroup WtServiceGroup
	err, ruleResult := WtServiceGroup.WtRuleService.GetWtRuleByUserId(idInfo.UserId)
	if err != nil {
		return err, wtRes.StatResult{}
	}


	//计算起始时间 起始时间是 5-0900 7-1000

	startTime := utils.GetTimeFromWeek(ruleResult.StartWeek, ruleResult.StartHour)
	endTime := utils.GetTimeFromWeek(ruleResult.EndWeek, ruleResult.EndHour)

	//在周报表中进行统计
	reports := ruleResult.Reporters
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
		} else {
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

func (wtOutputService *WtOutputService) ExportReportToExcel(info wt.StatDataSearch, excelPath string) (err error) {
	var WtServiceGroup WtServiceGroup
	err, reportResultList := WtServiceGroup.WtReportsService.getWtReportListForExcel(info)

	if err != nil {
		return err
	}

	if len(reportResultList) == 0 {
		return errors.New("没有任何数据可以导出!")
	}

	excel := excelize.NewFile()

	var titles []string
	titles = append(titles, "序号")
	titles = append(titles, "标题")
	titles = append(titles, "用户名")

	for _, content := range reportResultList[0].Contents {
		titles = append(titles, content.Title)
	}

	titles = append(titles, "创建时间")
	titles = append(titles, "更新时间")

	//sheetName := reportResultList[0].Header

	excel.SetSheetRow("Sheet1", "A1", &titles)

	for i, report := range reportResultList {
		axis := fmt.Sprintf("A%d", i+2)

		var excelContent []interface{}
		excelContent = append(excelContent, i + 1)
		excelContent = append(excelContent, report.Header)
		excelContent = append(excelContent, report.UserName)

		for _, content := range report.Contents {
			fromString, err := html2text.FromString(content.Content,  html2text.Options{TextOnly: true})
			if err != nil {
				return err
			}

			excelContent = append(excelContent, fromString)
		}

		excelContent = append(excelContent, report.CreatedAt.String())
		excelContent = append(excelContent, report.UpdatedAt.String())

		excel.SetSheetRow("Sheet1", axis, &excelContent)
	}
	err = excel.SaveAs(excelPath)

	return err
}
