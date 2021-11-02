package web

import (
	"goweb-gin-demo/global"
)

// file struct, 文件结构体
type File struct {
	global.GLOBAL_MODEL
	FileName     string
	FileMd5      string
	FilePath     string
	ExaFileChunk []FileChunk
	ChunkTotal   int
	IsFinish     bool
}

// file chunk struct, 切片结构体
type FileChunk struct {
	global.GLOBAL_MODEL
	ExaFileID       uint
	FileChunkNumber int
	FileChunkPath   string
}
