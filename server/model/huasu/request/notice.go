
package request

import (
	"github.com/kuihuar/huasu-admin/server/model/common/request"
	"time"
)

type NoticeSearch struct{
    CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
    request.PageInfo
}
