package wt

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"goweb-gin-demo/global"
	"goweb-gin-demo/model/common/request"
	"goweb-gin-demo/model/common/response"
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
		global.GLOBAL_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rewtOutput": rewtOutput}, c)
	}
}
