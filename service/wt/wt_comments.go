package wt

import (
	"goweb-gin-demo/global"
	"goweb-gin-demo/model/common/request"
	"goweb-gin-demo/model/wt"
	wtReq "goweb-gin-demo/model/wt/request"
)

type WtCommentService struct {
}

// CreateWtComment 创建WtComment记录
func (wtCommentService *WtCommentService) CreateWtComment(wtComment wt.WtComment) (err error) {
	err = global.GLOBAL_DB.Create(&wtComment).Error
	return err
}

// DeleteWtComment 删除WtComment记录
func (wtCommentService *WtCommentService)DeleteWtComment(wtComment wt.WtComment) (err error) {
	err = global.GLOBAL_DB.Delete(&wtComment).Error
	return err
}

// DeleteWtCommentByIds 批量删除WtComment记录
func (wtCommentService *WtCommentService)DeleteWtCommentByIds(ids request.IdsReq) (err error) {
	err = global.GLOBAL_DB.Delete(&[]wt.WtComment{},"id in ?",ids.Ids).Error
	return err
}

// UpdateWtComment 更新WtComment记录
func (wtCommentService *WtCommentService)UpdateWtComment(wtComment wt.WtComment) (err error) {
	err = global.GLOBAL_DB.Save(&wtComment).Error
	return err
}

// GetWtComment 根据id获取WtComment记录
func (wtCommentService *WtCommentService)GetWtComment(id uint) (err error, wtComment wt.WtComment) {
	err = global.GLOBAL_DB.Where("id = ?", id).First(&wtComment).Error
	return
}

// GetWtCommentInfoList 分页获取WtComment记录
func (wtCommentService *WtCommentService)GetWtCommentInfoList(info wtReq.WtCommentSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GLOBAL_DB.Model(&wt.WtComment{})
    var wtComments []wt.WtComment
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

    if info.ReportId > 0 {
		err = db.Where("report_id=?", info.ReportId).Limit(limit).Offset(offset).Order("created_at desc").Find(&wtComments).Error
	}else {
		err = db.Limit(limit).Offset(offset).Order("created_at desc").Find(&wtComments).Error
	}

	return err, wtComments, total
}
