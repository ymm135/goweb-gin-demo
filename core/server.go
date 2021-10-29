package core

import (
	"fmt"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/ymm135/goweb-gin-demo/global"
	"github.com/ymm135/goweb-gin-demo/initialize"
	"go.uber.org/zap"
	"time"
)

type server interface {
	ListenAndServe() error
}

func RunServer() {

	// 禁用控制台颜色
	gin.DisableConsoleColor()

	router := initialize.Routers()

	router.Static("/form-generator", "./resource/page")

	address := fmt.Sprintf(":%d", global.GLOBAL_CONFIG.System.Addr)
	s := initServer(address, router)
	// 保证文本顺序输出
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)
	global.GLOBAL_LOG.Info("server run success on ", zap.String("address", address))

	fmt.Printf(`
	当前版本:V2.4.5
	默认自动化文档地址:http://127.0.0.1%s/swagger/index.html
	默认前端文件运行地址:http://127.0.0.1:8080
`, address)
	global.GLOBAL_LOG.Error(s.ListenAndServe().Error())
}

func initServer(address string, router *gin.Engine) server {
	s := endless.NewServer(address, router)
	s.ReadHeaderTimeout = 10 * time.Millisecond
	s.WriteTimeout = 10 * time.Second
	s.MaxHeaderBytes = 1 << 20
	return s
}
