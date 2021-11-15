package response

import "goweb-gin-demo/model/system"

type ExaFileResponse struct {
	File system.FileUploadAndDownload `json:"file"`
}
