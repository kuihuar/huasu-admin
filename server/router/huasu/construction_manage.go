package huasu

import (
	"github.com/kuihuar/huasu-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ConstructionManageRouter struct {}

// InitConstructionManageRouter 初始化 工程管理 路由信息
func (s *ConstructionManageRouter) InitConstructionManageRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	constructionManageRouter := Router.Group("constructionManage").Use(middleware.OperationRecord())
	constructionManageRouterWithoutRecord := Router.Group("constructionManage")
	constructionManageRouterWithoutAuth := PublicRouter.Group("constructionManage")
	{
		constructionManageRouter.POST("createConstructionManage", constructionManageApi.CreateConstructionManage)   // 新建工程管理
		constructionManageRouter.DELETE("deleteConstructionManage", constructionManageApi.DeleteConstructionManage) // 删除工程管理
		constructionManageRouter.DELETE("deleteConstructionManageByIds", constructionManageApi.DeleteConstructionManageByIds) // 批量删除工程管理
		constructionManageRouter.PUT("updateConstructionManage", constructionManageApi.UpdateConstructionManage)    // 更新工程管理
	}
	{
		constructionManageRouterWithoutRecord.GET("findConstructionManage", constructionManageApi.FindConstructionManage)        // 根据ID获取工程管理
		constructionManageRouterWithoutRecord.GET("getConstructionManageList", constructionManageApi.GetConstructionManageList)  // 获取工程管理列表
	}
	{
	    constructionManageRouterWithoutAuth.GET("getConstructionManagePublic", constructionManageApi.GetConstructionManagePublic)  // 工程管理开放接口
	}
}
