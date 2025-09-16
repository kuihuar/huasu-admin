package huasu

import (
	"github.com/kuihuar/huasu-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type NewsRouter struct {}

// InitNewsRouter 初始化 新闻版块 路由信息
func (s *NewsRouter) InitNewsRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	newsRouter := Router.Group("news").Use(middleware.OperationRecord())
	newsRouterWithoutRecord := Router.Group("news")
	newsRouterWithoutAuth := PublicRouter.Group("news")
	{
		newsRouter.POST("createNews", newsApi.CreateNews)   // 新建新闻版块
		newsRouter.DELETE("deleteNews", newsApi.DeleteNews) // 删除新闻版块
		newsRouter.DELETE("deleteNewsByIds", newsApi.DeleteNewsByIds) // 批量删除新闻版块
		newsRouter.PUT("updateNews", newsApi.UpdateNews)    // 更新新闻版块
	}
	{
		newsRouterWithoutRecord.GET("findNews", newsApi.FindNews)        // 根据ID获取新闻版块
		newsRouterWithoutRecord.GET("getNewsList", newsApi.GetNewsList)  // 获取新闻版块列表
	}
	{
	    newsRouterWithoutAuth.GET("getNewsPublic", newsApi.GetNewsPublic)  // 新闻版块开放接口
	}
}
