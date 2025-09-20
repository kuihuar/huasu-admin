package huasu

import "github.com/kuihuar/huasu-admin/server/service"

type ApiGroup struct {
	CarouselManageApi
	NoticeApi
	NewsApi
	ConstructionManageApi
}

var (
	CarouselService           = service.ServiceGroupApp.HuasuServiceGroup.CarouselManageService
	noticeService             = service.ServiceGroupApp.HuasuServiceGroup.NoticeService
	newsService               = service.ServiceGroupApp.HuasuServiceGroup.NewsService
	constructionManageService = service.ServiceGroupApp.HuasuServiceGroup.ConstructionManageService
)
