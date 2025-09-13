package request

import (
	"github.com/kuihuar/huasu-admin/server/model/common/request"
	"github.com/kuihuar/huasu-admin/server/model/system"
)

type SysDictionaryDetailSearch struct {
	system.SysDictionaryDetail
	request.PageInfo
}
