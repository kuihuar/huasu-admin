package initialize

import (
	"github.com/kuihuar/huasu-admin/server/global"
	"github.com/kuihuar/huasu-admin/server/middleware"
	"github.com/kuihuar/huasu-admin/server/plugin/announcement/router"
	"github.com/gin-gonic/gin"
)

func Router(engine *gin.Engine) {
	public := engine.Group(global.GVA_CONFIG.System.RouterPrefix).Group("")
	private := engine.Group(global.GVA_CONFIG.System.RouterPrefix).Group("")
	private.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	router.Router.Info.Init(public, private)
}
