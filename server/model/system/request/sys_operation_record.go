package request

import (
	"github.com/kuihuar/huasu-admin/server/model/common/request"
	"github.com/kuihuar/huasu-admin/server/model/system"
)

type SysOperationRecordSearch struct {
	system.SysOperationRecord
	request.PageInfo
}
