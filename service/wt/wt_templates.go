package wt

import (
	"encoding/json"
	"goweb-gin-demo/global"
	"goweb-gin-demo/model/common/request"
	"goweb-gin-demo/model/wt"
	wtReq "goweb-gin-demo/model/wt/request"
	wtRes "goweb-gin-demo/model/wt/response"
)

type WtTemplateService struct {
}

// CreateWtTemplate 创建周报模板
func (wtTemplatesService *WtTemplateService) CreateWtTemplate(templateRes wtReq.WtTemplateRes) (err error) {
	template := voToTemplate(templateRes)
	err = global.GLOBAL_DB.Create(&template).Error
	return err
}


// DeleteWtTemplateByIds 批量删除周报模板
func (wtTemplatesService *WtTemplateService) DeleteWtTemplateByIds(ids request.IdsReq) (err error) {
	err = global.GLOBAL_DB.Delete(&[]wt.WtTemplate{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateWtTemplate 更新周报模板
func (wtTemplatesService *WtTemplateService) UpdateWtTemplate(templateRes wtReq.WtTemplateRes) (err error) {
	template := voToTemplate(templateRes)
	err = global.GLOBAL_DB.Updates(&template).Error
	return err
}

// GetWtTemplate 根据id获取周报模板
func (wtTemplatesService *WtTemplateService) GetWtTemplate(id uint) (err error, templateResult wtRes.WtTemplateResult) {
	var wtTemplates wt.WtTemplate
	err = global.GLOBAL_DB.Where("id = ?", id).First(&wtTemplates).Error
	result := templateToResult(wtTemplates)
	return err, result
}

// GetWtTemplateInfoList 分页获取周报模板
func (wtTemplatesService *WtTemplateService) GetWtTemplateInfoList(info wtReq.WtTemplateSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GLOBAL_DB.Model(&wt.WtTemplate{})
	var wtTemplatess []wt.WtTemplate
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&wtTemplatess).Error

	templateVOList := templateToVOs(wtTemplatess)

	return err, templateVOList, total
}

//数据转换一下, 需要把json数据转换为字符串
func voToTemplate(templateVO wtReq.WtTemplateRes) wt.WtTemplate {
	contentJson, _ := json.Marshal(templateVO.Contents)

	wtTemplate := wt.WtTemplate{
		GLOBAL_MODEL: global.GLOBAL_MODEL{ID: templateVO.ID},
		UserName:     templateVO.UserName,
		Header:       templateVO.Header,
		Contents:     string(contentJson),
	}
	return wtTemplate
}

// 批量转换 数据转换, 把字符串转换为json
func templateToVOs(templates []wt.WtTemplate) []wtRes.WtTemplateResult {
	var templateListResult []wtRes.WtTemplateResult
	for _, template := range templates {
		templateRes := templateToResult(template)
		templateListResult = append(templateListResult, templateRes)
	}
	return templateListResult
}

//单个转换
func templateToResult(template wt.WtTemplate) wtRes.WtTemplateResult {
	wtTemplateResult := wtRes.WtTemplateResult{}
	wtTemplateResult.GLOBAL_MODEL = template.GLOBAL_MODEL

	wtTemplateResult.UserName = template.UserName
	wtTemplateResult.Header = template.Header
	json.Unmarshal([]byte(template.Contents), &wtTemplateResult.Contents)
	return wtTemplateResult
}
