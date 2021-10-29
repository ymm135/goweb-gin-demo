package response

import "github.com/ymm135/goweb-gin-demo/model/web"

type SysMenusResponse struct {
	Menus []web.SysMenu `json:"menus"`
}

type SysBaseMenusResponse struct {
	Menus []web.SysBaseMenu `json:"menus"`
}

type SysBaseMenuResponse struct {
	Menu web.SysBaseMenu `json:"menu"`
}
