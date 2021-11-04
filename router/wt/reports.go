package wt

import (
	"github.com/gin-gonic/gin"
	"goweb-gin-demo/api"
	"goweb-gin-demo/middleware"
)

type WtReportsRouter struct {
}


// InitWtReportsRouter 初始化 WtReports 路由信息
func (s *WtReportsRouter) InitWtReportsRouter(Router *gin.RouterGroup) {
	wtReportsRouter := Router.Group("wtReports").Use(middleware.OperationRecord())
	wtReportsRouterWithoutRecord := Router.Group("wtReports")
	var wtReportsApi = api.ApiGroupApp.WtServiceGroup.WtReportsApi
	{
		wtReportsRouter.POST("createWtReports", wtReportsApi.CreateWtReports)   // 新建WtReports
		wtReportsRouter.DELETE("deleteWtReportsByIds", wtReportsApi.DeleteWtReportsByIds) // 批量删除WtReports
		wtReportsRouter.PUT("updateWtReports", wtReportsApi.UpdateWtReports)    // 更新WtReports
	}
	{
		wtReportsRouterWithoutRecord.GET("findWtReports", wtReportsApi.FindWtReports)        // 根据ID获取WtReports
		wtReportsRouterWithoutRecord.GET("getWtReportsList", wtReportsApi.GetWtReportsList)  // 获取WtReports列表
	}
}