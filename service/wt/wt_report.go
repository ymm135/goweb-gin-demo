package wt

import (
	"encoding/json"
	"goweb-gin-demo/global"
	"goweb-gin-demo/model/common/request"
	"goweb-gin-demo/model/wt"
	wtReq "goweb-gin-demo/model/wt/request"
	wtRes "goweb-gin-demo/model/wt/response"
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
func (wtReportsService *WtReportsService) GetWtReportsInfoList(info wtReq.WtReportsSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GLOBAL_DB.Model(&wt.WtReports{})
	var wtReportsList []wt.WtReports
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&wtReportsList).Error

	reportsVOList := reportsToVOs(wtReportsList)

	return err, reportsVOList, total
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
		SendTo:       string(sendToJson),
		Header:       reportsVO.Header,
		Contents:     string(contentJson),
		Pictures:     string(picturesJson),
		Attachments:  string(attachmentsJson),
	}
	return wtReports
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
	reportVO.Header = report.Header
	json.Unmarshal([]byte(report.SendTo), &reportVO.SendTo)
	json.Unmarshal([]byte(report.Contents), &reportVO.Contents)
	json.Unmarshal([]byte(report.Pictures), &reportVO.Pictures)
	json.Unmarshal([]byte(report.Attachments), &reportVO.Attachments)
	return reportVO
}
