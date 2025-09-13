package api

import "github.com/kuihuar/huasu-admin/server/plugin/announcement/service"

var (
	Api         = new(api)
	serviceInfo = service.Service.Info
)

type api struct{ Info info }
