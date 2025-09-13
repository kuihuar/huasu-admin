package initialize

import (
	_ "github.com/kuihuar/huasu-admin/server/source/example"
	_ "github.com/kuihuar/huasu-admin/server/source/system"
)

func init() {
	// do nothing,only import source package so that inits can be registered
}
