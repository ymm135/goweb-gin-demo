package web

import (
	"github.com/gin-gonic/gin"
	"goweb-gin-demo/api"
)

type FileUploadAndDownloadRouter struct {
}

func (e *FileUploadAndDownloadRouter) InitFileUploadAndDownloadRouter(Router *gin.RouterGroup) {
	fileUploadAndDownloadRouter := Router.Group("fileUploadAndDownload")
	var fileUploadAndDownloadApi = api.ApiGroupApp.ApiGroup.FileUploadAndDownloadApi
	{
		fileUploadAndDownloadRouter.POST("upload", fileUploadAndDownloadApi.UploadFile)                                 // 上传文件
		fileUploadAndDownloadRouter.GET("download", fileUploadAndDownloadApi.DownloadFile)                              // 文件下载
		fileUploadAndDownloadRouter.POST("getFileList", fileUploadAndDownloadApi.GetFileList)                           // 获取上传文件列表
		fileUploadAndDownloadRouter.POST("deleteFile", fileUploadAndDownloadApi.DeleteFile)                             // 删除指定文件
		fileUploadAndDownloadRouter.POST("breakpointContinue", fileUploadAndDownloadApi.BreakpointContinue)             // 断点续传
		fileUploadAndDownloadRouter.GET("findFile", fileUploadAndDownloadApi.FindFile)                                  // 查询当前文件成功的切片
		fileUploadAndDownloadRouter.POST("breakpointContinueFinish", fileUploadAndDownloadApi.BreakpointContinueFinish) // 查询当前文件成功的切片
		fileUploadAndDownloadRouter.POST("removeChunk", fileUploadAndDownloadApi.RemoveChunk)                           // 查询当前文件成功的切片
	}
}
