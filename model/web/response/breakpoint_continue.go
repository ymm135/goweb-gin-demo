package response

import "goweb-gin-demo/model/web"

type FilePathResponse struct {
	FilePath string `json:"filePath"`
}

type FileResponse struct {
	File web.File `json:"file"`
}
