package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/kuihuar/huasu-admin/server/router"
)

func holder(routers ...*gin.RouterGroup) {
	_ = routers
	_ = router.RouterGroupApp
}
func initBizRouter(routers ...*gin.RouterGroup) {
	privateGroup := routers[0]
	publicGroup := routers[1]
	holder(publicGroup, privateGroup)
	{
		huasuRouter := router.RouterGroupApp.Huasu
		huasuRouter.InitCarouselManageRouter(privateGroup, publicGroup)
		huasuRouter.InitNoticeRouter(privateGroup, publicGroup) // 占位方法，保证文件可以正确加载，避免go空变量检测报错，请勿删除。
		huasuRouter.InitNewsRouter(privateGroup, publicGroup)
	}
}
