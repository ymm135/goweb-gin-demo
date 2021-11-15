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

type WtCommentApi struct {
}

// CreateWtComment 创建周报评论
// @Tags WtComment
// @Summary 创建周报评论
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body wt.WtComment true "创建周报评论"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wtComment/createWtComment [post]
func (wtCommentApi *WtCommentApi) CreateWtComment(c *gin.Context) {
	var wtComment wt.WtComment
	_ = c.ShouldBindJSON(&wtComment)
	if err := wtCommentService.CreateWtComment(wtComment); err != nil {
		global.GLOBAL_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteWtCommentByIds 批量删除周报评论
// @Tags WtComment
// @Summary 批量删除周报评论
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除周报评论"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /wtComment/deleteWtCommentByIds [delete]
func (wtCommentApi *WtCommentApi) DeleteWtCommentByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := wtCommentService.DeleteWtCommentByIds(IDS); err != nil {
		global.GLOBAL_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateWtComment 更新周报评论
// @Tags WtComment
// @Summary 更新周报评论
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body wt.WtComment true "更新周报评论"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /wtComment/updateWtComment [put]
func (wtCommentApi *WtCommentApi) UpdateWtComment(c *gin.Context) {
	var wtComment wt.WtComment
	_ = c.ShouldBindJSON(&wtComment)
	if err := wtCommentService.UpdateWtComment(wtComment); err != nil {
		global.GLOBAL_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindWtComment 用id查询周报评论
// @Tags WtComment
// @Summary 用id查询周报评论
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.GetById true "用id查询周报评论"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /wtComment/findWtComment [get]
func (wtCommentApi *WtCommentApi) FindWtComment(c *gin.Context) {
	idStr := c.Query("id")
	id, _ := strconv.Atoi(idStr)

	if err, rewtComment := wtCommentService.GetWtComment(uint(id)); err != nil {
		global.GLOBAL_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rewtComment": rewtComment}, c)
	}
}

// GetWtCommentList 分页获取周报评论列表
// @Tags WtComment
// @Summary 分页获取周报评论列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query wtReq.WtCommentSearch true "分页获取周报评论列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wtComment/getWtCommentList [get]
func (wtCommentApi *WtCommentApi) GetWtCommentList(c *gin.Context) {
	var pageInfo wtReq.WtCommentSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := wtCommentService.GetWtCommentInfoList(pageInfo); err != nil {
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
