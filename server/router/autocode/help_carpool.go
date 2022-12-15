package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type HelpCarpoolRouter struct {
}

// InitHelpCarpoolRouter 初始化 HelpCarpool 路由信息
func (s *HelpCarpoolRouter) InitHelpCarpoolRouter(Router *gin.RouterGroup) {
	helpCarpoolRouter := Router.Group("helpCarpool").Use(middleware.OperationRecord())
	helpCarpoolRouterWithoutRecord := Router.Group("helpCarpool")
	var helpCarpoolApi = v1.ApiGroupApp.AutoCodeApiGroup.HelpCarpoolApi
	{
		helpCarpoolRouter.POST("createHelpCarpool", helpCarpoolApi.CreateHelpCarpool)   // 新建HelpCarpool
		helpCarpoolRouter.DELETE("deleteHelpCarpool", helpCarpoolApi.DeleteHelpCarpool) // 删除HelpCarpool
		helpCarpoolRouter.DELETE("deleteHelpCarpoolByIds", helpCarpoolApi.DeleteHelpCarpoolByIds) // 批量删除HelpCarpool
		helpCarpoolRouter.PUT("updateHelpCarpool", helpCarpoolApi.UpdateHelpCarpool)    // 更新HelpCarpool
	}
	{
		helpCarpoolRouterWithoutRecord.GET("findHelpCarpool", helpCarpoolApi.FindHelpCarpool)        // 根据ID获取HelpCarpool
		helpCarpoolRouterWithoutRecord.GET("getHelpCarpoolList", helpCarpoolApi.GetHelpCarpoolList)  // 获取HelpCarpool列表
	}
}
