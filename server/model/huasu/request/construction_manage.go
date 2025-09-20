
package request

import (
	"github.com/kuihuar/huasu-admin/server/model/common/request"
	"time"
)

type ConstructionManageSearch struct{
    CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
    request.PageInfo
}
