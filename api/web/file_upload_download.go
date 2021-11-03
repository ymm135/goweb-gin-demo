package web

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"goweb-gin-demo/global"
	"goweb-gin-demo/model/common/request"
	"goweb-gin-demo/model/common/response"
	"goweb-gin-demo/model/web"
	webRes "goweb-gin-demo/model/web/response"
	"net/http"
)

type FileUploadAndDownloadApi struct {
}

// @Tags FileUploadAndDownload
// @Summary 上传文件
// @Security ApiKeyAuth
// @accept multipart/form-data
// @Produce  application/json
// @Param file formData file true "上传文件示例"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"上传成功"}"
// @Router /fileUploadAndDownload/upload [post]
func (u *FileUploadAndDownloadApi) UploadFile(c *gin.Context) {
	var file web.FileUploadAndDownload
	noSave := c.DefaultQuery("noSave", "0")
	_, header, err := c.Request.FormFile("file")
	if err != nil {
		global.GLOBAL_LOG.Error("接收文件失败!", zap.Any("err", err))
		response.FailWithMessage("接收文件失败", c)
		return
	}
	err, file = fileUploadAndDownloadService.UploadFile(header, noSave) // 文件上传后拿到文件路径
	if err != nil {
		global.GLOBAL_LOG.Error("修改数据库链接失败!", zap.Any("err", err))
		response.FailWithMessage("修改数据库链接失败", c)
		return
	}
	response.OkWithDetailed(webRes.ExaFileResponse{File: file}, "上传成功", c)
}

// @Tags FileUploadAndDownload
// @Summary 文件下载
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param fileName query string true "待下载的文件名"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"文件下载成功"}"
// @Router /fileUploadAndDownload/download [get]
func (u *FileUploadAndDownloadApi) DownloadFile(c *gin.Context) {
	fileName := c.Query("fileName")

	//待下载的文件是存储本地的,如果存储在云盘，直接访问链接即可
	//如果没有记录，也会报错，record not found
	err, fileInfo := fileUploadAndDownloadService.FindFile(fileName)

	if err != nil {
		global.GLOBAL_LOG.Error("文件未找到!", zap.Any("err", err))
		c.Redirect(http.StatusFound, "/404")
		return
	}

	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename="+fileInfo.Name)
	c.Header("Content-Transfer-Encoding", "binary")

	c.File(fileInfo.Url)
	return
}

// @Tags FileUploadAndDownload
// @Summary 删除文件
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body web.FileUploadAndDownload true "传入文件里面id即可"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /fileUploadAndDownload/deleteFile [post]
func (u *FileUploadAndDownloadApi) DeleteFile(c *gin.Context) {
	var file web.FileUploadAndDownload
	_ = c.ShouldBindJSON(&file)
	if err := fileUploadAndDownloadService.DeleteFile(file); err != nil {
		global.GLOBAL_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// @Tags FileUploadAndDownload
// @Summary 分页文件列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "页码, 每页大小"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /fileUploadAndDownload/getFileList [post]
func (u *FileUploadAndDownloadApi) GetFileList(c *gin.Context) {
	var pageInfo request.PageInfo
	_ = c.ShouldBindJSON(&pageInfo)
	err, list, total := fileUploadAndDownloadService.GetFileRecordInfoList(pageInfo)
	if err != nil {
		global.GLOBAL_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
