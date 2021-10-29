package request

import (
	"github.com/ymm135/goweb-gin-demo/global"
	"github.com/ymm135/goweb-gin-demo/model/web"
)

// Add menu authority info structure
type AddMenuAuthorityInfo struct {
	Menus       []web.SysBaseMenu `json:"menus"`
	AuthorityId string               `json:"authorityId"` // 角色ID
}

func DefaultMenu() []web.SysBaseMenu {
	return []web.SysBaseMenu{{
		GLOBAL_MODEL: global.GLOBAL_MODEL{ID: 1},
		ParentId:  "0",
		Path:      "dashboard",
		Name:      "dashboard",
		Component: "view/dashboard/index.vue",
		Sort:      1,
		Meta: web.Meta{
			Title: "仪表盘",
			Icon:  "setting",
		},
	}}
}
