package wt

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"goweb-gin-demo/global"
	"goweb-gin-demo/model/common/request"
	"goweb-gin-demo/model/common/response"
	wtReq "goweb-gin-demo/model/wt/request"
	"strconv"
)

type WtRuleApi struct {
}

// CreateWtRule 创建统计规则
// @Tags WtRule
// @Summary 创建统计规则
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.WtRuleRes true "创建统计规则"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wtRule/createWtRule [post]
func (wtRuleApi *WtRuleApi) CreateWtRule(c *gin.Context) {
	var ruleRes wtReq.WtRuleRes
	_ = c.ShouldBindJSON(&ruleRes)

	if err := wtRuleService.CreateWtRule(ruleRes); err != nil {
		global.GLOBAL_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteWtRuleByIds 批量删除统计规则
// @Tags WtRule
// @Summary 批量删除统计规则
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除统计规则"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /wtRule/deleteWtRuleByIds [delete]
func (wtRuleApi *WtRuleApi) DeleteWtRuleByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := wtRuleService.DeleteWtRuleByIds(IDS); err != nil {
		global.GLOBAL_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateWtRule 更新统计规则
// @Tags WtRule
// @Summary 更新统计规则
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.WtRuleRes true "更新统计规则"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /wtRule/updateWtRule [put]
func (wtRuleApi *WtRuleApi) UpdateWtRule(c *gin.Context) {
	var ruleRes wtReq.WtRuleRes
	_ = c.ShouldBindJSON(&ruleRes)
	if err := wtRuleService.UpdateWtRule(ruleRes); err != nil {
		global.GLOBAL_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindWtRule 用id查询统计规则
// @Tags WtRule
// @Summary 用id查询统计规则
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.GetById true "用id查询统计规则"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /wtRule/findWtRule [get]
func (wtRuleApi *WtRuleApi) FindWtRule(c *gin.Context) {
	idStr := c.Query("id")
	id, _ := strconv.Atoi(idStr)

	if err, rewtRule := wtRuleService.GetWtRule(uint(id)); err != nil {
		global.GLOBAL_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rewtRule": rewtRule}, c)
	}
}

// GetWtRuleList 分页获取统计规则列表
// @Tags WtRule
// @Summary 分页获取统计规则列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query wtReq.WtRuleSearch true "分页获取统计规则列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wtRule/getWtRuleList [get]
func (wtRuleApi *WtRuleApi) GetWtRuleList(c *gin.Context) {
	var pageInfo wtReq.WtRuleSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := wtRuleService.GetWtRuleInfoList(pageInfo); err != nil {
		global.GLOBAL_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
