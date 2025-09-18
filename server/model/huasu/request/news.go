package request

import (
	"time"

	"github.com/kuihuar/huasu-admin/server/model/common/request"
)

type NewsSearch struct {
	CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
	request.PageInfo
	Category string `json:"category" form:"category"`
}
