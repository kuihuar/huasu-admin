package router

import (
	"github.com/kuihuar/huasu-admin/server/router/example"
	"github.com/kuihuar/huasu-admin/server/router/huasu"
	"github.com/kuihuar/huasu-admin/server/router/system"
)

var RouterGroupApp = new(RouterGroup)

type RouterGroup struct {
	System  system.RouterGroup
	Example example.RouterGroup
	Huasu   huasu.RouterGroup
}
