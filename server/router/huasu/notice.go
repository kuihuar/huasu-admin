package huasu

import (
	"github.com/kuihuar/huasu-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type NoticeRouter struct {}

// InitNoticeRouter 初始化 公告管理 路由信息
func (s *NoticeRouter) InitNoticeRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	noticeRouter := Router.Group("notice").Use(middleware.OperationRecord())
	noticeRouterWithoutRecord := Router.Group("notice")
	noticeRouterWithoutAuth := PublicRouter.Group("notice")
	{
		noticeRouter.POST("createNotice", noticeApi.CreateNotice)   // 新建公告管理
		noticeRouter.DELETE("deleteNotice", noticeApi.DeleteNotice) // 删除公告管理
		noticeRouter.DELETE("deleteNoticeByIds", noticeApi.DeleteNoticeByIds) // 批量删除公告管理
		noticeRouter.PUT("updateNotice", noticeApi.UpdateNotice)    // 更新公告管理
	}
	{
		noticeRouterWithoutRecord.GET("findNotice", noticeApi.FindNotice)        // 根据ID获取公告管理
		noticeRouterWithoutRecord.GET("getNoticeList", noticeApi.GetNoticeList)  // 获取公告管理列表
	}
	{
	    noticeRouterWithoutAuth.GET("getNoticePublic", noticeApi.GetNoticePublic)  // 公告管理开放接口
	}
}
