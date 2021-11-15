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

type WtTemplateApi struct {
}


// CreateWtTemplate 创建周报模板
// @Tags WtTemplate
// @Summary 创建周报模板
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.WtTemplateRes true "创建周报模板"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wtTemplates/createWtTemplate [post]
func (wtTemplatesApi *WtTemplateApi) CreateWtTemplate(c *gin.Context) {
	var templateRes wtReq.WtTemplateRes
	_ = c.ShouldBindJSON(&templateRes)
	if err := wtTemplatesService.CreateWtTemplate(templateRes); err != nil {
        global.GLOBAL_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteWtTemplateByIds 批量删除WtTemplate
// @Tags WtTemplate
// @Summary 批量删除WtTemplate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除周报模板"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /wtTemplates/deleteWtTemplateByIds [delete]
func (wtTemplatesApi *WtTemplateApi) DeleteWtTemplateByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := wtTemplatesService.DeleteWtTemplateByIds(IDS); err != nil {
        global.GLOBAL_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateWtTemplate 更新周报模板
// @Tags WtTemplate
// @Summary 更新周报模板
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.WtTemplateRes true "更新周报模板"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /wtTemplates/updateWtTemplate [put]
func (wtTemplatesApi *WtTemplateApi) UpdateWtTemplate(c *gin.Context) {
	var templateRes wtReq.WtTemplateRes
	_ = c.ShouldBindJSON(&templateRes)

	if err := wtTemplatesService.UpdateWtTemplate(templateRes); err != nil {
        global.GLOBAL_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindWtTemplate 用id查询周报模板
// @Tags WtTemplate
// @Summary 用id查询周报模板
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.GetById true "用id查询周报模板"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /wtTemplates/findWtTemplate [get]
func (wtTemplatesApi *WtTemplateApi) FindWtTemplate(c *gin.Context) {
	idStr := c.Query("id")
	id, _ := strconv.Atoi(idStr)

	if err, rewtTemplates := wtTemplatesService.GetWtTemplate(uint(id)); err != nil {
        global.GLOBAL_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rewtTemplates": rewtTemplates}, c)
	}
}

// GetWtTemplateList 分页获取周报模板列表
// @Tags WtTemplate
// @Summary 分页获取周报模板列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query wtReq.WtTemplateSearch true "分页获取周报模板列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wtTemplates/getWtTemplateList [get]
func (wtTemplatesApi *WtTemplateApi) GetWtTemplateList(c *gin.Context) {
	var pageInfo wtReq.WtTemplateSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := wtTemplatesService.GetWtTemplateInfoList(pageInfo); err != nil {
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
