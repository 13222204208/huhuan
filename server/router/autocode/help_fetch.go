package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type HelpFetchRouter struct {
}

// InitHelpFetchRouter 初始化 HelpFetch 路由信息
func (s *HelpFetchRouter) InitHelpFetchRouter(Router *gin.RouterGroup) {
	helpFetchRouter := Router.Group("helpFetch").Use(middleware.OperationRecord())
	helpFetchRouterWithoutRecord := Router.Group("helpFetch")
	var helpFetchApi = v1.ApiGroupApp.AutoCodeApiGroup.HelpFetchApi
	{
		helpFetchRouter.POST("createHelpFetch", helpFetchApi.CreateHelpFetch)   // 新建HelpFetch
		helpFetchRouter.DELETE("deleteHelpFetch", helpFetchApi.DeleteHelpFetch) // 删除HelpFetch
		helpFetchRouter.DELETE("deleteHelpFetchByIds", helpFetchApi.DeleteHelpFetchByIds) // 批量删除HelpFetch
		helpFetchRouter.PUT("updateHelpFetch", helpFetchApi.UpdateHelpFetch)    // 更新HelpFetch
	}
	{
		helpFetchRouterWithoutRecord.GET("findHelpFetch", helpFetchApi.FindHelpFetch)        // 根据ID获取HelpFetch
		helpFetchRouterWithoutRecord.GET("getHelpFetchList", helpFetchApi.GetHelpFetchList)  // 获取HelpFetch列表
	}
}
