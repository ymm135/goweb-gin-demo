package system

import (
	"errors"
	"mime/multipart"
	"strings"

	"goweb-gin-demo/global"
	"goweb-gin-demo/model/common/request"
	"goweb-gin-demo/model/web"
	"goweb-gin-demo/utils/upload"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: Upload
//@description: 创建文件上传记录
//@param: file model.FileUploadAndDownload
//@return: error

func (e *FileUploadAndDownloadService) Upload(file web.FileUploadAndDownload) error {
	return global.GLOBAL_DB.Create(&file).Error
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: FindFile
//@description: 查找文件
//@param: id uint
//@return: error, model.FileUploadAndDownload

func (e *FileUploadAndDownloadService) FindFile(key string) (error, web.FileUploadAndDownload) {
	var file web.FileUploadAndDownload
	err := global.GLOBAL_DB.Where("`key` = ?", key).First(&file).Error
	return err, file
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteFile
//@description: 删除文件记录
//@param: file model.FileUploadAndDownload
//@return: err error

func (e *FileUploadAndDownloadService) DeleteFile(file web.FileUploadAndDownload) (err error) {
	var fileFromDb web.FileUploadAndDownload
	err, fileFromDb = e.FindFile(file.Key)
	if err != nil {
		return
	}
	oss := upload.NewOss()
	if err = oss.DeleteFile(fileFromDb.Key); err != nil {
		return errors.New("文件删除失败")
	}
	err = global.GLOBAL_DB.Where("id = ?", file.ID).Unscoped().Delete(&file).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetFileRecordInfoList
//@description: 分页获取数据
//@param: info request.PageInfo
//@return: err error, list interface{}, total int64

func (e *FileUploadAndDownloadService) GetFileRecordInfoList(info request.PageInfo) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GLOBAL_DB.Model(&web.FileUploadAndDownload{})
	var fileLists []web.FileUploadAndDownload
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Order("updated_at desc").Find(&fileLists).Error
	return err, fileLists, total
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UploadFile
//@description: 根据配置文件判断是文件上传到本地或者七牛云
//@param: header *multipart.FileHeader, noSave string
//@return: err error, file model.FileUploadAndDownload

func (e *FileUploadAndDownloadService) UploadFile(header *multipart.FileHeader, noSave string) (err error, file web.FileUploadAndDownload) {
	oss := upload.NewOss()
	filePath, key, uploadErr := oss.UploadFile(header)
	if uploadErr != nil {
		panic(err)
	}
	if noSave == "0" {
		s := strings.Split(header.Filename, ".")
		f := web.FileUploadAndDownload{
			Url:  filePath,
			Name: header.Filename,
			Tag:  s[len(s)-1],
			Key:  key,
		}
		return e.Upload(f), f
	}
	return
}
