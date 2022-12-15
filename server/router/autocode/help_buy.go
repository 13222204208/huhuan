package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type HelpBuyRouter struct {
}

// InitHelpBuyRouter 初始化 HelpBuy 路由信息
func (s *HelpBuyRouter) InitHelpBuyRouter(Router *gin.RouterGroup) {
	helpBuyRouter := Router.Group("helpBuy").Use(middleware.OperationRecord())
	helpBuyRouterWithoutRecord := Router.Group("helpBuy")
	var helpBuyApi = v1.ApiGroupApp.AutoCodeApiGroup.HelpBuyApi
	{
		helpBuyRouter.POST("createHelpBuy", helpBuyApi.CreateHelpBuy)   // 新建HelpBuy
		helpBuyRouter.DELETE("deleteHelpBuy", helpBuyApi.DeleteHelpBuy) // 删除HelpBuy
		helpBuyRouter.DELETE("deleteHelpBuyByIds", helpBuyApi.DeleteHelpBuyByIds) // 批量删除HelpBuy
		helpBuyRouter.PUT("updateHelpBuy", helpBuyApi.UpdateHelpBuy)    // 更新HelpBuy
	}
	{
		helpBuyRouterWithoutRecord.GET("findHelpBuy", helpBuyApi.FindHelpBuy)        // 根据ID获取HelpBuy
		helpBuyRouterWithoutRecord.GET("getHelpBuyList", helpBuyApi.GetHelpBuyList)  // 获取HelpBuy列表
	}
}
