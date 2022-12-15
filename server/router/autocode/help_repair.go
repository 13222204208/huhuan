package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type HelpRepairRouter struct {
}

// InitHelpRepairRouter 初始化 HelpRepair 路由信息
func (s *HelpRepairRouter) InitHelpRepairRouter(Router *gin.RouterGroup) {
	helpRepairRouter := Router.Group("helpRepair").Use(middleware.OperationRecord())
	helpRepairRouterWithoutRecord := Router.Group("helpRepair")
	var helpRepairApi = v1.ApiGroupApp.AutoCodeApiGroup.HelpRepairApi
	{
		helpRepairRouter.POST("createHelpRepair", helpRepairApi.CreateHelpRepair)   // 新建HelpRepair
		helpRepairRouter.DELETE("deleteHelpRepair", helpRepairApi.DeleteHelpRepair) // 删除HelpRepair
		helpRepairRouter.DELETE("deleteHelpRepairByIds", helpRepairApi.DeleteHelpRepairByIds) // 批量删除HelpRepair
		helpRepairRouter.PUT("updateHelpRepair", helpRepairApi.UpdateHelpRepair)    // 更新HelpRepair
	}
	{
		helpRepairRouterWithoutRecord.GET("findHelpRepair", helpRepairApi.FindHelpRepair)        // 根据ID获取HelpRepair
		helpRepairRouterWithoutRecord.GET("getHelpRepairList", helpRepairApi.GetHelpRepairList)  // 获取HelpRepair列表
	}
}
