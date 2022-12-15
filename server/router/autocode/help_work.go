package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type HelpWorkRouter struct {
}

// InitHelpWorkRouter 初始化 HelpWork 路由信息
func (s *HelpWorkRouter) InitHelpWorkRouter(Router *gin.RouterGroup) {
	helpWorkRouter := Router.Group("helpWork").Use(middleware.OperationRecord())
	helpWorkRouterWithoutRecord := Router.Group("helpWork")
	var helpWorkApi = v1.ApiGroupApp.AutoCodeApiGroup.HelpWorkApi
	{
		helpWorkRouter.POST("createHelpWork", helpWorkApi.CreateHelpWork)   // 新建HelpWork
		helpWorkRouter.DELETE("deleteHelpWork", helpWorkApi.DeleteHelpWork) // 删除HelpWork
		helpWorkRouter.DELETE("deleteHelpWorkByIds", helpWorkApi.DeleteHelpWorkByIds) // 批量删除HelpWork
		helpWorkRouter.PUT("updateHelpWork", helpWorkApi.UpdateHelpWork)    // 更新HelpWork
	}
	{
		helpWorkRouterWithoutRecord.GET("findHelpWork", helpWorkApi.FindHelpWork)        // 根据ID获取HelpWork
		helpWorkRouterWithoutRecord.GET("getHelpWorkList", helpWorkApi.GetHelpWorkList)  // 获取HelpWork列表
	}
}
