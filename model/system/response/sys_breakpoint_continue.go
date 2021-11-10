package response

import "goweb-gin-demo/model/system"

type FilePathResponse struct {
	FilePath string `json:"filePath"`
}

type FileResponse struct {
	File system.File `json:"file"`
}
