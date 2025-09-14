package v1

import (
	"github.com/kuihuar/huasu-admin/server/api/v1/example"
	"github.com/kuihuar/huasu-admin/server/api/v1/huasu"
	"github.com/kuihuar/huasu-admin/server/api/v1/system"
)

var ApiGroupApp = new(ApiGroup)

type ApiGroup struct {
	SystemApiGroup  system.ApiGroup
	ExampleApiGroup example.ApiGroup
	HuasuApiGroup   huasu.ApiGroup
}
