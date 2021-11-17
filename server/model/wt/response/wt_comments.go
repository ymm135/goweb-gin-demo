package response

import (
	"goweb-gin-demo/model/wt"
)

type WtCommentResult struct {
	wt.WtComment
	NickName string `json:"nickName"`
}