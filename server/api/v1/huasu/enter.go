package huasu

import "github.com/kuihuar/huasu-admin/server/service"

type ApiGroup struct {
	CarouselManageApi
	NoticeApi
}

var (
	CarouselService = service.ServiceGroupApp.HuasuServiceGroup.CarouselManageService
	noticeService   = service.ServiceGroupApp.HuasuServiceGroup.NoticeService
)
