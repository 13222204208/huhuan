package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type HelpSecondRouter struct {
}

// InitHelpSecondRouter 初始化 HelpSecond 路由信息
func (s *HelpSecondRouter) InitHelpSecondRouter(Router *gin.RouterGroup) {
	helpSecondRouter := Router.Group("helpSecond").Use(middleware.OperationRecord())
	helpSecondRouterWithoutRecord := Router.Group("helpSecond")
	var helpSecondApi = v1.ApiGroupApp.AutoCodeApiGroup.HelpSecondApi
	{
		helpSecondRouter.POST("createHelpSecond", helpSecondApi.CreateHelpSecond)   // 新建HelpSecond
		helpSecondRouter.DELETE("deleteHelpSecond", helpSecondApi.DeleteHelpSecond) // 删除HelpSecond
		helpSecondRouter.DELETE("deleteHelpSecondByIds", helpSecondApi.DeleteHelpSecondByIds) // 批量删除HelpSecond
		helpSecondRouter.PUT("updateHelpSecond", helpSecondApi.UpdateHelpSecond)    // 更新HelpSecond
	}
	{
		helpSecondRouterWithoutRecord.GET("findHelpSecond", helpSecondApi.FindHelpSecond)        // 根据ID获取HelpSecond
		helpSecondRouterWithoutRecord.GET("getHelpSecondList", helpSecondApi.GetHelpSecondList)  // 获取HelpSecond列表
	}
}
