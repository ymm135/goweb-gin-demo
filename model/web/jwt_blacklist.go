package web

import (
	"github.com/ymm135/goweb-gin-demo/global"
)

type JwtBlacklist struct {
	global.GLOBAL_MODEL
	Jwt string `gorm:"type:text;comment:jwt"`
}
