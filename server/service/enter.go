package service

import (
	"github.com/kuihuar/huasu-admin/server/service/example"
	"github.com/kuihuar/huasu-admin/server/service/system"
)

var ServiceGroupApp = new(ServiceGroup)

type ServiceGroup struct {
	SystemServiceGroup  system.ServiceGroup
	ExampleServiceGroup example.ServiceGroup
}
