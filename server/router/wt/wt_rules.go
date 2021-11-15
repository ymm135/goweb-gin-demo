package wt

import (
	"github.com/gin-gonic/gin"
	"goweb-gin-demo/api"
	"goweb-gin-demo/middleware"
)

type WtRuleRouter struct {
}

// InitWtRuleRouter 初始化 统计规则 路由信息
func (s *WtRuleRouter) InitWtRuleRouter(Router *gin.RouterGroup) {
	wtRuleRouter := Router.Group("wtRule").Use(middleware.OperationRecord())
	wtRuleRouterWithoutRecord := Router.Group("wtRule")
	var wtRuleApi = api.ApiGroupApp.WtServiceGroup.WtRuleApi
	{
		wtRuleRouter.POST("createWtRule", wtRuleApi.CreateWtRule)   // 新建统计规则
		wtRuleRouter.DELETE("deleteWtRuleByIds", wtRuleApi.DeleteWtRuleByIds) // 批量删除统计规则
		wtRuleRouter.PUT("updateWtRule", wtRuleApi.UpdateWtRule)    // 更新统计规则
	}
	{
		wtRuleRouterWithoutRecord.GET("findWtRule", wtRuleApi.FindWtRule)        // 根据ID获取统计规则
		wtRuleRouterWithoutRecord.GET("getWtRuleList", wtRuleApi.GetWtRuleList)  // 获取统计规则列表
	}
}
