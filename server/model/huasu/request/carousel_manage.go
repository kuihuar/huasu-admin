
package request

import (
	"github.com/kuihuar/huasu-admin/server/model/common/request"
	"time"
)

type CarouselManageSearch struct{
    CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
    request.PageInfo
    Sort  string `json:"sort" form:"sort"`
    Order string `json:"order" form:"order"`
}
