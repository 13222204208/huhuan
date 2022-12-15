package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type HelpIncidentallyRouter struct {
}

// InitHelpIncidentallyRouter 初始化 HelpIncidentally 路由信息
func (s *HelpIncidentallyRouter) InitHelpIncidentallyRouter(Router *gin.RouterGroup) {
	helpIncidentallyRouter := Router.Group("helpIncidentally").Use(middleware.OperationRecord())
	helpIncidentallyRouterWithoutRecord := Router.Group("helpIncidentally")
	var helpIncidentallyApi = v1.ApiGroupApp.AutoCodeApiGroup.HelpIncidentallyApi
	{
		helpIncidentallyRouter.POST("createHelpIncidentally", helpIncidentallyApi.CreateHelpIncidentally)   // 新建HelpIncidentally
		helpIncidentallyRouter.DELETE("deleteHelpIncidentally", helpIncidentallyApi.DeleteHelpIncidentally) // 删除HelpIncidentally
		helpIncidentallyRouter.DELETE("deleteHelpIncidentallyByIds", helpIncidentallyApi.DeleteHelpIncidentallyByIds) // 批量删除HelpIncidentally
		helpIncidentallyRouter.PUT("updateHelpIncidentally", helpIncidentallyApi.UpdateHelpIncidentally)    // 更新HelpIncidentally
	}
	{
		helpIncidentallyRouterWithoutRecord.GET("findHelpIncidentally", helpIncidentallyApi.FindHelpIncidentally)        // 根据ID获取HelpIncidentally
		helpIncidentallyRouterWithoutRecord.GET("getHelpIncidentallyList", helpIncidentallyApi.GetHelpIncidentallyList)  // 获取HelpIncidentally列表
	}
}
