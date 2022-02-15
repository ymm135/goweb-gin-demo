package wt

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"goweb-gin-demo/global"
	"goweb-gin-demo/model/common/request"
	"goweb-gin-demo/model/common/response"
	"goweb-gin-demo/model/wt"
	wtReq "goweb-gin-demo/model/wt/request"
	"strconv"
)

type WtReportsApi struct {
}

// CreateWtReports 创建周报
// @Tags WtReports
// @Summary 创建周报
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body wtReq.WtReportsVO true "创建周报"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wtReports/createWtReports [post]
func (wtReportsApi *WtReportsApi) CreateWtReports(c *gin.Context) {
	var reportsVO wtReq.WtReportsVO
	_ = c.ShouldBindJSON(&reportsVO)

	if err := wtReportsService.CreateWtReports(reportsVO); err != nil {
		global.GLOBAL_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteWtReportsByIds 批量删除周报
// @Tags WtReports
// @Summary 批量删除周报
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除周报"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /wtReports/deleteWtReportsByIds [delete]
func (wtReportsApi *WtReportsApi) DeleteWtReportsByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := wtReportsService.DeleteWtReportsByIds(IDS); err != nil {
		global.GLOBAL_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateWtReports 更新周报
// @Tags WtReports
// @Summary 更新周报
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body wtReq.WtReportsVO true "更新周报"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /wtReports/updateWtReports [put]
func (wtReportsApi *WtReportsApi) UpdateWtReports(c *gin.Context) {
	var reportsVO wtReq.WtReportsVO
	_ = c.ShouldBindJSON(&reportsVO)
	if err := wtReportsService.UpdateWtReports(reportsVO); err != nil {
		global.GLOBAL_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindWtReports 用id查询WtReports
// @Tags WtReports
// @Summary 用id查询WtReports
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.GetById true "用id查询WtReports"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /wtReports/findWtReports [get]
func (wtReportsApi *WtReportsApi) FindWtReports(c *gin.Context) {
	idStr := c.Query("id")
	id, _ := strconv.Atoi(idStr)

	if err, rewtReports := wtReportsService.GetWtReports(uint(id)); err != nil {
		global.GLOBAL_LOG.Error("查询失败,id不存在!", zap.Any("err", err))
		response.FailWithMessage("查询失败,id不存在!", c)
	} else {
		response.OkWithData(gin.H{"rewtReports": rewtReports}, c)
	}
}

// GetWtReportsList 分页获取WtReports列表
// @Tags WtReports
// @Summary 分页获取WtReports列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query wtReq.WtReportsSearch true "分页获取WtReports列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wtReports/getWtReportsList [get]
func (wtReportsApi *WtReportsApi) GetWtReportsList(c *gin.Context) {
	var searchInfo wtReq.WtReportsSearch
	_ = c.ShouldBindQuery(&searchInfo)

	if err, list, total, ids := wtReportsService.GetWtReportsInfoList(searchInfo); err != nil {
		global.GLOBAL_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(wt.PageResult{
			List:     list,
			Ids:      ids,
			Total:    total,
			Page:     searchInfo.Page,
			PageSize: searchInfo.PageSize,
		}, "获取成功", c)
	}
}
