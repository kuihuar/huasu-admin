package request

import (
	"github.com/kuihuar/huasu-admin/server/model/common/request"
	"github.com/kuihuar/huasu-admin/server/model/system"
	"time"
)

type SysExportTemplateSearch struct {
	system.SysExportTemplate
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	request.PageInfo
}
