package response

import (
	"goweb-gin-demo/global"
	"goweb-gin-demo/model/wt"
)


type WtTemplateResult struct{
	global.GLOBAL_MODEL
	UserName string `json:"userName"`
	Header string `json:"header"`
	Contents []wt.Contents `json:"contents"`
}
