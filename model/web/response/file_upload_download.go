package response

import "goweb-gin-demo/model/web"

type ExaFileResponse struct {
	File web.FileUploadAndDownload `json:"file"`
}
