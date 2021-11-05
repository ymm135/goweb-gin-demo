package wt

import (
	"github.com/gin-gonic/gin"
	"goweb-gin-demo/api"
	"goweb-gin-demo/middleware"
)

type WtReportsRouter struct {
}


// InitWtReportsRouter 初始化周报路由信息
func (s *WtReportsRouter) InitWtReportsRouter(Router *gin.RouterGroup) {
	wtReportsRouter := Router.Group("wtReports").Use(middleware.OperationRecord())
	wtReportsRouterWithoutRecord := Router.Group("wtReports")
	var wtReportsApi = api.ApiGroupApp.WtServiceGroup.WtReportsApi
	{
		wtReportsRouter.POST("createWtReports", wtReportsApi.CreateWtReports)   // 新建周报
		wtReportsRouter.DELETE("deleteWtReportsByIds", wtReportsApi.DeleteWtReportsByIds) // 批量删除周报
		wtReportsRouter.PUT("updateWtReports", wtReportsApi.UpdateWtReports)    // 更新周报
	}
	{
		wtReportsRouterWithoutRecord.GET("findWtReports", wtReportsApi.FindWtReports)        // 根据ID获取周报
		wtReportsRouterWithoutRecord.GET("getWtReportsList", wtReportsApi.GetWtReportsList)  // 获取周报列表
	}

	wtTemplatesRouter := Router.Group("wtTemplates").Use(middleware.OperationRecord())
	wtTemplatesRouterWithoutRecord := Router.Group("wtTemplates")
	var wtTemplatesApi = api.ApiGroupApp.WtServiceGroup.WtTemplateApi
	{
		wtTemplatesRouter.POST("createWtTemplate", wtTemplatesApi.CreateWtTemplate)   // 新建周报模板
		wtTemplatesRouter.DELETE("deleteWtTemplateByIds", wtTemplatesApi.DeleteWtTemplateByIds) // 批量删除周报模板
		wtTemplatesRouter.PUT("updateWtTemplate", wtTemplatesApi.UpdateWtTemplate)    // 更新周报模板
	}
	{
		wtTemplatesRouterWithoutRecord.GET("findWtTemplate", wtTemplatesApi.FindWtTemplate)        // 根据ID获取周报模板
		wtTemplatesRouterWithoutRecord.GET("getWtTemplateList", wtTemplatesApi.GetWtTemplateList)  // 获取周报模板列表
	}

	wtCommentRouter := Router.Group("wtComment").Use(middleware.OperationRecord())
	wtCommentRouterWithoutRecord := Router.Group("wtComment")
	var wtCommentApi = api.ApiGroupApp.WtServiceGroup.WtCommentApi
	{
		wtCommentRouter.POST("createWtComment", wtCommentApi.CreateWtComment)   // 新建周报评论
		wtCommentRouter.DELETE("deleteWtCommentByIds", wtCommentApi.DeleteWtCommentByIds) // 批量删除周报评论
		wtCommentRouter.PUT("updateWtComment", wtCommentApi.UpdateWtComment)    // 更新周报评论
	}
	{
		wtCommentRouterWithoutRecord.GET("findWtComment", wtCommentApi.FindWtComment)        // 根据ID获取周报评论
		wtCommentRouterWithoutRecord.GET("getWtCommentList", wtCommentApi.GetWtCommentList)  // 获取周报评论列表
	}
}