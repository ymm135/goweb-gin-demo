package response

import "goweb-gin-demo/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
