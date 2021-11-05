package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"goweb-gin-demo/global"
	"goweb-gin-demo/middleware"
	"goweb-gin-demo/router"
)

func Routers() *gin.Engine {
	var Router = gin.Default()

	// 为用户头像和文件提供静态地址
	//Router.StaticFS()

	//排除Swagger路径
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	RouterGroup := router.RouterGroupApp

	PublicGroup := Router.Group("")
	{
		// 健康监测
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(200, "ok")
		})
	}
	{
		RouterGroup.InitBaseRouter(PublicGroup)
	}

	PrivateGroup := Router.Group("")
	//路由要经过jwt和cas校验
	PrivateGroup.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		RouterGroup.InitJwtRouter(PrivateGroup)
		RouterGroup.InitUserRouter(PrivateGroup)
		RouterGroup.InitMenuRouter(PrivateGroup)
		RouterGroup.InitSystemRouter(PrivateGroup)
		RouterGroup.InitFileUploadAndDownloadRouter(PrivateGroup)
		RouterGroup.InitAuthorityRouter(PrivateGroup)
		RouterGroup.InitWtReportsRouter(PrivateGroup)
		RouterGroup.InitWtRuleRouter(PrivateGroup)
	}

	global.GLOBAL_LOG.Info("router register success")

	return Router
}
