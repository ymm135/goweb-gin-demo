package response

import "github.com/ymm135/goweb-gin-demo/model/web"

type SysAuthorityResponse struct {
	Authority web.SysAuthority `json:"authority"`
}

type SysAuthorityCopyResponse struct {
	Authority      web.SysAuthority `json:"authority"`
	OldAuthorityId string              `json:"oldAuthorityId"` // 旧角色ID
}
