package wt

import (
	"github.com/gin-gonic/gin"
	"goweb-gin-demo/api"
)

type WtOutputRouter struct {
}

// InitWtOutputRouter 初始化 统计导出 路由信息
func (s *WtOutputRouter) InitWtOutputRouter(Router *gin.RouterGroup) {
	//wtOutputRouter := Router.Group("wtOutput").Use(middleware.OperationRecord())
	wtOutputRouterWithoutRecord := Router.Group("wtOutput")
	var wtOutputApi = api.ApiGroupApp.WtServiceGroup.WtOutputApi
	{
	}
	{
		wtOutputRouterWithoutRecord.GET("GetStatResult", wtOutputApi.GetStatResult)        // 根据用户ID获取统计规则
	}
}
