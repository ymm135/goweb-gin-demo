package middleware

import (
	"github.com/gin-gonic/gin"
	"goweb-gin-demo/global"
	"goweb-gin-demo/model/common/response"
	"goweb-gin-demo/model/system/request"
	"goweb-gin-demo/service"
)

var casbinService = service.ServiceGroupApp.SystemServiceGroup.CasbinService

// 拦截器
func CasbinHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, _ := c.Get("claims")
		waitUse := claims.(*request.CustomClaims)
		// 获取请求的URI
		obj := c.Request.URL.RequestURI()
		// 获取请求方法
		act := c.Request.Method
		// 获取用户的角色
		sub := waitUse.AuthorityId
		e := casbinService.Casbin()
		// 判断策略中是否存在, 查表匹配，比如: sub:100 ,obt:"/menu/getMenu", act:"Post"
		success, _ := e.Enforce(sub, obj, act)
		if global.GLOBAL_CONFIG.System.Env == "develop" || success {
			c.Next()
		} else {
			response.FailWithDetailed(gin.H{}, "权限不足", c)
			c.Abort()
			return
		}
	}
}
