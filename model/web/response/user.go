package response

import (
	"goweb-gin-demo/model/web"
)

type SysUserResponse struct {
	User web.SysUser `json:"user"`
}

type LoginResponse struct {
	User      web.SysUser `json:"user"`
	Token     string         `json:"token"`
	ExpiresAt int64          `json:"expiresAt"`
}
