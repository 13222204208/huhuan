package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type HelpGrouponRouter struct {
}

// InitHelpGrouponRouter 初始化 HelpGroupon 路由信息
func (s *HelpGrouponRouter) InitHelpGrouponRouter(Router *gin.RouterGroup) {
	helpGrouponRouter := Router.Group("helpGroupon").Use(middleware.OperationRecord())
	helpGrouponRouterWithoutRecord := Router.Group("helpGroupon")
	var helpGrouponApi = v1.ApiGroupApp.AutoCodeApiGroup.HelpGrouponApi
	{
		helpGrouponRouter.POST("createHelpGroupon", helpGrouponApi.CreateHelpGroupon)   // 新建HelpGroupon
		helpGrouponRouter.DELETE("deleteHelpGroupon", helpGrouponApi.DeleteHelpGroupon) // 删除HelpGroupon
		helpGrouponRouter.DELETE("deleteHelpGrouponByIds", helpGrouponApi.DeleteHelpGrouponByIds) // 批量删除HelpGroupon
		helpGrouponRouter.PUT("updateHelpGroupon", helpGrouponApi.UpdateHelpGroupon)    // 更新HelpGroupon
	}
	{
		helpGrouponRouterWithoutRecord.GET("findHelpGroupon", helpGrouponApi.FindHelpGroupon)        // 根据ID获取HelpGroupon
		helpGrouponRouterWithoutRecord.GET("getHelpGrouponList", helpGrouponApi.GetHelpGrouponList)  // 获取HelpGroupon列表
	}
}
