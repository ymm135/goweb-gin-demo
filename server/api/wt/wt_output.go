package wt

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"goweb-gin-demo/global"
	"goweb-gin-demo/model/common/request"
	"goweb-gin-demo/model/common/response"
	"goweb-gin-demo/model/wt"
	"goweb-gin-demo/utils"
)

type WtOutputApi struct {

}

// GetStatResult 根据用户id查询统计结果
// @Tags WtOutput
// @Summary 根据用户id查询统计结果
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.GetByUserID true "根据用户id查询统计结果"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /wtOutput/GetStatResult [get]
func (wtRuleApi *WtOutputApi) GetStatResult(c *gin.Context) {
	var idInfo request.GetByUserID
	c.ShouldBindQuery(&idInfo)

	if err, rewtOutput := wtOutputService.GetStatResult(idInfo); err != nil {
		global.GLOBAL_LOG.Error("统计失败,可能需要先创建统计规则!", zap.Any("err", err))
		response.FailWithMessage("统计失败,可能需要先创建统计规则!", c)
	} else {
		response.OkWithData(gin.H{"rewtOutput": rewtOutput}, c)
	}
}


// GetStatResult 把周报导出为excel
// @Tags WtOutput
// @Summary 把周报导出为excel
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query wt.StatDataSearch true "把周报导出为excel"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"导出成功"}"
// @Router /wtOutput/ExportReportToExcel [get]
func (wtRuleApi *WtOutputApi) ExportReportToExcel(c *gin.Context) {
	var searchInfo wt.StatDataSearch
	c.ShouldBindJSON(&searchInfo)

	filePath := global.GLOBAL_CONFIG.Excel.Dir + utils.GetExcelFileName()

	if err := wtOutputService.ExportReportToExcel(searchInfo, filePath); err != nil {
		global.GLOBAL_LOG.Error("导出失败!", zap.Any("err", err))
		response.FailWithMessage(err.Error(), c)
	}

	c.Writer.Header().Add("success", "true")
	c.File(filePath)
}
