package huasu

import api "github.com/kuihuar/huasu-admin/server/api/v1"

type RouterGroup struct {
	CarouselManageRouter
	NoticeRouter
	NewsRouter
}

var (
	CarouselApi = api.ApiGroupApp.HuasuApiGroup.CarouselManageApi
	noticeApi   = api.ApiGroupApp.HuasuApiGroup.NoticeApi
	newsApi     = api.ApiGroupApp.HuasuApiGroup.NewsApi
)
