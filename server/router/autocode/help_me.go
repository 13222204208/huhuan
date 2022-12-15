package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type HelpMeRouter struct {
}

// InitHelpMeRouter 初始化 HelpMe 路由信息
func (s *HelpMeRouter) InitHelpMeRouter(Router *gin.RouterGroup) {
	helpMeRouter := Router.Group("helpMe").Use(middleware.OperationRecord())
	helpMeRouterWithoutRecord := Router.Group("helpMe")
	var helpMeApi = v1.ApiGroupApp.AutoCodeApiGroup.HelpMeApi
	{
		helpMeRouter.POST("createHelpMe", helpMeApi.CreateHelpMe)             // 新建HelpMe
		helpMeRouter.DELETE("deleteHelpMe", helpMeApi.DeleteHelpMe)           // 删除HelpMe
		helpMeRouter.DELETE("deleteHelpMeByIds", helpMeApi.DeleteHelpMeByIds) // 批量删除HelpMe
		helpMeRouter.PUT("updateHelpMe", helpMeApi.UpdateHelpMe)              // 更新HelpMe
	}
	{
		helpMeRouterWithoutRecord.GET("findHelpMe", helpMeApi.FindHelpMe)       // 根据ID获取HelpMe
		helpMeRouterWithoutRecord.GET("getHelpMeList", helpMeApi.GetHelpMeList) // 获取HelpMe列表

		helpMeRouterWithoutRecord.GET("getHelpMeTitle", helpMeApi.GetHelpMeTitle) //获取helpMe标题及id
	}
}
