package huasu

import (
	"github.com/kuihuar/huasu-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CarouselManageRouter struct {}

// InitCarouselManageRouter 初始化 轮播图管理 路由信息
func (s *CarouselManageRouter) InitCarouselManageRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	CarouselRouter := Router.Group("Carousel").Use(middleware.OperationRecord())
	CarouselRouterWithoutRecord := Router.Group("Carousel")
	CarouselRouterWithoutAuth := PublicRouter.Group("Carousel")
	{
		CarouselRouter.POST("createCarouselManage", CarouselApi.CreateCarouselManage)   // 新建轮播图管理
		CarouselRouter.DELETE("deleteCarouselManage", CarouselApi.DeleteCarouselManage) // 删除轮播图管理
		CarouselRouter.DELETE("deleteCarouselManageByIds", CarouselApi.DeleteCarouselManageByIds) // 批量删除轮播图管理
		CarouselRouter.PUT("updateCarouselManage", CarouselApi.UpdateCarouselManage)    // 更新轮播图管理
	}
	{
		CarouselRouterWithoutRecord.GET("findCarouselManage", CarouselApi.FindCarouselManage)        // 根据ID获取轮播图管理
		CarouselRouterWithoutRecord.GET("getCarouselManageList", CarouselApi.GetCarouselManageList)  // 获取轮播图管理列表
	}
	{
	    CarouselRouterWithoutAuth.GET("getCarouselManagePublic", CarouselApi.GetCarouselManagePublic)  // 轮播图管理开放接口
	}
}
