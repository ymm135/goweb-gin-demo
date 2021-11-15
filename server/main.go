package main

import (
	"goweb-gin-demo/core"
	_ "goweb-gin-demo/docs"
	"goweb-gin-demo/global"
	"goweb-gin-demo/initialize"
)

func main() {
	global.GLOBAL_VP = core.Viper()      // 初始化Viper
	global.GLOBAL_LOG = core.Zap()       // 初始化zap日志库
	global.GLOBAL_DB = initialize.Gorm() // gorm连接数据库

	core.RunServer()
}
