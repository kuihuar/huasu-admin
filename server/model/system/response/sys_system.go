package response

import "github.com/kuihuar/huasu-admin/server/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
