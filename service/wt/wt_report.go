package wt

import (
	"encoding/json"
	"goweb-gin-demo/global"
	"goweb-gin-demo/model/common/request"
	"goweb-gin-demo/model/wt"
	wtReq "goweb-gin-demo/model/wt/request"
	wtRes "goweb-gin-demo/model/wt/response"
	"strconv"
)

type WtReportsService struct {
}

// CreateWtReports 创建周报
func (wtReportsService *WtReportsService) CreateWtReports(reportsVO wtReq.WtReportsVO) (err error) {
	wtReports := voToRrports(reportsVO)
	err = global.GLOBAL_DB.Create(&wtReports).Error
	return err
}

// DeleteWtReportsByIds 批量删除周报
func (wtReportsService *WtReportsService) DeleteWtReportsByIds(ids request.IdsReq) (err error) {
	err = global.GLOBAL_DB.Delete(&[]wt.WtReports{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateWtReports 更新周报
func (wtReportsService *WtReportsService) UpdateWtReports(reportsVO wtReq.WtReportsVO) (err error) {
	wtReports := voToRrports(reportsVO)
	err = global.GLOBAL_DB.Updates(&wtReports).Error
	return err
}

// GetWtReports 根据id获取周报
func (wtReportsService *WtReportsService) GetWtReports(id uint) (err error, reportsVO wtRes.WtReportsResult) {
	report := wt.WtReports{}
	err = global.GLOBAL_DB.Where("id = ?", id).First(&report).Error
	reportVO := reportToVO(report)

	return err, reportVO
}

// GetWtReportsInfoList 分页获取周报
func (wtReportsService *WtReportsService) GetWtReportsInfoList(info wtReq.WtReportsSearch) (err error, list []wtRes.WtReportsSearchResult, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)

	reportTable := global.GLOBAL_DB.Table("wt_reports")

	err = reportTable.Count(&total).Error
	if err != nil {
		return
	}

	var reportsSearchBOList []wtRes.WtReportsSearchBO

	if info.Page == 0 {
		limit = int(total)
	}

	//首选获取周报的ids
	var reportIds []uint
	reportTable.Select("`id`", offset, limit).Offset(offset).Limit(limit).Scan(&reportIds)

	querySql := "SELECT id, user_name, send_to, header, contents, pictures, attachments, created_at, updated_at, cmc.comment_count " +
		"FROM wt_reports " +
		"left join (SELECT report_id, count(report_id) as comment_count FROM wt_comments WHERE report_id in ? GROUP BY report_id) as cmc " +
		"on cmc.report_id = wt_reports.id " +
		"WHERE 1=1 "

	// 条件高级查询
	if info.UserId > 0 {
		querySql += " and id = " + strconv.Itoa(int(info.UserId))
	}

	if len(info.Content) != 0 {
		querySql += "  and contents LIKE '%" + info.Content + "%'"
	}

	if len(info.StartTime) != 0 && len(info.EndTime) != 0 {
		querySql += "  and created_at >= '" + info.StartTime + "' and created_at <= '" + info.EndTime + "'"
	}

	querySql += " LIMIT ? OFFSET ? "
	err = global.GLOBAL_DB.Raw(querySql, reportIds, limit, offset).Scan(&reportsSearchBOList).Error

	reportsSearchResultList := reportsToSearchResult(reportsSearchBOList)

	return err, reportsSearchResultList, total
}

//数据转换一下, 需要把json数据转换为字符串
func voToRrports(reportsVO wtReq.WtReportsVO) wt.WtReports {
	sendToJson, _ := json.Marshal(reportsVO.SendTo)
	contentJson, _ := json.Marshal(reportsVO.Contents)
	picturesJson, _ := json.Marshal(reportsVO.Pictures)
	attachmentsJson, _ := json.Marshal(reportsVO.Attachments)

	wtReports := wt.WtReports{
		GLOBAL_MODEL: global.GLOBAL_MODEL{ID: reportsVO.ID},
		UserName:     reportsVO.UserName,
		UserId:       reportsVO.UserId,
		SendTo:       string(sendToJson),
		Header:       reportsVO.Header,
		Contents:     string(contentJson),
		Pictures:     string(picturesJson),
		Attachments:  string(attachmentsJson),
	}
	return wtReports
}

// 批量转换 数据转换, 把字符串转换为json
func reportsToSearchResult(wtReportsList []wtRes.WtReportsSearchBO) []wtRes.WtReportsSearchResult {
	var reportsSearchResults []wtRes.WtReportsSearchResult
	for _, searchBO := range wtReportsList {
		reportVO := reportToSearchResult(searchBO)
		reportsSearchResults = append(reportsSearchResults, reportVO)
	}
	return reportsSearchResults
}

//单个转换
func reportToSearchResult(searchBO wtRes.WtReportsSearchBO) wtRes.WtReportsSearchResult {
	searchResult := wtRes.WtReportsSearchResult{}
	searchResult.GLOBAL_MODEL = searchBO.GLOBAL_MODEL

	searchResult.UserName = searchBO.UserName
	searchResult.UserId = searchBO.UserId
	searchResult.Header = searchBO.Header
	searchResult.CommentCount = searchBO.CommentCount

	json.Unmarshal([]byte(searchBO.SendTo), &searchResult.SendTo)
	json.Unmarshal([]byte(searchBO.Contents), &searchResult.Contents)
	json.Unmarshal([]byte(searchBO.Pictures), &searchResult.Pictures)
	json.Unmarshal([]byte(searchBO.Attachments), &searchResult.Attachments)
	return searchResult
}

// 批量转换 数据转换, 把字符串转换为json
func reportsToVOs(wtReportsList []wt.WtReports) []wtRes.WtReportsResult {
	var reportsVOList []wtRes.WtReportsResult
	for _, report := range wtReportsList {
		reportVO := reportToVO(report)
		reportsVOList = append(reportsVOList, reportVO)
	}
	return reportsVOList
}

//单个转换
func reportToVO(report wt.WtReports) wtRes.WtReportsResult {
	reportVO := wtRes.WtReportsResult{}
	reportVO.GLOBAL_MODEL = report.GLOBAL_MODEL

	reportVO.UserName = report.UserName
	reportVO.UserId = report.UserId
	reportVO.Header = report.Header
	json.Unmarshal([]byte(report.SendTo), &reportVO.SendTo)
	json.Unmarshal([]byte(report.Contents), &reportVO.Contents)
	json.Unmarshal([]byte(report.Pictures), &reportVO.Pictures)
	json.Unmarshal([]byte(report.Attachments), &reportVO.Attachments)
	return reportVO
}

func (wtReportsService *WtReportsService) getWtReportListForStat(statData wt.StatDataSearch) (err error, userIds []int) {
	var commitUserIds []int
	sql := "created_at >= '" + statData.StartTime + "' and created_at <= '" + statData.EndTime + "' and user_id in ? "
	err = global.GLOBAL_DB.Model(&wt.WtReports{}).Select("user_id").Where(sql, statData.UserIds).Scan(&commitUserIds).Error

	return err, commitUserIds
}

func (wtReportsService *WtReportsService) getWtReportListForExcel(statData wt.StatDataSearch) (err error, result []wtRes.WtReportsResult) {
	var wtReportList []wt.WtReports

	condition := " 1=1 "
	if len(statData.StartTime) != 0 {
		condition += "and created_at >= '" + statData.StartTime + "'"
	}

	if len(statData.EndTime) != 0 {
		condition += "and created_at <= '" + statData.EndTime + "'"
	}

	if len(statData.UserIds) != 0 {
		userIdStr := "( "
		for i, id := range statData.UserIds {
			if i < (len(statData.UserIds) - 1) {
				userIdStr += strconv.Itoa(id) + ", "
			} else {
				userIdStr += strconv.Itoa(id)
			}
		}
		userIdStr += " )"
		condition += " and user_id in " + userIdStr
	}

	err = global.GLOBAL_DB.Model(&wt.WtReports{}).Where(condition).Scan(&wtReportList).Error

	reportsResults := reportsToVOs(wtReportList)

	return err, reportsResults
}
